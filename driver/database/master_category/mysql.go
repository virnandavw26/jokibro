package master_category

import (
	"context"
	"jokibro/bussiness/master_category"
	"time"

	"gorm.io/gorm"
)

type masterCategoryRepository struct {
	conn *gorm.DB
}

func NewMasterCategoryRepository(conn *gorm.DB) master_category.Repository {
	return &masterCategoryRepository{
		conn: conn,
	}
}

func (r *masterCategoryRepository) Find(ctx context.Context, search string) ([]master_category.Domain, error) {
	res := []MasterCategory{}
	query := r.conn.Where("deleted_at", nil)

	if search != "" {
		query = query.Where("name like ?", "%"+search+"%")
	}

	err := query.Find(&res).Error
	if err != nil {
		return []master_category.Domain{}, err
	}

	masterCategoryDomain := []master_category.Domain{}
	for _, value := range res {
		masterCategoryDomain = append(masterCategoryDomain, *value.ToDomain())
	}

	return masterCategoryDomain, nil
}

func (r *masterCategoryRepository) FindByID(ctx context.Context, ID int) (master_category.Domain, error) {
	var res *MasterCategory
	err := r.conn.Where("deleted_at", nil).Where("id = ?", ID).First(&res).Error
	if err != nil {
		return master_category.Domain{}, err
	}

	return *res.ToDomain(), nil
}

func (r *masterCategoryRepository) Store(ctx context.Context, data *master_category.Domain) (res master_category.Domain, err error) {
	model := fromDomain(data)
	result := r.conn.Create(&model)
	if result.Error != nil {
		return master_category.Domain{}, result.Error
	}

	return *model.ToDomain(), err
}

func (r *masterCategoryRepository) Update(ctx context.Context, ID int, data *master_category.Domain) (res master_category.Domain, err error) {
	model := fromDomain(data)
	result := r.conn.Where("id = ?", ID).Save(&model)
	if result.Error != nil {
		return master_category.Domain{}, result.Error
	}

	return *model.ToDomain(), err
}

func (r *masterCategoryRepository) Delete(ctx context.Context, ID int) (err error) {
	model := MasterCategory{}
	result := r.conn.Model(&model).Where("id = ?", ID).Update("deleted_at", time.Now().UTC())
	if result.Error != nil {
		return result.Error
	}

	return err
}
