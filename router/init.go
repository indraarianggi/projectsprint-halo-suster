package router

import (
	"github.com/backend-magang/halo-suster/patient"
	"github.com/backend-magang/halo-suster/user"
	"github.com/labstack/echo/v4"
)

type RouterHandler struct {
	UserHandler    user.Handler
	PatientHandler patient.Handler
}

func InitRouter(server *echo.Echo, handlers RouterHandler) {
	InitUserRouter(server, handlers.UserHandler)
	InitMedicalRouter(server, handlers.PatientHandler)
}
