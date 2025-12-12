// server/blog/profile.go
package blog

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/models"
	"gorm.io/gorm"
)

func LoadProfile(db *gorm.DB, profilePath string) error {
	file, err := os.Open(profilePath)
	if err != nil {
		return fmt.Errorf("failed to open profile file: %w", err)
	}
	defer file.Close()

	profile := models.Profile{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}

		// Handle multi-line arrays
		if strings.HasPrefix(strings.TrimSpace(line), "interests = [") {
			interests, err := parseInterestsArray(scanner)
			if err != nil {
				return fmt.Errorf("failed to parse interests: %w", err)
			}
			interestsJSON, _ := json.Marshal(interests)
			profile.Interests = string(interestsJSON)
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "pic":
			profile.Pic = value
		case "name":
			profile.Name = value
		case "cake_day":
			profile.CakeDay = value
		case "bio":
			profile.Bio = value
		case "location":
			profile.Location = value
		case "timezone":
			profile.Timezone = value
		case "website":
			profile.Website = value
		case "email":
			profile.Email = value
		case "github":
			profile.Github = value
		case "bluesky":
			profile.Bluesky = value
		case "rss_feed":
			profile.RSSFeed = value
		case "bitcoin_donation_addr":
			profile.BitcoinDonationAddr = value
		case "ethereum_donation_addr":
			profile.EthereumDonationAddr = value
		case "solana_donation_addr":
			profile.SolanaDonationAddr = value
		case "monero_donation_addr":
			profile.MoneroDonationAddr = value
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading profile file: %w", err)
	}

	// Check if profile already exists
	var existing models.Profile
	result := db.First(&existing)
	if result.Error == nil {
		// Update existing profile
		profile.ID = existing.ID
		profile.CreatedAt = existing.CreatedAt
		if err := db.Save(&profile).Error; err != nil {
			return fmt.Errorf("failed to update profile: %w", err)
		}
	} else {
		// Create new profile
		if err := db.Create(&profile).Error; err != nil {
			return fmt.Errorf("failed to create profile: %w", err)
		}
	}

	fmt.Println("Profile loaded successfully")
	return nil
}

func parseInterestsArray(scanner *bufio.Scanner) ([]string, error) {
	var interests []string
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "]" {
			break
		}

		// Remove quotes and commas
		line = strings.Trim(line, "\", ")
		if line != "" {
			interests = append(interests, line)
		}
	}
	return interests, nil
}
