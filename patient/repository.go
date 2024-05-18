package patient

import (
	"context"
	"database/sql"

	"github.com/backend-magang/halo-suster/config"
	"github.com/backend-magang/halo-suster/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	Save(ctx context.Context, user models.Patient) (result models.Patient, err error)
	FindPatientByIdentityNumber(ctx context.Context, identityNumber int64) (result models.Patient, err error)
	FindPatients(ctx context.Context, request GetListPatientRequest) (result []models.Patient, err error)
}

type repository struct {
	db     *sqlx.DB
	config config.Config
	logger *logrus.Logger
}

func NewRepository(db *sqlx.DB, config config.Config, logger *logrus.Logger) Repository {
	return &repository{db, config, logger}
}

func (r *repository) Save(ctx context.Context, user models.Patient) (result models.Patient, err error) {
	query := `INSERT INTO patients (id, identity_number, name, phone_number, birth_date, gender, identity_image_url, created_at, updated_at) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) 
        RETURNING *`

	err = r.db.QueryRowxContext(ctx,
		query,
		user.ID,
		user.IdentityNumber,
		user.Name,
		user.PhoneNumber,
		user.BirthDate,
		user.Gender,
		user.IdentityImageUrl,
		user.CreatedAt,
		user.UpdatedAt,
	).StructScan(&result)

	if err != nil {
		r.logger.Errorf("[Repository][Patient][Save] failed to insert new patient, err: %s", err.Error())
		return
	}

	return
}

func (r *repository) FindPatientByIdentityNumber(ctx context.Context, identityNumber int64) (result models.Patient, err error) {
	query := `SELECT * FROM patients 
		WHERE identity_number = $1`

	err = r.db.QueryRowxContext(ctx, query, identityNumber).StructScan(&result)
	if err != nil && err != sql.ErrNoRows {
		r.logger.Errorf("[Repository][Patient][FindPatientByIdentityNumber] failed to find patient by identity number %d, err: %s", identityNumber, err.Error())
		return
	}

	return
}

func (r *repository) FindPatients(ctx context.Context, request GetListPatientRequest) (result []models.Patient, err error) {
	result = []models.Patient{}

	query, args := buildQueryGetListPatient(request, "id", "identity_number", "phone_number", "name", "birth_date", "gender", "created_at")
	query = r.db.Rebind(query)

	err = r.db.SelectContext(ctx, &result, query, args...)
	if err != nil {
		r.logger.Errorf("[Repository][Patient][FindPatients] failed to query, err: %s", err.Error())
		return result, err
	}

	return result, err
}
