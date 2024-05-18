package user

import (
	"fmt"
	"slices"
	"strings"

	"github.com/backend-magang/halo-suster/utils/constant"
	"github.com/spf13/cast"
)

func buildQueryGetListUser(req GetListUserRequest, fields ...string) (string, []interface{}) {
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
