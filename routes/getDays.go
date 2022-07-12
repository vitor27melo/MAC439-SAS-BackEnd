package routes

import (
	"backend/configs"
	"backend/model"
	"backend/tools"
	"github.com/labstack/echo/v4"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"net/http"
)

func GetDays(c echo.Context) error {
	driver, err := neo4j.NewDriver(configs.Neo4JURI, neo4j.BasicAuth(configs.Neo4JUsername, configs.Neo4JPassword, ""))
	if err != nil {
		panic(err)
		return err
	}
	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	greeting, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (n:Dia) RETURN n.data as data LIMIT 25",
			map[string]interface{}{})
		tools.CheckError(err)

		days := []model.Day{}

		for result.Next() {
			var day model.Day

			if data, found := result.Record().Get("data"); found {
				day.Data = data.(string)
			}

			days = append(days, day)
		}

		return days, nil
	})
	tools.CheckError(err)

	return c.JSON(http.StatusOK, greeting)
}
