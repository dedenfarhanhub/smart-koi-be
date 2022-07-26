package response

import (
	"github.com/dedenfarhanhub/smart-koi-be/business/calculate_productions"
	"github.com/dedenfarhanhub/smart-koi-be/controllers/history_productions/response"
)

type CalculateProductions struct {
	ID              int         `json:"id"`
	Production 		int64		`json:"production"`
	Stock    		int64 		`json:"stock"`
	MarketDemand	int64 		`json:"market_demand"`
	UserId			int 		`json:"user_id"`
	PercentageCalculation PercentageCalculation `json:"percentage_calculation"`
	LatestHistoryProduction response.HistoryProductions `json:"latest_history_production"`
}

type PercentageCalculation struct {
	Type		string 	`json:"type"`
	Percentage 	float64 `json:"percentage"`
}

func fromPercentageCalculationDomain(percentageCalculation calculate_productions.PercentageCalculationDomain) PercentageCalculation {
	return PercentageCalculation{
		Type: percentageCalculation.Type,
		Percentage: percentageCalculation.Percentage,
	}
}

func FromDomain(calculateProduction calculate_productions.Domain) CalculateProductions {
	return CalculateProductions{
		ID:                    calculateProduction.ID,
		Production:            calculateProduction.Production,
		Stock:                 calculateProduction.Stock,
		MarketDemand:          calculateProduction.MarketDemand,
		UserId:                calculateProduction.UserId,
		PercentageCalculation: fromPercentageCalculationDomain(calculateProduction.PercentageCalculation),
		LatestHistoryProduction: response.FromDomain(calculateProduction.LatestHistoryProduction),
	}
}