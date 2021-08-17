package response

import "jokibro/bussiness/admin_auth"

type AdminAuthResponse struct {
	Token string `json:"token"`
}

func FromDomain(domain *admin_auth.Domain) (res *AdminAuthResponse) {
	if domain != nil {
		res = &AdminAuthResponse{
			Token: domain.Token,
		}
	}

	return res
}
