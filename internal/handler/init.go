package handler

import (
	"github.com/backend-magang/halo-suster/internal/usecase"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Handler interface {
	// User
	RegisterIT(echo.Context) error
	LoginIT(echo.Context) error
	RegisterNurse(echo.Context) error
	LoginNurse(echo.Context) error
	SetPasswordNurse(echo.Context) error
	UpdateNurse(echo.Context) error
	DeleteNurse(echo.Context) error
	GetListUser(echo.Context) error

	// Patient
	AddPatient(echo.Context) error
	GetListPatient(echo.Context) error

	// Medical Record
	AddMedicalRecord(echo.Context) error
	GetListMedicalRecord(echo.Context) error
}

type handler struct {
	usecase usecase.Usecase
	logger  *logrus.Logger
}

func NewHandler(usecase usecase.Usecase, logger *logrus.Logger) Handler {
	return &handler{usecase, logger}
}
