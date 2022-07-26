package request

import "github.com/dedenfarhanhub/smart-koi-be/business/users"

type Users struct {
	Name     	string `json:"name"`
	Password 	string `json:"password"`
	Username    string `json:"username"`
	PhoneNumber	string `json:"phone_number"`
	Role 		string `json:"role"`
}

type UpdateUsers struct {
	Id			int 	`json:"id"`
	Name        string 	`json:"name"`
	Password	string 	`json:"password"`
	PhoneNumber string  `json:"phone_number"`
	Role  		string 	`json:"role"`
}

type FindUserByIdRequest struct {
	Id int `param:"id"`
}

func (req *Users) ToDomain() *users.Domain {
	return &users.Domain{
		Name:     		req.Name,
		Password: 		req.Password,
		Username:		req.Username,
		PhoneNumber:	req.PhoneNumber,
		Role: 			req.Role,
	}
}

func (req *UpdateUsers) ToUpdateDomain() *users.Domain {
	return &users.Domain{
		ID: 		 req.Id,
		Name:		 req.Name,
		PhoneNumber: req.PhoneNumber,
		Role:		 req.Role,
		Password: 	 req.Password,
	}
}
