package master_service

import (
	"context"
	"jokibro/bussiness/master_category"
	"time"
)

type Domain struct {
	ID               int
	MasterCategoryID int
	MasterCategory   *master_category.Domain
	Name             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time
}

type Usecase interface {
	Find(ctx context.Context, search string, masterCategoryID int) ([]Domain, error)
	FindByID(ctx context.Context, id int) (Domain, error)
	Store(ctx context.Context, data *Domain) (res Domain, err error)
	Update(ctx context.Context, ID int, data *Domain) (res Domain, err error)
	Delete(ctx context.Context, ID int) (err error)
}

type Repository interface {
	Find(ctx context.Context, search string, masterCategoryID int) ([]Domain, error)
	FindByID(ctx context.Context, id int) (Domain, error)
	Store(ctx context.Context, data *Domain) (res Domain, err error)
	Update(ctx context.Context, ID int, data *Domain) (res Domain, err error)
	Delete(ctx context.Context, ID int) (err error)
}
