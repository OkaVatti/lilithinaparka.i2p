package models

import (
	"time"

	"gorm.io/gorm"
)

type BlogPost struct {
	ID               uint           `gorm:"primarykey" json:"id"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
	Title            string         `gorm:"not null" json:"title"`
	Date             string         `gorm:"not null" json:"date"`
	Time             string         `json:"time"`
	Authors          string         `json:"authors"`    // JSON array as string
	Tags             string         `json:"tags"`       // JSON array as string
	Categories       string         `json:"categories"` // JSON array as string
	Draft            bool           `gorm:"default:false" json:"draft"`
	Share            bool           `gorm:"default:true" json:"share"`
	Slug             string         `gorm:"uniqueIndex;not null" json:"slug"`
	Layout           string         `json:"layout"`
	TOC              bool           `gorm:"default:false" json:"toc"`
	Comments         bool           `gorm:"default:false" json:"comments"`
	Math             bool           `gorm:"default:false" json:"math"`
	FeaturedImage    string         `json:"featured_image"`
	FeaturedImageAlt string         `json:"featured_image_alt"`
	FeaturedVideo    string         `json:"featured_video"`
	FeaturedVideoAlt string         `json:"featured_video_alt"`
	Summary          string         `json:"summary"`
	Content          string         `gorm:"type:text" json:"content"`
	FilePath         string         `gorm:"uniqueIndex" json:"file_path"`
}
