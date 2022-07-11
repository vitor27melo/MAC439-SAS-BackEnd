package routes

import (
	"backend/configs"
	"backend/model"
	"backend/assets"
	"database/sql"
	"log"
	"net/http"
	"github.com/labstack/echo/v4"
)

func GetCourses(c echo.Context) error {
	stmt := `
		SELECT
			id_disciplina,
			nome,
			sigla
		FROM
			disciplina;
	`

	courses := []model.Course{}

	db, err := sql.Open(configs.GetDBType(), configs.GetPostgresConnString())
	assets.CheckError(err)

	rows, e := db.Query(stmt)
	assets.CheckError(e)

	for rows.Next() {
		var course model.Course

		if err := rows.Scan(&course.Id_disciplina, &course.Nome, &course.Sigla); err != nil {
			log.Fatal(err)
		}
		courses = append(courses, course)
	}

	return c.JSON(http.StatusOK, courses)
}