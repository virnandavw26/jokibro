package request

import (
	"jokibro/bussiness/master_service"
	"time"
)

type MasterService struct {
	Name             string `json:"name"`
	MasterCategoryID int    `json:"master_category_id"`
}

func (req *MasterService) ToDomain() *master_service.Domain {
	return &master_service.Domain{
		Name:             req.Name,
		MasterCategoryID: req.MasterCategoryID,
		CreatedAt:        time.Now().UTC(),
		UpdatedAt:        time.Now().UTC(),
	}
}
