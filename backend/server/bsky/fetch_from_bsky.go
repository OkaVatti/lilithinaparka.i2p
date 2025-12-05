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

type BskyFeedResponse struct {
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
				} `json:"images"`
			} `json:"embed"`
		} `json:"post"`
	} `json:"feed"`
}

type BskyProfileResponse struct {
	DisplayName    string `json:"displayName"`
	Description    string `json:"description"`
	Avatar         string `json:"avatar"`
	Banner         string `json:"banner"`
	FollowersCount int    `json:"followersCount"`
	FollowsCount   int    `json:"followsCount"`
	PostsCount     int    `json:"postsCount"`
}

func FetchBskyPosts(db *gorm.DB, handle string) error {
	url := fmt.Sprintf("https://public.api.bsky.app/xrpc/app.bsky.feed.getAuthorFeed?actor=%s&limit=50", handle)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch BlueSky posts: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("BlueSky API returned status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	var feedResponse BskyFeedResponse
	if err := json.Unmarshal(body, &feedResponse); err != nil {
		return fmt.Errorf("failed to parse BlueSky response: %w", err)
	}

	for _, item := range feedResponse.Feed {
		post := item.Post

		var mediaURLs []string
		hasMedia := false
		if post.Embed != nil && post.Embed.Images != nil {
			hasMedia = true
			for _, img := range post.Embed.Images {
				mediaURLs = append(mediaURLs, img.Fullsize)
			}
		}

		mediaJSON, _ := json.Marshal(mediaURLs)

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
			HasMedia:    hasMedia,
			MediaURLs:   string(mediaJSON),
		}

		var existing models.BskyPost
		result := db.Where("uri = ?", post.URI).First(&existing)

		if result.Error == gorm.ErrRecordNotFound {
			if err := db.Create(&bskyPost).Error; err != nil {
				return fmt.Errorf("failed to create BlueSky post: %w", err)
			}
			fmt.Printf("Added BlueSky post: %s\n", post.URI)
		} else if result.Error == nil {
			bskyPost.ID = existing.ID
			if err := db.Save(&bskyPost).Error; err != nil {
				return fmt.Errorf("failed to update BlueSky post: %w", err)
			}
			fmt.Printf("Updated BlueSky post: %s\n", post.URI)
		} else {
			return fmt.Errorf("database error: %w", result.Error)
		}
	}

	return nil
}

func FetchBskyProfile(db *gorm.DB, handle string) error {
	url := fmt.Sprintf("https://public.api.bsky.app/xrpc/app.bsky.actor.getProfile?actor=%s", handle)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch BlueSky profile: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("BlueSky API returned status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	var profileResponse BskyProfileResponse
	if err := json.Unmarshal(body, &profileResponse); err != nil {
		return fmt.Errorf("failed to parse BlueSky profile response: %w", err)
	}

	var profile models.Profile
	result := db.First(&profile)

	if result.Error == gorm.ErrRecordNotFound {
		profile = models.Profile{}
	} else if result.Error != nil {
		return fmt.Errorf("database error: %w", result.Error)
	}

	profile.BskyDisplayName = profileResponse.DisplayName
	profile.BskyDescription = profileResponse.Description
	profile.BskyAvatar = profileResponse.Avatar
	profile.BskyBanner = profileResponse.Banner
	profile.BskyFollowersCount = profileResponse.FollowersCount
	profile.BskyFollowsCount = profileResponse.FollowsCount
	profile.BskyPostsCount = profileResponse.PostsCount

	if profile.ID == 0 {
		if err := db.Create(&profile).Error; err != nil {
			return fmt.Errorf("failed to create profile: %w", err)
		}
	} else {
		if err := db.Save(&profile).Error; err != nil {
			return fmt.Errorf("failed to update profile: %w", err)
		}
	}

	fmt.Println("Updated BlueSky profile information")
	return nil
}
