package transaction

import "be_crowdfunding/user"

type GetCampaignTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}

type CreateTransactionsInput struct {
	Amount     int `uri:"amount" binding:"required"`
	CampaignID int `uri:"campaign_id" binding:"required"`
	User       user.User
}
