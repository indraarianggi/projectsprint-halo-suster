package handler

import (
	"net/http"

	"github.com/backend-magang/halo-suster/models/input"
	"github.com/backend-magang/halo-suster/utils/helper"
	"github.com/backend-magang/halo-suster/utils/pkg"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

func (h *handler) RegisterIT(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	request := input.RegisterITRequest{}
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

	request := input.LoginRequest{}
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

	request := input.LoginRequest{}
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

	request := input.RegisterNurseRequest{}
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

	request := input.NurseAccessRequest{}
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

	request := input.UpdateNurseRequest{}
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

	request := input.DeleteNurseRequest{}
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

	request := input.GetListUserRequest{
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
