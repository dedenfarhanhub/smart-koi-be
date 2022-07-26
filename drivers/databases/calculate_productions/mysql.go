package calculate_productions

import (
	"context"
	"github.com/dedenfarhanhub/smart-koi-be/business/calculate_productions"
	"gorm.io/gorm"
)

type mysqlCalculateProductionsRepository struct {
	Conn *gorm.DB
}

func (m mysqlCalculateProductionsRepository) Store(ctx context.Context, data *calculate_productions.Domain) (calculate_productions.Domain, error) {
	rec := fromDomain(*data)

	result := m.Conn.Create(rec)
	if result.Error != nil {
		return calculate_productions.Domain{}, result.Error
	}

	return *rec.ToDomain(), nil
}

func NewMySQLCalculateProductionsRepository(conn *gorm.DB) calculate_productions.Repository {
	return &mysqlCalculateProductionsRepository{
		Conn: conn,
	}
}