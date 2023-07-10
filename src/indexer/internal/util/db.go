package util

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PageCrawl struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	URL         string             `bson:"url"`
	Title       string             `bson:"title"`
	BodyText    string             `bson:"bodyText"`
	MainText    string             `bson:"mainText"`
	Description string             `bson:"description"`
	RawHtml     string             `bson:"rawHtml"`
	CrawlDate   primitive.DateTime `bson:"crawlDate"`
	IndexedDate primitive.DateTime `bson:"indexedDate"`
}

func (page PageCrawl) String() string {
	return fmt.Sprintf("url: %s, title: %s", page.URL, page.Title)
}

func getClient(connString string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connString))
	defer cancel()
	if err != nil {
		panic(err)
	}

	return client
}

func GetCrawlPage(connString string) PageCrawl {
	// https://www.mongodb.com/docs/drivers/go/current/usage-examples/findOne/
	client := getClient(connString)
	collection := client.Database("mse").Collection("crawl")
	var result PageCrawl
	filter := bson.D{{Key: "url", Value: "https://www.tuebingen.de/"}}
	ctx := context.TODO()
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", result.String())
	return result
}

func GetAllCrawlPages(connString string, time ...time.Time) ([]PageCrawl, error) {
	client := getClient(connString)
	collection := client.Database("mse").Collection("crawl")
	var results []PageCrawl

	// Only get documents from time onwards, if time is specified
	filter := bson.D{}
	if len(time) > 0 {
		filter = bson.D{
			{Key: "crawlDate", Value: bson.D{{Key: "$gte", Value: primitive.NewDateTimeFromTime(time[0])}}},
		}
	}
	ctx := context.TODO()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to execute find query: %s", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var result PageCrawl
		if err := cursor.Decode(&result); err != nil {
			return nil, fmt.Errorf("failed to decode document: %s", err)
		}
		results = append(results, result)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %s", err)
	}

	return results, nil
}
