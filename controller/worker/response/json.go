package response

import (
	"jokibro/bussiness/worker"
	masterCategoryResp "jokibro/controller/master_service/response"
)

type Worker struct {
	ID              int                               `json:"id"`
	Name            string                            `json:"name"`
	MasterServiceID int                               `json:"master_service_id"`
	MasterService   *masterCategoryResp.MasterService `json:"master_service"`
	BirthDate       string                            `json:"birth_date"`
	Education       string                            `json:"education"`
	Address         string                            `json:"address"`
	Price           float64                           `json:"price"`
	Description     string                            `json:"description"`
	CreatedAt       string                            `json:"created_at"`
	UpdatedAt       string                            `json:"updated_at"`
	DeletedAt       string                            `json:"deleted_at"`
}

func FromDomain(domain *worker.Domain) (res *Worker) {
	if domain != nil {
		res = &Worker{
			ID:              domain.ID,
			Name:            domain.Name,
			MasterServiceID: domain.MasterServiceID,
			MasterService:   masterCategoryResp.FromDomain(domain.MasterService),
			Education:       domain.Education,
			Address:         domain.Address,
			Price:           domain.Price,
			Description:     domain.Description,
			CreatedAt:       domain.CreatedAt.Format("2006-01-01 15:04:05"),
			UpdatedAt:       domain.UpdatedAt.Format("2006-01-01 15:04:05"),
			DeletedAt:       domain.DeletedAt.Format("2006-01-01 15:04:05"),
		}

	}

	return res
}
