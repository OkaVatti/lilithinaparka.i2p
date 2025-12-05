package blog

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/models"
	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/utils"
	"gorm.io/gorm"
)

func ScanBlogPosts(db *gorm.DB, blogDir string) error {
	categories := []string{"Casual", "Interlude", "Serious"}

	for _, category := range categories {
		categoryPath := filepath.Join(blogDir, "posts", category)

		err := filepath.Walk(categoryPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
				content, err := os.ReadFile(path)
				if err != nil {
					return fmt.Errorf("failed to read file %s: %w", path, err)
				}

				fm, markdownContent, err := utils.ParseMarkdownWithFrontMatter(content)
				if err != nil {
					return fmt.Errorf("failed to parse front matter in %s: %w", path, err)
				}

				authorsJSON, _ := json.Marshal(fm.Authors)
				tagsJSON, _ := json.Marshal(fm.Tags)
				categoriesJSON, _ := json.Marshal(fm.Categories)

				blogPost := models.BlogPost{
					Title:            fm.Title,
					Date:             fm.Date,
					Time:             fm.Time,
					Authors:          string(authorsJSON),
					Tags:             string(tagsJSON),
					Categories:       string(categoriesJSON),
					Draft:            fm.Draft,
					Share:            fm.Share,
					Slug:             fm.Slug,
					Layout:           fm.Layout,
					TOC:              fm.TOC,
					Comments:         fm.Comments,
					Math:             fm.Math,
					FeaturedImage:    fm.FeaturedImage,
					FeaturedImageAlt: fm.FeaturedImageAlt,
					FeaturedVideo:    fm.FeaturedVideo,
					FeaturedVideoAlt: fm.FeaturedVideoAlt,
					Summary:          fm.Summary,
					Content:          markdownContent,
					FilePath:         path,
				}

				var existing models.BlogPost
				result := db.Where("file_path = ?", path).First(&existing)

				if result.Error == gorm.ErrRecordNotFound {
					if err := db.Create(&blogPost).Error; err != nil {
						return fmt.Errorf("failed to create blog post from %s: %w", path, err)
					}
					fmt.Printf("Added blog post: %s\n", fm.Title)
				} else if result.Error == nil {
					blogPost.ID = existing.ID
					if err := db.Save(&blogPost).Error; err != nil {
						return fmt.Errorf("failed to update blog post from %s: %w", path, err)
					}
					fmt.Printf("Updated blog post: %s\n", fm.Title)
				} else {
					return fmt.Errorf("database error for %s: %w", path, result.Error)
				}
			}

			return nil
		})

		if err != nil {
			return fmt.Errorf("failed to scan category %s: %w", category, err)
		}
	}

	return nil
}
