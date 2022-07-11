package routes

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"net/http"
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
			"MATCH (n:Dia) RETURN n LIMIT 25",
			map[string]interface{}{})
		if err != nil {
			return nil, err
		}

		if result.Next() {
			fmt.Print("%s", result.Record().Values[0])
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, greeting.Labels[0])
}
