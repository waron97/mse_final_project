package db

import (
	"context"
	"time"

	"ranker/src/util/core"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getClient() *mongo.Client {
	constants := core.GetConstants()
	connString := constants.MongoUri
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connString))
	defer cancel()
	if err != nil {
		panic(err)
	}

	return client
}
