# 定向规则引擎模块设计文档

## 1. 模块概述

定向规则引擎模块是DSP系统决策引擎层的核心组件之一，负责实现和管理广告投放的定向规则。该模块需要根据广告主设定的定向条件，对广告请求进行精准匹配，确保广告投放到目标受众。同时，通过规则优化和效果分析，不断提升定向的准确性和效率。

## 2. 功能列表

### 2.1 核心功能
1. 规则管理
2. 规则匹配
3. 规则优化
4. 效果分析
5. 定向预测

### 2.2 扩展功能
1. 智能定向
2. 实时优化
3. 规则推荐
4. 效果分析
5. 异常检测

## 3. 详细设计

### 3.1 系统架构
```
                    ┌─────────────┐
                    │  规则管理器  │
                    └──────┬──────┘
                           │
        ┌─────────┬────────┴───────┬─────────┐
        │         │                │         │
┌───────▼───┐ ┌───▼──────┐  ┌──────▼────┐ ┌──▼──────┐
│ 规则加载器 │ │ 规则执行器 │  │ 规则优化器 │ │ 效果分析器 │
└───────┬───┘ └─────┬────┘  └──────┬────┘ └────┬───┘
        │           │              │           │
        └───────────┼──────────────┼───────────┘
                    │              │
              ┌─────▼──────────────▼─────┐
              │      规则存储层          │
              └─────────────┬────────────┘
                            │
                    ┌───────▼───────┐
                    │  特征工程层   │
                    └───────────────┘
```

### 3.2 组件说明

#### 3.2.1 规则管理器
- 规则配置管理
- 规则版本控制
- 规则状态监控
- 规则操作审计

#### 3.2.2 规则加载器
- 规则加载
- 规则验证
- 规则更新
- 规则缓存

#### 3.2.3 规则执行器
- 规则匹配
- 特征提取
- 规则执行
- 结果验证

#### 3.2.4 规则优化器
- 参数优化
- 规则调整
- 效果评估
- 优化建议

#### 3.2.5 效果分析器
- 效果统计
- 趋势分析
- 效果预测
- 异常检测

## 4. 接口定义

### 4.1 外部接口
```go
// 定向规则接口
type TargetingEngine interface {
    // 匹配规则
    MatchRules(ctx context.Context, req *MatchRequest) (*MatchResponse, error)
    
    // 更新规则
    UpdateRule(ctx context.Context, rule *Rule) error
    
    // 获取规则效果
    GetRuleEffect(ctx context.Context, req *EffectRequest) (*EffectResponse, error)
    
    // 优化规则
    OptimizeRule(ctx context.Context, req *OptimizeRequest) (*OptimizeResponse, error)
}

// 匹配请求
type MatchRequest struct {
    ID          string            `json:"id"`
    User        *User             `json:"user"`
    Device      *Device           `json:"device"`
    Context     *Context          `json:"context"`
    Features    map[string]any    `json:"features"`
    Timestamp   time.Time         `json:"timestamp"`
}

// 匹配响应
type MatchResponse struct {
    ID          string            `json:"id"`
    Matched     bool              `json:"matched"`
    Rules       []Rule            `json:"rules"`
    Score       float64           `json:"score"`
    Details     map[string]any    `json:"details"`
    Timestamp   time.Time         `json:"timestamp"`
}

// 规则配置
type Rule struct {
    ID          string            `json:"id"`
    Name        string            `json:"name"`
    Type        RuleType          `json:"type"`
    Conditions  []Condition       `json:"conditions"`
    Actions     []Action          `json:"actions"`
    Priority    int               `json:"priority"`
    Status      RuleStatus        `json:"status"`
    UpdatedAt   time.Time         `json:"updated_at"`
}
```

### 4.2 内部接口
```go
// 规则执行接口
type RuleExecutor interface {
    Execute(ctx context.Context, req *ExecuteRequest) (*ExecuteResponse, error)
    Validate(ctx context.Context, rule *Rule) error
}

// 规则优化接口
type RuleOptimizer interface {
    Optimize(ctx context.Context, req *OptimizeRequest) (*OptimizeResponse, error)
    Evaluate(ctx context.Context, rule *Rule) (*Evaluation, error)
}

// 效果分析接口
type EffectAnalyzer interface {
    Analyze(ctx context.Context, req *AnalyzeRequest) (*AnalyzeResponse, error)
    Predict(ctx context.Context, req *PredictRequest) (*PredictResponse, error)
}
```

## 5. 数据结构

### 5.1 规则数据
```go
// 规则类型
type RuleType string

const (
    RuleTypeUser     RuleType = "user"
    RuleTypeDevice   RuleType = "device"
    RuleTypeContext  RuleType = "context"
    RuleTypeCustom   RuleType = "custom"
)

// 规则条件
type Condition struct {
    ID          string            `json:"id"`
    Type        ConditionType     `json:"type"`
    Field       string            `json:"field"`
    Operator    Operator          `json:"operator"`
    Value       any               `json:"value"`
    Weight      float64           `json:"weight"`
}

// 规则动作
type Action struct {
    ID          string            `json:"id"`
    Type        ActionType        `json:"type"`
    Parameters  map[string]any    `json:"parameters"`
    Priority    int               `json:"priority"`
}

// 特征数据
type Feature struct {
    ID          string            `json:"id"`
    Type        FeatureType       `json:"type"`
    Name        string            `json:"name"`
    Value       any               `json:"value"`
    Source      string            `json:"source"`
    Timestamp   time.Time         `json:"timestamp"`
}
```

### 5.2 效果数据
```go
// 规则效果
type RuleEffect struct {
    RuleID      string            `json:"rule_id"`
    Metrics     map[string]float64 `json:"metrics"`
    Trends      map[string][]Point `json:"trends"`
    Predictions map[string]float64 `json:"predictions"`
    UpdatedAt   time.Time         `json:"updated_at"`
}

// 优化建议
type Optimization struct {
    RuleID      string            `json:"rule_id"`
    Type        OptimizeType      `json:"type"`
    Parameters  map[string]any    `json:"parameters"`
    Expected    map[string]float64 `json:"expected"`
    Confidence  float64           `json:"confidence"`
    CreatedAt   time.Time         `json:"created_at"`
}

// 异常检测
type Anomaly struct {
    ID          string            `json:"id"`
    RuleID      string            `json:"rule_id"`
    Type        AnomalyType       `json:"type"`
    Metric      string            `json:"metric"`
    Value       float64           `json:"value"`
    Threshold   float64           `json:"threshold"`
    Severity    AnomalySeverity   `json:"severity"`
    DetectedAt  time.Time         `json:"detected_at"`
}
```

## 6. 实现路径

### 6.1 基础框架搭建
1. 项目初始化
   ```bash
   # 创建项目目录
   mkdir -p cmd/targeting_engine
   mkdir -p internal/{engine,executor,optimizer,analyzer}
   mkdir -p pkg/{models,utils,metrics}
   
   # 初始化 Go 模块
   go mod init min-dsp
   
   # 添加依赖
   go get -u github.com/go-redis/redis/v8
   go get -u github.com/olivere/elastic/v7
   go get -u github.com/antonmedv/expr
   ```

2. 配置文件
   ```yaml
   # config/targeting_engine.yaml
   engine:
     enabled: true
     mode: strict
     cache_size: 10000
     update_interval: 5m
   
   executor:
     max_concurrent: 1000
     timeout: 100ms
     retry_times: 3
   
   optimizer:
     update_interval: 1h
     batch_size: 1000
     learning_rate: 0.01
   
   analyzer:
     metrics_path: "config/metrics"
     update_interval: 5m
     retention_days: 30
   
   feature:
     types:
       - "user"
       - "device"
       - "context"
       - "custom"
     update_interval: 1m
     cache_size: 10000
   ```

### 6.2 核心功能实现

#### 6.2.1 规则引擎实现
```go
// internal/engine/targeting_engine.go
type TargetingEngine struct {
    config      *Config
    executor    RuleExecutor
    optimizer   RuleOptimizer
    analyzer    EffectAnalyzer
    cache       *Cache
}

func (e *TargetingEngine) MatchRules(ctx context.Context, req *MatchRequest) (*MatchResponse, error) {
    // 1. 特征提取
    features := e.extractFeatures(req)
    
    // 2. 规则匹配
    rules, err := e.matchRules(ctx, features)
    if err != nil {
        return nil, err
    }
    
    // 3. 规则执行
    executeReq := &ExecuteRequest{
        Rules:    rules,
        Features: features,
    }
    executeResp, err := e.executor.Execute(ctx, executeReq)
    if err != nil {
        return nil, err
    }
    
    // 4. 生成响应
    return &MatchResponse{
        ID:        req.ID,
        Matched:   executeResp.Matched,
        Rules:     executeResp.Rules,
        Score:     executeResp.Score,
        Details:   executeResp.Details,
        Timestamp: time.Now(),
    }, nil
}

func (e *TargetingEngine) UpdateRule(ctx context.Context, rule *Rule) error {
    // 1. 验证规则
    if err := e.executor.Validate(ctx, rule); err != nil {
        return err
    }
    
    // 2. 更新规则
    if err := e.updateRule(ctx, rule); err != nil {
        return err
    }
    
    // 3. 更新缓存
    e.cache.Set(ctx, getRuleKey(rule.ID), rule)
    
    // 4. 触发优化
    go e.triggerOptimization(ctx, rule)
    
    return nil
}
```

#### 6.2.2 规则执行器实现
```go
// internal/executor/rule_executor.go
type RuleExecutor struct {
    config      *Config
    cache       *Cache
}

func (e *RuleExecutor) Execute(ctx context.Context, req *ExecuteRequest) (*ExecuteResponse, error) {
    // 1. 规则过滤
    rules := e.filterRules(req.Rules, req.Features)
    
    // 2. 规则排序
    sort.Slice(rules, func(i, j int) bool {
        return rules[i].Priority > rules[j].Priority
    })
    
    // 3. 规则执行
    var matched bool
    var score float64
    var details = make(map[string]any)
    
    for _, rule := range rules {
        // 3.1 执行条件
        conditions := e.executeConditions(rule.Conditions, req.Features)
        if !conditions.Matched {
            continue
        }
        
        // 3.2 执行动作
        actions := e.executeActions(rule.Actions, req.Features)
        
        // 3.3 更新结果
        matched = true
        score += conditions.Score * rule.Priority
        details[rule.ID] = map[string]any{
            "conditions": conditions,
            "actions":    actions,
        }
    }
    
    return &ExecuteResponse{
        Matched:   matched,
        Rules:     rules,
        Score:     score,
        Details:   details,
    }, nil
}

func (e *RuleExecutor) executeConditions(conditions []Condition, features map[string]any) *ConditionResult {
    var matched bool
    var score float64
    
    for _, condition := range conditions {
        // 1. 获取特征值
        value, ok := features[condition.Field]
        if !ok {
            continue
        }
        
        // 2. 执行操作符
        result := e.executeOperator(condition.Operator, value, condition.Value)
        if !result {
            continue
        }
        
        // 3. 更新结果
        matched = true
        score += condition.Weight
    }
    
    return &ConditionResult{
        Matched: matched,
        Score:   score,
    }
}
```

#### 6.2.3 规则优化器实现
```go
// internal/optimizer/rule_optimizer.go
type RuleOptimizer struct {
    config      *Config
    analyzer    EffectAnalyzer
}

func (o *RuleOptimizer) Optimize(ctx context.Context, req *OptimizeRequest) (*OptimizeResponse, error) {
    // 1. 获取效果数据
    effect, err := o.analyzer.Analyze(ctx, &AnalyzeRequest{
        RuleID:    req.RuleID,
        TimeRange: req.TimeRange,
    })
    if err != nil {
        return nil, err
    }
    
    // 2. 分析优化空间
    space := o.analyzeOptimizationSpace(effect)
    
    // 3. 生成优化建议
    suggestions := o.generateSuggestions(space, req.Rule)
    
    // 4. 评估优化效果
    evaluation := o.evaluateOptimization(suggestions, effect)
    
    return &OptimizeResponse{
        Suggestions: suggestions,
        Evaluation:  evaluation,
        Confidence:  o.calculateConfidence(evaluation),
    }, nil
}

func (o *RuleOptimizer) evaluateOptimization(suggestions []*Suggestion, effect *RuleEffect) *Evaluation {
    // 1. 准备评估数据
    data := o.prepareEvaluationData(suggestions, effect)
    
    // 2. 训练评估模型
    model := o.trainEvaluationModel(data)
    
    // 3. 预测优化效果
    predictions := o.predictOptimizationEffect(model, suggestions)
    
    // 4. 生成评估报告
    return o.generateEvaluationReport(predictions, effect)
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
1. 使用连接池
2. 实现请求队列
3. 使用工作池
4. 实现请求合并

### 7.2 内存优化
1. 使用对象池
2. 实现数据分片
3. 优化数据结构
4. 控制缓存大小

### 7.3 性能监控
1. 请求延迟监控
2. 规则性能监控
3. 资源使用监控
4. 系统负载监控

## 8. 安全考虑

### 8.1 数据安全
1. 规则数据加密
2. 特征数据保护
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
// internal/engine/targeting_engine_test.go
func TestTargetingEngine_MatchRules(t *testing.T) {
    tests := []struct {
        name    string
        req     *MatchRequest
        want    *MatchResponse
        wantErr bool
    }{
        // 测试用例
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            e := NewTargetingEngine(config)
            got, err := e.MatchRules(context.Background(), tt.req)
            if (err != nil) != tt.wantErr {
                t.Errorf("MatchRules() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("MatchRules() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

### 9.2 性能测试
```go
// internal/engine/targeting_engine_benchmark_test.go
func BenchmarkTargetingEngine_MatchRules(b *testing.B) {
    engine := NewTargetingEngine(config)
    req := createTestRequest()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, err := engine.MatchRules(context.Background(), req)
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
RUN go build -o targeting_engine ./cmd/targeting_engine

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/targeting_engine .
COPY --from=builder /app/config ./config
EXPOSE 8080
CMD ["./targeting_engine"]
```

### 10.2 Kubernetes部署
```yaml
# k8s/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: targeting-engine
spec:
  replicas: 3
  selector:
    matchLabels:
      app: targeting-engine
  template:
    metadata:
      labels:
        app: targeting-engine
    spec:
      containers:
      - name: targeting-engine
        image: targeting-engine:latest
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
          name: targeting-engine-config
```

### 10.3 监控配置
```yaml
# prometheus/service-monitor.yaml
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: targeting-engine
spec:
  selector:
    matchLabels:
      app: targeting-engine
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