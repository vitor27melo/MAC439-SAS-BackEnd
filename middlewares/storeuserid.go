package middlewares

import (
	"backend/configs"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func StoreUserId(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("userId").(*jwt.Token)
		claims := userId.Claims.(*configs.JWTCustomClaims)
		fmt.Println(claims.UserId)

		return next(c)
	}
}
