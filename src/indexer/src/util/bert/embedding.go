package bert

import (
	"bytes"
	"encoding/json"
	"errors"
	"indexer/src/util/core"
	"io"
	"net/http"
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
		Type: "document",
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

func GetAvgEmbedding(emb []core.Vector) (core.Vector, error) {
	rows := len(emb)
	if rows == 0 {
		return nil, errors.New("embeddings empty")
	}

	columns := len(emb[0])
	averages := make(core.Vector, columns)

	for col := 0; col < columns; col++ {
		sum := 0.0
		for row := 0; row < rows; row++ {
			sum += emb[row][col]
		}
		averages[col] = sum / float64(rows)
	}
	return averages, nil
}

func IsAlive() bool {
	constants := core.GetConstants()
	req, err := http.NewRequest(http.MethodGet, constants.BertUri+"/health", nil)
	if err != nil {
		return false
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return false
	}
	defer res.Body.Close()
	return res.StatusCode == http.StatusOK

}
