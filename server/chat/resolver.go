package chat

import (
	"context"
	"log"

	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

type Response struct {
    Message  string `json:"message"`
    Response string `json:"response"`
}

func OnChat(ctx context.Context) func(...interface{}) {
	return func(data ...interface{}) {
		model, message := data[0].(string), data[1].(string)

		// TODO: Pull from database
		messages := []string{message}

		response, err := Chat(model, messages)
		if err != nil {
			log.Println("Error calling python API")
			log.Println(err)
			return
		}

		// Convert model and response to JSON
		msgResp := Response{Message: message, Response: response}

		// Send the response back to the frontend with the model name
		rt.EventsEmit(ctx, EventChatResponse, msgResp)
	}
}
