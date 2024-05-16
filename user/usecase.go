package user

import (
	"context"
	"database/sql"
	"net/http"
	"strings"
	"time"

	"github.com/backend-magang/halo-suster/config"
	"github.com/backend-magang/halo-suster/middleware"
	"github.com/backend-magang/halo-suster/models"
	"github.com/backend-magang/halo-suster/utils/constant"
	"github.com/backend-magang/halo-suster/utils/helper"
	"github.com/backend-magang/halo-suster/utils/lib"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

type Usecase interface {
	RegisterIT(context.Context, RegisterITRequest) helper.StandardResponse
}

type usecase struct {
	repository Repository
	config     config.Config
	logger     *logrus.Logger
}

func NewUsecase(repository Repository, config config.Config, logger *logrus.Logger) Usecase {
	return &usecase{repository, config, logger}
}

func (u *usecase) RegisterIT(ctx context.Context, request RegisterITRequest) helper.StandardResponse {
	var (
		user    = models.User{}
		newUser = models.User{}
		now     = time.Now()
	)

	hashedPassword := helper.HashPassword(request.Password, cast.ToInt(u.config.BcryptSalt))
	newUser = models.User{
		ID:       helper.NewULID(),
		NIP:      request.NIP,
		Name:     request.Name,
		Role:     constant.ROLE_IT,
		Password: sql.NullString{String: hashedPassword, Valid: true},
		// IdentityImageUrl: "",
		CreatedAt: now,
		UpdatedAt: now,
	}

	user, err := u.repository.Save(ctx, newUser)
	if err != nil {
		if strings.Contains(err.Error(), lib.ErrConstraintKey.Error()) {
			return helper.StandardResponse{Code: http.StatusConflict, Message: constant.DUPLICATE_NIP, Error: err}
		}
		return helper.StandardResponse{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	token, _ := middleware.GenerateToken(user)
	userAsResponse := models.UserWithToken{
		ID:    user.ID,
		NIP:   user.NIP,
		Name:  user.Name,
		Token: token,
	}

	return helper.StandardResponse{Code: http.StatusCreated, Message: constant.SUCCESS_REGISTER_USER, Data: userAsResponse}
}
