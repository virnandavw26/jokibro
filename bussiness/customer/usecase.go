package customer

import (
	"context"
	"jokibro/helper/encrypt"
	"time"
)

type CustomerUsecase struct {
	customerRespository Repository
	contextTimeout      time.Duration
}

func NewCustomerUsecase(timeout time.Duration, r Repository) Usecase {
	return &CustomerUsecase{
		contextTimeout:      timeout,
		customerRespository: r,
	}
}

func (uc *CustomerUsecase) Find(ctx context.Context) ([]Domain, error) {
	resp, err := uc.customerRespository.Find(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return resp, nil
}

func (uc *CustomerUsecase) FindByID(ctx context.Context, id int) (Domain, error) {
	resp, err := uc.customerRespository.FindByID(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return resp, nil
}

func (uc *CustomerUsecase) FindByEmail(ctx context.Context, email string) (Domain, error) {
	resp, err := uc.customerRespository.FindByEmail(ctx, email)
	if err != nil {
		return Domain{}, err
	}

	return resp, nil
}
func (uc *CustomerUsecase) Store(ctx context.Context, data *Domain) (res Domain, err error) {
	data.Password, err = encrypt.Hash(data.Password)
	if err != nil {
		return res, err
	}

	res, err = uc.customerRespository.Store(ctx, data)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *CustomerUsecase) Update(ctx context.Context, ID int, data *Domain) (res Domain, err error) {
	_, err = uc.FindByID(ctx, ID)
	if err != nil {
		return res, err
	}

	data.UpdatedAt = time.Now().UTC()

	res, err = uc.customerRespository.Update(ctx, ID, &*data)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *CustomerUsecase) Delete(ctx context.Context, ID int) (err error) {
	_, err = uc.FindByID(ctx, ID)
	if err != nil {
		return err
	}
	err = uc.customerRespository.Delete(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}
