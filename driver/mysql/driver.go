package mysql_driver

import (
	"database/sql"
	"fmt"
	adminRepo "jokibro/driver/database/admin"
	customerRepo "jokibro/driver/database/customer"
	masterCategoryRepo "jokibro/driver/database/master_category"
	masterServiceRepo "jokibro/driver/database/master_service"
	transactionRepo "jokibro/driver/database/transaction"
	workerRepo "jokibro/driver/database/worker"
	"jokibro/helper/encrypt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigDB struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBDatabase string
}

func (config *ConfigDB) InitialDB() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBDatabase)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	Migrate(db)
	Seeder(db)

	return db
}

func Migrate(DB *gorm.DB) {
	DB.AutoMigrate(&masterCategoryRepo.MasterCategory{}, &masterServiceRepo.MasterService{}, &workerRepo.Worker{}, &adminRepo.Admin{}, &customerRepo.Customer{}, &transactionRepo.Transaction{})
}

func Seeder(db *gorm.DB) {
	// Seed admins
	admin := []adminRepo.Admin{}
	db.Find(&admin)
	if len(admin) == 0 {
		password, _ := encrypt.Hash("admin")
		var admin = []adminRepo.Admin{
			{Name: "Superadmin", Email: "superadmin@admin.com", Password: sql.NullString{String: password, Valid: true}, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		}
		db.Create(&admin)
	}

}
