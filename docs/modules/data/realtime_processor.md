# 实时数据处理模块设计文档

## 1. 模块概述

实时数据处理模块是DSP系统数据处理层的核心组件之一，负责处理系统中的实时数据流，包括广告请求数据、用户行为数据、竞价数据等。该模块需要具备高吞吐、低延迟的数据处理能力，通过流式处理技术对数据进行实时分析、转换和聚合，为决策引擎层提供及时的数据支持。

## 2. 功能列表

### 2.1 核心功能
1. 实时数据接入
2. 数据流处理
3. 实时数据分析
4. 数据转换聚合
5. 实时数据分发

### 2.2 扩展功能
1. 数据质量控制
2. 实时监控告警
3. 数据回溯处理
4. 处理性能优化
5. 容错恢复机制

## 3. 详细设计

### 3.1 系统架构
```
                    ┌─────────────┐
                    │  数据接入层  │
                    └──────┬──────┘
                           │
                    ┌──────▼──────┐
                    │  消息队列   │
                    └──────┬──────┘
                           │
         ┌─────────┬───────┴───────┬─────────┐
         │         │               │         │
┌────────▼───┐ ┌───▼──────┐  ┌────▼────┐ ┌──▼──────┐
│ 请求处理器  │ │ 行为处理器 │  │ 竞价处理器 │ │ 统计处理器 │
└───────┬────┘ └─────┬────┘  └────┬────┘ └────┬────┘
        │            │            │           │
        └────────────┼────────────┼───────────┘
                     │            │
               ┌─────▼────────────▼─────┐
               │      数据聚合层         │
               └─────────────┬──────────┘
                             │
                     ┌───────▼───────┐
                     │  数据分发层   │
                     └───────────────┘
```

### 3.2 组件说明

#### 3.2.1 数据接入层
- 数据源接入管理
- 数据格式验证
- 数据预处理
- 负载均衡

#### 3.2.2 消息队列
- Kafka集群管理
- 消息分区管理
- 消息持久化
- 消息重试机制

#### 3.2.3 处理器组件
- 请求处理器：处理广告请求数据
- 行为处理器：处理用户行为数据
- 竞价处理器：处理实时竞价数据
- 统计处理器：处理实时统计数据

#### 3.2.4 数据聚合层
- 数据聚合计算
- 实时指标统计
- 数据关联分析
- 结果缓存管理

#### 3.2.5 数据分发层
- 数据路由分发
- 订阅管理
- 数据推送
- 负载均衡

## 4. 接口定义

### 4.1 外部接口
```go
// 实时处理服务接口
type RealtimeProcessor interface {
    // 处理数据流
    ProcessStream(ctx context.Context, stream *DataStream) (*ProcessResult, error)
    
    // 获取处理状态
    GetStatus(ctx context.Context) (*ProcessStatus, error)
    
    // 更新处理配置
    UpdateConfig(ctx context.Context, config *ProcessConfig) error
    
    // 获取处理统计
    GetStats(ctx context.Context) (*ProcessStats, error)
}

// 数据流
type DataStream struct {
    ID          string            `json:"id"`
    Type        StreamType        `json:"type"`
    Source      string            `json:"source"`
    Data        []byte            `json:"data"`
    Timestamp   time.Time         `json:"timestamp"`
    Metadata    map[string]string `json:"metadata"`
}

// 处理结果
type ProcessResult struct {
    ID          string            `json:"id"`
    Status      string            `json:"status"`
    Data        []byte            `json:"data"`
    Timestamp   time.Time         `json:"timestamp"`
    Metrics     *ProcessMetrics   `json:"metrics"`
}
```

### 4.2 内部接口
```go
// 数据处理器接口
type DataProcessor interface {
    // 处理数据
    Process(ctx context.Context, data *ProcessData) (*ProcessResult, error)
    
    // 初始化处理器
    Init(ctx context.Context, config *ProcessorConfig) error
    
    // 关闭处理器
    Close(ctx context.Context) error
}

// 数据聚合器接口
type DataAggregator interface {
    // 聚合数据
    Aggregate(ctx context.Context, data []*ProcessData) (*AggregateResult, error)
    
    // 更新聚合规则
    UpdateRules(ctx context.Context, rules []*AggregateRule) error
    
    // 获取聚合结果
    GetResults(ctx context.Context, query *AggregateQuery) (*AggregateResult, error)
}

// 数据分发器接口
type DataDistributor interface {
    // 分发数据
    Distribute(ctx context.Context, data *DistributeData) error
    
    // 添加订阅者
    AddSubscriber(ctx context.Context, subscriber *Subscriber) error
    
    // 移除订阅者
    RemoveSubscriber(ctx context.Context, subscriberID string) error
}
```

## 5. 数据结构

### 5.1 处理数据
```go
// 处理数据类型
type ProcessDataType string

const (
    ProcessTypeRequest  ProcessDataType = "request"
    ProcessTypeBehavior ProcessDataType = "behavior"
    ProcessTypeBid      ProcessDataType = "bid"
    ProcessTypeStats    ProcessDataType = "stats"
)

// 处理数据
type ProcessData struct {
    Type        ProcessDataType    `json:"type"`
    ID          string             `json:"id"`
    Data        any                `json:"data"`
    Timestamp   time.Time          `json:"timestamp"`
    Metadata    map[string]string  `json:"metadata"`
}

// 处理配置
type ProcessConfig struct {
    Type        ProcessDataType    `json:"type"`
    BatchSize   int                `json:"batch_size"`
    Timeout     time.Duration      `json:"timeout"`
    RetryCount  int                `json:"retry_count"`
    Options     map[string]any     `json:"options"`
}
```

### 5.2 聚合数据
```go
// 聚合规则
type AggregateRule struct {
    ID          string             `json:"id"`
    Type        string             `json:"type"`
    Fields      []string           `json:"fields"`
    Window      time.Duration      `json:"window"`
    Function    string             `json:"function"`
    Filter      string             `json:"filter"`
}

// 聚合结果
type AggregateResult struct {
    RuleID      string             `json:"rule_id"`
    Data        map[string]any     `json:"data"`
    Count       int64              `json:"count"`
    StartTime   time.Time          `json:"start_time"`
    EndTime     time.Time          `json:"end_time"`
}

// 处理指标
type ProcessMetrics struct {
    Processed   int64              `json:"processed"`
    Failed      int64              `json:"failed"`
    Latency     time.Duration      `json:"latency"`
    Throughput  float64            `json:"throughput"`
    UpdatedAt   time.Time          `json:"updated_at"`
}
```

## 6. 实现路径

### 6.1 基础框架搭建
1. 项目初始化
   ```bash
   # 创建项目目录
   mkdir -p cmd/processor
   mkdir -p internal/{ingestion,process,aggregate,distribute}
   mkdir -p pkg/{kafka,metrics,utils}
   
   # 初始化 Go 模块
   go mod init min-dsp
   
   # 添加依赖
   go get -u github.com/Shopify/sarama
   go get -u github.com/prometheus/client_golang
   go get -u go.uber.org/zap
   ```

2. 配置文件
   ```yaml
   # config/processor.yaml
   kafka:
     brokers:
       - "localhost:9092"
       - "localhost:9093"
     topics:
       request: "dsp.requests"
       behavior: "dsp.behaviors"
       bid: "dsp.bids"
       stats: "dsp.stats"
     consumer:
       group: "dsp-processor"
       batch_size: 1000
       batch_timeout: "100ms"
   
   processor:
     workers: 10
     batch_size: 1000
     timeout: "1s"
     retry_count: 3
     metrics:
       port: 9090
       path: "/metrics"
   
   aggregator:
     rules:
       - id: "request_stats"
         type: "count"
         window: "1m"
         fields: ["source", "type"]
       - id: "bid_stats"
         type: "avg"
         window: "5m"
         fields: ["price", "win"]
   
   distributor:
     subscribers:
       - id: "rtb_engine"
         topics: ["request_stats", "bid_stats"]
       - id: "monitor"
         topics: ["*"]
   
   monitoring:
     log_level: "info"
     alert_threshold: 0.9
     metrics_interval: "15s"
   ```

### 6.2 核心功能实现

#### 6.2.1 数据接入实现
```go
// internal/ingestion/ingestion.go
type DataIngestion struct {
    config      *Config
    kafka       *KafkaClient
    processors  map[string]DataProcessor
    metrics     *Metrics
}

func (i *DataIngestion) Start(ctx context.Context) error {
    // 1. 初始化Kafka客户端
    if err := i.kafka.Init(ctx); err != nil {
        return err
    }
    
    // 2. 启动消费者
    for _, topic := range i.config.Kafka.Topics {
        go i.consumeTopic(ctx, topic)
    }
    
    // 3. 启动指标收集
    go i.collectMetrics(ctx)
    
    return nil
}

func (i *DataIngestion) consumeTopic(ctx context.Context, topic string) {
    for {
        select {
        case <-ctx.Done():
            return
        default:
            // 1. 获取消息批次
            messages, err := i.kafka.ConsumeBatch(ctx, topic)
            if err != nil {
                i.metrics.RecordError("consume_error", err)
                continue
            }
            
            // 2. 处理消息
            if err := i.processMessages(ctx, messages); err != nil {
                i.metrics.RecordError("process_error", err)
            }
        }
    }
}
```

#### 6.2.2 处理器实现
```go
// internal/process/processor.go
type Processor struct {
    config      *Config
    aggregator  DataAggregator
    distributor DataDistributor
    metrics     *Metrics
}

func (p *Processor) Process(ctx context.Context, data *ProcessData) (*ProcessResult, error) {
    // 1. 数据预处理
    processed, err := p.preprocess(ctx, data)
    if err != nil {
        return nil, err
    }
    
    // 2. 数据聚合
    aggregated, err := p.aggregator.Aggregate(ctx, []*ProcessData{processed})
    if err != nil {
        return nil, err
    }
    
    // 3. 数据分发
    if err := p.distributor.Distribute(ctx, &DistributeData{
        Data: aggregated,
        Type: data.Type,
    }); err != nil {
        return nil, err
    }
    
    // 4. 记录指标
    p.metrics.RecordProcess(data.Type, time.Since(data.Timestamp))
    
    return &ProcessResult{
        Status: "success",
        Data:   processed.Data,
        Metrics: &ProcessMetrics{
            Processed: 1,
            Latency:  time.Since(data.Timestamp),
        },
    }, nil
}
```

#### 6.2.3 聚合器实现
```go
// internal/aggregate/aggregator.go
type Aggregator struct {
    config      *Config
    rules       map[string]*AggregateRule
    cache       *Cache
    metrics     *Metrics
}

func (a *Aggregator) Aggregate(ctx context.Context, data []*ProcessData) (*AggregateResult, error) {
    // 1. 获取聚合规则
    rules := a.getRules(data[0].Type)
    
    // 2. 并发聚合
    var wg sync.WaitGroup
    results := make(chan *AggregateResult, len(rules))
    errors := make(chan error, len(rules))
    
    for _, rule := range rules {
        wg.Add(1)
        go func(r *AggregateRule) {
            defer wg.Done()
            if result, err := a.aggregateRule(ctx, r, data); err != nil {
                errors <- err
            } else {
                results <- result
            }
        }(rule)
    }
    
    // 3. 等待完成
    wg.Wait()
    close(results)
    close(errors)
    
    // 4. 处理结果
    return a.processResults(results, errors)
}

func (a *Aggregator) aggregateRule(ctx context.Context, rule *AggregateRule, data []*ProcessData) (*AggregateResult, error) {
    // 1. 数据过滤
    filtered := a.filterData(rule, data)
    
    // 2. 数据分组
    groups := a.groupData(rule, filtered)
    
    // 3. 聚合计算
    return a.calculateAggregate(rule, groups)
}
```

### 6.3 中间件实现

#### 6.3.1 监控中间件
```go
// internal/middleware/monitor.go
func MonitorMiddleware(metrics *Metrics) gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        
        c.Next()
        
        // 记录指标
        metrics.RecordAPI(
            c.Request.URL.Path,
            c.Writer.Status(),
            time.Since(start),
        )
    }
}
```

#### 6.3.2 限流中间件
```go
// internal/middleware/ratelimit.go
func RateLimitMiddleware(limiter *RateLimiter) gin.HandlerFunc {
    return func(c *gin.Context) {
        // 1. 获取限流器
        limit := limiter.GetLimiter(c.Request.URL.Path)
        
        // 2. 检查限流
        if !limit.Allow() {
            c.AbortWithStatus(http.StatusTooManyRequests)
            return
        }
        
        c.Next()
    }
}
```

## 7. 性能考虑

### 7.1 处理优化
1. 批量处理
2. 并行处理
3. 内存优化
4. 缓存策略

### 7.2 资源优化
1. CPU使用优化
2. 内存使用优化
3. 网络使用优化
4. 磁盘使用优化

### 7.3 性能监控
1. 处理性能
2. 资源使用
3. 系统负载
4. 服务质量

## 8. 安全考虑

### 8.1 数据安全
1. 数据加密
2. 访问控制
3. 数据验证
4. 审计日志

### 8.2 运行安全
1. 服务认证
2. 权限管理
3. 异常处理
4. 故障恢复

## 9. 测试方案

### 9.1 单元测试
```go
// internal/process/processor_test.go
func TestProcessor_Process(t *testing.T) {
    tests := []struct {
        name    string
        data    *ProcessData
        want    *ProcessResult
        wantErr bool
    }{
        // 测试用例
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            p := NewProcessor(config)
            got, err := p.Process(context.Background(), tt.data)
            if (err != nil) != tt.wantErr {
                t.Errorf("Process() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("Process() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

### 9.2 性能测试
```go
// internal/process/processor_benchmark_test.go
func BenchmarkProcessor_Process(b *testing.B) {
    processor := NewProcessor(config)
    data := createTestData()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, err := processor.Process(context.Background(), data)
        if err != nil {
            b.Fatal(err)
        }
    }
}
```

### 9.3 集成测试
1. 端到端测试
2. 性能测试
3. 压力测试
4. 故障测试

## 10. 部署方案

### 10.1 容器化部署
```dockerfile
# Dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o processor ./cmd/processor

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/processor .
COPY --from=builder /app/config ./config
EXPOSE 8080
CMD ["./processor"]
```

### 10.2 Kubernetes部署
```yaml
# k8s/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: processor
spec:
  replicas: 3
  selector:
    matchLabels:
      app: processor
  template:
    metadata:
      labels:
        app: processor
    spec:
      containers:
      - name: processor
        image: processor:latest
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
          name: processor-config
```

### 10.3 监控配置
```yaml
# prometheus/service-monitor.yaml
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: processor
spec:
  selector:
    matchLabels:
      app: processor
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
1. 数据处理工具
2. 性能分析工具
3. 故障诊断工具
4. 监控管理工具 