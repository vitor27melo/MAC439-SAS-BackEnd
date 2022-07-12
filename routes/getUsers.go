package routes

import (
	"backend/tools"
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var coll *mongo.Collection

func GetUsers(c echo.Context) error {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://mac439:Atila_Iamarino@mac439.fsau9s1.mongodb.net/?retryWrites=true&w=majority"))
	tools.CheckError(err)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	tools.CheckError(err)

	defer client.Disconnect(ctx)

	coll = client.Database("test").Collection("usuario")

	findOptions := options.Find()
	findOptions.SetLimit(5)
	cursor, err := coll.Find(context.TODO(), bson.D{{}}, findOptions)
	tools.CheckError(err)

	var results []bson.M
	err = cursor.All(context.TODO(), &results)
	tools.CheckError(err)

	return c.JSON(http.StatusOK, results)
}
