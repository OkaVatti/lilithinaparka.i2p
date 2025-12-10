package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// BlueSkyItem is a simplified representation
type BlueSkyItem struct {
	ID      string    `json:"id"`
	Handle  string    `json:"handle"`
	Text    string    `json:"text"`
	Created time.Time `json:"created_at"`
}

// fetchBlueSky fetches posts from Bluesky API for the configured handle.
// Requires env BLUESKY_USER and BLUESKY_TOKEN (optional).
func fetchBlueSky() ([]BlueSkyItem, error) {
	user := os.Getenv("BLUESKY_USER")
	token := os.Getenv("BLUESKY_TOKEN")
	if user == "" {
		return nil, fmt.Errorf("BLUESKY_USER not set")
	}
	// The BlueSky API is under rapid change; this performs a basic request to public feed
	// NOTE: change endpoint if the API evolves. This is a best-effort example.
	url := fmt.Sprintf("https://bsky.social/xrpc/app.bsky.feed.getAuthorFeed?actor=%s", user)
	req, _ := http.NewRequest("GET", url, nil)
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	// We conservatively try to unmarshal items; structures vary in official API.
	var parsed map[string]any
	if err := json.Unmarshal(body, &parsed); err != nil {
		return nil, err
	}
	items := []BlueSkyItem{}
	// Try common shapes: { feed: { threads: [{ post: { text, etc } }] } }
	if feed, ok := parsed["feed"].(map[string]any); ok {
		if threads, ok := feed["threads"].([]any); ok {
			for _, t := range threads {
				if tm, ok := t.(map[string]any); ok {
					if post, ok := tm["post"].(map[string]any); ok {
						text := ""
						if txt, ok := post["text"].(string); ok {
							text = txt
						}
						created := time.Now()
						if cstr, ok := post["createdAt"].(string); ok {
							if ts, err := time.Parse(time.RFC3339, cstr); err == nil {
								created = ts
							}
						}
						id := ""
						if idv, ok := post["uri"].(string); ok {
							id = idv
						}
						items = append(items, BlueSkyItem{
							ID:      id,
							Handle:  user,
							Text:    text,
							Created: created,
						})
					}
				}
			}
		}
	}
	return items, nil
}
