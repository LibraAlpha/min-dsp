# 广告主管理模块设计文档

## 1. 模块概述

广告主管理模块是DSP系统运营管理层的核心组件之一，负责管理广告主账户、资质、预算、投放策略等核心业务数据。该模块需要提供完整的广告主生命周期管理功能，包括账户管理、资质审核、预算管理、投放管理等功能，同时确保数据的安全性和操作的合规性。

## 2. 功能列表

### 2.1 核心功能
1. 账户管理
2. 资质管理
3. 预算管理
4. 投放管理
5. 结算管理

### 2.2 扩展功能
1. 数据统计分析
2. 操作日志审计
3. 权限管理
4. 通知提醒
5. 报表导出

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
│ 账户管理    │ │ 资质管理  │  │ 预算管理 │ │ 投放管理 │
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
- 账户管理：广告主账户的创建、修改、查询、删除等
- 资质管理：资质审核、资质更新、资质验证等
- 预算管理：预算设置、预算调整、预算监控等
- 投放管理：投放策略、投放控制、投放监控等

#### 3.2.3 数据层
- 数据访问
- 数据验证
- 数据转换
- 数据缓存

#### 3.2.4 存储层
- 关系型数据库
- 缓存数据库
- 文件存储
- 日志存储

## 4. 接口定义

### 4.1 外部接口
```go
// 广告主管理服务接口
type AdvertiserManager interface {
    // 账户管理
    CreateAccount(ctx context.Context, req *CreateAccountRequest) (*AccountResponse, error)
    UpdateAccount(ctx context.Context, req *UpdateAccountRequest) (*AccountResponse, error)
    GetAccount(ctx context.Context, req *GetAccountRequest) (*AccountResponse, error)
    ListAccounts(ctx context.Context, req *ListAccountsRequest) (*ListAccountsResponse, error)
    
    // 资质管理
    SubmitQualification(ctx context.Context, req *SubmitQualificationRequest) (*QualificationResponse, error)
    ReviewQualification(ctx context.Context, req *ReviewQualificationRequest) (*QualificationResponse, error)
    GetQualification(ctx context.Context, req *GetQualificationRequest) (*QualificationResponse, error)
    
    // 预算管理
    SetBudget(ctx context.Context, req *SetBudgetRequest) (*BudgetResponse, error)
    UpdateBudget(ctx context.Context, req *UpdateBudgetRequest) (*BudgetResponse, error)
    GetBudget(ctx context.Context, req *GetBudgetRequest) (*BudgetResponse, error)
    
    // 投放管理
    SetDelivery(ctx context.Context, req *SetDeliveryRequest) (*DeliveryResponse, error)
    UpdateDelivery(ctx context.Context, req *UpdateDeliveryRequest) (*DeliveryResponse, error)
    GetDelivery(ctx context.Context, req *GetDeliveryRequest) (*DeliveryResponse, error)
}

// 账户请求
type CreateAccountRequest struct {
    Name        string            `json:"name"`
    Type        AccountType       `json:"type"`
    Contact     *ContactInfo      `json:"contact"`
    Settings    *AccountSettings  `json:"settings"`
    Metadata    map[string]string `json:"metadata"`
}

// 资质请求
type SubmitQualificationRequest struct {
    AccountID   string            `json:"account_id"`
    Type        QualificationType `json:"type"`
    Documents   []*Document       `json:"documents"`
    ExpiryDate  time.Time         `json:"expiry_date"`
    Metadata    map[string]string `json:"metadata"`
}

// 预算请求
type SetBudgetRequest struct {
    AccountID   string            `json:"account_id"`
    Type        BudgetType        `json:"type"`
    Amount      float64           `json:"amount"`
    Period      BudgetPeriod      `json:"period"`
    Settings    *BudgetSettings   `json:"settings"`
    Metadata    map[string]string `json:"metadata"`
}

// 投放请求
type SetDeliveryRequest struct {
    AccountID   string            `json:"account_id"`
    Type        DeliveryType      `json:"type"`
    Settings    *DeliverySettings `json:"settings"`
    Schedule    *DeliverySchedule `json:"schedule"`
    Metadata    map[string]string `json:"metadata"`
}
```

### 4.2 内部接口
```go
// 账户服务接口
type AccountService interface {
    // 创建账户
    Create(ctx context.Context, account *Account) (*Account, error)
    
    // 更新账户
    Update(ctx context.Context, account *Account) (*Account, error)
    
    // 获取账户
    Get(ctx context.Context, id string) (*Account, error)
    
    // 查询账户
    Query(ctx context.Context, query *AccountQuery) ([]*Account, error)
}

// 资质服务接口
type QualificationService interface {
    // 提交资质
    Submit(ctx context.Context, qualification *Qualification) (*Qualification, error)
    
    // 审核资质
    Review(ctx context.Context, review *QualificationReview) (*Qualification, error)
    
    // 获取资质
    Get(ctx context.Context, id string) (*Qualification, error)
    
    // 验证资质
    Validate(ctx context.Context, id string) (*ValidationResult, error)
}

// 预算服务接口
type BudgetService interface {
    // 设置预算
    Set(ctx context.Context, budget *Budget) (*Budget, error)
    
    // 更新预算
    Update(ctx context.Context, budget *Budget) (*Budget, error)
    
    // 获取预算
    Get(ctx context.Context, id string) (*Budget, error)
    
    // 检查预算
    Check(ctx context.Context, id string, amount float64) (*CheckResult, error)
}

// 投放服务接口
type DeliveryService interface {
    // 设置投放
    Set(ctx context.Context, delivery *Delivery) (*Delivery, error)
    
    // 更新投放
    Update(ctx context.Context, delivery *Delivery) (*Delivery, error)
    
    // 获取投放
    Get(ctx context.Context, id string) (*Delivery, error)
    
    // 控制投放
    Control(ctx context.Context, id string, action DeliveryAction) error
}
```

## 5. 数据结构

### 5.1 账户数据
```go
// 账户类型
type AccountType string

const (
    AccountTypeIndividual AccountType = "individual"
    AccountTypeCompany    AccountType = "company"
    AccountTypeAgency     AccountType = "agency"
)

// 账户信息
type Account struct {
    ID          string            `json:"id"`
    Name        string            `json:"name"`
    Type        AccountType       `json:"type"`
    Status      AccountStatus     `json:"status"`
    Contact     *ContactInfo      `json:"contact"`
    Settings    *AccountSettings  `json:"settings"`
    CreatedAt   time.Time         `json:"created_at"`
    UpdatedAt   time.Time         `json:"updated_at"`
    Metadata    map[string]string `json:"metadata"`
}

// 联系信息
type ContactInfo struct {
    Name        string            `json:"name"`
    Email       string            `json:"email"`
    Phone       string            `json:"phone"`
    Address     string            `json:"address"`
    Position    string            `json:"position"`
}

// 账户设置
type AccountSettings struct {
    Language    string            `json:"language"`
    Timezone    string            `json:"timezone"`
    Currency    string            `json:"currency"`
    Notifications []string        `json:"notifications"`
}
```

### 5.2 资质数据
```go
// 资质类型
type QualificationType string

const (
    QualificationTypeBusiness  QualificationType = "business"
    QualificationTypePersonal  QualificationType = "personal"
    QualificationTypeSpecial   QualificationType = "special"
)

// 资质信息
type Qualification struct {
    ID          string            `json:"id"`
    AccountID   string            `json:"account_id"`
    Type        QualificationType `json:"type"`
    Status      ReviewStatus      `json:"status"`
    Documents   []*Document       `json:"documents"`
    Review      *ReviewInfo       `json:"review"`
    ExpiryDate  time.Time         `json:"expiry_date"`
    CreatedAt   time.Time         `json:"created_at"`
    UpdatedAt   time.Time         `json:"updated_at"`
}

// 文档信息
type Document struct {
    ID          string            `json:"id"`
    Type        string            `json:"type"`
    Name        string            `json:"name"`
    URL         string            `json:"url"`
    Hash        string            `json:"hash"`
    Status      string            `json:"status"`
    CreatedAt   time.Time         `json:"created_at"`
}

// 审核信息
type ReviewInfo struct {
    Reviewer    string            `json:"reviewer"`
    Status      ReviewStatus      `json:"status"`
    Comment     string            `json:"comment"`
    ReviewedAt  time.Time         `json:"reviewed_at"`
}
```

### 5.3 预算数据
```go
// 预算类型
type BudgetType string

const (
    BudgetTypeDaily    BudgetType = "daily"
    BudgetTypeWeekly   BudgetType = "weekly"
    BudgetTypeMonthly  BudgetType = "monthly"
)

// 预算信息
type Budget struct {
    ID          string            `json:"id"`
    AccountID   string            `json:"account_id"`
    Type        BudgetType        `json:"type"`
    Amount      float64           `json:"amount"`
    Used        float64           `json:"used"`
    Period      BudgetPeriod      `json:"period"`
    Status      BudgetStatus      `json:"status"`
    Settings    *BudgetSettings   `json:"settings"`
    CreatedAt   time.Time         `json:"created_at"`
    UpdatedAt   time.Time         `json:"updated_at"`
}

// 预算设置
type BudgetSettings struct {
    AlertThreshold float64         `json:"alert_threshold"`
    AutoAdjust     bool           `json:"auto_adjust"`
    AdjustLimit    float64         `json:"adjust_limit"`
    Notifications  []string        `json:"notifications"`
}

// 预算周期
type BudgetPeriod struct {
    StartTime   time.Time         `json:"start_time"`
    EndTime     time.Time         `json:"end_time"`
    Recurring   bool              `json:"recurring"`
    Interval    string            `json:"interval"`
}
```

### 5.4 投放数据
```go
// 投放类型
type DeliveryType string

const (
    DeliveryTypeRTB    DeliveryType = "rtb"
    DeliveryTypePDB    DeliveryType = "pdb"
    DeliveryTypePMP    DeliveryType = "pmp"
)

// 投放信息
type Delivery struct {
    ID          string            `json:"id"`
    AccountID   string            `json:"account_id"`
    Type        DeliveryType      `json:"type"`
    Status      DeliveryStatus    `json:"status"`
    Settings    *DeliverySettings `json:"settings"`
    Schedule    *DeliverySchedule `json:"schedule"`
    Stats       *DeliveryStats    `json:"stats"`
    CreatedAt   time.Time         `json:"created_at"`
    UpdatedAt   time.Time         `json:"updated_at"`
}

// 投放设置
type DeliverySettings struct {
    Targeting   *TargetingRules   `json:"targeting"`
    Bidding     *BiddingRules     `json:"bidding"`
    Creative    *CreativeRules    `json:"creative"`
    Frequency   *FrequencyRules   `json:"frequency"`
}

// 投放计划
type DeliverySchedule struct {
    StartTime   time.Time         `json:"start_time"`
    EndTime     time.Time         `json:"end_time"`
    TimeRanges  []*TimeRange      `json:"time_ranges"`
    DaysOfWeek  []int            `json:"days_of_week"`
    Timezone    string            `json:"timezone"`
}

// 投放统计
type DeliveryStats struct {
    Impressions int64             `json:"impressions"`
    Clicks      int64             `json:"clicks"`
    Conversions int64             `json:"conversions"`
    Spend       float64           `json:"spend"`
    UpdatedAt   time.Time         `json:"updated_at"`
}
```

## 6. 实现路径

### 6.1 基础框架搭建
1. 项目初始化
   ```bash
   # 创建项目目录
   mkdir -p cmd/manager
   mkdir -p internal/{account,qualification,budget,delivery}
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
   # config/manager.yaml
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
   
   storage:
     type: "local"
     path: "/data/uploads"
     max_size: 10485760
     allowed_types:
       - "image/jpeg"
       - "image/png"
       - "application/pdf"
   
   security:
     jwt_secret: "******"
     token_expire: "24h"
     password_salt: "******"
   
   notification:
     email:
       host: "smtp.example.com"
       port: 587
       username: "noreply@example.com"
       password: "******"
     sms:
       provider: "aliyun"
       access_key: "******"
       access_secret: "******"
   
   monitoring:
     log_level: "info"
     metrics_port: 9090
     alert_threshold: 0.9
   ```

### 6.2 核心功能实现

#### 6.2.1 账户管理实现
```go
// internal/account/service.go
type AccountService struct {
    db          *gorm.DB
    cache       *redis.Client
    storage     StorageService
    notifier    NotificationService
    logger      *zap.Logger
}

func (s *AccountService) Create(ctx context.Context, req *CreateAccountRequest) (*Account, error) {
    // 1. 验证请求
    if err := s.validateCreateRequest(req); err != nil {
        return nil, err
    }
    
    // 2. 创建账户
    account := &Account{
        Name:     req.Name,
        Type:     req.Type,
        Contact:  req.Contact,
        Settings: req.Settings,
        Status:   AccountStatusPending,
    }
    
    // 3. 保存账户
    if err := s.db.Create(account).Error; err != nil {
        return nil, err
    }
    
    // 4. 缓存账户
    if err := s.cacheAccount(account); err != nil {
        s.logger.Warn("failed to cache account", zap.Error(err))
    }
    
    // 5. 发送通知
    s.notifier.SendAccountCreated(account)
    
    return account, nil
}

func (s *AccountService) Update(ctx context.Context, req *UpdateAccountRequest) (*Account, error) {
    // 1. 获取账户
    account, err := s.Get(ctx, req.ID)
    if err != nil {
        return nil, err
    }
    
    // 2. 更新账户
    if err := s.db.Model(account).Updates(req.Updates).Error; err != nil {
        return nil, err
    }
    
    // 3. 更新缓存
    if err := s.cacheAccount(account); err != nil {
        s.logger.Warn("failed to update account cache", zap.Error(err))
    }
    
    // 4. 发送通知
    s.notifier.SendAccountUpdated(account)
    
    return account, nil
}
```

#### 6.2.2 资质管理实现
```go
// internal/qualification/service.go
type QualificationService struct {
    db          *gorm.DB
    storage     StorageService
    notifier    NotificationService
    logger      *zap.Logger
}

func (s *QualificationService) Submit(ctx context.Context, req *SubmitQualificationRequest) (*Qualification, error) {
    // 1. 验证请求
    if err := s.validateSubmitRequest(req); err != nil {
        return nil, err
    }
    
    // 2. 上传文档
    documents, err := s.uploadDocuments(req.Documents)
    if err != nil {
        return nil, err
    }
    
    // 3. 创建资质
    qualification := &Qualification{
        AccountID:  req.AccountID,
        Type:       req.Type,
        Documents:  documents,
        Status:     ReviewStatusPending,
        ExpiryDate: req.ExpiryDate,
    }
    
    // 4. 保存资质
    if err := s.db.Create(qualification).Error; err != nil {
        return nil, err
    }
    
    // 5. 发送通知
    s.notifier.SendQualificationSubmitted(qualification)
    
    return qualification, nil
}

func (s *QualificationService) Review(ctx context.Context, req *ReviewQualificationRequest) (*Qualification, error) {
    // 1. 获取资质
    qualification, err := s.Get(ctx, req.ID)
    if err != nil {
        return nil, err
    }
    
    // 2. 更新审核状态
    qualification.Review = &ReviewInfo{
        Reviewer:   req.Reviewer,
        Status:     req.Status,
        Comment:    req.Comment,
        ReviewedAt: time.Now(),
    }
    
    // 3. 保存更新
    if err := s.db.Save(qualification).Error; err != nil {
        return nil, err
    }
    
    // 4. 发送通知
    s.notifier.SendQualificationReviewed(qualification)
    
    return qualification, nil
}
```

#### 6.2.3 预算管理实现
```go
// internal/budget/service.go
type BudgetService struct {
    db          *gorm.DB
    cache       *redis.Client
    notifier    NotificationService
    logger      *zap.Logger
}

func (s *BudgetService) Set(ctx context.Context, req *SetBudgetRequest) (*Budget, error) {
    // 1. 验证请求
    if err := s.validateSetRequest(req); err != nil {
        return nil, err
    }
    
    // 2. 创建预算
    budget := &Budget{
        AccountID:  req.AccountID,
        Type:       req.Type,
        Amount:     req.Amount,
        Period:     req.Period,
        Status:     BudgetStatusActive,
        Settings:   req.Settings,
    }
    
    // 3. 保存预算
    if err := s.db.Create(budget).Error; err != nil {
        return nil, err
    }
    
    // 4. 缓存预算
    if err := s.cacheBudget(budget); err != nil {
        s.logger.Warn("failed to cache budget", zap.Error(err))
    }
    
    // 5. 发送通知
    s.notifier.SendBudgetSet(budget)
    
    return budget, nil
}

func (s *BudgetService) Check(ctx context.Context, id string, amount float64) (*CheckResult, error) {
    // 1. 获取预算
    budget, err := s.Get(ctx, id)
    if err != nil {
        return nil, err
    }
    
    // 2. 检查预算
    if budget.Used+amount > budget.Amount {
        return &CheckResult{
            Allowed: false,
            Reason:  "insufficient_budget",
        }, nil
    }
    
    // 3. 更新使用量
    if err := s.db.Model(budget).Update("used", gorm.Expr("used + ?", amount)).Error; err != nil {
        return nil, err
    }
    
    // 4. 检查告警阈值
    if budget.Settings.AlertThreshold > 0 && 
       (budget.Used+amount)/budget.Amount >= budget.Settings.AlertThreshold {
        s.notifier.SendBudgetAlert(budget)
    }
    
    return &CheckResult{
        Allowed: true,
        Reason:  "success",
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
// internal/account/service_test.go
func TestAccountService_Create(t *testing.T) {
    tests := []struct {
        name    string
        req     *CreateAccountRequest
        want    *Account
        wantErr bool
    }{
        // 测试用例
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            s := NewAccountService(db, cache, storage, notifier, logger)
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
// internal/account/service_benchmark_test.go
func BenchmarkAccountService_Create(b *testing.B) {
    s := NewAccountService(db, cache, storage, notifier, logger)
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
RUN go build -o manager ./cmd/manager

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/manager .
COPY --from=builder /app/config ./config
EXPOSE 8080
CMD ["./manager"]
```

### 10.2 Kubernetes部署
```yaml
# k8s/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: manager
spec:
  replicas: 3
  selector:
    matchLabels:
      app: manager
  template:
    metadata:
      labels:
        app: manager
    spec:
      containers:
      - name: manager
        image: manager:latest
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
        - name: uploads
          mountPath: /data/uploads
      volumes:
      - name: config
        configMap:
          name: manager-config
      - name: uploads
        persistentVolumeClaim:
          claimName: manager-uploads
```

### 10.3 监控配置
```yaml
# prometheus/service-monitor.yaml
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: manager
spec:
  selector:
    matchLabels:
      app: manager
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
1. 账户管理工具
2. 资质管理工具
3. 预算管理工具
4. 投放管理工具 