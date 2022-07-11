package routes

//import (
//	"github.com/labstack/echo/v4"
//	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
//)
//
//func GetDays(c echo.Context) error {
//	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
//	defer session.Close()
//
//	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
//		result, err := tx.Run("CREATE (a:Person {name: $name})", map[string]interface{}{"name": name})
//		if err != nil {
//			return nil, err
//		}
//
//		return result.Consume()
//	})
//
//	return err
//}
