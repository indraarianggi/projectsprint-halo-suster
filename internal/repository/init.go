package repository

import (
	"context"

	"github.com/backend-magang/halo-suster/config"
	"github.com/backend-magang/halo-suster/models/entity"
	"github.com/backend-magang/halo-suster/models/input"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	// User
	SaveUser(ctx context.Context, user entity.User) (result entity.User, err error)
	FindUserByNIP(ctx context.Context, nip int64) (result entity.User, err error)
	FindUserByID(ctx context.Context, userId string) (result entity.User, err error)
	FindUsers(ctx context.Context, request input.GetListUserRequest) (result []entity.User, err error)
	UpdateUser(ctx context.Context, user entity.User) (result entity.User, err error)
	DeleteUser(ctx context.Context, userId string) (err error)

	// Patient
	SavePatient(ctx context.Context, user entity.Patient) (result entity.Patient, err error)
	FindPatientByIdentityNumber(ctx context.Context, identityNumber int64) (result entity.Patient, err error)
	FindPatients(ctx context.Context, request input.GetListPatientRequest) (result []entity.Patient, err error)

	// Medical Record
	SaveMedicalRecord(ctx context.Context, medicalRecord entity.MedicalRecord) (result entity.MedicalRecord, err error)
	FindMedicalRecords(ctx context.Context, request input.GetListMedicalRecordRequest) (result []entity.MedicalRecordResult, err error)
}

type repository struct {
	db     *sqlx.DB
	logger *logrus.Logger
	config config.Config
}

func NewRepository(db *sqlx.DB, config config.Config, logger *logrus.Logger) Repository {
	return &repository{
		db,
		logger,
		config,
	}
}
