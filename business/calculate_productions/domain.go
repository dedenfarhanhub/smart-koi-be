package calculate_productions

import (
	"context"
	"github.com/dedenfarhanhub/smart-koi-be/business/history_productions"
	"time"
)

type Domain struct {
	ID        		int
	Production  	int64
	Stock     		int64
	MarketDemand    int64
	UserId      	int
	PercentageCalculation PercentageCalculationDomain
	LatestHistoryProduction history_productions.Domain
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
}

type PercentageCalculationDomain struct {
	Type 		string
	Percentage 	float64
}

type UseCase interface {
	Store(ctx context.Context, data *Domain) (Domain, error)
}

type Repository interface {
	Store(ctx context.Context, data *Domain) (Domain, error)
}