package master_category

import (
	"database/sql"
	"jokibro/bussiness/master_category"
	"time"
)

type MasterCategory struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

func fromDomain(domain *master_category.Domain) *MasterCategory {
	return &MasterCategory{
		Name:      domain.Name,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func (model *MasterCategory) ToDomain() (domain *master_category.Domain) {
	if model != nil {
		domain = &master_category.Domain{
			ID:        model.ID,
			Name:      model.Name,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
			DeletedAt: model.DeletedAt.Time,
		}
	}

	return domain
}
