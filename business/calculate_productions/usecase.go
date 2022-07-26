package calculate_productions

import (
	"context"
	"github.com/dedenfarhanhub/smart-koi-be/business"
	"github.com/dedenfarhanhub/smart-koi-be/business/history_productions"
	"github.com/dedenfarhanhub/smart-koi-be/helper/calculation"
	"github.com/dedenfarhanhub/smart-koi-be/helper/logging"
	"math"
	"time"
)

type CalculateProductionUseCase struct {
	calculateProductionRepository 	Repository
	historyProductionRepository 	history_productions.Repository
	contextTimeout 					time.Duration
	logger         					logging.Logger
}

func (c CalculateProductionUseCase) Store(ctx context.Context, data *Domain) (Domain, error) {
	request := map[string]interface{}{
		"stock": 			data.Stock,
		"market_demand":	data.MarketDemand,
		"user_id": 			data.UserId,
	}

	allHistoryProductions, err := c.historyProductionRepository.GetAll(ctx)
	if err != nil {
		return Domain{}, business.ErrInternalServer
	}

	if len(allHistoryProductions) == 0 {
		return Domain{}, business.ErrHistoryProductionResource
	}

	calculationFuzzy(data, allHistoryProductions)
	if data.Production <= 0 {
		return Domain{}, business.ErrInternalServer
	}

	savedHistoryProduction, err := c.calculateProductionRepository.Store(ctx, data)
	latestPeriodDate, _ := c.historyProductionRepository.GetLatestPeriodDate(ctx)
	calculationPercentage(data, latestPeriodDate)
	savedHistoryProduction.LatestHistoryProduction = latestPeriodDate
	savedHistoryProduction.PercentageCalculation = data.PercentageCalculation
	if err != nil {
		result := map[string]interface{}{
			"error": err.Error(),
		}
		c.logger.LogEntry(request, result).Error(err.Error())
		return Domain{}, err
	}

	result := map[string]interface{}{
		"success": "true",
	}
	c.logger.LogEntry(request, result).Info("incoming request")
	return savedHistoryProduction, nil
}

func calculationFuzzy(data *Domain, allHistoryProductions []history_productions.Domain){
	var historyValueMarketDemands []int64
	var historyValueStocks []int64
	var historyValueProductions []int64

	for _, value := range allHistoryProductions {
		historyValueMarketDemands = append(historyValueMarketDemands, value.MarketDemand)
		historyValueStocks = append(historyValueStocks, value.Stock)
		historyValueProductions = append(historyValueProductions, value.Production)
	}

	minMarketDemand, maxMarketDemand := calculation.FindMinAndMax(historyValueMarketDemands)

	var marketDemandDown float64

	if data.MarketDemand <= minMarketDemand {
		marketDemandDown = 1
	} else if minMarketDemand < data.MarketDemand && data.MarketDemand < maxMarketDemand {
		marketDemandDown = calculation.RoundFloat(float64(maxMarketDemand - data.MarketDemand) / float64(maxMarketDemand - minMarketDemand), 2)
	} else {
		marketDemandDown = 0
	}

	var marketDemandUp float64

	if data.MarketDemand <= minMarketDemand {
		marketDemandUp = 0
	} else if minMarketDemand < data.MarketDemand && data.MarketDemand < maxMarketDemand {
		marketDemandUp = calculation.RoundFloat(float64(data.MarketDemand - minMarketDemand) / float64(maxMarketDemand - minMarketDemand), 2)
	} else {
		marketDemandUp = 1
	}

	minStock, maxStock := calculation.FindMinAndMax(historyValueStocks)

	var stockLittle float64

	if data.Stock <= minStock {
		stockLittle = 1
	} else if minStock < data.Stock && data.Stock < maxStock {
		stockLittle = calculation.RoundFloat(float64(maxStock - data.Stock) / float64(maxStock - minStock), 2)
	} else {
		stockLittle = 0
	}

	var stockMany float64

	if data.Stock <= minStock {
		stockMany = 0
	} else if minStock < data.Stock && data.Stock < maxStock {
		stockMany = calculation.RoundFloat(float64(data.Stock - minStock) / float64(maxStock - minStock), 2)
	} else {
		stockMany = 1
	}

	minProduction, maxProduction := calculation.FindMinAndMax(historyValueProductions)

	alphaPredicateOne := calculation.FindMinFloat(marketDemandDown, stockMany)
	productionOne := productionReduce(alphaPredicateOne, float64(minProduction), float64(maxProduction))

	alphaPredicateTwo := calculation.FindMinFloat(marketDemandDown, stockLittle)
	productionTwo := productionReduce(alphaPredicateTwo, float64(minProduction), float64(maxProduction))

	alphaPredicateThree := calculation.FindMinFloat(marketDemandUp, stockMany)
	productionThree := productionIncrease(alphaPredicateThree, float64(minProduction), float64(maxProduction))

	alphaPredicateFour := calculation.FindMinFloat(marketDemandUp, stockLittle)
	productionFour := productionIncrease(alphaPredicateFour, float64(minProduction), float64(maxProduction))

	production := ((alphaPredicateOne * productionOne) + (alphaPredicateTwo * productionTwo) + (alphaPredicateThree * productionThree) + (alphaPredicateFour * productionFour)) / (alphaPredicateOne + alphaPredicateTwo + alphaPredicateThree + alphaPredicateFour)
	data.Production = int64(math.Round(production))
}

func calculationPercentage(data *Domain, latestPeriodDate history_productions.Domain){
	percentage := calculation.RoundFloat(float64(data.Production - latestPeriodDate.Production) / float64(latestPeriodDate.Production) * 100, 2)
	data.PercentageCalculation.Percentage = percentage
	if percentage > 0 {
		data.PercentageCalculation.Type = "UP"
	} else {
		data.PercentageCalculation.Type = "DOWN"
	}
}

func productionReduce(alphaPredicate float64, minProduction float64, maxProduction float64) float64 {
	if alphaPredicate == 1 {
		return minProduction
	}

	if alphaPredicate == 0 {
		return maxProduction
	}

	return calculation.RoundFloat(maxProduction - (alphaPredicate * (maxProduction - minProduction)), 2)
}

func productionIncrease(alphaPredicate float64, minProduction float64, maxProduction float64) float64 {
	if alphaPredicate == 1 {
		return maxProduction
	}

	if alphaPredicate == 0 {
		return minProduction
	}

	return calculation.RoundFloat(((maxProduction - minProduction) * alphaPredicate) + minProduction, 2)
}

func NewCalculateProductionUseCase(cpr Repository, hpr history_productions.Repository, timeout time.Duration, logger logging.Logger) UseCase {
	return &CalculateProductionUseCase{
		calculateProductionRepository: cpr,
		historyProductionRepository: hpr,
		contextTimeout: timeout,
		logger:         logger,
	}
}