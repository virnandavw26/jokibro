package customer

import (
	"jokibro/app/middleware"
	"jokibro/bussiness/customer"
	"jokibro/controller"
	"jokibro/controller/customer/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CustomerController struct {
	customerUsecase customer.Usecase
}

func NewCustomerController(e *echo.Echo, uc customer.Usecase) *CustomerController {
	return &CustomerController{
		customerUsecase: uc,
	}
}

func (ctrl *CustomerController) FindByToken(c echo.Context) error {
	ctx := c.Request().Context()
	customer := middleware.GetCustomer(c)

	res, err := ctrl.customerUsecase.FindByID(ctx, customer.ID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(&res))
}
