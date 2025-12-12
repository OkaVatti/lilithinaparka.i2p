// cmd/bsky_sync/main.go
package main

import (
	"log"

	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/bsky"
	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	cfg := config.LoadConfig()

	db, err := gorm.Open(sqlite.Open(cfg.Database.Path), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	enhanced := bsky.NewBskyEnhanced(db, cfg.External.BskyHandle)
	if err := enhanced.SyncProfileAndPosts(); err != nil {
		log.Fatal(err)
	}

	// Print stats
	stats, err := enhanced.GetFeedStats()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("BlueSky sync completed: %+v", stats)
}
