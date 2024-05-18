package main

import (
	"github.com/backend-magang/halo-suster/config"
	"github.com/backend-magang/halo-suster/driver"
	"github.com/backend-magang/halo-suster/internal/handler"
	"github.com/backend-magang/halo-suster/internal/repository"
	"github.com/backend-magang/halo-suster/internal/usecase"
	"github.com/backend-magang/halo-suster/middleware"
	"github.com/backend-magang/halo-suster/router"
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

	repository := repository.NewRepository(dbClient, config, logger)
	usecase := usecase.NewUsecase(repository, config, logger)
	handler := handler.NewHandler(usecase, logger)

	router.InitRouter(server, handler)

	server.Start(":8080")
}
