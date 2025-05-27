package models

import (
	"time"
)

// Event 事件信息表
type Event struct {
	ID         int64     `gorm:"primaryKey;column:id" json:"id"`
	UserID     int64     `gorm:"not null;column:user_id" json:"user_id"`
	TaskID     int64     `gorm:"not null;column:task_id" json:"task_id"`
	CoordX     float64   `gorm:"column:coord_x" json:"coord_x"`
	CoordY     float64   `gorm:"column:coord_y" json:"coord_y"`
	CoordZ     float64   `gorm:"column:coord_z" json:"coord_z"`
	Status     int       `gorm:"not null;column:status" json:"status"`
	Level      int       `gorm:"not null;default:0;column:level" json:"level"`
	EventType  int       `gorm:"not null;default:0;column:event_type" json:"event_type"`
	Desc       string    `gorm:"type:text;column:desc" json:"desc"`
	Image      string    `gorm:"type:text;column:image" json:"image"`
	IsDeleted  int       `gorm:"not null;default:0;column:is_deleted" json:"is_deleted"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

// TableName 设置 Event 表名为 t_event
func (Event) TableName() string {
	return "t_event"
}
