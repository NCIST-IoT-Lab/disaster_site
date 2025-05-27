// filepath: /root/disaster_site_information_management_system/internal/handlers/auth_handler.go
package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"disaster_site_information_management_system/internal/utils"
)

// AuthHandler 处理认证相关的请求
type AuthHandler struct {
	DB *gorm.DB
}

// NewAuthHandler 创建一个新的AuthHandler
func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{DB: db}
}

// Login 用户登录
func (h *AuthHandler) Login(c *gin.Context) {
	// 暂时简单实现，返回成功信息
	utils.SuccessResponse(c, gin.H{"token": "sample-token"})
}

// Register 用户注册
func (h *AuthHandler) Register(c *gin.Context) {
	// 暂时简单实现，返回成功信息
	utils.SuccessResponse(c, gin.H{"message": "用户注册成功"})
}