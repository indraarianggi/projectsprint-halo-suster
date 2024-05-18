package router

import (
	"github.com/backend-magang/halo-suster/internal/handler"
	"github.com/backend-magang/halo-suster/middleware"
	"github.com/labstack/echo/v4"
)

func InitUploadImageRouter(e *echo.Echo, handler handler.Handler) {
	v1 := e.Group("/v1")
	image := v1.Group("/image", middleware.TokenValidation())

	image.POST("", handler.UploadImage)
}
