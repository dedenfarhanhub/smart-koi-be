package response

import "github.com/dedenfarhanhub/smart-koi-be/business/users"

type Profile struct {
	ID               int                `json:"id"`
	Name             string             `json:"name"`
	Username         string             `json:"username"`
	PhoneNumber      string             `json:"phone_number"`
	Role             string             `json:"role"`
}

func FromDomain(userDomain users.Domain) Profile {
	return Profile{
		ID:               userDomain.ID,
		Name:             userDomain.Name,
		Username:         userDomain.Username,
		PhoneNumber: 	  userDomain.PhoneNumber,
		Role:             userDomain.Role,
	}
}
