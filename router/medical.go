package router

import (
	"github.com/backend-magang/halo-suster/middleware"
	"github.com/backend-magang/halo-suster/patient"
	"github.com/labstack/echo/v4"
)

func InitMedicalRouter(e *echo.Echo, handler patient.Handler) {
	v1 := e.Group("/v1")
	medical := v1.Group("/medical")
	patient := medical.Group("/patient", middleware.TokenValidation())

	patient.POST("", handler.AddPatient)
	patient.GET("", handler.GetListPatient)
}
