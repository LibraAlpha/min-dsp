package model

// Account 账户结构体
// 可扩展字段如ID、用户名、密码、余额等

type Account struct {
	ID       int
	Username string
	Password string
	Balance  int
}
