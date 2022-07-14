package configs

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
)

type JWTCustomClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

var JWTSecret = []byte("Atila Iamarino")

var config = middleware.JWTConfig{
	Claims:     &JWTCustomClaims{},
	SigningKey: JWTSecret,
}

func CreateJWT(userId string) (string, error) {
	claims := &JWTCustomClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 10).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(JWTSecret)
}
