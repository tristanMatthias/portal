package events

import (
	"context"
)

// Mapping of event names to event handlers
var Events = map[string]func(...interface{}){}

func GenerateEventHandlers(ctx context.Context) {
}
