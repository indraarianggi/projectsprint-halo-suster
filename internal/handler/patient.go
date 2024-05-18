package handler

import (
	"net/http"

	"github.com/backend-magang/halo-suster/models/input"
	"github.com/backend-magang/halo-suster/utils/helper"
	"github.com/backend-magang/halo-suster/utils/pkg"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

func (h *handler) AddPatient(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	request := input.AddPatientRequest{}
	err = pkg.BindValidate(c, &request)
	if err != nil {
		return helper.WriteResponse(c, helper.StandardResponse{Code: http.StatusBadRequest, Error: err})
	}

	response := h.usecase.AddPatient(ctx, request)
	return helper.WriteResponse(c, response)
}

func (h *handler) GetListPatient(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	request := input.GetListPatientRequest{
		IdentityNumber: c.QueryParam("identityNumber"),
		PhoneNumber:    c.QueryParam("phoneNumber"),
		Name:           c.QueryParam("name"),
		CreatedAt:      c.QueryParam("createdAt"),
		Limit:          c.QueryParam("limit"),
		Offset:         c.QueryParam("offset"),
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

	response := h.usecase.GetListPatient(ctx, request)
	return helper.WriteResponse(c, response)
}
