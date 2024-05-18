package user

import (
	"context"
	"database/sql"
	"time"

	"github.com/backend-magang/halo-suster/config"
	"github.com/backend-magang/halo-suster/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	Save(ctx context.Context, user models.User) (result models.User, err error)
	FindUserByNIP(ctx context.Context, nip int64) (result models.User, err error)
	FindUserByID(ctx context.Context, userId string) (result models.User, err error)
	FindUser(ctx context.Context, request GetListUserRequest) (result []models.User, err error)
	UpdateUser(ctx context.Context, user models.User) (result models.User, err error)
	DeleteUser(ctx context.Context, userId string) (err error)
}

type repository struct {
	db     *sqlx.DB
	config config.Config
	logger *logrus.Logger
}

func NewRepository(db *sqlx.DB, config config.Config, logger *logrus.Logger) Repository {
	return &repository{db, config, logger}
}

func (r *repository) Save(ctx context.Context, user models.User) (result models.User, err error) {
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
		r.logger.Errorf("[Repository][User][Save] failed to insert new user, err: %s", err.Error())
		return
	}

	return
}

func (r *repository) FindUserByNIP(ctx context.Context, nip int64) (result models.User, err error) {
	query := `SELECT * FROM users 
		WHERE nip = $1 AND deleted_at IS NULL`

	err = r.db.QueryRowxContext(ctx, query, nip).StructScan(&result)
	if err != nil && err != sql.ErrNoRows {
		r.logger.Errorf("[Repository][User][FindUserByNIP] failed to find user by nip %d, err: %s", nip, err.Error())
		return
	}

	return
}

func (r *repository) FindUserByID(ctx context.Context, userId string) (result models.User, err error) {
	query := `SELECT * FROM users 
		WHERE id = $1 AND deleted_at IS NULL`

	err = r.db.QueryRowxContext(ctx, query, userId).StructScan(&result)
	if err != nil && err != sql.ErrNoRows {
		r.logger.Errorf("[Repository][User][FindUserByID] failed to find user by nip %s, err: %s", userId, err.Error())
		return
	}

	return
}

func (r *repository) FindUser(ctx context.Context, request GetListUserRequest) ([]models.User, error) {
	result := []models.User{}

	query, args := buildQueryGetListUser(request, "id", "nip", "name", "created_at")
	query = r.db.Rebind(query)

	err := r.db.SelectContext(ctx, &result, query, args...)
	if err != nil {
		r.logger.Errorf("[Repository][User][FindUser] failed to query, err: %s", err.Error())
		return result, err
	}

	return result, err
}

func (r *repository) UpdateUser(ctx context.Context, user models.User) (result models.User, err error) {
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
		result models.User
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
