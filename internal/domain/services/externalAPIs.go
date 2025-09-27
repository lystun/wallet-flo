package services

import (
	"context"

	"wallet-flo/internal/domain/dto"
)

type ExternalAPIs interface {
	GetPairExchangeRate(ctx context.Context, baseCurrency, targetCurrency string) (dto.ExchangeRate, error)
	GetUSDExchangeRate(ctx context.Context) (dto.ExchangeRate, error)
	GetUserCountry(ctx context.Context, ipAddress string) (string, error)
}
