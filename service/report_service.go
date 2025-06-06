package service

// ReportService 报表相关业务逻辑
// 这里只做结构预留，后续可扩展具体实现

type ReportService struct{}

func (s *ReportService) GetReport() (int, int, float64, error) {
	// 简单报表统计逻辑
	return 100, 10, 123.45, nil
}
