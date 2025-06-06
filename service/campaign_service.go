package service

// CampaignService 投放相关业务逻辑
// 这里只做结构预留，后续可扩展具体实现

type CampaignService struct{}

func (s *CampaignService) Create() error {
	// 创建投放逻辑
	return nil
}

func (s *CampaignService) List() ([]string, error) {
	// 查询投放列表逻辑
	return []string{"campaign1", "campaign2"}, nil
}
