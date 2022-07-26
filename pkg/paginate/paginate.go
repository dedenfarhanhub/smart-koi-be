package paginate

import (
	"github.com/dedenfarhanhub/smart-koi-be/helper/pagination"
	"gorm.io/gorm"
	"math"
)

func Paginate(value interface{}, pagination *pagination.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)
	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages
	pagination.PrevPage   = pagination.Page > 1
	pagination.NextPage   = (totalPages - pagination.Page) > 1
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}

