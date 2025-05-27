// filepath: /root/disaster_site_information_management_system/internal/models/user.go
package models

import (
	"time"
)

// User 用户信息表
type User struct {
	ID         int64     `gorm:"primaryKey;column:id" json:"id"`
	Name       string    `gorm:"size:250;not null;column:name" json:"name"`
	UserType   int       `gorm:"not null;default:0;column:user_type" json:"user_type"`
	Avatar     string    `gorm:"type:text;column:avatar" json:"avatar"`
	IsDeleted  int       `gorm:"not null;default:0;column:is_deleted" json:"is_deleted"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

// TableName 设置 User 表名为 t_user
func (User) TableName() string {
	return "t_user"
}
