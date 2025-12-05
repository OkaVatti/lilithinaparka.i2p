# Lilith Parker's I2P Blog - Backend

A Go backend server that scans markdown blog posts, fetches BlueSky data, and serves content via REST API.

## Tech Stack

- **Language**: Go 1.25.4
- **Framework**: Echo v4
- **ORM**: GORM
- **Database**: SQLite
- **APIs**: BlueSky AT Protocol

## Features

- Markdown blog post scanning with frontmatter parsing
- SQLite database for content storage
- BlueSky API integration for social posts
- Profile management from static config files
- Automatic background sync (every 5 minutes)
- RESTful API endpoints
- CORS enabled for frontend integration

## Setup

1. Install Go dependencies:
```bash
go mod download
```

2. Create the blog directory structure:
```bash
mkdir -p blog/posts/Casual
mkdir -p blog/posts/Interlude
mkdir -p blog/posts/Serious
mkdir -p blog/profile
mkdir -p blog/assets/images
```

3. Create `blog/profile/info.txt` with your profile information

4. Add markdown blog posts to the appropriate category directories

5. Run the server:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

### Blog Posts
- `GET /api/blog/posts` - Get all blog posts (query: `?drafts=true`)
- `GET /api/blog/posts/:slug` - Get post by slug
- `GET /api/blog/category/:category` - Get posts by category
- `GET /api/blog/tag/:tag` - Get posts by tag
- `POST /api/blog/rescan` - Manually trigger blog post rescan

### BlueSky
- `GET /api/bsky/posts` - Get BlueSky posts (query: `?limit=50`)
- `GET /api/bsky/post` - Get specific post (query: `?uri=...`)
- `POST /api/bsky/refresh` - Manually refresh BlueSky posts

### Profile
- `GET /api/profile` - Get profile information
- `POST /api/profile/refresh` - Manually refresh BlueSky profile data

### Health
- `GET /api/health` - Health check endpoint

## Database Schema

### BlogPost
- Stores parsed markdown posts with frontmatter
- Includes metadata like title, date, tags, categories
- Tracks draft status and file paths

### BskyPost
- Stores BlueSky posts fetched via API
- Includes engagement metrics (likes, reposts, etc.)
- Handles media attachments

### Profile
- Stores static profile info from `info.txt`
- Includes BlueSky profile data
- Crypto donation addresses

## Background Sync

The server automatically syncs data every 5 minutes:
1. Rescans blog posts for changes
2. Fetches new BlueSky posts
3. Updates BlueSky profile information

## Markdown Format

Blog posts should follow this frontmatter format:
```markdown
---
title: Post Title
date: 2025-12-3
time: 09:13
authors: [Author Name]
tags: [tag1, tag2]
categories: [Casual]
draft: false
share: true
slug: post-slug
layout: post
toc: true
comments: true
math: false
featured_image: "path/to/image.png"
summary: "Post summary"
---

# Post Content

Your markdown content here...
```

## Deployment

For production:
1. Build the binary: `go build -o blog-server main.go`
2. Ensure the `blog/` directory is present with content
3. Run the server: `./blog-server`
4. Configure I2P tunnel to expose port 8080

## License

MIT