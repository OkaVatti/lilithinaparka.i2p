package models

import (
	"time"

	"gorm.io/gorm"
)

type AnalyticsEvent struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	SessionID string         `gorm:"index;not null" json:"session_id"`
	EventType string         `gorm:"index;not null" json:"event_type"` // pageview, click, search, game_start, game_end, etc.
	Page      string         `json:"page"`
	Referrer  string         `json:"referrer"`
	UserAgent string         `json:"user_agent"`
	IPHash    string         `json:"ip_hash"`
	Data      string         `gorm:"type:text" json:"data"` // JSON data specific to event
}

type DailyStats struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	Date       string    `gorm:"uniqueIndex;not null" json:"date"` // YYYY-MM-DD
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	PageViews  int       `gorm:"default:0" json:"page_views"`
	Visitors   int       `gorm:"default:0" json:"visitors"`
	Sessions   int       `gorm:"default:0" json:"sessions"`
	BlogViews  int       `gorm:"default:0" json:"blog_views"`
	GamePlays  int       `gorm:"default:0" json:"game_plays"`
	MediaViews int       `gorm:"default:0" json:"media_views"`
	Data       string    `gorm:"type:text" json:"data"` // Additional stats
}

type PopularContent struct {
	ContentType string `json:"content_type"` // blog, game, media
	ContentID   uint   `json:"content_id"`
	Title       string `json:"title"`
	Views       int    `json:"views"`
	Engagement  int    `json:"engagement"` // clicks, likes, etc.
}
