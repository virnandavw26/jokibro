package response

import (
	"jokibro/bussiness/transaction"
	customerResp "jokibro/controller/customer/response"
	masterServiceResp "jokibro/controller/master_service/response"
	workerResp "jokibro/controller/worker/response"
)

type Transaction struct {
	ID              int                              `json:"id"`
	CustomerID      int                              `json:"customer_id"`
	Customer        *customerResp.Customer           `json:"customer"`
	WorkerID        int                              `json:"worker_id"`
	Worker          *workerResp.Worker               `json:"worker"`
	MasterServiceID int                              `json:"master_service_id"`
	MasterService   *masterServiceResp.MasterService `json:"master_service"`
	StartWorkingAt  string                           `json:"start_working_at"`
	EndWorkingAt    string                           `json:"end_working_at"`
	Price           float64                          `json:"price"`
	Status          string                           `json:"status"`
	CreatedAt       string                           `json:"created_at"`
	UpdatedAt       string                           `json:"updated_at"`
	DeletedAt       string                           `json:"deleted_at"`
}

func FromDomain(domain *transaction.Domain) (res *Transaction) {
	if domain != nil {
		res = &Transaction{
			ID:              domain.ID,
			CustomerID:      domain.CustomerID,
			Customer:        customerResp.FromDomain(domain.Customer),
			WorkerID:        domain.WorkerID,
			Worker:          workerResp.FromDomain(domain.Worker),
			MasterServiceID: domain.MasterServiceID,
			MasterService:   masterServiceResp.FromDomain(domain.MasterService),
			StartWorkingAt:  domain.StartWorkingAt.UTC().Format("2006-01-02 15:04:05"),
			EndWorkingAt:    domain.EndWorkingAt.UTC().Format("2006-01-02 15:04:05"),
			Price:           domain.Price,
			Status:          domain.Status,
			CreatedAt:       domain.CreatedAt.UTC().Format("2006-01-02 15:04:05"),
			UpdatedAt:       domain.UpdatedAt.UTC().Format("2006-01-02 15:04:05"),
			DeletedAt:       domain.DeletedAt.UTC().Format("2006-01-02 15:04:05"),
		}

	}

	return res
}
