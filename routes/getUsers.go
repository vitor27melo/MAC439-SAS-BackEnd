package routes

import (
	"backend/configs"
	"backend/tools"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUsers(c echo.Context) error {
	client, ctx := configs.GetMongoClient()
	defer client.Disconnect(ctx)

	coll := client.Database("test").Collection("usuario")

	findOptions := options.Find()
	findOptions.SetLimit(5)
	cursor, err := coll.Find(context.TODO(), bson.D{{}}, findOptions)
	tools.CheckError(err)

	var results []bson.M
	err = cursor.All(context.TODO(), &results)
	tools.CheckError(err)

	return c.JSON(http.StatusOK, results)
}
