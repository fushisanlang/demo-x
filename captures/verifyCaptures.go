package captures

import (
	"demo-x/scenes"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func VerifyCaptures(app *tview.Application) *tview.Flex {
	scene1 := scenes.VerifyScene()

	scene1.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			app.SetRoot(StartCaptures(app), true)
			return nil
		}
		return event
	})
	return scene1

}
