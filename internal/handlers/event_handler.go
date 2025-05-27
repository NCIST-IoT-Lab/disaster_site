package handlers

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"disaster_site_information_management_system/internal/models"
	"disaster_site_information_management_system/internal/utils"
	tmpl "disaster_site_information_management_system/json_tmpl"

	"github.com/samber/lo"

	"github.com/cassius0924/jtgo"
)

// EventHandler 处理事件相关的请求
type EventHandler struct {
	ctx context.Context
	DB  *gorm.DB
}

// NewEventHandler 创建一个新的EventHandler
func NewEventHandler(ctx context.Context, db *gorm.DB) *EventHandler {
	return &EventHandler{ctx: ctx, DB: db}
}

// GetAllEvents 获取所有事件信息
func (h *EventHandler) GetAllEvents(c *gin.Context) {
	var (
		events []models.Event
		users  []models.User
		wg     utils.Group
	)
	wg.Go(func() error {
		result := h.DB.Find(&events, "is_deleted = ?", 0)
		if result.Error != nil {
			return errors.New("获取事件信息失败")
		}
		return nil
	})
	wg.Go(func() error {
		result := h.DB.Find(&users, "is_deleted = ?", 0)
		if result.Error != nil {
			return errors.New("获取用户信息失败")
		}
		return nil
	})
	if err := wg.Wait(); err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}

	engine, err := jtgo.GetEngine(h.ctx, "event", tmpl.GetEventTemplate())
	if err != nil {
		utils.InternalServerErrorResponse(c, "获取模板失败\n" + err.Error())
		return
	}

	userIDToName := lo.SliceToMap(users, func(user models.User) (int64, string) {
		return user.ID, user.Name
	})

	dataset := map[string]any{
		"events":       events,
		"userIDToName": userIDToName,
	}

	result, err := engine.WithDataset(dataset).Run()
	if err != nil {
		utils.InternalServerErrorResponse(c, "渲染模板失败")
		return
	}

	log.Println("dataset:", dataset)
	utils.SuccessResponse(c, result)
}

// GetEvent 通过ID获取事件信息
func (h *EventHandler) GetEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequestResponse(c, "无效的ID")
		return
	}

	var event models.Event
	result := h.DB.First(&event, id)
	if result.Error != nil {
		utils.NotFoundResponse(c, "事件信息不存在")
		return
	}

	utils.SuccessResponse(c, event)
}

// CreateEvent 创建新的事件信息
func (h *EventHandler) CreateEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		utils.BadRequestResponse(c, err.Error())
		return
	}

	result := h.DB.Create(&event)
	if result.Error != nil {
		utils.InternalServerErrorResponse(c, "创建事件信息失败")
		return
	}

	utils.SuccessResponse(c, event)
}

// UpdateEvent 更新事件信息
func (h *EventHandler) UpdateEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequestResponse(c, "无效的ID")
		return
	}

	var event models.Event
	result := h.DB.First(&event, id)
	if result.Error != nil {
		utils.NotFoundResponse(c, "事件信息不存在")
		return
	}

	// 绑定更新的数据
	if err := c.ShouldBindJSON(&event); err != nil {
		utils.BadRequestResponse(c, err.Error())
		return
	}

	h.DB.Save(&event)
	utils.SuccessResponse(c, event)
}

// DeleteEvent 删除事件信息
func (h *EventHandler) DeleteEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.BadRequestResponse(c, "无效的ID")
		return
	}

	var event models.Event
	result := h.DB.First(&event, id)
	if result.Error != nil {
		utils.NotFoundResponse(c, "事件信息不存在")
		return
	}

	h.DB.Delete(&event)
	utils.SuccessResponse(c, gin.H{"message": "事件信息已删除"})
}
