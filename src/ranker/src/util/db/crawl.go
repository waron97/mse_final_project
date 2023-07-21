package db

import (
	"context"
	"ranker/src/util/core"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPage(docId string) PageCrawl {
	documentId, err := primitive.ObjectIDFromHex(docId)
	if err != nil {
		core.ErrPanic(err)
	}
	client := getClient()
	var page PageCrawl
	ctx := context.TODO()
	filter := bson.D{{Key: "_id", Value: documentId}}
	err = client.Database("mse").Collection("crawl").FindOne(ctx, filter).Decode(&page)
	if err != nil {
		core.ErrPanic(err)
	}
	return page
}
