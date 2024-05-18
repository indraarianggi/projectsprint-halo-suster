package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/backend-magang/halo-suster/models/entity"
	"github.com/backend-magang/halo-suster/models/input"
	"github.com/backend-magang/halo-suster/utils/helper"
)

func (r *repository) SaveUser(ctx context.Context, user entity.User) (result entity.User, err error) {
	query := `INSERT INTO users (id, nip, name, role, password, identity_image_url, created_at, updated_at) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
        RETURNING *`

	err = r.db.QueryRowxContext(ctx,
		query,
		user.ID,
		user.NIP,
		user.Name,
		user.Role,
		user.Password,
		user.IdentityImageUrl,
		user.CreatedAt,
		user.UpdatedAt,
	).StructScan(&result)

	if err != nil {
		r.logger.Errorf("[Repository][User][SaveUser] failed to insert new user, err: %s", err.Error())
		return
	}

	return
}

func (r *repository) FindUserByNIP(ctx context.Context, nip int64) (result entity.User, err error) {
	query := `SELECT * FROM users 
		WHERE nip = $1 AND deleted_at IS NULL`

	err = r.db.QueryRowxContext(ctx, query, nip).StructScan(&result)
	if err != nil && err != sql.ErrNoRows {
		r.logger.Errorf("[Repository][User][FindUserByNIP] failed to find user by nip %d, err: %s", nip, err.Error())
		return
	}

	return
}

func (r *repository) FindUserByID(ctx context.Context, userId string) (result entity.User, err error) {
	query := `SELECT * FROM users 
		WHERE id = $1 AND deleted_at IS NULL`

	err = r.db.QueryRowxContext(ctx, query, userId).StructScan(&result)
	if err != nil && err != sql.ErrNoRows {
		r.logger.Errorf("[Repository][User][FindUserByID] failed to find user by nip %s, err: %s", userId, err.Error())
		return
	}

	return
}

func (r *repository) FindUsers(ctx context.Context, request input.GetListUserRequest) (result []entity.User, err error) {
	result = []entity.User{}

	query, args := helper.BuildQueryGetListUser(request, "id", "nip", "name", "created_at")
	query = r.db.Rebind(query)

	err = r.db.SelectContext(ctx, &result, query, args...)
	if err != nil {
		r.logger.Errorf("[Repository][User][FindUsers] failed to query, err: %s", err.Error())
		return result, err
	}

	return result, err
}

func (r *repository) UpdateUser(ctx context.Context, user entity.User) (result entity.User, err error) {
	query := `UPDATE users SET 
		nip = $1,
		name = $2,
		role = $3,
		password = $4,
		identity_image_url = $5,
		updated_at = $6 
	WHERE id = $7 RETURNING *`

	err = r.db.QueryRowxContext(ctx,
		query,
		user.NIP,
		user.Name,
		user.Role,
		user.Password,
		user.IdentityImageUrl,
		user.UpdatedAt,
		user.ID,
	).StructScan(&result)

	if err != nil {
		r.logger.Errorf("[Repository][User][UpdateUser] failed to update user, err: %s", err.Error())
		return
	}

	return
}

func (r *repository) DeleteUser(ctx context.Context, userId string) (err error) {
	var (
		now    = time.Now()
		result entity.User
		args   = []interface{}{now, sql.NullTime{Time: now, Valid: true}, userId}
	)

	query := `UPDATE users SET 
		updated_at = $1,
		deleted_at = $2 
	WHERE id = $3 AND deleted_at IS NULL RETURNING *`

	err = r.db.QueryRowxContext(ctx, query, args...).StructScan(&result)
	if err != nil {
		r.logger.Errorf("[Repository][Users][DeleteUser] failed to delete user, err: %s", err.Error())
		return
	}

	return
}
