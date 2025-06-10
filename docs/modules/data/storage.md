# 数据存储模块设计文档

## 1. 模块概述

数据存储模块是DSP系统数据处理层的核心组件之一，负责管理和维护系统中的各类数据。该模块需要支持多种数据类型的存储需求，包括实时数据、离线数据、用户数据、广告数据等，同时确保数据的安全性、可靠性和高效访问。通过统一的数据访问接口和存储策略，为上层应用提供稳定、高效的数据服务。

## 2. 功能列表

### 2.1 核心功能
1. 数据存储管理
2. 数据访问控制
3. 数据备份恢复
4. 数据生命周期管理
5. 数据一致性保证

### 2.2 扩展功能
1. 数据分片管理
2. 数据压缩存储
3. 数据加密存储
4. 数据访问审计
5. 存储性能优化

## 3. 详细设计

### 3.1 系统架构
```
                    ┌─────────────┐
                    │  存储管理器  │
                    └──────┬──────┘
                           │
        ┌─────────┬────────┴───────┬─────────┐
        │         │                │         │
┌───────▼───┐ ┌───▼──────┐  ┌──────▼────┐ ┌──▼──────┐
│ 关系型存储 │ │ 缓存存储  │  │ 文档存储  │ │ 时序存储 │
└───────┬───┘ └─────┬────┘  └──────┬────┘ └────┬───┘
        │           │              │           │
        └───────────┼──────────────┼───────────┘
                    │              │
              ┌─────▼──────────────▼─────┐
              │      数据访问层          │
              └─────────────┬────────────┘
                            │
                    ┌───────▼───────┐
                    │  数据安全层   │
                    └───────────────┘
```

### 3.2 组件说明

#### 3.2.1 存储管理器
- 存储策略管理
- 存储资源调度
- 存储性能监控
- 存储容量管理

#### 3.2.2 关系型存储
- MySQL集群管理
- 数据分片管理
- 主从复制管理
- 数据备份恢复

#### 3.2.3 缓存存储
- Redis集群管理
- 缓存策略管理
- 数据过期管理
- 内存使用优化

#### 3.2.4 文档存储
- MongoDB集群管理
- 文档索引管理
- 数据分片管理
- 查询优化管理

#### 3.2.5 时序存储
- InfluxDB集群管理
- 数据压缩管理
- 数据保留策略
- 查询性能优化

## 4. 接口定义

### 4.1 外部接口
```go
// 存储服务接口
type StorageService interface {
    // 数据写入
    Write(ctx context.Context, req *WriteRequest) (*WriteResponse, error)
    
    // 数据读取
    Read(ctx context.Context, req *ReadRequest) (*ReadResponse, error)
    
    // 数据删除
    Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error)
    
    // 数据更新
    Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error)
}

// 写入请求
type WriteRequest struct {
    Type        StorageType        `json:"type"`
    Key         string             `json:"key"`
    Value       any                `json:"value"`
    Options     *WriteOptions      `json:"options"`
    Timestamp   time.Time          `json:"timestamp"`
}

// 读取请求
type ReadRequest struct {
    Type        StorageType        `json:"type"`
    Key         string             `json:"key"`
    Options     *ReadOptions       `json:"options"`
    Timestamp   time.Time          `json:"timestamp"`
}

// 存储选项
type StorageOptions struct {
    TTL         time.Duration      `json:"ttl"`
    Compression bool               `json:"compression"`
    Encryption  bool               `json:"encryption"`
    Replication int                `json:"replication"`
}
```

### 4.2 内部接口
```go
// 存储引擎接口
type StorageEngine interface {
    // 初始化存储
    Init(ctx context.Context, config *Config) error
    
    // 关闭存储
    Close(ctx context.Context) error
    
    // 健康检查
    HealthCheck(ctx context.Context) (*HealthStatus, error)
}

// 数据访问接口
type DataAccess interface {
    // 批量写入
    BatchWrite(ctx context.Context, items []*WriteItem) error
    
    // 批量读取
    BatchRead(ctx context.Context, keys []string) (map[string]any, error)
    
    // 范围查询
    RangeQuery(ctx context.Context, req *RangeRequest) (*RangeResponse, error)
}

// 数据管理接口
type DataManager interface {
    // 数据备份
    Backup(ctx context.Context, req *BackupRequest) (*BackupResponse, error)
    
    // 数据恢复
    Restore(ctx context.Context, req *RestoreRequest) (*RestoreResponse, error)
    
    // 数据清理
    Cleanup(ctx context.Context, req *CleanupRequest) (*CleanupResponse, error)
}
```

## 5. 数据结构

### 5.1 存储数据
```go
// 存储类型
type StorageType string

const (
    StorageTypeMySQL    StorageType = "mysql"
    StorageTypeRedis    StorageType = "redis"
    StorageTypeMongoDB  StorageType = "mongodb"
    StorageTypeInfluxDB StorageType = "influxdb"
)

// 存储配置
type StorageConfig struct {
    Type        StorageType        `json:"type"`
    Host        string             `json:"host"`
    Port        int                `json:"port"`
    Database    string             `json:"database"`
    Username    string             `json:"username"`
    Password    string             `json:"password"`
    Options     map[string]any     `json:"options"`
}

// 存储状态
type StorageStatus struct {
    Type        StorageType        `json:"type"`
    Status      string             `json:"status"`
    Capacity    int64              `json:"capacity"`
    Used        int64              `json:"used"`
    Connections int                `json:"connections"`
    UpdatedAt   time.Time          `json:"updated_at"`
}
```

### 5.2 访问数据
```go
// 访问记录
type AccessLog struct {
    ID          string            `json:"id"`
    Type        AccessType        `json:"type"`
    Key         string            `json:"key"`
    User        string            `json:"user"`
    IP          string            `json:"ip"`
    Timestamp   time.Time         `json:"timestamp"`
    Duration    time.Duration     `json:"duration"`
    Status      string            `json:"status"`
}

// 性能指标
type PerformanceMetrics struct {
    Type        StorageType        `json:"type"`
    QPS         float64            `json:"qps"`
    Latency     time.Duration      `json:"latency"`
    ErrorRate   float64            `json:"error_rate"`
    Connections int                `json:"connections"`
    UpdatedAt   time.Time          `json:"updated_at"`
}

// 存储统计
type StorageStats struct {
    Type        StorageType        `json:"type"`
    Reads       int64              `json:"reads"`
    Writes      int64              `json:"writes"`
    Deletes     int64              `json:"deletes"`
    Updates     int64              `json:"updates"`
    Errors      int64              `json:"errors"`
    UpdatedAt   time.Time          `json:"updated_at"`
}
```

## 6. 实现路径

### 6.1 基础框架搭建
1. 项目初始化
   ```bash
   # 创建项目目录
   mkdir -p cmd/storage
   mkdir -p internal/{manager,engine,access,security}
   mkdir -p pkg/{mysql,redis,mongodb,influxdb}
   
   # 初始化 Go 模块
   go mod init min-dsp
   
   # 添加依赖
   go get -u github.com/go-sql-driver/mysql
   go get -u github.com/go-redis/redis/v8
   go get -u go.mongodb.org/mongo-driver/mongo
   go get -u github.com/influxdata/influxdb-client-go/v2
   ```

2. 配置文件
   ```yaml
   # config/storage.yaml
   mysql:
     master:
       host: "localhost"
       port: 3306
       database: "dsp"
       username: "root"
       password: "******"
     slaves:
       - host: "slave1"
         port: 3306
       - host: "slave2"
         port: 3306
   
   redis:
     master:
       host: "localhost"
       port: 6379
       password: "******"
     slaves:
       - host: "slave1"
         port: 6379
       - host: "slave2"
         port: 6379
   
   mongodb:
     uri: "mongodb://localhost:27017"
     database: "dsp"
     username: "admin"
     password: "******"
   
   influxdb:
     url: "http://localhost:8086"
     token: "******"
     org: "dsp"
     bucket: "metrics"
   
   security:
     encryption_key: "******"
     access_key: "******"
     tls_enabled: true
   
   monitoring:
     metrics_port: 9090
     log_level: "info"
     alert_threshold: 0.9
   ```

### 6.2 核心功能实现

#### 6.2.1 存储管理器实现
```go
// internal/manager/storage_manager.go
type StorageManager struct {
    config      *Config
    engines     map[StorageType]StorageEngine
    monitor     *StorageMonitor
}

func (m *StorageManager) Write(ctx context.Context, req *WriteRequest) (*WriteResponse, error) {
    // 1. 获取存储引擎
    engine, err := m.getEngine(req.Type)
    if err != nil {
        return nil, err
    }
    
    // 2. 数据预处理
    data, err := m.preprocessData(req.Value)
    if err != nil {
        return nil, err
    }
    
    // 3. 写入数据
    result, err := engine.Write(ctx, &WriteItem{
        Key:   req.Key,
        Value: data,
        Options: req.Options,
    })
    if err != nil {
        return nil, err
    }
    
    // 4. 记录访问日志
    m.recordAccessLog(ctx, AccessTypeWrite, req.Key)
    
    return &WriteResponse{
        Status:  "success",
        Message: "write successful",
        Data:    result,
    }, nil
}

func (m *StorageManager) getEngine(storageType StorageType) (StorageEngine, error) {
    engine, ok := m.engines[storageType]
    if !ok {
        return nil, ErrStorageTypeNotSupported
    }
    return engine, nil
}
```

#### 6.2.2 存储引擎实现
```go
// internal/engine/mysql_engine.go
type MySQLEngine struct {
    config      *Config
    master      *sql.DB
    slaves      []*sql.DB
    monitor     *EngineMonitor
}

func (e *MySQLEngine) Write(ctx context.Context, item *WriteItem) (*WriteResult, error) {
    // 1. 准备SQL语句
    query := e.prepareQuery(item)
    
    // 2. 执行写入
    result, err := e.master.ExecContext(ctx, query, item.Value)
    if err != nil {
        return nil, err
    }
    
    // 3. 检查主从同步
    if err := e.checkReplication(ctx); err != nil {
        e.monitor.RecordError("replication_error", err)
    }
    
    return &WriteResult{
        Affected: result.RowsAffected(),
        LastID:   result.LastInsertId(),
    }, nil
}

func (e *MySQLEngine) Read(ctx context.Context, item *ReadItem) (*ReadResult, error) {
    // 1. 选择从库
    slave := e.selectSlave()
    
    // 2. 准备查询
    query := e.prepareQuery(item)
    
    // 3. 执行查询
    rows, err := slave.QueryContext(ctx, query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    // 4. 处理结果
    return e.processResults(rows)
}
```

#### 6.2.3 数据访问实现
```go
// internal/access/data_access.go
type DataAccess struct {
    config      *Config
    manager     *StorageManager
    cache       *Cache
}

func (a *DataAccess) BatchWrite(ctx context.Context, items []*WriteItem) error {
    // 1. 数据分组
    groups := a.groupItems(items)
    
    // 2. 并发写入
    var wg sync.WaitGroup
    errChan := make(chan error, len(groups))
    
    for _, group := range groups {
        wg.Add(1)
        go func(g *ItemGroup) {
            defer wg.Done()
            if err := a.writeGroup(ctx, g); err != nil {
                errChan <- err
            }
        }(group)
    }
    
    // 3. 等待完成
    wg.Wait()
    close(errChan)
    
    // 4. 检查错误
    for err := range errChan {
        if err != nil {
            return err
        }
    }
    
    return nil
}

func (a *DataAccess) writeGroup(ctx context.Context, group *ItemGroup) error {
    // 1. 开启事务
    tx, err := a.beginTransaction(ctx)
    if err != nil {
        return err
    }
    defer tx.Rollback()
    
    // 2. 批量写入
    for _, item := range group.Items {
        if err := a.writeItem(ctx, tx, item); err != nil {
            return err
        }
    }
    
    // 3. 提交事务
    return tx.Commit()
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
        metrics.RecordStorage(
            c.Request.URL.Path,
            c.Writer.Status(),
            time.Since(start),
        )
    }
}
```

#### 6.3.2 安全中间件
```go
// internal/middleware/security.go
func SecurityMiddleware(security *Security) gin.HandlerFunc {
    return func(c *gin.Context) {
        // 1. 验证访问权限
        if err := security.ValidateAccess(c); err != nil {
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }
        
        // 2. 加密敏感数据
        if err := security.EncryptSensitiveData(c); err != nil {
            c.AbortWithStatus(http.StatusInternalServerError)
            return
        }
        
        c.Next()
        
        // 3. 记录审计日志
        security.RecordAuditLog(c)
    }
}
```

## 7. 性能考虑

### 7.1 存储优化
1. 数据分片
2. 索引优化
3. 缓存策略
4. 连接池管理

### 7.2 访问优化
1. 批量操作
2. 异步处理
3. 读写分离
4. 负载均衡

### 7.3 性能监控
1. 存储性能
2. 访问性能
3. 资源使用
4. 系统负载

## 8. 安全考虑

### 8.1 数据安全
1. 数据加密
2. 访问控制
3. 数据备份
4. 审计日志

### 8.2 运行安全
1. 连接加密
2. 权限管理
3. 异常处理
4. 故障恢复

## 9. 测试方案

### 9.1 单元测试
```go
// internal/engine/mysql_engine_test.go
func TestMySQLEngine_Write(t *testing.T) {
    tests := []struct {
        name    string
        item    *WriteItem
        want    *WriteResult
        wantErr bool
    }{
        // 测试用例
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            e := NewMySQLEngine(config)
            got, err := e.Write(context.Background(), tt.item)
            if (err != nil) != tt.wantErr {
                t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("Write() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

### 9.2 性能测试
```go
// internal/engine/mysql_engine_benchmark_test.go
func BenchmarkMySQLEngine_Write(b *testing.B) {
    engine := NewMySQLEngine(config)
    item := createTestItem()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, err := engine.Write(context.Background(), item)
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
RUN go build -o storage ./cmd/storage

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/storage .
COPY --from=builder /app/config ./config
EXPOSE 8080
CMD ["./storage"]
```

### 10.2 Kubernetes部署
```yaml
# k8s/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: storage
spec:
  replicas: 3
  selector:
    matchLabels:
      app: storage
  template:
    metadata:
      labels:
        app: storage
    spec:
      containers:
      - name: storage
        image: storage:latest
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
        - name: data
          mountPath: /data
      volumes:
      - name: config
        configMap:
          name: storage-config
      - name: data
        persistentVolumeClaim:
          claimName: storage-data
```

### 10.3 监控配置
```yaml
# prometheus/service-monitor.yaml
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: storage
spec:
  selector:
    matchLabels:
      app: storage
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
1. 数据管理工具
2. 性能分析工具
3. 故障诊断工具
4. 备份恢复工具 