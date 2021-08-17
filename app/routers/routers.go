package routers

import (
	"jokibro/app/middleware"
	"jokibro/controller/admin"
	"jokibro/controller/admin_auth"
	"jokibro/controller/customer"
	"jokibro/controller/customer_auth"
	"jokibro/controller/master_category"
	"jokibro/controller/master_service"
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
}

func (c *ControllerList) RouteRegister(e *echo.Echo) {
	r := e.Group("/api/v1")

	adminAuthRouter := r.Group("/admin-auth")
	adminAuthRouter.POST("/login", c.AdminAuthController.Login)

	c.JWTMiddleware.Role = "admin"

	adminRouter := r.Group("/admin")
	adminRouter.Use(c.JWTMiddleware.VerifyRole)
	adminRouter.POST("/token", c.Admincontroller.FindByToken)

	masterCategoryRouter := r.Group("/master-category")
	masterCategoryRouter.GET("", c.MasterCategoryController.Find)
	masterCategoryRouter.GET("/id/:id", c.MasterCategoryController.FindByID)
	masterCategoryRouter.POST("", c.MasterCategoryController.Store, c.JWTMiddleware.VerifyRole)
	masterCategoryRouter.PUT("/id/:id", c.MasterCategoryController.Update, c.JWTMiddleware.VerifyRole)
	masterCategoryRouter.DELETE("/id/:id", c.MasterCategoryController.Delete, c.JWTMiddleware.VerifyRole)

	masterServiceRouter := r.Group("/master-service")
	masterServiceRouter.GET("", c.MasterServiceController.Find)
	masterServiceRouter.GET("/id/:id", c.MasterServiceController.FindByID)
	masterServiceRouter.POST("", c.MasterServiceController.Store, c.JWTMiddleware.VerifyRole)
	masterServiceRouter.PUT("/id/:id", c.MasterServiceController.Update, c.JWTMiddleware.VerifyRole)
	masterServiceRouter.DELETE("/id/:id", c.MasterServiceController.Delete, c.JWTMiddleware.VerifyRole)

	workerRouter := r.Group("/worker")
	workerRouter.GET("", c.WorkerController.Find)
	workerRouter.GET("/id/:id", c.WorkerController.FindByID)
	workerRouter.POST("", c.WorkerController.Store, c.JWTMiddleware.VerifyRole)
	workerRouter.PUT("/id/:id", c.WorkerController.Update, c.JWTMiddleware.VerifyRole)
	workerRouter.DELETE("/id/:id", c.WorkerController.Delete, c.JWTMiddleware.VerifyRole)

	customerAuthRouter := r.Group("/customer-auth")
	customerAuthRouter.POST("/login", c.CustomerAuthController.Login)
	customerAuthRouter.POST("/register", c.CustomerAuthController.Register)

	c.JWTMiddleware.Role = "customer"

	customerRouter := r.Group("/customer")
	customerRouter.Use(c.JWTMiddleware.VerifyRole)
	customerRouter.GET("/token", c.CustomerController.FindByToken)
}
