package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/blog"
	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/bsky"
	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/config"
	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/handlers"
	mw "github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/middleware"
	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/models"
	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/utils"
	"golang.org/x/time/rate"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db        *gorm.DB
	cfg       *config.Config
	sseServer *utils.SSEServer
)

func main() {
	log.Println("Starting lilithinaparka.i2p backend server...")

	// Load configuration
	cfg = config.LoadConfig()
	log.Printf("Configuration loaded: Host=%s, Port=%s", cfg.Server.Host, cfg.Server.Port)

	// Initialize database
	var err error
	db, err = initDatabase(cfg.Database.Path)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	log.Println("Database initialized successfully")

	// Initialize SSE server
	sseServer = utils.NewSSEServer()

	// Load initial data
	if err := loadInitialData(); err != nil {
		log.Printf("Warning: Failed to load initial data: %v", err)
	}

	// Start background sync
	startBackgroundSync()

	// Initialize Echo server
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// Configure middleware
	setupMiddleware(e)

	// Setup routes
	setupRoutes(e)

	// Start server
	addr := cfg.Server.Host + ":" + cfg.Server.Port
	log.Printf("Server starting on %s", addr)

	// Graceful shutdown
	go func() {
		if err := e.Start(addr); err != nil {
			log.Printf("Server error: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	if err := e.Close(); err != nil {
		log.Printf("Error during shutdown: %v", err)
	}
}

func initDatabase(dbPath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	// Enable WAL mode for better concurrency
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(1)
	if _, err := sqlDB.Exec("PRAGMA journal_mode=WAL"); err != nil {
		return nil, err
	}
	if _, err := sqlDB.Exec("PRAGMA busy_timeout=5000"); err != nil {
		return nil, err
	}

	// Auto-migrate all models
	if err := db.AutoMigrate(
		&models.BlogPost{},
		&models.BskyPost{},
		&models.Profile{},
		&models.User{},
		&models.Game{},
		&models.GameScore{},
		&models.GameSave{},
		&models.MediaItem{},
		&models.MediaCategory{},
		&models.AnalyticsEvent{},
		&models.DailyStats{},
		&models.SearchIndex{},
		&models.Notification{},
		&models.NotificationPreferences{},
	); err != nil {
		return nil, err
	}

	return db, nil
}

func startBackgroundSync() {
	ticker := time.NewTicker(5 * time.Minute)
	go func() {
		for range ticker.C {
			log.Println("Running background sync...")

			if err := blog.ScanBlogPosts(db, "./blog"); err != nil {
				log.Printf("Error rescanning blog posts: %v", err)
			}

			bskyHandle := cfg.External.BskyHandle
			if bskyHandle != "" {
				if err := bsky.FetchBskyPosts(db, bskyHandle); err != nil {
					log.Printf("Error fetching BlueSky posts: %v", err)
				}
				if err := bsky.FetchBskyProfile(db, bskyHandle); err != nil {
					log.Printf("Error fetching BlueSky profile: %v", err)
				}
			}
		}
	}()
}

func setupMiddleware(e *echo.Echo) {
	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: cfg.Security.CORSOrigins,
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// Logger
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${status} ${method} ${uri} ${latency_human}\n",
	}))

	// Recover from panics
	e.Use(middleware.Recover())

	// Rate limiting - Convert int to rate.Limit
	rateLimit := rate.Limit(cfg.Security.RateLimit)
	e.Use(middleware.RateLimiterWithConfig(middleware.RateLimiterConfig{
		Store: middleware.NewRateLimiterMemoryStore(rateLimit),
	}))

	// JWT middleware (optional authentication)
	e.Use(mw.JWTMiddleware())

	// Security headers
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "DENY",
		ContentSecurityPolicy: "default-src 'self'",
	}))
}

// In main.go, update the loadInitialData function
func loadInitialData() error {
	log.Println("Loading initial data...")

	// Load profile using enhanced function
	if err := blog.LoadProfile(db, "./blog/profile/info.txt"); err != nil {
		log.Printf("Warning: Failed to load profile: %v", err)
	}

	// Use enhanced blog processor
	blogProcessor := blog.NewBlogProcessor(db, "./blog")
	if err := blogProcessor.ProcessAllPosts(); err != nil {
		log.Printf("Warning: Failed to process blog posts: %v", err)
	}

	// Use enhanced BlueSky sync
	bskyHandle := cfg.External.BskyHandle
	if bskyHandle != "" {
		bskyEnhanced := bsky.NewBskyEnhanced(db, bskyHandle)
		if err := bskyEnhanced.SyncProfileAndPosts(); err != nil {
			log.Printf("Warning: Failed to sync BlueSky data: %v", err)
		}
	}

	log.Println("Initial data load complete")
	return nil
}

func setupRoutes(e *echo.Echo) {
	// Initialize handlers
	blogHandlers := handlers.NewBlogHandlers(db)
	bskyHandlers := handlers.NewBskyHandlers(db)
	profileHandlers := handlers.NewProfileHandlers(db)
	authHandlers := handlers.NewAuthHandlers(db)
	gameHandlers := handlers.NewGameHandlers(db)
	mediaHandlers := handlers.NewMediaHandlers(db)
	searchHandlers := handlers.NewSearchHandlers(db)
	analyticsHandlers := handlers.NewAnalyticsHandlers(db)
	notificationHandlers := handlers.NewNotificationHandlers(db)

	// API group
	api := e.Group("/api")

	// Health check
	api.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"status": "ok",
			"time":   time.Now().Format(time.RFC3339),
		})
	})

	// Auth routes
	auth := api.Group("/auth")
	auth.POST("/login", authHandlers.Login)
	auth.POST("/logout", authHandlers.Logout)
	auth.GET("/status", authHandlers.GetStatus, mw.RequireAuth())

	// Blog routes
	blogGroup := api.Group("/blog")
	blogGroup.GET("/posts", blogHandlers.GetAllPosts)
	blogGroup.GET("/posts/:slug", blogHandlers.GetPostBySlug)
	blogGroup.GET("/category/:category", blogHandlers.GetPostsByCategory)
	blogGroup.GET("/tag/:tag", blogHandlers.GetPostsByTag)
	blogGroup.POST("/rescan", blogHandlers.RescanPosts, mw.RequireAuth(), mw.RequireAdmin())

	// BlueSky routes
	bskyGroup := api.Group("/bsky")
	bskyGroup.GET("/posts", bskyHandlers.GetAllPosts)
	bskyGroup.GET("/post", bskyHandlers.GetPostByURI)
	bskyGroup.POST("/refresh", bskyHandlers.RefreshPosts, mw.RequireAuth(), mw.RequireAdmin())
	bskyGroup.POST("/profile/refresh", bskyHandlers.RefreshProfile, mw.RequireAuth(), mw.RequireAdmin())
	bskyGroup.POST("/refresh", bskyHandlers.RefreshPosts, mw.RequireAuth(), mw.RequireAdmin())
	bskyGroup.POST("/profile/refresh", bskyHandlers.RefreshProfile, mw.RequireAuth(), mw.RequireAdmin())
	bskyGroup.GET("/stats", func(c echo.Context) error {
		bskyEnhanced := bsky.NewBskyEnhanced(db, cfg.External.BskyHandle)
		stats, err := bskyEnhanced.GetFeedStats()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to get BlueSky stats",
			})
		}
		return c.JSON(http.StatusOK, stats)
	})

	// Profile routes
	profileGroup := api.Group("/profile")
	profileGroup.GET("", profileHandlers.GetProfile)
	profileGroup.PUT("", profileHandlers.UpdateProfile, mw.RequireAuth(), mw.RequireAdmin())

	// Game routes
	gamesGroup := api.Group("/games")
	gamesGroup.GET("", gameHandlers.GetAllGames)
	gamesGroup.GET("/:slug", gameHandlers.GetGameBySlug)
	gamesGroup.GET("/:slug/leaderboard", gameHandlers.GetLeaderboard)
	gamesGroup.GET("/:slug/recent", gameHandlers.GetRecentScores)
	gamesGroup.POST("/:slug/score", gameHandlers.SubmitScore)
	gamesGroup.POST("", gameHandlers.AdminCreateGame, mw.RequireAuth(), mw.RequireAdmin())
	gamesGroup.PUT("/:slug", gameHandlers.AdminUpdateGame, mw.RequireAuth(), mw.RequireAdmin())

	// Media routes
	mediaGroup := api.Group("/media")
	mediaGroup.GET("", mediaHandlers.GetMedia)
	mediaGroup.GET("/categories", mediaHandlers.GetMediaCategories)
	mediaGroup.GET("/:id", mediaHandlers.GetMediaItem)
	mediaGroup.POST("/upload", mediaHandlers.UploadMedia, mw.RequireAuth(), mw.RequireAdmin())

	// Search routes
	searchGroup := api.Group("/search")
	searchGroup.GET("", searchHandlers.Search)
	searchGroup.GET("/autocomplete", searchHandlers.AutoComplete)
	searchGroup.POST("/rebuild", searchHandlers.RebuildIndex, mw.RequireAuth(), mw.RequireAdmin())

	// Analytics routes
	analyticsGroup := api.Group("/analytics")
	analyticsGroup.POST("/track", analyticsHandlers.TrackEvent)
	analyticsGroup.GET("/stats", analyticsHandlers.GetStats, mw.RequireAuth(), mw.RequireAdmin())
	analyticsGroup.GET("/realtime", analyticsHandlers.GetRealtimeStats, mw.RequireAuth(), mw.RequireAdmin())
	analyticsGroup.GET("/export", analyticsHandlers.ExportData, mw.RequireAuth(), mw.RequireAdmin())

	// Notification routes
	notifGroup := api.Group("/notifications")
	notifGroup.GET("", notificationHandlers.GetNotifications, mw.RequireAuth())
	notifGroup.PUT("/:id/read", notificationHandlers.MarkAsRead, mw.RequireAuth())
	notifGroup.DELETE("/:id", notificationHandlers.DeleteNotification, mw.RequireAuth())
	notifGroup.GET("/preferences", notificationHandlers.GetPreferences, mw.RequireAuth())
	notifGroup.PUT("/preferences", notificationHandlers.UpdatePreferences, mw.RequireAuth())

	// SSE endpoint
	api.GET("/events", func(c echo.Context) error {
		return sseServer.HandleSSE(c)
	})

	// Static file serving (for media)
	e.Static("/media", "./media")
}
