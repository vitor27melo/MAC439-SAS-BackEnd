package routes

import (
	"backend/assets"
	"backend/configs"
	"backend/model"
	"database/sql"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
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
