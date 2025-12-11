package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/disintegration/imaging"
	"github.com/labstack/echo/v4"
	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/models"
	"github.com/rwcarlsen/goexif/exif"
	"gorm.io/gorm"
)

type MediaHandlers struct {
	DB *gorm.DB
}

func NewMediaHandlers(db *gorm.DB) *MediaHandlers {
	return &MediaHandlers{DB: db}
}

func (h *MediaHandlers) UploadMedia(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid form data",
		})
	}

	files := form.File["files"]
	var uploadedMedia []models.MediaItem

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			continue
		}
		defer file.Close()

		// Create upload directory if it doesn't exist
		uploadDir := "./media/uploads"
		if err := os.MkdirAll(uploadDir, 0755); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to create upload directory",
			})
		}

		// Generate unique filename
		ext := filepath.Ext(fileHeader.Filename)
		filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
		filePath := filepath.Join(uploadDir, filename)

		// Save the file
		dst, err := os.Create(filePath)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to save file",
			})
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to copy file",
			})
		}

		// Get file info
		fileInfo, _ := dst.Stat()

		// Process image/video
		mediaItem := models.MediaItem{
			Title:       strings.TrimSuffix(fileHeader.Filename, ext),
			FileName:    filename,
			FileSize:    fileInfo.Size(),
			MimeType:    fileHeader.Header.Get("Content-Type"),
			Category:    c.FormValue("category"),
			Tags:        c.FormValue("tags"),
			Description: c.FormValue("description"),
			Artist:      c.FormValue("artist"),
			IsPublic:    c.FormValue("is_public") != "false",
		}

		// Parse EXIF data for images
		if strings.HasPrefix(mediaItem.MimeType, "image/") {
			if err := h.processImage(&mediaItem, filePath); err != nil {
				fmt.Printf("Failed to process image: %v\n", err)
			}
		}

		// Save to database
		if err := h.DB.Create(&mediaItem).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to save media to database",
			})
		}

		uploadedMedia = append(uploadedMedia, mediaItem)
	}

	return c.JSON(http.StatusCreated, uploadedMedia)
}

func (h *MediaHandlers) processImage(media *models.MediaItem, filePath string) error {
	// Open image for processing
	img, err := imaging.Open(filePath)
	if err != nil {
		return err
	}

	// Get dimensions
	media.Width = img.Bounds().Dx()
	media.Height = img.Bounds().Dy()

	// Create thumbnail
	thumb := imaging.Resize(img, 300, 0, imaging.Lanczos)
	thumbPath := strings.Replace(filePath, "/uploads/", "/thumbnails/", 1)

	// Create thumbnail directory
	thumbDir := filepath.Dir(thumbPath)
	if err := os.MkdirAll(thumbDir, 0755); err != nil {
		return err
	}

	if err := imaging.Save(thumb, thumbPath); err != nil {
		return err
	}
	media.Thumbnail = strings.TrimPrefix(thumbPath, "./media")

	// Parse EXIF data
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	if x, err := exif.Decode(file); err == nil {
		exifData := make(map[string]interface{})

		if date, err := x.DateTime(); err == nil {
			exifData["date_time"] = date.Format(time.RFC3339)
		}

		if lat, long, err := x.LatLong(); err == nil {
			exifData["latitude"] = lat
			exifData["longitude"] = long
		}

		if model, err := x.Get(exif.Model); err == nil {
			exifData["camera_model"], _ = model.StringVal()
		}

		if maker, err := x.Get(exif.Make); err == nil {
			exifData["camera_make"], _ = maker.StringVal()
		}

		if exposure, err := x.Get(exif.ExposureTime); err == nil {
			exifData["exposure_time"], _ = exposure.Rat(0)
		}

		if fNumber, err := x.Get(exif.FNumber); err == nil {
			exifData["f_number"], _ = fNumber.Rat(0)
		}

		if iso, err := x.Get(exif.ISOSpeedRatings); err == nil {
			exifData["iso"], _ = iso.Int(0)
		}

		if exifJSON, err := json.Marshal(exifData); err == nil {
			media.EXIF = string(exifJSON)
		}
	}

	return nil
}

func (h *MediaHandlers) GetMedia(c echo.Context) error {
	query := h.DB.Model(&models.MediaItem{}).Where("is_public = ?", true)

	// Apply filters
	if category := c.QueryParam("category"); category != "" {
		query = query.Where("category = ?", category)
	}

	if tag := c.QueryParam("tag"); tag != "" {
		query = query.Where("tags LIKE ?", "%"+tag+"%")
	}

	if search := c.QueryParam("search"); search != "" {
		query = query.Where("title LIKE ? OR description LIKE ? OR artist LIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	// Pagination
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}
	offset := (page - 1) * limit

	var total int64
	query.Count(&total)

	var media []models.MediaItem
	if err := query.Order("created_at DESC").Offset(offset).Limit(limit).Find(&media).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch media",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"items": media,
		"pagination": map[string]interface{}{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": (total + int64(limit) - 1) / int64(limit),
		},
	})
}

func (h *MediaHandlers) GetMediaCategories(c echo.Context) error {
	var categories []models.MediaCategory
	if err := h.DB.Where("is_public = ?", true).Order("name ASC").Find(&categories).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch categories",
		})
	}
	return c.JSON(http.StatusOK, categories)
}

func (h *MediaHandlers) GetMediaItem(c echo.Context) error {
	id := c.Param("id")

	var media models.MediaItem
	if err := h.DB.Where("id = ? AND is_public = ?", id, true).First(&media).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Media not found",
		})
	}

	// Increment view count
	h.DB.Model(&media).Update("views", media.Views+1)

	return c.JSON(http.StatusOK, media)
}
