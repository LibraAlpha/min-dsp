package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetReport 获取投放报表
func GetReport(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"impressions": 100, "clicks": 10, "spend": 123.45})
}
