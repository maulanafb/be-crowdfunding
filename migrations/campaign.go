package migrations

import (
	"time"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	type Campaign struct {
		ID               int
		UserID           int
		Name             string
		ShortDescription string
		Description      string
		Perks            string
		BackerCount      int
		GoalAmount       int
		CurrentAmount    int
		Slug             string
		CreatedAt        time.Time
		UpdatedAt        time.Time
	}

	type CampaignImage struct {
		ID         int
		CampaignID int
		FileName   string
		IsPrimary  int
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}

	return db.AutoMigrate(&Campaign{}, &CampaignImage{})
}
