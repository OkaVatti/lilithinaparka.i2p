# lilithinaparka.i2p - Backend API Server

Clean, production-ready Go backend for the I2P blog platform.

## Features

- RESTful API for blog posts, media, games, and profiles
- BlueSky social feed integration
- SQLite database with WAL mode for better concurrency
- JWT authentication for admin endpoints
- Real-time updates via Server-Sent Events (SSE)
- Game leaderboard system
- Full-text search with indexing
- Analytics and notification system
- Docker support for easy deployment

## Quick Start

### Prerequisites

- Go 1.21 or higher
- SQLite 3.40+
- FFmpeg (for media processing)
- Make (optional, for convenience)

### Installation

1. Clone and navigate to the backend directory:
```bash
cd backend
```

2. Copy environment file and configure:
```bash
cp .env.example .env
# Edit .env with your settings
```

3. Initialize project structure:
```bash
make init
```

4. Install dependencies:
```bash
go mod download
```

5. Run the server:
```bash
make run
```

The server will start on `http://localhost:8080` by default.

## Configuration

All configuration is done via environment variables or the `.env` file:

### Server Settings
- `HOST` - Server host (default: 127.0.0.1)
- `PORT` - Server port (default: 8080)

### Database
- `DB_PATH` - SQLite database path (default: ./blog.db)

### Security
- `JWT_SECRET` - Secret key for JWT tokens (CHANGE THIS!)
- `IP_SALT` - Salt for IP hashing (CHANGE THIS!)
- `CORS_ORIGIN_*` - Allowed CORS origins

### External Services
- `BSKY_HANDLE` - BlueSky handle for feed integration

### Admin
- `ADMIN_USERNAME` - Admin username
- `ADMIN_PASSWORD` - Admin password (CHANGE THIS!)

## Project Structure

```
backend/
├── main.go                    # Application entry point
├── server/
│   ├── blog/                 # Blog post handling
│   ├── bsky/                 # BlueSky integration
│   ├── config/               # Configuration management
│   ├── games/                # Game engine
│   ├── handlers/             # HTTP handlers
│   ├── middleware/           # Authentication & rate limiting
│   ├── models/               # Database models
│   └── utils/                # Utility functions
├── blog/                     # Content directory
│   ├── posts/               # Markdown blog posts
│   ├── profile/             # Profile information
│   └── assets/              # Images and media
└── media/                    # Uploaded media files
```

## API Endpoints

### Public Endpoints

#### Health Check
```
GET /api/health
```

#### Blog Posts
```
GET /api/blog/posts              # List all posts
GET /api/blog/posts/:slug        # Get post by slug
GET /api/blog/category/:category # Posts by category
GET /api/blog/tag/:tag           # Posts by tag
```

#### BlueSky
```
GET /api/bsky/posts              # Get BlueSky posts
GET /api/bsky/post?uri=...       # Get specific post
```

#### Profile
```
GET /api/profile                 # Get public profile
```

#### Games
```
GET /api/games                   # List all games
GET /api/games/:slug             # Get game details
GET /api/games/:slug/leaderboard # Get leaderboard
POST /api/games/:slug/score      # Submit score
```

#### Search
```
GET /api/search?query=...        # Search content
GET /api/search/autocomplete?q=... # Autocomplete suggestions
```

#### Media
```
GET /api/media                   # List media
GET /api/media/:id               # Get media item
GET /api/media/categories        # List categories
```

### Protected Endpoints (Admin Only)

All admin endpoints require JWT authentication via `Authorization: Bearer <token>` header.

#### Authentication
```
POST /api/auth/login             # Admin login
POST /api/auth/logout            # Logout
GET /api/auth/status             # Check auth status
```

#### Admin Operations
```
POST /api/blog/rescan            # Rescan blog posts
POST /api/bsky/refresh           # Refresh BlueSky posts
POST /api/bsky/profile/refresh   # Refresh BlueSky profile
PUT /api/profile                 # Update profile
POST /api/media/upload           # Upload media
POST /api/games                  # Create game
PUT /api/games/:slug             # Update game
POST /api/search/rebuild         # Rebuild search index
```

## Development

### Running in Development Mode

With auto-reload (requires [air](https://github.com/cosmtrek/air)):
```bash
make dev
```

### Running Tests
```bash
make test
```

### Code Formatting
```bash
make fmt
```

### Linting (requires golangci-lint)
```bash
make lint
```

## Docker Deployment

### Build Docker Image
```bash
make docker-build
```

### Run with Docker
```bash
make docker-run
```

### Run with Docker Compose
```bash
make docker-dev
```

### Stop Docker Compose
```bash
make docker-stop
```

## Database Management

### Backup Database
```bash
make backup
```

This creates a timestamped backup in the `backups/` directory.

### Database Schema

The database schema is automatically created/updated using GORM's AutoMigrate feature. Key tables:

- `blog_posts` - Blog post content and metadata
- `bsky_posts` - Cached BlueSky posts
- `profiles` - User profile information
- `games` - Game definitions
- `game_scores` - Game leaderboard scores
- `media_items` - Uploaded media metadata
- `search_indices` - Full-text search index
- `analytics_events` - Analytics tracking
- `notifications` - User notifications

## I2P Deployment

### I2P Tunnel Configuration

1. In your I2P router console, create a new HTTP server tunnel:
   - Target host: 127.0.0.1
   - Target port: 8080
   - Tunnel length: 3 (inbound/outbound)
   - Tunnel quantity: 3

2. Save your tunnel's `.b32.i2p` address

3. Update CORS settings in `.env`:
```bash
CORS_ORIGIN_2=http://your-address.b32.i2p
```

### Systemd Service (Linux)

Create `/etc/systemd/system/lilithinaparka.service`:

```ini
[Unit]
Description=lilithinaparka.i2p Backend
After=network.target i2p.service
Requires=i2p.service

[Service]
Type=simple
User=your-user
WorkingDirectory=/path/to/backend
ExecStart=/path/to/backend/dist/server
Restart=on-failure
RestartSec=5
Environment="HOST=127.0.0.1"
Environment="PORT=8080"

[Install]
WantedBy=multi-user.target
```

Enable and start:
```bash
sudo systemctl enable lilithinaparka
sudo systemctl start lilithinaparka
sudo systemctl status lilithinaparka
```

## Troubleshooting

### Database Locked Errors
- Ensure only one instance is running
- Check that WAL mode is enabled
- Increase busy timeout in config

### BlueSky Integration Fails
- Verify BSKY_HANDLE is correct
- Check network connectivity
- Review API rate limits

### CORS Errors
- Verify CORS_ORIGIN settings match your frontend URL exactly
- Check that the backend is running and accessible

### High Memory Usage
- Adjust rate limiting settings
- Clear old analytics events periodically
- Monitor SSE client connections

## Performance Optimization

### Database
- WAL mode is enabled by default for better concurrency
- Indexes are automatically created for frequently queried fields
- Consider periodic VACUUM operations for large databases

### Caching
- Static assets are served with appropriate cache headers
- Consider adding Redis for session storage in high-traffic scenarios

### Rate Limiting
- Adjust `RATE_LIMIT` and `RATE_LIMIT_BURST` based on expected traffic
- Consider IP-based or endpoint-specific limits

## Security Considerations

### Production Checklist
- [ ] Change all default passwords in `.env`
- [ ] Set strong JWT_SECRET (32+ random characters)
- [ ] Set unique IP_SALT for hashing
- [ ] Configure proper CORS origins
- [ ] Enable HTTPS (via reverse proxy)
- [ ] Set up regular database backups
- [ ] Monitor logs for suspicious activity
- [ ] Keep dependencies updated
- [ ] Run behind I2P tunnel for anonymity

### Admin Access
- Admin credentials are stored securely with bcrypt hashing
- JWT tokens expire after 24 hours
- All admin endpoints require authentication

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests: `make test`
5. Format code: `make fmt`
6. Submit a pull request

## License

See main project README for license information.

## Support

For issues and questions:
- Check the troubleshooting section
- Review server logs
- Open an issue on GitHub (if public)
- Contact via I2P email (see profile)