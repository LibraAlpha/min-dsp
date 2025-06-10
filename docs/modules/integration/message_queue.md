# 消息队列模块设计文档

## 1. 模块概述

消息队列模块是DSP系统集成层的核心组件之一，负责系统内部各模块之间的异步通信、消息解耦、流量削峰和任务分发等功能。该模块需要支持高吞吐量、低延迟、高可靠性的消息处理，同时确保消息的顺序性、可靠投递和消息追踪。

## 2. 功能列表

### 2.1 核心功能
1. 消息发布
2. 消息订阅
3. 消息存储
4. 消息路由
5. 消息重试

### 2.2 扩展功能
1. 消息监控
2. 消息追踪
3. 消息过滤
4. 消息转换
5. 消息压缩

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
│ 消息发布    │ │ 消息订阅  │  │ 消息存储 │ │ 消息路由 │
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
- 消息发布：消息创建、消息验证、消息发送等
- 消息订阅：主题订阅、消息接收、消息处理等
- 消息存储：消息持久化、消息索引、消息清理等
- 消息路由：消息分发、负载均衡、消息过滤等

#### 3.2.3 数据层
- 消息存储
- 消息缓存
- 消息同步
- 消息备份

#### 3.2.4 存储层
- 消息队列
- 分布式缓存
- 关系型数据库
- 文件存储

## 4. 接口定义

### 4.1 外部接口
```go
// 消息队列服务接口
type MessageQueue interface {
    // 消息发布
    Publish(ctx context.Context, req *PublishRequest) (*PublishResponse, error)
    BatchPublish(ctx context.Context, req *BatchPublishRequest) (*BatchPublishResponse, error)
    
    // 消息订阅
    Subscribe(ctx context.Context, req *SubscribeRequest) (*SubscribeResponse, error)
    Unsubscribe(ctx context.Context, req *UnsubscribeRequest) error
    
    // 消息管理
    GetMessage(ctx context.Context, req *GetMessageRequest) (*MessageResponse, error)
    DeleteMessage(ctx context.Context, req *DeleteMessageRequest) error
    ListMessages(ctx context.Context, req *ListMessagesRequest) (*ListMessagesResponse, error)
    
    // 主题管理
    CreateTopic(ctx context.Context, req *CreateTopicRequest) (*TopicResponse, error)
    DeleteTopic(ctx context.Context, req *DeleteTopicRequest) error
    ListTopics(ctx context.Context, req *ListTopicsRequest) (*ListTopicsResponse, error)
}

// 消息发布请求
type PublishRequest struct {
    Topic       string            `json:"topic"`
    Message     *Message          `json:"message"`
    Properties  map[string]string `json:"properties"`
    Metadata    map[string]string `json:"metadata"`
}

// 消息订阅请求
type SubscribeRequest struct {
    Topic       string            `json:"topic"`
    Group       string            `json:"group"`
    Handler     MessageHandler    `json:"handler"`
    Properties  map[string]string `json:"properties"`
    Metadata    map[string]string `json:"metadata"`
}

// 获取消息请求
type GetMessageRequest struct {
    Topic       string            `json:"topic"`
    MessageID   string            `json:"message_id"`
    Properties  map[string]string `json:"properties"`
}

// 创建主题请求
type CreateTopicRequest struct {
    Name        string            `json:"name"`
    Type        TopicType         `json:"type"`
    Properties  map[string]string `json:"properties"`
    Metadata    map[string]string `json:"metadata"`
}
```

### 4.2 内部接口
```go
// 消息存储服务接口
type StorageService interface {
    // 存储消息
    Store(ctx context.Context, message *Message) error
    
    // 获取消息
    Get(ctx context.Context, id string) (*Message, error)
    
    // 删除消息
    Delete(ctx context.Context, id string) error
    
    // 列出消息
    List(ctx context.Context, filter *MessageFilter) ([]*Message, error)
}

// 消息路由服务接口
type RouterService interface {
    // 路由消息
    Route(ctx context.Context, message *Message) error
    
    // 添加路由
    AddRoute(ctx context.Context, route *Route) error
    
    // 删除路由
    DeleteRoute(ctx context.Context, routeID string) error
    
    // 获取路由
    GetRoute(ctx context.Context, routeID string) (*Route, error)
}

// 消息处理服务接口
type ProcessorService interface {
    // 处理消息
    Process(ctx context.Context, message *Message) error
    
    // 重试消息
    Retry(ctx context.Context, message *Message) error
    
    // 死信处理
    HandleDeadLetter(ctx context.Context, message *Message) error
    
    // 获取统计
    GetStats(ctx context.Context, topic string) (*ProcessStats, error)
}

// 消息追踪服务接口
type TraceService interface {
    // 追踪消息
    Trace(ctx context.Context, message *Message) error
    
    // 获取追踪
    GetTrace(ctx context.Context, messageID string) (*Trace, error)
    
    // 获取统计
    GetStats(ctx context.Context, filter *TraceFilter) (*TraceStats, error)
}
```

## 5. 数据结构

### 5.1 消息数据
```go
// 消息类型
type MessageType string

const (
    MessageTypeNormal   MessageType = "normal"
    MessageTypeOrder    MessageType = "order"
    MessageTypeDelay    MessageType = "delay"
    MessageTypeDead     MessageType = "dead"
)

// 消息信息
type Message struct {
    ID          string            `json:"id"`
    Topic       string            `json:"topic"`
    Type        MessageType       `json:"type"`
    Content     []byte            `json:"content"`
    Properties  map[string]string `json:"properties"`
    Status      MessageStatus     `json:"status"`
    CreatedAt   time.Time         `json:"created_at"`
    UpdatedAt   time.Time         `json:"updated_at"`
    Metadata    map[string]string `json:"metadata"`
}

// 消息状态
type MessageStatus struct {
    Status      string            `json:"status"`
    Retries     int               `json:"retries"`
    NextRetry   time.Time         `json:"next_retry"`
    Error       string            `json:"error"`
    UpdatedAt   time.Time         `json:"updated_at"`
}
```

### 5.2 主题数据
```go
// 主题类型
type TopicType string

const (
    TopicTypeNormal   TopicType = "normal"
    TopicTypeOrder    TopicType = "order"
    TopicTypeDelay    TopicType = "delay"
)

// 主题信息
type Topic struct {
    ID          string            `json:"id"`
    Name        string            `json:"name"`
    Type        TopicType         `json:"type"`
    Properties  *TopicProperties  `json:"properties"`
    Status      TopicStatus       `json:"status"`
    CreatedAt   time.Time         `json:"created_at"`
    UpdatedAt   time.Time         `json:"updated_at"`
    Metadata    map[string]string `json:"metadata"`
}

// 主题属性
type TopicProperties struct {
    Retention   int               `json:"retention"`
    MaxSize     int64             `json:"max_size"`
    MaxMessages int64             `json:"max_messages"`
    Compression bool              `json:"compression"`
}

// 主题状态
type TopicStatus struct {
    Status      string            `json:"status"`
    Messages    int64             `json:"messages"`
    Size        int64             `json:"size"`
    UpdatedAt   time.Time         `json:"updated_at"`
}
```

### 5.3 订阅数据
```go
// 订阅信息
type Subscription struct {
    ID          string            `json:"id"`
    Topic       string            `json:"topic"`
    Group       string            `json:"group"`
    Handler     string            `json:"handler"`
    Properties  *SubProperties    `json:"properties"`
    Status      SubStatus         `json:"status"`
    CreatedAt   time.Time         `json:"created_at"`
    UpdatedAt   time.Time         `json:"updated_at"`
    Metadata    map[string]string `json:"metadata"`
}

// 订阅属性
type SubProperties struct {
    BatchSize   int               `json:"batch_size"`
    Timeout     int               `json:"timeout"`
    Retries     int               `json:"retries"`
    Backoff     int               `json:"backoff"`
}

// 订阅状态
type SubStatus struct {
    Status      string            `json:"status"`
    Offset      int64             `json:"offset"`
    Lag         int64             `json:"lag"`
    UpdatedAt   time.Time         `json:"updated_at"`
}
```

### 5.4 追踪数据
```go
// 追踪信息
type Trace struct {
    ID          string            `json:"id"`
    MessageID   string            `json:"message_id"`
    Topic       string            `json:"topic"`
    Events      []*TraceEvent     `json:"events"`
    Status      TraceStatus       `json:"status"`
    CreatedAt   time.Time         `json:"created_at"`
    UpdatedAt   time.Time         `json:"updated_at"`
}

// 追踪事件
type TraceEvent struct {
    Type        string            `json:"type"`
    Status      string            `json:"status"`
    Details     string            `json:"details"`
    Timestamp   time.Time         `json:"timestamp"`
}

// 追踪状态
type TraceStatus struct {
    Status      string            `json:"status"`
    Events      int               `json:"events"`
    Duration    int64             `json:"duration"`
    UpdatedAt   time.Time         `json:"updated_at"`
}
```

## 6. 实现路径

### 6.1 基础框架搭建
1. 项目初始化
   ```bash
   # 创建项目目录
   mkdir -p cmd/queue
   mkdir -p internal/{publisher,subscriber,storage,router}
   mkdir -p pkg/{api,service,utils}
   
   # 初始化 Go 模块
   go mod init min-dsp
   
   # 添加依赖
   go get -u github.com/gin-gonic/gin
   go get -u gorm.io/gorm
   go get -u github.com/go-redis/redis/v8
   go get -u github.com/apache/rocketmq-client-go/v2
   go get -u go.uber.org/zap
   ```

2. 配置文件
   ```yaml
   # config/queue.yaml
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
   
   rocketmq:
     name_server:
       - "localhost:9876"
     group_name: "dsp"
     instance_name: "queue"
     access_key: "******"
     secret_key: "******"
   
   storage:
     type: "rocketmq"
     topic_prefix: "dsp_"
     group_prefix: "dsp_"
   
   message:
     retention: 7
     max_size: 1048576
     max_messages: 1000000
     compression: true
   
   monitoring:
     log_level: "info"
     metrics_port: 9090
     alert_threshold: 0.9
   ```

### 6.2 核心功能实现

#### 6.2.1 消息发布实现
```go
// internal/publisher/service.go
type PublisherService struct {
    db          *gorm.DB
    cache       *redis.Client
    mq          *rocketmq.Producer
    storage     StorageService
    router      RouterService
    trace       TraceService
    logger      *zap.Logger
}

func (s *PublisherService) Publish(ctx context.Context, req *PublishRequest) (*PublishResponse, error) {
    // 1. 验证请求
    if err := s.validatePublishRequest(req); err != nil {
        return nil, err
    }
    
    // 2. 创建消息
    message := &Message{
        Topic:      req.Topic,
        Type:       MessageTypeNormal,
        Content:    req.Message.Content,
        Properties: req.Properties,
        Status:     MessageStatus{Status: "created"},
    }
    
    // 3. 存储消息
    if err := s.storage.Store(ctx, message); err != nil {
        return nil, err
    }
    
    // 4. 发送消息
    if err := s.sendMessage(ctx, message); err != nil {
        return nil, err
    }
    
    // 5. 追踪消息
    if err := s.trace.Trace(ctx, message); err != nil {
        s.logger.Warn("failed to trace message", zap.Error(err))
    }
    
    return &PublishResponse{MessageID: message.ID}, nil
}

func (s *PublisherService) sendMessage(ctx context.Context, message *Message) error {
    // 1. 构建消息
    msg := &rocketmq.Message{
        Topic: message.Topic,
        Body:  message.Content,
    }
    
    // 2. 设置属性
    for k, v := range message.Properties {
        msg.WithProperty(k, v)
    }
    
    // 3. 发送消息
    result, err := s.mq.SendSync(ctx, msg)
    if err != nil {
        return err
    }
    
    // 4. 更新状态
    message.Status.Status = "sent"
    message.Status.UpdatedAt = time.Now()
    
    return s.storage.Store(ctx, message)
}
```

#### 6.2.2 消息订阅实现
```go
// internal/subscriber/service.go
type SubscriberService struct {
    db          *gorm.DB
    cache       *redis.Client
    mq          *rocketmq.PushConsumer
    storage     StorageService
    processor   ProcessorService
    trace       TraceService
    logger      *zap.Logger
}

func (s *SubscriberService) Subscribe(ctx context.Context, req *SubscribeRequest) (*SubscribeResponse, error) {
    // 1. 验证请求
    if err := s.validateSubscribeRequest(req); err != nil {
        return nil, err
    }
    
    // 2. 创建订阅
    sub := &Subscription{
        Topic:      req.Topic,
        Group:      req.Group,
        Handler:    req.Handler,
        Properties: req.Properties,
        Status:     SubStatus{Status: "active"},
    }
    
    // 3. 存储订阅
    if err := s.storage.Store(ctx, sub); err != nil {
        return nil, err
    }
    
    // 4. 启动消费
    go s.startConsume(ctx, sub)
    
    return &SubscribeResponse{Subscription: sub}, nil
}

func (s *SubscriberService) startConsume(ctx context.Context, sub *Subscription) {
    // 1. 订阅主题
    err := s.mq.Subscribe(sub.Topic, sub.Group, func(ctx context.Context, msgs ...*rocketmq.MessageExt) error {
        for _, msg := range msgs {
            // 2. 处理消息
            if err := s.processMessage(ctx, msg, sub); err != nil {
                s.logger.Error("failed to process message", zap.Error(err))
                continue
            }
        }
        return nil
    })
    
    if err != nil {
        s.logger.Error("failed to subscribe topic", zap.Error(err))
        return
    }
    
    // 3. 启动消费者
    if err := s.mq.Start(); err != nil {
        s.logger.Error("failed to start consumer", zap.Error(err))
        return
    }
    
    // 4. 等待退出
    <-ctx.Done()
    s.mq.Shutdown()
}

func (s *SubscriberService) processMessage(ctx context.Context, msg *rocketmq.MessageExt, sub *Subscription) error {
    // 1. 获取消息
    message, err := s.storage.Get(ctx, msg.MsgId)
    if err != nil {
        return err
    }
    
    // 2. 处理消息
    if err := s.processor.Process(ctx, message); err != nil {
        // 3. 重试消息
        if message.Status.Retries < sub.Properties.Retries {
            return s.processor.Retry(ctx, message)
        }
        
        // 4. 处理死信
        return s.processor.HandleDeadLetter(ctx, message)
    }
    
    // 5. 更新状态
    message.Status.Status = "processed"
    message.Status.UpdatedAt = time.Now()
    
    return s.storage.Store(ctx, message)
}
```

#### 6.2.3 消息路由实现
```go
// internal/router/service.go
type RouterService struct {
    db          *gorm.DB
    cache       *redis.Client
    logger      *zap.Logger
}

func (s *RouterService) Route(ctx context.Context, message *Message) error {
    // 1. 获取路由
    routes, err := s.getRoutes(ctx, message.Topic)
    if err != nil {
        return err
    }
    
    // 2. 过滤路由
    routes = s.filterRoutes(routes, message)
    
    // 3. 执行路由
    for _, route := range routes {
        if err := s.executeRoute(ctx, route, message); err != nil {
            s.logger.Error("failed to execute route", zap.Error(err))
            continue
        }
    }
    
    return nil
}

func (s *RouterService) filterRoutes(routes []*Route, message *Message) []*Route {
    filtered := make([]*Route, 0)
    
    for _, route := range routes {
        // 1. 检查条件
        if !s.matchCondition(route.Condition, message) {
            continue
        }
        
        // 2. 检查过滤器
        if !s.matchFilter(route.Filter, message) {
            continue
        }
        
        filtered = append(filtered, route)
    }
    
    return filtered
}

func (s *RouterService) executeRoute(ctx context.Context, route *Route, message *Message) error {
    // 1. 转换消息
    if route.Transform != nil {
        if err := s.transformMessage(message, route.Transform); err != nil {
            return err
        }
    }
    
    // 2. 发送消息
    if err := s.sendMessage(ctx, route, message); err != nil {
        return err
    }
    
    // 3. 更新统计
    return s.updateStats(ctx, route)
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
// internal/publisher/service_test.go
func TestPublisherService_Publish(t *testing.T) {
    tests := []struct {
        name    string
        req     *PublishRequest
        want    *PublishResponse
        wantErr bool
    }{
        // 测试用例
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            s := NewPublisherService(db, cache, mq, storage, router, trace, logger)
            got, err := s.Publish(context.Background(), tt.req)
            if (err != nil) != tt.wantErr {
                t.Errorf("Publish() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("Publish() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

### 9.2 性能测试
```go
// internal/publisher/service_benchmark_test.go
func BenchmarkPublisherService_Publish(b *testing.B) {
    s := NewPublisherService(db, cache, mq, storage, router, trace, logger)
    req := createTestRequest()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, err := s.Publish(context.Background(), req)
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
RUN go build -o queue ./cmd/queue

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/queue .
COPY --from=builder /app/config ./config
EXPOSE 8080
CMD ["./queue"]
```

### 10.2 Kubernetes部署
```yaml
# k8s/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: queue
spec:
  replicas: 3
  selector:
    matchLabels:
      app: queue
  template:
    metadata:
      labels:
        app: queue
    spec:
      containers:
      - name: queue
        image: queue:latest
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
          name: queue-config
```

### 10.3 监控配置
```yaml
# prometheus/service-monitor.yaml
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: queue
spec:
  selector:
    matchLabels:
      app: queue
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
1. 消息管理工具
2. 消息监控工具
3. 消息诊断工具
4. 消息备份工具 