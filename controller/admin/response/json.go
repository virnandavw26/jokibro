package response

import (
	"jokibro/bussiness/admin"
)

type Worker struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

func FromDomain(domain *admin.Domain) (res *Worker) {
	if domain != nil {
		res = &Worker{
			ID:        domain.ID,
			Name:      domain.Name,
			Email:     domain.Email,
			CreatedAt: domain.CreatedAt.UTC().Format("2006-01-02 15:04:05"),
			UpdatedAt: domain.UpdatedAt.UTC().Format("2006-01-02 15:04:05"),
			DeletedAt: domain.DeletedAt.UTC().Format("2006-01-02 15:04:05"),
		}
	}

	return res
}
