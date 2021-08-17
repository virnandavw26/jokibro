package customer_auth

import (
	"context"
	"jokibro/bussiness/customer"
)

type Domain struct {
	Email    string
	Password string
	Token    string
}

type Usecase interface {
	Login(ctx context.Context, data *Domain) (res Domain, err error)
	Register(ctx context.Context, data *customer.Domain) (res Domain, err error)
}
