package main

import (
	"log"
	"net/http"
	"time"

	adminAuthUsecase "jokibro/bussiness/admin_auth"
	adminAuthController "jokibro/controller/admin_auth"

	customerAuthUsecase "jokibro/bussiness/customer_auth"
	customerAuthController "jokibro/controller/customer_auth"

	customerUsecase "jokibro/bussiness/customer"
	customerController "jokibro/controller/customer"
	customerRepository "jokibro/driver/database/customer"

	adminUsecase "jokibro/bussiness/admin"
	adminController "jokibro/controller/admin"
	adminRepository "jokibro/driver/database/admin"

	masterCategoryUsecase "jokibro/bussiness/master_category"
	masterCategoryController "jokibro/controller/master_category"
	masterCategoryRepository "jokibro/driver/database/master_category"

	masterServiceUsecase "jokibro/bussiness/master_service"
	masterServiceController "jokibro/controller/master_service"
	masterServiceRepository "jokibro/driver/database/master_service"

	workerUsecase "jokibro/bussiness/worker"
	workerController "jokibro/controller/worker"
	workerRepository "jokibro/driver/database/worker"

	transactionUsecase "jokibro/bussiness/transaction"
	transactionController "jokibro/controller/transaction"
	transactionRepository "jokibro/driver/database/transaction"

	newsUsecase "jokibro/bussiness/news"
	newsController "jokibro/controller/news"
	newsRepository "jokibro/driver/thirdparties/news"

	_dbHelper "jokibro/driver/mysql"

	"jokibro/app/middleware"
	_routes "jokibro/app/routers"

	"github.com/labstack/echo/v4"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`app/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	configdb := _dbHelper.ConfigDB{
		DBUsername: viper.GetString(`database.user`),
		DBPassword: viper.GetString(`database.pass`),
		DBHost:     viper.GetString(`database.host`),
		DBPort:     viper.GetString(`database.port`),
		DBDatabase: viper.GetString(`database.name`),
	}
	db := configdb.InitialDB()

	configJWT := middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	e := echo.New()

	adminRepo := adminRepository.NewAdminRepository(db)
	adminUc := adminUsecase.NewAdminUsecase(timeoutContext, adminRepo)
	adminCtrl := adminController.NewAdminController(e, adminUc)

	adminAuthUc := adminAuthUsecase.NewAdminAuthUsecase(timeoutContext, adminRepo, &configJWT)
	adminAuthCtrl := adminAuthController.NewAdminAuthController(e, adminAuthUc)

	masterCategoryRepo := masterCategoryRepository.NewMasterCategoryRepository(db)
	masterCategoryUc := masterCategoryUsecase.NewMasterCategoryUsecase(timeoutContext, masterCategoryRepo)
	masterCategoryCtrl := masterCategoryController.NewMasterCategoryController(e, masterCategoryUc)

	masterServiceRepo := masterServiceRepository.NewMasterServiceRepository(db)
	masterServiceUc := masterServiceUsecase.NewMasterServiceUsecase(timeoutContext, masterServiceRepo)
	masterServiceCtrl := masterServiceController.NewMasterServiceController(e, masterServiceUc)

	workerRepo := workerRepository.NewWorkerRepository(db)
	workerUc := workerUsecase.NewWorkerUsecase(timeoutContext, workerRepo)
	workerCtrl := workerController.NewWorkerController(e, workerUc)

	customerRepo := customerRepository.NewCustomerRepository(db)
	customerUc := customerUsecase.NewCustomerUsecase(timeoutContext, customerRepo)
	customerCtrl := customerController.NewCustomerController(e, customerUc)

	customerAuthUc := customerAuthUsecase.NewCustomerAuthUsecase(timeoutContext, customerRepo, &configJWT)
	customerAuthCtrl := customerAuthController.NewCustomerAuthController(e, customerAuthUc)

	transactionRepo := transactionRepository.NewTransactionRepository(db)
	transactionUc := transactionUsecase.NewTransactionUsecase(timeoutContext, transactionRepo, workerRepo)
	transactionCtrl := transactionController.NewTransactionController(e, transactionUc)

	newsRepo := newsRepository.NewNewsRepository(viper.GetString(`newsapi.host`), viper.GetString(`newsapi.key`), &http.Client{})
	newsUc := newsUsecase.NewNewsUsecase(timeoutContext, newsRepo)
	newsCtrl := newsController.NewNewsController(e, newsUc)

	routesInit := _routes.ControllerList{
		JWTMiddleware:            &configJWT,
		AdminAuthController:      *adminAuthCtrl,
		Admincontroller:          *adminCtrl,
		MasterCategoryController: *masterCategoryCtrl,
		MasterServiceController:  *masterServiceCtrl,
		WorkerController:         *workerCtrl,
		CustomerAuthController:   *customerAuthCtrl,
		CustomerController:       *customerCtrl,
		TransactionController:    *transactionCtrl,
		NewsController:           *newsCtrl,
	}

	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))

}
