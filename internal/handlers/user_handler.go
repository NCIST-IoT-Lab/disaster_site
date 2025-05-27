package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"disaster_site_information_management_system/internal/models"
	"disaster_site_information_management_system/internal/utils"
)

// UserHandler 处理用户相关的请求
type UserHandler struct {
	DB *gorm.DB
}

// NewUserHandler 创建一个新的UserHandler
func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{DB: db}
}

// GetAllUsers 获取所有用户信息
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	var users []models.User
	result := h.DB.Find(&users)
	if result.Error != nil {
		utils.InternalServerErrorResponse(c, "获取用户信息失败")
		return
	}

	utils.SuccessResponse(c, users)
}

// GetUser 通过ID获取用户信息
func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequestResponse(c, "无效的ID")
		return
	}

	var user models.User
	result := h.DB.First(&user, id)
	if result.Error != nil {
		utils.NotFoundResponse(c, "用户不存在")
		return
	}

	utils.SuccessResponse(c, user)
}

// CreateUser 创建新的用户
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.BadRequestResponse(c, err.Error())
		return
	}

	result := h.DB.Create(&user)
	if result.Error != nil {
		utils.InternalServerErrorResponse(c, "创建用户失败")
		return
	}

	utils.SuccessResponse(c, user)
}

// UpdateUser 更新用户信息
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequestResponse(c, "无效的ID")
		return
	}

	var user models.User
	result := h.DB.First(&user, id)
	if result.Error != nil {
		utils.NotFoundResponse(c, "用户不存在")
		return
	}

	// 绑定更新的数据
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.BadRequestResponse(c, err.Error())
		return
	}

	h.DB.Save(&user)
	utils.SuccessResponse(c, user)
}

// DeleteUser 删除用户
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequestResponse(c, "无效的ID")
		return
	}

	var user models.User
	result := h.DB.First(&user, id)
	if result.Error != nil {
		utils.NotFoundResponse(c, "用户不存在")
		return
	}

	h.DB.Delete(&user)
	utils.SuccessResponse(c, gin.H{"message": "用户已删除"})
}
