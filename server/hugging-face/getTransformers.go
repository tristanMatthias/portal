package huggingface

import (
	"fmt"
	"log"

	"github.com/monaco-io/request"
)

type APIQueryParams struct {
	Limit     int
	Search    string
	Sort      string
	Direction int
}

type HFModel struct {
	ID          string `json:"id"`
	ModelID     string `json:"model_id"`
	PipelineTag string `json:"pipeline_tag"`
	Private     bool   `json:"private"`
	Downloads   int    `json:"downloads"`
}

func GetTransformers(params APIQueryParams) ([]HFModel, error) {
	var result []HFModel
	// Set the API endpoint URL and add optional parameters
	apiURL := "https://huggingface.co/api/models"

	var query = map[string]string{
		"task": "text-generation",
	}

	if params.Limit != 0 {
		query["limit"] = fmt.Sprintf("%d", params.Limit)
	}
	if params.Search != "" {
		query["search"] = params.Search
	}
	if params.Sort != "" {
		query["sort"] = params.Sort
	}
	if params.Direction != 0 {
		query["direction"] = fmt.Sprintf("%d", params.Direction)
	}

	// Send an HTTP GET request to the API endpoint
	c := request.Client{
		URL:    apiURL,
		Method: "GET",
		Query:  query,
	}
	resp := c.Send().Scan(&result)

	if !resp.OK() {
		// handle error
		log.Println(resp.Error())
		return nil, resp.Error()
	}

	return result, nil
}
