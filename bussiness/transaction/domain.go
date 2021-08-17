package transaction

import (
	"context"
	"jokibro/bussiness/customer"
	"jokibro/bussiness/master_service"
	"jokibro/bussiness/worker"
	"time"
)

type Domain struct {
	ID              int
	CustomerID      int
	Customer        *customer.Domain
	WorkerID        int
	Worker          *worker.Domain
	MasterServiceID int
	MasterService   *master_service.Domain
	StartWorkingAt  time.Time
	EndWorkingAt    time.Time
	Price           float64
	Status          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}

type Usecase interface {
	Find(ctx context.Context, customerID int, startWorkingAt time.Time, endWorkingAt time.Time) ([]Domain, error)
	FindByID(ctx context.Context, id int) (Domain, error)
	Store(ctx context.Context, data *Domain) (res Domain, err error)
	Update(ctx context.Context, ID int, data *Domain) (res Domain, err error)
	UpdateStatus(ctx context.Context, ID int, status string) (err error)
	Delete(ctx context.Context, ID int) (err error)
}

type Repository interface {
	Find(ctx context.Context, customerID int, startWorkingAt time.Time, endWorkingAt time.Time) ([]Domain, error)
	FindByID(ctx context.Context, id int) (Domain, error)
	Store(ctx context.Context, data *Domain) (res Domain, err error)
	Update(ctx context.Context, ID int, data *Domain) (res Domain, err error)
	UpdateStatus(ctx context.Context, ID int, status string) (err error)
	Delete(ctx context.Context, ID int) (err error)
}
