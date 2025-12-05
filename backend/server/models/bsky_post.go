package models

import (
	"time"

	"gorm.io/gorm"
)

type BskyPost struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	URI         string         `gorm:"uniqueIndex;not null" json:"uri"`
	CID         string         `json:"cid"`
	Author      string         `json:"author"`
	Text        string         `gorm:"type:text" json:"text"`
	PostedAt    time.Time      `json:"posted_at"`
	ReplyCount  int            `json:"reply_count"`
	RepostCount int            `json:"repost_count"`
	LikeCount   int            `json:"like_count"`
	QuoteCount  int            `json:"quote_count"`
	HasMedia    bool           `json:"has_media"`
	MediaURLs   string         `json:"media_urls"` // JSON array as string
}
