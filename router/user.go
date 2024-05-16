package router

import (
	"github.com/backend-magang/halo-suster/user"
	"github.com/labstack/echo/v4"
)

func InitUserRouter(e *echo.Echo, handler user.Handler) {
	v1 := e.Group("/v1")
	user := v1.Group("/user")
	it := user.Group("/it")

	it.POST("/register", handler.RegisterIT)
}
