package routes

import (
	"backend/configs"
	"backend/tools"
	"context"
	"database/sql"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

var stmtFindCpf = `
	SELECT
		cpf
	FROM
		usuario
	WHERE
		id_usuario = $1;
`

func GetUserFiles(c echo.Context) error {
	client, ctx := configs.GetMongoClient()
	defer client.Disconnect(ctx)

	db, err := sql.Open(configs.GetDBType(), configs.GetPostgresConnString())
	tools.CheckError(err)

	defer db.Close()

	userId := c.Get("userId")
	var cpf string
	err = db.QueryRow(stmtQuery, userId).Scan(&cpf)
	tools.CheckError(err)

	coll := client.Database("test").Collection("usuario")

	findOptions := options.Find()
	cursor, err := coll.Find(context.TODO(), bson.D{{}}, findOptions)

	var results []bson.M
	err = cursor.All(context.TODO(), &results)

	return c.JSON(http.StatusOK, results)

}
