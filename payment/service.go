package payment

import (
	"be_crowdfunding/user"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type service struct {
}

type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
}

func NewService() *service {
	return &service{}
}

func (service *service) GetPaymentURL(transaction Transaction, user user.User) (string, error) {
	// Load environment variables from .env
	godotenv.Load()
	// if err != nil {
	// 	return "", err
	// }

	// Read values from environment variables
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")

	// 1. Initiate Snap client
	var s = snap.Client{}
	s.New(serverKey, midtrans.Sandbox)

	// 2. Initiate Snap request
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			Email: user.Email,
			FName: user.Name,
		},
	}
	// fmt.Println("isi req", req)
	// 3. Request create Snap transaction to Midtrans
	snapResp, _ := s.CreateTransaction(req)
	fmt.Println("BaCA INI !!! : ", snapResp.RedirectURL)
	// if err != nil {
	// 	return "", err
	// }
	return snapResp.RedirectURL, nil
}
