package server

import (
	"context"
	"portal/server/chat"
	"portal/server/events"
	huggingface "portal/server/hugging-face"
	"portal/server/model"
)

type ServerModule struct {
	ctx context.Context
	Controllers map[string]interface{}
}

// Server creates a new App application struct
func Server() *ServerModule {
	controllers := make(map[string]interface{})
	controllers["chat"] = chat.ChatController()
	controllers["model"] = model.ModelController()
	controllers["huggingface"] = huggingface.HuggingFaceController()

	return &ServerModule{
		Controllers: controllers,
	}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (s *ServerModule) Startup(ctx context.Context) {
	s.ctx = ctx
	// Loop through the controllers and call their startup methods
	for _, controller := range s.Controllers {
		// If the controller doesn't have a startup method, it will be skipped
		i, ok := controller.(interface {
			Startup(context.Context)
		})

		if ok {
			i.Startup(ctx)
		}
	}
	events.GenerateEventHandlers(ctx)
}
