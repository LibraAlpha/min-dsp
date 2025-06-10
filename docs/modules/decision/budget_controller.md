# 预算控制模块设计文档

## 1. 模块概述

预算控制模块是DSP系统决策引擎层的核心组件之一，负责管理和控制广告主的预算使用。该模块需要确保广告主的预算在指定时间内合理分配和使用，同时保证预算使用的准确性和实时性。在高并发场景下，需要保证预算控制的原子性和一致性。

## 2. 功能列表

### 2.1 核心功能
1. 预算分配管理
2. 实时预算控制
3. 预算使用统计
4. 预算告警管理
5. 预算调整控制

### 2.2 扩展功能
1. 智能预算分配
2. 预算使用预测
3. 预算优化建议
4. 预算使用分析
5. 预算异常检测

## 3. 详细设计

### 3.1 系统架构
```
                    ┌─────────────┐
                    │  预算管理器  │
                    └──────┬──────┘
                           │
        ┌─────────┬────────┴───────┬─────────┐
        │         │                │         │
┌───────▼───┐ ┌───▼──────┐  ┌──────▼────┐ ┌──▼──────┐
│ 预算分配器 │ │ 预算控制器 │  │ 预算统计器 │ │ 告警管理器 │
└───────┬───┘ └─────┬────┘  └──────┬────┘ └────┬───┘
        │           │              │           │
        └───────────┼──────────────┼───────────┘
                    │              │
              ┌─────▼──────────────▼─────┐
              │      预算存储层          │
              └─────────────┬────────────┘
                            │
                    ┌───────▼───────┐
                    │  缓存层       │
                    └───────────────┘
```

### 3.2 组件说明

#### 3.2.1 预算管理器
- 预算策略管理
- 预算规则配置
- 预算状态监控
- 预算操作审计

#### 3.2.2 预算分配器
- 预算初始化
- 预算分配策略
- 预算调整
- 预算回收

#### 3.2.3 预算控制器
- 预算检查
- 预算预留
- 预算扣减
- 预算释放

#### 3.2.4 预算统计器
- 实时统计
- 历史统计
- 趋势分析
- 报表生成

#### 3.2.5 告警管理器
- 告警规则配置
- 告警触发
- 告警通知
- 告警处理

## 4. 接口定义

### 4.1 外部接口
```go
// 预算控制接口
type BudgetController interface {
    // 检查预算
    CheckBudget(ctx context.Context, req *BudgetCheckRequest) (*BudgetCheckResponse, error)
    
    // 预留预算
    ReserveBudget(ctx context.Context, req *BudgetReserveRequest) (*BudgetReserveResponse, error)
    
    // 扣减预算
    DeductBudget(ctx context.Context, req *BudgetDeductRequest) (*BudgetDeductResponse, error)
    
    // 释放预算
    ReleaseBudget(ctx context.Context, req *BudgetReleaseRequest) (*BudgetReleaseResponse, error)
    
    // 获取预算统计
    GetBudgetStats(ctx context.Context, req *BudgetStatsRequest) (*BudgetStatsResponse, error)
}

// 预算检查请求
type BudgetCheckRequest struct {
    AdvertiserID string            `json:"advertiser_id"`
    CampaignID   string            `json:"campaign_id"`
    Amount       float64           `json:"amount"`
    Timestamp    time.Time         `json:"timestamp"`
}

// 预算检查响应
type BudgetCheckResponse struct {
    Available    bool              `json:"available"`
    Remaining    float64           `json:"remaining"`
    Reserved     float64           `json:"reserved"`
    Used         float64           `json:"used"`
    Timestamp    time.Time         `json:"timestamp"`
}

// 预算预留请求
type BudgetReserveRequest struct {
    AdvertiserID string            `json:"advertiser_id"`
    CampaignID   string            `json:"campaign_id"`
    Amount       float64           `json:"amount"`
    ExpireTime   time.Time         `json:"expire_time"`
    Timestamp    time.Time         `json:"timestamp"`
}

// 预算预留响应
type BudgetReserveResponse struct {
    Success      bool              `json:"success"`
    ReserveID    string            `json:"reserve_id"`
    ExpireTime   time.Time         `json:"expire_time"`
    Timestamp    time.Time         `json:"timestamp"`
}
```

### 4.2 内部接口
```go
// 预算分配接口
type BudgetAllocator interface {
    AllocateBudget(ctx context.Context, req *AllocateRequest) (*AllocateResponse, error)
    AdjustBudget(ctx context.Context, req *AdjustRequest) (*AdjustResponse, error)
}

// 预算统计接口
type BudgetStatsCollector interface {
    CollectStats(ctx context.Context, req *CollectRequest) (*CollectResponse, error)
    GetStats(ctx context.Context, req *StatsRequest) (*StatsResponse, error)
}

// 告警管理接口
type AlertManager interface {
    CheckAlert(ctx context.Context, stats *BudgetStats) ([]*Alert, error)
    SendAlert(ctx context.Context, alert *Alert) error
}
```

## 5. 数据结构

### 5.1 预算数据
```go
// 预算配置
type BudgetConfig struct {
    AdvertiserID string            `json:"advertiser_id"`
    CampaignID   string            `json:"campaign_id"`
    DailyBudget  float64           `json:"daily_budget"`
    TotalBudget  float64           `json:"total_budget"`
    StartTime    time.Time         `json:"start_time"`
    EndTime      time.Time         `json:"end_time"`
    TimeZone     string            `json:"time_zone"`
    Pacing       *PacingConfig     `json:"pacing"`
}

// 预算状态
type BudgetStatus struct {
    AdvertiserID string            `json:"advertiser_id"`
    CampaignID   string            `json:"campaign_id"`
    TotalBudget  float64           `json:"total_budget"`
    UsedBudget   float64           `json:"used_budget"`
    Reserved     float64           `json:"reserved"`
    Available    float64           `json:"available"`
    Status       BudgetState       `json:"status"`
    UpdatedAt    time.Time         `json:"updated_at"`
}

// 预算使用记录
type BudgetUsage struct {
    ID           string            `json:"id"`
    AdvertiserID string            `json:"advertiser_id"`
    CampaignID   string            `json:"campaign_id"`
    Amount       float64           `json:"amount"`
    Type         UsageType         `json:"type"`
    Status       UsageStatus       `json:"status"`
    Timestamp    time.Time         `json:"timestamp"`
}
```

### 5.2 告警数据
```go
// 告警规则
type AlertRule struct {
    ID           string            `json:"id"`
    Name         string            `json:"name"`
    Type         AlertType         `json:"type"`
    Threshold    float64           `json:"threshold"`
    Condition    AlertCondition    `json:"condition"`
    Action       AlertAction       `json:"action"`
    Enabled      bool              `json:"enabled"`
}

// 告警记录
type Alert struct {
    ID           string            `json:"id"`
    RuleID       string            `json:"rule_id"`
    Type         AlertType         `json:"type"`
    Level        AlertLevel        `json:"level"`
    Content      string            `json:"content"`
    Status       AlertStatus       `json:"status"`
    CreatedAt    time.Time         `json:"created_at"`
    UpdatedAt    time.Time         `json:"updated_at"`
}
```

## 6. 实现路径

### 6.1 基础框架搭建
1. 项目初始化
   ```bash
   # 创建项目目录
   mkdir -p cmd/budget_controller
   mkdir -p internal/{controller,allocator,stats,alert}
   mkdir -p pkg/{models,utils,metrics}
   
   # 初始化 Go 模块
   go mod init min-dsp
   
   # 添加依赖
   go get -u github.com/go-redis/redis/v8
   go get -u github.com/go-redis/redis_rate/v9
   ```

2. 配置文件
   ```yaml
   # config/budget_controller.yaml
   controller:
     enabled: true
     mode: strict
     cache_size: 10000
     update_interval: 1s
   
   allocator:
     strategy: "proportional"
     update_interval: 5m
     min_alloc: 0.01
   
   stats:
     redis:
       addr: "localhost:6379"
       db: 0
       pool_size: 100
     update_interval: 1s
     retention_days: 30
   
   alert:
     rules_path: "config/rules/alert"
     update_interval: 1m
     notification:
       email:
         enabled: true
         smtp_host: "smtp.example.com"
         smtp_port: 587
       webhook:
         enabled: true
         url: "http://alert.example.com/webhook"
   ```

### 6.2 核心功能实现

#### 6.2.1 预算控制器实现
```go
// internal/controller/budget_controller.go
type BudgetController struct {
    config      *Config
    allocator   BudgetAllocator
    stats       BudgetStatsCollector
    alert       AlertManager
    cache       *Cache
}

func (c *BudgetController) CheckBudget(ctx context.Context, req *BudgetCheckRequest) (*BudgetCheckResponse, error) {
    // 1. 获取预算状态
    status, err := c.getBudgetStatus(ctx, req.AdvertiserID, req.CampaignID)
    if err != nil {
        return nil, err
    }
    
    // 2. 检查预算可用性
    available := c.checkBudgetAvailable(status, req.Amount)
    
    // 3. 获取预算统计
    stats, err := c.stats.GetStats(ctx, &StatsRequest{
        AdvertiserID: req.AdvertiserID,
        CampaignID:   req.CampaignID,
    })
    if err != nil {
        return nil, err
    }
    
    return &BudgetCheckResponse{
        Available: available,
        Remaining: status.Available,
        Reserved:  status.Reserved,
        Used:      stats.Used,
        Timestamp: time.Now(),
    }, nil
}

func (c *BudgetController) ReserveBudget(ctx context.Context, req *BudgetReserveRequest) (*BudgetReserveResponse, error) {
    // 1. 检查预算
    check, err := c.CheckBudget(ctx, &BudgetCheckRequest{
        AdvertiserID: req.AdvertiserID,
        CampaignID:   req.CampaignID,
        Amount:       req.Amount,
        Timestamp:    req.Timestamp,
    })
    if err != nil {
        return nil, err
    }
    if !check.Available {
        return &BudgetReserveResponse{
            Success: false,
        }, nil
    }
    
    // 2. 预留预算
    reserveID := uuid.New().String()
    err = c.cache.Set(ctx, getReserveKey(reserveID), &BudgetReserve{
        ID:           reserveID,
        AdvertiserID: req.AdvertiserID,
        CampaignID:   req.CampaignID,
        Amount:       req.Amount,
        ExpireTime:   req.ExpireTime,
        Status:       ReserveStatusPending,
        CreatedAt:    time.Now(),
    })
    if err != nil {
        return nil, err
    }
    
    // 3. 更新预算状态
    err = c.updateBudgetStatus(ctx, req.AdvertiserID, req.CampaignID, req.Amount)
    if err != nil {
        c.cache.Del(ctx, getReserveKey(reserveID))
        return nil, err
    }
    
    return &BudgetReserveResponse{
        Success:    true,
        ReserveID:  reserveID,
        ExpireTime: req.ExpireTime,
        Timestamp:  time.Now(),
    }, nil
}
```

#### 6.2.2 预算分配器实现
```go
// internal/allocator/budget_allocator.go
type BudgetAllocator struct {
    config      *Config
    cache       *Cache
}

func (a *BudgetAllocator) AllocateBudget(ctx context.Context, req *AllocateRequest) (*AllocateResponse, error) {
    // 1. 获取预算配置
    config, err := a.getBudgetConfig(ctx, req.AdvertiserID, req.CampaignID)
    if err != nil {
        return nil, err
    }
    
    // 2. 计算分配金额
    amount := a.calculateAllocation(config, req)
    
    // 3. 更新预算状态
    err = a.updateBudgetStatus(ctx, req.AdvertiserID, req.CampaignID, amount)
    if err != nil {
        return nil, err
    }
    
    return &AllocateResponse{
        Amount:    amount,
        Timestamp: time.Now(),
    }, nil
}

func (a *BudgetAllocator) calculateAllocation(config *BudgetConfig, req *AllocateRequest) float64 {
    switch config.Pacing.Strategy {
    case PacingStrategyProportional:
        return a.calculateProportionalAllocation(config, req)
    case PacingStrategyEven:
        return a.calculateEvenAllocation(config, req)
    case PacingStrategyAggressive:
        return a.calculateAggressiveAllocation(config, req)
    default:
        return a.calculateDefaultAllocation(config, req)
    }
}
```

#### 6.2.3 预算统计器实现
```go
// internal/stats/budget_stats.go
type BudgetStatsCollector struct {
    config      *Config
    redis       *redis.Client
}

func (c *BudgetStatsCollector) CollectStats(ctx context.Context, req *CollectRequest) (*CollectResponse, error) {
    // 1. 收集统计数据
    stats := &BudgetStats{
        AdvertiserID: req.AdvertiserID,
        CampaignID:   req.CampaignID,
        Used:         req.Amount,
        Timestamp:    time.Now(),
    }
    
    // 2. 更新Redis统计
    pipe := c.redis.Pipeline()
    pipe.HIncrByFloat(ctx, getStatsKey(req.AdvertiserID, req.CampaignID), "used", req.Amount)
    pipe.Expire(ctx, getStatsKey(req.AdvertiserID, req.CampaignID), 24*time.Hour)
    _, err := pipe.Exec(ctx)
    if err != nil {
        return nil, err
    }
    
    // 3. 检查告警
    alerts, err := c.checkAlerts(ctx, stats)
    if err != nil {
        return nil, err
    }
    
    return &CollectResponse{
        Stats:  stats,
        Alerts: alerts,
    }, nil
}
```

### 6.3 中间件实现

#### 6.3.1 缓存中间件
```go
// internal/middleware/cache.go
func CacheMiddleware(cache *Cache) gin.HandlerFunc {
    return func(c *gin.Context) {
        // 1. 检查缓存
        if result, ok := cache.Get(c.Request.Context(), getRequestKey(c)); ok {
            c.JSON(http.StatusOK, result)
            c.Abort()
            return
        }
        
        c.Next()
        
        // 2. 设置缓存
        if c.Writer.Status() == http.StatusOK {
            cache.Set(c.Request.Context(), getRequestKey(c), c.Keys["result"])
        }
    }
}
```

#### 6.3.2 监控中间件
```go
// internal/middleware/monitor.go
func MonitorMiddleware(metrics *Metrics) gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        
        c.Next()
        
        // 记录指标
        metrics.RecordRequest(
            c.Request.URL.Path,
            c.Writer.Status(),
            time.Since(start),
        )
    }
}
```

## 7. 性能考虑

### 7.1 高并发处理
1. 使用Redis实现分布式锁
2. 实现预算预留机制
3. 使用批量处理
4. 实现异步处理

### 7.2 内存优化
1. 使用对象池
2. 实现数据分片
3. 优化数据结构
4. 控制缓存大小

### 7.3 性能监控
1. 请求延迟监控
2. 预算操作监控
3. 资源使用监控
4. 系统负载监控

## 8. 安全考虑

### 8.1 数据安全
1. 预算数据加密
2. 操作日志记录
3. 访问权限控制
4. 数据备份恢复

### 8.2 操作安全
1. 操作验证
2. 并发控制
3. 异常处理
4. 审计日志

## 9. 测试方案

### 9.1 单元测试
```go
// internal/controller/budget_controller_test.go
func TestBudgetController_CheckBudget(t *testing.T) {
    tests := []struct {
        name    string
        req     *BudgetCheckRequest
        want    *BudgetCheckResponse
        wantErr bool
    }{
        // 测试用例
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            c := NewBudgetController(config)
            got, err := c.CheckBudget(context.Background(), tt.req)
            if (err != nil) != tt.wantErr {
                t.Errorf("CheckBudget() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("CheckBudget() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

### 9.2 性能测试
```go
// internal/controller/budget_controller_benchmark_test.go
func BenchmarkBudgetController_CheckBudget(b *testing.B) {
    controller := NewBudgetController(config)
    req := createTestRequest()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, err := controller.CheckBudget(context.Background(), req)
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
RUN go build -o budget_controller ./cmd/budget_controller

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/budget_controller .
COPY --from=builder /app/config ./config
EXPOSE 8080
CMD ["./budget_controller"]
```

### 10.2 Kubernetes部署
```yaml
# k8s/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: budget-controller
spec:
  replicas: 3
  selector:
    matchLabels:
      app: budget-controller
  template:
    metadata:
      labels:
        app: budget-controller
    spec:
      containers:
      - name: budget-controller
        image: budget-controller:latest
        ports:
        - containerPort: 8080
        resources:
          requests:
            memory: "1Gi"
            cpu: "1000m"
          limits:
            memory: "2Gi"
            cpu: "2000m"
        volumeMounts:
        - name: config
          mountPath: /app/config
      volumes:
      - name: config
        configMap:
          name: budget-controller-config
```

### 10.3 监控配置
```yaml
# prometheus/service-monitor.yaml
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: budget-controller
spec:
  selector:
    matchLabels:
      app: budget-controller
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