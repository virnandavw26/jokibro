package transaction

import (
	"database/sql"
	"jokibro/bussiness/transaction"
	"jokibro/driver/database/customer"
	"jokibro/driver/database/master_service"
	"jokibro/driver/database/worker"
	"time"
)

type Transaction struct {
	ID              int
	CustomerID      int
	Customer        *customer.Customer `gorm:"foreignKey:CustomerID"`
	WorkerID        int
	Worker          *worker.Worker `gorm:"foreignKey:WorkerID"`
	MasterServiceID int
	MasterService   *master_service.MasterService `gorm:"foreignKey:MasterServiceID"`
	StartWorkingAt  sql.NullTime
	EndWorkingAt    sql.NullTime
	Price           sql.NullFloat64
	Status          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       sql.NullTime
}

func fromDomain(domain *transaction.Domain) *Transaction {
	return &Transaction{
		CustomerID:      domain.CustomerID,
		WorkerID:        domain.WorkerID,
		MasterServiceID: domain.MasterServiceID,
		StartWorkingAt:  sql.NullTime{Time: domain.StartWorkingAt, Valid: true},
		EndWorkingAt:    sql.NullTime{Time: domain.EndWorkingAt, Valid: true},
		Price:           sql.NullFloat64{Float64: domain.Price, Valid: true},
		Status:          domain.Status,
		CreatedAt:       domain.CreatedAt,
		UpdatedAt:       domain.UpdatedAt,
	}
}

func (model *Transaction) ToDomain() (domain *transaction.Domain) {
	if model != nil {
		domain = &transaction.Domain{
			ID:              model.ID,
			CustomerID:      model.CustomerID,
			Customer:        model.Customer.ToDomain(),
			WorkerID:        model.WorkerID,
			Worker:          model.Worker.ToDomain(),
			MasterServiceID: model.MasterServiceID,
			MasterService:   model.MasterService.ToDomain(),
			StartWorkingAt:  model.StartWorkingAt.Time,
			EndWorkingAt:    model.EndWorkingAt.Time,
			Price:           model.Price.Float64,
			Status:          model.Status,
			CreatedAt:       model.CreatedAt,
			UpdatedAt:       model.UpdatedAt,
			DeletedAt:       model.DeletedAt.Time,
		}
	}
	return domain
}
