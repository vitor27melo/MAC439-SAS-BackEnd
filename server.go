package main

import (
	"backend/routes"
	"net/http"

	"github.com/labstack/echo/v4"
)

func healthz(c echo.Context) error {
	return c.String(http.StatusOK, "O pai ta off!")
}

func main() {
	e := echo.New()
	e.GET("/healthz", healthz)
	e.GET("/courses", routes.GetCourses)
	e.GET("/days", routes.GetDays)
	e.GET("/users", routes.GetUsers)
	e.Start(":1323")
}
