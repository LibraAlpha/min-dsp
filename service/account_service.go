package service

// AccountService 账户相关业务逻辑
// 这里只做结构预留，后续可扩展具体实现

type AccountService struct{}

func (s *AccountService) Register() error {
	// 注册逻辑
	return nil
}

func (s *AccountService) Login() error {
	// 登录逻辑
	return nil
}

func (s *AccountService) GetBalance() (int, error) {
	// 查询余额逻辑
	return 1000, nil
}
