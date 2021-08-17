package transaction

import (
	"jokibro/app/middleware"
	"jokibro/bussiness/transaction"
	"jokibro/controller"
	"jokibro/controller/transaction/request"
	"jokibro/controller/transaction/response"
	"jokibro/helper/messages"
	"jokibro/helper/str"
	"time"

	"net/http"

	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	transactionUsecase transaction.Usecase
}

func NewTransactionController(e *echo.Echo, u transaction.Usecase) *TransactionController {
	return &TransactionController{
		transactionUsecase: u,
	}
}

func (ctrl *TransactionController) Find(c echo.Context) error {
	ctx := c.Request().Context()
	customer := middleware.GetCustomer(c)
	customerID := 0
	if customer != nil {
		customerID = customer.ID
	}

	data, err := ctrl.transactionUsecase.Find(ctx, customerID, time.Time{}, time.Time{})
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	res := []response.Transaction{}
	for _, value := range data {
		res = append(res, *response.FromDomain(&value))
	}

	return controller.NewSuccessResponse(c, res)
}

func (ctrl *TransactionController) FindByID(c echo.Context) error {
	ctx := c.Request().Context()
	ID := str.StringToInt(c.Param("id"))

	res, err := ctrl.transactionUsecase.FindByID(ctx, ID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(&res))
}

func (ctrl *TransactionController) Store(c echo.Context) error {
	ctx := c.Request().Context()
	customer := middleware.GetCustomer(c)

	req := request.Transaction{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	reqBody, err := req.ToDomain()
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	reqBody.CustomerID = customer.ID

	_, err = ctrl.transactionUsecase.Store(ctx, reqBody)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	return controller.NewSuccessResponse(c, messages.BaseResponseMessageInserted)
}

func (ctrl *TransactionController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	ID := str.StringToInt(c.Param("id"))

	req := request.Transaction{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	reqBody, err := req.ToDomain()
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	_, err = ctrl.transactionUsecase.Update(ctx, ID, reqBody)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	return controller.NewSuccessResponse(c, messages.BaseResponseMessageUpdated)
}

func (ctrl *TransactionController) UpdateStatus(c echo.Context) error {
	ctx := c.Request().Context()
	ID := str.StringToInt(c.Param("id"))

	req := request.Transaction{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.transactionUsecase.UpdateStatus(ctx, ID, req.Status)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	return controller.NewSuccessResponse(c, messages.BaseResponseMessageUpdated)
}

func (ctrl *TransactionController) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	ID := str.StringToInt(c.Param("id"))

	err := ctrl.transactionUsecase.Delete(ctx, ID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	return controller.NewSuccessResponse(c, messages.BaseResponseMessageUpdated)
}
