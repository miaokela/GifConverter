package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"gif-converter/utils"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	calculator := utils.NewCalculator()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "gif-converter",
		Width:  600,
		Height: 420,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			calculator, // 绑定新的结构体
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
