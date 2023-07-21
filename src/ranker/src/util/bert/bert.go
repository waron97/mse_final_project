package bert

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"ranker/src/util/core"
)

type response struct {
	Embeddings []core.Vector `json:"embeddings"`
}

type request struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

// ToDo - replace with BERT embeddings
func GetEmbeddings(text string) []core.Vector {
	constants := core.GetConstants()
	body := request{
		Text: text,
		Type: "query",
	}
	jsonPayload, err := json.Marshal(body)
	if err != nil {
		return nil
	}
	bodyReader := bytes.NewReader(jsonPayload)
	req, err := http.NewRequest(http.MethodPost, constants.BertUri+"/encode", bodyReader)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer res.Body.Close()
	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil
	}
	var resp response
	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		return nil
	}
	return resp.Embeddings
}
