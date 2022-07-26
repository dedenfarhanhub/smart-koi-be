package users

import (
	"github.com/dedenfarhanhub/smart-koi-be/business/users"
	"time"
)

type Users struct {
	ID        	int `gorm:"primary_key" json:"id"`
	Name      	string
	Password  	string
	Username  	string
	PhoneNumber	string
	Role      	string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}

func (rec *Users) ToDomain() (res *users.Domain) {
	if rec != nil {
		res = &users.Domain{
			ID:        		rec.ID,
			Name:      		rec.Name,
			Password:  		rec.Password,
			Username:     	rec.Username,
			PhoneNumber:	rec.PhoneNumber,
			Role:      	 	rec.Role,
			CreatedAt:   	rec.CreatedAt,
			UpdatedAt:		rec.UpdatedAt,
		}
	}
	return res
}

func fromDomain(userDomain users.Domain) *Users {
	return &Users{
		ID:        		userDomain.ID,
		Name:      		userDomain.Name,
		Password:  		userDomain.Password,
		Username:  		userDomain.Username,
		PhoneNumber:	userDomain.PhoneNumber,
		Role:      		userDomain.Role,
		CreatedAt: 		userDomain.CreatedAt,
		UpdatedAt: 		userDomain.UpdatedAt,
	}
}