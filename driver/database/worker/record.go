package worker

import (
	"database/sql"
	"jokibro/bussiness/worker"
	"jokibro/driver/database/master_service"
	"time"
)

type Worker struct {
	ID              int
	MasterServiceID int
	MasterService   *master_service.MasterService `gorm:"foreignKey:MasterServiceID"`
	Name            string
	BirthDate       sql.NullTime
	Education       sql.NullString
	Address         sql.NullString
	Price           sql.NullFloat64
	Description     sql.NullString
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       sql.NullTime
}

func fromDomain(domain *worker.Domain) *Worker {
	return &Worker{
		Name:            domain.Name,
		MasterServiceID: domain.MasterServiceID,
		BirthDate:       sql.NullTime{Time: domain.BirthDate, Valid: true},
		Education:       sql.NullString{String: domain.Education, Valid: true},
		Address:         sql.NullString{String: domain.Address, Valid: true},
		Price:           sql.NullFloat64{Float64: domain.Price, Valid: true},
		Description:     sql.NullString{String: domain.Description, Valid: true},
		CreatedAt:       domain.CreatedAt,
		UpdatedAt:       domain.UpdatedAt,
	}
}

func (model *Worker) ToDomain() (domain *worker.Domain) {
	if model != nil {
		domain = &worker.Domain{
			ID:              model.ID,
			Name:            model.Name,
			MasterServiceID: model.MasterServiceID,
			MasterService:   model.MasterService.ToDomain(),
			BirthDate:       model.BirthDate.Time,
			Education:       model.Education.String,
			Address:         model.Address.String,
			Price:           model.Price.Float64,
			Description:     model.Description.String,
			CreatedAt:       model.CreatedAt,
			UpdatedAt:       model.UpdatedAt,
			DeletedAt:       model.DeletedAt.Time,
		}
	}
	return domain
}
