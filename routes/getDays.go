package routes

import (
	"backend/model"
	"backend/tools"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func GetDays(c echo.Context) error {
	driver, err := neo4j.NewDriver("neo4j+s://3184388f.databases.neo4j.io:7687", neo4j.BasicAuth("neo4j", "ArbFLmo6VLVl1qHLUMzQalzXmOwBlsFqpyafiX7a218", ""))
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
