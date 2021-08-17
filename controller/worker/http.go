package worker

import (
	"jokibro/bussiness/worker"
	"jokibro/controller"
	"jokibro/controller/worker/request"
	"jokibro/controller/worker/response"
	"jokibro/helper/messages"
	"jokibro/helper/str"

	"net/http"

	"github.com/labstack/echo/v4"
)

type WorkerController struct {
	workerUsecase worker.Usecase
}

func NewWorkerController(e *echo.Echo, u worker.Usecase) *WorkerController {
	return &WorkerController{
		workerUsecase: u,
	}
}

func (ctrl *WorkerController) Find(c echo.Context) error {
	ctx := c.Request().Context()

	data, err := ctrl.workerUsecase.Find(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	res := []response.Worker{}
	for _, value := range data {
		res = append(res, *response.FromDomain(&value))
	}

	return controller.NewSuccessResponse(c, res)
}

func (ctrl *WorkerController) FindByID(c echo.Context) error {
	ctx := c.Request().Context()
	ID := str.StringToInt(c.Param("id"))

	res, err := ctrl.workerUsecase.FindByID(ctx, ID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(&res))
}

func (ctrl *WorkerController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Worker{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	reqBody, err := req.ToDomain()
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	_, err = ctrl.workerUsecase.Store(ctx, reqBody)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	return controller.NewSuccessResponse(c, messages.BaseResponseMessageInserted)
}

func (ctrl *WorkerController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	ID := str.StringToInt(c.Param("id"))

	req := request.Worker{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	reqBody, err := req.ToDomain()
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	_, err = ctrl.workerUsecase.Update(ctx, ID, reqBody)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	return controller.NewSuccessResponse(c, messages.BaseResponseMessageUpdated)
}

func (ctrl *WorkerController) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	ID := str.StringToInt(c.Param("id"))

	err := ctrl.workerUsecase.Delete(ctx, ID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	return controller.NewSuccessResponse(c, messages.BaseResponseMessageUpdated)
}
