package router

import (
	"github.com/backend-magang/halo-suster/middleware"
	"github.com/backend-magang/halo-suster/user"
	"github.com/backend-magang/halo-suster/utils/constant"
	"github.com/labstack/echo/v4"
)

func InitUserRouter(e *echo.Echo, handler user.Handler) {
	v1 := e.Group("/v1")
	user := v1.Group("/user")
	it := user.Group("/it")
	nurse := user.Group("/nurse")

	it.POST("/register", handler.RegisterIT)
	it.POST("/login", handler.LoginIT)

	nurse.POST("/login", handler.LoginNurse)
	nurse.POST("/register", handler.RegisterNurse, middleware.TokenValidation(constant.ROLE_IT))
	nurse.POST("/:id/access", handler.SetPasswordNurse, middleware.TokenValidation(constant.ROLE_IT))
}
