package middleware

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/backend-magang/halo-suster/config"
	"github.com/backend-magang/halo-suster/models/entity"
	"github.com/backend-magang/halo-suster/utils/constant"
	"github.com/backend-magang/halo-suster/utils/helper"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

const userClaimKey = "authUser"

func getJWTSecretKey() string {
	return config.Load().JWTSecret
}

func GenerateToken(user entity.User) (t string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["nip"] = user.NIP
	claims["name"] = user.Name
	claims["role"] = user.Role
	claims["expiredAt"] = time.Now().Add(8 * time.Hour)

	t, err = token.SignedString([]byte(getJWTSecretKey()))
	if err != nil {
		log.Println("[Middleware] failed to signed jwt token, err: ", err)
		return
	}

	return
}

func IsTokenExpired(t time.Time) bool {
	now := time.Now()
	return t.Before(now)
}

func TokenValidation(args ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var authorizedRole string
			if len(args) != 0 {
				authorizedRole = args[0]
			}

			var userClaims = entity.UserClaims{}
			requestAuthHeader := c.Request().Header.Get("Authorization")

			if !strings.Contains(requestAuthHeader, "Bearer") {
				return helper.WriteResponse(c, helper.StandardResponse{
					Code:    http.StatusUnauthorized,
					Message: constant.INVALID_TOKEN,
				})
			}

			requestToken := strings.Split(requestAuthHeader, " ")[1]

			token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("invalid token signing method")
				}
				return []byte(getJWTSecretKey()), nil
			})

			if err != nil {
				return helper.WriteResponse(c, helper.StandardResponse{
					Code:    http.StatusUnauthorized,
					Message: constant.INVALID_SIGNING_METHOD,
					Error:   err,
				})
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return helper.WriteResponse(c, helper.StandardResponse{
					Code:    http.StatusUnauthorized,
					Message: constant.INVALID_TOKEN_CLAIMS,
					Error:   errors.New(constant.INVALID_TOKEN_CLAIMS),
				})
			}

			userClaims.ID = cast.ToString(claims["id"])
			userClaims.NIP = cast.ToInt64(claims["nip"])
			userClaims.Name = cast.ToString(claims["name"])
			userClaims.Role = cast.ToString(claims["role"])
			expiredAtStr := cast.ToString(claims["expiredAt"])
			userClaims.ExpiredAt, _ = time.Parse(time.RFC3339, expiredAtStr)

			if IsTokenExpired(userClaims.ExpiredAt) {
				return helper.WriteResponse(c, helper.StandardResponse{
					Code:    http.StatusUnauthorized,
					Message: constant.TOKEN_EXPIRED,
					Error:   errors.New(constant.TOKEN_EXPIRED),
				})
			}

			if authorizedRole != "" && userClaims.Role != authorizedRole {
				return helper.WriteResponse(c, helper.StandardResponse{
					Code:    http.StatusUnauthorized,
					Message: "unauthorized, you are not " + authorizedRole,
					Error:   errors.New("unauthorized, you are not " + authorizedRole),
				})
			}

			c.Set(userClaimKey, userClaims)
			return next(c)
		}
	}
}

func GetUserClaims(ctx echo.Context) entity.UserClaims {
	return ctx.Get(userClaimKey).(entity.UserClaims)
}
