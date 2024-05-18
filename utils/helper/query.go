package helper

import (
	"fmt"
	"slices"
	"strings"

	"github.com/backend-magang/halo-suster/models/input"
	"github.com/backend-magang/halo-suster/utils/constant"
	"github.com/spf13/cast"
)

func BuildQueryGetListUser(req input.GetListUserRequest, fields ...string) (string, []interface{}) {
	var (
		queryBuilder   strings.Builder
		args           []interface{}
		selectedFields string
		orderByClause  string
	)

	if len(fields) == 0 {
		selectedFields = "*"
	} else {
		selectedFields = strings.Join(fields, ", ")
	}

	queryBuilder.WriteString("SELECT " + selectedFields + " FROM users WHERE 1=1")

	if req.ID != "" {
		queryBuilder.WriteString(" AND id = ?")
		args = append(args, req.ID)
	}

	if cast.ToInt(req.NIP) != 0 {
		queryBuilder.WriteString(" AND CAST(nip AS TEXT) LIKE ?")
		args = append(args, req.NIP+"%")
	}

	if req.Name != "" {
		queryBuilder.WriteString(" AND name ILIKE ?")
		args = append(args, "%"+req.Name+"%")
	}

	if req.Role != "" && slices.Contains([]string{constant.ROLE_IT, constant.ROLE_NURSE}, req.Role) {
		queryBuilder.WriteString(" AND role = ?")
		args = append(args, req.Role)
	}

	if cast.ToBool(req.IsDeleted) {
		queryBuilder.WriteString(" AND deleted_at IS NOT NULL")
	} else {
		queryBuilder.WriteString(" AND deleted_at IS NULL")
	}

	if req.CreatedAt != "" && slices.Contains([]string{"asc", "desc"}, req.CreatedAt) {
		if orderByClause != "" {
			orderByClause += fmt.Sprintf(", created_at %s", strings.ToUpper(req.CreatedAt))
		} else {
			orderByClause += fmt.Sprintf("created_at %s", strings.ToUpper(req.CreatedAt))
		}
	}

	if orderByClause != "" {
		queryBuilder.WriteString(" ORDER BY " + orderByClause)
	} else {
		queryBuilder.WriteString(" ORDER BY created_at DESC")
	}

	queryBuilder.WriteString(" LIMIT ? OFFSET ?")
	args = append(args, cast.ToInt(req.Limit), cast.ToInt(req.Offset))

	return queryBuilder.String(), args
}

func BuildQueryGetListPatient(req input.GetListPatientRequest, fields ...string) (string, []interface{}) {
	var (
		queryBuilder   strings.Builder
		args           []interface{}
		selectedFields string
		orderByClause  string
	)

	if len(fields) == 0 {
		selectedFields = "*"
	} else {
		selectedFields = strings.Join(fields, ", ")
	}

	queryBuilder.WriteString("SELECT " + selectedFields + " FROM patients WHERE 1=1")

	if cast.ToInt(req.IdentityNumber) != 0 {
		queryBuilder.WriteString(" AND CAST(identity_number AS TEXT) LIKE ?")
		args = append(args, req.IdentityNumber+"%")
	}

	if req.PhoneNumber != "" {
		queryBuilder.WriteString(" AND phone_number ILIKE ?")
		args = append(args, "+"+req.PhoneNumber+"%")
	}

	if req.Name != "" {
		queryBuilder.WriteString(" AND name ILIKE ?")
		args = append(args, "%"+req.Name+"%")
	}

	if req.CreatedAt != "" && slices.Contains([]string{"asc", "desc"}, req.CreatedAt) {
		if orderByClause != "" {
			orderByClause += fmt.Sprintf(", created_at %s", strings.ToUpper(req.CreatedAt))
		} else {
			orderByClause += fmt.Sprintf("created_at %s", strings.ToUpper(req.CreatedAt))
		}
	}

	if orderByClause != "" {
		queryBuilder.WriteString(" ORDER BY " + orderByClause)
	} else {
		queryBuilder.WriteString(" ORDER BY created_at DESC")
	}

	queryBuilder.WriteString(" LIMIT ? OFFSET ?")
	args = append(args, cast.ToInt(req.Limit), cast.ToInt(req.Offset))

	return queryBuilder.String(), args
}

func BuildQueryGetListMedicalRecord(req input.GetListMedicalRecordRequest, fields ...string) (string, []interface{}) {
	var (
		queryBuilder strings.Builder
		args         []interface{}
		// selectedFields string
		orderByClause string
	)

	// if len(fields) == 0 {
	// 	selectedFields = "*"
	// } else {
	// 	selectedFields = strings.Join(fields, ", ")
	// }

	// queryBuilder.WriteString("SELECT " + selectedFields + " FROM users WHERE 1=1")
	queryBuilder.WriteString(`
		SELECT
			medical_records.id AS medical_record_id,
			medical_records.symptoms,
			medical_records.medications,
			medical_records.created_at,
			identity_detail.identity_number,
			identity_detail.phone_number,
			identity_detail.name,
			identity_detail.birth_date,
			identity_detail.gender,
			identity_detail.identity_image_url,
			created_by.id AS created_by_id,
			created_by.nip AS created_by_nip,
			created_by.name AS created_by_name 
		FROM 
			medical_records
		JOIN 
			patients AS identity_detail ON medical_records.patient_id = identity_detail.id 
		JOIN 
			users AS created_by ON medical_records.created_by_id = created_by.id 
		WHERE
			1=1
	`)

	if cast.ToInt(req.IdentityNumber) != 0 {
		queryBuilder.WriteString(" AND CAST(identity_number AS TEXT) LIKE ?")
		args = append(args, req.IdentityNumber+"%")
	}

	if req.CreatedByID != "" {
		queryBuilder.WriteString(" AND created_by_id = ?")
		args = append(args, req.CreatedByID)
	}

	if cast.ToInt(req.CreatedByNIP) != 0 {
		queryBuilder.WriteString(" AND CAST(created_by_nip AS TEXT) LIKE ?")
		args = append(args, req.CreatedByNIP+"%")
	}

	if req.CreatedAt != "" && slices.Contains([]string{"asc", "desc"}, req.CreatedAt) {
		if orderByClause != "" {
			orderByClause += fmt.Sprintf(", created_at %s", strings.ToUpper(req.CreatedAt))
		} else {
			orderByClause += fmt.Sprintf("created_at %s", strings.ToUpper(req.CreatedAt))
		}
	}

	if orderByClause != "" {
		queryBuilder.WriteString(" ORDER BY " + orderByClause)
	} else {
		queryBuilder.WriteString(" ORDER BY created_at DESC")
	}

	queryBuilder.WriteString(" LIMIT ? OFFSET ?")
	args = append(args, cast.ToInt(req.Limit), cast.ToInt(req.Offset))

	return queryBuilder.String(), args
}
