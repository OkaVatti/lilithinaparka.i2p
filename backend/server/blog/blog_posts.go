package blog

import (
	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/models"
	"gorm.io/gorm"
)

func GetAllBlogPosts(db *gorm.DB, includeDrafts bool) ([]models.BlogPost, error) {
	var posts []models.BlogPost
	query := db.Order("date DESC, time DESC")

	if !includeDrafts {
		query = query.Where("draft = ?", false)
	}

	if err := query.Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func GetBlogPostBySlug(db *gorm.DB, slug string) (*models.BlogPost, error) {
	var post models.BlogPost
	if err := db.Where("slug = ?", slug).First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func GetBlogPostsByCategory(db *gorm.DB, category string, includeDrafts bool) ([]models.BlogPost, error) {
	var posts []models.BlogPost
	query := db.Where("categories LIKE ?", "%"+category+"%").Order("date DESC, time DESC")

	if !includeDrafts {
		query = query.Where("draft = ?", false)
	}

	if err := query.Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func GetBlogPostsByTag(db *gorm.DB, tag string, includeDrafts bool) ([]models.BlogPost, error) {
	var posts []models.BlogPost
	query := db.Where("tags LIKE ?", "%"+tag+"%").Order("date DESC, time DESC")

	if !includeDrafts {
		query = query.Where("draft = ?", false)
	}

	if err := query.Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}
