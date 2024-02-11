package transaction

import (
	"be_crowdfunding/campaign"
	"errors"
)

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

type Service interface {
	GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
	GetTransactionByUserID(userID int) ([]Transaction, error)
	CreateTransactions(input CreateTransactionsInput) (Transaction, error)
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error) {
	// get campaign
	//check campaign.userID != user_id yang melakukan request

	campaign, err := s.campaignRepository.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("Not an owner of a campaign")
	}
	transactions, err := s.repository.GetCampaignID(input.ID)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
func (s *service) GetTransactionByUserID(userID int) ([]Transaction, error) {
	transactions, err := s.repository.GetByUserID(userID)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (s *service) CreateTransactions(input CreateTransactionsInput) (Transaction, error) {
	transaction := Transaction{}
	transaction.CampaignID = input.CampaignID
	transaction.Amount = input.Amount
	transaction.UserID = input.User.ID
	transaction.Status = "pending"
	// transaction.Code = ""
	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}
	return newTransaction, nil
}
