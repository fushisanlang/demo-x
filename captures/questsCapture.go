package captures

import (
	"demo-x/scenes"
	"demo-x/service"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func QuestsCapture(app *tview.Application) *tview.Flex {

	scene1 := scenes.QuestsScene(questCode)
	scene1.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRight && questCode < 3 {
			questCode = questCode + 1
			app.SetRoot(QuestsCapture(app), true)
		} else if event.Key() == tcell.KeyLeft && questCode > 1 {
			questCode = questCode - 1
			app.SetRoot(QuestsCapture(app), true)
		} else if event.Key() == tcell.KeyEsc {
			app.SetRoot(CaptureMap[service.GameData.UserSaveId](app), true)
		}
		return event
	})
	return scene1

}
