package usecase

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/backend-magang/halo-suster/models/entity"
	"github.com/backend-magang/halo-suster/models/input"
	"github.com/backend-magang/halo-suster/utils/constant"
	"github.com/backend-magang/halo-suster/utils/helper"
	"github.com/backend-magang/halo-suster/utils/lib"
)

func (u *usecase) AddPatient(ctx context.Context, request input.AddPatientRequest) helper.StandardResponse {
	var (
		newPatient   entity.Patient
		patient      entity.Patient
		dataResponse entity.Patient
		err          error
		now          = time.Now()
	)

	birtDate, err := time.Parse(time.RFC3339, request.BirthDate)
	if err != nil {
		return helper.StandardResponse{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	newPatient = entity.Patient{
		ID:               helper.NewULID(),
		IdentityNumber:   request.IdentityNumber,
		Name:             request.Name,
		PhoneNumber:      request.PhoneNumber,
		BirthDate:        birtDate,
		Gender:           request.Gender,
		IdentityImageUrl: request.IdentityImageUrl,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	// save new patient to database
	patient, err = u.repository.SavePatient(ctx, newPatient)
	if err != nil {
		if strings.Contains(err.Error(), lib.ErrConstraintKey.Error()) {
			return helper.StandardResponse{Code: http.StatusConflict, Message: constant.DUPLICATE_IDENTITY_NUMBER, Error: err}
		}
		return helper.StandardResponse{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	dataResponse = entity.Patient{
		ID:               patient.ID,
		IdentityNumber:   patient.IdentityNumber,
		PhoneNumber:      patient.PhoneNumber,
		Name:             patient.Name,
		BirthDate:        patient.BirthDate,
		Gender:           patient.Gender,
		IdentityImageUrl: patient.IdentityImageUrl,
		CreatedAt:        patient.CreatedAt,
	}

	return helper.StandardResponse{Code: http.StatusCreated, Message: constant.SUCCESS_ADD_PATIENT, Data: dataResponse}
}

func (u *usecase) GetListPatient(ctx context.Context, request input.GetListPatientRequest) helper.StandardResponse {
	var (
		patients []entity.Patient
		err      error
	)

	patients, err = u.repository.FindPatients(ctx, request)
	if err != nil {
		return helper.StandardResponse{Code: http.StatusInternalServerError, Message: constant.FAILED_GET_PATIENTS, Error: err}
	}

	return helper.StandardResponse{Code: http.StatusOK, Message: constant.SUCCESS, Data: patients}
}
