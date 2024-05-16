package helper

import (
	"net/http"

	"github.com/backend-magang/halo-suster/utils/constant"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type StandardResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   error       `json:"error"`
}

type HttpResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error,omitempty"`
}

func WriteResponse(c echo.Context, response StandardResponse) error {
	status := constant.SUCCESS

	if response.Code > 299 {
		status = constant.FAILED
	}

	var errorResponse interface{} = nil
	if response.Error != nil {
		errorResponse = response.Error.Error()
	}

	if response.Message == "" {
		response.Message = http.StatusText(response.Code)
	}

	return c.JSON(response.Code, HttpResponse{
		Code:    response.Code,
		Status:  status,
		Message: response.Message,
		Data:    response.Data,
		Error:   errorResponse,
	})
}

func FormatValidationError(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
