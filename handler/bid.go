package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Bid 处理广告竞价请求
func Bid(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"price": 1.23, "ad": "ad_content"})
}
