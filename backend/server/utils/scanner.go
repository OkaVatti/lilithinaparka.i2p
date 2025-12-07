package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/radovskyb/watcher"
)

var postFolders = []string{"Casual", "Interlude", "Serious"}

type frontMatter struct {
	Title         string   `yaml:"title"`
	Date          string   `yaml:"date"`
	Time          string   `yaml:"time"`
	Authors       []string `yaml:"authors"`
	Tags          []string `yaml:"tags"`
	Categories    []string `yaml:"categories"`
	Draft         bool     `yaml:"draft"`
	Share         bool     `yaml:"share"`
	Slug          string   `yaml:"slug"`
	Summary       string   `yaml:"summary"`
	FeaturedImage string   `yaml:"featured_image"`
}

func scanPosts(db *DB) error {
	base := "./blog/posts"
	for _, cat := range postFolders {
		dir := filepath.Join(base, cat)
		if err := ensure(dir); err != nil {
			return err
		}
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			return err
		}
		for _, f := range files {
			if f.IsDir() {
				continue
			}
			if !strings.HasSuffix(f.Name(), ".md") {
				continue
			}
			path := filepath.Join(dir, f.Name())
			if err := parseAndUpsertFile(db, path); err != nil {
				log.Printf("error parsing %s: %v", path, err)
			}
		}
	}

	// start watcher for live changes
	go watchFolder(db, "./blog/posts")
	return nil
}

func watchFolder(db *DB, path string) {
	w := watcher.New()
	w.SetMaxEvents(1)
	_ = w.AddRecursive(path)

	go func() {
		for {
			select {
			case event := <-w.Event:
				// only handle create/modify
				if event.Op == watcher.Write || event.Op == watcher.Create {
					p := event.Path
					if strings.HasSuffix(p, ".md") {
						if err := parseAndUpsertFile(db, p); err != nil {
							log.Println("watch parse:", err)
						} else {
							// notify SSE clients
							notifyNewPost(p)
						}
					}
				}
			case err := <-w.Error:
				log.Println("watch error:", err)
			case <-w.Closed:
				return
			}
		}
	}()

	if err := w.Start(time.Millisecond * 500); err != nil {
		log.Println("watch start err:", err)
	}
}

func parseAndUpsertFile(db *DB, path string) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	fm, md, err := splitFrontMatter(content)
	if err != nil {
		return err
	}
	var m frontMatter
	if err := yaml.Unmarshal(fm, &m); err != nil {
		return err
	}
	// build Post
	slug := m.Slug
	if slug == "" {
		_, name := filepath.Split(path)
		slug = slugFromFilename(name)
	}
	// combine date/time
	dateStr := m.Date
	if m.Time != "" {
		dateStr = fmt.Sprintf("%s %s", m.Date, m.Time)
	}
	parsedDate, _ := time.Parse("2006-1-2 15:04", dateStr)
	p := Post{
		Slug:          slug,
		Title:         m.Title,
		Date:          parsedDate,
		Authors:       joinTags(m.Authors),
		Tags:          joinTags(m.Tags),
		Categories:    joinTags(m.Categories),
		Draft:         m.Draft,
		Share:         m.Share,
		FeaturedImage: m.FeaturedImage,
		Summary:       m.Summary,
		Content:       string(md),
	}
	return db.UpsertPost(&p)
}

func splitFrontMatter(content []byte) ([]byte, []byte, error) {
	trim := bytes.TrimSpace(content)
	if !bytes.HasPrefix(trim, []byte("---")) {
		// no frontmatter; put empty
		return []byte{}, trim, nil
	}
	parts := bytes.SplitN(trim, []byte("---"), 3)
	if len(parts) < 3 {
		return nil, nil, fmt.Errorf("invalid frontmatter")
	}
	// parts[1] is YAML, parts[2] is rest
	return bytes.TrimSpace(parts[1]), bytes.TrimSpace(parts[2]), nil
}
