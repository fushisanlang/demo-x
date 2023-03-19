package main

import (
	"demo-x/captures"

	"github.com/rivo/tview"
)

func main() {

	app := tview.NewApplication()

	startScene := captures.VerifyCapture(app)
	//@todo: 需要判断是否有存档

	app.SetRoot(startScene, true)
	//// Run the application.
	if err := app.Run(); err != nil {
		panic(err)
	}
}
