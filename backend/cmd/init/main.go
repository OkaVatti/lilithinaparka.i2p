package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/config"
	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/middleware"
	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/models"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	fmt.Println("=== lilithinaparka.i2p Database Initialization ===")
	fmt.Println()

	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db, err := gorm.Open(sqlite.Open(cfg.Database.Path), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Connected to database:", cfg.Database.Path)

	// Enable WAL mode
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database connection: %v", err)
	}
	sqlDB.SetMaxOpenConns(1)
	if _, err := sqlDB.Exec("PRAGMA journal_mode=WAL"); err != nil {
		log.Fatalf("Failed to enable WAL mode: %v", err)
	}
	if _, err := sqlDB.Exec("PRAGMA busy_timeout=5000"); err != nil {
		log.Fatalf("Failed to set busy timeout: %v", err)
	}

	fmt.Println("Database configuration applied")

	// Auto-migrate all models
	fmt.Println("\nMigrating database schema...")
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
		log.Fatalf("Failed to migrate database: %v", err)
	}

	fmt.Println("Database schema migrated successfully")

	// Check if admin user exists
	var adminCount int64
	db.Model(&models.User{}).Where("role = ?", "admin").Count(&adminCount)

	if adminCount > 0 {
		fmt.Println("\nAdmin user already exists. Skipping creation.")
	} else {
		fmt.Println("\nNo admin user found. Creating admin account...")
		if err := createAdminUser(db, cfg); err != nil {
			log.Fatalf("Failed to create admin user: %v", err)
		}
	}

	// Create sample game entries if none exist
	var gameCount int64
	db.Model(&models.Game{}).Count(&gameCount)

	if gameCount == 0 {
		fmt.Println("\nCreating sample game entries...")
		if err := createSampleGames(db); err != nil {
			log.Printf("Warning: Failed to create sample games: %v", err)
		}
	}

	fmt.Println("\n=== Initialization Complete ===")
	fmt.Println("\nYou can now start the server with: make run")
	fmt.Println("Or: go run main.go")
}

func createAdminUser(db *gorm.DB, cfg *config.Config) error {
	reader := bufio.NewReader(os.Stdin)

	// Get username
	fmt.Print("Enter admin username [admin]: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)
	if username == "" {
		username = "admin"
	}

	// Get password
	fmt.Print("Enter admin password (min 8 characters): ")
	passwordBytes, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return fmt.Errorf("failed to read password: %w", err)
	}
	password := string(passwordBytes)
	fmt.Println()

	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters")
	}

	// Confirm password
	fmt.Print("Confirm admin password: ")
	confirmBytes, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return fmt.Errorf("failed to read password: %w", err)
	}
	confirm := string(confirmBytes)
	fmt.Println()

	if password != confirm {
		return fmt.Errorf("passwords do not match")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Create admin user
	admin := models.User{
		Username:     username,
		Email:        username + "@localhost",
		PasswordHash: string(hashedPassword),
		Role:         "admin",
	}

	if err := db.Create(&admin).Error; err != nil {
		return fmt.Errorf("failed to create admin user: %w", err)
	}

	// Generate JWT token for testing
	token, err := middleware.GenerateToken(admin.ID, admin.Username, admin.Role)
	if err != nil {
		log.Printf("Warning: Failed to generate token: %v", err)
	} else {
		fmt.Println("\n--- Admin Account Created ---")
		fmt.Printf("Username: %s\n", username)
		fmt.Printf("Role: admin\n")
		fmt.Println("\nTest JWT Token (expires in 24h):")
		fmt.Println(token)
	}

	return nil
}

func createSampleGames(db *gorm.DB) error {
	games := []models.Game{
		{
			Slug:           "tetris",
			Name:           "Tetris",
			Description:    "Classic block puzzle game",
			Category:       "Puzzle",
			Tags:           `["puzzle","arcade","classic"]`,
			MinPlayers:     1,
			MaxPlayers:     1,
			HasLeaderboard: true,
			Version:        "1.0.0",
			Config:         `{"canvas_width":300,"canvas_height":600}`,
		},
		{
			Slug:           "sudoku",
			Name:           "Sudoku",
			Description:    "Logic-based number placement puzzle",
			Category:       "Puzzle",
			Tags:           `["puzzle","logic","numbers"]`,
			MinPlayers:     1,
			MaxPlayers:     1,
			HasLeaderboard: true,
			Version:        "1.0.0",
			Config:         `{"grid_size":9}`,
		},
		{
			Slug:               "bones",
			Name:               "Bones",
			Description:        "Tabletop card game inspired by Inscryption",
			Category:           "Strategy",
			Tags:               `["card","strategy","deck-building"]`,
			MinPlayers:         1,
			MaxPlayers:         4,
			MultiplayerSupport: true,
			HasLeaderboard:     true,
			Version:            "1.0.0",
			Config:             `{"deck_size":50,"starting_cards":5}`,
		},
		{
			Slug:           "gems",
			Name:           "Gems",
			Description:    "Match-3 puzzle game",
			Category:       "Puzzle",
			Tags:           `["match3","puzzle","casual"]`,
			MinPlayers:     1,
			MaxPlayers:     1,
			HasLeaderboard: true,
			Version:        "1.0.0",
			Config:         `{"board_width":8,"board_height":8}`,
		},
	}

	for _, game := range games {
		if err := db.Create(&game).Error; err != nil {
			return err
		}
		fmt.Printf("  - Created game: %s\n", game.Name)
	}

	return nil
}
