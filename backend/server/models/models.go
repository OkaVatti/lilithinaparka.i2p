package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Slug          string         `gorm:"uniqueIndex" json:"slug"`
	Title         string         `json:"title"`
	Date          time.Time      `json:"date"`
	Authors       string         `json:"authors"`
	Tags          string         `json:"tags"`
	Categories    string         `json:"categories"`
	Draft         bool           `json:"draft"`
	Share         bool           `json:"share"`
	Layout        string         `json:"layout"`
	TOC           bool           `json:"toc"`
	Comments      bool           `json:"comments"`
	Math          bool           `json:"math"`
	FeaturedImage string         `json:"featured_image"`
	Summary       string         `json:"summary"`
	Content       string         `gorm:"type:text" json:"content"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}
