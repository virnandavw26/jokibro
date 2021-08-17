package admin

import (
	"jokibro/app/middleware"
	"jokibro/bussiness/admin"
	"jokibro/controller"
	"jokibro/controller/admin/response"

	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminController struct {
	adminUsecase admin.Usecase
}

func NewAdminController(e *echo.Echo, u admin.Usecase) *AdminController {
	return &AdminController{
		adminUsecase: u,
	}
}

func (ctrl *AdminController) FindByToken(c echo.Context) error {
	ctx := c.Request().Context()
	admin := middleware.GetAdmin(c)

	res, err := ctrl.adminUsecase.FindByID(ctx, admin.ID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(&res))
}
