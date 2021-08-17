package admin

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Usecase interface {
	FindByID(ctx context.Context, id int) (Domain, error)
	FindByEmail(ctx context.Context, email string) (Domain, error)
}

type Repository interface {
	FindByID(ctx context.Context, id int) (Domain, error)
	FindByEmail(ctx context.Context, email string) (Domain, error)
}
