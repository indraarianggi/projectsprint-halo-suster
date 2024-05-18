package repository

import (
	"context"

	"github.com/backend-magang/halo-suster/models/entity"
	"github.com/backend-magang/halo-suster/models/input"
	"github.com/backend-magang/halo-suster/utils/helper"
)

func (r *repository) SaveMedicalRecord(ctx context.Context, medicalRecord entity.MedicalRecord) (result entity.MedicalRecord, err error) {
	query := `INSERT INTO medical_records (id, patient_id, identity_number, created_by_id, created_by_nip, symptoms, medications, created_at, updated_at) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) 
        RETURNING *`

	err = r.db.QueryRowxContext(ctx,
		query,
		medicalRecord.ID,
		medicalRecord.PatientID,
		medicalRecord.IdentityNumber,
		medicalRecord.CreatedByID,
		medicalRecord.CreatedByNIP,
		medicalRecord.Symptoms,
		medicalRecord.Medications,
		medicalRecord.CreatedAt,
		medicalRecord.UpdatedAt,
	).StructScan(&result)

	if err != nil {
		r.logger.Errorf("[Repository][MedicalRecord][SaveMedicalRecord] failed to insert new medical record, err: %s", err.Error())
		return
	}

	return
}

func (r *repository) FindMedicalRecords(ctx context.Context, request input.GetListMedicalRecordRequest) (result []entity.MedicalRecordResponse, err error) {
	result = []entity.MedicalRecordResponse{}

	query, args := helper.BuildQueryGetListMedicalRecord(request)
	query = r.db.Rebind(query)

	err = r.db.SelectContext(ctx, &result, query, args...)
	if err != nil {
		r.logger.Errorf("[Repository][MedicalRecord][FindMedicalRecords] failed to query, err: %s", err.Error())
		return result, err
	}

	return result, err
}
