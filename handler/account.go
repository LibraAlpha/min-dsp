package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterAccount 账户注册
func RegisterAccount(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "注册成功"})
}

// LoginAccount 账户登录
func LoginAccount(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "登录成功"})
}

// GetBalance 查询账户余额
func GetBalance(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"balance": 1000})
}
