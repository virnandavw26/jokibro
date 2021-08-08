package routers

import (
	controller "jokibro/app/http/controller"

	"github.com/labstack/echo/v4"
)

type CategoryRoute struct {
	RouteGroup *echo.Group
	Controller *controller.Controller
}

func (r CategoryRoute) RegisterRoute() {

	c := controller.CategoryController{Controller: r.Controller}

	r.RouteGroup.GET("/category", c.CategoryController)
}
