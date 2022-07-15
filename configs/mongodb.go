package configs

import (
	"backend/tools"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoClient() (*mongo.Client, context.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://mac439:Atila_Iamarino@mac439.fsau9s1.mongodb.net/?retryWrites=true&w=majority"))
	tools.CheckError(err)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	tools.CheckError(err)

	return client, ctx
}
