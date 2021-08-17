package master_service

import (
	"database/sql"
	"jokibro/bussiness/master_service"
	"jokibro/driver/database/master_category"
	"time"
)

type MasterService struct {
	ID               int
	MasterCategoryID int
	MasterCategory   *master_category.MasterCategory `gorm:"foreignKey:MasterCategoryID"`
	Name             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        sql.NullTime
}

func fromDomain(domain *master_service.Domain) *MasterService {
	return &MasterService{
		Name:             domain.Name,
		MasterCategoryID: domain.MasterCategoryID,
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
	}
}

func (model *MasterService) ToDomain() (domain *master_service.Domain) {
	if model != nil {
		domain = &master_service.Domain{
			ID:               model.ID,
			Name:             model.Name,
			MasterCategoryID: model.MasterCategoryID,
			MasterCategory:   model.MasterCategory.ToDomain(),
			CreatedAt:        model.CreatedAt,
			UpdatedAt:        model.UpdatedAt,
			DeletedAt:        model.DeletedAt.Time,
		}
	}
	return domain
}
