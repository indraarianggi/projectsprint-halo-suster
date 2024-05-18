package patient

import (
	"fmt"
	"slices"
	"strings"

	"github.com/spf13/cast"
)

func buildQueryGetListPatient(req GetListPatientRequest, fields ...string) (string, []interface{}) {
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
