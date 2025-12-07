// backend/server/handlers/handle_auth.go
package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/middleware"
	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthHandlers struct {
	DB *gorm.DB
}

func NewAuthHandlers(db *gorm.DB) *AuthHandlers {
	return &AuthHandlers{DB: db}
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
	User      struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Role     string `json:"role"`
	} `json:"user"`
}

func (h *AuthHandlers) Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request",
		})
	}

	var user models.User
	if err := h.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid credentials",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid credentials",
		})
	}

	token, err := middleware.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to generate token",
		})
	}

	user.LastLogin = time.Now()
	h.DB.Save(&user)

	response := LoginResponse{
		Token:     token,
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	}
	response.User.ID = user.ID
	response.User.Username = user.Username
	response.User.Role = user.Role

	return c.JSON(http.StatusOK, response)
}

func (h *AuthHandlers) Logout(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Logged out successfully",
	})
}

func (h *AuthHandlers) GetStatus(c echo.Context) error {
	userID := c.Get("user_id")
	username := c.Get("username")
	role := c.Get("role")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"authenticated": true,
		"user_id":       userID,
		"username":      username,
		"role":          role,
	})
}
