package admin

import (
	"context"
	"time"
)

type adminUsecase struct {
	adminRespository Repository
	contextTimeout   time.Duration
}

func NewAdminUsecase(timeout time.Duration, r Repository) Usecase {
	return &adminUsecase{
		contextTimeout:   timeout,
		adminRespository: r,
	}
}

func (uc *adminUsecase) FindByID(ctx context.Context, id int) (Domain, error) {
	resp, err := uc.adminRespository.FindByID(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return resp, nil
}

func (uc *adminUsecase) FindByEmail(ctx context.Context, email string) (Domain, error) {
	resp, err := uc.adminRespository.FindByEmail(ctx, email)
	if err != nil {
		return Domain{}, err
	}

	return resp, nil
}
