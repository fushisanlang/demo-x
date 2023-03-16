package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	// Create a new application.
	app := tview.NewApplication()

	// Create two scenes.
	scene1 := tview.NewFlex()
	text1 := tview.NewTextView().SetText("This is Scene 1")
	scene1.AddItem(text1, 0, 1, false)

	scene2 := tview.NewFlex()
	text2 := tview.NewTextView().SetText("This is Scene 2")
	scene2.AddItem(text2, 0, 1, false)

	// Set the application's root scene to the first scene.
	app.SetRoot(scene1, true)

	// Set an input capture for the first scene to switch to the second scene when the user presses the F1 key.
	scene1.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyF1 {
			app.SetRoot(scene2, true)
			return nil
		}
		return event
	})

	// Set an input capture for the second scene to switch to the first scene when the user presses the F2 key.
	scene2.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyF1 {
			app.SetRoot(scene1, true)
			return nil
		}
		return event
	})

	// Run the application.
	if err := app.Run(); err != nil {
		panic(err)
	}
}
