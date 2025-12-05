package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/blog"
	"gorm.io/gorm"
)

type BlogHandlers struct {
	DB *gorm.DB
}

func NewBlogHandlers(db *gorm.DB) *BlogHandlers {
	return &BlogHandlers{DB: db}
}

func (h *BlogHandlers) GetAllPosts(c echo.Context) error {
	includeDrafts := c.QueryParam("drafts") == "true"

	posts, err := blog.GetAllBlogPosts(h.DB, includeDrafts)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch blog posts",
		})
	}

	return c.JSON(http.StatusOK, posts)
}

func (h *BlogHandlers) GetPostBySlug(c echo.Context) error {
	slug := c.Param("slug")

	post, err := blog.GetBlogPostBySlug(h.DB, slug)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "Blog post not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch blog post",
		})
	}

	return c.JSON(http.StatusOK, post)
}

func (h *BlogHandlers) GetPostsByCategory(c echo.Context) error {
	category := c.Param("category")
	includeDrafts := c.QueryParam("drafts") == "true"

	posts, err := blog.GetBlogPostsByCategory(h.DB, category, includeDrafts)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch blog posts",
		})
	}

	return c.JSON(http.StatusOK, posts)
}

func (h *BlogHandlers) GetPostsByTag(c echo.Context) error {
	tag := c.Param("tag")
	includeDrafts := c.QueryParam("drafts") == "true"

	posts, err := blog.GetBlogPostsByTag(h.DB, tag, includeDrafts)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch blog posts",
		})
	}

	return c.JSON(http.StatusOK, posts)
}

func (h *BlogHandlers) RescanPosts(c echo.Context) error {
	blogDir := "./blog"

	if err := blog.ScanBlogPosts(h.DB, blogDir); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to rescan blog posts",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Blog posts rescanned successfully",
	})
}
