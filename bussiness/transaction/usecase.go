package transaction

import (
	"context"
	"jokibro/bussiness/worker"
	"jokibro/helper/messages"
	"math"
	"time"
)

type TransactionUsecase struct {
	transactionRespository Repository
	workerRepository       worker.Repository
	contextTimeout         time.Duration
}

func NewTransactionUsecase(timeout time.Duration, r Repository, workerRepository worker.Repository) Usecase {
	return &TransactionUsecase{
		contextTimeout:         timeout,
		transactionRespository: r,
		workerRepository:       workerRepository,
	}
}

func (uc *TransactionUsecase) Find(ctx context.Context, customerID int, startWorkingAt time.Time, endWorkingAt time.Time) ([]Domain, error) {
	resp, err := uc.transactionRespository.Find(ctx, customerID, startWorkingAt, endWorkingAt)
	if err != nil {
		return []Domain{}, err
	}

	return resp, nil
}

func (uc *TransactionUsecase) FindByID(ctx context.Context, id int) (Domain, error) {
	resp, err := uc.transactionRespository.FindByID(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return resp, nil
}

func (uc *TransactionUsecase) Store(ctx context.Context, data *Domain) (res Domain, err error) {
	worker, err := uc.workerRepository.FindByID(ctx, data.WorkerID)
	if err != nil {
		return res, err
	}
	data.MasterServiceID = worker.MasterServiceID

	transactions, err := uc.transactionRespository.Find(ctx, 0, data.StartWorkingAt, data.EndWorkingAt)
	if err != nil {
		return res, err
	}
	if len(transactions) > 0 {
		return res, messages.ErrWorkerIntersectWorkingDate
	}

	data.Price = uc.CalculatePrice(worker.Price, data.StartWorkingAt, data.EndWorkingAt)
	res, err = uc.transactionRespository.Store(ctx, data)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *TransactionUsecase) CalculatePrice(price float64, startWorkingAt time.Time, endWorkingAt time.Time) (res float64) {
	diff := endWorkingAt.Sub(startWorkingAt)

	diffDay := math.Floor(diff.Hours() / 24)

	modulusHour := int64(diff.Hours()) % 24

	totalHours := (diffDay * 8) + float64(modulusHour)

	return totalHours * price
}

func (uc *TransactionUsecase) Update(ctx context.Context, ID int, data *Domain) (res Domain, err error) {
	_, err = uc.FindByID(ctx, ID)
	if err != nil {
		return res, err
	}

	data.UpdatedAt = time.Now().UTC()

	res, err = uc.transactionRespository.Update(ctx, ID, data)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *TransactionUsecase) UpdateStatus(ctx context.Context, ID int, status string) (err error) {
	_, err = uc.FindByID(ctx, ID)
	if err != nil {
		return err
	}

	err = uc.transactionRespository.UpdateStatus(ctx, ID, status)
	if err != nil {
		return err
	}

	return nil
}

func (uc *TransactionUsecase) Delete(ctx context.Context, ID int) (err error) {
	_, err = uc.FindByID(ctx, ID)
	if err != nil {
		return err
	}
	err = uc.transactionRespository.Delete(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}
