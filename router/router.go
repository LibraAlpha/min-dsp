package router

import (
	"min-dsp/handler"

	"github.com/gin-gonic/gin"
)

// SetupRouter 初始化Gin路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 广告主账户相关路由
	account := r.Group("/account")
	{
		account.POST("/register", handler.RegisterAccount)
		account.POST("/login", handler.LoginAccount)
		account.GET("/balance", handler.GetBalance)
	}

	// 广告位管理
	adslot := r.Group("/adslot")
	{
		adslot.POST("/create", handler.CreateAdSlot)
		adslot.GET("/list", handler.ListAdSlot)
	}

	// 投放管理
	campaign := r.Group("/campaign")
	{
		campaign.POST("/create", handler.CreateCampaign)
		campaign.GET("/list", handler.ListCampaign)
	}

	// 竞价请求
	r.POST("/bid", handler.Bid)

	// 报表统计
	r.GET("/report", handler.GetReport)

	return r
}
