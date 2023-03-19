package captures

import (
	"demo-x/scenes"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func StartCapture(app *tview.Application) *tview.Flex {

	scene1 := scenes.StartScene(startCode)
	scene1.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter && startCode < startMaxCode {
			startCode = startCode + 1
			app.SetRoot(StartCapture(app), true)
			return nil
			//} else if event.Key() == tcell.KeyEnter {
			//	panic(1)
		}
		return event
	})
	return scene1

}
