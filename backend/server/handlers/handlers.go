package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/models"

	"github.com/labstack/echo/v4"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
)

// DB wrapper
type DB struct {
	Conn *gorm.DB
}

func (db *DB) UpsertPost(p *models.Post) error {
	var existing models.Post
	tx := db.Conn.Where("slug = ?", p.Slug).First(&existing)
	if tx.Error == nil {
		// update
		existing.Title = p.Title
		existing.Authors = p.Authors
		existing.Tags = p.Tags
		existing.Categories = p.Categories
		existing.Draft = p.Draft
		existing.Share = p.Share
		existing.FeaturedImage = p.FeaturedImage
		existing.Summary = p.Summary
		existing.Content = p.Content
		existing.Date = p.Date
		return db.Conn.Save(&existing).Error
	}
	// create
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	return db.Conn.Create(p).Error
}

func (db *DB) ListPosts(c echo.Context) error {
	// query params
	category := c.QueryParam("category")
	showDraft := c.QueryParam("draft") == "true"
	var posts []models.Post
	q := db.Conn.Order("date desc")
	if category != "" {
		q = q.Where("categories LIKE ?", "%"+category+"%")
	}
	if !showDraft {
		q = q.Where("draft = ?", false)
	}
	if err := q.Find(&posts).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, posts)
}

func (db *DB) GetPost(c echo.Context) error {
	slug := c.Param("slug")
	var post models.Post
	if err := db.Conn.Where("slug = ?", slug).First(&post).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "not found"})
	}
	return c.JSON(http.StatusOK, post)
}

func (db *DB) CreatePost(c echo.Context) error {
	req := new(models.Post)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	// if slug missing, generate
	if req.Slug == "" {
		req.Slug = slugFromFilename(req.Title)
	}
	if req.Date.IsZero() {
		req.Date = time.Now()
	}
	// write markdown file to disk under first category
	cat := "Casual"
	if req.Categories != "" {
		cat = req.Categories
	}
	path := "./blog/posts/" + cat + "/" + req.Slug + ".md"

	// create frontmatter YAML
	front := map[string]interface{}{
		"title":      req.Title,
		"date":       req.Date.Format("2006-1-2"),
		"time":       req.Date.Format("15:04"),
		"authors":    req.Authors,
		"tags":       req.Tags,
		"categories": req.Categories,
		"draft":      req.Draft,
		"share":      req.Share,
		"slug":       req.Slug,
		"summary":    req.Summary,
	}
	y, _ := yaml.Marshal(front)
	md := "---\n" + string(y) + "\n---\n\n" + req.Content
	if err := writeFileAtomic(path, []byte(md)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	// Upsert to DB
	if err := db.UpsertPost(req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	// notify SSE
	notifyNewPost(path)
	return c.JSON(http.StatusCreated, req)
}

func (db *DB) UpdatePost(c echo.Context) error {
	slug := c.Param("slug")
	req := new(models.Post)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	// find existing
	var existing models.Post
	if err := db.Conn.Where("slug = ?", slug).First(&existing).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "not found"})
	}
	// update fields
	if req.Title != "" {
		existing.Title = req.Title
	}
	if req.Content != "" {
		existing.Content = req.Content
	}
	existing.Tags = req.Tags
	existing.Authors = req.Authors
	existing.Categories = req.Categories
	existing.Draft = req.Draft
	existing.Summary = req.Summary
	existing.UpdatedAt = time.Now()

	if err := db.Conn.Save(&existing).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	// write markdown back to disk
	cat := existing.Categories
	if cat == "" {
		cat = "Casual"
	}
	path := "./blog/posts/" + cat + "/" + existing.Slug + ".md"
	front := map[string]interface{}{
		"title":      existing.Title,
		"date":       existing.Date.Format("2006-1-2"),
		"time":       existing.Date.Format("15:04"),
		"authors":    existing.Authors,
		"tags":       existing.Tags,
		"categories": existing.Categories,
		"draft":      existing.Draft,
		"share":      existing.Share,
		"slug":       existing.Slug,
		"summary":    existing.Summary,
	}
	y, _ := yaml.Marshal(front)
	md := "---\n" + string(y) + "\n---\n\n" + existing.Content
	if err := writeFileAtomic(path, []byte(md)); err != nil {
		// log but continue
	}
	notifyNewPost(path)
	return c.JSON(http.StatusOK, existing)
}

func (db *DB) DeletePost(c echo.Context) error {
	slug := c.Param("slug")
	if err := db.Conn.Where("slug = ?", slug).Delete(&models.Post{}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	// attempt to delete file variants
	base := "./blog/posts"
	for _, cat := range postFolders {
		path := base + "/" + cat + "/" + slug + ".md"
		_ = os.Remove(path)
	}
	return c.NoContent(http.StatusNoContent)
}
