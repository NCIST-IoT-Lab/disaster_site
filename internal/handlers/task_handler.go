package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"disaster_site_information_management_system/internal/models"
	"disaster_site_information_management_system/internal/utils"
)

// TaskHandler 处理任务相关的请求
type TaskHandler struct {
	DB *gorm.DB
}

// NewTaskHandler 创建一个新的TaskHandler
func NewTaskHandler(db *gorm.DB) *TaskHandler {
	return &TaskHandler{DB: db}
}

// GetAllTasks 获取所有任务信息
func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	var tasks []models.Task
	result := h.DB.Find(&tasks)
	if result.Error != nil {
		utils.InternalServerErrorResponse(c, "获取任务信息失败")
		return
	}

	utils.SuccessResponse(c, tasks)
}

// GetTask 通过ID获取任务信息
func (h *TaskHandler) GetTask(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequestResponse(c, "无效的ID")
		return
	}

	var task models.Task
	result := h.DB.First(&task, id)
	if result.Error != nil {
		utils.NotFoundResponse(c, "任务不存在")
		return
	}

	utils.SuccessResponse(c, task)
}

// CreateTask 创建新的任务
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		utils.BadRequestResponse(c, err.Error())
		return
	}

	result := h.DB.Create(&task)
	if result.Error != nil {
		utils.InternalServerErrorResponse(c, "创建任务失败")
		return
	}

	utils.SuccessResponse(c, task)
}

// UpdateTask 更新任务信息
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequestResponse(c, "无效的ID")
		return
	}

	var task models.Task
	result := h.DB.First(&task, id)
	if result.Error != nil {
		utils.NotFoundResponse(c, "任务不存在")
		return
	}

	// 绑定更新的数据
	if err := c.ShouldBindJSON(&task); err != nil {
		utils.BadRequestResponse(c, err.Error())
		return
	}

	h.DB.Save(&task)
	utils.SuccessResponse(c, task)
}

// DeleteTask 删除任务
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequestResponse(c, "无效的ID")
		return
	}

	var task models.Task
	result := h.DB.First(&task, id)
	if result.Error != nil {
		utils.NotFoundResponse(c, "任务不存在")
		return
	}

	h.DB.Delete(&task)
	utils.SuccessResponse(c, gin.H{"message": "任务已删除"})
}
