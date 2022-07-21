package routes

import (
	"backend/configs"
	"backend/model"
	"backend/tools"
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTimeSchedule(c echo.Context) error {
	stmt := `
		select
		    id_disciplina,
			sigla,
			dia_semana,
			d.nome,
			inicio,
		    fim
		from
			usuario u
		inner join
			usuario_oferecimento_recorrente ofr on u.id_usuario = ofr.fk_usuario
		inner join
			oferecimento_recorrente o on ofr.fk_oferecimento_recorrente = o.id_oferecimento_recorrente
		inner join
			disciplina d on o.fk_disciplina = d.id_disciplina
		where id_usuario = $1
	`

	courses := []model.Oferecimento{}

	db, err := sql.Open(configs.GetDBType(), configs.GetPostgresConnString())
	tools.CheckError(err)

	defer db.Close()

	userId := c.Get("userId")

	rows, err := db.Query(stmt, userId)
	tools.CheckError(err)

	for rows.Next() {
		var course model.Oferecimento

		if err := rows.Scan(&course.Id_disciplina, &course.Sigla, &course.DiaSemana, &course.Nome, &course.Inicio, &course.Fim); err != nil {
			log.Fatal(err)
		}
		courses = append(courses, course)
	}

	return c.JSON(http.StatusOK, courses)
}
