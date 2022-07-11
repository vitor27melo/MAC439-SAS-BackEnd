package main

import (
	"backend/routes"
	"github.com/labstack/echo/v4"
	"net/http"
)

func healthz(c echo.Context) error {
	return c.String(http.StatusOK, "O pai ta off!")
}

func main() {
	e := echo.New()
	e.GET("/healthz", healthz)
	e.GET("/courses", routes.GetCourses)
	e.Start(":1323")

}
