package models

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	ID                   uint           `gorm:"primarykey" json:"id"`
	CreatedAt            time.Time      `json:"created_at"`
	UpdatedAt            time.Time      `json:"updated_at"`
	DeletedAt            gorm.DeletedAt `gorm:"index" json:"-"`
	Pic                  string         `json:"pic"`
	Name                 string         `json:"name"`
	CakeDay              string         `json:"cake_day"`
	Bio                  string         `json:"bio"`
	Interests            string         `json:"interests"` // JSON array as string
	Location             string         `json:"location"`
	Timezone             string         `json:"timezone"`
	Website              string         `json:"website"`
	Email                string         `json:"email"`
	Github               string         `json:"github"`
	Bluesky              string         `json:"bluesky"`
	RSSFeed              string         `json:"rss_feed"`
	BitcoinDonationAddr  string         `json:"bitcoin_donation_addr"`
	EthereumDonationAddr string         `json:"ethereum_donation_addr"`
	SolanaDonationAddr   string         `json:"solana_donation_addr"`
	MoneroDonationAddr   string         `json:"monero_donation_addr"`
	// BlueSky profile data
	BskyDisplayName    string `json:"bsky_display_name"`
	BskyDescription    string `json:"bsky_description"`
	BskyAvatar         string `json:"bsky_avatar"`
	BskyBanner         string `json:"bsky_banner"`
	BskyFollowersCount int    `json:"bsky_followers_count"`
	BskyFollowsCount   int    `json:"bsky_follows_count"`
	BskyPostsCount     int    `json:"bsky_posts_count"`
}
