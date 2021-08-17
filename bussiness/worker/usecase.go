package worker

import (
	"context"
	"time"
)

type WorkerUsecase struct {
	workerRespository Repository
	contextTimeout    time.Duration
}

func NewWorkerUsecase(timeout time.Duration, r Repository) Usecase {
	return &WorkerUsecase{
		contextTimeout:    timeout,
		workerRespository: r,
	}
}

func (uc *WorkerUsecase) Find(ctx context.Context) ([]Domain, error) {
	resp, err := uc.workerRespository.Find(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return resp, nil
}

func (uc *WorkerUsecase) FindByID(ctx context.Context, id int) (Domain, error) {
	resp, err := uc.workerRespository.FindByID(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return resp, nil
}

func (uc *WorkerUsecase) Store(ctx context.Context, data *Domain) (Domain, error) {
	res, err := uc.workerRespository.Store(ctx, data)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *WorkerUsecase) Update(ctx context.Context, ID int, data *Domain) (res Domain, err error) {
	_, err = uc.FindByID(ctx, ID)
	if err != nil {
		return res, err
	}

	data.UpdatedAt = time.Now().UTC()

	res, err = uc.workerRespository.Update(ctx, ID, &*data)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *WorkerUsecase) Delete(ctx context.Context, ID int) (err error) {
	_, err = uc.FindByID(ctx, ID)
	if err != nil {
		return err
	}
	err = uc.workerRespository.Delete(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}
