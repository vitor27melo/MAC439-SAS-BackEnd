package main

import (
	"backend/configs"
	"backend/middlewares"
	"backend/routes"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var config = middleware.JWTConfig{
	Claims:     &configs.JWTCustomClaims{},
	SigningKey: configs.JWTSecret,
}

func healthz(c echo.Context) error {
	return c.String(http.StatusOK, "It's alive!!!\nhttps://www.youtube.com/watch?v=xos2MnVxe-c")
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.GET("/healthz", healthz)
	e.POST("/login", routes.Login)
	e.POST("/register", routes.Register)

	userGroup := e.Group("/user")
	{
		userGroup.Use(middleware.JWTWithConfig(config))
		userGroup.Use(middlewares.StoreUserId)
		userGroup.GET("/courses", routes.GetCourses)
		userGroup.GET("/days", routes.GetDays)
		userGroup.GET("/list", routes.GetUsers)
		userGroup.GET("/course-schedule", routes.GetTimeSchedule)
		userGroup.GET("/files-list", routes.GetUserFiles)
		userGroup.POST("/upload-file", routes.UploadFile)
		userGroup.GET("/download-file/:name", routes.DownloadFile)
		userGroup.GET("/risk", routes.CalculateRisk)
		userGroup.POST("/report-covid", routes.ReportCovid)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}
	e.Start(":" + port)
}
