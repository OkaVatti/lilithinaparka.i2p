// cmd/blog_processor/main.go
package main

import (
	"log"

	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/blog"
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

	processor := blog.NewBlogProcessor(db, "./blog")
	if err := processor.ProcessAllPosts(); err != nil {
		log.Fatal(err)
	}

	log.Println("Blog processing completed")
}
