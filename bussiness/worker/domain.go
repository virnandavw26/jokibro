package worker

import (
	"context"
	"jokibro/bussiness/master_service"
	"time"
)

type Domain struct {
	ID              int
	MasterServiceID int
	MasterService   *master_service.Domain
	Name            string
	BirthDate       time.Time
	Education       string
	Address         string
	Price           float64
	Description     string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}

type Usecase interface {
	Find(ctx context.Context) ([]Domain, error)
	FindByID(ctx context.Context, id int) (Domain, error)
	Store(ctx context.Context, data *Domain) (res Domain, err error)
	Update(ctx context.Context, ID int, data *Domain) (res Domain, err error)
	Delete(ctx context.Context, ID int) (err error)
}

type Repository interface {
	Find(ctx context.Context) ([]Domain, error)
	FindByID(ctx context.Context, id int) (Domain, error)
	Store(ctx context.Context, data *Domain) (res Domain, err error)
	Update(ctx context.Context, ID int, data *Domain) (res Domain, err error)
	Delete(ctx context.Context, ID int) (err error)
}
