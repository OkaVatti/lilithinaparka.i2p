package utils

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

type SSEClient struct {
	id      string
	channel chan string
}

type SSEServer struct {
	clients map[string]*SSEClient
	mutex   sync.RWMutex
}

func NewSSEServer() *SSEServer {
	return &SSEServer{
		clients: make(map[string]*SSEClient),
	}
}

func (s *SSEServer) AddClient(id string) *SSEClient {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	client := &SSEClient{
		id:      id,
		channel: make(chan string, 10),
	}
	s.clients[id] = client
	return client
}

func (s *SSEServer) RemoveClient(id string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if client, exists := s.clients[id]; exists {
		close(client.channel)
		delete(s.clients, id)
	}
}

func (s *SSEServer) Broadcast(message string) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for _, client := range s.clients {
		select {
		case client.channel <- message:
		default:
			// Channel full, skip this client
		}
	}
}

func (s *SSEServer) HandleSSE(c echo.Context) error {
	clientID := fmt.Sprintf("%d", time.Now().UnixNano())
	client := s.AddClient(clientID)
	defer s.RemoveClient(clientID)

	c.Response().Header().Set("Content-Type", "text/event-stream")
	c.Response().Header().Set("Cache-Control", "no-cache")
	c.Response().Header().Set("Connection", "keep-alive")
	c.Response().Header().Set("X-Accel-Buffering", "no")
	c.Response().WriteHeader(200)

	// Send initial connection message
	fmt.Fprintf(c.Response(), "data: %s\n\n", `{"type":"connected","client_id":"`+clientID+`"}`)
	c.Response().Flush()

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case msg := <-client.channel:
			fmt.Fprintf(c.Response(), "data: %s\n\n", msg)
			c.Response().Flush()
		case <-ticker.C:
			// Send keepalive
			fmt.Fprintf(c.Response(), ": keepalive\n\n")
			c.Response().Flush()
		case <-c.Request().Context().Done():
			return nil
		}
	}
}

// Helper functions for common notifications

func NotifyNewPost(path string) {
	data := map[string]string{
		"type": "new_post",
		"path": path,
	}
	if jsonData, err := json.Marshal(data); err == nil {
		// This would be called from a global SSE server instance
		// Implementation depends on how you structure the app
		fmt.Printf("New post notification: %s\n", string(jsonData))
	}
}

func NotifyNewScore(gameSlug, alias string, score int64) {
	data := map[string]interface{}{
		"type":  "new_score",
		"game":  gameSlug,
		"alias": alias,
		"score": score,
	}
	if jsonData, err := json.Marshal(data); err == nil {
		fmt.Printf("New score notification: %s\n", string(jsonData))
	}
}

func NotifyBskyUpdate(postCount int) {
	data := map[string]interface{}{
		"type":       "bsky_update",
		"post_count": postCount,
	}
	if jsonData, err := json.Marshal(data); err == nil {
		fmt.Printf("BlueSky update notification: %s\n", string(jsonData))
	}
}
