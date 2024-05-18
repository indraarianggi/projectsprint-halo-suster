package patient

import (
	"net/http"

	"github.com/backend-magang/halo-suster/utils/helper"
	"github.com/backend-magang/halo-suster/utils/pkg"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

type Handler interface {
	AddPatient(echo.Context) error
	GetListPatient(echo.Context) error
}

type handler struct {
	usecase Usecase
	logger  *logrus.Logger
}

func NewHandler(usecase Usecase, logger *logrus.Logger) Handler {
	return &handler{usecase, logger}
}

func (h *handler) AddPatient(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	request := AddPatientRequest{}
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

	request := GetListPatientRequest{
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
