package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateCampaign 创建投放
func CreateCampaign(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "投放创建成功"})
}

// ListCampaign 投放列表
func ListCampaign(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"campaigns": []string{"campaign1", "campaign2"}})
}
