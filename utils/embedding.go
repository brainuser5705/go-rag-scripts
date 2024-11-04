package utils

import (
	"context"

	"github.com/carlmjohnson/requests"
)

// Input data into the API call
type embedData struct {
	Input string `json:"input"`
}

// Response format from LM Studios
type embedRes struct {
	Object string `json:"object"`
	Data []( struct {
		Object string
		Embedding []float32
		Index int
	} ) `json:"data"`
	Model string `json:"model"`
	Usage struct {
		PromptToken int
		TotalToken int
	} `json:"usage"`
}

func Embed(s string) []float32 {

	reqData := embedData{Input: s}

	var resStruct embedRes

	err := requests.
	URL("http://127.0.0.1:1234/v1/embeddings").
	Accept("application/json").
	BodyJSON(&reqData).
	ToJSON(&resStruct).
	Fetch(context.Background())

	if err != nil {
		panic(err)
	}

	return resStruct.Data[0].Embedding
}
