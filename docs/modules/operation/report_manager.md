# 报表统计模块设计文档

## 1. 模块概述

报表统计模块是DSP系统运营管理层的核心组件之一，负责收集、处理、分析和展示各类业务数据，为运营决策提供数据支持。该模块需要支持多维度的数据统计、灵活的报表配置、实时的数据更新、丰富的可视化展示等功能，同时确保数据的准确性和及时性。

## 2. 功能列表

### 2.1 核心功能
1. 数据采集
2. 数据处理
3. 报表生成
4. 数据展示
5. 数据导出

### 2.2 扩展功能
1. 自定义报表
2. 数据预警
3. 趋势分析
4. 竞品分析
5. 数据大屏

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
│ 数据采集    │ │ 数据处理  │  │ 报表生成 │ │ 数据展示 │
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
- 数据采集：实时数据采集、离线数据采集、数据清洗等
- 数据处理：数据转换、数据聚合、数据计算等
- 报表生成：报表模板、报表配置、报表生成等
- 数据展示：数据可视化、数据大屏、数据导出等

#### 3.2.3 数据层
- 数据访问
- 数据验证
- 数据转换
- 数据缓存

#### 3.2.4 存储层
- 关系型数据库
- 缓存数据库
- 时序数据库
- 数据仓库

## 4. 接口定义

### 4.1 外部接口
```go
// 报表管理服务接口
type ReportManager interface {
    // 数据采集
    CollectData(ctx context.Context, req *CollectDataRequest) (*CollectDataResponse, error)
    GetDataStatus(ctx context.Context, req *GetDataStatusRequest) (*DataStatusResponse, error)
    
    // 报表管理
    CreateReport(ctx context.Context, req *CreateReportRequest) (*ReportResponse, error)
    UpdateReport(ctx context.Context, req *UpdateReportRequest) (*ReportResponse, error)
    GetReport(ctx context.Context, req *GetReportRequest) (*ReportResponse, error)
    ListReports(ctx context.Context, req *ListReportsRequest) (*ListReportsResponse, error)
    
    // 数据查询
    QueryData(ctx context.Context, req *QueryDataRequest) (*QueryDataResponse, error)
    GetMetrics(ctx context.Context, req *GetMetricsRequest) (*MetricsResponse, error)
    
    // 数据导出
    ExportData(ctx context.Context, req *ExportDataRequest) (*ExportDataResponse, error)
    GetExportStatus(ctx context.Context, req *GetExportStatusRequest) (*ExportStatusResponse, error)
}

// 数据采集请求
type CollectDataRequest struct {
    Type        DataType          `json:"type"`
    Source      string            `json:"source"`
    StartTime   time.Time         `json:"start_time"`
    EndTime     time.Time         `json:"end_time"`
    Parameters  map[string]string `json:"parameters"`
    Metadata    map[string]string `json:"metadata"`
}

// 报表创建请求
type CreateReportRequest struct {
    Name        string            `json:"name"`
    Type        ReportType        `json:"type"`
    Template    string            `json:"template"`
    Config      *ReportConfig     `json:"config"`
    Schedule    *ReportSchedule   `json:"schedule"`
    Recipients  []string          `json:"recipients"`
    Metadata    map[string]string `json:"metadata"`
}

// 数据查询请求
type QueryDataRequest struct {
    ReportID    string            `json:"report_id"`
    StartTime   time.Time         `json:"start_time"`
    EndTime     time.Time         `json:"end_time"`
    Dimensions  []string          `json:"dimensions"`
    Metrics     []string          `json:"metrics"`
    Filters     []*QueryFilter    `json:"filters"`
    Sort        []*QuerySort      `json:"sort"`
    Page        *QueryPage        `json:"page"`
}

// 数据导出请求
type ExportDataRequest struct {
    ReportID    string            `json:"report_id"`
    Format      ExportFormat      `json:"format"`
    StartTime   time.Time         `json:"start_time"`
    EndTime     time.Time         `json:"end_time"`
    Parameters  map[string]string `json:"parameters"`
}
```

### 4.2 内部接口
```go
// 数据采集服务接口
type CollectorService interface {
    // 采集数据
    Collect(ctx context.Context, request *CollectRequest) error
    
    // 获取状态
    GetStatus(ctx context.Context, id string) (*CollectStatus, error)
    
    // 验证数据
    Validate(ctx context.Context, data *CollectData) (*ValidationResult, error)
}

// 数据处理服务接口
type ProcessorService interface {
    // 处理数据
    Process(ctx context.Context, data *ProcessData) (*ProcessResult, error)
    
    // 转换数据
    Transform(ctx context.Context, data *TransformData) (*TransformResult, error)
    
    // 聚合数据
    Aggregate(ctx context.Context, data *AggregateData) (*AggregateResult, error)
}

// 报表服务接口
type ReportService interface {
    // 创建报表
    Create(ctx context.Context, report *Report) (*Report, error)
    
    // 更新报表
    Update(ctx context.Context, report *Report) (*Report, error)
    
    // 获取报表
    Get(ctx context.Context, id string) (*Report, error)
    
    // 生成报表
    Generate(ctx context.Context, id string) (*GenerateResult, error)
}

// 查询服务接口
type QueryService interface {
    // 查询数据
    Query(ctx context.Context, query *DataQuery) (*QueryResult, error)
    
    // 获取指标
    GetMetrics(ctx context.Context, query *MetricsQuery) (*MetricsResult, error)
    
    // 导出数据
    Export(ctx context.Context, request *ExportRequest) (*ExportResult, error)
}
```

## 5. 数据结构

### 5.1 数据采集
```go
// 数据类型
type DataType string

const (
    DataTypeRealTime  DataType = "realtime"
    DataTypeOffline   DataType = "offline"
    DataTypeBatch     DataType = "batch"
)

// 采集数据
type CollectData struct {
    ID          string            `json:"id"`
    Type        DataType          `json:"type"`
    Source      string            `json:"source"`
    Data        interface{}       `json:"data"`
    Status      CollectStatus     `json:"status"`
    StartTime   time.Time         `json:"start_time"`
    EndTime     time.Time         `json:"end_time"`
    CreatedAt   time.Time         `json:"created_at"`
    UpdatedAt   time.Time         `json:"updated_at"`
}

// 采集状态
type CollectStatus struct {
    Status      string            `json:"status"`
    Progress    float64           `json:"progress"`
    Total       int64             `json:"total"`
    Processed   int64             `json:"processed"`
    Failed      int64             `json:"failed"`
    Error       string            `json:"error"`
    UpdatedAt   time.Time         `json:"updated_at"`
}
```

### 5.2 报表定义
```go
// 报表类型
type ReportType string

const (
    ReportTypeDaily    ReportType = "daily"
    ReportTypeWeekly   ReportType = "weekly"
    ReportTypeMonthly  ReportType = "monthly"
    ReportTypeCustom   ReportType = "custom"
)

// 报表信息
type Report struct {
    ID          string            `json:"id"`
    Name        string            `json:"name"`
    Type        ReportType        `json:"type"`
    Template    string            `json:"template"`
    Config      *ReportConfig     `json:"config"`
    Schedule    *ReportSchedule   `json:"schedule"`
    Recipients  []string          `json:"recipients"`
    Status      ReportStatus      `json:"status"`
    CreatedAt   time.Time         `json:"created_at"`
    UpdatedAt   time.Time         `json:"updated_at"`
    Metadata    map[string]string `json:"metadata"`
}

// 报表配置
type ReportConfig struct {
    Dimensions  []string          `json:"dimensions"`
    Metrics     []string          `json:"metrics"`
    Filters     []*ReportFilter   `json:"filters"`
    Sort        []*ReportSort     `json:"sort"`
    Format      *ReportFormat     `json:"format"`
    Charts      []*ReportChart    `json:"charts"`
}

// 报表计划
type ReportSchedule struct {
    Type        string            `json:"type"`
    Interval    string            `json:"interval"`
    StartTime   time.Time         `json:"start_time"`
    EndTime     time.Time         `json:"end_time"`
    Timezone    string            `json:"timezone"`
    LastRun     time.Time         `json:"last_run"`
    NextRun     time.Time         `json:"next_run"`
}
```

### 5.3 查询数据
```go
// 查询结果
type QueryResult struct {
    ReportID    string            `json:"report_id"`
    TimeRange   *TimeRange        `json:"time_range"`
    Dimensions  []string          `json:"dimensions"`
    Metrics     []string          `json:"metrics"`
    Data        []*QueryRow       `json:"data"`
    Total       int64             `json:"total"`
    Page        *QueryPage        `json:"page"`
}

// 查询行
type QueryRow struct {
    Dimensions  map[string]string `json:"dimensions"`
    Metrics     map[string]float64 `json:"metrics"`
    Timestamp   time.Time         `json:"timestamp"`
}

// 查询分页
type QueryPage struct {
    PageNum     int               `json:"page_num"`
    PageSize    int               `json:"page_size"`
    TotalPages  int               `json:"total_pages"`
    HasMore     bool              `json:"has_more"`
}
```

### 5.4 导出数据
```go
// 导出格式
type ExportFormat string

const (
    ExportFormatCSV    ExportFormat = "csv"
    ExportFormatExcel  ExportFormat = "excel"
    ExportFormatPDF    ExportFormat = "pdf"
    ExportFormatJSON   ExportFormat = "json"
)

// 导出结果
type ExportResult struct {
    ID          string            `json:"id"`
    ReportID    string            `json:"report_id"`
    Format      ExportFormat      `json:"format"`
    Status      ExportStatus      `json:"status"`
    URL         string            `json:"url"`
    Size        int64             `json:"size"`
    CreatedAt   time.Time         `json:"created_at"`
    ExpiredAt   time.Time         `json:"expired_at"`
}

// 导出状态
type ExportStatus struct {
    Status      string            `json:"status"`
    Progress    float64           `json:"progress"`
    Error       string            `json:"error"`
    UpdatedAt   time.Time         `json:"updated_at"`
}
```

## 6. 实现路径

### 6.1 基础框架搭建
1. 项目初始化
   ```bash
   # 创建项目目录
   mkdir -p cmd/report
   mkdir -p internal/{collector,processor,report,query}
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
   # config/report.yaml
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
   
   clickhouse:
     host: "localhost"
     port: 9000
     database: "dsp"
     username: "default"
     password: "******"
   
   storage:
     type: "s3"
     bucket: "dsp-reports"
     region: "us-east-1"
     access_key: "******"
     secret_key: "******"
   
   monitoring:
     log_level: "info"
     metrics_port: 9090
     alert_threshold: 0.9
   ```

### 6.2 核心功能实现

#### 6.2.1 数据采集实现
```go
// internal/collector/service.go
type CollectorService struct {
    db          *gorm.DB
    cache       *redis.Client
    influx      *influxdb.Client
    clickhouse  *clickhouse.Conn
    logger      *zap.Logger
}

func (s *CollectorService) Collect(ctx context.Context, req *CollectDataRequest) error {
    // 1. 验证请求
    if err := s.validateCollectRequest(req); err != nil {
        return err
    }
    
    // 2. 创建采集任务
    task := &CollectTask{
        Type:       req.Type,
        Source:     req.Source,
        StartTime:  req.StartTime,
        EndTime:    req.EndTime,
        Status:     CollectStatusPending,
    }
    
    // 3. 保存任务
    if err := s.db.Create(task).Error; err != nil {
        return err
    }
    
    // 4. 启动采集
    go s.startCollect(task)
    
    return nil
}

func (s *CollectorService) startCollect(task *CollectTask) {
    // 1. 更新状态
    task.Status = CollectStatusRunning
    s.db.Save(task)
    
    // 2. 采集数据
    data, err := s.collectData(task)
    if err != nil {
        task.Status = CollectStatusFailed
        task.Error = err.Error()
        s.db.Save(task)
        return
    }
    
    // 3. 处理数据
    if err := s.processData(data); err != nil {
        task.Status = CollectStatusFailed
        task.Error = err.Error()
        s.db.Save(task)
        return
    }
    
    // 4. 更新状态
    task.Status = CollectStatusCompleted
    s.db.Save(task)
}
```

#### 6.2.2 报表生成实现
```go
// internal/report/service.go
type ReportService struct {
    db          *gorm.DB
    cache       *redis.Client
    storage     StorageService
    query       QueryService
    logger      *zap.Logger
}

func (s *ReportService) Create(ctx context.Context, req *CreateReportRequest) (*Report, error) {
    // 1. 验证请求
    if err := s.validateCreateRequest(req); err != nil {
        return nil, err
    }
    
    // 2. 创建报表
    report := &Report{
        Name:       req.Name,
        Type:       req.Type,
        Template:   req.Template,
        Config:     req.Config,
        Schedule:   req.Schedule,
        Recipients: req.Recipients,
        Status:     ReportStatusActive,
    }
    
    // 3. 保存报表
    if err := s.db.Create(report).Error; err != nil {
        return nil, err
    }
    
    // 4. 缓存报表
    if err := s.cacheReport(report); err != nil {
        s.logger.Warn("failed to cache report", zap.Error(err))
    }
    
    // 5. 启动调度
    if report.Schedule != nil {
        go s.startSchedule(report)
    }
    
    return report, nil
}

func (s *ReportService) Generate(ctx context.Context, id string) (*GenerateResult, error) {
    // 1. 获取报表
    report, err := s.Get(ctx, id)
    if err != nil {
        return nil, err
    }
    
    // 2. 查询数据
    data, err := s.query.Query(ctx, &DataQuery{
        ReportID:   id,
        Config:     report.Config,
        TimeRange:  s.getTimeRange(report),
    })
    if err != nil {
        return nil, err
    }
    
    // 3. 生成报表
    result, err := s.generateReport(report, data)
    if err != nil {
        return nil, err
    }
    
    // 4. 保存报表
    if err := s.saveReport(result); err != nil {
        return nil, err
    }
    
    // 5. 发送通知
    if len(report.Recipients) > 0 {
        go s.sendNotification(report, result)
    }
    
    return result, nil
}
```

#### 6.2.3 数据查询实现
```go
// internal/query/service.go
type QueryService struct {
    db          *gorm.DB
    cache       *redis.Client
    clickhouse  *clickhouse.Conn
    logger      *zap.Logger
}

func (s *QueryService) Query(ctx context.Context, query *DataQuery) (*QueryResult, error) {
    // 1. 构建查询
    sql := s.buildQuery(query)
    
    // 2. 执行查询
    rows, err := s.clickhouse.Query(ctx, sql)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    // 3. 处理结果
    result := &QueryResult{
        ReportID:   query.ReportID,
        TimeRange:  query.TimeRange,
        Dimensions: query.Config.Dimensions,
        Metrics:    query.Config.Metrics,
        Data:       make([]*QueryRow, 0),
    }
    
    for rows.Next() {
        row := &QueryRow{
            Dimensions: make(map[string]string),
            Metrics:    make(map[string]float64),
        }
        
        if err := rows.Scan(&row); err != nil {
            return nil, err
        }
        
        result.Data = append(result.Data, row)
    }
    
    // 4. 处理分页
    if query.Page != nil {
        result.Total = int64(len(result.Data))
        result.Page = s.processPagination(result.Data, query.Page)
    }
    
    return result, nil
}

func (s *QueryService) GetMetrics(ctx context.Context, query *MetricsQuery) (*MetricsResult, error) {
    // 1. 构建查询
    sql := s.buildMetricsQuery(query)
    
    // 2. 执行查询
    rows, err := s.clickhouse.Query(ctx, sql)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    // 3. 处理结果
    result := &MetricsResult{
        Metrics: make(map[string]float64),
    }
    
    for rows.Next() {
        var metric string
        var value float64
        
        if err := rows.Scan(&metric, &value); err != nil {
            return nil, err
        }
        
        result.Metrics[metric] = value
    }
    
    return result, nil
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
// internal/report/service_test.go
func TestReportService_Create(t *testing.T) {
    tests := []struct {
        name    string
        req     *CreateReportRequest
        want    *Report
        wantErr bool
    }{
        // 测试用例
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            s := NewReportService(db, cache, storage, query, logger)
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
// internal/report/service_benchmark_test.go
func BenchmarkReportService_Create(b *testing.B) {
    s := NewReportService(db, cache, storage, query, logger)
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
RUN go build -o report ./cmd/report

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/report .
COPY --from=builder /app/config ./config
EXPOSE 8080
CMD ["./report"]
```

### 10.2 Kubernetes部署
```yaml
# k8s/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: report
spec:
  replicas: 3
  selector:
    matchLabels:
      app: report
  template:
    metadata:
      labels:
        app: report
    spec:
      containers:
      - name: report
        image: report:latest
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
          name: report-config
```

### 10.3 监控配置
```yaml
# prometheus/service-monitor.yaml
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: report
spec:
  selector:
    matchLabels:
      app: report
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
1. 数据采集工具
2. 报表管理工具
3. 数据查询工具
4. 导出管理工具 