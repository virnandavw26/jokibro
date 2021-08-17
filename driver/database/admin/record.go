package admin

import (
	"database/sql"
	"jokibro/bussiness/admin"
	"time"
)

type Admin struct {
	ID        int
	Name      string
	Email     string
	Password  sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

func fromDomain(domain *admin.Domain) *Admin {
	return &Admin{
		Name:      domain.Name,
		Email:     domain.Email,
		Password:  sql.NullString{String: domain.Password, Valid: true},
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func (model *Admin) ToDomain() (domain *admin.Domain) {
	if model != nil {
		domain = &admin.Domain{
			ID:        model.ID,
			Name:      model.Name,
			Email:     model.Email,
			Password:  model.Password.String,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
			DeletedAt: model.DeletedAt.Time,
		}
	}
	return domain
}
