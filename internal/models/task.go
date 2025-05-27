package models

import (
	"time"
)

// Task 任务信息表
type Task struct {
	ID         int64     `gorm:"primaryKey;column:id" json:"id"`
	Title      string    `gorm:"size:250;column:title" json:"title"`
	IsDeleted  int       `gorm:"not null;default:0;column:is_deleted" json:"is_deleted"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

// TableName 设置 Task 表名为 t_task
func (Task) TableName() string {
	return "t_task"
}
