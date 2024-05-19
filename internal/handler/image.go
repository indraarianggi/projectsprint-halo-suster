package handler

import (
	"net/http"

	"github.com/backend-magang/halo-suster/utils/helper"
	"github.com/labstack/echo/v4"
)

func (h *handler) UploadImage(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	response := h.usecase.UploadImage(ctx, fileHeader)
	return helper.WriteResponse(c, response)
}
