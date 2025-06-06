package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateAdSlot 创建广告位
func CreateAdSlot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "广告位创建成功"})
}

// ListAdSlot 广告位列表
func ListAdSlot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"adslots": []string{"slot1", "slot2"}})
}
