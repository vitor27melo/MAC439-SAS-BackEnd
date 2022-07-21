package routes

import (
	"backend/configs"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	UserId   string
	Name     string
	Cpf      string
}

func Login(c echo.Context) (err error) {
	stmtQuery := `
		SELECT
			id_usuario,
			nome,
			cpf
		FROM
			usuario
		WHERE
			(cpf = $1 OR email = $1)
			AND senha = $2;
	`

	cred := new(Credentials)
	cred.Username = c.FormValue("username")
	cred.Password = c.FormValue("password")

	if cred.Username == "" || cred.Password == "" {
		return c.JSON(http.StatusNotFound, "Usuário ou senha não encontrados!")

	}

	db, err := sql.Open(configs.GetDBType(), configs.GetPostgresConnString())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Erro na conexão com o banco de dados.")
	}

	defer db.Close()

	err = db.QueryRow(stmtQuery, cred.Username, cred.Password).Scan(&cred.UserId, &cred.Name, &cred.Cpf)

	if err != nil {
		return c.JSON(http.StatusNotFound, "Usuário ou senha não encontrados!")
	}

	token, err := configs.CreateJWT(cred.UserId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Erro na criação do JWT.")
	}
	return c.JSON(http.StatusOK, map[string]string{"token": token, "nome": cred.Name, "cpf": cred.Cpf})
}
