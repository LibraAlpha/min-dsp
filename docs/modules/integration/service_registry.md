# 服务注册与发现模块设计文档

## 1. 模块概述

服务注册与发现模块是DSP系统集成层的核心组件之一，负责管理微服务架构中的服务注册、服务发现、健康检查和服务路由等功能。该模块需要支持服务的自动注册、服务状态监控、负载均衡、服务熔断等功能，同时确保服务的高可用性和可靠性。

## 2. 功能列表

### 2.1 核心功能
1. 服务注册
2. 服务发现
3. 健康检查
4. 负载均衡
5. 服务路由

### 2.2 扩展功能
1. 服务监控
2. 服务熔断
3. 服务限流
4. 服务降级
5. 服务追踪

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
│ 服务注册    │ │ 服务发现  │  │ 健康检查 │ │ 负载均衡 │
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
- 服务注册：服务注册、服务更新、服务注销等
- 服务发现：服务查询、服务订阅、服务通知等
- 健康检查：状态检查、心跳检测、故障检测等
- 负载均衡：负载策略、服务选择、权重调整等

#### 3.2.3 数据层
- 服务存储
- 服务缓存
- 服务同步
- 服务备份

#### 3.2.4 存储层
- 服务注册中心
- 分布式缓存
- 关系型数据库
- 文件存储

## 4. 接口定义

### 4.1 外部接口
```go
// 服务注册与发现接口
type ServiceRegistry interface {
    // 服务注册
    Register(ctx context.Context, req *RegisterRequest) (*ServiceResponse, error)
    Deregister(ctx context.Context, req *DeregisterRequest) error
    Update(ctx context.Context, req *UpdateRequest) (*ServiceResponse, error)
    
    // 服务发现
    Discover(ctx context.Context, req *DiscoverRequest) (*DiscoverResponse, error)
    Subscribe(ctx context.Context, req *SubscribeRequest) (*SubscribeResponse, error)
    Unsubscribe(ctx context.Context, req *UnsubscribeRequest) error
    
    // 健康检查
    HealthCheck(ctx context.Context, req *HealthCheckRequest) (*HealthCheckResponse, error)
    GetHealth(ctx context.Context, req *GetHealthRequest) (*HealthResponse, error)
    
    // 负载均衡
    GetService(ctx context.Context, req *GetServiceRequest) (*ServiceResponse, error)
    ListServices(ctx context.Context, req *ListServicesRequest) (*ListServicesResponse, error)
}

// 服务注册请求
type RegisterRequest struct {
    Service     *Service         `json:"service"`
    Instance    *Instance        `json:"instance"`
    Metadata    map[string]string `json:"metadata"`
}

// 服务发现请求
type DiscoverRequest struct {
    ServiceName string            `json:"service_name"`
    Tags        []string          `json:"tags"`
    Version     string            `json:"version"`
    Metadata    map[string]string `json:"metadata"`
}

// 健康检查请求
type HealthCheckRequest struct {
    ServiceID   string            `json:"service_id"`
    InstanceID  string            `json:"instance_id"`
    Status      HealthStatus      `json:"status"`
    Metadata    map[string]string `json:"metadata"`
}

// 获取服务请求
type GetServiceRequest struct {
    ServiceName string            `json:"service_name"`
    Strategy    LoadStrategy      `json:"strategy"`
    Metadata    map[string]string `json:"metadata"`
}
```

### 4.2 内部接口
```go
// 服务存储接口
type StorageService interface {
    // 存储服务
    Store(ctx context.Context, service *Service) error
    
    // 获取服务
    Get(ctx context.Context, id string) (*Service, error)
    
    // 删除服务
    Delete(ctx context.Context, id string) error
    
    // 列出服务
    List(ctx context.Context, filter *ServiceFilter) ([]*Service, error)
}

// 服务发现接口
type DiscoveryService interface {
    // 发现服务
    Discover(ctx context.Context, request *DiscoverRequest) ([]*Service, error)
    
    // 订阅服务
    Subscribe(ctx context.Context, request *SubscribeRequest) error
    
    // 取消订阅
    Unsubscribe(ctx context.Context, request *UnsubscribeRequest) error
    
    // 通知更新
    Notify(ctx context.Context, service *Service) error
}

// 健康检查接口
type HealthService interface {
    // 检查健康
    Check(ctx context.Context, service *Service) (*HealthResult, error)
    
    // 更新状态
    UpdateStatus(ctx context.Context, service *Service, status HealthStatus) error
    
    // 获取状态
    GetStatus(ctx context.Context, service *Service) (*HealthStatus, error)
}

// 负载均衡接口
type LoadBalancer interface {
    // 选择服务
    Select(ctx context.Context, services []*Service, strategy LoadStrategy) (*Service, error)
    
    // 更新权重
    UpdateWeight(ctx context.Context, service *Service, weight int) error
    
    // 获取统计
    GetStats(ctx context.Context, service *Service) (*LoadStats, error)
}
```

## 5. 数据结构

### 5.1 服务数据
```go
// 服务信息
type Service struct {
    ID          string            `json:"id"`
    Name        string            `json:"name"`
    Version     string            `json:"version"`
    Type        ServiceType       `json:"type"`
    Status      ServiceStatus     `json:"status"`
    Instances   []*Instance       `json:"instances"`
    Metadata    map[string]string `json:"metadata"`
    CreatedAt   time.Time         `json:"created_at"`
    UpdatedAt   time.Time         `json:"updated_at"`
}

// 实例信息
type Instance struct {
    ID          string            `json:"id"`
    ServiceID   string            `json:"service_id"`
    Host        string            `json:"host"`
    Port        int               `json:"port"`
    Protocol    string            `json:"protocol"`
    Status      InstanceStatus    `json:"status"`
    Weight      int               `json:"weight"`
    Metadata    map[string]string `json:"metadata"`
    CreatedAt   time.Time         `json:"created_at"`
    UpdatedAt   time.Time         `json:"updated_at"`
}

// 服务状态
type ServiceStatus struct {
    Status      string            `json:"status"`
    Health      HealthStatus      `json:"health"`
    Instances   int               `json:"instances"`
    LastCheck   time.Time         `json:"last_check"`
    Error       string            `json:"error"`
}
```

### 5.2 健康数据
```go
// 健康状态
type HealthStatus string

const (
    HealthStatusUp      HealthStatus = "UP"
    HealthStatusDown    HealthStatus = "DOWN"
    HealthStatusUnknown HealthStatus = "UNKNOWN"
)

// 健康信息
type Health struct {
    ID          string            `json:"id"`
    ServiceID   string            `json:"service_id"`
    InstanceID  string            `json:"instance_id"`
    Status      HealthStatus      `json:"status"`
    Details     *HealthDetails    `json:"details"`
    LastCheck   time.Time         `json:"last_check"`
    UpdatedAt   time.Time         `json:"updated_at"`
}

// 健康详情
type HealthDetails struct {
    CPU         float64           `json:"cpu"`
    Memory      float64           `json:"memory"`
    Disk        float64           `json:"disk"`
    Network     *NetworkStats     `json:"network"`
    Errors      []string          `json:"errors"`
}
```

### 5.3 负载数据
```go
// 负载策略
type LoadStrategy string

const (
    LoadStrategyRandom     LoadStrategy = "random"
    LoadStrategyRoundRobin LoadStrategy = "round_robin"
    LoadStrategyWeight     LoadStrategy = "weight"
    LoadStrategyLeastConn  LoadStrategy = "least_conn"
)

// 负载统计
type LoadStats struct {
    ServiceID   string            `json:"service_id"`
    InstanceID  string            `json:"instance_id"`
    Requests    int64             `json:"requests"`
    Errors      int64             `json:"errors"`
    Latency     float64           `json:"latency"`
    CPU         float64           `json:"cpu"`
    Memory      float64           `json:"memory"`
    UpdatedAt   time.Time         `json:"updated_at"`
}
```

### 5.4 订阅数据
```go
// 订阅信息
type Subscription struct {
    ID          string            `json:"id"`
    ServiceName string            `json:"service_name"`
    ClientID    string            `json:"client_id"`
    Callback    string            `json:"callback"`
    Status      SubStatus         `json:"status"`
    CreatedAt   time.Time         `json:"created_at"`
    UpdatedAt   time.Time         `json:"updated_at"`
    Metadata    map[string]string `json:"metadata"`
}

// 订阅状态
type SubStatus struct {
    Status      string            `json:"status"`
    LastNotify  time.Time         `json:"last_notify"`
    Error       string            `json:"error"`
    RetryCount  int               `json:"retry_count"`
}
```

## 6. 实现路径

### 6.1 基础框架搭建
1. 项目初始化
   ```bash
   # 创建项目目录
   mkdir -p cmd/registry
   mkdir -p internal/{register,discovery,health,loadbalancer}
   mkdir -p pkg/{api,service,utils}
   
   # 初始化 Go 模块
   go mod init min-dsp
   
   # 添加依赖
   go get -u github.com/gin-gonic/gin
   go get -u gorm.io/gorm
   go get -u github.com/go-redis/redis/v8
   go get -u go.etcd.io/etcd/client/v3
   go get -u go.uber.org/zap
   ```

2. 配置文件
   ```yaml
   # config/registry.yaml
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
   
   etcd:
     endpoints:
       - "localhost:2379"
     dial_timeout: "5s"
     username: "root"
     password: "******"
   
   registry:
     type: "etcd"
     prefix: "/dsp/services"
     timeout: "5s"
     ttl: "30s"
   
   health:
     interval: "10s"
     timeout: "5s"
     retries: 3
   
   loadbalancer:
     strategy: "weight"
     weight:
       cpu: 0.4
       memory: 0.3
       latency: 0.3
   
   monitoring:
     log_level: "info"
     metrics_port: 9090
     alert_threshold: 0.9
   ```

### 6.2 核心功能实现

#### 6.2.1 服务注册实现
```go
// internal/register/service.go
type RegisterService struct {
    db          *gorm.DB
    cache       *redis.Client
    etcd        *clientv3.Client
    storage     StorageService
    discovery   DiscoveryService
    health      HealthService
    logger      *zap.Logger
}

func (s *RegisterService) Register(ctx context.Context, req *RegisterRequest) (*ServiceResponse, error) {
    // 1. 验证请求
    if err := s.validateRegisterRequest(req); err != nil {
        return nil, err
    }
    
    // 2. 创建服务
    service := &Service{
        Name:       req.Service.Name,
        Version:    req.Service.Version,
        Type:       req.Service.Type,
        Status:     ServiceStatus{Status: "created"},
        Instances:  []*Instance{req.Instance},
    }
    
    // 3. 存储服务
    if err := s.storage.Store(ctx, service); err != nil {
        return nil, err
    }
    
    // 4. 更新缓存
    if err := s.cacheService(service); err != nil {
        s.logger.Warn("failed to cache service", zap.Error(err))
    }
    
    // 5. 启动健康检查
    go s.startHealthCheck(service)
    
    return &ServiceResponse{Service: service}, nil
}

func (s *RegisterService) Deregister(ctx context.Context, req *DeregisterRequest) error {
    // 1. 获取服务
    service, err := s.storage.Get(ctx, req.ServiceID)
    if err != nil {
        return err
    }
    
    // 2. 移除实例
    for i, instance := range service.Instances {
        if instance.ID == req.InstanceID {
            service.Instances = append(service.Instances[:i], service.Instances[i+1:]...)
            break
        }
    }
    
    // 3. 更新服务
    if len(service.Instances) == 0 {
        if err := s.storage.Delete(ctx, service.ID); err != nil {
            return err
        }
    } else {
        if err := s.storage.Store(ctx, service); err != nil {
            return err
        }
    }
    
    // 4. 更新缓存
    if err := s.cacheService(service); err != nil {
        s.logger.Warn("failed to cache service", zap.Error(err))
    }
    
    return nil
}
```

#### 6.2.2 服务发现实现
```go
// internal/discovery/service.go
type DiscoveryService struct {
    etcd        *clientv3.Client
    cache       *redis.Client
    logger      *zap.Logger
}

func (s *DiscoveryService) Discover(ctx context.Context, req *DiscoverRequest) (*DiscoverResponse, error) {
    // 1. 从缓存获取
    services, err := s.getFromCache(ctx, req.ServiceName)
    if err == nil && len(services) > 0 {
        return &DiscoverResponse{Services: services}, nil
    }
    
    // 2. 从存储获取
    services, err = s.getFromStorage(ctx, req.ServiceName)
    if err != nil {
        return nil, err
    }
    
    // 3. 更新缓存
    if err := s.cacheServices(services); err != nil {
        s.logger.Warn("failed to cache services", zap.Error(err))
    }
    
    return &DiscoverResponse{Services: services}, nil
}

func (s *DiscoveryService) Subscribe(ctx context.Context, req *SubscribeRequest) (*SubscribeResponse, error) {
    // 1. 创建订阅
    sub := &Subscription{
        ServiceName: req.ServiceName,
        ClientID:    req.ClientID,
        Callback:    req.Callback,
        Status:      SubStatus{Status: "active"},
    }
    
    // 2. 存储订阅
    key := fmt.Sprintf("/dsp/subscription/%s/%s", req.ServiceName, req.ClientID)
    value, err := json.Marshal(sub)
    if err != nil {
        return nil, err
    }
    
    _, err = s.etcd.Put(ctx, key, string(value))
    if err != nil {
        return nil, err
    }
    
    // 3. 监听变更
    go s.watchService(ctx, req.ServiceName, sub)
    
    return &SubscribeResponse{Subscription: sub}, nil
}

func (s *DiscoveryService) watchService(ctx context.Context, serviceName string, sub *Subscription) {
    watchKey := fmt.Sprintf("/dsp/services/%s", serviceName)
    watchCh := s.etcd.Watch(ctx, watchKey)
    
    for {
        select {
        case <-ctx.Done():
            return
        case resp := <-watchCh:
            for _, event := range resp.Events {
                if event.Type == clientv3.EventTypePut {
                    var service Service
                    if err := json.Unmarshal(event.Kv.Value, &service); err != nil {
                        s.logger.Error("failed to unmarshal service", zap.Error(err))
                        continue
                    }
                    
                    if err := s.notifySubscriber(ctx, sub, &service); err != nil {
                        s.logger.Error("failed to notify subscriber", zap.Error(err))
                    }
                }
            }
        }
    }
}
```

#### 6.2.3 健康检查实现
```go
// internal/health/service.go
type HealthService struct {
    db          *gorm.DB
    cache       *redis.Client
    logger      *zap.Logger
}

func (s *HealthService) Check(ctx context.Context, service *Service) (*HealthResult, error) {
    // 1. 检查实例
    for _, instance := range service.Instances {
        // 2. 检查连接
        if err := s.checkConnection(instance); err != nil {
            s.updateStatus(ctx, service, instance, HealthStatusDown)
            continue
        }
        
        // 3. 检查资源
        if err := s.checkResources(instance); err != nil {
            s.updateStatus(ctx, service, instance, HealthStatusDown)
            continue
        }
        
        // 4. 更新状态
        s.updateStatus(ctx, service, instance, HealthStatusUp)
    }
    
    // 5. 更新服务状态
    return s.updateServiceStatus(ctx, service)
}

func (s *HealthService) checkConnection(instance *Instance) error {
    // 1. 构建请求
    url := fmt.Sprintf("%s://%s:%d/health", instance.Protocol, instance.Host, instance.Port)
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return err
    }
    
    // 2. 设置超时
    client := &http.Client{Timeout: 5 * time.Second}
    
    // 3. 发送请求
    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    
    // 4. 检查响应
    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }
    
    return nil
}

func (s *HealthService) checkResources(instance *Instance) error {
    // 1. 获取资源使用情况
    stats, err := s.getResourceStats(instance)
    if err != nil {
        return err
    }
    
    // 2. 检查CPU使用率
    if stats.CPU > 0.9 {
        return fmt.Errorf("high CPU usage: %.2f", stats.CPU)
    }
    
    // 3. 检查内存使用率
    if stats.Memory > 0.9 {
        return fmt.Errorf("high memory usage: %.2f", stats.Memory)
    }
    
    // 4. 检查磁盘使用率
    if stats.Disk > 0.9 {
        return fmt.Errorf("high disk usage: %.2f", stats.Disk)
    }
    
    return nil
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
// internal/register/service_test.go
func TestRegisterService_Register(t *testing.T) {
    tests := []struct {
        name    string
        req     *RegisterRequest
        want    *ServiceResponse
        wantErr bool
    }{
        // 测试用例
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            s := NewRegisterService(db, cache, etcd, storage, discovery, health, logger)
            got, err := s.Register(context.Background(), tt.req)
            if (err != nil) != tt.wantErr {
                t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("Register() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

### 9.2 性能测试
```go
// internal/register/service_benchmark_test.go
func BenchmarkRegisterService_Register(b *testing.B) {
    s := NewRegisterService(db, cache, etcd, storage, discovery, health, logger)
    req := createTestRequest()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, err := s.Register(context.Background(), req)
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
RUN go build -o registry ./cmd/registry

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/registry .
COPY --from=builder /app/config ./config
EXPOSE 8080
CMD ["./registry"]
```

### 10.2 Kubernetes部署
```yaml
# k8s/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: registry
spec:
  replicas: 3
  selector:
    matchLabels:
      app: registry
  template:
    metadata:
      labels:
        app: registry
    spec:
      containers:
      - name: registry
        image: registry:latest
        ports:
        - containerPort: 8080
        resources:
          requests:
            memory: "1Gi"
            cpu: "500m"
          limits:
            memory: "2Gi"
            cpu: "1000m"
        volumeMounts:
        - name: config
          mountPath: /app/config
      volumes:
      - name: config
        configMap:
          name: registry-config
```

### 10.3 监控配置
```yaml
# prometheus/service-monitor.yaml
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: registry
spec:
  selector:
    matchLabels:
      app: registry
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
1. 服务管理工具
2. 服务监控工具
3. 服务诊断工具
4. 服务备份工具 