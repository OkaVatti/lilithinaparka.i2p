// backend/server/sse.go
package utils

import (
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
	c.Response().WriteHeader(200)

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case msg := <-client.channel:
			fmt.Fprintf(c.Response(), "data: %s\n\n", msg)
			c.Response().Flush()
		case <-ticker.C:
			fmt.Fprintf(c.Response(), ": keepalive\n\n")
			c.Response().Flush()
		case <-c.Request().Context().Done():
			return nil
		}
	}
}

var sseServer = NewSSEServer()

func notifyNewPost(path string) {
	sseServer.Broadcast(fmt.Sprintf(`{"type":"new_post","path":"%s"}`, path))
}
