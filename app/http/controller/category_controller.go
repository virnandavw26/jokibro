package controllers

import (
	"jokibro/app/http/request"
	service "jokibro/app/service/category_service"
	"jokibro/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	*Controller
}

func (controller CategoryController) CategoryController(c echo.Context) error {
	uc := service.CategoryUsecase{DB: controller.DB}
	res := helpers.BaseResponse{}
	req := new(request.CategoryRequest)

	if err := c.Bind(req); err != nil {
		res = helpers.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: helpers.MessageFailed,
			Data:    nil,
		}

		return c.JSON(http.StatusOK, res)
	}
	data, err := uc.FindAll(req.Name)
	if err != nil {
		res = helpers.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: helpers.MessageFailed,
			Data:    nil,
		}

		return c.JSON(http.StatusOK, res)
	}

	res = helpers.BaseResponse{
		Code:    http.StatusOK,
		Message: helpers.MessageSuccess,
		Data:    data,
	}

	return c.JSON(http.StatusOK, res)

}
