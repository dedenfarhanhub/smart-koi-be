package request

import "github.com/dedenfarhanhub/smart-koi-be/business/users"

type UpdateProfile struct {
	Name            string `json:"name"`
	Password  		string `json:"password"`
	PhoneNumber		string `json:"phone_number"`
}


func (req *UpdateProfile) ToUpdateDomain() *users.Domain {
	return &users.Domain{
		Name:		req.Name,
		Password:	req.Password,
		PhoneNumber: req.PhoneNumber,
	}
}
