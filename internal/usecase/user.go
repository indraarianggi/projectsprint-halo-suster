package usecase

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/backend-magang/halo-suster/middleware"
	"github.com/backend-magang/halo-suster/models/entity"
	"github.com/backend-magang/halo-suster/models/input"
	"github.com/backend-magang/halo-suster/utils/constant"
	"github.com/backend-magang/halo-suster/utils/helper"
	"github.com/backend-magang/halo-suster/utils/lib"
	"github.com/spf13/cast"
	"golang.org/x/crypto/bcrypt"
)

func (u *usecase) RegisterIT(ctx context.Context, request input.RegisterITRequest) helper.StandardResponse {
	var (
		newUser      entity.User
		user         entity.User
		dataResponse entity.UserWithToken
		token        string
		err          error
		now          = time.Now()
	)

	// generate hashed password
	hashedPassword := helper.HashPassword(request.Password, cast.ToInt(u.config.BcryptSalt))
	newUser = entity.User{
		ID:       helper.NewULID(),
		NIP:      request.NIP,
		Name:     request.Name,
		Role:     constant.ROLE_IT,
		Password: sql.NullString{String: hashedPassword, Valid: true},
		// IdentityImageUrl: "",
		CreatedAt: now,
		UpdatedAt: now,
	}

	// save new user to database
	user, err = u.repository.SaveUser(ctx, newUser)
	if err != nil {
		if strings.Contains(err.Error(), lib.ErrConstraintKey.Error()) {
			return helper.StandardResponse{Code: http.StatusConflict, Message: constant.DUPLICATE_NIP, Error: err}
		}
		return helper.StandardResponse{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	// generate token
	token, _ = middleware.GenerateToken(user)
	dataResponse = entity.UserWithToken{
		ID:    user.ID,
		NIP:   user.NIP,
		Name:  user.Name,
		Token: token,
	}

	return helper.StandardResponse{Code: http.StatusCreated, Message: constant.SUCCESS_REGISTER_USER, Data: dataResponse}
}

func (u *usecase) LoginIT(ctx context.Context, request input.LoginRequest) helper.StandardResponse {
	var (
		user         entity.User
		dataResponse entity.UserWithToken
		token        string
		err          error
	)

	// find user by nip
	user, err = u.repository.FindUserByNIP(ctx, request.NIP)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.StandardResponse{Code: http.StatusNotFound, Message: constant.USER_NOT_FOUND, Error: err}
		}
		return helper.StandardResponse{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	// check user role code, must be 615 (it)
	userNIPStr := strconv.FormatInt(user.NIP, 10)
	userRoleCode, err := strconv.Atoi(userNIPStr[:3])
	if err != nil {
		return helper.StandardResponse{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	} else if userRoleCode != constant.ROLE_CODE_IT {
		return helper.StandardResponse{Code: http.StatusNotFound, Message: constant.USER_NOT_FOUND}
	}

	// check user password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(request.Password))
	if err != nil {
		return helper.StandardResponse{Code: http.StatusBadRequest, Message: constant.FAILED_LOGIN}
	}

	// generate token
	token, _ = middleware.GenerateToken(user)
	dataResponse = entity.UserWithToken{
		ID:    user.ID,
		NIP:   user.NIP,
		Name:  user.Name,
		Token: token,
	}

	return helper.StandardResponse{Code: http.StatusOK, Message: constant.SUCCESS_LOGIN, Data: dataResponse}
}

func (u *usecase) LoginNurse(ctx context.Context, request input.LoginRequest) helper.StandardResponse {
	var (
		user         entity.User
		dataResponse entity.UserWithToken
		token        string
		err          error
	)

	// find user by nip
	user, err = u.repository.FindUserByNIP(ctx, request.NIP)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.StandardResponse{Code: http.StatusNotFound, Message: constant.USER_NOT_FOUND, Error: err}
		}
		return helper.StandardResponse{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	// check user role code, must be 303 (nurse)
	userNIPStr := strconv.FormatInt(user.NIP, 10)
	userRoleCode, err := strconv.Atoi(userNIPStr[:3])
	if err != nil {
		return helper.StandardResponse{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	} else if userRoleCode != constant.ROLE_CODE_NURSE {
		return helper.StandardResponse{Code: http.StatusNotFound, Message: constant.USER_NOT_FOUND}
	}

	// check user password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(request.Password))
	if err != nil {
		return helper.StandardResponse{Code: http.StatusBadRequest, Message: constant.FAILED_LOGIN}
	}

	// generate token
	token, _ = middleware.GenerateToken(user)
	dataResponse = entity.UserWithToken{
		ID:    user.ID,
		NIP:   user.NIP,
		Name:  user.Name,
		Token: token,
	}

	return helper.StandardResponse{Code: http.StatusOK, Message: constant.SUCCESS_LOGIN, Data: dataResponse}
}

func (u *usecase) RegisterNurse(ctx context.Context, request input.RegisterNurseRequest) helper.StandardResponse {
	var (
		newUser      entity.User
		user         entity.User
		dataResponse entity.User
		err          error
		now          = time.Now()
	)

	newUser = entity.User{
		ID:               helper.NewULID(),
		NIP:              request.NIP,
		Name:             request.Name,
		Role:             constant.ROLE_NURSE,
		IdentityImageUrl: request.IdentityImageUrl,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	// save new user to database
	user, err = u.repository.SaveUser(ctx, newUser)
	if err != nil {
		if strings.Contains(err.Error(), lib.ErrConstraintKey.Error()) {
			return helper.StandardResponse{Code: http.StatusConflict, Message: constant.DUPLICATE_NIP, Error: err}
		}
		return helper.StandardResponse{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	dataResponse = entity.User{
		ID:        user.ID,
		NIP:       user.NIP,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}

	return helper.StandardResponse{Code: http.StatusCreated, Message: constant.SUCCESS_REGISTER_USER, Data: dataResponse}
}

func (u *usecase) SetPasswordNurse(ctx context.Context, request input.NurseAccessRequest) helper.StandardResponse {
	// find user by id
	user, err := u.repository.FindUserByID(ctx, request.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.StandardResponse{Code: http.StatusNotFound, Message: constant.USER_NOT_FOUND, Error: err}
		}
		return helper.StandardResponse{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	// check user role code, must be 303 (nurse)
	userNIPStr := strconv.FormatInt(user.NIP, 10)
	userRoleCode, err := strconv.Atoi(userNIPStr[:3])
	if err != nil {
		return helper.StandardResponse{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	} else if userRoleCode != constant.ROLE_CODE_NURSE {
		return helper.StandardResponse{Code: http.StatusNotFound, Message: constant.USER_NOT_FOUND}
	}

	// generate hashed password
	hashedPassword := helper.HashPassword(request.Password, cast.ToInt(u.config.BcryptSalt))
	user.Password = sql.NullString{String: hashedPassword, Valid: true}
	user.UpdatedAt = time.Now()

	// update user data in database
	_, err = u.repository.UpdateUser(ctx, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.StandardResponse{Code: http.StatusNotFound, Message: constant.USER_NOT_FOUND, Error: err}
		}
		return helper.StandardResponse{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	return helper.StandardResponse{Code: http.StatusOK, Message: constant.SUCCESS, Data: nil}
}

func (u *usecase) UpdateNurse(ctx context.Context, request input.UpdateNurseRequest) helper.StandardResponse {
	var (
		updatedUser  entity.User
		user         entity.User
		dataResponse entity.User
		err          error
		now          = time.Now()
	)

	// find user by id
	user, err = u.repository.FindUserByID(ctx, request.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.StandardResponse{Code: http.StatusNotFound, Message: constant.USER_NOT_FOUND, Error: err}
		}
		return helper.StandardResponse{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	// check user role code, must be 303 (nurse)
	userNIPStr := strconv.FormatInt(user.NIP, 10)
	userRoleCode, err := strconv.Atoi(userNIPStr[:3])
	if err != nil {
		return helper.StandardResponse{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	} else if userRoleCode != constant.ROLE_CODE_NURSE {
		return helper.StandardResponse{Code: http.StatusNotFound, Message: constant.USER_NOT_FOUND}
	}

	updatedUser = entity.User{
		ID:               request.ID,
		NIP:              request.NIP,
		Name:             request.Name,
		Role:             user.Role,
		Password:         user.Password,
		IdentityImageUrl: user.IdentityImageUrl,
		UpdatedAt:        now,
	}

	// update user data in database
	user, err = u.repository.UpdateUser(ctx, updatedUser)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.StandardResponse{Code: http.StatusNotFound, Message: constant.USER_NOT_FOUND, Error: err}
		} else if strings.Contains(err.Error(), lib.ErrConstraintKey.Error()) {
			return helper.StandardResponse{Code: http.StatusConflict, Message: constant.DUPLICATE_NIP, Error: err}
		}
		return helper.StandardResponse{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	dataResponse = entity.User{
		ID:        user.ID,
		NIP:       user.NIP,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}

	return helper.StandardResponse{Code: http.StatusOK, Message: constant.SUCCESS_UPDATE_USER, Data: dataResponse}
}

func (u *usecase) DeleteNurse(ctx context.Context, request input.DeleteNurseRequest) helper.StandardResponse {
	// find user by id
	user, err := u.repository.FindUserByID(ctx, request.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.StandardResponse{Code: http.StatusNotFound, Message: constant.USER_NOT_FOUND, Error: err}
		}
		return helper.StandardResponse{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	// check user role code, must be 303 (nurse)
	userNIPStr := strconv.FormatInt(user.NIP, 10)
	userRoleCode, err := strconv.Atoi(userNIPStr[:3])
	if err != nil {
		return helper.StandardResponse{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	} else if userRoleCode != constant.ROLE_CODE_NURSE {
		return helper.StandardResponse{Code: http.StatusNotFound, Message: constant.USER_NOT_FOUND}
	}

	// delete user
	err = u.repository.DeleteUser(ctx, request.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.StandardResponse{Code: http.StatusNotFound, Message: constant.USER_NOT_FOUND, Error: err}
		}
		return helper.StandardResponse{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	return helper.StandardResponse{Code: http.StatusOK, Message: constant.SUCCESS, Data: nil}
}

func (u *usecase) GetListUser(ctx context.Context, request input.GetListUserRequest) helper.StandardResponse {
	var (
		users []entity.User
		err   error
	)

	users, err = u.repository.FindUsers(ctx, request)
	if err != nil {
		return helper.StandardResponse{Code: http.StatusInternalServerError, Message: constant.FAILED_GET_USERS, Error: err}
	}

	return helper.StandardResponse{Code: http.StatusOK, Message: constant.SUCCESS, Data: users}
}
