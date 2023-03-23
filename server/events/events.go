package events

import (
	"context"

	chat "portal/server/chat"
	model "portal/server/model"

	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

// Mapping of event names to event handlers
var Events = map[string]func(...interface{}){}

func GenerateEventHandlers(ctx context.Context) {
	rt.EventsOn(ctx, model.EventDownload, model.OnDownload(ctx))
	rt.EventsOn(ctx, chat.EventChat, chat.OnChat(ctx))
}
