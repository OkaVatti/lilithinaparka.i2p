package utils

import (
	"fmt"
	"io"
	"net/http"
	"time"
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

func FetchBskyData(endpoint string) ([]byte, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}
