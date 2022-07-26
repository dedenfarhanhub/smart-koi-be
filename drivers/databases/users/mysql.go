package users

import (
	"context"
	"github.com/dedenfarhanhub/smart-koi-be/business/users"
	"github.com/dedenfarhanhub/smart-koi-be/helper/pagination"
	"github.com/dedenfarhanhub/smart-koi-be/pkg/paginate"
	"gorm.io/gorm"
	"time"
)

type mysqlUsersRepository struct {
	Conn *gorm.DB
}

func (m mysqlUsersRepository) Destroy(ctx context.Context, id int) (users.Domain, error) {
	user := Users{}
	result := m.Conn.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	m.Conn.Delete(&user)
	return *user.ToDomain(), nil
}

func (m mysqlUsersRepository) Fetch(ctx context.Context, pagination pagination.Pagination, keyword string) (pagination.Pagination, []users.Domain, error) {
	var listUser []*Users

	collection := m.Conn.Where("role not in(?)", "ADMIN")

	if keyword != "" {
		collection.Where("username LIKE ? OR name LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}

	collection.Scopes(paginate.Paginate(listUser, &pagination, collection)).Find(&listUser)

	var allUsers []users.Domain
	for _, value := range listUser {
		allUsers = append(allUsers, *value.ToDomain())
	}
	return pagination, allUsers, nil
}

func (m mysqlUsersRepository) GetByID(ctx context.Context, id int) (users.Domain, error) {
	userById := Users{}
	result := m.Conn.Where("id = ?", id).First(&userById)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}
	return *userById.ToDomain(), nil
}

func (m mysqlUsersRepository) Update(ctx context.Context, data *users.Domain, id int) (users.Domain, error) {
	userUpdate := fromDomain(*data)
	userUpdate.ID = id
	userUpdate.UpdatedAt = time.Now().UTC()

	result := m.Conn.Where("id = ?", id).Updates(&userUpdate)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return *userUpdate.ToDomain(), nil
}

func (m mysqlUsersRepository) Store(ctx context.Context, data *users.Domain) (users.Domain, error) {
	rec := fromDomain(*data)

	result := m.Conn.Create(rec)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return *rec.ToDomain(), nil
}

func (m mysqlUsersRepository) GetByUsername(ctx context.Context, username string) (users.Domain, error) {
	rec := Users{}
	err := m.Conn.Where("username = ?", username).First(&rec).Error
	if err != nil {
		return users.Domain{}, err
	}
	return *rec.ToDomain(), nil
}

func NewMySQLUsersRepository(conn *gorm.DB) users.Repository {
	return &mysqlUsersRepository{
		Conn: conn,
	}
}
