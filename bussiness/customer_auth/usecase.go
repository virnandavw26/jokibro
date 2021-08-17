package customer_auth

import (
	"context"
	"jokibro/app/middleware"
	"jokibro/bussiness/customer"
	"jokibro/helper/encrypt"
	"jokibro/helper/messages"
	"strings"
	"time"

	"gorm.io/gorm"
)

type customerAuthUsecase struct {
	customerRepository customer.Repository
	contextTimeout     time.Duration
	jwtAuth            *middleware.ConfigJWT
}

func NewCustomerAuthUsecase(timeout time.Duration, customerRepo customer.Repository, jwt *middleware.ConfigJWT) Usecase {
	return &customerAuthUsecase{
		customerRepository: customerRepo,
		jwtAuth:            jwt,
		contextTimeout:     timeout,
	}
}

func (uc customerAuthUsecase) Register(ctx context.Context, data *customer.Domain) (res Domain, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	customer, err := uc.customerRepository.FindByEmail(ctx, data.Email)
	if err != nil {
		if err.Error() != gorm.ErrRecordNotFound.Error() {
			return res, err
		}
	}

	if customer.ID != 0 {
		return res, messages.ErrDataAlreadyExist
	}

	data.Password, err = encrypt.Hash(data.Password)
	if err != nil {
		return res, err
	}

	customer, err = uc.customerRepository.Store(ctx, data)
	if err != nil {
		return res, err
	}

	res.Token = uc.jwtAuth.GenerateToken(customer.ID, "customer")

	return res, err
}

func (uc customerAuthUsecase) Login(ctx context.Context, data *Domain) (res Domain, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if strings.TrimSpace(data.Email) == "" && strings.TrimSpace(data.Password) == "" {
		return res, messages.ErrUsernamePasswordNotFound
	}

	customer, err := uc.customerRepository.FindByEmail(ctx, data.Email)
	if err != nil {
		return res, err
	}

	if !encrypt.ValidateHash(data.Password, customer.Password) {
		return res, messages.ErrInvalidCred
	}

	res.Token = uc.jwtAuth.GenerateToken(customer.ID, "customer")

	return res, err
}
