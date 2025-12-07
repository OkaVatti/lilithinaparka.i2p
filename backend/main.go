package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/blog"
	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/bsky"
	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/handlers"
	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Auto migrate schemas
	if err := db.AutoMigrate(
		&models.BlogPost{},
		&models.BskyPost{},
		&models.Profile{},
	); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("Database initialized successfully")
	return db, nil
}

func parseProfileInfo(filePath string) (*models.Profile, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open profile info file: %w", err)
	}
	defer file.Close()

	profile := &models.Profile{}
	scanner := bufio.NewScanner(file)
	var currentKey string
	var arrayBuilder strings.Builder
	inArray := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Handle array continuation
		if inArray {
			arrayBuilder.WriteString(line)
			if strings.Contains(line, "]") {
				inArray = false
				arrayStr := arrayBuilder.String()
				// Parse array
				arrayStr = strings.TrimSpace(arrayStr)
				arrayStr = strings.TrimPrefix(arrayStr, "[")
				arrayStr = strings.TrimSuffix(arrayStr, "]")

				var items []string
				for _, item := range strings.Split(arrayStr, ",") {
					item = strings.TrimSpace(item)
					item = strings.Trim(item, "\"")
					if item != "" {
						items = append(items, item)
					}
				}

				itemsJSON, _ := json.Marshal(items)
				if currentKey == "interests" {
					profile.Interests = string(itemsJSON)
				}

				arrayBuilder.Reset()
				currentKey = ""
			}
			continue
		}

		// Parse key-value pairs
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Check if value is an array
		if strings.HasPrefix(value, "[") {
			currentKey = key
			inArray = true
			arrayBuilder.WriteString(value)
			if strings.Contains(value, "]") {
				inArray = false
				arrayStr := value
				arrayStr = strings.TrimPrefix(arrayStr, "[")
				arrayStr = strings.TrimSuffix(arrayStr, "]")

				var items []string
				for _, item := range strings.Split(arrayStr, ",") {
					item = strings.TrimSpace(item)
					item = strings.Trim(item, "\"")
					if item != "" {
						items = append(items, item)
					}
				}

				itemsJSON, _ := json.Marshal(items)
				if key == "interests" {
					profile.Interests = string(itemsJSON)
				}

				arrayBuilder.Reset()
				currentKey = ""
			}
			continue
		}

		// Remove quotes from value
		value = strings.Trim(value, "\"")

		// Map keys to profile fields
		switch key {
		case "pic":
			profile.Pic = value
		case "name":
			profile.Name = value
		case "cake_day":
			profile.CakeDay = value
		case "bio":
			profile.Bio = value
		case "location":
			profile.Location = value
		case "timezone":
			profile.Timezone = value
		case "website":
			profile.Website = value
		case "email":
			profile.Email = value
		case "github":
			profile.Github = value
		case "bluesky":
			profile.Bluesky = value
		case "rss_feed":
			profile.RSSFeed = value
		case "bitcoin_donation_addr":
			profile.BitcoinDonationAddr = value
		case "ethereum_donation_addr":
			profile.EthereumDonationAddr = value
		case "solana_donation_addr":
			profile.SolanaDonationAddr = value
		case "monero_donation_addr":
			profile.MoneroDonationAddr = value
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading profile info file: %w", err)
	}

	return profile, nil
}

func loadProfile(db *gorm.DB, profilePath string) error {
	profile, err := parseProfileInfo(profilePath)
	if err != nil {
		return fmt.Errorf("failed to parse profile info: %w", err)
	}

	var existing models.Profile
	result := db.First(&existing)

	if result.Error == gorm.ErrRecordNotFound {
		if err := db.Create(profile).Error; err != nil {
			return fmt.Errorf("failed to create profile: %w", err)
		}
		log.Println("Profile created successfully")
	} else if result.Error == nil {
		profile.ID = existing.ID
		// Preserve BlueSky data
		profile.BskyDisplayName = existing.BskyDisplayName
		profile.BskyDescription = existing.BskyDescription
		profile.BskyAvatar = existing.BskyAvatar
		profile.BskyBanner = existing.BskyBanner
		profile.BskyFollowersCount = existing.BskyFollowersCount
		profile.BskyFollowsCount = existing.BskyFollowsCount
		profile.BskyPostsCount = existing.BskyPostsCount

		if err := db.Save(profile).Error; err != nil {
			return fmt.Errorf("failed to update profile: %w", err)
		}
		log.Println("Profile updated successfully")
	} else {
		return fmt.Errorf("database error: %w", result.Error)
	}

	return nil
}

func startBackgroundSync(db *gorm.DB) {
	ticker := time.NewTicker(5 * time.Minute)
	go func() {
		for range ticker.C {
			log.Println("Starting background sync...")

			// Rescan blog posts
			if err := blog.ScanBlogPosts(db, "./blog"); err != nil {
				log.Printf("Error rescanning blog posts: %v\n", err)
			}

			// Refresh BlueSky posts
			if err := bsky.FetchBskyPosts(db, "lilithinaparka.bsky.social"); err != nil {
				log.Printf("Error fetching BlueSky posts: %v\n", err)
			}

			// Refresh BlueSky profile
			if err := bsky.FetchBskyProfile(db, "lilithinaparka.bsky.social"); err != nil {
				log.Printf("Error fetching BlueSky profile: %v\n", err)
			}

			log.Println("Background sync completed")
		}
	}()
}

func main() {
	// Initialize database
	db, err := initDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Load profile from info.txt
	if err := loadProfile(db, "./blog/profile/info.txt"); err != nil {
		log.Printf("Warning: Failed to load profile: %v", err)
	}

	// Initial data sync
	log.Println("Performing initial data sync...")

	if err := blog.ScanBlogPosts(db, "./blog"); err != nil {
		log.Printf("Warning: Failed to scan blog posts: %v", err)
	}

	if err := bsky.FetchBskyPosts(db, "lilithinaparka.bsky.social"); err != nil {
		log.Printf("Warning: Failed to fetch BlueSky posts: %v", err)
	}

	if err := bsky.FetchBskyProfile(db, "lilithinaparka.bsky.social"); err != nil {
		log.Printf("Warning: Failed to fetch BlueSky profile: %v", err)
	}

	// Start background sync
	startBackgroundSync(db)

	// Initialize Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Handlers
	blogHandlers := handlers.NewBlogHandlers(db)
	bskyHandlers := handlers.NewBskyHandlers(db)
	profileHandlers := handlers.NewProfileHandlers(db)

	// Blog routes
	e.GET("/api/blog/posts", blogHandlers.GetAllPosts)
	e.GET("/api/blog/posts/:slug", blogHandlers.GetPostBySlug)
	e.GET("/api/blog/category/:category", blogHandlers.GetPostsByCategory)
	e.GET("/api/blog/tag/:tag", blogHandlers.GetPostsByTag)
	e.POST("/api/blog/rescan", blogHandlers.RescanPosts)

	// BlueSky routes
	e.GET("/api/bsky/posts", bskyHandlers.GetAllPosts)
	e.GET("/api/bsky/post", bskyHandlers.GetPostByURI)
	e.POST("/api/bsky/refresh", bskyHandlers.RefreshPosts)

	// Profile routes
	e.GET("/api/profile", profileHandlers.GetProfile)
	e.POST("/api/profile/refresh", bskyHandlers.RefreshProfile)
	e.PUT("/api/profile", profileHandlers.UpdateProfile)

	// SSE endpoint for real-time updates
	e.GET("/api/events", sseServer.HandleSSE)

	// Health check
	e.GET("/api/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"status": "ok",
			"time":   time.Now().Format(time.RFC3339),
		})
	})
}
