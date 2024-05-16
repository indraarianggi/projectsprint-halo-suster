package user

import (
	"net/http"

	"github.com/backend-magang/halo-suster/utils/helper"
	"github.com/backend-magang/halo-suster/utils/pkg"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Handler interface {
	RegisterIT(echo.Context) error
}

type handler struct {
	usecase Usecase
	logger  *logrus.Logger
}

func NewHandler(usecase Usecase, logger *logrus.Logger) Handler {
	return &handler{usecase, logger}
}

func (h *handler) RegisterIT(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	request := RegisterITRequest{}
	err = pkg.BindValidate(c, &request)
	if err != nil {
		return helper.WriteResponse(c, helper.StandardResponse{Code: http.StatusBadRequest, Error: err})
	}

	response := h.usecase.RegisterIT(ctx, request)
	return helper.WriteResponse(c, response)
}
