package chat

import (
	"log"
)

type Response struct {
	Prompt  string `json:"prompt"`
	Response string `json:"response"`
}

// Return a function that will be called that returns a Response or error
func (m *Chat) ActionChat(model string, prompt string) (Response, error) {
	// TODO: Pull from database
	messages := []string{prompt}

	response, err := PythonChat(model, messages)
	if err != nil {
		log.Println("Error calling python API")
		log.Println(err)
		return Response{}, err
	}

	// Convert model and response to JSON
	msgResp := Response{Prompt: prompt, Response: response}

	return msgResp, nil
}
