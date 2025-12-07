// backend/server/handlers/handle_games.go
package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/models"
	"gorm.io/gorm"
)

type GameHandlers struct {
	DB *gorm.DB
}

func NewGameHandlers(db *gorm.DB) *GameHandlers {
	return &GameHandlers{DB: db}
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
		Alias string `json:"alias" validate:"required"`
		Score int64  `json:"score" validate:"required"`
		Level int    `json:"level"`
		Data  string `json:"data"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request",
		})
	}

	ipHash := hashIP(c.RealIP())

	gameScore := models.GameScore{
		GameID: game.ID,
		Alias:  req.Alias,
		Score:  req.Score,
		Level:  req.Level,
		Data:   req.Data,
		IPHash: ipHash,
	}

	if err := h.DB.Create(&gameScore).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to submit score",
		})
	}

	return c.JSON(http.StatusCreated, gameScore)
}

func (h *GameHandlers) SaveGame(c echo.Context) error {
	slug := c.Param("slug")

	var game models.Game
	if err := h.DB.Where("slug = ?", slug).First(&game).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Game not found",
		})
	}

	var req struct {
		Alias    string `json:"alias" validate:"required"`
		SaveData string `json:"save_data" validate:"required"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request",
		})
	}

	ipHash := hashIP(c.RealIP())

	var existingSave models.GameSave
	result := h.DB.Where("game_id = ? AND alias = ? AND ip_hash = ?", game.ID, req.Alias, ipHash).First(&existingSave)

	if result.Error == gorm.ErrRecordNotFound {
		gameSave := models.GameSave{
			GameID:   game.ID,
			Alias:    req.Alias,
			SaveData: req.SaveData,
			IPHash:   ipHash,
		}
		if err := h.DB.Create(&gameSave).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to save game",
			})
		}
		return c.JSON(http.StatusCreated, gameSave)
	}

	existingSave.SaveData = req.SaveData
	if err := h.DB.Save(&existingSave).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update save",
		})
	}

	return c.JSON(http.StatusOK, existingSave)
}

func (h *GameHandlers) LoadGame(c echo.Context) error {
	slug := c.Param("slug")
	alias := c.QueryParam("alias")

	if alias == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Alias parameter required",
		})
	}

	var game models.Game
	if err := h.DB.Where("slug = ?", slug).First(&game).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Game not found",
		})
	}

	ipHash := hashIP(c.RealIP())

	var gameSave models.GameSave
	if err := h.DB.Where("game_id = ? AND alias = ? AND ip_hash = ?", game.ID, alias, ipHash).
		First(&gameSave).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "No save found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to load game",
		})
	}

	return c.JSON(http.StatusOK, gameSave)
}

func hashIP(ip string) string {
	hash := sha256.Sum256([]byte(ip + "salt-change-this"))
	return hex.EncodeToString(hash[:])
}
