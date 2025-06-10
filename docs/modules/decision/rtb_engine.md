# 实时竞价模块设计文档

## 1. 模块概述

实时竞价模块是DSP系统决策引擎层的核心组件，负责根据广告请求和广告主策略进行实时竞价决策。该模块需要在高并发、低延迟的环境下，综合考虑广告主预算、定向条件、出价策略等多个因素，做出最优的竞价决策。

## 2. 功能列表

### 2.1 核心功能
1. 实时竞价决策
2. 出价策略执行
3. 预算控制
4. 定向匹配
5. 竞价优化

### 2.2 扩展功能
1. 智能出价
2. 实时调价
3. 竞争分析
4. 效果预测
5. 策略优化

## 3. 详细设计

### 3.1 系统架构
```
                    ┌─────────────┐
                    │  请求接收器  │
                    └──────┬──────┘
                           │
                    ┌──────▼──────┐
                    │  策略匹配器  │
                    └──────┬──────┘
                           │
        ┌─────────┬────────┴───────┬─────────┐
        │         │                │         │
┌───────▼───┐ ┌───▼──────┐  ┌──────▼────┐ ┌──▼──────┐
│ 预算检查器 │ │ 定向匹配器 │  │ 出价计算器 │ │ 优化器   │
└───────┬───┘ └─────┬────┘  └──────┬────┘ └────┬───┘
        │           │              │           │
        └───────────┼──────────────┼───────────┘
                    │              │
              ┌─────▼──────────────▼─────┐
              │      竞价决策器          │
              └─────────────┬────────────┘
                            │
                    ┌───────▼───────┐
                    │  响应生成器   │
                    └───────────────┘
```

### 3.2 组件说明

#### 3.2.1 请求接收器
- 接收预处理请求
- 请求分类
- 请求分发
- 流量控制

#### 3.2.2 策略匹配器
- 策略加载
- 策略匹配
- 策略优先级
- 策略更新

#### 3.2.3 预算检查器
- 预算查询
- 预算预留
- 预算更新
- 预算告警

#### 3.2.4 定向匹配器
- 用户定向
- 设备定向
- 地域定向
- 上下文定向

#### 3.2.5 出价计算器
- 基础出价
- 策略出价
- 智能调价
- 竞争分析

#### 3.2.6 优化器
- 效果预测
- 策略优化
- 参数调优
- 实时反馈

#### 3.2.7 竞价决策器
- 综合决策
- 优先级处理
- 竞争处理
- 结果验证

#### 3.2.8 响应生成器
- 响应组装
- 数据补充
- 格式转换
- 响应发送

## 4. 接口定义

### 4.1 外部接口
```go
// 实时竞价接口
type RTBEngine interface {
    // 竞价请求
    Bid(ctx context.Context, req *BidRequest) (*BidResponse, error)
    
    // 更新策略
    UpdateStrategy(ctx context.Context, strategy *BidStrategy) error
    
    // 获取统计
    GetStats(ctx context.Context) (*BidStats, error)
}

// 竞价请求
type BidRequest struct {
    ID          string            `json:"id"`
    Imp         []Imp             `json:"imp"`
    App         *App              `json:"app,omitempty"`
    Device      *Device           `json:"device,omitempty"`
    User        *User             `json:"user,omitempty"`
    Context     *Context          `json:"context,omitempty"`
    Ext         map[string]any    `json:"ext,omitempty"`
}

// 竞价响应
type BidResponse struct {
    ID          string            `json:"id"`
    BidID       string            `json:"bidid"`
    SeatBid     []SeatBid         `json:"seatbid"`
    Ext         map[string]any    `json:"ext,omitempty"`
}

// 竞价策略
type BidStrategy struct {
    ID          string            `json:"id"`
    Name        string            `json:"name"`
    Type        StrategyType      `json:"type"`
    Priority    int               `json:"priority"`
    Budget      *Budget           `json:"budget"`
    Targeting   *Targeting        `json:"targeting"`
    Bidding     *Bidding          `json:"bidding"`
    Optimization *Optimization    `json:"optimization"`
}
```

### 4.2 内部接口
```go
// 预算检查接口
type BudgetChecker interface {
    CheckBudget(ctx context.Context, strategy *BidStrategy) (*BudgetCheck, error)
}

// 定向匹配接口
type TargetingMatcher interface {
    MatchTargeting(ctx context.Context, req *BidRequest, targeting *Targeting) (*TargetingMatch, error)
}

// 出价计算接口
type BidCalculator interface {
    CalculateBid(ctx context.Context, req *BidRequest, strategy *BidStrategy) (*BidPrice, error)
}

// 优化器接口
type Optimizer interface {
    Optimize(ctx context.Context, req *BidRequest, strategy *BidStrategy) (*Optimization, error)
}
```

## 5. 数据结构

### 5.1 策略数据
```go
// 预算配置
type Budget struct {
    Daily       float64           `json:"daily"`
    Total       float64           `json:"total"`
    StartTime   time.Time         `json:"start_time"`
    EndTime     time.Time         `json:"end_time"`
    TimeZone    string            `json:"time_zone"`
    Pacing      *Pacing           `json:"pacing"`
}

// 定向配置
type Targeting struct {
    User        *UserTargeting    `json:"user,omitempty"`
    Device      *DeviceTargeting  `json:"device,omitempty"`
    Geo         *GeoTargeting     `json:"geo,omitempty"`
    Context     *ContextTargeting `json:"context,omitempty"`
    Custom      []CustomTargeting `json:"custom,omitempty"`
}

// 出价配置
type Bidding struct {
    BasePrice   float64           `json:"base_price"`
    MaxPrice    float64           `json:"max_price"`
    MinPrice    float64           `json:"min_price"`
    Strategy    BiddingStrategy   `json:"strategy"`
    Parameters  map[string]any    `json:"parameters"`
}

// 优化配置
type Optimization struct {
    Goal        OptimizationGoal  `json:"goal"`
    Metrics     []Metric          `json:"metrics"`
    Constraints []Constraint       `json:"constraints"`
    Parameters  map[string]any    `json:"parameters"`
}
```

### 5.2 竞价数据
```go
// 竞价结果
type BidResult struct {
    RequestID   string            `json:"request_id"`
    StrategyID  string            `json:"strategy_id"`
    Price       float64           `json:"price"`
    Win         bool              `json:"win"`
    Reason      string            `json:"reason"`
    Metrics     map[string]float64 `json:"metrics"`
    Timestamp   time.Time         `json:"timestamp"`
}

// 预算检查
type BudgetCheck struct {
    StrategyID  string            `json:"strategy_id"`
    Available   bool              `json:"available"`
    Remaining   float64           `json:"remaining"`
    Reserved    float64           `json:"reserved"`
    Timestamp   time.Time         `json:"timestamp"`
}

// 定向匹配
type TargetingMatch struct {
    StrategyID  string            `json:"strategy_id"`
    Matched     bool              `json:"matched"`
    Score       float64           `json:"score"`
    Details     map[string]any    `json:"details"`
    Timestamp   time.Time         `json:"timestamp"`
}

// 出价计算
type BidPrice struct {
    StrategyID  string            `json:"strategy_id"`
    BasePrice   float64           `json:"base_price"`
    FinalPrice  float64           `json:"final_price"`
    Factors     []Factor          `json:"factors"`
    Timestamp   time.Time         `json:"timestamp"`
}
```

## 6. 实现路径

### 6.1 基础框架搭建
1. 项目初始化
   ```bash
   # 创建项目目录
   mkdir -p cmd/rtb_engine
   mkdir -p internal/{engine,strategy,budget,targeting,bidding,optimizer}
   mkdir -p pkg/{models,utils,metrics}
   
   # 初始化 Go 模块
   go mod init min-dsp
   
   # 添加依赖
   go get -u github.com/go-redis/redis/v8
   go get -u github.com/olivere/elastic/v7
   ```

2. 配置文件
   ```yaml
   # config/rtb_engine.yaml
   engine:
     enabled: true
     mode: strict
     cache_size: 10000
     update_interval: 5m
   
   strategy:
     rules_path: "config/rules/strategy"
     update_interval: 1m
     max_strategies: 1000
   
   budget:
     redis:
       addr: "localhost:6379"
       db: 0
       pool_size: 100
     update_interval: 1s
     reserve_timeout: 5s
   
   targeting:
     cache_size: 10000
     update_interval: 5m
     strict_mode: true
   
   bidding:
     model_path: "models/bidding"
     update_interval: 1h
     learning_rate: 0.01
   
   optimizer:
     metrics_path: "config/metrics"
     update_interval: 5m
     batch_size: 1000
   ```

### 6.2 核心功能实现

#### 6.2.1 实时竞价引擎实现
```go
// internal/engine/rtb_engine.go
type RTBEngine struct {
    config      *Config
    strategy    StrategyManager
    budget      BudgetChecker
    targeting   TargetingMatcher
    bidding     BidCalculator
    optimizer   Optimizer
    cache       *Cache
}

func (e *RTBEngine) Bid(ctx context.Context, req *BidRequest) (*BidResponse, error) {
    // 1. 策略匹配
    strategies, err := e.strategy.MatchStrategies(ctx, req)
    if err != nil {
        return nil, err
    }
    
    // 2. 预算检查
    availableStrategies := make([]*BidStrategy, 0)
    for _, strategy := range strategies {
        check, err := e.budget.CheckBudget(ctx, strategy)
        if err != nil {
            continue
        }
        if check.Available {
            availableStrategies = append(availableStrategies, strategy)
        }
    }
    
    // 3. 定向匹配
    matchedStrategies := make([]*BidStrategy, 0)
    for _, strategy := range availableStrategies {
        match, err := e.targeting.MatchTargeting(ctx, req, strategy.Targeting)
        if err != nil {
            continue
        }
        if match.Matched {
            matchedStrategies = append(matchedStrategies, strategy)
        }
    }
    
    // 4. 出价计算
    bids := make([]*Bid, 0)
    for _, strategy := range matchedStrategies {
        price, err := e.bidding.CalculateBid(ctx, req, strategy)
        if err != nil {
            continue
        }
        
        // 5. 优化处理
        optimization, err := e.optimizer.Optimize(ctx, req, strategy)
        if err != nil {
            continue
        }
        
        // 6. 生成出价
        bid := &Bid{
            StrategyID: strategy.ID,
            Price:     price.FinalPrice,
            AdID:      strategy.AdID,
            Creative:  strategy.Creative,
            Ext:       optimization.Parameters,
        }
        bids = append(bids, bid)
    }
    
    // 7. 响应生成
    response := e.generateResponse(req, bids)
    
    return response, nil
}
```

#### 6.2.2 策略管理器实现
```go
// internal/strategy/strategy_manager.go
type StrategyManager struct {
    strategies  map[string]*BidStrategy
    config      *Config
    cache       *Cache
}

func (m *StrategyManager) MatchStrategies(ctx context.Context, req *BidRequest) ([]*BidStrategy, error) {
    // 1. 获取策略列表
    strategies := m.getActiveStrategies()
    
    // 2. 策略过滤
    filtered := make([]*BidStrategy, 0)
    for _, strategy := range strategies {
        if m.isStrategyValid(strategy) {
            filtered = append(filtered, strategy)
        }
    }
    
    // 3. 策略排序
    sort.Slice(filtered, func(i, j int) bool {
        return filtered[i].Priority > filtered[j].Priority
    })
    
    return filtered, nil
}

func (m *StrategyManager) UpdateStrategy(ctx context.Context, strategy *BidStrategy) error {
    // 1. 策略验证
    if err := m.validateStrategy(strategy); err != nil {
        return err
    }
    
    // 2. 策略更新
    m.strategies[strategy.ID] = strategy
    
    // 3. 缓存更新
    m.cache.Set(ctx, getStrategyKey(strategy.ID), strategy)
    
    return nil
}
```

#### 6.2.3 出价计算器实现
```go
// internal/bidding/bid_calculator.go
type BidCalculator struct {
    model       *Model
    config      *Config
    cache       *Cache
}

func (c *BidCalculator) CalculateBid(ctx context.Context, req *BidRequest, strategy *BidStrategy) (*BidPrice, error) {
    // 1. 基础出价
    basePrice := strategy.Bidding.BasePrice
    
    // 2. 策略调整
    adjustments := c.calculateAdjustments(req, strategy)
    
    // 3. 模型预测
    prediction, err := c.model.Predict(ctx, req, strategy)
    if err != nil {
        return nil, err
    }
    
    // 4. 价格计算
    finalPrice := c.calculateFinalPrice(basePrice, adjustments, prediction)
    
    // 5. 价格限制
    finalPrice = c.applyPriceLimits(finalPrice, strategy.Bidding)
    
    return &BidPrice{
        StrategyID: strategy.ID,
        BasePrice:  basePrice,
        FinalPrice: finalPrice,
        Factors:    adjustments,
        Timestamp:  time.Now(),
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
        if result, ok := cache.Get(c.Request.Context(), getRequestID(c)); ok {
            c.JSON(http.StatusOK, result)
            c.Abort()
            return
        }
        
        c.Next()
        
        // 2. 设置缓存
        if c.Writer.Status() == http.StatusOK {
            cache.Set(c.Request.Context(), getRequestID(c), c.Keys["result"])
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

### 8.1 策略安全
1. 策略验证
2. 策略版本控制
3. 策略备份
4. 策略审计

### 8.2 数据安全
1. 敏感数据加密
2. 数据脱敏
3. 访问控制
4. 数据备份

## 9. 测试方案

### 9.1 单元测试
```go
// internal/engine/rtb_engine_test.go
func TestRTBEngine_Bid(t *testing.T) {
    tests := []struct {
        name    string
        req     *BidRequest
        want    *BidResponse
        wantErr bool
    }{
        // 测试用例
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            e := NewRTBEngine(config)
            got, err := e.Bid(context.Background(), tt.req)
            if (err != nil) != tt.wantErr {
                t.Errorf("Bid() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("Bid() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

### 9.2 性能测试
```go
// internal/engine/rtb_engine_benchmark_test.go
func BenchmarkRTBEngine_Bid(b *testing.B) {
    engine := NewRTBEngine(config)
    req := createTestRequest()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, err := engine.Bid(context.Background(), req)
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
RUN go build -o rtb_engine ./cmd/rtb_engine

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/rtb_engine .
COPY --from=builder /app/config ./config
COPY --from=builder /app/models ./models
EXPOSE 8080
CMD ["./rtb_engine"]
```

### 10.2 Kubernetes部署
```yaml
# k8s/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: rtb-engine
spec:
  replicas: 3
  selector:
    matchLabels:
      app: rtb-engine
  template:
    metadata:
      labels:
        app: rtb-engine
    spec:
      containers:
      - name: rtb-engine
        image: rtb-engine:latest
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
        - name: models
          mountPath: /app/models
        - name: config
          mountPath: /app/config
      volumes:
      - name: models
        configMap:
          name: rtb-engine-models
      - name: config
        configMap:
          name: rtb-engine-config
```

### 10.3 监控配置
```yaml
# prometheus/service-monitor.yaml
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: rtb-engine
spec:
  selector:
    matchLabels:
      app: rtb-engine
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