package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"unicode"

	"github.com/labstack/echo/v4"
	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/models"
	"gorm.io/gorm"
)

type SearchHandlers struct {
	DB *gorm.DB
}

func NewSearchHandlers(db *gorm.DB) *SearchHandlers {
	return &SearchHandlers{DB: db}
}

type SearchRequest struct {
	Query   string   `json:"query" validate:"required,min=1"`
	Types   []string `json:"types"` // blog, media, game, page
	Limit   int      `json:"limit" validate:"min=1,max=100"`
	Page    int      `json:"page" validate:"min=1"`
	SortBy  string   `json:"sort_by"`  // relevance, date, title
	SortDir string   `json:"sort_dir"` // asc, desc
}

type SearchResponse struct {
	Results     []models.SearchResult `json:"results"`
	Total       int64                 `json:"total"`
	Page        int                   `json:"page"`
	TotalPages  int                   `json:"total_pages"`
	Query       string                `json:"query"`
	Suggestions []string              `json:"suggestions"`
}

func (h *SearchHandlers) Search(c echo.Context) error {
	var req SearchRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid search request",
		})
	}

	if req.Limit == 0 {
		req.Limit = 20
	}
	if req.Page == 0 {
		req.Page = 1
	}
	if req.SortBy == "" {
		req.SortBy = "relevance"
	}
	if req.SortDir == "" {
		req.SortDir = "desc"
	}

	// Clean and tokenize query
	query := strings.ToLower(strings.TrimSpace(req.Query))
	keywords := tokenizeQuery(query)

	if len(keywords) == 0 {
		return c.JSON(http.StatusOK, SearchResponse{
			Results:    []models.SearchResult{},
			Total:      0,
			Page:       req.Page,
			TotalPages: 0,
			Query:      query,
		})
	}

	// Build search query
	dbQuery := h.DB.Model(&models.SearchIndex{}).Where("is_public = ?", true)

	// Filter by types if specified
	if len(req.Types) > 0 {
		dbQuery = dbQuery.Where("type IN ?", req.Types)
	}

	// Build full-text search conditions
	var conditions []string
	var args []interface{}

	for _, keyword := range keywords {
		conditions = append(conditions, "(title LIKE ? OR content LIKE ?)")
		args = append(args, "%"+keyword+"%", "%"+keyword+"%")
	}

	if len(conditions) > 0 {
		dbQuery = dbQuery.Where(strings.Join(conditions, " OR "), args...)
	}

	// Count total results
	var total int64
	dbQuery.Count(&total)

	// Apply pagination
	offset := (req.Page - 1) * req.Limit
	dbQuery = dbQuery.Offset(offset).Limit(req.Limit)

	// Apply sorting
	switch req.SortBy {
	case "date":
		dbQuery = dbQuery.Order("created_at " + req.SortDir)
	case "title":
		dbQuery = dbQuery.Order("title " + req.SortDir)
	default: // relevance
		dbQuery = dbQuery.Order("weight DESC").Order("created_at DESC")
	}

	// Execute query
	var indices []models.SearchIndex
	if err := dbQuery.Find(&indices).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to execute search",
		})
	}

	// Convert to search results
	results := make([]models.SearchResult, 0, len(indices))
	for _, index := range indices {
		result := models.SearchResult{
			Type:   index.Type,
			ItemID: index.ItemID,
			Title:  index.Title,
			Slug:   index.Slug,
			Date:   index.CreatedAt,
		}

		// Generate description and highlights
		description, highlights := generateSnippet(index.Content, keywords)
		result.Description = description
		result.Highlights = highlights

		// Calculate relevance score
		result.Score = calculateRelevance(index, keywords)

		results = append(results, result)
	}

	// Generate search suggestions
	suggestions := h.generateSuggestions(query, keywords)

	response := SearchResponse{
		Results:     results,
		Total:       total,
		Page:        req.Page,
		TotalPages:  (int(total) + req.Limit - 1) / req.Limit,
		Query:       query,
		Suggestions: suggestions,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *SearchHandlers) AutoComplete(c echo.Context) error {
	query := c.QueryParam("q")
	if len(query) < 2 {
		return c.JSON(http.StatusOK, []string{})
	}

	var suggestions []struct {
		Title string `json:"title"`
		Type  string `json:"type"`
	}

	dbQuery := h.DB.Model(&models.SearchIndex{}).
		Select("DISTINCT title, type").
		Where("is_public = ? AND title LIKE ?", true, query+"%").
		Order("weight DESC").
		Limit(10)

	if err := dbQuery.Find(&suggestions).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch suggestions",
		})
	}

	// Also search in content for common phrases
	//
	//var contentSuggestions []struct {
	//	Phrase string `json:"phrase"`
	//}

	// This would require a more sophisticated approach with a phrases table
	// For now, we'll return title-based suggestions

	return c.JSON(http.StatusOK, suggestions)
}

func (h *SearchHandlers) generateSuggestions(query string, keywords []string) []string {
	// Simple suggestion logic - can be enhanced with a proper search dictionary
	var suggestions []string

	// Common misspellings and alternatives
	commonAlternatives := map[string][]string{
		"programming": {"code", "development", "software", "coding"},
		"blog":        {"article", "post", "writing"},
		"game":        {"play", "gaming", "entertainment"},
		"media":       {"image", "video", "photo", "picture"},
	}

	for _, keyword := range keywords {
		if alts, ok := commonAlternatives[keyword]; ok {
			suggestions = append(suggestions, alts...)
		}
	}

	// Remove duplicates
	seen := make(map[string]bool)
	unique := []string{}
	for _, s := range suggestions {
		if !seen[s] {
			seen[s] = true
			unique = append(unique, s)
		}
	}

	return unique
}

func (h *SearchHandlers) RebuildIndex(c echo.Context) error {
	// Clear existing index
	if err := h.DB.Exec("DELETE FROM search_indices").Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to clear search index",
		})
	}

	// Index blog posts
	var blogPosts []models.BlogPost
	if err := h.DB.Where("draft = ?", false).Find(&blogPosts).Error; err == nil {
		for _, post := range blogPosts {
			var tags []string
			json.Unmarshal([]byte(post.Tags), &tags)

			index := models.SearchIndex{
				Type:     "blog",
				ItemID:   post.ID,
				Title:    post.Title,
				Content:  post.Summary + " " + post.Content,
				Tags:     post.Tags,
				Slug:     post.Slug,
				Weight:   1.0,
				IsPublic: true,
			}
			h.DB.Create(&index)
		}
	}

	// Index games
	var games []models.Game
	if err := h.DB.Find(&games).Error; err == nil {
		for _, game := range games {
			index := models.SearchIndex{
				Type:     "game",
				ItemID:   game.ID,
				Title:    game.Name,
				Content:  game.Description,
				Tags:     game.Tags,
				Slug:     game.Slug,
				Weight:   0.8,
				IsPublic: true,
			}
			h.DB.Create(&index)
		}
	}

	// Index media items
	var media []models.MediaItem
	if err := h.DB.Where("is_public = ?", true).Find(&media).Error; err == nil {
		for _, item := range media {
			index := models.SearchIndex{
				Type:     "media",
				ItemID:   item.ID,
				Title:    item.Title,
				Content:  item.Description,
				Tags:     item.Tags,
				Slug:     item.FileName,
				Weight:   0.6,
				IsPublic: true,
			}
			h.DB.Create(&index)
		}
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Search index rebuilt successfully",
	})
}

// Helper functions
func tokenizeQuery(query string) []string {
	// Simple tokenization - split by spaces and remove punctuation
	f := func(c rune) bool {
		return unicode.IsSpace(c) || unicode.IsPunct(c)
	}
	tokens := strings.FieldsFunc(strings.ToLower(query), f)

	// Remove stop words
	stopWords := map[string]bool{
		"the": true, "and": true, "or": true, "a": true, "an": true,
		"in": true, "on": true, "at": true, "to": true, "for": true,
		"of": true, "with": true, "by": true,
	}

	var filtered []string
	for _, token := range tokens {
		if !stopWords[token] && len(token) > 1 {
			filtered = append(filtered, token)
		}
	}

	return filtered
}

func generateSnippet(content string, keywords []string) (string, []string) {
	// Create a snippet of around 200 characters with keywords highlighted
	const snippetLength = 200

	if len(content) <= snippetLength {
		return content, []string{}
	}

	// Find first occurrence of any keyword
	lowerContent := strings.ToLower(content)
	var firstPos int = -1

	for _, keyword := range keywords {
		if pos := strings.Index(lowerContent, keyword); pos != -1 {
			if firstPos == -1 || pos < firstPos {
				firstPos = pos
			}
		}
	}

	// If no keyword found, start from beginning
	if firstPos == -1 {
		firstPos = 0
	}

	// Adjust to start at word boundary
	for firstPos > 0 && content[firstPos-1] != ' ' && content[firstPos-1] != '\n' {
		firstPos--
	}

	end := firstPos + snippetLength
	if end > len(content) {
		end = len(content)
	}

	snippet := content[firstPos:end]
	if end < len(content) {
		snippet += "..."
	}

	// Generate highlights
	var highlights []string
	for _, keyword := range keywords {
		if strings.Contains(lowerContent, keyword) {
			highlights = append(highlights, keyword)
		}
	}

	return snippet, highlights
}

func calculateRelevance(index models.SearchIndex, keywords []string) float64 {
	score := index.Weight

	// Title matches are worth more
	lowerTitle := strings.ToLower(index.Title)
	for _, keyword := range keywords {
		if strings.Contains(lowerTitle, keyword) {
			score += 0.5
		}
	}

	// Tag matches
	var tags []string
	json.Unmarshal([]byte(index.Tags), &tags)
	for _, tag := range tags {
		lowerTag := strings.ToLower(tag)
		for _, keyword := range keywords {
			if strings.Contains(lowerTag, keyword) {
				score += 0.3
				break
			}
		}
	}

	return score
}
