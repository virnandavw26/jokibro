package admin

import (
	"context"
	"jokibro/bussiness/admin"

	"gorm.io/gorm"
)

type adminRepository struct {
	conn *gorm.DB
}

func NewAdminRepository(conn *gorm.DB) admin.Repository {
	return &adminRepository{
		conn: conn,
	}
}

func (r *adminRepository) FindByID(ctx context.Context, ID int) (admin.Domain, error) {
	var res *Admin
	err := r.conn.Where("deleted_at", nil).Where("id = ?", ID).First(&res).Error
	if err != nil {
		return admin.Domain{}, err
	}

	return *res.ToDomain(), nil
}

func (r *adminRepository) FindByEmail(ctx context.Context, email string) (admin.Domain, error) {
	var res *Admin
	err := r.conn.Where("deleted_at", nil).Where("email = ?", email).First(&res).Error
	if err != nil {
		return admin.Domain{}, err
	}

	return *res.ToDomain(), nil
}
