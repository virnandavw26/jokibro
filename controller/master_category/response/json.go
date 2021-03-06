package response

import (
	"jokibro/bussiness/master_category"
)

type MasterCategory struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

func FromDomain(domain *master_category.Domain) (res *MasterCategory) {
	if domain != nil {
		res = &MasterCategory{
			ID:        domain.ID,
			Name:      domain.Name,
			CreatedAt: domain.CreatedAt.UTC().Format("2006-01-02 15:04:05"),
			UpdatedAt: domain.UpdatedAt.UTC().Format("2006-01-02 15:04:05"),
			DeletedAt: domain.DeletedAt.UTC().Format("2006-01-02 15:04:05"),
		}
	}
	return res
}
