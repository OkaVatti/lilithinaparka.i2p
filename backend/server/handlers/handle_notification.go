package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/models"
	"gorm.io/gorm"
)

type NotificationHandlers struct {
	DB *gorm.DB
}

func NewNotificationHandlers(db *gorm.DB) *NotificationHandlers {
	return &NotificationHandlers{DB: db}
}

func (h *NotificationHandlers) GetNotifications(c echo.Context) error {
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Authentication required",
		})
	}

	unreadOnly := c.QueryParam("unread") == "true"
	limitStr := c.QueryParam("limit")
	limit := 50
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	query := h.DB.Where("user_id = ?", userID)
	if unreadOnly {
		query = query.Where("read = ?", false)
	}

	// Remove expired notifications
	h.DB.Where("expires_at IS NOT NULL AND expires_at < ?", time.Now()).Delete(&models.Notification{})

	var notifications []models.Notification
	if err := query.Order("created_at DESC").Limit(limit).Find(&notifications).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch notifications",
		})
	}

	return c.JSON(http.StatusOK, notifications)
}

func (h *NotificationHandlers) MarkAsRead(c echo.Context) error {
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Authentication required",
		})
	}

	id := c.Param("id")
	if id == "all" {
		// Mark all as read
		if err := h.DB.Model(&models.Notification{}).
			Where("user_id = ? AND read = ?", userID, false).
			Update("read", true).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to update notifications",
			})
		}
		return c.JSON(http.StatusOK, map[string]string{
			"message": "All notifications marked as read",
		})
	}

	// Mark specific notification as read
	notificationID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid notification ID",
		})
	}

	var notification models.Notification
	if err := h.DB.Where("id = ? AND user_id = ?", notificationID, userID).First(&notification).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Notification not found",
		})
	}

	notification.Read = true
	if err := h.DB.Save(&notification).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update notification",
		})
	}

	return c.JSON(http.StatusOK, notification)
}

func (h *NotificationHandlers) DeleteNotification(c echo.Context) error {
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Authentication required",
		})
	}

	id := c.Param("id")
	notificationID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid notification ID",
		})
	}

	if err := h.DB.Where("id = ? AND user_id = ?", notificationID, userID).Delete(&models.Notification{}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to delete notification",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Notification deleted",
	})
}

func (h *NotificationHandlers) GetPreferences(c echo.Context) error {
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Authentication required",
		})
	}

	var preferences models.NotificationPreferences
	if err := h.DB.Where("user_id = ?", userID).First(&preferences).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create default preferences
			preferences = models.NotificationPreferences{
				UserID:          userID.(uint),
				EmailEnabled:    true,
				PushEnabled:     true,
				InAppEnabled:    true,
				NotifyOnComment: true,
				NotifyOnLike:    true,
				NotifyOnFollow:  true,
				NotifyOnPost:    true,
			}
			h.DB.Create(&preferences)
		} else {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to fetch preferences",
			})
		}
	}

	return c.JSON(http.StatusOK, preferences)
}

func (h *NotificationHandlers) UpdatePreferences(c echo.Context) error {
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Authentication required",
		})
	}

	var preferences models.NotificationPreferences
	if err := c.Bind(&preferences); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid preferences data",
		})
	}

	var existing models.NotificationPreferences
	if err := h.DB.Where("user_id = ?", userID).First(&existing).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			preferences.UserID = userID.(uint)
			if err := h.DB.Create(&preferences).Error; err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{
					"error": "Failed to create preferences",
				})
			}
		} else {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to update preferences",
			})
		}
	} else {
		preferences.ID = existing.ID
		preferences.UserID = existing.UserID
		if err := h.DB.Save(&preferences).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to update preferences",
			})
		}
	}

	return c.JSON(http.StatusOK, preferences)
}

// Utility function to create notifications
func (h *NotificationHandlers) CreateNotification(userID uint, nType, title, message, actionURL string, data interface{}, expiresIn time.Duration) error {
	var dataJSON string
	if data != nil {
		bytes, err := json.Marshal(data)
		if err != nil {
			return err
		}
		dataJSON = string(bytes)
	}

	var expiresAt *time.Time
	if expiresIn > 0 {
		exp := time.Now().Add(expiresIn)
		expiresAt = &exp
	}

	notification := models.Notification{
		UserID:    userID,
		Type:      nType,
		Title:     title,
		Message:   message,
		ActionURL: actionURL,
		Data:      dataJSON,
		ExpiresAt: expiresAt,
	}

	return h.DB.Create(&notification).Error
}

// System-wide notifications
func (h *NotificationHandlers) SendSystemNotification(title, message string) error {
	// Get all users who have in-app notifications enabled
	var users []struct {
		ID uint `json:"id"`
	}

	if err := h.DB.Model(&models.User{}).Select("id").Find(&users).Error; err != nil {
		return err
	}

	for _, user := range users {
		h.CreateNotification(user.ID, "info", title, message, "", nil, 7*24*time.Hour)
	}

	return nil
}
