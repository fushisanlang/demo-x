package main

import (
	"demo-x/captures"

	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	startScene := captures.StartScene(app)
	//startScene2 := scenes.StartScene2(app)

	// Set the application's root scene to the first scene.
	app.SetRoot(startScene, true)
	// Run the application.
	if err := app.Run(); err != nil {
		panic(err)
	}
}
