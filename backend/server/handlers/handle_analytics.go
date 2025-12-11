package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/models"
	"gorm.io/gorm"
)

type AnalyticsHandlers struct {
	DB *gorm.DB
}

func NewAnalyticsHandlers(db *gorm.DB) *AnalyticsHandlers {
	return &AnalyticsHandlers{DB: db}
}

type TrackEventRequest struct {
	EventType string                 `json:"event_type" validate:"required"`
	Page      string                 `json:"page"`
	Referrer  string                 `json:"referrer"`
	Data      map[string]interface{} `json:"data"`
}

func (h *AnalyticsHandlers) TrackEvent(c echo.Context) error {
	var req TrackEventRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid event data",
		})
	}

	// Generate session ID from headers
	sessionID := generateSessionID(c.Request())

	// Hash IP for privacy
	ipHash := hashString(c.RealIP() + c.Request().UserAgent())

	// Serialize data
	dataJSON, _ := json.Marshal(req.Data)

	event := models.AnalyticsEvent{
		SessionID: sessionID,
		EventType: req.EventType,
		Page:      req.Page,
		Referrer:  req.Referrer,
		UserAgent: c.Request().UserAgent(),
		IPHash:    ipHash,
		Data:      string(dataJSON),
	}

	if err := h.DB.Create(&event).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to track event",
		})
	}

	// Update daily stats
	h.updateDailyStats(req.EventType)

	// Update content-specific analytics
	if req.EventType == "blog_view" {
		if blogID, ok := req.Data["blog_id"].(float64); ok {
			h.incrementBlogViews(uint(blogID))
		}
	} else if req.EventType == "game_play" {
		if gameSlug, ok := req.Data["game_slug"].(string); ok {
			h.incrementGamePlays(gameSlug)
		}
	}

	return c.JSON(http.StatusOK, map[string]string{
		"status": "tracked",
	})
}

func (h *AnalyticsHandlers) GetStats(c echo.Context) error {
	period := c.QueryParam("period") // day, week, month, year
	if period == "" {
		period = "week"
	}

	var stats struct {
		TotalPageViews int64                   `json:"total_page_views"`
		TotalVisitors  int64                   `json:"total_visitors"`
		TotalSessions  int64                   `json:"total_sessions"`
		PopularBlogs   []models.PopularContent `json:"popular_blogs"`
		PopularGames   []models.PopularContent `json:"popular_games"`
		RecentActivity []models.AnalyticsEvent `json:"recent_activity"`
	}

	// Calculate date range
	now := time.Now()
	var startDate time.Time

	switch period {
	case "day":
		startDate = now.AddDate(0, 0, -1)
	case "week":
		startDate = now.AddDate(0, 0, -7)
	case "month":
		startDate = now.AddDate(0, -1, 0)
	case "year":
		startDate = now.AddDate(-1, 0, 0)
	default:
		startDate = now.AddDate(0, 0, -7)
	}

	// Get total page views
	h.DB.Model(&models.AnalyticsEvent{}).
		Where("event_type = ? AND created_at >= ?", "pageview", startDate).
		Count(&stats.TotalPageViews)

	// Get unique visitors (by session ID)
	h.DB.Model(&models.AnalyticsEvent{}).
		Select("COUNT(DISTINCT session_id)").
		Where("created_at >= ?", startDate).
		Scan(&stats.TotalVisitors)

	// Get total sessions
	h.DB.Model(&models.AnalyticsEvent{}).
		Select("COUNT(DISTINCT session_id)").
		Where("created_at >= ?", startDate).
		Scan(&stats.TotalSessions)

	// Get popular blogs
	var blogViews []struct {
		Data  string `json:"data"`
		Count int    `json:"count"`
	}

	h.DB.Model(&models.AnalyticsEvent{}).
		Select("data, COUNT(*) as count").
		Where("event_type = ? AND created_at >= ?", "blog_view", startDate).
		Group("data").
		Order("count DESC").
		Limit(5).
		Find(&blogViews)

	for _, bv := range blogViews {
		var data map[string]interface{}
		json.Unmarshal([]byte(bv.Data), &data)

		if blogID, ok := data["blog_id"].(float64); ok {
			var blog models.BlogPost
			if err := h.DB.First(&blog, uint(blogID)).Error; err == nil {
				stats.PopularBlogs = append(stats.PopularBlogs, models.PopularContent{
					ContentType: "blog",
					ContentID:   uint(blogID),
					Title:       blog.Title,
					Views:       bv.Count,
				})
			}
		}
	}

	// Get popular games
	var gamePlays []struct {
		Data  string `json:"data"`
		Count int    `json:"count"`
	}

	h.DB.Model(&models.AnalyticsEvent{}).
		Select("data, COUNT(*) as count").
		Where("event_type = ? AND created_at >= ?", "game_play", startDate).
		Group("data").
		Order("count DESC").
		Limit(5).
		Find(&gamePlays)

	for _, gp := range gamePlays {
		var data map[string]interface{}
		json.Unmarshal([]byte(gp.Data), &data)

		if gameSlug, ok := data["game_slug"].(string); ok {
			var game models.Game
			if err := h.DB.Where("slug = ?", gameSlug).First(&game).Error; err == nil {
				stats.PopularGames = append(stats.PopularGames, models.PopularContent{
					ContentType: "game",
					ContentID:   game.ID,
					Title:       game.Name,
					Engagement:  gp.Count,
				})
			}
		}
	}

	// Get recent activity
	h.DB.Where("created_at >= ?", startDate).
		Order("created_at DESC").
		Limit(20).
		Find(&stats.RecentActivity)

	return c.JSON(http.StatusOK, stats)
}

func (h *AnalyticsHandlers) GetRealtimeStats(c echo.Context) error {
	// Get stats for last hour
	oneHourAgo := time.Now().Add(-1 * time.Hour)

	var stats struct {
		ActiveUsers   int64 `json:"active_users"`
		PageViews     int64 `json:"page_views"`
		CurrentVisits []struct {
			Page      string    `json:"page"`
			UserAgent string    `json:"user_agent"`
			Time      time.Time `json:"time"`
		} `json:"current_visits"`
	}

	// Active users (sessions in last 5 minutes)
	fiveMinutesAgo := time.Now().Add(-5 * time.Minute)
	h.DB.Model(&models.AnalyticsEvent{}).
		Select("COUNT(DISTINCT session_id)").
		Where("created_at >= ?", fiveMinutesAgo).
		Scan(&stats.ActiveUsers)

	// Page views in last hour
	h.DB.Model(&models.AnalyticsEvent{}).
		Where("event_type = ? AND created_at >= ?", "pageview", oneHourAgo).
		Count(&stats.PageViews)

	// Current visits (last 15 minutes)
	fifteenMinutesAgo := time.Now().Add(-15 * time.Minute)
	h.DB.Model(&models.AnalyticsEvent{}).
		Select("DISTINCT ON (session_id) page, user_agent, created_at as time").
		Where("created_at >= ?", fifteenMinutesAgo).
		Order("session_id, created_at DESC").
		Limit(10).
		Find(&stats.CurrentVisits)

	return c.JSON(http.StatusOK, stats)
}

func (h *AnalyticsHandlers) ExportData(c echo.Context) error {
	format := c.QueryParam("format") // json, csv
	if format == "" {
		format = "json"
	}

	var events []models.AnalyticsEvent
	if err := h.DB.Find(&events).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to export data",
		})
	}

	if format == "csv" {
		// Generate CSV
		csvData := "ID,SessionID,EventType,Page,Referrer,UserAgent,CreatedAt\n"
		for _, event := range events {
			csvData += fmt.Sprintf("%d,%s,%s,%s,%s,%s,%s\n",
				event.ID, event.SessionID, event.EventType,
				event.Page, event.Referrer, event.UserAgent,
				event.CreatedAt.Format(time.RFC3339))
		}

		c.Response().Header().Set("Content-Type", "text/csv")
		c.Response().Header().Set("Content-Disposition", "attachment; filename=analytics_export.csv")
		return c.String(http.StatusOK, csvData)
	}

	// Default to JSON
	return c.JSON(http.StatusOK, events)
}

// Helper functions
func generateSessionID(r *http.Request) string {
	// Create a unique session ID from various headers
	components := []string{
		r.Header.Get("User-Agent"),
		r.Header.Get("Accept-Language"),
		r.RemoteAddr,
		time.Now().Format("20060102"),
	}

	combined := ""
	for _, comp := range components {
		combined += comp
	}

	return hashString(combined)
}

func hashString(s string) string {
	hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(hash[:16]) // Use first 16 chars
}

func (h *AnalyticsHandlers) updateDailyStats(eventType string) {
	today := time.Now().Format("2006-01-02")

	var stats models.DailyStats
	result := h.DB.Where("date = ?", today).First(&stats)

	if result.Error == gorm.ErrRecordNotFound {
		stats = models.DailyStats{
			Date: today,
		}
		h.DB.Create(&stats)
	}

	// Update appropriate counter
	switch eventType {
	case "pageview":
		stats.PageViews++
	case "blog_view":
		stats.BlogViews++
	case "game_play":
		stats.GamePlays++
	case "media_view":
		stats.MediaViews++
	}

	// Update unique visitors and sessions would require more complex logic
	h.DB.Save(&stats)
}

func (h *AnalyticsHandlers) incrementBlogViews(blogID uint) {
	h.DB.Model(&models.BlogPost{}).
		Where("id = ?", blogID).
		UpdateColumn("views", gorm.Expr("views + ?", 1))
}

func (h *AnalyticsHandlers) incrementGamePlays(gameSlug string) {
	var game models.Game
	if err := h.DB.Where("slug = ?", gameSlug).First(&game).Error; err == nil {
		// Update game play count in config or separate table
		// For now, we'll just track in analytics
	}
}
