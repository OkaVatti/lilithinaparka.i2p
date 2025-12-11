package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/models"
	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/utils"
	"gorm.io/gorm"
)

type GameHandlers struct {
	DB *gorm.DB
}

func NewGameHandlers(db *gorm.DB) *GameHandlers {
	return &GameHandlers{DB: db}
}

type GameConfig struct {
	CanvasWidth  int    `json:"canvas_width"`
	CanvasHeight int    `json:"canvas_height"`
	Instructions string `json:"instructions"`
	Controls     struct {
		Up     string `json:"up"`
		Down   string `json:"down"`
		Left   string `json:"left"`
		Right  string `json:"right"`
		Action string `json:"action"`
	} `json:"controls"`
	Scoring struct {
		PointsPerItem int `json:"points_per_item"`
		BonusPoints   int `json:"bonus_points"`
		TimeBonus     int `json:"time_bonus"`
	} `json:"scoring"`
}

func (h *GameHandlers) GetAllGames(c echo.Context) error {
	var games []models.Game
	if err := h.DB.Order("name ASC").Find(&games).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch games",
		})
	}
	return c.JSON(http.StatusOK, games)
}

func (h *GameHandlers) GetGameBySlug(c echo.Context) error {
	slug := c.Param("slug")
	var game models.Game
	if err := h.DB.Where("slug = ?", slug).First(&game).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "Game not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch game",
		})
	}

	return c.JSON(http.StatusOK, game)
}

func (h *GameHandlers) GetLeaderboard(c echo.Context) error {
	slug := c.Param("slug")
	limitStr := c.QueryParam("limit")
	limit := 100
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	var game models.Game
	if err := h.DB.Where("slug = ?", slug).First(&game).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Game not found",
		})
	}

	var scores []models.GameScore
	if err := h.DB.Where("game_id = ?", game.ID).
		Order("score DESC").
		Limit(limit).
		Find(&scores).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch leaderboard",
		})
	}

	return c.JSON(http.StatusOK, scores)
}

func (h *GameHandlers) SubmitScore(c echo.Context) error {
	slug := c.Param("slug")

	var game models.Game
	if err := h.DB.Where("slug = ?", slug).First(&game).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Game not found",
		})
	}

	var req struct {
		Alias string          `json:"alias" validate:"required,max=20"`
		Score int64           `json:"score" validate:"required,min=0"`
		Level int             `json:"level" validate:"min=1"`
		Data  json.RawMessage `json:"data"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request",
		})
	}

	// Basic validation
	if req.Alias == "" || len(req.Alias) > 20 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Alias must be between 1 and 20 characters",
		})
	}

	if req.Score < 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Score cannot be negative",
		})
	}

	if req.Level < 1 {
		req.Level = 1
	}

	ipHash := hashIP(c.RealIP())

	gameScore := models.GameScore{
		GameID: game.ID,
		Alias:  req.Alias,
		Score:  req.Score,
		Level:  req.Level,
		Data:   string(req.Data),
		IPHash: ipHash,
	}

	if err := h.DB.Create(&gameScore).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to submit score",
		})
	}

	// Broadcast new score via SSE
	utils.NotifyNewScore(slug, req.Alias, req.Score)

	return c.JSON(http.StatusCreated, gameScore)
}

func (h *GameHandlers) GetRecentScores(c echo.Context) error {
	slug := c.Param("slug")
	limitStr := c.QueryParam("limit")
	limit := 20
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	var game models.Game
	if err := h.DB.Where("slug = ?", slug).First(&game).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Game not found",
		})
	}

	var scores []models.GameScore
	if err := h.DB.Where("game_id = ?", game.ID).
		Order("created_at DESC").
		Limit(limit).
		Find(&scores).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch recent scores",
		})
	}

	return c.JSON(http.StatusOK, scores)
}

func (h *GameHandlers) AdminCreateGame(c echo.Context) error {
	var game models.Game
	if err := c.Bind(&game); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid game data",
		})
	}

	// Validate required fields
	if game.Slug == "" || game.Name == "" || game.Category == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Slug, name, and category are required",
		})
	}

	if err := h.DB.Create(&game).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create game",
		})
	}

	return c.JSON(http.StatusCreated, game)
}

func (h *GameHandlers) AdminUpdateGame(c echo.Context) error {
	slug := c.Param("slug")

	var existing models.Game
	if err := h.DB.Where("slug = ?", slug).First(&existing).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Game not found",
		})
	}

	var updates models.Game
	if err := c.Bind(&updates); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid update data",
		})
	}

	updates.ID = existing.ID
	if err := h.DB.Save(&updates).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update game",
		})
	}

	return c.JSON(http.StatusOK, updates)
}

func hashIP(ip string) string {
	salt := os.Getenv("IP_SALT")
	if salt == "" {
		salt = "default-salt-change-this"
	}
	hash := sha256.Sum256([]byte(ip + salt))
	return hex.EncodeToString(hash[:])
}
