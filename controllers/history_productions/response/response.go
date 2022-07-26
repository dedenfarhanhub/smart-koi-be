package response

import (
	"github.com/dedenfarhanhub/smart-koi-be/business/history_productions"
	"time"
)

type HistoryProductions struct {
	ID              int         `json:"id"`
	PeriodDate     	time.Time 	`json:"period_date"`
	Production 		int64		`json:"production"`
	Stock    		int64 		`json:"stock"`
	MarketDemand	int64 		`json:"market_demand"`
	UserId			int 		`json:"user_id"`
}

type HistoryProductionStats struct {
	MaxProduction 	int64 `json:"max_production"`
	MinProduction 	int64 `json:"min_production"`
	MaxStock		int64 `json:"max_stock"`
	MinStock		int64 `json:"min_stock"`
	MaxMarketDemand int64 `json:"max_market_demand"`
	MinMarketDemand int64 `json:"min_market_demand"`
}

func FromDomain(historyProduction history_productions.Domain) HistoryProductions {
	return HistoryProductions{
		ID:    			historyProduction.ID,
		PeriodDate:		historyProduction.PeriodDate,
		Production:		historyProduction.Production,
		Stock: 			historyProduction.Stock,
		MarketDemand:	historyProduction.MarketDemand,
		UserId: 		historyProduction.UserId,
	}
}

func FromListDomain(historyProductionsDomain []history_productions.Domain) *[]HistoryProductions {
	var historyProductions []HistoryProductions
	for _, value := range historyProductionsDomain {
		historyProduction := HistoryProductions{
			ID:       		value.ID,
			PeriodDate:		value.PeriodDate,
			Production:		value.Production,
			Stock: 			value.Stock,
			MarketDemand:	value.MarketDemand,
			UserId: 		value.UserId,
		}
		historyProductions = append(historyProductions, historyProduction)
	}
	return &historyProductions
}

func FromHistoryProductionStatDomain(historyProductionStatDomain history_productions.HistoryProductionStatDomain) HistoryProductionStats {
	return HistoryProductionStats{
		MaxProduction: historyProductionStatDomain.MaxProduction,
		MinProduction: historyProductionStatDomain.MinProduction,
		MaxStock: historyProductionStatDomain.MaxStock,
		MinStock: historyProductionStatDomain.MinStock,
		MaxMarketDemand: historyProductionStatDomain.MaxMarketDemand,
		MinMarketDemand: historyProductionStatDomain.MinMarketDemand,
	}
}
