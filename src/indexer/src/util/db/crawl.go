package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getCrawlPageFilter() bson.D {
	filter := bson.D{
		{Key: "$and", Value: bson.A{
			bson.D{{Key: "indexedDate", Value: bson.D{{Key: "$exists", Value: false}}}},
			bson.D{{Key: "relevant", Value: true}},
		}},
	}

	return filter
}

func CountCrawlPages() int {
	client := getClient()
	collection := client.Database("mse").Collection("crawl")
	ctx := context.TODO()
	count, err := collection.CountDocuments(ctx, getCrawlPageFilter())

	if err != nil {
		panic(err)
	}

	return int(count)
}

func GetCrawlPages(channel chan PageCrawl) int {
	client := getClient()
	collection := client.Database("mse").Collection("crawl")
	var count int = 0

	ctx := context.TODO()

	filter := getCrawlPageFilter()
	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		panic(err)
	}

	for cursor.Next(ctx) {
		var result PageCrawl
		if err := cursor.Decode(&result); err != nil {
			// this is a fatal error because we expect a
			// specific number of documents to be added
			// to the channel
			panic(err)
		} else {
			channel <- result
			count++
		}
	}

	defer cursor.Close(ctx)

	close(channel)
	return count
}

func MarkPageIndexed(id string) {
	client := getClient()
	collection := client.Database("mse").Collection("crawl")

	ctx := context.TODO()
	docId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{
		{Key: "_id", Value: docId},
	}
	currentDate := primitive.DateTime(time.Now().UnixNano() / int64(time.Millisecond))
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "indexedDate", Value: currentDate},
		}},
	}
	_, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		panic(err)
	}
}
