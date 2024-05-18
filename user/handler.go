package user

import (
	"net/http"

	"github.com/backend-magang/halo-suster/utils/helper"
	"github.com/backend-magang/halo-suster/utils/pkg"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

type Handler interface {
	RegisterIT(echo.Context) error
	LoginIT(echo.Context) error
	RegisterNurse(echo.Context) error
	LoginNurse(echo.Context) error
	SetPasswordNurse(echo.Context) error
	UpdateNurse(echo.Context) error
	DeleteNurse(echo.Context) error
	GetListUser(echo.Context) error
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

func (h *handler) LoginIT(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	request := LoginRequest{}
	err = pkg.BindValidate(c, &request)
	if err != nil {
		return helper.WriteResponse(c, helper.StandardResponse{Code: http.StatusBadRequest, Error: err})
	}

	response := h.usecase.LoginIT(ctx, request)
	return helper.WriteResponse(c, response)
}

func (h *handler) LoginNurse(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	request := LoginRequest{}
	err = pkg.BindValidate(c, &request)
	if err != nil {
		return helper.WriteResponse(c, helper.StandardResponse{Code: http.StatusBadRequest, Error: err})
	}

	response := h.usecase.LoginNurse(ctx, request)
	return helper.WriteResponse(c, response)
}

func (h *handler) RegisterNurse(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	request := RegisterNurseRequest{}
	err = pkg.BindValidate(c, &request)
	if err != nil {
		return helper.WriteResponse(c, helper.StandardResponse{Code: http.StatusBadRequest, Error: err})
	}

	response := h.usecase.RegisterNurse(ctx, request)
	return helper.WriteResponse(c, response)
}

func (h *handler) SetPasswordNurse(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	request := NurseAccessRequest{}
	err = pkg.BindValidate(c, &request)
	if err != nil {
		return helper.WriteResponse(c, helper.StandardResponse{Code: http.StatusBadRequest, Error: err})
	}

	response := h.usecase.SetPasswordNurse(ctx, request)
	return helper.WriteResponse(c, response)
}

func (h *handler) UpdateNurse(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	request := UpdateNurseRequest{}
	err = pkg.BindValidate(c, &request)
	if err != nil {
		return helper.WriteResponse(c, helper.StandardResponse{Code: http.StatusBadRequest, Error: err})
	}

	response := h.usecase.UpdateNurse(ctx, request)
	return helper.WriteResponse(c, response)
}

func (h *handler) DeleteNurse(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	request := DeleteNurseRequest{}
	err = pkg.BindValidate(c, &request)
	if err != nil {
		return helper.WriteResponse(c, helper.StandardResponse{Code: http.StatusBadRequest, Error: err})
	}

	response := h.usecase.DeleteNurse(ctx, request)
	return helper.WriteResponse(c, response)
}

func (h *handler) GetListUser(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	request := GetListUserRequest{
		ID:        c.QueryParam("userId"),
		NIP:       c.QueryParam("nip"),
		Name:      c.QueryParam("name"),
		Role:      c.QueryParam("role"),
		CreatedAt: c.QueryParam("createdAt"),
		IsDeleted: c.QueryParam("isDeleted"),
		Limit:     c.QueryParam("limit"),
		Offset:    c.QueryParam("offset"),
	}

	err = pkg.BindValidate(c, &request)
	if err != nil {
		return helper.WriteResponse(c, helper.StandardResponse{Code: http.StatusBadRequest, Error: err})
	}

	if cast.ToInt(request.Limit) == 0 {
		request.Limit = "5"
	}

	if cast.ToInt(request.Offset) == 0 {
		request.Offset = "0"
	}

	response := h.usecase.GetListUser(ctx, request)
	return helper.WriteResponse(c, response)
}
