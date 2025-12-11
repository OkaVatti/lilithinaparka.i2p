package models

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	UserID    uint           `gorm:"index" json:"user_id"`
	Type      string         `gorm:"index;not null" json:"type"` // info, success, warning, error
	Title     string         `json:"title"`
	Message   string         `gorm:"type:text" json:"message"`
	Read      bool           `gorm:"default:false" json:"read"`
	ActionURL string         `json:"action_url"`
	Data      string         `gorm:"type:text" json:"data"` // JSON data
	ExpiresAt *time.Time     `json:"expires_at"`
}

type NotificationPreferences struct {
	ID              uint           `gorm:"primarykey" json:"id"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
	UserID          uint           `gorm:"uniqueIndex" json:"user_id"`
	EmailEnabled    bool           `gorm:"default:true" json:"email_enabled"`
	PushEnabled     bool           `gorm:"default:true" json:"push_enabled"`
	InAppEnabled    bool           `gorm:"default:true" json:"in_app_enabled"`
	NotifyOnComment bool           `gorm:"default:true" json:"notify_on_comment"`
	NotifyOnLike    bool           `gorm:"default:true" json:"notify_on_like"`
	NotifyOnFollow  bool           `gorm:"default:true" json:"notify_on_follow"`
	NotifyOnPost    bool           `gorm:"default:true" json:"notify_on_post"`
	QuietHoursStart string         `json:"quiet_hours_start"` // HH:MM format
	QuietHoursEnd   string         `json:"quiet_hours_end"`   // HH:MM format
}
