package model

// BidRequest 竞价请求结构体
// 可扩展字段如广告位ID、用户信息等

type BidRequest struct {
	AdSlotID int
	UserID   int
}

// BidResponse 竞价响应结构体
// 可扩展字段如出价、广告内容等

type BidResponse struct {
	Price float64
	Ad    string
}
