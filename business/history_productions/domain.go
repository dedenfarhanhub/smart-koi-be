package history_productions

import (
	"context"
	"github.com/dedenfarhanhub/smart-koi-be/helper/pagination"
	"time"
)

type Domain struct {
	ID        		int
	PeriodDate      time.Time
	Production  	int64
	Stock     		int64
	MarketDemand    int64
	UserId      	int
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
}

type HistoryProductionStatDomain struct {
	MaxProduction 	int64
	MinProduction 	int64
	MaxStock 		int64
	MinStock 		int64
	MaxMarketDemand int64
	MinMarketDemand int64
}

type UseCase interface {
	Store(ctx context.Context, data *Domain) (Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, data *Domain, id int) (Domain, error)
	Destroy(ctx context.Context, id int) (Domain, error)
	Fetch(ctx context.Context, pagination pagination.Pagination, startDate string, endDate string) (pagination.Pagination, []Domain, error)
	Barchart(ctx context.Context) ([]Domain, error)
	Stat(ctx context.Context) (HistoryProductionStatDomain, error)
}

type Repository interface {
	Fetch(ctx context.Context, pagination pagination.Pagination, startDate string, endDate string) (pagination.Pagination, []Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	GetByPeriodDate(ctx context.Context, periodDate time.Time) (Domain, error)
	GetByPeriodDateNotInId(ctx context.Context, periodDate time.Time, id int) (Domain, error)
	Update(ctx context.Context, data *Domain, id int) (Domain, error)
	Store(ctx context.Context, data *Domain) (Domain, error)
	Destroy(ctx context.Context, id int) (Domain, error)
	GetLatestPeriodDate(ctx context.Context) (Domain, error)
}