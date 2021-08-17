package admin_auth

import (
	"context"
)

type Domain struct {
	Email    string
	Password string
	Token    string
}

type Usecase interface {
	Login(ctx context.Context, data *Domain) (res Domain, err error)
}
