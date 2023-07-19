package indexer

type Document struct {
	Id   string
	Text string
}

func NewDocument(id, text string) *Document {
	return &Document{
		Id:   id,
		Text: text,
	}
}

func (d *Document) String() string {
	return "<" + d.Id + ": " + d.Text + ">"
}
