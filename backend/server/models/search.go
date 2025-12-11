package models

import (
	"time"

	"gorm.io/gorm"
)

type SearchIndex struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Type      string         `gorm:"index;not null" json:"type"` // 'blog', 'media', 'game', 'page'
	ItemID    uint           `gorm:"index;not null" json:"item_id"`
	Title     string         `gorm:"not null" json:"title"`
	Content   string         `gorm:"type:text" json:"content"`
	Tags      string         `gorm:"type:text" json:"tags"` // JSON array
	Slug      string         `gorm:"index" json:"slug"`
	Weight    float64        `gorm:"default:1.0" json:"weight"`
	IsPublic  bool           `gorm:"default:true" json:"is_public"`
}

type SearchResult struct {
	Type        string    `json:"type"`
	ItemID      uint      `json:"item_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Slug        string    `json:"slug"`
	Date        time.Time `json:"date"`
	Score       float64   `json:"score"`
	Highlights  []string  `json:"highlights"`
}
