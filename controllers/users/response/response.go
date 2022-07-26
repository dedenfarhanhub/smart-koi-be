package response

import (
	"github.com/dedenfarhanhub/smart-koi-be/business/users"
)

type Users struct {
	ID               int                `gorm:"primary_key" json:"id"`
	Name             string             `json:"name"`
	Username         string             `json:"username"`
	PhoneNumber      string             `json:"phone_number"`
	Role             string             `json:"role"`
}

func FromDomain(userDomain users.Domain) Users {
	return Users{
		ID:              	 userDomain.ID,
		Name:             	userDomain.Name,
		Username:         	userDomain.Username,
		PhoneNumber:		userDomain.PhoneNumber,
		Role:             	userDomain.Role,
	}
}

func FromListDomain(usersDomain []users.Domain) *[]Users {
	var allUsers []Users
	for _, value := range usersDomain {
		user := Users{
			ID:       		value.ID,
			Name:           value.Name,
			Username:       value.Username,
			PhoneNumber:    value.PhoneNumber,
			Role:           value.Role,
		}
		allUsers = append(allUsers, user)
	}
	return &allUsers
}