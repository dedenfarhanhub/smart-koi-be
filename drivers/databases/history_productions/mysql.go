package history_productions

import (
	"context"
	"github.com/dedenfarhanhub/smart-koi-be/business/history_productions"
	"github.com/dedenfarhanhub/smart-koi-be/constans"
	"github.com/dedenfarhanhub/smart-koi-be/helper/pagination"
	"github.com/dedenfarhanhub/smart-koi-be/pkg/paginate"
	"gorm.io/gorm"
	"time"
)

type mysqlHistoryProductionsRepository struct {
	Conn *gorm.DB
}

func (m mysqlHistoryProductionsRepository) GetLatestPeriodDate(ctx context.Context) (history_productions.Domain, error) {
	historyProduction := HistoryProductions{}
	result := m.Conn.Order("period_date desc").Last(&historyProduction)
	if result.Error != nil {
		return history_productions.Domain{}, result.Error
	}
	return *historyProduction.ToDomain(), nil
}

func (m mysqlHistoryProductionsRepository) GetAll(ctx context.Context) ([]history_productions.Domain, error) {
	var historyProductions []*HistoryProductions

	m.Conn.Find(&historyProductions)
	var allHistoryProductions []history_productions.Domain
	for _, value := range historyProductions {
		allHistoryProductions = append(allHistoryProductions, *value.ToDomain())
	}

	return allHistoryProductions, nil
}

func (m mysqlHistoryProductionsRepository) Fetch(ctx context.Context, pagination pagination.Pagination, startDate string, endDate string) (pagination.Pagination, []history_productions.Domain, error) {
	var historyProductions []*HistoryProductions

	collection := m.Conn.Where("1 = ?", 1)

	if startDate != "" {
		parsingStartDate, _ := time.Parse(constans.DateFormat, startDate)
		collection.Where("period_date >= DATE(?)", parsingStartDate)
	}

	if endDate != "" {
		parsingEndDate, _ := time.Parse(constans.DateFormat, endDate)
		collection.Where("period_date <= DATE(?)", parsingEndDate)
	}

	collection.Scopes(paginate.Paginate(historyProductions, &pagination, collection)).Find(&historyProductions)

	var allHistoryProductions []history_productions.Domain
	for _, value := range historyProductions {
		allHistoryProductions = append(allHistoryProductions, *value.ToDomain())
	}
	return pagination, allHistoryProductions, nil
}

func (m mysqlHistoryProductionsRepository) GetByID(ctx context.Context, id int) (history_productions.Domain, error) {
	historyProduction := HistoryProductions{}
	result := m.Conn.Where("id = ?", id).First(&historyProduction)
	if result.Error != nil {
		return history_productions.Domain{}, result.Error
	}
	return *historyProduction.ToDomain(), nil
}

func (m mysqlHistoryProductionsRepository) GetByPeriodDate(ctx context.Context, periodDate time.Time) (history_productions.Domain, error) {
	historyProduction := HistoryProductions{}
	result := m.Conn.Where("period_date = DATE(?)", periodDate).First(&historyProduction)
	if result.Error != nil {
		return history_productions.Domain{}, result.Error
	}
	return *historyProduction.ToDomain(), nil
}

func (m mysqlHistoryProductionsRepository) GetByPeriodDateNotInId(ctx context.Context, periodDate time.Time, id int) (history_productions.Domain, error) {
	historyProduction := HistoryProductions{}
	result := m.Conn.Where("period_date = DATE(?) and id not in(?)", periodDate, id).First(&historyProduction)
	if result.Error != nil {
		return history_productions.Domain{}, result.Error
	}
	return *historyProduction.ToDomain(), nil
}

func (m mysqlHistoryProductionsRepository) Update(ctx context.Context, data *history_productions.Domain, id int) (history_productions.Domain, error) {
	historyProductionUpdate := fromDomain(*data)
	historyProductionUpdate.ID = id
	historyProductionUpdate.UpdatedAt = time.Now().UTC()

	result := m.Conn.Where("id = ?", id).Updates(&historyProductionUpdate)
	if result.Error != nil {
		return history_productions.Domain{}, result.Error
	}

	return *historyProductionUpdate.ToDomain(), nil
}

func (m mysqlHistoryProductionsRepository) Store(ctx context.Context, data *history_productions.Domain) (history_productions.Domain, error) {
	rec := fromDomain(*data)

	result := m.Conn.Create(rec)
	if result.Error != nil {
		return history_productions.Domain{}, result.Error
	}

	return *rec.ToDomain(), nil
}

func (m mysqlHistoryProductionsRepository) Destroy(ctx context.Context, id int) (history_productions.Domain, error) {
	historyProduction := HistoryProductions{}
	result := m.Conn.Where("id = ?", id).First(&historyProduction)
	if result.Error != nil {
		return history_productions.Domain{}, result.Error
	}

	m.Conn.Delete(&historyProduction)
	return *historyProduction.ToDomain(), nil
}

func NewMySQLHistoryProductionsRepository(conn *gorm.DB) history_productions.Repository {
	return &mysqlHistoryProductionsRepository{
		Conn: conn,
	}
}