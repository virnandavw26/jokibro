package response

import "jokibro/bussiness/customer_auth"

type CustomerAuthResponse struct {
	Token string `json:"token"`
}

func FromDomain(domain *customer_auth.Domain) (res *CustomerAuthResponse) {
	if domain != nil {
		res = &CustomerAuthResponse{
			Token: domain.Token,
		}
	}

	return res
}
