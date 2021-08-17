package customer

import (
	"context"
	"jokibro/bussiness/customer"
	"time"

	"gorm.io/gorm"
)

type customerRepository struct {
	conn *gorm.DB
}

func NewCustomerRepository(conn *gorm.DB) customer.Repository {
	return &customerRepository{
		conn: conn,
	}
}

func (r *customerRepository) Find(ctx context.Context) ([]customer.Domain, error) {
	res := []Customer{}
	err := r.conn.Preload("MasterService").Where("deleted_at", nil).Find(&res).Error
	if err != nil {
		return []customer.Domain{}, err
	}

	customerDomain := []customer.Domain{}
	for _, value := range res {
		customerDomain = append(customerDomain, *value.ToDomain())
	}

	return customerDomain, nil
}

func (r *customerRepository) FindByID(ctx context.Context, ID int) (customer.Domain, error) {
	var res *Customer
	err := r.conn.Where("deleted_at", nil).Where("id = ?", ID).First(&res).Error
	if err != nil {
		return customer.Domain{}, err
	}

	return *res.ToDomain(), nil
}

func (r *customerRepository) FindByEmail(ctx context.Context, email string) (customer.Domain, error) {
	var res *Customer
	err := r.conn.Where("deleted_at", nil).Where("email = ?", email).First(&res).Error
	if err != nil {
		return customer.Domain{}, err
	}

	return *res.ToDomain(), nil
}

func (r *customerRepository) Store(ctx context.Context, data *customer.Domain) (res customer.Domain, err error) {
	model := fromDomain(data)
	result := r.conn.Create(&model)
	if result.Error != nil {
		return customer.Domain{}, result.Error
	}

	return *model.ToDomain(), err
}

func (r *customerRepository) Update(ctx context.Context, ID int, data *customer.Domain) (res customer.Domain, err error) {
	model := fromDomain(data)
	result := r.conn.Where("id = ?", ID).Save(&model)
	if result.Error != nil {
		return customer.Domain{}, result.Error
	}

	return *model.ToDomain(), err
}

func (r *customerRepository) Delete(ctx context.Context, ID int) (err error) {
	model := Customer{}
	result := r.conn.Model(&model).Where("id = ?", ID).Update("deleted_at", time.Now().UTC())
	if result.Error != nil {
		return result.Error
	}

	return err
}
