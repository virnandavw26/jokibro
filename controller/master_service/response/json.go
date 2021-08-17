package response

import (
	"jokibro/bussiness/master_service"
	masterCategoryResp "jokibro/controller/master_category/response"
)

type MasterService struct {
	ID               int                                `json:"id"`
	Name             string                             `json:"name"`
	MasterCategoryID int                                `json:"master_category_id"`
	MasterCategory   *masterCategoryResp.MasterCategory `json:"master_category"`
	CreatedAt        string                             `json:"created_at"`
	UpdatedAt        string                             `json:"updated_at"`
	DeletedAt        string                             `json:"deleted_at"`
}

func FromDomain(domain *master_service.Domain) (res *MasterService) {
	if domain != nil {
		res = &MasterService{
			ID:               domain.ID,
			Name:             domain.Name,
			MasterCategoryID: domain.MasterCategoryID,
			MasterCategory:   masterCategoryResp.FromDomain(domain.MasterCategory),
			CreatedAt:        domain.CreatedAt.Format("2006-01-01 15:04:05"),
			UpdatedAt:        domain.UpdatedAt.Format("2006-01-01 15:04:05"),
			DeletedAt:        domain.DeletedAt.Format("2006-01-01 15:04:05"),
		}

	}

	return res
}
