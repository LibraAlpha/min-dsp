package model

// Campaign 投放结构体
// 可扩展字段如ID、名称、出价、预算、定向等

type Campaign struct {
	ID     int
	Name   string
	Bid    float64
	Budget int
	Target string
}
