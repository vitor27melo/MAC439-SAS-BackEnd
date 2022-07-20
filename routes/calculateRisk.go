package routes

import (
	"backend/configs"
	"backend/tools"

	"github.com/labstack/echo/v4"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"

	"fmt"
	"strconv"
)

func CalculateRisk(c echo.Context, name string) float64 {
	riskLevel := 0.0

	driver, err := neo4j.NewDriver(configs.Neo4JURI, neo4j.BasicAuth(configs.Neo4JUsername, configs.Neo4JPassword, ""))
	if err != nil {
		panic(err)
	}
	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	directProbs, err := neo4j.Collect(session.Run("MATCH (self:User)-[:PARTICIPOU]->(e1:Evento)<-[:PARTICIPOU]-(other:User)-[:SUSPEITA]->(p:ProbCovid) WHERE self.nome = $name RETURN p.confiança AS probabilities", map[string]interface{}{}))
	tools.CheckError(err)

	for _, p := range directProbs {
		if prob, found := p.Get("confiança"); found {
			n, err := strconv.ParseFloat(prob.(string))
			tools.CheckError(err)
			if n-1 > riskLevel {
				riskLevel = n - 1
			}
		}

	}

	//indirectProbs, err := neo4j.Collect(session.Run("MATCH (self:User)-[:PARTICIPOU]->(e1:Evento)-[:ACONTECEU]->(d:Dia)<-[:ACONTECEU]<-(e2:Evento)<-[:PARTICIPOU]-(other:User)-[:SUSPEITA]->(p:ProbCovid) WHERE self.nome = $name RETURN p.confiança AS probabilities", nil))
	//tools.CheckError(err)

	/*

		dp, ip, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, interface{}, error) {

			directProb, err := transaction.Run(
				"MATCH (self:User)-[:PARTICIPOU]->(e1:Evento)<-[:PARTICIPOU]-(other:User)-[:SUSPEITA]->(p:ProbCovid) WHERE self.nome = $name RETURN p.confiança AS probabilities",
				map[string]interface{}{})
			tools.CheckError(err)

			indirectProb, err := transaction.Run(
				"MATCH (self:User)-[:PARTICIPOU]->(e1:Evento)-[:ACONTECEU]->(d:Dia)<-[:ACONTECEU]<-(e2:Evento)<-[:PARTICIPOU]-(other:User)-[:SUSPEITA]->(p:ProbCovid) WHERE self.nome = $name RETURN p.confiança AS probabilities",
				map[string]interface{}{})
			tools.CheckError(err)

			return directProb,indirectProb
		})

	*/

	fmt.Print(c, name, err)
	return riskLevel
}
