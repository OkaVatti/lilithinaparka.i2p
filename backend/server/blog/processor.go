// server/blog/processor.go
package blog

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/models"
	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/utils"
	"gorm.io/gorm"
)

type BlogProcessor struct {
	DB      *gorm.DB
	BlogDir string
}

func NewBlogProcessor(db *gorm.DB, blogDir string) *BlogProcessor {
	return &BlogProcessor{
		DB:      db,
		BlogDir: blogDir,
	}
}

func (bp *BlogProcessor) ProcessAllPosts() error {
	fmt.Println("Processing all blog posts...")

	// Process by categories
	categories := []string{"Casual", "Interlude", "Serious"}
	totalProcessed := 0

	for _, category := range categories {
		categoryPath := filepath.Join(bp.BlogDir, "posts", category)
		count, err := bp.processCategory(categoryPath, category)
		if err != nil {
			return fmt.Errorf("error processing category %s: %w", category, err)
		}
		totalProcessed += count
		fmt.Printf("  Category %s: %d posts processed\n", category, count)
	}

	fmt.Printf("Total posts processed: %d\n", totalProcessed)
	return nil
}

func (bp *BlogProcessor) processCategory(categoryPath, category string) (int, error) {
	count := 0

	err := filepath.Walk(categoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			if err := bp.processMarkdownFile(path, category); err != nil {
				fmt.Printf("Warning: Failed to process %s: %v\n", path, err)
				return nil // Continue with other files
			}
			count++
		}
		return nil
	})

	return count, err
}

func (bp *BlogProcessor) processMarkdownFile(path, category string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", path, err)
	}

	fm, markdownContent, err := utils.ParseMarkdownWithFrontMatter(content)
	if err != nil {
		return fmt.Errorf("failed to parse front matter in %s: %w", path, err)
	}

	// Parse date and time to validate format
	_, err = time.Parse("2006-1-2", fm.Date)
	if err != nil {
		// Log warning but continue processing
		fmt.Printf("Warning: Invalid date format '%s' in %s, using current date\n", fm.Date, path)
		// Optionally set a default date
		fm.Date = time.Now().Format("2006-01-02")
	}

	// Generate slug if not provided
	slug := fm.Slug
	if slug == "" {
		slug = generateSlugFromTitle(fm.Title)
	}

	// Create blog post model
	blogPost := models.BlogPost{
		Title:            fm.Title,
		Date:             fm.Date,
		Time:             fm.Time,
		Slug:             slug,
		Categories:       fmt.Sprintf(`["%s"]`, category),
		Draft:            fm.Draft,
		Content:          markdownContent,
		Summary:          fm.Summary,
		Tags:             fmt.Sprintf(`["%s"]`, strings.Join(fm.Tags, `","`)),
		Authors:          fmt.Sprintf(`["%s"]`, strings.Join(fm.Authors, `","`)),
		FeaturedImage:    fm.FeaturedImage,
		FeaturedImageAlt: fm.FeaturedImageAlt,
		Layout:           fm.Layout,
		TOC:              fm.TOC,
		Comments:         fm.Comments,
		Math:             fm.Math,
		Share:            fm.Share,
		FilePath:         path,
	}

	// Check if post already exists
	var existing models.BlogPost
	result := bp.DB.Where("slug = ?", slug).First(&existing)

	if result.Error == nil {
		// Update existing post
		blogPost.ID = existing.ID
		blogPost.CreatedAt = existing.CreatedAt
		if err := bp.DB.Save(&blogPost).Error; err != nil {
			return fmt.Errorf("failed to update blog post: %w", err)
		}
		fmt.Printf("  Updated: %s\n", fm.Title)
	} else {
		// Create new post
		if err := bp.DB.Create(&blogPost).Error; err != nil {
			return fmt.Errorf("failed to create blog post: %w", err)
		}
		fmt.Printf("  Added: %s\n", fm.Title)
	}

	return nil
}

func generateSlugFromTitle(title string) string {
	slug := strings.ToLower(title)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, ".", "")
	slug = strings.ReplaceAll(slug, ",", "")
	slug = strings.ReplaceAll(slug, ":", "")
	slug = strings.ReplaceAll(slug, ";", "")
	slug = strings.ReplaceAll(slug, "'", "")
	slug = strings.ReplaceAll(slug, "\"", "")
	slug = strings.ReplaceAll(slug, "!", "")
	slug = strings.ReplaceAll(slug, "?", "")
	return slug
}
