package service

// AdSlotService 广告位相关业务逻辑
// 这里只做结构预留，后续可扩展具体实现

type AdSlotService struct{}

func (s *AdSlotService) Create() error {
	// 创建广告位逻辑
	return nil
}

func (s *AdSlotService) List() ([]string, error) {
	// 查询广告位列表逻辑
	return []string{"slot1", "slot2"}, nil
}
