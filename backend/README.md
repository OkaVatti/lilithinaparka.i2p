# lilithinaparka.i2p - Backend API Server

## Overview

This is the backend API server for the lilithinaparka.i2p portal. Built in Go, it provides a secure, performant REST API that manages blog content, user profiles, game leaderboards, media files, and external service integration (BlueSky). It is designed to run as a hidden service within the I2P network

## ✨ Core Features

- Unified Content API: 
    - CRUD operations for Markdown blog posts, BlueSky sync, art, and video metadata.
- Game State Management: 
    - Handles global leaderboards, score submission, and game save persistence.
- File & Media Processing: 
    - Ingests Markdown and media from the filesystem, generates thumbnails, and provides adaptive video streaming.
- External Service Integration: 
    - Scheduled synchronization with the BlueSky API (AT Protocol) and webhook handling.
- Security-First Architecture: 
    - JWT authentication for admin endpoints, rate limiting, and hardened headers. SQL injection prevented via parameterized queries.
- I2P-Optimized: 
    - Configured for higher latency tolerance and lower bandwidth consumption typical of garlic routing.

## 🛠️ Tech Stack
- Language: Go 1.21+
- Database: SQLite 3.40+ (with WAL mode for performance)
- Router: gorilla/mux
- ORM/Data Layer: gorm.io/gorm
- Config Management: viper or environment variables
- Logging: Structured JSON logging with slog
- Markdown Processing: goldmark with frontmatter

## 🚀 Getting Started
Prerequisites
- Go: Version 1.21 or higher.
- SQLite: Command-line tools (sqlite3) are helpful for debugging.
- FFmpeg: Required for video thumbnail and transcode generation.
- I2P Router: For full integration testing.

## Installation & Setup

### Clone and Navigate:
```bash

git clone <your-repository-url>
cd backend
```

### Install Dependencies:
```bash
go mod download
```

## Configure Environment:

### Copy the example environment file and set your variables:
```bash
cp .env.example .env
# Edit .env with your settings
```

Key configuration includes DB_PATH, BSKY_APP_PASSWORD, JWT_SECRET, and server HOST/PORT.

## Initialize the Database:

### Run migrations and seed initial data (like default admin user, game definitions):
```bash
go run cmd/migrate/main.go
go run cmd/seed/main.go
```

## Development

### Start the development server with file watching:
```bash
go run main.go --dev
```

The API server will start, typically at http://localhost:8080. An API explorer (like Swagger UI, if configured) may be available.

## Building for Production

### Create an optimized binary:
```bash
go build -ldflags="-s -w" -o dist/server main.go
```

### Run the binary:
```bash
./dist/server --config ./config/production.yaml
```

## 📁 Project Structure

### A detailed breakdown of the backend's internal organization:
```text
backend/
├── cmd/                      # Application entry points (CLI tools)
├── internal/                 # Private application code
│   ├── api/handlers/        # HTTP request handlers
│   ├── api/middleware/      # CORS, logging, auth, rate limiting[citation:4]
│   ├── database/            # Models, migrations, and repositories
│   ├── services/            # Core business logic (blog, bsky, games)
│   └── utils/               # Shared utilities (cache, logger, validator)
├── blog/                    # Content directory (Markdown, images, videos)
├── db/                      # SQLite database file and migration scripts
├── config/                  # Configuration files (YAML/JSON)
├── go.mod
├── main.go                  # Server entry point
└── Dockerfile
```

## 🔧 Configuration

### Configuration is managed through environment variables and/or YAML files, prioritized in this order:
- Command-line flags
- Environment variables (e.g., SERVER_PORT, DB_PATH)
- Configuration file (e.g., config/production.yaml)
- Default values in code

> Security Note: Never commit files containing secrets (.env, config/local.yaml) to version control.

## 🗄️ Database Management
- Migrations: Database schema changes are managed using SQL migration files in db/migrations/. Use the cmd/migrate tool to apply them.
- Seeding: Initial data (admin user, game entries) is populated via cmd/seed.
- Backups: Implement a regular backup strategy for the SQLite file, especially before running migrations.

## 🔌 API Specification

### The backend provides a comprehensive REST API. Key endpoints include:

| Method | Endpoint             | Description                              | Auth Required                  |
| ------ | -------------------- | ---------------------------------------- | ------------------------------ |
| GET    | /api/health          | Server Health Check                      | No                             |
| GET    | /api/blog/posts      | Paginated list of blog posts             | No                             |
| POST   | /api/auth/login      | Admin login (Requires JWT + Private-Key) | No                             |
| POST   | /api/games/:id/score | Submit a game score                      | No (rate-limited)              |
| POST   | /api/media/upload    | Upload an Image or Video                 | YES (Admin JWT + Private Key)  |

### CORS Policy: 

In production, the Access-Control-Allow-Origin header is strictly set to your I2P eepsite address (e.g., http://your-site.b32.i2p). 
During development, it can be set to http://localhost:3000

## 🔒 Security & Hardening

This backend is designed for deployment on I2P, which adds inherent network-layer privacy

Additional measures include:
- Authentication: JWT-based auth for admin routes. Passwords hashed with argon2id.
- Input Validation: All incoming data is validated using struct tags and custom validators before processing.
- Rate Limiting: Implemented globally and per-endpoint (e.g., on /api/games/*/score) to prevent abuse.
- Headers: Security headers like X-Frame-Options: DENY, X-Content-Type-Options: nosniff are set by middleware.

## 📊 Deployment

- I2P Tunnel Configuration
    - To expose the backend as an I2P eepsite, you must configure an HTTP tunnel in your I2P router console. The tunnel should point to the backend server's host and port (e.g., 127.0.0.1:8080)
- Systemd Service (Linux)
    - For production deployments, a systemd service file ensures the backend starts automatically and restarts on failure. An example is provided in the deployment/ directory.
- Docker Deployment
    - A Dockerfile and docker-compose.yml are provided for containerized deployment, which is highly recommended for consistency.

### Build and run:
```bash
docker-compose up --build -d
```

### View logs:
```bash
docker-compose logs -f
```

## 🔍 Monitoring & Troubleshooting
- Logs: Check structured JSON logs for request details and errors. Log level can be set via LOG_LEVEL env var.
- Health Endpoint: GET /api/health returns server status and database connectivity.
- Common Issues:
    - "Database is locked": Ensure only one instance of the backend is writing to the SQLite file.
    - CORS errors from frontend: Verify the CORS origin setting in the backend configuration matches the frontend's origin exactly
    - BlueSky sync failing: Check the BSKY_APP_PASSWORD and handle in the .env file. Review service logs for API errors.