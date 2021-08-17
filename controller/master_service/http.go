package master_service

import (
	"jokibro/bussiness/master_service"
	"jokibro/controller"
	"jokibro/controller/master_service/request"
	"jokibro/controller/master_service/response"
	"jokibro/helper/messages"
	"jokibro/helper/str"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

type MasterServiceController struct {
	masterServiceUsecase master_service.Usecase
}

func NewMasterServiceController(e *echo.Echo, u master_service.Usecase) *MasterServiceController {
	return &MasterServiceController{
		masterServiceUsecase: u,
	}
}

func (ctrl *MasterServiceController) Find(c echo.Context) error {
	ctx := c.Request().Context()

	masterCategoryID, _ := strconv.Atoi(c.QueryParam("master_category_id"))
	search := c.QueryParam("search")

	data, err := ctrl.masterServiceUsecase.Find(ctx, search, masterCategoryID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	res := []response.MasterService{}
	for _, value := range data {
		res = append(res, *response.FromDomain(&value))
	}

	return controller.NewSuccessResponse(c, res)
}

func (ctrl *MasterServiceController) FindByID(c echo.Context) error {
	ctx := c.Request().Context()
	ID := str.StringToInt(c.Param("id"))

	res, err := ctrl.masterServiceUsecase.FindByID(ctx, ID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(&res))
}

func (ctrl *MasterServiceController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.MasterService{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	_, err := ctrl.masterServiceUsecase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	return controller.NewSuccessResponse(c, messages.BaseResponseMessageInserted)
}

func (ctrl *MasterServiceController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	ID := str.StringToInt(c.Param("id"))

	req := request.MasterService{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	_, err := ctrl.masterServiceUsecase.Update(ctx, ID, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	return controller.NewSuccessResponse(c, messages.BaseResponseMessageUpdated)
}

func (ctrl *MasterServiceController) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	ID := str.StringToInt(c.Param("id"))

	err := ctrl.masterServiceUsecase.Delete(ctx, ID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	return controller.NewSuccessResponse(c, messages.BaseResponseMessageUpdated)
}
