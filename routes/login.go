package routes

import (
	"backend/configs"
	"backend/tools"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Credentials struct {
	username string
	password string
	userId   string
}

func Login(c echo.Context) error {
	stmtQuery := `
		SELECT
			id_usuario
		FROM
			usuario
		WHERE
			cpf = ?
			AND senha = ?;
	`

	cred := &Credentials{}
	err := c.Bind(&cred)
	tools.CheckError(err)

	db, err := sql.Open(configs.GetDBType(), configs.GetPostgresConnString())
	tools.CheckError(err)

	defer db.Close()

	err = db.QueryRow(stmtQuery, cred.username, cred.password).Scan(&cred.userId)
	tools.CheckError(err)

	token, err := configs.CreateJWT(cred.userId)
	tools.CheckError(err)

	// stmtExec := `
	// 	UPDATE usuario
	// 	SET token = ?
	// 	WHERE id_usuario = ?
	// `

	// db.Exec(stmtExec, token, cred.userId)

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
