package usecase

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/backend-magang/halo-suster/models/entity"
	"github.com/backend-magang/halo-suster/models/input"
	"github.com/backend-magang/halo-suster/utils/constant"
	"github.com/backend-magang/halo-suster/utils/helper"
)

func (u *usecase) AddMedicalRecord(ctx context.Context, request input.AddMedicalRecordRequest) helper.StandardResponse {
	var (
		patient          entity.Patient
		newMedicalRecord entity.MedicalRecord
		dataResponse     entity.MedicalRecord
		err              error
		now              = time.Now()
	)

	// get detail patient by identity number
	patient, err = u.repository.FindPatientByIdentityNumber(ctx, request.IdentityNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.StandardResponse{Code: http.StatusBadRequest, Message: constant.PATIENT_NOT_FOUND, Error: err}
		}
		return helper.StandardResponse{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	newMedicalRecord = entity.MedicalRecord{
		ID:             helper.NewULID(),
		PatientID:      patient.ID,
		IdentityNumber: patient.IdentityNumber,
		Symptoms:       request.Symptoms,
		Medications:    request.Medications,
		CreatedByID:    request.CreatedByID,
		CreatedByNIP:   request.CreatedByNIP,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	// save new medical record to database
	_, err = u.repository.SaveMedicalRecord(ctx, newMedicalRecord)
	if err != nil {
		return helper.StandardResponse{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	dataResponse = entity.MedicalRecord{
		ID:             newMedicalRecord.ID,
		IdentityNumber: newMedicalRecord.IdentityNumber,
		Symptoms:       newMedicalRecord.Symptoms,
		Medications:    newMedicalRecord.Medications,
		CreatedByID:    newMedicalRecord.CreatedByID,
		CreatedAt:      newMedicalRecord.CreatedAt,
	}

	return helper.StandardResponse{Code: http.StatusCreated, Message: constant.SUCCESS_ADD_PATIENT, Data: dataResponse}
}

func (u *usecase) GetListMedicalRecord(ctx context.Context, request input.GetListMedicalRecordRequest) helper.StandardResponse {
	var (
		medicalRecords []entity.MedicalRecordResponse
		err            error
	)

	result, err := u.repository.FindMedicalRecords(ctx, request)
	if err != nil {
		return helper.StandardResponse{Code: http.StatusInternalServerError, Message: constant.FAILED_GET_MEDICAL_RECORDS, Error: err}
	}

	for _, record := range result {
		medicalRecords = append(medicalRecords, entity.MedicalRecordResponse{
			ID: record.ID,
			IdentityDetail: entity.Patient{
				IdentityNumber:   record.PatientIdentityNumber,
				PhoneNumber:      record.PatientPhoneNumber,
				Name:             record.PatientName,
				BirthDate:        record.PatientBirthDate,
				Gender:           record.PatientGender,
				IdentityImageUrl: record.PatientIdentityImageUrl,
			},
			Symptoms:    record.Symptoms,
			Medications: record.Medications,
			CreatedBy: entity.User{
				ID:   record.CreatedByID,
				NIP:  record.CreatedByNIP,
				Name: record.CreatedByName,
			},
			CreatedAt: record.CreatedAt,
		})
	}

	return helper.StandardResponse{Code: http.StatusOK, Message: constant.SUCCESS, Data: medicalRecords}
}
