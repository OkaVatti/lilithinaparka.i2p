// backend/server/models/game.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type Game struct {
	ID                 uint           `gorm:"primarykey" json:"id"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"-"`
	Slug               string         `gorm:"uniqueIndex;not null" json:"slug"`
	Name               string         `gorm:"not null" json:"name"`
	Description        string         `json:"description"`
	Category           string         `gorm:"not null" json:"category"`
	Tags               string         `gorm:"type:json;default:'[]'" json:"tags"`
	MinPlayers         int            `gorm:"default:1" json:"min_players"`
	MaxPlayers         int            `gorm:"default:1" json:"max_players"`
	MultiplayerSupport bool           `gorm:"default:false" json:"multiplayer_supported"`
	HasLeaderboard     bool           `gorm:"default:true" json:"has_leaderboard"`
	Version            string         `gorm:"default:'1.0.0'" json:"version"`
	Config             string         `gorm:"type:json;default:'{}'" json:"config"`
}

type GameScore struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	GameID    uint           `gorm:"not null;index" json:"game_id"`
	Alias     string         `gorm:"not null" json:"alias"`
	Score     int64          `gorm:"not null;index" json:"score"`
	Level     int            `gorm:"default:1" json:"level"`
	Data      string         `gorm:"type:json;default:'{}'" json:"data"`
	IPHash    string         `gorm:"index" json:"ip_hash"`
}

type GameSave struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	GameID    uint           `gorm:"not null;index" json:"game_id"`
	Alias     string         `gorm:"not null" json:"alias"`
	SaveData  string         `gorm:"type:text;not null" json:"save_data"`
	IPHash    string         `gorm:"index" json:"ip_hash"`
}
