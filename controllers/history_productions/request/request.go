package request

import (
	"github.com/dedenfarhanhub/smart-koi-be/business/history_productions"
	"github.com/dedenfarhanhub/smart-koi-be/constans"
	"time"
)

type HistoryProduction struct {
	PeriodDate     	string 	`json:"period_date"`
	Production 		int64 	`json:"production"`
	Stock    		int64 	`json:"stock"`
	MarketDemand	int64 	`json:"market_demand"`
}

type UpdateHistoryProduction struct {
	Id				int 	`json:"id"`
	PeriodDate     	string 	`json:"period_date"`
	Production 		int64	`json:"production"`
	Stock    		int64 	`json:"stock"`
	MarketDemand	int64 	`json:"market_demand"`
}

type FindHistoryProductionByIdRequest struct {
	Id int `param:"id"`
}

func (req *HistoryProduction) ToDomain(userId int) *history_productions.Domain {
	parsePeriodDate, _ := time.Parse(constans.DateFormat, req.PeriodDate)
	return &history_productions.Domain{
		PeriodDate:     parsePeriodDate,
		Production: 	req.Production,
		Stock:    		req.Stock,
		MarketDemand:	req.MarketDemand,
		UserId: 		userId,
	}
}

func (req *UpdateHistoryProduction) ToUpdateDomain(userId int) *history_productions.Domain {
	parsePeriodDate, _ := time.Parse(constans.DateFormat, req.PeriodDate)
	return &history_productions.Domain{
		ID: 			req.Id,
		PeriodDate:     parsePeriodDate,
		Production: 	req.Production,
		Stock:    		req.Stock,
		MarketDemand:	req.MarketDemand,
		UserId: 		userId,
	}
}

