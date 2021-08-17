package worker

import (
	"context"
	"jokibro/bussiness/worker"
	"time"

	"gorm.io/gorm"
)

type workerRepository struct {
	conn *gorm.DB
}

func NewWorkerRepository(conn *gorm.DB) worker.Repository {
	return &workerRepository{
		conn: conn,
	}
}

func (r *workerRepository) Find(ctx context.Context) ([]worker.Domain, error) {
	res := []Worker{}
	err := r.conn.Preload("MasterService").Where("deleted_at", nil).Find(&res).Error
	if err != nil {
		return []worker.Domain{}, err
	}

	workerDomain := []worker.Domain{}
	for _, value := range res {
		workerDomain = append(workerDomain, *value.ToDomain())
	}

	return workerDomain, nil
}

func (r *workerRepository) FindByID(ctx context.Context, ID int) (worker.Domain, error) {
	var res *Worker
	err := r.conn.Preload("MasterService").Where("deleted_at", nil).Where("id = ?", ID).First(&res).Error
	if err != nil {
		return worker.Domain{}, err
	}

	return *res.ToDomain(), nil
}

func (r *workerRepository) Store(ctx context.Context, data *worker.Domain) (res worker.Domain, err error) {
	model := fromDomain(data)
	result := r.conn.Create(&model)
	if result.Error != nil {
		return worker.Domain{}, result.Error
	}

	return *model.ToDomain(), err
}

func (r *workerRepository) Update(ctx context.Context, ID int, data *worker.Domain) (res worker.Domain, err error) {
	model := fromDomain(data)
	result := r.conn.Where("id = ?", ID).Save(&model)
	if result.Error != nil {
		return worker.Domain{}, result.Error
	}

	return *model.ToDomain(), err
}

func (r *workerRepository) Delete(ctx context.Context, ID int) (err error) {
	model := Worker{}
	result := r.conn.Model(&model).Where("id = ?", ID).Update("deleted_at", time.Now().UTC())
	if result.Error != nil {
		return result.Error
	}

	return err
}
