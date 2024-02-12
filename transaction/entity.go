package transaction

import (
	"be_crowdfunding/campaign"
	"be_crowdfunding/user"
	"time"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	PaymentURL string
	Campaign   campaign.Campaign
	User       user.User
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
