// server/bsky/enhanced.go
package bsky

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/models"
	"gorm.io/gorm"
)

type BskyEnhanced struct {
	DB        *gorm.DB
	Handle    string
	RateLimit time.Duration
}

func NewBskyEnhanced(db *gorm.DB, handle string) *BskyEnhanced {
	return &BskyEnhanced{
		DB:        db,
		Handle:    handle,
		RateLimit: time.Second * 2, // Rate limit to avoid API throttling
	}
}

func (be *BskyEnhanced) SyncProfileAndPosts() error {
	fmt.Println("Starting BlueSky sync...")

	// Sync profile
	if err := be.syncProfile(); err != nil {
		fmt.Printf("Warning: Failed to sync profile: %v\n", err)
	}

	// Sync posts
	if err := be.syncPosts(); err != nil {
		fmt.Printf("Warning: Failed to sync posts: %v\n", err)
	}

	fmt.Println("BlueSky sync completed")
	return nil
}

func (be *BskyEnhanced) syncProfile() error {
	url := fmt.Sprintf("https://public.api.bsky.app/xrpc/app.bsky.actor.getProfile?actor=%s", be.Handle)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch profile: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("BlueSky API returned status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	var profileResponse struct {
		DisplayName    string `json:"displayName"`
		Description    string `json:"description"`
		Avatar         string `json:"avatar"`
		Banner         string `json:"banner"`
		FollowersCount int    `json:"followersCount"`
		FollowsCount   int    `json:"followsCount"`
		PostsCount     int    `json:"postsCount"`
	}

	if err := json.Unmarshal(body, &profileResponse); err != nil {
		return fmt.Errorf("failed to parse profile response: %w", err)
	}

	// Update or create profile
	var profile models.Profile
	if err := be.DB.First(&profile).Error; err != nil {
		// Create new profile
		profile = models.Profile{}
	}

	profile.BskyDisplayName = profileResponse.DisplayName
	profile.BskyDescription = profileResponse.Description
	profile.BskyAvatar = profileResponse.Avatar
	profile.BskyBanner = profileResponse.Banner
	profile.BskyFollowersCount = profileResponse.FollowersCount
	profile.BskyFollowsCount = profileResponse.FollowsCount
	profile.BskyPostsCount = profileResponse.PostsCount

	if profile.ID == 0 {
		if err := be.DB.Create(&profile).Error; err != nil {
			return fmt.Errorf("failed to create profile: %w", err)
		}
	} else {
		if err := be.DB.Save(&profile).Error; err != nil {
			return fmt.Errorf("failed to update profile: %w", err)
		}
	}

	fmt.Printf("Profile synced: %s\n", profileResponse.DisplayName)
	return nil
}

func (be *BskyEnhanced) syncPosts() error {
	url := fmt.Sprintf("https://public.api.bsky.app/xrpc/app.bsky.feed.getAuthorFeed?actor=%s&limit=100", be.Handle)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch posts: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("BlueSky API returned status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	var feedResponse struct {
		Feed []struct {
			Post struct {
				URI    string `json:"uri"`
				CID    string `json:"cid"`
				Author struct {
					Handle      string `json:"handle"`
					DisplayName string `json:"displayName"`
				} `json:"author"`
				Record struct {
					Text      string    `json:"text"`
					CreatedAt time.Time `json:"createdAt"`
				} `json:"record"`
				ReplyCount  int `json:"replyCount"`
				RepostCount int `json:"repostCount"`
				LikeCount   int `json:"likeCount"`
				QuoteCount  int `json:"quoteCount"`
				Embed       *struct {
					Images []struct {
						Fullsize string `json:"fullsize"`
						Alt      string `json:"alt"`
					} `json:"images"`
				} `json:"embed"`
			} `json:"post"`
		} `json:"feed"`
	}

	if err := json.Unmarshal(body, &feedResponse); err != nil {
		return fmt.Errorf("failed to parse feed response: %w", err)
	}

	postCount := 0
	for _, item := range feedResponse.Feed {
		post := item.Post

		// Check if post already exists
		var existing models.BskyPost
		result := be.DB.Where("uri = ?", post.URI).First(&existing)

		// Process media
		var mediaURLs []string
		var mediaAlts []string
		if post.Embed != nil && post.Embed.Images != nil {
			for _, img := range post.Embed.Images {
				mediaURLs = append(mediaURLs, img.Fullsize)
				mediaAlts = append(mediaAlts, img.Alt)
			}
		}

		mediaData := map[string]interface{}{
			"urls": mediaURLs,
			"alts": mediaAlts,
		}
		mediaJSON, _ := json.Marshal(mediaData)

		bskyPost := models.BskyPost{
			URI:         post.URI,
			CID:         post.CID,
			Author:      post.Author.Handle,
			Text:        post.Record.Text,
			PostedAt:    post.Record.CreatedAt,
			ReplyCount:  post.ReplyCount,
			RepostCount: post.RepostCount,
			LikeCount:   post.LikeCount,
			QuoteCount:  post.QuoteCount,
			HasMedia:    len(mediaURLs) > 0,
			MediaURLs:   string(mediaJSON),
		}

		if result.Error != nil {
			// Create new post
			if err := be.DB.Create(&bskyPost).Error; err != nil {
				fmt.Printf("Warning: Failed to create post %s: %v\n", post.URI, err)
				continue
			}
		} else {
			// Update existing post
			bskyPost.ID = existing.ID
			if err := be.DB.Save(&bskyPost).Error; err != nil {
				fmt.Printf("Warning: Failed to update post %s: %v\n", post.URI, err)
				continue
			}
		}

		postCount++
		time.Sleep(be.RateLimit) // Rate limiting
	}

	fmt.Printf("Synced %d BlueSky posts\n", postCount)
	return nil
}

func (be *BskyEnhanced) GetFeedStats() (map[string]interface{}, error) {
	var totalPosts int64
	var totalLikes int64
	var totalReposts int64

	be.DB.Model(&models.BskyPost{}).Count(&totalPosts)

	// Calculate totals
	var posts []models.BskyPost
	if err := be.DB.Find(&posts).Error; err != nil {
		return nil, err
	}

	for _, post := range posts {
		totalLikes += int64(post.LikeCount)
		totalReposts += int64(post.RepostCount)
	}

	return map[string]interface{}{
		"total_posts":   totalPosts,
		"total_likes":   totalLikes,
		"total_reposts": totalReposts,
		"handle":        be.Handle,
		"last_sync":     time.Now().Format(time.RFC3339),
	}, nil
}
