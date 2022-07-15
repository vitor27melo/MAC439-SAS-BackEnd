package routes

import (
	"backend/configs"
	"backend/tools"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	UserId   string
}

func Login(c echo.Context) (err error) {
	stmtQuery := `
		SELECT
			id_usuario
		FROM
			usuario
		WHERE
			cpf = $1
			AND senha = $2;
	`

	cred := new(Credentials)
	if err = c.Bind(cred); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	db, err := sql.Open(configs.GetDBType(), configs.GetPostgresConnString())
	tools.CheckError(err)

	defer db.Close()

	err = db.QueryRow(stmtQuery, cred.Username, cred.Password).Scan(&cred.UserId)
	tools.CheckError(err)

	token, err := configs.CreateJWT(cred.UserId)
	tools.CheckError(err)

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
