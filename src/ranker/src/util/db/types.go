package db

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Passage struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Text string             `bson:"text"`
}

func (passage Passage) String() string {
	return fmt.Sprintf("{Passage id: %s, length: %d}", passage.ID, len(passage.Text))
}

type PageCrawl struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	URL         string             `bson:"url"`
	Title       string             `bson:"title"`
	BodyText    string             `bson:"bodyText"`
	MainText    string             `bson:"mainText"`
	Description string             `bson:"description"`
	RawHtml     string             `bson:"rawHtml"`
	Passages    []Passage          `bson:"passages"`
	CrawlDate   primitive.DateTime `bson:"crawlDate"`
	IndexedDate primitive.DateTime `bson:"indexedDate"`
}

func (page PageCrawl) String() string {
	return fmt.Sprintf("{CrawlPage url: %s, title: %s}", page.URL, page.Title)
}

func (page PageCrawl) GetPassage(passageId string) string {
	for _, passage := range page.Passages {
		if passage.ID.Hex() == passageId {
			return passage.Text
		}
	}
	return ""
}
