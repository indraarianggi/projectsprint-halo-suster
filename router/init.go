package router

import (
	"github.com/backend-magang/halo-suster/user"
	"github.com/labstack/echo/v4"
)

type RouterHandler struct {
	UserHandler user.Handler
}

func InitRouter(server *echo.Echo, handlers RouterHandler) {
	InitUserRouter(server, handlers.UserHandler)
}
