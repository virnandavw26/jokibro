package transaction

import (
	"context"
	"database/sql"
	"jokibro/bussiness/transaction"
	"time"

	"gorm.io/gorm"
)

type transactionRepository struct {
	conn *gorm.DB
}

func NewTransactionRepository(conn *gorm.DB) transaction.Repository {
	return &transactionRepository{
		conn: conn,
	}
}

func (r *transactionRepository) Find(ctx context.Context, customerID int, startWorkingAt time.Time, endWorkingAt time.Time) ([]transaction.Domain, error) {
	res := []Transaction{}

	query := r.conn.Preload("Customer").Preload("Worker").Preload("MasterService").Where("deleted_at", nil)

	var nullTime sql.NullTime
	if startWorkingAt != nullTime.Time {
		groupWhere := r.conn.Where(r.conn.Where("start_working_at <= ? AND end_working_at >= ? AND status IN ?", startWorkingAt, startWorkingAt, []string{"pending", "paid"}))
		if endWorkingAt != nullTime.Time {
			groupWhere = groupWhere.Or(r.conn.Where("start_working_at <= ? AND end_working_at >= ? AND status IN ?", endWorkingAt, endWorkingAt, []string{"pending", "paid"}))
		}
		query = query.Where(groupWhere)
	}

	if customerID != 0 {
		query = query.Where("customer_id", customerID)
	}

	err := query.Find(&res).Error
	if err != nil {
		return []transaction.Domain{}, err
	}

	transactionDomain := []transaction.Domain{}
	for _, value := range res {
		transactionDomain = append(transactionDomain, *value.ToDomain())
	}

	return transactionDomain, nil
}

func (r *transactionRepository) FindByID(ctx context.Context, ID int) (transaction.Domain, error) {
	var res *Transaction
	err := r.conn.Preload("Customer").Preload("Worker").Preload("MasterService").Where("deleted_at", nil).Where("id = ?", ID).First(&res).Error
	if err != nil {
		return transaction.Domain{}, err
	}

	return *res.ToDomain(), nil
}

func (r *transactionRepository) Store(ctx context.Context, data *transaction.Domain) (res transaction.Domain, err error) {
	model := fromDomain(data)
	result := r.conn.Create(&model)
	if result.Error != nil {
		return transaction.Domain{}, result.Error
	}

	return *model.ToDomain(), err
}

func (r *transactionRepository) Update(ctx context.Context, ID int, data *transaction.Domain) (res transaction.Domain, err error) {
	model := fromDomain(data)
	result := r.conn.Where("id = ?", ID).Save(&model)
	if result.Error != nil {
		return transaction.Domain{}, result.Error
	}

	return *model.ToDomain(), err
}

func (r *transactionRepository) UpdateStatus(ctx context.Context, ID int, status string) (err error) {
	transaction := Transaction{}
	result := r.conn.Model(&transaction).Where("id = ?", ID).Update("status", status)
	if result.Error != nil {
		return result.Error
	}

	return err
}

func (r *transactionRepository) Delete(ctx context.Context, ID int) (err error) {
	model := Transaction{}
	result := r.conn.Model(&model).Where("id = ?", ID).Update("deleted_at", time.Now().UTC())
	if result.Error != nil {
		return result.Error
	}

	return err
}
