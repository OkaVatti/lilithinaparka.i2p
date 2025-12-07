// backend/server/handlers/handle_profile.go
package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/models"
	"gorm.io/gorm"
)

type ProfileHandlers struct {
	DB *gorm.DB
}

func NewProfileHandlers(db *gorm.DB) *ProfileHandlers {
	return &ProfileHandlers{DB: db}
}

func (h *ProfileHandlers) GetProfile(c echo.Context) error {
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

func (h *ProfileHandlers) UpdateProfile(c echo.Context) error {
	var profile models.Profile
	if err := c.Bind(&profile); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid profile data",
		})
	}

	var existing models.Profile
	if err := h.DB.First(&existing).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Profile not found",
		})
	}

	profile.ID = existing.ID
	if err := h.DB.Save(&profile).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update profile",
		})
	}

	return c.JSON(http.StatusOK, profile)
}
