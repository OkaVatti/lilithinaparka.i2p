package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Server struct {
		Host string
		Port string
	}
	Database struct {
		Path string
	}
	Security struct {
		JWTSecret      string
		CORSOrigins    []string
		RateLimit      int
		RateLimitBurst int
	}
	External struct {
		BskyHandle    string
		BskyAppKey    string
		BskyAppSecret string
	}
	Media struct {
		UploadPath    string
		MaxUploadSize int64
		AllowedTypes  []string
		ThumbnailSize int
	}
	Admin struct {
		Username string
		Password string
	}
}

func LoadConfig() *Config {
	// Load .env file if exists
	godotenv.Load()

	var config Config

	// Server configuration
	config.Server.Host = getEnv("HOST", "127.0.0.1")
	config.Server.Port = getEnv("PORT", "8080")

	// Database configuration
	config.Database.Path = getEnv("DB_PATH", "./blog.db")

	// Security configuration
	config.Security.JWTSecret = getEnv("JWT_SECRET", "your-secret-key-change-this")
	config.Security.CORSOrigins = []string{
		"http://localhost:3000",
		"https://localhost:3000",
		"http://127.0.0.1:3000",
	}
	config.Security.RateLimit = getEnvAsInt("RATE_LIMIT", 100)
	config.Security.RateLimitBurst = getEnvAsInt("RATE_LIMIT_BURST", 30)

	// External services
	config.External.BskyHandle = getEnv("BSKY_HANDLE", "lilithinaparka.bsky.social")
	config.External.BskyAppKey = getEnv("BSKY_APP_KEY", "")
	config.External.BskyAppSecret = getEnv("BSKY_APP_SECRET", "")

	// Media configuration
	config.Media.UploadPath = getEnv("MEDIA_UPLOAD_PATH", "./media/uploads")
	config.Media.MaxUploadSize = getEnvAsInt64("MAX_UPLOAD_SIZE", 50*1024*1024) // 50MB
	config.Media.AllowedTypes = []string{
		"image/jpeg", "image/png", "image/gif", "image/webp",
		"video/mp4", "video/webm", "video/ogg",
	}
	config.Media.ThumbnailSize = getEnvAsInt("THUMBNAIL_SIZE", 300)

	// Admin credentials
	config.Admin.Username = getEnv("ADMIN_USERNAME", "admin")
	config.Admin.Password = getEnv("ADMIN_PASSWORD", "admin123")

	return &config
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsInt64(key string, defaultValue int64) int64 {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseInt(valueStr, 10, 64); err == nil {
		return value
	}
	return defaultValue
}
