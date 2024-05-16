package main

import (
	"github.com/backend-magang/halo-suster/config"
	"github.com/backend-magang/halo-suster/driver"
	"github.com/backend-magang/halo-suster/middleware"
	"github.com/backend-magang/halo-suster/router"
	"github.com/backend-magang/halo-suster/user"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	server := echo.New()

	// Load config
	config := config.Load()
	logger := logrus.New()

	dbClient := driver.InitPostgres(config)

	middleware.InitMiddleware(server)

	userRepository := user.NewRepository(dbClient, config, logger)
	userUsecase := user.NewUsecase(userRepository, config, logger)
	userHandler := user.NewHandler(userUsecase, logger)

	router.InitRouter(server, router.RouterHandler{UserHandler: userHandler})

	server.Start(":8080")
}
