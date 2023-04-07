package captures

import (
	"demo-x/scenes"
	"demo-x/service"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func BagCapture(app *tview.Application) *tview.Flex {

	scene1 := scenes.BagScene()
	scene1.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter && startCode < startMaxCode {
			startCode = startCode + 1
			app.SetRoot(StartCapture(app), true)
			return nil
		} else if event.Key() == tcell.KeyEsc {
			app.SetRoot(CaptureMap[service.GameData.UserSaveId](app), true)
		}
		return event
	})
	return scene1

}
