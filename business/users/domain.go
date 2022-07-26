package users

import (
	"context"
	"github.com/dedenfarhanhub/smart-koi-be/helper/pagination"
	"time"
)

type Domain struct {
	ID        	int
	Name      	string
	Password  	string
	Username  	string
	PhoneNumber string
	Role      	string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}

type UseCase interface {
	Login(ctx context.Context, username, password string, sso bool) (string, Domain, error)
	Store(ctx context.Context, data *Domain, sso bool) (Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, data *Domain, id int) (Domain, error)
	Destroy(ctx context.Context, id int) (Domain, error)
	Fetch(ctx context.Context, pagination pagination.Pagination, keyword string) (pagination.Pagination, []Domain, error)
}

type Repository interface {
	Fetch(ctx context.Context, pagination pagination.Pagination, keyword string) (pagination.Pagination, []Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, data *Domain, id int) (Domain, error)
	Store(ctx context.Context, data *Domain) (Domain, error)
	Destroy(ctx context.Context, id int) (Domain, error)
	GetByUsername(ctx context.Context, username string) (Domain, error)
}
