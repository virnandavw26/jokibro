package news

import (
	"jokibro/bussiness/news"
	"jokibro/controller"
	"jokibro/controller/news/response"

	"net/http"

	"github.com/labstack/echo/v4"
)

type NewsController struct {
	newsUsecase news.Usecase
}

func NewNewsController(e *echo.Echo, u news.Usecase) *NewsController {
	return &NewsController{
		newsUsecase: u,
	}
}

func (ctrl *NewsController) Find(c echo.Context) error {
	ctx := c.Request().Context()

	data, err := ctrl.newsUsecase.Find(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	res := []response.News{}
	for _, value := range data {
		res = append(res, *response.FromDomain(&value))
	}

	return controller.NewSuccessResponse(c, res)
}
