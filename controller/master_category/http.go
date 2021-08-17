package master_category

import (
	"jokibro/bussiness/master_category"
	"jokibro/controller"
	"jokibro/controller/master_category/request"
	"jokibro/controller/master_category/response"
	"jokibro/helper/messages"
	"jokibro/helper/str"

	"net/http"

	"github.com/labstack/echo/v4"
)

type MasterCategoryController struct {
	masterCategoryUsecase master_category.Usecase
}

func NewMasterCategoryController(e *echo.Echo, u master_category.Usecase) *MasterCategoryController {
	return &MasterCategoryController{
		masterCategoryUsecase: u,
	}
}

func (ctrl *MasterCategoryController) Find(c echo.Context) error {
	ctx := c.Request().Context()

	search := c.QueryParam("search")

	data, err := ctrl.masterCategoryUsecase.Find(ctx, search)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	res := []response.MasterCategory{}
	for _, value := range data {
		res = append(res, *response.FromDomain(&value))
	}

	return controller.NewSuccessResponse(c, res)
}

func (ctrl *MasterCategoryController) FindByID(c echo.Context) error {
	ctx := c.Request().Context()
	ID := str.StringToInt(c.Param("id"))

	res, err := ctrl.masterCategoryUsecase.FindByID(ctx, ID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(&res))
}

func (ctrl *MasterCategoryController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.MasterCategory{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	_, err := ctrl.masterCategoryUsecase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	return controller.NewSuccessResponse(c, messages.BaseResponseMessageInserted)
}

func (ctrl *MasterCategoryController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	ID := str.StringToInt(c.Param("id"))

	req := request.MasterCategory{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	_, err := ctrl.masterCategoryUsecase.Update(ctx, ID, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	return controller.NewSuccessResponse(c, messages.BaseResponseMessageUpdated)
}

func (ctrl *MasterCategoryController) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	ID := str.StringToInt(c.Param("id"))

	err := ctrl.masterCategoryUsecase.Delete(ctx, ID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	return controller.NewSuccessResponse(c, messages.BaseResponseMessageUpdated)
}
