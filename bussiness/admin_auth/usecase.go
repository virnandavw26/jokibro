package admin_auth

import (
	"context"
	"jokibro/app/middleware"
	"jokibro/bussiness/admin"
	"jokibro/helper/encrypt"
	"jokibro/helper/messages"
	"strings"
	"time"
)

type adminAuthUsecase struct {
	adminRepository admin.Repository
	contextTimeout  time.Duration
	jwtAuth         *middleware.ConfigJWT
}

func NewAdminAuthUsecase(timeout time.Duration, adminRepo admin.Repository, jwt *middleware.ConfigJWT) Usecase {
	return &adminAuthUsecase{
		adminRepository: adminRepo,
		jwtAuth:         jwt,
		contextTimeout:  timeout,
	}
}

func (uc adminAuthUsecase) Login(ctx context.Context, data *Domain) (res Domain, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if strings.TrimSpace(data.Email) == "" && strings.TrimSpace(data.Password) == "" {
		return res, messages.ErrUsernamePasswordNotFound
	}

	admin, err := uc.adminRepository.FindByEmail(ctx, data.Email)
	if err != nil {
		return res, err
	}

	if !encrypt.ValidateHash(data.Password, admin.Password) {
		return res, messages.ErrInvalidCred
	}

	res.Token = uc.jwtAuth.GenerateToken(admin.ID, "admin")

	return res, err
}
