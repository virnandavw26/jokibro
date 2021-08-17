package request

import (
	"jokibro/bussiness/admin_auth"
)

type AdminAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func (req *AdminAuth) ToDomain() *admin_auth.Domain {
	return &admin_auth.Domain{
		Email:    req.Email,
		Password: req.Password,
	}
}
