package router

import (
	"github.com/backend-magang/halo-suster/internal/handler"
	"github.com/backend-magang/halo-suster/middleware"
	"github.com/labstack/echo/v4"
)

func InitMedicalRouter(e *echo.Echo, handler handler.Handler) {
	v1 := e.Group("/v1")
	medical := v1.Group("/medical")
	patient := medical.Group("/patient", middleware.TokenValidation())
	record := medical.Group("/record", middleware.TokenValidation())

	patient.POST("", handler.AddPatient)
	patient.GET("", handler.GetListPatient)

	record.POST("", handler.AddMedicalRecord)
	record.GET("", handler.GetListMedicalRecord)
}
