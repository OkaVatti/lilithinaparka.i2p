# lilithinaparka.i2p
- Lilith's I2P Blog - Frontend

A minimal, TUI-inspired blog frontend built with Nuxt 4, Deno, and TailwindCSS v4.

## Tech Stack

- **Framework**: Nuxt 4
- **Runtime**: Deno
- **State Management**: Pinia
- **Styling**: TailwindCSS v4
- **Language**: TypeScript
- **Markdown**: marked

## Features

- Dynamic markdown rendering from SQLite backend
- Real-time blog post updates
- BlueSky social feed integration
- Profile page with dynamic info fetched from the database.
- Terminal/retro aesthetic inspired by classic hacker sites
- Responsive design
- SEO-friendly

## Setup

1. Install dependencies:
```bash
npm install
```

2. Configure environment variables:
Create a `.env` file with:
```
NUXT_PUBLIC_API_BASE=http://localhost:8080/api
```

3. Run development server:
```bash
deno run dev
```

4. Build for production:
```bash
deno run build
```

5. Preview production build:
```bash
deno run preview
```

## Project Structure

##### Backend

```
backend/
├── cmd/
│   ├── server/         # Main server binary
│   ├── migrate/        # Database migrations
│   ├── seed/           # Database seeding
│   ├── watcher/        # File system watcher for content
│   └── test/           # Test utilities
├── internal/
│   ├── api/
│   │   ├── handlers/   # HTTP request handlers
│   │   ├── middleware/ # CORS, logging, auth middleware
│   │   └── routes/     # Route definitions
│   ├── config/
│   │   └── config.go   # Configuration management
│   ├── database/
│   │   ├── models/     # GORM models
│   │   ├── migrations/ # SQL migration files
│   │   └── repository/ # Data access layer
│   ├── services/
│   │   ├── blog/       # Blog post processing
│   │   ├── bsky/       # BlueSky integration
│   │   ├── games/      # Game state management
│   │   ├── media/      # Image/video processing
│   │   └── profile/    # Profile management
│   ├── utils/
│   │   ├── cache/      # Redis client wrapper
│   │   ├── logger/     # Structured logging
│   │   └── validator/  # Input validation
│   └── types/          # Internal type definitions
├── blog/               # Content directory (git ignored)
│   ├── assets/
│   ├── posts/
│   │   ├── 2024/
│   │   │   ├── 01/
│   │   │   │   └── welcome-to-i2p.md
│   │   │   └── metadata.json
│   │   └── _drafts/
│   ├── profile/
│   │   ├── info.yaml
│   │   ├── keys.json
│   │   └── socials.json
│   └── videos/
│       ├── transcodes/  # Generated transcodes
│       └── thumbs/      # Generated thumbnails
├── db/
│   ├── migrations/      # Goose migration files
│   └── seeds/          # Seed data
├── static/             # Static files served directly
│   ├── fonts/
│   ├── icons/
│   └── robots.txt
├── go.mod
├── go.sum
├── main.go            # Application entry point
├── Dockerfile
└── docker-compose.yml
```

##### Frontend

```
frontend/
├── /apps
|   ├── assets/
|   │   ├── css/
|   │   │   ├── main.css
|   │   │   ├── multiverse.css
|   │   │   ├── colours.css      # Color scheme definitions
|   │   │   ├── themes.css       # Theme definitions
|   │   │   ├── animations/      # Keyframe animations
|   │   │   ├── pages/           # Page-specific styles
|   │   │   ├── templates/       # Template-specific styles
|   │   │   └── components/      # Component-specific styles
|   │   ├── fonts/               # Custom font files
|   │   ├── images/              # Static images
|   │   └── sounds/              # UI sound effects
|   ├── components/
|   │   ├── Art/
|   │   │   ├── ArtPostCard.vue
|   │   │   ├── ArtPostContainer.vue
|   │   │   ├── Gallery.vue
|   │   │   ├── MasonryGrid.vue
|   │   │   └── Lightbox.vue
|   │   ├── Blog/
|   │   │   ├── BlogPostCard.vue
|   │   │   ├── BlogPostContainer.vue
|   │   │   ├── BskyPostCard.vue
|   │   │   ├── BskyThread.vue
|   │   │   ├── MarkdownRenderer.vue
|   │   │   ├── TableOfContents.vue
|   │   │   └── CodeBlock.vue
|   │   ├── Games/
|   │   │   ├── common/
|   │   │   │   ├── GameCard.vue
|   │   │   │   ├── GameContainer.vue
|   │   │   │   ├── GameList.vue
|   │   │   │   ├── GameAlias.vue
|   │   │   │   ├── Leaderboard.vue
|   │   │   │   ├── GamepadDetector.vue
|   │   │   │   └── SaveManager.vue
|   │   │   └── [game-specific directories]
|   │   ├── Profile/
|   │   │   ├── About.vue
|   │   │   ├── Donations.vue
|   │   │   ├── Onymous.vue
|   │   │   ├── Pic.vue
|   │   │   ├── PublicKeys.vue
|   │   │   ├── SocialLinks.vue
|   │   │   └── ProfileCard.vue
|   │   ├── Universal/
|   │   │   ├── Layout/
|   │   │   │   ├── Header.vue
|   │   │   │   ├── Navbar.vue
|   │   │   │   ├── Sidebar.vue
|   │   │   │   ├── Footer.vue
|   │   │   │   └── Breadcrumbs.vue
|   │   │   ├── UI/
|   │   │   │   ├── Button.vue
|   │   │   │   ├── Card.vue
|   │   │   │   ├── Modal.vue
|   │   │   │   ├── Toast.vue
|   │   │   │   ├── Tooltip.vue
|   │   │   │   ├── LoadingSpinner.vue
|   │   │   │   └── ProgressBar.vue
|   │   │   ├── Theme/
|   │   │   │   ├── ColorPicker.vue
|   │   │   │   ├── ThemePicker.vue
|   │   │   │   ├── AnimatedBackground.vue
|   │   │   │   └── CRTFilter.vue
|   │   │   └── Search/
|   │   │       ├── SearchBar.vue
|   │   │       └── SearchResults.vue
|   │   └── Video/
|   │       ├── Player.vue
|   │       ├── Playlist.vue
|   │       └── Thumbnail.vue
|   ├── composables/
|   │   ├── useApi.ts           # API client with retry logic
|   │   ├── useBsky.ts          # BlueSky integration
|   │   ├── useMarkdown.ts      # Markdown processing
|   │   ├── useGames.ts         # Game state management
|   │   ├── useTheme.ts         # Theme management
|   │   ├── useSearch.ts        # Client-side search
|   │   ├── useWebSocket.ts     # WebSocket connections
|   │   ├── useLocalStorage.ts  # Local storage wrapper
|   │   └── useDebounce.ts      # Debounce utilities
|   ├── layouts/
|   │   ├── default.vue
|   │   ├── blog.vue           # Blog-specific layout
|   │   ├── game.vue           # Game-specific layout
|   │   └── admin.vue          # Admin layout (protected)
|   └── pages/
|       ├── index.vue
|       ├── art/
|       │   ├── index.vue
|       │   └── [slug].vue
|       ├── blog/
|       │   ├── index.vue
|       │   ├── [slug].vue
|       │   ├── categories/
|       │   │   └── [category].vue
|       │   └── tags/
|       │       └── [tag].vue
|       ├── games/
|       │   ├── index.vue
|       │   ├── [game].vue
|       │   └── leaderboards.vue
|       ├── videos/
|       │   ├── index.vue
|       │   └── [video].vue
|       ├── search.vue
|       ├── settings.vue
|       └── admin/
|           ├── index.vue
|           ├── posts/
|           │   ├── index.vue
|           │   └── new.vue
|           └── media.vue
├── stores/
│   ├── art.ts
│   ├── blog.ts
│   ├── bsky.ts
│   ├── games.ts
│   ├── profile.ts
│   ├── theme.ts
│   ├── ui.ts
│   └── user.ts
├── types/
│   ├── api.ts            # API response types
│   ├── blog.ts           # Blog-related types
│   ├── games.ts          # Game-related types
│   ├── theme.ts          # Theme-related types
│   └── index.ts          # Re-exports
├── utils/
│   ├── api/
│   │   ├── client.ts     # Axios instance
│   │   └── errors.ts     # Error handling
│   ├── format/
│   │   ├── date.ts       # Date formatting
│   │   ├── markdown.ts   # Markdown utilities
│   │   └── text.ts       # Text utilities
│   ├── validation/
│   │   ├── schemas.ts    # Zod schemas
│   │   └── validators.ts # Validation functions
│   └── constants/
│       └── index.ts      # App constants
├── middleware/
│   ├── auth.ts           # Authentication middleware
│   └── rate-limit.ts     # Rate limiting
├── plugins/
│   ├── vuetify.ts        # UI component library
│   ├── analytics.ts      # Privacy-focused analytics
│   └── i2p.ts            # I2P-specific utilities
├── app.vue
├── nuxt.config.ts
├── tailwind.config.js
├── tsconfig.json
├── deno.json
├── Dockerfile
└── README.md
```


## API Endpoints

Authentication (JWT-based)
```text

POST   /api/auth/login        - Admin login
POST   /api/auth/logout       - Logout
POST   /api/auth/refresh      - Refresh token
GET    /api/auth/status       - Check auth status
```

Blog Posts
```text

GET    /api/blog/posts                    - List all posts (paginated)
GET    /api/blog/posts/:id                - Get single post by ID
GET    /api/blog/posts/slug/:slug         - Get single post by slug
GET    /api/blog/posts/category/:category - Filter by category
GET    /api/blog/posts/tag/:tag           - Filter by tag
GET    /api/blog/posts/search?q=query     - Full-text search
POST   /api/blog/posts                    - Create new post (admin)
PUT    /api/blog/posts/:id                - Update post (admin)
DELETE /api/blog/posts/:id                - Delete post (admin)
GET    /api/blog/stats                    - Blog statistics
```

BlueSky Integration
```text

GET    /api/bsky/posts                    - List cached BlueSky posts
GET    /api/bsky/posts/:id                - Get single BlueSky post
POST   /api/bsky/sync                     - Manual sync (admin)
GET    /api/bsky/profile                  - Get BlueSky profile info
GET    /api/bsky/feed                     - Get live feed (proxied)
```

Profile
```text

GET    /api/profile                       - Get public profile
GET    /api/profile/admin                 - Get full profile (admin)
PUT    /api/profile                       - Update profile (admin)
GET    /api/profile/socials               - Get social links
```

Games
```text

GET    /api/games                        - List all games
GET    /api/games/:id                    - Get game details
GET    /api/games/:id/leaderboard        - Get game leaderboard
POST   /api/games/:id/score              - Submit score
POST   /api/games/:id/save               - Save game state
GET    /api/games/:id/save               - Load game state
GET    /api/games/categories             - List game categories
GET    /api/games/tags                   - List game tags
```

Media
```text

GET    /api/media/images                 - List images
GET    /api/media/images/:id             - Get image metadata
POST   /api/media/upload                 - Upload media (admin)
DELETE /api/media/:id                    - Delete media (admin)
GET    /api/videos                       - List videos
GET    /api/videos/:id                   - Get video metadata
GET    /api/videos/:id/stream            - Stream video
GET    /api/videos/:id/thumbnail         - Get video thumbnail
```

Comments (Webmentions)
```text

GET    /api/comments/:postId             - Get comments for post
POST   /api/comments                     - Add comment (moderated)
POST   /api/webmention                   - Accept webmentions
```

System
```text

GET    /api/health                       - Health check
GET    /api/metrics                      - Prometheus metrics
GET    /api/config/themes                - Available themes
GET    /api/feed/rss                     - RSS feed
GET    /api/feed/atom                    - Atom feed
GET    /api/sitemap.xml                  - Dynamic sitemap
```

```
Database Schema
Tables
posts
sql

CREATE TABLE posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    uuid TEXT UNIQUE NOT NULL,
    title TEXT NOT NULL,
    slug TEXT UNIQUE NOT NULL,
    short_title TEXT,
    excerpt TEXT,
    content TEXT NOT NULL,
    content_html TEXT NOT NULL,
    author_id INTEGER NOT NULL,
    category TEXT,
    tags JSON DEFAULT '[]',
    status TEXT DEFAULT 'draft', -- draft, published, archived
    featured BOOLEAN DEFAULT FALSE,
    word_count INTEGER DEFAULT 0,
    reading_time INTEGER DEFAULT 0, -- in minutes
    published_at DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (author_id) REFERENCES users(id)
);

CREATE INDEX idx_posts_status ON posts(status);
CREATE INDEX idx_posts_published ON posts(published_at);
CREATE INDEX idx_posts_category ON posts(category);
CREATE VIRTUAL TABLE posts_fts USING fts5(title, excerpt, content, tags);
```

bsky_posts
```sql

CREATE TABLE bsky_posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    bsky_id TEXT UNIQUE NOT NULL,
    uri TEXT NOT NULL,
    author_did TEXT NOT NULL,
    author_handle TEXT NOT NULL,
    text TEXT NOT NULL,
    facets JSON DEFAULT '[]',
    embed JSON,
    reply_count INTEGER DEFAULT 0,
    repost_count INTEGER DEFAULT 0,
    like_count INTEGER DEFAULT 0,
    indexed_at DATETIME NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    fetched_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_bsky_author ON bsky_posts(author_did);
CREATE INDEX idx_bsky_created ON bsky_posts(created_at);
```

games
```sql

CREATE TABLE games (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    slug TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    category TEXT NOT NULL,
    tags JSON DEFAULT '[]',
    min_players INTEGER DEFAULT 1,
    max_players INTEGER DEFAULT 1,
    multiplayer_supported BOOLEAN DEFAULT FALSE,
    has_leaderboard BOOLEAN DEFAULT TRUE,
    version TEXT DEFAULT '1.0.0',
    config JSON DEFAULT '{}',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE game_scores (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    game_id INTEGER NOT NULL,
    alias TEXT NOT NULL,
    score INTEGER NOT NULL,
    level INTEGER DEFAULT 1,
    data JSON DEFAULT '{}', -- Additional game-specific data
    ip_hash TEXT, -- Hashed IP for rate limiting
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (game_id) REFERENCES games(id),
    UNIQUE(game_id, alias)
);

CREATE INDEX idx_scores_game ON game_scores(game_id, score DESC);
CREATE INDEX idx_scores_alias ON game_scores(alias);
```

media
```sql

CREATE TABLE media (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    uuid TEXT UNIQUE NOT NULL,
    filename TEXT NOT NULL,
    original_name TEXT NOT NULL,
    path TEXT NOT NULL,
    mime_type TEXT NOT NULL,
    size INTEGER NOT NULL,
    width INTEGER,
    height INTEGER,
    duration INTEGER, -- For videos, in seconds
    alt_text TEXT,
    caption TEXT,
    tags JSON DEFAULT '[]',
    uploaded_by INTEGER,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (uploaded_by) REFERENCES users(id)
);
```

users
```sql

CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE,
    password_hash TEXT NOT NULL,
    role TEXT DEFAULT 'user', -- user, author, admin
    profile JSON DEFAULT '{}',
    last_login DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

settings
```sql

CREATE TABLE settings (
    key TEXT PRIMARY KEY,
    value TEXT NOT NULL,
    type TEXT NOT NULL, -- string, number, boolean, json
    category TEXT DEFAULT 'general',
    description TEXT,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

## Styling & Theming

### Design System

Typography Scale:
```css

:root {
  --text-xs: 0.75rem;     /* 12px */
  --text-sm: 0.875rem;    /* 14px */
  --text-base: 1rem;      /* 16px */
  --text-lg: 1.125rem;    /* 18px */
  --text-xl: 1.25rem;     /* 20px */
  --text-2xl: 1.5rem;     /* 24px */
  --text-3xl: 1.875rem;   /* 30px */
  --text-4xl: 2.25rem;    /* 36px */
  --text-5xl: 3rem;       /* 48px */
}
```

Spacing Scale:
```css

:root {
  --space-1: 0.25rem;     /* 4px */
  --space-2: 0.5rem;      /* 8px */
  --space-3: 0.75rem;     /* 12px */
  --space-4: 1rem;        /* 16px */
  --space-6: 1.5rem;      /* 24px */
  --space-8: 2rem;        /* 32px */
  --space-12: 3rem;       /* 48px */
  --space-16: 4rem;       /* 64px */
}
```

### Color Schemes

Available color schemes defined in colours.css:

- Dracula (Default Dark)
- Nord (Arctic)
- Gruvbox (Warm)
- Solarized (Scientific)
- One Dark (Modern)
- Tokyo Night (Vibrant)
- Rose Pine (Elegant)
- Catppuccin (Pastel)
- Everforest (Natural)
- Kanagawa (Japanese)
- Monokai (Classic)
- Terminal (Green-on-black)

Each scheme includes:

- Primary, secondary, accent colors
- Background gradients
- Syntax highlighting colors
- Status colors (success, warning, error)

### Themes

Seasonal themes with animated ASCII backgrounds:

Winter:
- Snowfall animation
- Icy blue accents
- CRT with "cold" scanlines

Spring:
- Cherry blossom particles
- Pastel colors
- Gentle bloom effects

Summer:
- Sun rays animation
- Vibrant colors
- Heat haze effect

Fall:
- Falling leaves
- Warm amber tones
- Subtle grain texture

## Component Styling Guidelines

Blog Posts:
- Monospace font for body text
- Line height: 1.6 for readability
- Code blocks with syntax highlighting
- Subtle box shadows for depth

Games:
- Pixel-perfect canvas rendering
- Retro color palettes (C64, Apple II, etc.)
- Scanline overlays optional
- Responsive scaling maintain aspect ratio

Navigation:
- Fixed position on desktop
- Bottom navigation on mobile
- Hover effects with glow
- Active state clearly indicated

## Game Specifications
Bones (Tabletop Card Game)
    - Genre: Strategy, Deck-building
    - Players: 1-4 (I2P multiplayer)
    - Inspiration: Inscryption, Slay the Spire
    - Features:
        - Procedurally generated encounters
        - 150+ unique cards
        - 5 different factions
        - Campaign mode with story
        - Daily challenges
        - Mod support via JSON
Crypt (Dungeon Crawler)
- Genre: Roguelike, Action
- Players: 1-2 (co-op via I2P)
- Inspiration: Original Legend of Zelda, ADOM
- Features:
    - Randomly generated dungeons (10+ biomes)
    - Permadeath with meta-progression
    - 8 character classes
    - Crafting and enchanting system
    - Boss fights with unique mechanics
    - ASCII art with optional tilesets

Gems (Match-3 Puzzle)
- Genre: Puzzle, Casual
- Players: 1
- Inspiration: Bejeweled, Candy Crush
- Features:
    - Infinite level generation
    - 7 gem types with special effects
    - Power-ups and combos
    - Daily puzzle challenges
    - Global leaderboard for high scores

NetClicker (Incremental Game)
- Genre: Incremental, Simulation
- Players: 1
- Inspiration: Cookie Clicker, Universal Paperclips
- Theme: Network routing and I2P
- Features:
    - Prestige system with upgrades
    - Research tree with 50+ nodes
    - Automated network management
    - Events and achievements
    - Offline progress calculation

NekoRunner (Endless Runner)
- Genre: Arcade, Casual
- Players: 1
- Aesthetic: Japanese ASCII art
- Features:
    - Procedurally generated obstacles
    - 5 playable characters (cat breeds)
    - Power-ups and temporary abilities
    - Seasonal themes and events
    - Local high score tracking

Noire (Advanced Chess)
- Genre: Strategy, Board Game
- Players: 1-2 (I2P multiplayer)
- Board: 16x16 with special squares
- Pieces: Standard chess + 4 new piece types
- Features:
    - AI with 5 difficulty levels
    - Time controls and chess clocks
    - Move analysis and hints
    - Tournament system
    - Puzzle mode (1000+ positions

Sudoku (Logic Puzzle)
    Genre: Puzzle, Logic
    Players: 1
    Variants:
        Classic 9x9
        Samurai (overlapping)
        Killer (cage sums)
        Ronin (extra constraints; Fog of War)
    Features:
        10,000+ puzzles database
        5 difficulty levels
        Hint system with explanations
        Puzzle generator
        Daily challenge

Tetris (Block Puzzle)
- Genre: Puzzle, Arcade
- Players: 1-2 (competitive I2P)
- Inspiration: Classic Tetris, Tetris 99
- Features:
    - Multiple game modes (Marathon, Sprint, Ultra)
    - Customizable controls
    - Ghost piece and hold piece
    - Next piece preview (1-5 pieces)
    - T-spin detection and scoring
    - Local and global leaderboards

Game Engine Requirements
- 60 FPS target on modern browsers
- Input handling for keyboard, mouse, gamepad
- Sound effects via Web Audio API
- Local storage for save games
- Graceful degradation for older browsers

## Security Considerations

### Frontend Security
- Content Security Policy (CSP) strict headers
- Subresource Integrity (SRI) for CDN assets
- XSS Protection: DOMPurify for user-generated content
- CSRF Tokens for state-changing operations
- Rate Limiting: Client-side request throttling
- Input Validation: Zod schemas for all API calls

### Backend Security
- SQL Injection Prevention: Prepared statements only
- XSS Protection: Context-aware output encoding
- File Upload Security:
    - MIME type verification
    - File size limits (10MB images, 100MB videos)
    - Virus scanning via ClamAV (optional)
    - Secure file storage outside webroot
- Authentication:
    - JWT with short expiry (15 minutes)
    - Refresh tokens with rotation
    - Private Key for Authorization
    - Password hashing via Argon2id
- Rate Limiting: Per-IP and per-endpoint limits
- Headers:
    - HSTS with preload
    - X-Frame-Options: DENY
    - X-Content-Type-Options: nosniff
    - Referrer-Policy: strict-origin-when-cross-origin

### I2P-Specific Security
- Hidden Service Authentication: Optional client certificates
- Traffic Analysis Protection: All traffic over I2P
- No Clearnet Fallback: Strict I2P-only operation
- Resource Isolation: Separate databases per service
- Minimal Logging: No IP addresses, only request paths

### Data Privacy
- Anonymized Analytics: No personal identifiers
- Cookie Consent: No tracking cookies by default
- Data Minimization: Only collect essential data
- Right to Deletion: User data deletion endpoint
- Encryption at Rest: SQLite database encryption option

## I2P Deployment

### Tunnel Configuration

```ini

# i2ptunnel.config for lilithinaparka

[lilithinaparka]
type = http
host = 127.0.0.1
port = 8080
inbound.length = 3
outbound.length = 3
inbound.lengthVariance = 1
outbound.lengthVariance = 1
inbound.quantity = 3
outbound.quantity = 3
inbound.backupQuantity = 2
outbound.backupQuantity = 2
i2cp.leaseSetType = 1
i2cp.leaseSetKeyType = 0
i2cp.leaseSetPrivate = false
i2cp.reduceIdle = 1200000
i2cp.reduceQuantity = 4
i2cp.closeOnIdle = false
i2cp.closeIdleTime = 600000
i2cp.encryptLeaseSet = true
```

### Server Configuration

```yaml

# config/i2p.yaml
i2p:
  enabled: true
  host: "127.0.0.1"
  port: 8080
  base32: ".b32.i2p"  # Will be replaced with actual address
  name: "lilithinaparka"
  description: "TUI-inspired blog and games portal"
  tags: ["blog", "games", "retro", "privacy"]
  
server:
  read_timeout: 45s
  write_timeout: 45s
  idle_timeout: 60s
  max_header_bytes: 1048576
  
cache:
  static_assets: 3600  # 1 hour
  api_responses: 300   # 5 minutes
  markdown_rendered: 1800  # 30 minutes
```

### Deployment Steps

#### Build Production Assets:

```bash

# Build frontend
cd frontend
deno task build

# Build backend
cd ../backend
go build -ldflags="-s -w -X main.version=$(git describe --tags)" -o server

# Create deployment bundle
mkdir -p deploy
cp -r frontend/.output deploy/public
cp backend/server deploy/
cp -r backend/blog deploy/
cp backend/config/production.yaml deploy/config.yaml
```

#### Docker Deployment:

```dockerfile

FROM alpine:3.18

# Install dependencies
RUN apk add --no-cache \
    ca-certificates \
    sqlite \
    ffmpeg \
    tini

# Create non-root user
RUN adduser -D -u 1000 appuser

# Copy application
COPY --from=builder /app/server /app/
COPY --from=builder /app/public /app/public/
COPY --from=builder /app/blog /app/blog/
COPY --from=builder /app/config.yaml /app/

# Setup permissions
RUN chown -R appuser:appuser /app
USER appuser

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/api/health || exit 1

EXPOSE 8080
ENTRYPOINT ["/sbin/tini", "--"]
CMD ["/app/server", "--config", "/app/config.yaml"]
```

#### Systemd Service:

```ini

# /etc/systemd/system/lilithinaparka.service
[Unit]
Description=Lilith in a Parka I2P Blog
After=network.target i2p.service
Requires=i2p.service

[Service]
Type=simple
User=lilith
Group=lilith
WorkingDirectory=/opt/lilithinaparka
ExecStart=/opt/lilithinaparka/server --config /opt/lilithinaparka/config.yaml
Restart=on-failure
RestartSec=5
Environment="GOMAXPROCS=2"
StandardOutput=journal
StandardError=journal
SyslogIdentifier=lilithinaparka

# Security hardening
NoNewPrivileges=true
PrivateTmp=true
ProtectSystem=strict
ReadWritePaths=/opt/lilithinaparka/db /opt/lilithinaparka/cache
ProtectHome=true
ProtectKernelTunables=true
ProtectKernelModules=true
ProtectControlGroups=true

[Install]
WantedBy=multi-user.target
```

#### Update I2P Tunnel:

```bash

# Register .i2p address
i2prouter register <base64_private_key> lilithinaparka

# Verify tunnel is running
i2prouter tunnels
```

### Monitoring & Analytics

#### Health Metrics
```go

type Metrics struct {
    Uptime               time.Duration
    TotalRequests        int64
    RequestRate          float64   // requests per second
    ErrorRate            float64   // errors per second
    ActiveConnections    int
    DatabaseConnections  int
    MemoryUsage          uint64    // in bytes
    GoroutineCount       int
    ResponseTime         Histogram // percentiles
    CacheHitRate         float64
}
```

Prometheus Metrics
- http_requests_total (counter)
- http_request_duration_seconds (histogram)
- database_queries_total (counter)
- cache_hits_total (counter)
- cache_misses_total (counter)
- memory_usage_bytes (gauge)
- goroutine_count (gauge)

### Logging

#### Structured JSON logging with different levels:
```json

{
  "timestamp": "2024-01-15T10:30:00Z",
  "level": "INFO",
  "service": "blog-api",
  "method": "GET",
  "path": "/api/blog/posts",
  "status": 200,
  "duration_ms": 45.2,
  "user_agent": "Lynx/2.8.9",
  "request_id": "req_abc123"
}
```

#### Alerting Rules
- Error rate > 5% for 5 minutes
- Response time p95 > 2 seconds
- Memory usage > 80% for 10 minutes
- Database connections > 90% of max

## Roadmap

### Phase 1: Core Blog
- Basic markdown rendering
- SQLite backend
- TUI aesthetic
- Profile page
- Basic theming

### Phase 2: Content Expansion
- Art gallery with EXIF display
- Video streaming backend
- BlueSky integration
- Search functionality
- RSS/Atom feeds

### Phase 3: Interactive Features
- First 3 games (Bones, Crypt, Gems)
- Leaderboard system
- Comment system
- Enhanced theming
- Offline support

### Phase 4: Advanced Features
- Remaining 5 games
- I2P multiplayer
- Admin interface
- Analytics dashboard
- API documentation

### Phase 5: Polish & Scale
- Performance optimizations
- Accessibility audit
- Internationalization
- Federation protocols
- Mobile applications

## Contributing

### Development Process
- Fork the repository
- Create a feature branch: git checkout -b feature/amazing-feature
- Commit changes: git commit -m 'Add amazing feature'
- Push to branch: git push origin feature/amazing-feature
- Open a Pull Request

### Code Style
- Frontend: ESLint + Prettier configuration provided
- Backend: gofmt with goimports
- Commits: Conventional Commits specification
- Documentation: Code comments for exported functions

### Testing Requirements
- Unit tests for business logic
- Integration tests for API endpoints
- E2E tests for critical user flows
- Performance tests for game components

### Issue Labels
- bug: Something isn't working
- enhancement: New feature or request
- documentation: Documentation improvements
- good first issue: Good for newcomers
- help wanted: Extra attention needed