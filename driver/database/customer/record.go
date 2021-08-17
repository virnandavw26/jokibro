package customer

import (
	"database/sql"
	"jokibro/bussiness/customer"
	"time"
)

type Customer struct {
	ID        int
	Name      string
	Email     string
	Password  sql.NullString
	Address   sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

func fromDomain(domain *customer.Domain) *Customer {
	return &Customer{
		Name:      domain.Name,
		Email:     domain.Email,
		Password:  sql.NullString{String: domain.Password, Valid: true},
		Address:   sql.NullString{String: domain.Address, Valid: true},
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func (model *Customer) ToDomain() (domain *customer.Domain) {
	if model != nil {
		domain = &customer.Domain{
			ID:        model.ID,
			Name:      model.Name,
			Email:     model.Email,
			Password:  model.Password.String,
			Address:   model.Address.String,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
			DeletedAt: model.DeletedAt.Time,
		}
	}
	return domain
}
