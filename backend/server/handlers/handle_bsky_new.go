// server/handlers/handle_bsky_new.go
package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/bsky"
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
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "BlueSky post not found",
		})
	}

	return c.JSON(http.StatusOK, post)
}

func (h *BskyHandlers) RefreshPosts(c echo.Context) error {
	// You'll need to pass config or get from environment
	handle := "lilithinaparka.bsky.social" // Replace with config value

	if err := bsky.FetchBskyPosts(h.DB, handle); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to refresh BlueSky posts",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "BlueSky posts refreshed successfully",
	})
}

func (h *BskyHandlers) RefreshProfile(c echo.Context) error {
	handle := "lilithinaparka.bsky.social" // Replace with config value

	if err := bsky.FetchBskyProfile(h.DB, handle); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to refresh BlueSky profile",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "BlueSky profile refreshed successfully",
	})
}
