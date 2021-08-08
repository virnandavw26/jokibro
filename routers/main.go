package routers

import (
	"os"

	controller "jokibro/app/http/controller"

	"github.com/labstack/echo/v4"
)

type MainRoute struct {
	Controller *controller.Controller
	RouteGroup *echo.Group
}

func (route MainRoute) RegisterRoute() {
	e := echo.New()

	v1 := e.Group("/api/v1")

	categoryRoute := CategoryRoute{RouteGroup: v1, Controller: route.Controller}
	categoryRoute.RegisterRoute()

	e.Logger.Fatal(e.Start(os.Getenv("APP_HOST")))
}
