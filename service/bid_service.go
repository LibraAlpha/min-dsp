package service

// BidService 竞价相关业务逻辑
// 这里只做结构预留，后续可扩展具体实现

type BidService struct{}

func (s *BidService) Bid() (float64, string, error) {
	// 简单竞价逻辑
	return 1.23, "ad_content", nil
}
