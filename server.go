package main

import (
	"backend/routes"
	"net/http"
	"github.com/labstack/echo/v4"
)

func healthz(c echo.Context) error {
	return c.String(http.StatusOK, "O pai ta on!")
}

func main() {
	e := echo.New()
	e.GET("/healthz", healthz)
	e.GET("/courses", routes.GetCourses)
	e.Start(":1323")
}
