package history_productions

import (
	"github.com/dedenfarhanhub/smart-koi-be/business/history_productions"
	"time"
)

type HistoryProductions struct {
	ID        		int 		`gorm:"primary_key" json:"id"`
	PeriodDate  	time.Time 	`gorm:"column:period_date"`
	Production  	int64     	`gorm:"column:production"`
	Stock   		int64     	`gorm:"column:stock"`
	MarketDemand	int64     	`gorm:"column:market_demand"`
	UserId			int     	`gorm:"column:user_id"`
	CreatedAt 		time.Time 	`gorm:"column:created_at"`
	UpdatedAt 		time.Time 	`gorm:"column:updated_at"`
}

func (rec *HistoryProductions) ToDomain() (res *history_productions.Domain) {
	if rec != nil {
		res = &history_productions.Domain{
			ID:        		rec.ID,
			PeriodDate:     rec.PeriodDate,
			Production:  	rec.Production,
			Stock:     		rec.Stock,
			MarketDemand:	rec.MarketDemand,
			UserId: 		rec.UserId,
			CreatedAt: 		rec.CreatedAt,
			UpdatedAt: 		rec.UpdatedAt,
		}
	}
	return res
}

func fromDomain(historyProductionDomain history_productions.Domain) *HistoryProductions {
	return &HistoryProductions{
		ID:        		historyProductionDomain.ID,
		PeriodDate:     historyProductionDomain.PeriodDate,
		Production:  	historyProductionDomain.Production,
		Stock:     		historyProductionDomain.Stock,
		MarketDemand:	historyProductionDomain.MarketDemand,
		UserId: 		historyProductionDomain.UserId,
		CreatedAt: 		historyProductionDomain.CreatedAt,
		UpdatedAt: 		historyProductionDomain.UpdatedAt,
	}
}
