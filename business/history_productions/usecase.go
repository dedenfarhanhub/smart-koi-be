package history_productions

import (
	"context"
	"github.com/dedenfarhanhub/smart-koi-be/business"
	"github.com/dedenfarhanhub/smart-koi-be/helper/calculation"
	"github.com/dedenfarhanhub/smart-koi-be/helper/logging"
	"github.com/dedenfarhanhub/smart-koi-be/helper/pagination"
	"strings"
	"time"
)

type HistoryProductionUseCase struct {
	historyProductionRepository Repository
	contextTimeout 				time.Duration
	logger         				logging.Logger
}

func (h HistoryProductionUseCase) Stat(ctx context.Context) (HistoryProductionStatDomain, error) {
	allHistoryProductions, err := h.historyProductionRepository.GetAll(ctx)

	var historyValueMarketDemands []int64
	var historyValueStocks []int64
	var historyValueProductions []int64

	for _, value := range allHistoryProductions {
		historyValueMarketDemands = append(historyValueMarketDemands, value.MarketDemand)
		historyValueStocks = append(historyValueStocks, value.Stock)
		historyValueProductions = append(historyValueProductions, value.Production)
	}

	minMarketDemand, maxMarketDemand := calculation.FindMinAndMax(historyValueMarketDemands)
	minStock, maxStock := calculation.FindMinAndMax(historyValueStocks)
	minProduction, maxProduction := calculation.FindMinAndMax(historyValueProductions)

	var historyProductionStatDomain = HistoryProductionStatDomain{}

	historyProductionStatDomain.MaxProduction 	= maxProduction
	historyProductionStatDomain.MinProduction 	= minProduction
	historyProductionStatDomain.MaxStock 		= maxStock
	historyProductionStatDomain.MinStock 		= minStock
	historyProductionStatDomain.MaxMarketDemand = maxMarketDemand
	historyProductionStatDomain.MinMarketDemand = minMarketDemand

	if err != nil {
		return HistoryProductionStatDomain{}, business.ErrInternalServer
	}

	return historyProductionStatDomain, nil
}

func (h HistoryProductionUseCase) Barchart(ctx context.Context) ([]Domain, error) {
	req := pagination.Pagination{}
	req.Page = 1
	req.Limit = 3
	_, allHistoryProductions, err := h.historyProductionRepository.Fetch(ctx, req, "", "")
	if err != nil {
		return []Domain{}, business.ErrInternalServer
	}

	return allHistoryProductions, nil
}

func (h HistoryProductionUseCase) Store(ctx context.Context, data *Domain) (Domain, error) {
	request := map[string]interface{}{
		"period_date": 		data.PeriodDate,
		"production": 		data.Production,
		"stock": 			data.Stock,
		"market_demand":	data.MarketDemand,
		"user_id": 			data.UserId,
	}

	existedHistoryProduction, err := h.historyProductionRepository.GetByPeriodDate(ctx, data.PeriodDate)

	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			result := map[string]interface{}{
				"success": "false",
				"error":   err.Error(),
			}
			h.logger.LogEntry(request, result).Error(err.Error())
			return Domain{}, err
		}
	}

	if existedHistoryProduction != (Domain{}) {
		return Domain{}, business.ErrDuplicateData
	}

	savedHistoryProduction, err := h.historyProductionRepository.Store(ctx, data)
	if err != nil {
		result := map[string]interface{}{
			"error": err.Error(),
		}
		h.logger.LogEntry(request, result).Error(err.Error())
		return Domain{}, err
	}

	result := map[string]interface{}{
		"success": "true",
	}
	h.logger.LogEntry(request, result).Info("incoming request")

	return savedHistoryProduction, nil
}

func (h HistoryProductionUseCase) GetByID(ctx context.Context, id int) (Domain, error) {
	result, err := h.historyProductionRepository.GetByID(ctx, id)
	if err != nil {
		return Domain{}, business.ErrNotFound
	}

	return result, nil
}

func (h HistoryProductionUseCase) Update(ctx context.Context, data *Domain, id int) (Domain, error) {
	request := map[string]interface{}{
		"id":       		id,
		"period_date": 		data.PeriodDate,
		"production": 		data.Production,
		"stock": 			data.Stock,
		"market_demand":	data.MarketDemand,
		"user_id": 			data.UserId,
	}

	_, err := h.historyProductionRepository.GetByID(ctx, id)

	if err != nil {
		return Domain{}, business.ErrNotFound
	}

	existedHistoryProduction, err := h.historyProductionRepository.GetByPeriodDateNotInId(ctx, data.PeriodDate, id)

	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			result := map[string]interface{}{
				"success": "false",
				"error":   err.Error(),
			}
			h.logger.LogEntry(request, result).Error(err.Error())
			return Domain{}, err
		}
	}

	if existedHistoryProduction != (Domain{}) {
		return Domain{}, business.ErrDuplicateData
	}

	savedHistoryProduction, err := h.historyProductionRepository.Update(ctx, data, id)
	if err != nil {
		result := map[string]interface{}{
			"error": err.Error(),
		}
		h.logger.LogEntry(request, result).Error(err.Error())
		return Domain{}, err
	}

	result := map[string]interface{}{
		"success": "true",
	}
	h.logger.LogEntry(request, result).Info("incoming request")

	return savedHistoryProduction, nil
}

func (h HistoryProductionUseCase) Destroy(ctx context.Context, id int) (Domain, error) {
	result, err := h.historyProductionRepository.Destroy(ctx, id)
	if err != nil {
		return Domain{}, business.ErrNotFound
	}

	return result, nil
}

func (h HistoryProductionUseCase) Fetch(ctx context.Context, pagination pagination.Pagination, startDate string, endDate string) (pagination.Pagination, []Domain, error) {
	result, allHistoryProductions, err := h.historyProductionRepository.Fetch(ctx, pagination, startDate, endDate)
	if err != nil {
		result := map[string]interface{}{
			"error": err.Error(),
		}
		h.logger.LogEntry("can't get all data history_productions", result).Error(err.Error())

		return pagination, []Domain{}, err
	}

	h.logger.LogEntry("success to get all data history_productions", nil).Info("incoming request")

	return result, allHistoryProductions, nil
}

func NewHistoryProductionUseCase(hpr Repository, timeout time.Duration, logger logging.Logger) UseCase {
	return &HistoryProductionUseCase{
		historyProductionRepository: hpr,
		contextTimeout: timeout,
		logger:         logger,
	}
}