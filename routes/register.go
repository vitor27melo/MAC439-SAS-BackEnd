package routes

import (
	"backend/configs"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) (err error) {
	stmtQuery := `
		INSERT INTO
		    usuario(nome, email, cpf, senha)
		VALUES ($1, $2, $3, $4)

	`

	cred := new(Credentials)
	cred.Username = c.FormValue("username")
	cred.Password = c.FormValue("password")
	cred.Cpf = c.FormValue("cpf")
	cred.Name = c.FormValue("name")

	if cred.Username == "" || cred.Password == "" {
		return c.JSON(http.StatusBadRequest, "Informações insuficientes!")

	}

	db, err := sql.Open(configs.GetDBType(), configs.GetPostgresConnString())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Erro na conexão com o banco de dados.")
	}

	defer db.Close()

	_, err = db.Exec(stmtQuery, cred.Name, cred.Username, cred.Cpf, cred.Password)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "Informações insuficientes ou inconsistentes!")
	}

	token, err := configs.CreateJWT(cred.UserId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Erro na criação do JWT.")
	}
	return c.JSON(http.StatusOK, map[string]string{"token": token, "nome": cred.Name})
}
