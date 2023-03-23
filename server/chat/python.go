package chat

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const pythonApi = "http://localhost:8000"

// Call the python API at localhost:8080/chat/<model>
func PythonChat(model string, messages []string) (string, error) {
	// Convert messages to JSON string
	messageData, err := json.Marshal(map[string]interface{}{
		"messages": messages,
	})

	if err != nil {
		return "", err
	}

	// Send POST request to python API
	resp, err := http.Post(pythonApi+"/chat/"+model, "application/json", bytes.NewBuffer(messageData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
