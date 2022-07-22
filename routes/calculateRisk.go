package routes

import (
	"backend/configs"
	"backend/tools"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"

	"fmt"
)

func CalculateRisk(c echo.Context) error {
	riskLevel := 0
	url := c.Request().RequestURI
	url_list := strings.Split(url, "/")
	name := url_list[len(url_list)-1]

	driver, err := neo4j.NewDriver(configs.Neo4JURI, neo4j.BasicAuth(configs.Neo4JUsername, configs.Neo4JPassword, ""))
	if err != nil {
		panic(err)
	}
	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	directProbs, err := neo4j.Collect(session.Run("MATCH (self:User)-[:PARTICIPOU]->(e1:Evento)<-[:PARTICIPOU]-(other:User)-[:SUSPEITA]->(p:ProbCovid) WHERE self.cpf = $cpf RETURN p.confianca AS probabilities", map[string]interface{}{"cpf": name}))
	tools.CheckError(err)

	for _, p := range directProbs {
		n := p.Values[0].(int64)
		//fmt.Print(n)
		if int(n)-1 > riskLevel {
			riskLevel = int(n) - 1
		}

	}

	indirectProbs, err := neo4j.Collect(session.Run("MATCH (self:User)-[:PARTICIPOU]->(e1:Evento)-[:ACONTECEU]->(d:Dia)<-[:ACONTECEU]-(e2:Evento)-[:PARTICIPOU]-(other:User)-[:SUSPEITA]->(p:ProbCovid) WHERE self.cpf = $cpf RETURN p.confianca AS probabilities", map[string]interface{}{"$cpf": name}))
	tools.CheckError(err)

	for _, p := range indirectProbs {
		n := p.Values[0].(int64)
		fmt.Print(n)
		if int(n)-1 > riskLevel {
			riskLevel = int(n) - 2
		}

	}

	// Update ProbCovid node
	// session.Run("")

	return c.JSON(http.StatusOK, riskLevel)
}
