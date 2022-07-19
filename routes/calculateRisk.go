package routes

import (
	"backend/configs"

	"github.com/labstack/echo/v4"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func CalculateRisk(c echo.Context, id string) float64 {
	riskLevel := 0.0

	driver, err := neo4j.NewDriver(configs.Neo4JURI, neo4j.BasicAuth(configs.Neo4JUsername, configs.Neo4JPassword, ""))
	if err != nil {
		panic(err)
	}
	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	/*
		stmtQuery := `
		SELECT
			id_usuario
		FROM
			usuario
		WHERE
			cpf = $1
			AND senha = $2; `
	*/

	return riskLevel
}
