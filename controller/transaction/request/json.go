package request

import (
	"jokibro/bussiness/transaction"
	"jokibro/helper/messages"
	"time"
)

type Transaction struct {
	Name           string  `json:"name"`
	WorkerID       int     `json:"worker_id"`
	StartWorkingAt string  `json:"start_working_at"`
	EndWorkingAt   string  `json:"end_working_at"`
	Status         string  `json:"status"`
	Price          float64 `json:"price"`
}

func (req *Transaction) ToDomain() (res *transaction.Domain, err error) {

	startWorkingAt, err := time.Parse("2006-01-02 15:04 -0700", req.StartWorkingAt)
	if err != nil {
		return nil, err
	}
	endWorkingAt, err := time.Parse("2006-01-02 15:04 -0700", req.EndWorkingAt)
	if err != nil {
		return nil, err
	}

	if startWorkingAt.UTC().Before(time.Now().UTC()) {
		return nil, messages.ErrInvalidStartWorkingAt
	}
	if endWorkingAt.UTC().Before(startWorkingAt) {
		return nil, messages.ErrInvalidEndWorkingAt
	}

	return &transaction.Domain{
		WorkerID:       req.WorkerID,
		Status:         "pending",
		StartWorkingAt: startWorkingAt.UTC(),
		EndWorkingAt:   endWorkingAt.UTC(),
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
	}, err
}
