package master_category

import (
	"context"
	"time"
)

type MasterCategoryUsecase struct {
	masterCategoryRespository Repository
	contextTimeout            time.Duration
}

func NewMasterCategoryUsecase(timeout time.Duration, r Repository) Usecase {
	return &MasterCategoryUsecase{
		contextTimeout:            timeout,
		masterCategoryRespository: r,
	}
}

func (uc *MasterCategoryUsecase) Find(ctx context.Context, search string) ([]Domain, error) {
	resp, err := uc.masterCategoryRespository.Find(ctx, search)
	if err != nil {
		return []Domain{}, err
	}

	return resp, nil
}

func (uc *MasterCategoryUsecase) FindByID(ctx context.Context, id int) (Domain, error) {
	resp, err := uc.masterCategoryRespository.FindByID(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return resp, nil
}

func (uc *MasterCategoryUsecase) Store(ctx context.Context, data *Domain) (Domain, error) {
	res, err := uc.masterCategoryRespository.Store(ctx, data)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *MasterCategoryUsecase) Update(ctx context.Context, ID int, data *Domain) (res Domain, err error) {
	_, err = uc.FindByID(ctx, ID)
	if err != nil {
		return res, err
	}

	data.UpdatedAt = time.Now().UTC()

	res, err = uc.masterCategoryRespository.Update(ctx, ID, data)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *MasterCategoryUsecase) Delete(ctx context.Context, ID int) (err error) {
	_, err = uc.FindByID(ctx, ID)
	if err != nil {
		return err
	}
	err = uc.masterCategoryRespository.Delete(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}
