package master_service

import (
	"context"
	"time"
)

type MasterServiceUsecase struct {
	masterServiceRespository Repository
	contextTimeout           time.Duration
}

func NewMasterServiceUsecase(timeout time.Duration, r Repository) Usecase {
	return &MasterServiceUsecase{
		contextTimeout:           timeout,
		masterServiceRespository: r,
	}
}

func (uc *MasterServiceUsecase) Find(ctx context.Context, search string, masterCategoryID int) ([]Domain, error) {
	resp, err := uc.masterServiceRespository.Find(ctx, search, masterCategoryID)
	if err != nil {
		return []Domain{}, err
	}

	return resp, nil
}

func (uc *MasterServiceUsecase) FindByID(ctx context.Context, id int) (Domain, error) {
	resp, err := uc.masterServiceRespository.FindByID(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return resp, nil
}

func (uc *MasterServiceUsecase) Store(ctx context.Context, data *Domain) (Domain, error) {
	res, err := uc.masterServiceRespository.Store(ctx, data)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *MasterServiceUsecase) Update(ctx context.Context, ID int, data *Domain) (res Domain, err error) {
	_, err = uc.FindByID(ctx, ID)
	if err != nil {
		return res, err
	}

	data.UpdatedAt = time.Now().UTC()

	res, err = uc.masterServiceRespository.Update(ctx, ID, data)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *MasterServiceUsecase) Delete(ctx context.Context, ID int) (err error) {
	_, err = uc.FindByID(ctx, ID)
	if err != nil {
		return err
	}
	err = uc.masterServiceRespository.Delete(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}
