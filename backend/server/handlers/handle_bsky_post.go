package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/bsky"
	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/models"
	"gorm.io/gorm"
)

type BskyHandlers struct {
	DB *gorm.DB
}

func NewBskyHandlers(db *gorm.DB) *BskyHandlers {
	return &BskyHandlers{DB: db}
}

func (h *BskyHandlers) GetAllPosts(c echo.Context) error {
	limitStr := c.QueryParam("limit")
	limit := 50

	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	posts, err := bsky.GetAllBskyPosts(h.DB, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch BlueSky posts",
		})
	}

	return c.JSON(http.StatusOK, posts)
}

func (h *BskyHandlers) GetPostByURI(c echo.Context) error {
	uri := c.QueryParam("uri")

	if uri == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "URI parameter is required",
		})
	}

	post, err := bsky.GetBskyPostByURI(h.DB, uri)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "BlueSky post not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch BlueSky post",
		})
	}

	return c.JSON(http.StatusOK, post)
}

func (h *BskyHandlers) RefreshPosts(c echo.Context) error {
	handle := "lilithinaparka.bsky.social"

	if err := bsky.FetchBskyPosts(h.DB, handle); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to refresh BlueSky posts",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "BlueSky posts refreshed successfully",
	})
}

func (h *BskyHandlers) GetProfile(c echo.Context) error {
	var profile models.Profile
	if err := h.DB.First(&profile).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "Profile not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch profile",
		})
	}

	return c.JSON(http.StatusOK, profile)
}

func (h *BskyHandlers) RefreshProfile(c echo.Context) error {
	handle := "lilithinaparka.bsky.social"

	if err := bsky.FetchBskyProfile(h.DB, handle); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to refresh BlueSky profile",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "BlueSky profile refreshed successfully",
	})
}
