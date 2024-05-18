package usecase

import (
	"context"

	"github.com/backend-magang/halo-suster/config"
	"github.com/backend-magang/halo-suster/internal/repository"
	"github.com/backend-magang/halo-suster/models/input"
	"github.com/backend-magang/halo-suster/utils/helper"
	"github.com/sirupsen/logrus"
)

type Usecase interface {
	// User
	RegisterIT(context.Context, input.RegisterITRequest) helper.StandardResponse
	LoginIT(context.Context, input.LoginRequest) helper.StandardResponse
	RegisterNurse(context.Context, input.RegisterNurseRequest) helper.StandardResponse
	LoginNurse(context.Context, input.LoginRequest) helper.StandardResponse
	SetPasswordNurse(context.Context, input.NurseAccessRequest) helper.StandardResponse
	UpdateNurse(context.Context, input.UpdateNurseRequest) helper.StandardResponse
	DeleteNurse(context.Context, input.DeleteNurseRequest) helper.StandardResponse
	GetListUser(context.Context, input.GetListUserRequest) helper.StandardResponse

	// Patient
	AddPatient(context.Context, input.AddPatientRequest) helper.StandardResponse
	GetListPatient(context.Context, input.GetListPatientRequest) helper.StandardResponse

	// Medical Record
	AddMedicalRecord(context.Context, input.AddMedicalRecordRequest) helper.StandardResponse
	GetListMedicalRecord(context.Context, input.GetListMedicalRecordRequest) helper.StandardResponse
}

type usecase struct {
	repository repository.Repository
	config     config.Config
	logger     *logrus.Logger
}

func NewUsecase(repository repository.Repository, config config.Config, logger *logrus.Logger) Usecase {
	return &usecase{repository, config, logger}
}
