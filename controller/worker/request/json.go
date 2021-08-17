package request

import (
	"jokibro/bussiness/worker"
	"time"
)

type Worker struct {
	Name            string  `json:"name"`
	MasterServiceID int     `json:"master_service_id"`
	BirthDate       string  `json:"birth_date"`
	Education       string  `json:"education"`
	Address         string  `json:"address"`
	Price           float64 `json:"price"`
	Description     string  `json:"description"`
}

func (req *Worker) ToDomain() (res *worker.Domain, err error) {
	birthDate, err := time.Parse("2006-01-02", req.BirthDate)
	return &worker.Domain{
		Name:            req.Name,
		MasterServiceID: req.MasterServiceID,
		BirthDate:       birthDate,
		Education:       req.Education,
		Address:         req.Address,
		Price:           req.Price,
		Description:     req.Description,
		CreatedAt:       time.Now().UTC(),
		UpdatedAt:       time.Now().UTC(),
	}, err
}
