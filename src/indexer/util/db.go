package util

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PageCrawl struct {
	url         string
	title       string
	bodyText    string
	mainText    string
	description string
	rawHtml     string
	crawlDate   string
	indexedDate string
}

func (page PageCrawl) String() string {
	return fmt.Sprintf("url: %s, title: %s", page.url, page.title)
}

func getClient() *mongo.Client {
	connString := "mongodb://db:27017"
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connString))
	defer cancel()
	if err != nil {
		panic(err)
	}

	return client
}

func GetCrawlPage() string {
	// https://www.mongodb.com/docs/drivers/go/current/usage-examples/findOne/
	client := getClient()
	collection := client.Database("mse").Collection("crawl")
	var result PageCrawl
	filter := bson.D{{"url", "https://www.tuebingen.de/"}}
	ctx := context.TODO()
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println(json.MarshalIndent(result, "", "  "))
	return "function executed in go"
}
