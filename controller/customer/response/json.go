package response

import "jokibro/bussiness/customer"

type Customer struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

func FromDomain(domain *customer.Domain) (res *Customer) {
	if domain != nil {
		res = &Customer{
			ID:        domain.ID,
			Name:      domain.Name,
			Email:     domain.Email,
			Address:   domain.Address,
			CreatedAt: domain.CreatedAt.UTC().Format("2006-01-02 15:04:05"),
			UpdatedAt: domain.UpdatedAt.UTC().Format("2006-01-02 15:04:05"),
			DeletedAt: domain.DeletedAt.UTC().Format("2006-01-02 15:04:05"),
		}
	}

	return res
}
