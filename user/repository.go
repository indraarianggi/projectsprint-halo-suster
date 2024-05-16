package user

import (
	"context"

	"github.com/backend-magang/halo-suster/config"
	"github.com/backend-magang/halo-suster/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	Save(context.Context, models.User) (models.User, error)
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
