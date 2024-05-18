package router

import (
	"github.com/backend-magang/halo-suster/internal/handler"
	"github.com/labstack/echo/v4"
)

func InitRouter(server *echo.Echo, handler handler.Handler) {
	InitUserRouter(server, handler)
	InitMedicalRouter(server, handler)
}
