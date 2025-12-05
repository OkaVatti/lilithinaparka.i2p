package utils

import (
	"bufio"
	"bytes"
	"strings"

	"gopkg.in/yaml.v3"
)

type FrontMatter struct {
	Title            string   `yaml:"title"`
	Date             string   `yaml:"date"`
	Time             string   `yaml:"time"`
	Authors          []string `yaml:"authors"`
	Tags             []string `yaml:"tags"`
	Categories       []string `yaml:"categories"`
	Draft            bool     `yaml:"draft"`
	Share            bool     `yaml:"share"`
	Slug             string   `yaml:"slug"`
	Layout           string   `yaml:"layout"`
	TOC              bool     `yaml:"toc"`
	Comments         bool     `yaml:"comments"`
	Math             bool     `yaml:"math"`
	FeaturedImage    string   `yaml:"featured_image"`
	FeaturedImageAlt string   `yaml:"featured_image_alt"`
	FeaturedVideo    string   `yaml:"featured_video"`
	FeaturedVideoAlt string   `yaml:"featured_video_alt"`
	Summary          string   `yaml:"summary"`
}

func ParseMarkdownWithFrontMatter(content []byte) (*FrontMatter, string, error) {
	scanner := bufio.NewScanner(bytes.NewReader(content))

	var frontMatterLines []string
	var contentLines []string
	inFrontMatter := false
	frontMatterCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "---" {
			frontMatterCount++
			if frontMatterCount == 1 {
				inFrontMatter = true
				continue
			} else if frontMatterCount == 2 {
				inFrontMatter = false
				continue
			}
		}

		if inFrontMatter {
			frontMatterLines = append(frontMatterLines, line)
		} else if frontMatterCount >= 2 {
			contentLines = append(contentLines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, "", err
	}

	var fm FrontMatter
	if len(frontMatterLines) > 0 {
		yamlContent := strings.Join(frontMatterLines, "\n")
		if err := yaml.Unmarshal([]byte(yamlContent), &fm); err != nil {
			return nil, "", err
		}
	}

	markdownContent := strings.Join(contentLines, "\n")
	return &fm, markdownContent, nil
}
