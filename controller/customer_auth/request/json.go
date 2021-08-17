package request

import (
	"jokibro/bussiness/customer_auth"
)

type CustomerAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func (req *CustomerAuth) ToDomain() *customer_auth.Domain {
	return &customer_auth.Domain{
		Email:    req.Email,
		Password: req.Password,
	}
}
