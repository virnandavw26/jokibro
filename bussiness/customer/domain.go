package customer

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	Name      string
	Email     string
	Password  string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Usecase interface {
	Find(ctx context.Context) ([]Domain, error)
	FindByID(ctx context.Context, id int) (Domain, error)
	FindByEmail(ctx context.Context, email string) (Domain, error)
	Store(ctx context.Context, data *Domain) (res Domain, err error)
	Update(ctx context.Context, ID int, data *Domain) (res Domain, err error)
	Delete(ctx context.Context, ID int) (err error)
}

type Repository interface {
	Find(ctx context.Context) ([]Domain, error)
	FindByID(ctx context.Context, id int) (Domain, error)
	FindByEmail(ctx context.Context, email string) (Domain, error)
	Store(ctx context.Context, data *Domain) (res Domain, err error)
	Update(ctx context.Context, ID int, data *Domain) (res Domain, err error)
	Delete(ctx context.Context, ID int) (err error)
}
