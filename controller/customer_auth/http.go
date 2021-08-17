package customer_auth

import (
	"jokibro/bussiness/customer_auth"
	"jokibro/controller"
	customerReq "jokibro/controller/customer/request"
	customerAuthReq "jokibro/controller/customer_auth/request"
	"jokibro/controller/customer_auth/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CustomerAuthController struct {
	customerAuthUsecase customer_auth.Usecase
}

func NewCustomerAuthController(e *echo.Echo, uc customer_auth.Usecase) *CustomerAuthController {
	return &CustomerAuthController{
		customerAuthUsecase: uc,
	}
}

func (ctrl *CustomerAuthController) Register(c echo.Context) error {
	ctx := c.Request().Context()

	req := customerReq.Customer{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.customerAuthUsecase.Register(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnauthorized, err)
	}

	return controller.NewSuccessResponse(c, resp)

}

func (ctrl *CustomerAuthController) Login(c echo.Context) error {
	ctx := c.Request().Context()

	req := customerAuthReq.CustomerAuth{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.customerAuthUsecase.Login(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnauthorized, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(&resp))

}
