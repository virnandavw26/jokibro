package request

import (
	"jokibro/bussiness/customer"
)

type Customer struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

func (req *Customer) ToDomain() *customer.Domain {
	return &customer.Domain{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Address:  req.Address,
	}
}
