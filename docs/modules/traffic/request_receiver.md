# 广告请求接收模块设计文档

## 1. 模块概述

广告请求接收模块是DSP系统的入口模块，负责接收和处理来自广告交易平台的RTB请求。该模块需要处理高并发的请求，确保请求的可靠性和实时性。

## 2. 功能列表

### 2.1 核心功能
1. OpenRTB协议支持
2. 请求验证和过滤
3. 请求解析和处理
4. 请求分发和负载均衡
5. 异常处理和重试机制

### 2.2 扩展功能
1. 协议版本兼容
2. 自定义扩展字段
3. 请求日志记录
4. 性能监控
5. 限流控制

## 3. 详细设计

### 3.1 系统架构
```
                    ┌─────────────┐
                    │  负载均衡层  │
                    └──────┬──────┘
                           │
                    ┌──────▼──────┐
                    │  接入层服务  │
                    └──────┬──────┘
                           │
┌─────────────┐    ┌──────▼──────┐    ┌─────────────┐
│  请求验证器  │◄───┤  请求处理器  ├───►│  请求分发器  │
└─────────────┘    └──────┬──────┘    └─────────────┘
                          │
                    ┌─────▼─────┐
                    │ 日志记录器 │
                    └───────────┘
```

### 3.2 组件说明

#### 3.2.1 负载均衡层
- 使用 Nginx/LVS 实现负载均衡
- 支持多种负载均衡策略
- 实现健康检查
- 支持 SSL 终止

#### 3.2.2 接入层服务
- 基于 Go 语言实现
- 使用 Gin 框架处理 HTTP 请求
- 实现请求限流
- 支持优雅关闭

#### 3.2.3 请求验证器
- 实现请求签名验证
- 时间戳校验
- IP 白名单检查
- 请求频率限制

#### 3.2.4 请求处理器
- OpenRTB 协议解析
- 请求数据验证
- 异常处理
- 请求转换

#### 3.2.5 请求分发器
- 请求优先级队列
- 负载均衡策略
- 超时控制
- 失败重试

#### 3.2.6 日志记录器
- 请求日志记录
- 性能指标收集
- 异常日志记录
- 审计日志

## 4. 接口定义

### 4.1 HTTP 接口
```go
// 广告请求接口
POST /api/v1/bid
Content-Type: application/json

// 请求体
{
    "id": "bid_request_id",
    "imp": [...],
    "app": {...},
    "device": {...},
    "user": {...}
}

// 响应体
{
    "id": "bid_request_id",
    "bidid": "bid_id",
    "seatbid": [...]
}
```

### 4.2 内部接口
```go
// 请求验证接口
type RequestValidator interface {
    ValidateRequest(ctx context.Context, req *BidRequest) error
}

// 请求处理接口
type RequestProcessor interface {
    ProcessRequest(ctx context.Context, req *BidRequest) (*BidResponse, error)
}

// 请求分发接口
type RequestDispatcher interface {
    DispatchRequest(ctx context.Context, req *BidRequest) error
}
```

## 5. 数据结构

### 5.1 请求结构
```go
// BidRequest 竞价请求
type BidRequest struct {
    ID      string    `json:"id"`
    Imp     []Imp     `json:"imp"`
    App     *App      `json:"app,omitempty"`
    Device  *Device   `json:"device,omitempty"`
    User    *User     `json:"user,omitempty"`
    Context *Context  `json:"context,omitempty"`
}

// Imp 展示对象
type Imp struct {
    ID       string   `json:"id"`
    Banner   *Banner  `json:"banner,omitempty"`
    Video    *Video   `json:"video,omitempty"`
    Native   *Native  `json:"native,omitempty"`
    BidFloor float64  `json:"bidfloor"`
}

// 其他相关结构体定义...
```

### 5.2 响应结构
```go
// BidResponse 竞价响应
type BidResponse struct {
    ID      string     `json:"id"`
    BidID   string     `json:"bidid"`
    SeatBid []SeatBid  `json:"seatbid"`
}

// SeatBid 席位出价
type SeatBid struct {
    Bid []Bid `json:"bid"`
}

// Bid 出价信息
type Bid struct {
    ID      string  `json:"id"`
    ImpID   string  `json:"impid"`
    Price   float64 `json:"price"`
    AdM     string  `json:"adm"`
    AdID    string  `json:"adid"`
    NURL    string  `json:"nurl,omitempty"`
}
```

## 6. 实现路径

### 6.1 基础框架搭建
1. 项目初始化
   ```bash
   # 创建项目目录
   mkdir -p cmd/request_receiver
   mkdir -p internal/{handler,service,model,middleware}
   mkdir -p pkg/{validator,parser,dispatcher}
   
   # 初始化 Go 模块
   go mod init min-dsp
   
   # 添加依赖
   go get -u github.com/gin-gonic/gin
   go get -u github.com/go-playground/validator/v10
   ```

2. 配置文件
   ```yaml
   # config/request_receiver.yaml
   server:
     port: 8080
     read_timeout: 5s
     write_timeout: 5s
   
   validator:
     max_request_size: 1048576  # 1MB
     request_timeout: 100ms
     ip_whitelist:
       - 192.168.1.0/24
   
   dispatcher:
     worker_pool_size: 1000
     queue_size: 10000
     retry_times: 3
     retry_interval: 100ms
   ```

### 6.2 核心功能实现

#### 6.2.1 请求验证器实现
```go
// pkg/validator/request_validator.go
type RequestValidator struct {
    config *Config
    cache  *Cache
}

func (v *RequestValidator) ValidateRequest(ctx context.Context, req *BidRequest) error {
    // 1. 基础验证
    if err := v.validateBasic(req); err != nil {
        return err
    }
    
    // 2. 签名验证
    if err := v.validateSignature(req); err != nil {
        return err
    }
    
    // 3. 时间戳验证
    if err := v.validateTimestamp(req); err != nil {
        return err
    }
    
    // 4. IP验证
    if err := v.validateIP(ctx, req); err != nil {
        return err
    }
    
    // 5. 频率限制
    if err := v.validateRateLimit(ctx, req); err != nil {
        return err
    }
    
    return nil
}
```

#### 6.2.2 请求处理器实现
```go
// pkg/parser/request_processor.go
type RequestProcessor struct {
    validator *validator.RequestValidator
    parser    *OpenRTBParser
}

func (p *RequestProcessor) ProcessRequest(ctx context.Context, req *BidRequest) (*BidResponse, error) {
    // 1. 请求验证
    if err := p.validator.ValidateRequest(ctx, req); err != nil {
        return nil, err
    }
    
    // 2. 请求解析
    bidRequest, err := p.parser.ParseRequest(req)
    if err != nil {
        return nil, err
    }
    
    // 3. 请求处理
    bidResponse, err := p.processBidRequest(ctx, bidRequest)
    if err != nil {
        return nil, err
    }
    
    // 4. 响应生成
    return p.parser.GenerateResponse(bidResponse), nil
}
```

#### 6.2.3 请求分发器实现
```go
// pkg/dispatcher/request_dispatcher.go
type RequestDispatcher struct {
    config     *Config
    workerPool *WorkerPool
    queue      *PriorityQueue
}

func (d *RequestDispatcher) DispatchRequest(ctx context.Context, req *BidRequest) error {
    // 1. 请求入队
    task := &Task{
        Request: req,
        Priority: d.calculatePriority(req),
        Timeout:  d.config.RequestTimeout,
    }
    
    if err := d.queue.Push(task); err != nil {
        return err
    }
    
    // 2. 等待处理
    select {
    case result := <-task.Result:
        return result.Error
    case <-ctx.Done():
        return ctx.Err()
    case <-time.After(d.config.RequestTimeout):
        return ErrRequestTimeout
    }
}
```

### 6.3 中间件实现

#### 6.3.1 限流中间件
```go
// internal/middleware/ratelimit.go
func RateLimit(cfg *Config) gin.HandlerFunc {
    limiter := rate.NewLimiter(rate.Limit(cfg.RateLimit), cfg.BurstLimit)
    
    return func(c *gin.Context) {
        if !limiter.Allow() {
            c.JSON(http.StatusTooManyRequests, gin.H{
                "error": "rate limit exceeded",
            })
            c.Abort()
            return
        }
        c.Next()
    }
}
```

#### 6.3.2 日志中间件
```go
// internal/middleware/logger.go
func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        path := c.Request.URL.Path
        
        c.Next()
        
        latency := time.Since(start)
        status := c.Writer.Status()
        
        logger.Info("request processed",
            zap.String("path", path),
            zap.Int("status", status),
            zap.Duration("latency", latency),
        )
    }
}
```

## 7. 性能考虑

### 7.1 高并发处理
1. 使用连接池
2. 实现请求队列
3. 使用工作池
4. 实现请求合并

### 7.2 内存优化
1. 对象池复用
2. 内存预分配
3. 大对象处理
4. GC优化

### 7.3 性能监控
1. 请求延迟监控
2. 内存使用监控
3. GC监控
4. 系统资源监控

## 8. 安全考虑

### 8.1 请求安全
1. 请求签名验证
2. 时间戳校验
3. IP白名单
4. 请求频率限制

### 8.2 数据安全
1. 敏感数据加密
2. 数据脱敏
3. 日志脱敏
4. 数据备份

## 9. 测试方案

### 9.1 单元测试
```go
// pkg/validator/request_validator_test.go
func TestRequestValidator_ValidateRequest(t *testing.T) {
    tests := []struct {
        name    string
        req     *BidRequest
        wantErr bool
    }{
        // 测试用例
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            v := NewRequestValidator(config)
            err := v.ValidateRequest(context.Background(), tt.req)
            if (err != nil) != tt.wantErr {
                t.Errorf("ValidateRequest() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
```

### 9.2 性能测试
```go
// pkg/dispatcher/request_dispatcher_benchmark_test.go
func BenchmarkRequestDispatcher_DispatchRequest(b *testing.B) {
    dispatcher := NewRequestDispatcher(config)
    req := createTestRequest()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        err := dispatcher.DispatchRequest(context.Background(), req)
        if err != nil {
            b.Fatal(err)
        }
    }
}
```

### 9.3 压力测试
1. 使用 wrk 进行压力测试
2. 使用 JMeter 进行负载测试
3. 使用自定义工具进行长期稳定性测试

## 10. 部署方案

### 10.1 容器化部署
```dockerfile
# Dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o request_receiver ./cmd/request_receiver

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/request_receiver .
COPY --from=builder /app/config ./config
EXPOSE 8080
CMD ["./request_receiver"]
```

### 10.2 Kubernetes部署
```yaml
# k8s/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: request-receiver
spec:
  replicas: 3
  selector:
    matchLabels:
      app: request-receiver
  template:
    metadata:
      labels:
        app: request-receiver
    spec:
      containers:
      - name: request-receiver
        image: request-receiver:latest
        ports:
        - containerPort: 8080
        resources:
          requests:
            memory: "512Mi"
            cpu: "500m"
          limits:
            memory: "1Gi"
            cpu: "1000m"
```

### 10.3 监控配置
```yaml
# prometheus/service-monitor.yaml
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: request-receiver
spec:
  selector:
    matchLabels:
      app: request-receiver
  endpoints:
  - port: metrics
    interval: 15s
```

## 11. 运维支持

### 11.1 日志收集
1. 使用 ELK 收集日志
2. 实现结构化日志
3. 日志分级处理
4. 日志轮转

### 11.2 监控告警
1. 使用 Prometheus 收集指标
2. 使用 Grafana 展示监控
3. 配置告警规则
4. 实现告警通知

### 11.3 运维工具
1. 健康检查接口
2. 优雅关闭
3. 配置热更新
4. 性能分析工具 