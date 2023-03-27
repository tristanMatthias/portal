package main

import (
	"embed"
	"os"
	s "portal/server"
	"portal/server/lib"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create the config directory if it doesn't exist
	os.Mkdir(lib.ConfigPath(""), 0755)

	// Create an instance of the server structure
	server := s.Server()

	// Reduce server controller to an array of interfaces
	controllers := make([]interface{}, len(server.Controllers))
	i := 0
	for _, v := range server.Controllers {
		controllers[i] = v
		i++
	}

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Portal",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        server.Startup,
		Bind: append([]interface{}{
			server,
		}, controllers...),
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
