package calculate_productions

import (
	"github.com/dedenfarhanhub/smart-koi-be/business/calculate_productions"
	"time"
)

type CalculateProductions struct {
	ID        		int 		`gorm:"primary_key" json:"id"`
	Production  	int64     	`gorm:"column:production"`
	Stock   		int64     	`gorm:"column:stock"`
	MarketDemand	int64     	`gorm:"column:market_demand"`
	UserId			int     	`gorm:"column:user_id"`
	CreatedAt 		time.Time 	`gorm:"column:created_at"`
	UpdatedAt 		time.Time 	`gorm:"column:updated_at"`
}

func (rec *CalculateProductions) ToDomain() (res *calculate_productions.Domain) {
	if rec != nil {
		res = &calculate_productions.Domain{
			ID:        		rec.ID,
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

func fromDomain(calculateProductionDomain calculate_productions.Domain) *CalculateProductions {
	return &CalculateProductions{
		ID:        		calculateProductionDomain.ID,
		Production:  	calculateProductionDomain.Production,
		Stock:     		calculateProductionDomain.Stock,
		MarketDemand:	calculateProductionDomain.MarketDemand,
		UserId: 		calculateProductionDomain.UserId,
		CreatedAt: 		calculateProductionDomain.CreatedAt,
		UpdatedAt: 		calculateProductionDomain.UpdatedAt,
	}
}
