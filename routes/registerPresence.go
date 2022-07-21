package routes

import (
	"backend/configs"
	"backend/tools"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"

	"fmt"
)

func RegisterPresence(c echo.Context) error {
	cpf := c.FormValue("cpf")
	date := c.FormValue("data")
	sigla := c.FormValue("sigla")

	driver, err := neo4j.NewDriver(configs.Neo4JURI, neo4j.BasicAuth(configs.Neo4JUsername, configs.Neo4JPassword, ""))
	if err != nil {
		panic(err)
	}
	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	roleCall, err := session.Run("MATCH (u:User) WHERE u.cpf = $cpf MATCH (e:Evento)-[:ACONTECEU]->(d:Dia) WHERE d.data = $date AND e.codigo = $sigla CREATE (u)-[:PARTICIPOU]->(e)", map[string]interface{}{"cpf": cpf, "date": date, "sigla": sigla})
	fmt.Print(roleCall)
	tools.CheckError(err)

	return c.JSON(http.StatusOK, nil)
}
