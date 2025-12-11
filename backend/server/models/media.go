package models

import (
	"time"

	"gorm.io/gorm"
)

type MediaItem struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Title       string         `gorm:"not null" json:"title"`
	Description string         `json:"description"`
	FileName    string         `gorm:"not null" json:"file_name"`
	FileSize    int64          `json:"file_size"`
	MimeType    string         `json:"mime_type"`
	Width       int            `json:"width"`
	Height      int            `json:"height"`
	Duration    int            `json:"duration"` // for videos, in seconds
	Thumbnail   string         `json:"thumbnail"`
	Category    string         `json:"category"`
	Tags        string         `gorm:"type:text" json:"tags"` // JSON array
	IsPublic    bool           `gorm:"default:true" json:"is_public"`
	Views       int            `gorm:"default:0" json:"views"`
	Likes       int            `gorm:"default:0" json:"likes"`
	Artist      string         `json:"artist"`
	Year        int            `json:"year"`
	EXIF        string         `gorm:"type:text" json:"exif"` // JSON object of EXIF data
}

type MediaCategory struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Name        string         `gorm:"uniqueIndex;not null" json:"name"`
	Description string         `json:"description"`
	Slug        string         `gorm:"uniqueIndex;not null" json:"slug"`
	ItemCount   int            `gorm:"default:0" json:"item_count"`
	IsPublic    bool           `gorm:"default:true" json:"is_public"`
}
