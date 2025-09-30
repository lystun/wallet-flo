package dto

import "wallet-flo/pkg/config"

type Wallet struct {
	Models
	Pin  int64 `json:"pin"`
	UserID   string  `json:"user_id"`
	Currency string  `json:"currency"`
	Balance  float64 `json:"balance"`
	Email  string `json:"email"`
	AccountNo  string `json:"account_no"`
	PhoneNumber  string `json:"phone_number"`
	WalletCategory  string `json:"wallet_category"` // define this as an enum
	WalletType  string `json:"wallet_type"` // define this as an enum
}

type DepositRequest struct {
	Amount   float64 `json:"amount" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
}

type TransferRequest struct {
	FromCurrency  string  `json:"from_currency" binding:"required"`
	ToCurrency    string  `json:"to_currency" binding:"required"`
	FromAmount    float64 `json:"from_amount"`
	ToAmount      float64 `json:"to_amount"`
	ReceiverEmail string  `json:"receiver_email"`
}

func IsValidCurrency(currency string) bool {
	switch currency {
	case string(config.USD), string(config.EUR), string(config.XAF), string(config.NGN):
		return true
	default:
		return false
	}
}
