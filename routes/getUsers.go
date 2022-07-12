package routes

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

var coll *mongo.Collection

func GetUsers(c echo.Context) error {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://mac439:Atila_Iamarino@mac439.fsau9s1.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	coll = client.Database("test").Collection("usuario")

	findOptions := options.Find()
	findOptions.SetLimit(5)
	cursor, err := coll.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		panic(err)
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}
	return c.JSON(http.StatusOK, "")
}
