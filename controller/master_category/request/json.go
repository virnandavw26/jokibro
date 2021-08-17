package request

import (
	"jokibro/bussiness/master_category"
	"time"
)

type MasterCategory struct {
	Name string `json:"name"`
}

func (req *MasterCategory) ToDomain() *master_category.Domain {
	return &master_category.Domain{
		Name:      req.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}
