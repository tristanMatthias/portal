package server

import (
	"context"
	events "portal/server/events"
)

// App struct
type App struct {
	ctx context.Context
}

// Portal creates a new App application struct
func Portal() *App {
	return &App{}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	events.GenerateEventHandlers(ctx)
}

// // Greet returns a greeting for the given name
// func (a *App) Greet(name string) string {
// 	return fmt.Sprintf("Hello %s, It's show time!", name)
// }
