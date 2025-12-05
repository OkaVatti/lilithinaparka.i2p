package bsky

import (
	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/models"
	"gorm.io/gorm"
)

func GetAllBskyPosts(db *gorm.DB, limit int) ([]models.BskyPost, error) {
	var posts []models.BskyPost
	query := db.Order("posted_at DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	if err := query.Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func GetBskyPostByURI(db *gorm.DB, uri string) (*models.BskyPost, error) {
	var post models.BskyPost
	if err := db.Where("uri = ?", uri).First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}
