package master_service

import (
	"context"
	"jokibro/bussiness/master_service"
	"time"

	"gorm.io/gorm"
)

type masterServiceRepository struct {
	conn *gorm.DB
}

func NewMasterServiceRepository(conn *gorm.DB) master_service.Repository {
	return &masterServiceRepository{
		conn: conn,
	}
}

func (r *masterServiceRepository) Find(ctx context.Context, search string, masterCategoryID int) ([]master_service.Domain, error) {
	res := []MasterService{}

	query := r.conn.Preload("MasterCategory").Where("deleted_at", nil)

	if masterCategoryID != 0 {
		query = query.Where("master_category_id", masterCategoryID)
	}

	if search != "" {
		query = query.Where("name like ?", "%"+search+"%")
	}

	err := query.Find(&res).Error
	if err != nil {
		return []master_service.Domain{}, err
	}

	masterServiceDomain := []master_service.Domain{}
	for _, value := range res {
		masterServiceDomain = append(masterServiceDomain, *value.ToDomain())
	}

	return masterServiceDomain, nil
}

func (r *masterServiceRepository) FindByID(ctx context.Context, ID int) (master_service.Domain, error) {
	var res *MasterService
	err := r.conn.Preload("MasterCategory").Where("deleted_at", nil).Where("id = ?", ID).First(&res).Error
	if err != nil {
		return master_service.Domain{}, err
	}

	return *res.ToDomain(), nil
}

func (r *masterServiceRepository) Store(ctx context.Context, data *master_service.Domain) (res master_service.Domain, err error) {
	model := fromDomain(data)
	result := r.conn.Create(&model)
	if result.Error != nil {
		return master_service.Domain{}, result.Error
	}

	return *model.ToDomain(), err
}

func (r *masterServiceRepository) Update(ctx context.Context, ID int, data *master_service.Domain) (res master_service.Domain, err error) {
	model := fromDomain(data)
	result := r.conn.Where("id = ?", ID).Save(&model)
	if result.Error != nil {
		return master_service.Domain{}, result.Error
	}

	return *model.ToDomain(), err
}

func (r *masterServiceRepository) Delete(ctx context.Context, ID int) (err error) {
	model := MasterService{}
	result := r.conn.Model(&model).Where("id = ?", ID).Update("deleted_at", time.Now().UTC())
	if result.Error != nil {
		return result.Error
	}

	return err
}
