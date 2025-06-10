# 广告位管理模块设计文档

## 1. 模块概述

广告位管理模块是DSP系统运营管理层的核心组件之一，负责管理广告位信息、投放规则、创意规格、流量质量等关键数据。该模块需要提供完整的广告位生命周期管理功能，包括广告位创建、审核、配置、监控等功能，同时确保广告位的可用性和投放效果。

## 2. 功能列表

### 2.1 核心功能
1. 广告位管理
2. 规格管理
3. 流量管理
4. 质量监控
5. 投放控制

### 2.2 扩展功能
1. 数据分析
2. 效果评估
3. 智能优化
4. 异常检测
5. 报表统计

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
│ 广告位管理  │ │ 规格管理  │  │ 流量管理 │ │ 质量监控 │
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
- 广告位管理：广告位的创建、修改、查询、删除等
- 规格管理：创意规格、尺寸规格、格式规格等
- 流量管理：流量统计、流量分析、流量预测等
- 质量监控：质量评估、异常检测、效果分析等

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
// 广告位管理服务接口
type AdPlacementManager interface {
    // 广告位管理
    CreatePlacement(ctx context.Context, req *CreatePlacementRequest) (*PlacementResponse, error)
    UpdatePlacement(ctx context.Context, req *UpdatePlacementRequest) (*PlacementResponse, error)
    GetPlacement(ctx context.Context, req *GetPlacementRequest) (*PlacementResponse, error)
    ListPlacements(ctx context.Context, req *ListPlacementsRequest) (*ListPlacementsResponse, error)
    
    // 规格管理
    CreateSpec(ctx context.Context, req *CreateSpecRequest) (*SpecResponse, error)
    UpdateSpec(ctx context.Context, req *UpdateSpecRequest) (*SpecResponse, error)
    GetSpec(ctx context.Context, req *GetSpecRequest) (*SpecResponse, error)
    
    // 流量管理
    GetTrafficStats(ctx context.Context, req *GetTrafficStatsRequest) (*TrafficStatsResponse, error)
    GetTrafficAnalysis(ctx context.Context, req *GetTrafficAnalysisRequest) (*TrafficAnalysisResponse, error)
    
    // 质量监控
    GetQualityMetrics(ctx context.Context, req *GetQualityMetricsRequest) (*QualityMetricsResponse, error)
    GetQualityReport(ctx context.Context, req *GetQualityReportRequest) (*QualityReportResponse, error)
}

// 广告位请求
type CreatePlacementRequest struct {
    Name        string            `json:"name"`
    Type        PlacementType     `json:"type"`
    Publisher   string            `json:"publisher"`
    Category    string            `json:"category"`
    Specs       []*SpecInfo       `json:"specs"`
    Settings    *PlacementSettings `json:"settings"`
    Metadata    map[string]string `json:"metadata"`
}

// 规格请求
type CreateSpecRequest struct {
    PlacementID string            `json:"placement_id"`
    Type        SpecType          `json:"type"`
    Format      string            `json:"format"`
    Size        *SizeInfo         `json:"size"`
    Settings    *SpecSettings     `json:"settings"`
    Metadata    map[string]string `json:"metadata"`
}

// 流量统计请求
type GetTrafficStatsRequest struct {
    PlacementID string            `json:"placement_id"`
    StartTime   time.Time         `json:"start_time"`
    EndTime     time.Time         `json:"end_time"`
    Granularity string            `json:"granularity"`
    Metrics     []string          `json:"metrics"`
}

// 质量指标请求
type GetQualityMetricsRequest struct {
    PlacementID string            `json:"placement_id"`
    StartTime   time.Time         `json:"start_time"`
    EndTime     time.Time         `json:"end_time"`
    Metrics     []string          `json:"metrics"`
}
```

### 4.2 内部接口
```go
// 广告位服务接口
type PlacementService interface {
    // 创建广告位
    Create(ctx context.Context, placement *Placement) (*Placement, error)
    
    // 更新广告位
    Update(ctx context.Context, placement *Placement) (*Placement, error)
    
    // 获取广告位
    Get(ctx context.Context, id string) (*Placement, error)
    
    // 查询广告位
    Query(ctx context.Context, query *PlacementQuery) ([]*Placement, error)
}

// 规格服务接口
type SpecService interface {
    // 创建规格
    Create(ctx context.Context, spec *Spec) (*Spec, error)
    
    // 更新规格
    Update(ctx context.Context, spec *Spec) (*Spec, error)
    
    // 获取规格
    Get(ctx context.Context, id string) (*Spec, error)
    
    // 验证规格
    Validate(ctx context.Context, spec *Spec) (*ValidationResult, error)
}

// 流量服务接口
type TrafficService interface {
    // 获取统计数据
    GetStats(ctx context.Context, query *TrafficQuery) (*TrafficStats, error)
    
    // 获取分析数据
    GetAnalysis(ctx context.Context, query *TrafficQuery) (*TrafficAnalysis, error)
    
    // 更新统计数据
    UpdateStats(ctx context.Context, stats *TrafficStats) error
}

// 质量服务接口
type QualityService interface {
    // 获取质量指标
    GetMetrics(ctx context.Context, query *QualityQuery) (*QualityMetrics, error)
    
    // 获取质量报告
    GetReport(ctx context.Context, query *QualityQuery) (*QualityReport, error)
    
    // 更新质量指标
    UpdateMetrics(ctx context.Context, metrics *QualityMetrics) error
}
```

## 5. 数据结构

### 5.1 广告位数据
```go
// 广告位类型
type PlacementType string

const (
    PlacementTypeBanner    PlacementType = "banner"
    PlacementTypeVideo     PlacementType = "video"
    PlacementTypeNative    PlacementType = "native"
    PlacementTypeInterstitial PlacementType = "interstitial"
)

// 广告位信息
type Placement struct {
    ID          string            `json:"id"`
    Name        string            `json:"name"`
    Type        PlacementType     `json:"type"`
    Publisher   string            `json:"publisher"`
    Category    string            `json:"category"`
    Status      PlacementStatus   `json:"status"`
    Specs       []*Spec           `json:"specs"`
    Settings    *PlacementSettings `json:"settings"`
    Stats       *PlacementStats   `json:"stats"`
    CreatedAt   time.Time         `json:"created_at"`
    UpdatedAt   time.Time         `json:"updated_at"`
    Metadata    map[string]string `json:"metadata"`
}

// 广告位设置
type PlacementSettings struct {
    FloorPrice  float64           `json:"floor_price"`
    MaxPrice    float64           `json:"max_price"`
    Timeout     int               `json:"timeout"`
    Frequency   *FrequencyRules   `json:"frequency"`
    Targeting   *TargetingRules   `json:"targeting"`
    Blocking    *BlockingRules    `json:"blocking"`
}

// 广告位统计
type PlacementStats struct {
    Impressions int64             `json:"impressions"`
    Clicks      int64             `json:"clicks"`
    CTR         float64           `json:"ctr"`
    Revenue     float64           `json:"revenue"`
    UpdatedAt   time.Time         `json:"updated_at"`
}
```

### 5.2 规格数据
```go
// 规格类型
type SpecType string

const (
    SpecTypeCreative SpecType = "creative"
    SpecTypeSize     SpecType = "size"
    SpecTypeFormat   SpecType = "format"
)

// 规格信息
type Spec struct {
    ID          string            `json:"id"`
    PlacementID string            `json:"placement_id"`
    Type        SpecType          `json:"type"`
    Format      string            `json:"format"`
    Size        *SizeInfo         `json:"size"`
    Status      SpecStatus        `json:"status"`
    Settings    *SpecSettings     `json:"settings"`
    CreatedAt   time.Time         `json:"created_at"`
    UpdatedAt   time.Time         `json:"updated_at"`
}

// 尺寸信息
type SizeInfo struct {
    Width       int               `json:"width"`
    Height      int               `json:"height"`
    AspectRatio string            `json:"aspect_ratio"`
    Unit        string            `json:"unit"`
}

// 规格设置
type SpecSettings struct {
    MaxSize     int64             `json:"max_size"`
    MinSize     int64             `json:"min_size"`
    Formats     []string          `json:"formats"`
    Requirements []string          `json:"requirements"`
}
```

### 5.3 流量数据
```go
// 流量统计
type TrafficStats struct {
    PlacementID string            `json:"placement_id"`
    TimeRange   *TimeRange        `json:"time_range"`
    Metrics     map[string]float64 `json:"metrics"`
    Dimensions  map[string]string `json:"dimensions"`
    UpdatedAt   time.Time         `json:"updated_at"`
}

// 流量分析
type TrafficAnalysis struct {
    PlacementID string            `json:"placement_id"`
    TimeRange   *TimeRange        `json:"time_range"`
    Trends      []*TrendData      `json:"trends"`
    Patterns    []*PatternData    `json:"patterns"`
    Predictions []*PredictionData `json:"predictions"`
}

// 趋势数据
type TrendData struct {
    Metric      string            `json:"metric"`
    Values      []float64         `json:"values"`
    Timestamps  []time.Time       `json:"timestamps"`
    Change      float64           `json:"change"`
}

// 模式数据
type PatternData struct {
    Type        string            `json:"type"`
    Pattern     string            `json:"pattern"`
    Confidence  float64           `json:"confidence"`
    Examples    []string          `json:"examples"`
}
```

### 5.4 质量数据
```go
// 质量指标
type QualityMetrics struct {
    PlacementID string            `json:"placement_id"`
    TimeRange   *TimeRange        `json:"time_range"`
    Metrics     map[string]float64 `json:"metrics"`
    Scores      map[string]float64 `json:"scores"`
    UpdatedAt   time.Time         `json:"updated_at"`
}

// 质量报告
type QualityReport struct {
    PlacementID string            `json:"placement_id"`
    TimeRange   *TimeRange        `json:"time_range"`
    Summary     *QualitySummary   `json:"summary"`
    Details     []*QualityDetail  `json:"details"`
    Issues      []*QualityIssue   `json:"issues"`
    Suggestions []*QualitySuggestion `json:"suggestions"`
}

// 质量摘要
type QualitySummary struct {
    OverallScore float64           `json:"overall_score"`
    CategoryScores map[string]float64 `json:"category_scores"`
    KeyMetrics   map[string]float64 `json:"key_metrics"`
    Status       string            `json:"status"`
}

// 质量详情
type QualityDetail struct {
    Category    string            `json:"category"`
    Metrics     map[string]float64 `json:"metrics"`
    Score       float64           `json:"score"`
    Trend       string            `json:"trend"`
}

// 质量问题
type QualityIssue struct {
    Type        string            `json:"type"`
    Level       string            `json:"level"`
    Description string            `json:"description"`
    Impact      string            `json:"impact"`
    Solution    string            `json:"solution"`
}
```

## 6. 实现路径

### 6.1 基础框架搭建
1. 项目初始化
   ```bash
   # 创建项目目录
   mkdir -p cmd/placement
   mkdir -p internal/{placement,spec,traffic,quality}
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
   # config/placement.yaml
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
   
   quality:
     check_interval: "5m"
     metrics_retention: "30d"
     alert_thresholds:
       ctr: 0.001
       viewability: 0.5
       latency: 1000
   ```

### 6.2 核心功能实现

#### 6.2.1 广告位管理实现
```go
// internal/placement/service.go
type PlacementService struct {
    db          *gorm.DB
    cache       *redis.Client
    influx      *influxdb.Client
    logger      *zap.Logger
}

func (s *PlacementService) Create(ctx context.Context, req *CreatePlacementRequest) (*Placement, error) {
    // 1. 验证请求
    if err := s.validateCreateRequest(req); err != nil {
        return nil, err
    }
    
    // 2. 创建广告位
    placement := &Placement{
        Name:      req.Name,
        Type:      req.Type,
        Publisher: req.Publisher,
        Category:  req.Category,
        Specs:     req.Specs,
        Settings:  req.Settings,
        Status:    PlacementStatusPending,
    }
    
    // 3. 保存广告位
    if err := s.db.Create(placement).Error; err != nil {
        return nil, err
    }
    
    // 4. 缓存广告位
    if err := s.cachePlacement(placement); err != nil {
        s.logger.Warn("failed to cache placement", zap.Error(err))
    }
    
    // 5. 初始化统计
    if err := s.initPlacementStats(placement); err != nil {
        s.logger.Warn("failed to init placement stats", zap.Error(err))
    }
    
    return placement, nil
}

func (s *PlacementService) Update(ctx context.Context, req *UpdatePlacementRequest) (*Placement, error) {
    // 1. 获取广告位
    placement, err := s.Get(ctx, req.ID)
    if err != nil {
        return nil, err
    }
    
    // 2. 更新广告位
    if err := s.db.Model(placement).Updates(req.Updates).Error; err != nil {
        return nil, err
    }
    
    // 3. 更新缓存
    if err := s.cachePlacement(placement); err != nil {
        s.logger.Warn("failed to update placement cache", zap.Error(err))
    }
    
    return placement, nil
}
```

#### 6.2.2 规格管理实现
```go
// internal/spec/service.go
type SpecService struct {
    db          *gorm.DB
    cache       *redis.Client
    logger      *zap.Logger
}

func (s *SpecService) Create(ctx context.Context, req *CreateSpecRequest) (*Spec, error) {
    // 1. 验证请求
    if err := s.validateCreateRequest(req); err != nil {
        return nil, err
    }
    
    // 2. 创建规格
    spec := &Spec{
        PlacementID: req.PlacementID,
        Type:        req.Type,
        Format:      req.Format,
        Size:        req.Size,
        Settings:    req.Settings,
        Status:      SpecStatusActive,
    }
    
    // 3. 保存规格
    if err := s.db.Create(spec).Error; err != nil {
        return nil, err
    }
    
    // 4. 缓存规格
    if err := s.cacheSpec(spec); err != nil {
        s.logger.Warn("failed to cache spec", zap.Error(err))
    }
    
    return spec, nil
}

func (s *SpecService) Validate(ctx context.Context, spec *Spec) (*ValidationResult, error) {
    // 1. 验证尺寸
    if err := s.validateSize(spec.Size); err != nil {
        return nil, err
    }
    
    // 2. 验证格式
    if err := s.validateFormat(spec.Format); err != nil {
        return nil, err
    }
    
    // 3. 验证设置
    if err := s.validateSettings(spec.Settings); err != nil {
        return nil, err
    }
    
    return &ValidationResult{
        Valid: true,
        Message: "validation successful",
    }, nil
}
```

#### 6.2.3 流量管理实现
```go
// internal/traffic/service.go
type TrafficService struct {
    db          *gorm.DB
    influx      *influxdb.Client
    logger      *zap.Logger
}

func (s *TrafficService) GetStats(ctx context.Context, query *TrafficQuery) (*TrafficStats, error) {
    // 1. 构建查询
    fluxQuery := s.buildStatsQuery(query)
    
    // 2. 执行查询
    result, err := s.influx.Query(ctx, fluxQuery)
    if err != nil {
        return nil, err
    }
    
    // 3. 处理结果
    stats, err := s.processStatsResult(result)
    if err != nil {
        return nil, err
    }
    
    return stats, nil
}

func (s *TrafficService) GetAnalysis(ctx context.Context, query *TrafficQuery) (*TrafficAnalysis, error) {
    // 1. 获取统计数据
    stats, err := s.GetStats(ctx, query)
    if err != nil {
        return nil, err
    }
    
    // 2. 分析趋势
    trends, err := s.analyzeTrends(stats)
    if err != nil {
        return nil, err
    }
    
    // 3. 分析模式
    patterns, err := s.analyzePatterns(stats)
    if err != nil {
        return nil, err
    }
    
    // 4. 生成预测
    predictions, err := s.generatePredictions(stats)
    if err != nil {
        return nil, err
    }
    
    return &TrafficAnalysis{
        PlacementID: query.PlacementID,
        TimeRange:   query.TimeRange,
        Trends:      trends,
        Patterns:    patterns,
        Predictions: predictions,
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
// internal/placement/service_test.go
func TestPlacementService_Create(t *testing.T) {
    tests := []struct {
        name    string
        req     *CreatePlacementRequest
        want    *Placement
        wantErr bool
    }{
        // 测试用例
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            s := NewPlacementService(db, cache, influx, logger)
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
// internal/placement/service_benchmark_test.go
func BenchmarkPlacementService_Create(b *testing.B) {
    s := NewPlacementService(db, cache, influx, logger)
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
RUN go build -o placement ./cmd/placement

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/placement .
COPY --from=builder /app/config ./config
EXPOSE 8080
CMD ["./placement"]
```

### 10.2 Kubernetes部署
```yaml
# k8s/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: placement
spec:
  replicas: 3
  selector:
    matchLabels:
      app: placement
  template:
    metadata:
      labels:
        app: placement
    spec:
      containers:
      - name: placement
        image: placement:latest
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
          name: placement-config
```

### 10.3 监控配置
```yaml
# prometheus/service-monitor.yaml
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: placement
spec:
  selector:
    matchLabels:
      app: placement
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
1. 广告位管理工具
2. 规格管理工具
3. 流量分析工具
4. 质量监控工具 