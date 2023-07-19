package indexer

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

func getClient() *mongo.Client {
	constants := GetConstants()
	connString := constants.mongoUri
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connString))
	defer cancel()
	if err != nil {
		panic(err)
	}

	return client
}

func GetCrawlPage() PageCrawl {
	// https://www.mongodb.com/docs/drivers/go/current/usage-examples/findOne/
	client := getClient()
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

func StoreAllCrawlPages(i *Index) error {
	fmt.Println("[StoreAllCrawlPages] routine start")
	client := getClient()
	collection := client.Database("mse").Collection("crawl")

	ctx := context.TODO()

	filter := bson.D{}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to execute find query: %s", err)
	}
	defer cursor.Close(ctx)

	// store documents
	for cursor.Next(ctx) {
		var result PageCrawl
		if err := cursor.Decode(&result); err != nil {
			return fmt.Errorf("failed to decode document: %s", err)
		}
		doc := NewDocument(result.ID.Hex(), result.BodyText)
		i.Store(doc)
	}

	if err := cursor.Err(); err != nil {
		return fmt.Errorf("cursor error: %s", err)
	}
	fmt.Println(time.Now(), " - Indexing finished")
	return nil
}
