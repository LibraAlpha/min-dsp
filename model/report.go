package model

// Report 报表结构体
// 可扩展字段如曝光、点击、消耗等

type Report struct {
	Impressions int
	Clicks      int
	Spend       float64
}
