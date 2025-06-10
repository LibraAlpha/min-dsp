# 投放管理模块设计文档

## 1. 模块概述

投放管理模块是DSP系统运营管理层的核心组件之一，负责管理广告投放的全生命周期，包括投放策略制定、投放执行、效果监控、优化调整等。该模块需要确保广告投放的精准性、高效性和可控性，同时提供完整的投放数据分析和优化建议。

## 2. 功能列表

### 2.1 核心功能
1. 投放策略管理
2. 投放执行控制
3. 投放效果监控
4. 投放优化调整
5. 投放数据分析

### 2.2 扩展功能
1. 智能投放优化
2. 实时效果分析
3. 竞品分析
4. 异常检测
5. 投放预测

## 3. 详细设计

### 3.1 系统架构
```
                    ┌─────────────┐
                    │  接入层     │
                    └──────┬──────┘
                           │
                    ┌──────▼──────┐
                    │  业务层     │
                    └──────┬──────┘
                           │
         ┌─────────┬───────┴───────┬─────────┐
         │         │               │         │
┌────────▼───┐ ┌───▼──────┐  ┌────▼────┐ ┌──▼──────┐
│ 策略管理    │ │ 执行控制  │  │ 效果监控 │ │ 优化调整 │
└───────┬────┘ └─────┬────┘  └────┬────┘ └────┬────┘
        │            │            │           │
        └────────────┼────────────┼───────────┘
                     │            │
               ┌─────▼────────────▼─────┐
               │      数据层            │
               └─────────────┬──────────┘
                             │
                     ┌───────▼───────┐
                     │  存储层       │
                     └───────────────┘
```

### 3.2 组件说明

#### 3.2.1 接入层
- API接口管理
- 请求验证
- 权限控制
- 访问日志

#### 3.2.2 业务层
- 策略管理：投放策略的创建、修改、查询、删除等
- 执行控制：投放任务的启动、暂停、恢复、停止等
- 效果监控：实时监控投放效果、异常检测、告警等
- 优化调整：基于效果数据的自动优化和手动调整

#### 3.2.3 数据层
- 数据访问
- 数据验证
- 数据转换
- 数据缓存

#### 3.2.4 存储层
- 关系型数据库
- 缓存数据库
- 时序数据库
- 日志存储

## 4. 接口定义

### 4.1 外部接口
```go
// 投放管理服务接口
type DeliveryManager interface {
    // 策略管理
    CreateStrategy(ctx context.Context, req *CreateStrategyRequest) (*StrategyResponse, error)
    UpdateStrategy(ctx context.Context, req *UpdateStrategyRequest) (*StrategyResponse, error)
    GetStrategy(ctx context.Context, req *GetStrategyRequest) (*StrategyResponse, error)
    ListStrategies(ctx context.Context, req *ListStrategiesRequest) (*ListStrategiesResponse, error)
    
    // 执行控制
    StartDelivery(ctx context.Context, req *StartDeliveryRequest) (*DeliveryResponse, error)
    PauseDelivery(ctx context.Context, req *PauseDeliveryRequest) (*DeliveryResponse, error)
    ResumeDelivery(ctx context.Context, req *ResumeDeliveryRequest) (*DeliveryResponse, error)
    StopDelivery(ctx context.Context, req *StopDeliveryRequest) (*DeliveryResponse, error)
    
    // 效果监控
    GetDeliveryMetrics(ctx context.Context, req *GetDeliveryMetricsRequest) (*MetricsResponse, error)
    GetDeliveryReport(ctx context.Context, req *GetDeliveryReportRequest) (*ReportResponse, error)
    
    // 优化调整
    OptimizeDelivery(ctx context.Context, req *OptimizeDeliveryRequest) (*OptimizeResponse, error)
    AdjustDelivery(ctx context.Context, req *AdjustDeliveryRequest) (*AdjustResponse, error)
}

// 策略请求
type CreateStrategyRequest struct {
    Name        string            `json:"name"`
    Type        StrategyType      `json:"type"`
    Advertiser  string            `json:"advertiser"`
    Campaign    string            `json:"campaign"`
    Settings    *StrategySettings `json:"settings"`
    Targeting   *TargetingRules   `json:"targeting"`
    Budget      *BudgetRules      `json:"budget"`
    Schedule    *ScheduleRules    `json:"schedule"`
    Metadata    map[string]string `json:"metadata"`
}

// 投放请求
type StartDeliveryRequest struct {
    StrategyID  string            `json:"strategy_id"`
    StartTime   time.Time         `json:"start_time"`
    EndTime     time.Time         `json:"end_time"`
    Settings    *DeliverySettings `json:"settings"`
    Metadata    map[string]string `json:"metadata"`
}

// 监控请求
type GetDeliveryMetricsRequest struct {
    StrategyID  string            `json:"strategy_id"`
    StartTime   time.Time         `json:"start_time"`
    EndTime     time.Time         `json:"end_time"`
    Metrics     []string          `json:"metrics"`
    Granularity string            `json:"granularity"`
}

// 优化请求
type OptimizeDeliveryRequest struct {
    StrategyID  string            `json:"strategy_id"`
    Type        OptimizeType      `json:"type"`
    Parameters  map[string]string `json:"parameters"`
    Constraints *OptimizeConstraints `json:"constraints"`
}
```

### 4.2 内部接口
```go
// 策略服务接口
type StrategyService interface {
    // 创建策略
    Create(ctx context.Context, strategy *Strategy) (*Strategy, error)
    
    // 更新策略
    Update(ctx context.Context, strategy *Strategy) (*Strategy, error)
    
    // 获取策略
    Get(ctx context.Context, id string) (*Strategy, error)
    
    // 查询策略
    Query(ctx context.Context, query *StrategyQuery) ([]*Strategy, error)
}

// 执行服务接口
type ExecutionService interface {
    // 启动投放
    Start(ctx context.Context, delivery *Delivery) error
    
    // 暂停投放
    Pause(ctx context.Context, id string) error
    
    // 恢复投放
    Resume(ctx context.Context, id string) error
    
    // 停止投放
    Stop(ctx context.Context, id string) error
}

// 监控服务接口
type MonitorService interface {
    // 获取指标
    GetMetrics(ctx context.Context, query *MetricsQuery) (*DeliveryMetrics, error)
    
    // 获取报告
    GetReport(ctx context.Context, query *ReportQuery) (*DeliveryReport, error)
    
    // 更新指标
    UpdateMetrics(ctx context.Context, metrics *DeliveryMetrics) error
}

// 优化服务接口
type OptimizeService interface {
    // 优化投放
    Optimize(ctx context.Context, request *OptimizeRequest) (*OptimizeResult, error)
    
    // 调整投放
    Adjust(ctx context.Context, request *AdjustRequest) (*AdjustResult, error)
    
    // 获取建议
    GetSuggestions(ctx context.Context, query *SuggestionQuery) ([]*Suggestion, error)
}
```

## 5. 数据结构

### 5.1 策略数据
```go
// 策略类型
type StrategyType string

const (
    StrategyTypeRTB    StrategyType = "rtb"
    StrategyTypePDB    StrategyType = "pdb"
    StrategyTypePMP    StrategyType = "pmp"
)

// 策略信息
type Strategy struct {
    ID          string            `json:"id"`
    Name        string            `json:"name"`
    Type        StrategyType      `json:"type"`
    Advertiser  string            `json:"advertiser"`
    Campaign    string            `json:"campaign"`
    Status      StrategyStatus    `json:"status"`
    Settings    *StrategySettings `json:"settings"`
    Targeting   *TargetingRules   `json:"targeting"`
    Budget      *BudgetRules      `json:"budget"`
    Schedule    *ScheduleRules    `json:"schedule"`
    Stats       *StrategyStats    `json:"stats"`
    CreatedAt   time.Time         `json:"created_at"`
    UpdatedAt   time.Time         `json:"updated_at"`
    Metadata    map[string]string `json:"metadata"`
}

// 策略设置
type StrategySettings struct {
    Bidding     *BiddingRules     `json:"bidding"`
    Creative    *CreativeRules    `json:"creative"`
    Frequency   *FrequencyRules   `json:"frequency"`
    Optimization *OptimizationRules `json:"optimization"`
}

// 策略统计
type StrategyStats struct {
    Impressions int64             `json:"impressions"`
    Clicks      int64             `json:"clicks"`
    Conversions int64             `json:"conversions"`
    Spend       float64           `json:"spend"`
    Revenue     float64           `json:"revenue"`
    UpdatedAt   time.Time         `json:"updated_at"`
}
```

### 5.2 投放数据
```go
// 投放状态
type DeliveryStatus string

const (
    DeliveryStatusPending  DeliveryStatus = "pending"
    DeliveryStatusRunning  DeliveryStatus = "running"
    DeliveryStatusPaused   DeliveryStatus = "paused"
    DeliveryStatusStopped  DeliveryStatus = "stopped"
    DeliveryStatusError    DeliveryStatus = "error"
)

// 投放信息
type Delivery struct {
    ID          string            `json:"id"`
    StrategyID  string            `json:"strategy_id"`
    Status      DeliveryStatus    `json:"status"`
    StartTime   time.Time         `json:"start_time"`
    EndTime     time.Time         `json:"end_time"`
    Settings    *DeliverySettings `json:"settings"`
    Progress    *DeliveryProgress `json:"progress"`
    Stats       *DeliveryStats    `json:"stats"`
    CreatedAt   time.Time         `json:"created_at"`
    UpdatedAt   time.Time         `json:"updated_at"`
}

// 投放设置
type DeliverySettings struct {
    Priority    int               `json:"priority"`
    MaxBid      float64           `json:"max_bid"`
    MinBid      float64           `json:"min_bid"`
    DailyBudget float64           `json:"daily_budget"`
    TimeRanges  []*TimeRange      `json:"time_ranges"`
    GeoRanges   []*GeoRange       `json:"geo_ranges"`
}

// 投放进度
type DeliveryProgress struct {
    Current     float64           `json:"current"`
    Target      float64           `json:"target"`
    Percentage  float64           `json:"percentage"`
    Remaining   time.Duration     `json:"remaining"`
    UpdatedAt   time.Time         `json:"updated_at"`
}
```

### 5.3 监控数据
```go
// 投放指标
type DeliveryMetrics struct {
    StrategyID  string            `json:"strategy_id"`
    TimeRange   *TimeRange        `json:"time_range"`
    Metrics     map[string]float64 `json:"metrics"`
    Dimensions  map[string]string `json:"dimensions"`
    UpdatedAt   time.Time         `json:"updated_at"`
}

// 投放报告
type DeliveryReport struct {
    StrategyID  string            `json:"strategy_id"`
    TimeRange   *TimeRange        `json:"time_range"`
    Summary     *ReportSummary    `json:"summary"`
    Details     []*ReportDetail   `json:"details"`
    Trends      []*ReportTrend    `json:"trends"`
    Issues      []*ReportIssue    `json:"issues"`
}

// 报告摘要
type ReportSummary struct {
    Impressions int64             `json:"impressions"`
    Clicks      int64             `json:"clicks"`
    CTR         float64           `json:"ctr"`
    Conversions int64             `json:"conversions"`
    CVR         float64           `json:"cvr"`
    Spend       float64           `json:"spend"`
    Revenue     float64           `json:"revenue"`
    ROI         float64           `json:"roi"`
}

// 报告详情
type ReportDetail struct {
    Dimension   string            `json:"dimension"`
    Value       string            `json:"value"`
    Metrics     map[string]float64 `json:"metrics"`
    Percentage  float64           `json:"percentage"`
    Trend       string            `json:"trend"`
}
```

### 5.4 优化数据
```go
// 优化类型
type OptimizeType string

const (
    OptimizeTypeBid      OptimizeType = "bid"
    OptimizeTypeTarget   OptimizeType = "target"
    OptimizeTypeCreative OptimizeType = "creative"
    OptimizeTypeBudget   OptimizeType = "budget"
)

// 优化结果
type OptimizeResult struct {
    StrategyID  string            `json:"strategy_id"`
    Type        OptimizeType      `json:"type"`
    Changes     []*OptimizeChange `json:"changes"`
    Impact      *OptimizeImpact   `json:"impact"`
    Confidence  float64           `json:"confidence"`
    CreatedAt   time.Time         `json:"created_at"`
}

// 优化变更
type OptimizeChange struct {
    Field       string            `json:"field"`
    OldValue    interface{}       `json:"old_value"`
    NewValue    interface{}       `json:"new_value"`
    Reason      string            `json:"reason"`
}

// 优化影响
type OptimizeImpact struct {
    Impressions float64           `json:"impressions"`
    Clicks      float64           `json:"clicks"`
    CTR         float64           `json:"ctr"`
    Spend       float64           `json:"spend"`
    Revenue     float64           `json:"revenue"`
    ROI         float64           `json:"roi"`
}
```

## 6. 实现路径

### 6.1 基础框架搭建
1. 项目初始化
   ```bash
   # 创建项目目录
   mkdir -p cmd/delivery
   mkdir -p internal/{strategy,execution,monitor,optimize}
   mkdir -p pkg/{api,service,storage,utils}
   
   # 初始化 Go 模块
   go mod init min-dsp
   
   # 添加依赖
   go get -u github.com/gin-gonic/gin
   go get -u gorm.io/gorm
   go get -u github.com/go-redis/redis/v8
   go get -u go.uber.org/zap
   ```

2. 配置文件
   ```yaml
   # config/delivery.yaml
   server:
     port: 8080
     mode: "release"
     timeout: "30s"
   
   database:
     driver: "mysql"
     host: "localhost"
     port: 3306
     database: "dsp"
     username: "root"
     password: "******"
     max_idle_conns: 10
     max_open_conns: 100
   
   redis:
     host: "localhost"
     port: 6379
     password: "******"
     db: 0
     pool_size: 100
   
   influxdb:
     url: "http://localhost:8086"
     token: "******"
     org: "dsp"
     bucket: "metrics"
   
   monitoring:
     log_level: "info"
     metrics_port: 9090
     alert_threshold: 0.9
   
   optimization:
     check_interval: "5m"
     update_interval: "1h"
     min_confidence: 0.8
     max_changes: 5
   ```

### 6.2 核心功能实现

#### 6.2.1 策略管理实现
```go
// internal/strategy/service.go
type StrategyService struct {
    db          *gorm.DB
    cache       *redis.Client
    influx      *influxdb.Client
    logger      *zap.Logger
}

func (s *StrategyService) Create(ctx context.Context, req *CreateStrategyRequest) (*Strategy, error) {
    // 1. 验证请求
    if err := s.validateCreateRequest(req); err != nil {
        return nil, err
    }
    
    // 2. 创建策略
    strategy := &Strategy{
        Name:       req.Name,
        Type:       req.Type,
        Advertiser: req.Advertiser,
        Campaign:   req.Campaign,
        Settings:   req.Settings,
        Targeting:  req.Targeting,
        Budget:     req.Budget,
        Schedule:   req.Schedule,
        Status:     StrategyStatusPending,
    }
    
    // 3. 保存策略
    if err := s.db.Create(strategy).Error; err != nil {
        return nil, err
    }
    
    // 4. 缓存策略
    if err := s.cacheStrategy(strategy); err != nil {
        s.logger.Warn("failed to cache strategy", zap.Error(err))
    }
    
    // 5. 初始化统计
    if err := s.initStrategyStats(strategy); err != nil {
        s.logger.Warn("failed to init strategy stats", zap.Error(err))
    }
    
    return strategy, nil
}

func (s *StrategyService) Update(ctx context.Context, req *UpdateStrategyRequest) (*Strategy, error) {
    // 1. 获取策略
    strategy, err := s.Get(ctx, req.ID)
    if err != nil {
        return nil, err
    }
    
    // 2. 更新策略
    if err := s.db.Model(strategy).Updates(req.Updates).Error; err != nil {
        return nil, err
    }
    
    // 3. 更新缓存
    if err := s.cacheStrategy(strategy); err != nil {
        s.logger.Warn("failed to update strategy cache", zap.Error(err))
    }
    
    return strategy, nil
}
```

#### 6.2.2 执行控制实现
```go
// internal/execution/service.go
type ExecutionService struct {
    db          *gorm.DB
    cache       *redis.Client
    monitor     MonitorService
    logger      *zap.Logger
}

func (s *ExecutionService) Start(ctx context.Context, req *StartDeliveryRequest) error {
    // 1. 获取策略
    strategy, err := s.getStrategy(ctx, req.StrategyID)
    if err != nil {
        return err
    }
    
    // 2. 创建投放
    delivery := &Delivery{
        StrategyID: req.StrategyID,
        Status:     DeliveryStatusPending,
        StartTime:  req.StartTime,
        EndTime:    req.EndTime,
        Settings:   req.Settings,
    }
    
    // 3. 保存投放
    if err := s.db.Create(delivery).Error; err != nil {
        return err
    }
    
    // 4. 启动投放
    if err := s.startDelivery(delivery); err != nil {
        delivery.Status = DeliveryStatusError
        s.db.Save(delivery)
        return err
    }
    
    // 5. 更新状态
    delivery.Status = DeliveryStatusRunning
    if err := s.db.Save(delivery).Error; err != nil {
        return err
    }
    
    // 6. 启动监控
    go s.monitorDelivery(delivery)
    
    return nil
}

func (s *ExecutionService) Pause(ctx context.Context, id string) error {
    // 1. 获取投放
    delivery, err := s.getDelivery(ctx, id)
    if err != nil {
        return err
    }
    
    // 2. 暂停投放
    if err := s.pauseDelivery(delivery); err != nil {
        return err
    }
    
    // 3. 更新状态
    delivery.Status = DeliveryStatusPaused
    if err := s.db.Save(delivery).Error; err != nil {
        return err
    }
    
    return nil
}
```

#### 6.2.3 监控服务实现
```go
// internal/monitor/service.go
type MonitorService struct {
    db          *gorm.DB
    influx      *influxdb.Client
    logger      *zap.Logger
}

func (s *MonitorService) GetMetrics(ctx context.Context, query *MetricsQuery) (*DeliveryMetrics, error) {
    // 1. 构建查询
    fluxQuery := s.buildMetricsQuery(query)
    
    // 2. 执行查询
    result, err := s.influx.Query(ctx, fluxQuery)
    if err != nil {
        return nil, err
    }
    
    // 3. 处理结果
    metrics, err := s.processMetricsResult(result)
    if err != nil {
        return nil, err
    }
    
    return metrics, nil
}

func (s *MonitorService) GetReport(ctx context.Context, query *ReportQuery) (*DeliveryReport, error) {
    // 1. 获取指标
    metrics, err := s.GetMetrics(ctx, query.MetricsQuery)
    if err != nil {
        return nil, err
    }
    
    // 2. 生成摘要
    summary, err := s.generateSummary(metrics)
    if err != nil {
        return nil, err
    }
    
    // 3. 生成详情
    details, err := s.generateDetails(metrics)
    if err != nil {
        return nil, err
    }
    
    // 4. 分析趋势
    trends, err := s.analyzeTrends(metrics)
    if err != nil {
        return nil, err
    }
    
    // 5. 检测问题
    issues, err := s.detectIssues(metrics)
    if err != nil {
        return nil, err
    }
    
    return &DeliveryReport{
        StrategyID: query.StrategyID,
        TimeRange:  query.TimeRange,
        Summary:    summary,
        Details:    details,
        Trends:     trends,
        Issues:     issues,
    }, nil
}
```

### 6.3 中间件实现

#### 6.3.1 认证中间件
```go
// internal/middleware/auth.go
func AuthMiddleware(jwt *JWT) gin.HandlerFunc {
    return func(c *gin.Context) {
        // 1. 获取令牌
        token := c.GetHeader("Authorization")
        if token == "" {
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }
        
        // 2. 验证令牌
        claims, err := jwt.ValidateToken(token)
        if err != nil {
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }
        
        // 3. 设置上下文
        c.Set("user_id", claims.UserID)
        c.Set("user_role", claims.Role)
        
        c.Next()
    }
}
```

#### 6.3.2 日志中间件
```go
// internal/middleware/logger.go
func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        path := c.Request.URL.Path
        method := c.Request.Method
        
        c.Next()
        
        // 记录访问日志
        logger.Info("api access",
            zap.String("path", path),
            zap.String("method", method),
            zap.Int("status", c.Writer.Status()),
            zap.Duration("latency", time.Since(start)),
            zap.String("ip", c.ClientIP()),
            zap.String("user_id", c.GetString("user_id")),
        )
    }
}
```

## 7. 性能考虑

### 7.1 数据库优化
1. 索引优化
2. 查询优化
3. 连接池管理
4. 分库分表

### 7.2 缓存优化
1. 多级缓存
2. 缓存预热
3. 缓存更新
4. 缓存清理

### 7.3 接口优化
1. 接口限流
2. 数据压缩
3. 批量处理
4. 异步处理

## 8. 安全考虑

### 8.1 数据安全
1. 数据加密
2. 访问控制
3. 数据备份
4. 审计日志

### 8.2 运行安全
1. 身份认证
2. 权限管理
3. 操作审计
4. 安全防护

## 9. 测试方案

### 9.1 单元测试
```go
// internal/strategy/service_test.go
func TestStrategyService_Create(t *testing.T) {
    tests := []struct {
        name    string
        req     *CreateStrategyRequest
        want    *Strategy
        wantErr bool
    }{
        // 测试用例
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            s := NewStrategyService(db, cache, influx, logger)
            got, err := s.Create(context.Background(), tt.req)
            if (err != nil) != tt.wantErr {
                t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("Create() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

### 9.2 性能测试
```go
// internal/strategy/service_benchmark_test.go
func BenchmarkStrategyService_Create(b *testing.B) {
    s := NewStrategyService(db, cache, influx, logger)
    req := createTestRequest()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, err := s.Create(context.Background(), req)
        if err != nil {
            b.Fatal(err)
        }
    }
}
```

### 9.3 集成测试
1. 接口测试
2. 功能测试
3. 性能测试
4. 安全测试

## 10. 部署方案

### 10.1 容器化部署
```dockerfile
# Dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o delivery ./cmd/delivery

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/delivery .
COPY --from=builder /app/config ./config
EXPOSE 8080
CMD ["./delivery"]
```

### 10.2 Kubernetes部署
```yaml
# k8s/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: delivery
spec:
  replicas: 3
  selector:
    matchLabels:
      app: delivery
  template:
    metadata:
      labels:
        app: delivery
    spec:
      containers:
      - name: delivery
        image: delivery:latest
        ports:
        - containerPort: 8080
        resources:
          requests:
            memory: "2Gi"
            cpu: "1000m"
          limits:
            memory: "4Gi"
            cpu: "2000m"
        volumeMounts:
        - name: config
          mountPath: /app/config
      volumes:
      - name: config
        configMap:
          name: delivery-config
```

### 10.3 监控配置
```yaml
# prometheus/service-monitor.yaml
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: delivery
spec:
  selector:
    matchLabels:
      app: delivery
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
1. 策略管理工具
2. 投放控制工具
3. 效果分析工具
4. 优化建议工具 