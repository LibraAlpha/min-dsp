# 日志管理模块设计文档

## 1. 模块概述

日志管理模块是DSP系统集成层的核心组件之一，负责系统各类日志的收集、存储、分析和展示等功能。该模块需要支持多源日志采集、实时日志处理、分布式日志存储、智能日志分析，以及灵活的日志查询和可视化展示。

## 2. 功能列表

### 2.1 核心功能
1. 日志采集
2. 日志存储
3. 日志分析
4. 日志查询
5. 日志展示

### 2.2 扩展功能
1. 日志告警
2. 日志统计
3. 日志审计
4. 日志备份
5. 日志清理

## 3. 详细设计

### 3.1 系统架构
```
                    ┌─────────────┐
                    │  接入层     │
                    └──────┬──────┘
                           │
                    ┌──────▼──────┐
                    │  采集层     │
                    └──────┬──────┘
                           │
         ┌─────────┬───────┴───────┬─────────┐
         │         │               │         │
┌────────▼───┐ ┌───▼──────┐  ┌────▼────┐ ┌──▼──────┐
│ 系统日志    │ │ 应用日志  │  │ 访问日志 │ │ 审计日志 │
└───────┬────┘ └─────┬────┘  └────┬────┘ └────┬────┘
        │            │            │           │
        └────────────┼────────────┼───────────┘
                     │            │
               ┌─────▼────────────▼─────┐
               │      存储层            │
               └─────────────┬──────────┘
                             │
         ┌─────────┬─────────┼─────────┬─────────┐
         │         │         │         │         │
┌────────▼───┐ ┌───▼────┐ ┌──▼────┐ ┌──▼────┐ ┌──▼────┐
│ 实时存储    │ │ 离线存储 │ │ 备份存储 │ │ 归档存储 │ │ 缓存存储 │
└───────┬────┘ └────┬────┘ └───┬────┘ └───┬────┘ └───┬────┘
        │           │          │          │          │
        └───────────┼──────────┼──────────┼──────────┘
                    │          │          │
               ┌────▼──────────▼──────────▼────┐
               │           分析层              │
               └─────────────┬─────────────────┘
                             │
         ┌─────────┬─────────┼─────────┬─────────┐
         │         │         │         │         │
┌────────▼───┐ ┌───▼────┐ ┌──▼────┐ ┌──▼────┐ ┌──▼────┐
│ 日志分析    │ │ 日志告警 │ │ 日志统计 │ │ 日志审计 │ │ 日志展示 │
└────────────┘ └─────────┘ └────────┘ └────────┘ └────────┘
```

### 3.2 组件说明

#### 3.2.1 接入层
- API接口管理
- 数据验证
- 访问控制
- 请求日志

#### 3.2.2 采集层
- 系统日志：操作系统、中间件等日志
- 应用日志：业务应用、服务等日志
- 访问日志：用户访问、API调用等日志
- 审计日志：操作审计、安全审计等日志

#### 3.2.3 存储层
- 实时存储：Elasticsearch等
- 离线存储：HDFS等
- 备份存储：对象存储等
- 归档存储：冷数据存储
- 缓存存储：Redis等

#### 3.2.4 分析层
- 日志分析：日志解析、特征提取
- 日志告警：异常检测、告警通知
- 日志统计：统计分析、趋势分析
- 日志审计：安全审计、合规审计
- 日志展示：可视化展示、报表生成

## 4. 接口定义

### 4.1 外部接口
```go
// 日志管理服务接口
type LogManager interface {
    // 日志采集
    CollectLog(ctx context.Context, req *CollectRequest) error
    BatchCollect(ctx context.Context, req *BatchCollectRequest) error
    
    // 日志查询
    QueryLog(ctx context.Context, req *QueryRequest) (*QueryResponse, error)
    SearchLog(ctx context.Context, req *SearchRequest) (*SearchResponse, error)
    
    // 日志分析
    AnalyzeLog(ctx context.Context, req *AnalyzeRequest) (*AnalyzeResponse, error)
    GetStats(ctx context.Context, req *StatsRequest) (*StatsResponse, error)
    
    // 日志管理
    ExportLog(ctx context.Context, req *ExportRequest) error
    CleanLog(ctx context.Context, req *CleanRequest) error
    BackupLog(ctx context.Context, req *BackupRequest) error
}

// 日志采集请求
type CollectRequest struct {
    Log         *Log             `json:"log"`
    Source      string           `json:"source"`
    Type        LogType          `json:"type"`
    Metadata    map[string]string `json:"metadata"`
}

// 日志查询请求
type QueryRequest struct {
    Source      string           `json:"source"`
    Type        LogType          `json:"type"`
    StartTime   time.Time        `json:"start_time"`
    EndTime     time.Time        `json:"end_time"`
    Filter      *LogFilter       `json:"filter"`
    Page        *PageRequest     `json:"page"`
}

// 日志分析请求
type AnalyzeRequest struct {
    Source      string           `json:"source"`
    Type        LogType          `json:"type"`
    StartTime   time.Time        `json:"start_time"`
    EndTime     time.Time        `json:"end_time"`
    Analysis    *AnalysisConfig  `json:"analysis"`
}
```

### 4.2 内部接口
```go
// 日志采集服务接口
type CollectorService interface {
    // 采集日志
    Collect(ctx context.Context, log *Log) error
    
    // 批量采集
    BatchCollect(ctx context.Context, logs []*Log) error
    
    // 获取采集器
    GetCollector(ctx context.Context, source string) (Collector, error)
    
    // 列出采集器
    ListCollectors(ctx context.Context) ([]Collector, error)
}

// 日志存储服务接口
type StorageService interface {
    // 存储日志
    Store(ctx context.Context, log *Log) error
    
    // 批量存储
    BatchStore(ctx context.Context, logs []*Log) error
    
    // 查询日志
    Query(ctx context.Context, filter *LogFilter) ([]*Log, error)
    
    // 删除日志
    Delete(ctx context.Context, filter *LogFilter) error
}

// 日志分析服务接口
type AnalyzerService interface {
    // 分析日志
    Analyze(ctx context.Context, logs []*Log) (*Analysis, error)
    
    // 获取分析
    GetAnalysis(ctx context.Context, id string) (*Analysis, error)
    
    // 获取统计
    GetStats(ctx context.Context, filter *AnalysisFilter) (*AnalysisStats, error)
}

// 日志告警服务接口
type AlertService interface {
    // 检测告警
    Detect(ctx context.Context, log *Log) error
    
    // 添加规则
    AddRule(ctx context.Context, rule *AlertRule) error
    
    // 删除规则
    DeleteRule(ctx context.Context, ruleID string) error
    
    // 获取规则
    GetRule(ctx context.Context, ruleID string) (*AlertRule, error)
}
```

## 5. 数据结构

### 5.1 日志数据
```go
// 日志类型
type LogType string

const (
    LogTypeSystem   LogType = "system"
    LogTypeApp      LogType = "app"
    LogTypeAccess   LogType = "access"
    LogTypeAudit    LogType = "audit"
)

// 日志级别
type LogLevel string

const (
    LogLevelDebug   LogLevel = "debug"
    LogLevelInfo    LogLevel = "info"
    LogLevelWarning LogLevel = "warning"
    LogLevelError   LogLevel = "error"
    LogLevelFatal   LogLevel = "fatal"
)

// 日志信息
type Log struct {
    ID          string            `json:"id"`
    Source      string            `json:"source"`
    Type        LogType           `json:"type"`
    Level       LogLevel          `json:"level"`
    Content     string            `json:"content"`
    Fields      map[string]string `json:"fields"`
    Timestamp   time.Time         `json:"timestamp"`
    Metadata    map[string]string `json:"metadata"`
}

// 日志配置
type LogConfig struct {
    Source      string            `json:"source"`
    Type        LogType           `json:"type"`
    Level       LogLevel          `json:"level"`
    Format      string            `json:"format"`
    Retention   string            `json:"retention"`
    Metadata    map[string]string `json:"metadata"`
}
```

### 5.2 分析数据
```go
// 分析类型
type AnalysisType string

const (
    AnalysisTypePattern  AnalysisType = "pattern"
    AnalysisTypeTrend    AnalysisType = "trend"
    AnalysisTypeAnomaly  AnalysisType = "anomaly"
)

// 分析信息
type Analysis struct {
    ID          string            `json:"id"`
    Type        AnalysisType      `json:"type"`
    Source      string            `json:"source"`
    Pattern     string            `json:"pattern"`
    Result      interface{}       `json:"result"`
    Status      string            `json:"status"`
    CreatedAt   time.Time         `json:"created_at"`
    UpdatedAt   time.Time         `json:"updated_at"`
    Metadata    map[string]string `json:"metadata"`
}

// 分析配置
type AnalysisConfig struct {
    Type        AnalysisType      `json:"type"`
    Pattern     string            `json:"pattern"`
    Parameters  map[string]string `json:"parameters"`
    Metadata    map[string]string `json:"metadata"`
}

// 分析统计
type AnalysisStats struct {
    Total       int               `json:"total"`
    Patterns    int               `json:"patterns"`
    Anomalies   int               `json:"anomalies"`
    UpdatedAt   time.Time         `json:"updated_at"`
}
```

### 5.3 告警数据
```go
// 告警规则
type AlertRule struct {
    ID          string            `json:"id"`
    Name        string            `json:"name"`
    Type        LogType           `json:"type"`
    Pattern     string            `json:"pattern"`
    Condition   string            `json:"condition"`
    Threshold   int               `json:"threshold"`
    Duration    string            `json:"duration"`
    Notify      *NotifyConfig     `json:"notify"`
    Metadata    map[string]string `json:"metadata"`
}

// 告警信息
type Alert struct {
    ID          string            `json:"id"`
    Rule        *AlertRule        `json:"rule"`
    Log         *Log              `json:"log"`
    Status      string            `json:"status"`
    CreatedAt   time.Time         `json:"created_at"`
    UpdatedAt   time.Time         `json:"updated_at"`
    Metadata    map[string]string `json:"metadata"`
}

// 告警统计
type AlertStats struct {
    Total       int               `json:"total"`
    Firing      int               `json:"firing"`
    Resolved    int               `json:"resolved"`
    UpdatedAt   time.Time         `json:"updated_at"`
}
```

### 5.4 查询数据
```go
// 日志过滤器
type LogFilter struct {
    Source      string            `json:"source"`
    Type        LogType           `json:"type"`
    Level       LogLevel          `json:"level"`
    StartTime   time.Time         `json:"start_time"`
    EndTime     time.Time         `json:"end_time"`
    Fields      map[string]string `json:"fields"`
    Query       string            `json:"query"`
}

// 分页请求
type PageRequest struct {
    Page        int               `json:"page"`
    Size        int               `json:"size"`
    Sort        string            `json:"sort"`
    Order       string            `json:"order"`
}

// 查询结果
type QueryResponse struct {
    Total       int64             `json:"total"`
    Page        int               `json:"page"`
    Size        int               `json:"size"`
    Logs        []*Log            `json:"logs"`
}
```

## 6. 实现路径

### 6.1 基础框架搭建
1. 项目初始化
   ```bash
   # 创建项目目录
   mkdir -p cmd/log
   mkdir -p internal/{collector,storage,analyzer,alert}
   mkdir -p pkg/{api,service,utils}
   
   # 初始化 Go 模块
   go mod init min-dsp
   
   # 添加依赖
   go get -u github.com/gin-gonic/gin
   go get -u gorm.io/gorm
   go get -u github.com/go-redis/redis/v8
   go get -u github.com/elastic/go-elasticsearch/v8
   go get -u go.uber.org/zap
   ```

2. 配置文件
   ```yaml
   # config/log.yaml
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
   
   elasticsearch:
     hosts:
       - "http://localhost:9200"
     username: "elastic"
     password: "******"
     index_prefix: "dsp_"
   
   log:
     level: "info"
     format: "json"
     retention: "30d"
     max_size: 100
     max_backups: 10
   
   collector:
     batch_size: 1000
     flush_interval: "5s"
     retry_times: 3
   
   analyzer:
     workers: 10
     batch_size: 100
     timeout: "30s"
   ```

### 6.2 核心功能实现

#### 6.2.1 日志采集实现
```go
// internal/collector/service.go
type CollectorService struct {
    db          *gorm.DB
    cache       *redis.Client
    es          *elasticsearch.Client
    logger      *zap.Logger
}

func (s *CollectorService) Collect(ctx context.Context, log *Log) error {
    // 1. 验证日志
    if err := s.validateLog(log); err != nil {
        return err
    }
    
    // 2. 处理日志
    if err := s.processLog(ctx, log); err != nil {
        return err
    }
    
    // 3. 存储日志
    if err := s.storeLog(ctx, log); err != nil {
        return err
    }
    
    // 4. 分析日志
    if err := s.analyzeLog(ctx, log); err != nil {
        s.logger.Warn("failed to analyze log", zap.Error(err))
    }
    
    return nil
}

func (s *CollectorService) processLog(ctx context.Context, log *Log) error {
    // 1. 解析日志
    if err := s.parseLog(log); err != nil {
        return err
    }
    
    // 2. 提取字段
    if err := s.extractFields(log); err != nil {
        return err
    }
    
    // 3. 过滤日志
    if s.filterLog(log) {
        return nil
    }
    
    // 4. 转换日志
    return s.transformLog(log)
}

func (s *CollectorService) storeLog(ctx context.Context, log *Log) error {
    // 1. 存储到 Elasticsearch
    doc := map[string]interface{}{
        "source":    log.Source,
        "type":      log.Type,
        "level":     log.Level,
        "content":   log.Content,
        "fields":    log.Fields,
        "timestamp": log.Timestamp,
    }
    
    _, err := s.es.Index().
        Index(s.getIndex(log)).
        Document(doc).
        Do(ctx)
    
    if err != nil {
        return err
    }
    
    // 2. 更新缓存
    if err := s.updateCache(ctx, log); err != nil {
        s.logger.Warn("failed to update cache", zap.Error(err))
    }
    
    return nil
}
```

#### 6.2.2 日志分析实现
```go
// internal/analyzer/service.go
type AnalyzerService struct {
    db          *gorm.DB
    cache       *redis.Client
    logger      *zap.Logger
}

func (s *AnalyzerService) Analyze(ctx context.Context, logs []*Log) (*Analysis, error) {
    // 1. 获取分析配置
    config, err := s.getConfig(ctx)
    if err != nil {
        return nil, err
    }
    
    // 2. 分析日志
    switch config.Type {
    case AnalysisTypePattern:
        return s.analyzePattern(ctx, logs, config)
    case AnalysisTypeTrend:
        return s.analyzeTrend(ctx, logs, config)
    case AnalysisTypeAnomaly:
        return s.analyzeAnomaly(ctx, logs, config)
    default:
        return nil, fmt.Errorf("unsupported analysis type: %s", config.Type)
    }
}

func (s *AnalyzerService) analyzePattern(ctx context.Context, logs []*Log, config *AnalysisConfig) (*Analysis, error) {
    // 1. 编译模式
    pattern, err := regexp.Compile(config.Pattern)
    if err != nil {
        return nil, err
    }
    
    // 2. 匹配模式
    matches := make([]string, 0)
    for _, log := range logs {
        if pattern.MatchString(log.Content) {
            matches = append(matches, log.Content)
        }
    }
    
    // 3. 创建分析
    return &Analysis{
        Type:    AnalysisTypePattern,
        Pattern: config.Pattern,
        Result:  matches,
        Status:  "completed",
    }, nil
}

func (s *AnalyzerService) analyzeTrend(ctx context.Context, logs []*Log, config *AnalysisConfig) (*Analysis, error) {
    // 1. 按时间分组
    groups := make(map[time.Time][]*Log)
    for _, log := range logs {
        key := log.Timestamp.Truncate(time.Hour)
        groups[key] = append(groups[key], log)
    }
    
    // 2. 计算趋势
    trend := make(map[time.Time]int)
    for t, group := range groups {
        trend[t] = len(group)
    }
    
    // 3. 创建分析
    return &Analysis{
        Type:    AnalysisTypeTrend,
        Pattern: config.Pattern,
        Result:  trend,
        Status:  "completed",
    }, nil
}
```

#### 6.2.3 日志告警实现
```go
// internal/alert/service.go
type AlertService struct {
    db          *gorm.DB
    cache       *redis.Client
    logger      *zap.Logger
}

func (s *AlertService) Detect(ctx context.Context, log *Log) error {
    // 1. 获取规则
    rules, err := s.getRules(ctx, log.Type)
    if err != nil {
        return err
    }
    
    // 2. 检测告警
    for _, rule := range rules {
        if err := s.checkRule(ctx, rule, log); err != nil {
            s.logger.Warn("failed to check rule", zap.Error(err))
            continue
        }
    }
    
    return nil
}

func (s *AlertService) checkRule(ctx context.Context, rule *AlertRule, log *Log) error {
    // 1. 检查模式
    if !s.matchPattern(rule.Pattern, log) {
        return nil
    }
    
    // 2. 检查条件
    if !s.matchCondition(rule.Condition, log) {
        return nil
    }
    
    // 3. 检查阈值
    if !s.checkThreshold(ctx, rule, log) {
        return nil
    }
    
    // 4. 创建告警
    alert := &Alert{
        Rule:      rule,
        Log:       log,
        Status:    "firing",
    }
    
    // 5. 存储告警
    if err := s.storeAlert(ctx, alert); err != nil {
        return err
    }
    
    // 6. 发送通知
    return s.notifyAlert(ctx, alert)
}

func (s *AlertService) matchPattern(pattern string, log *Log) bool {
    // 1. 编译模式
    re, err := regexp.Compile(pattern)
    if err != nil {
        return false
    }
    
    // 2. 匹配内容
    return re.MatchString(log.Content)
}

func (s *AlertService) matchCondition(condition string, log *Log) bool {
    // 1. 解析条件
    expr, err := parser.ParseExpr(condition)
    if err != nil {
        return false
    }
    
    // 2. 计算条件
    result, err := eval.Eval(expr, map[string]interface{}{
        "log": log,
    })
    if err != nil {
        return false
    }
    
    return result.(bool)
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
// internal/collector/service_test.go
func TestCollectorService_Collect(t *testing.T) {
    tests := []struct {
        name    string
        log     *Log
        wantErr bool
    }{
        // 测试用例
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            s := NewCollectorService(db, cache, es, logger)
            err := s.Collect(context.Background(), tt.log)
            if (err != nil) != tt.wantErr {
                t.Errorf("Collect() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
        })
    }
}
```

### 9.2 性能测试
```go
// internal/collector/service_benchmark_test.go
func BenchmarkCollectorService_Collect(b *testing.B) {
    s := NewCollectorService(db, cache, es, logger)
    log := createTestLog()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        err := s.Collect(context.Background(), log)
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
RUN go build -o log ./cmd/log

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/log .
COPY --from=builder /app/config ./config
EXPOSE 8080
CMD ["./log"]
```

### 10.2 Kubernetes部署
```yaml
# k8s/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: log
spec:
  replicas: 3
  selector:
    matchLabels:
      app: log
  template:
    metadata:
      labels:
        app: log
    spec:
      containers:
      - name: log
        image: log:latest
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
          name: log-config
```

### 10.3 监控配置
```yaml
# prometheus/service-monitor.yaml
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: log
spec:
  selector:
    matchLabels:
      app: log
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
1. 日志管理工具
2. 日志分析工具
3. 日志诊断工具
4. 日志备份工具 