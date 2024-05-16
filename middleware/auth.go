package middleware

import (
	"log"
	"time"

	"github.com/backend-magang/halo-suster/config"
	"github.com/backend-magang/halo-suster/models"
	"github.com/golang-jwt/jwt/v4"
)

func getJWTSecretKey() string {
	return config.Load().JWTSecret
}

func GenerateToken(user models.User) (t string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["nip"] = user.NIP
	claims["name"] = user.Name
	claims["role"] = user.Role
	claims["expired_at"] = time.Now().Add(8 * time.Hour)

	t, err = token.SignedString([]byte(getJWTSecretKey()))
	if err != nil {
		log.Println("[Middleware] failed to signed jwt token, err: ", err)
		return
	}

	return
}
