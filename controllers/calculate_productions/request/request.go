package request

import (
	"github.com/dedenfarhanhub/smart-koi-be/business/calculate_productions"
)

type CalculateProduction struct {
	Stock    		int64 	`json:"stock"`
	MarketDemand	int64 	`json:"market_demand"`
}

func (req *CalculateProduction) ToDomain(userId int) *calculate_productions.Domain {
	return &calculate_productions.Domain{
		Stock:    		req.Stock,
		MarketDemand:	req.MarketDemand,
		UserId: 		userId,
	}
}