package middlewares

import (
	"backend/configs"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func StoreUserId(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*configs.JWTCustomClaims)
		userId := claims.UserId
		c.Set("userId", userId)
		return next(c)
	}
}
