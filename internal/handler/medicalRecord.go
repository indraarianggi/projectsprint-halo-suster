package handler

import (
	"net/http"

	"github.com/backend-magang/halo-suster/middleware"
	"github.com/backend-magang/halo-suster/models/input"
	"github.com/backend-magang/halo-suster/utils/helper"
	"github.com/backend-magang/halo-suster/utils/pkg"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

func (h *handler) AddMedicalRecord(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	userClaims := middleware.GetUserClaims(c)

	request := input.AddMedicalRecordRequest{
		CreatedByID:  userClaims.ID,
		CreatedByNIP: userClaims.NIP,
	}

	err = pkg.BindValidate(c, &request)
	if err != nil {
		return helper.WriteResponse(c, helper.StandardResponse{Code: http.StatusBadRequest, Error: err})
	}

	response := h.usecase.AddMedicalRecord(ctx, request)
	return helper.WriteResponse(c, response)
}

func (h *handler) GetListMedicalRecord(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	request := input.GetListMedicalRecordRequest{
		IdentityNumber: c.QueryParam("identityDetail.identityNumber"),
		CreatedByID:    c.QueryParam("createdBy.userId"),
		CreatedByNIP:   c.QueryParam("createdBy.nip"),
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

	response := h.usecase.GetListMedicalRecord(ctx, request)
	return helper.WriteResponse(c, response)
}
