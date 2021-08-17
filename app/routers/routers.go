package routers

import (
	"jokibro/app/middleware"
	"jokibro/controller/admin"
	"jokibro/controller/admin_auth"
	"jokibro/controller/customer"
	"jokibro/controller/customer_auth"
	"jokibro/controller/master_category"
	"jokibro/controller/master_service"
	"jokibro/controller/transaction"
	"jokibro/controller/worker"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	JWTMiddleware            *middleware.ConfigJWT
	AdminAuthController      admin_auth.AdminAuthController
	CustomerAuthController   customer_auth.CustomerAuthController
	Admincontroller          admin.AdminController
	CustomerController       customer.CustomerController
	MasterCategoryController master_category.MasterCategoryController
	MasterServiceController  master_service.MasterServiceController
	WorkerController         worker.WorkerController
	TransactionController    transaction.TransactionController
}

func (c *ControllerList) RouteRegister(e *echo.Echo) {
	adminMiddleware := *c.JWTMiddleware
	adminMiddleware.Role = "admin"
	customerMiddleware := *c.JWTMiddleware
	customerMiddleware.Role = "customer"

	r := e.Group("/api/v1")

	adminAuthRouter := r.Group("/admin-auth")
	adminAuthRouter.POST("/login", c.AdminAuthController.Login)

	adminRouter := r.Group("/admin")
	adminRouter.Use(adminMiddleware.VerifyRole)
	adminRouter.POST("/token", c.Admincontroller.FindByToken)

	adminTransactionRouter := adminRouter.Group("/order")
	adminTransactionRouter.GET("", c.TransactionController.Find)
	adminTransactionRouter.GET("/id/:id", c.TransactionController.FindByID)
	adminTransactionRouter.PUT("/id/:id", c.TransactionController.UpdateStatus)

	masterCategoryRouter := r.Group("/master-category")
	masterCategoryRouter.GET("", c.MasterCategoryController.Find)
	masterCategoryRouter.GET("/id/:id", c.MasterCategoryController.FindByID)
	masterCategoryRouter.POST("", c.MasterCategoryController.Store, adminMiddleware.VerifyRole)
	masterCategoryRouter.PUT("/id/:id", c.MasterCategoryController.Update, adminMiddleware.VerifyRole)
	masterCategoryRouter.DELETE("/id/:id", c.MasterCategoryController.Delete, adminMiddleware.VerifyRole)

	masterServiceRouter := r.Group("/master-service")
	masterServiceRouter.GET("", c.MasterServiceController.Find)
	masterServiceRouter.GET("/id/:id", c.MasterServiceController.FindByID)
	masterServiceRouter.POST("", c.MasterServiceController.Store, adminMiddleware.VerifyRole)
	masterServiceRouter.PUT("/id/:id", c.MasterServiceController.Update, adminMiddleware.VerifyRole)
	masterServiceRouter.DELETE("/id/:id", c.MasterServiceController.Delete, adminMiddleware.VerifyRole)

	workerRouter := r.Group("/worker")
	workerRouter.GET("", c.WorkerController.Find)
	workerRouter.GET("/id/:id", c.WorkerController.FindByID)
	workerRouter.POST("", c.WorkerController.Store, adminMiddleware.VerifyRole)
	workerRouter.PUT("/id/:id", c.WorkerController.Update, adminMiddleware.VerifyRole)
	workerRouter.DELETE("/id/:id", c.WorkerController.Delete, adminMiddleware.VerifyRole)

	customerAuthRouter := r.Group("/customer-auth")
	customerAuthRouter.POST("/login", c.CustomerAuthController.Login)
	customerAuthRouter.POST("/register", c.CustomerAuthController.Register)

	customerRouter := r.Group("/customer")
	customerRouter.Use(customerMiddleware.VerifyRole)
	customerRouter.GET("/token", c.CustomerController.FindByToken)

	transactionRouter := r.Group("/order")
	transactionRouter.GET("", c.TransactionController.Find, customerMiddleware.VerifyRole)
	transactionRouter.GET("/id/:id", c.TransactionController.FindByID, customerMiddleware.VerifyRole)
	transactionRouter.POST("", c.TransactionController.Store, customerMiddleware.VerifyRole)
}
