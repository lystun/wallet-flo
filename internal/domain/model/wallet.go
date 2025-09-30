package model

import (
	"wallet-flo/internal/domain/dto"
)

type Wallet struct {
	Models
	Pin            string
	UserID         string `gorm:"unique"`
	Currency       string `gorm:"unique"`
	Email          string `gorm:"unique"`
	AccountNo      string `gorm:"unique"`
	PhoneNumber    string `gorm:"unique"`
	Balance        float64
	WalletCategory string
	WalletType     string
}

func (w *Wallet) ToWalletDTO() dto.Wallet {
	out := dto.Wallet{
		Models: dto.Models{
			ID:        w.ID,
			CreatedAt: w.CreatedAt,
			UpdatedAt: w.UpdatedAt,
		},
		UserID:   w.UserID,
		Currency: w.Currency,
		Balance:  w.Balance,
	}

	return out
}

func FromWalletDTO(w dto.Wallet) Wallet {
	response := Wallet{
		Models: Models{
			ID:        w.ID,
			CreatedAt: w.CreatedAt,
			UpdatedAt: w.UpdatedAt,
		},
		UserID:   w.UserID,
		Currency: w.Currency,
		Balance:  w.Balance,
	}

	return response
}
