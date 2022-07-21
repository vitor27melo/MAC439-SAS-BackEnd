package routes

import (
	"backend/configs"
	"backend/tools"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"

	"fmt"
)

func ReportCovid(c echo.Context) error {
	name := c.FormValue("nome")
	date := c.FormValue("data")

	driver, err := neo4j.NewDriver(configs.Neo4JURI, neo4j.BasicAuth(configs.Neo4JUsername, configs.Neo4JPassword, ""))
	if err != nil {
		panic(err)
	}
	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	createProb, err := neo4j.Collect(session.Run("MATCH (u:User) WHERE u.nome = $name MATCH (d:Dia) WHERE d.data = $date CREATE (p:ProbCovid{confiança:10}),(p)-[:ACONTECEU]->(d),(u)-[:SUSPEITA]->(p)", map[string]interface{}{"name": name, "date": date}))
	fmt.Print(createProb)
	tools.CheckError(err)

	return c.JSON(http.StatusOK, nil)
}
