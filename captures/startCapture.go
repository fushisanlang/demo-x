package captures

import (
	"demo-x/scenes"
	"demo-x/service"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func StartCapture(app *tview.Application) *tview.Flex {
	service.GameData.UserSaveId = 1
	scene1 := scenes.StartScene(startCode)
	scene1.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter && startCode < startMaxCode {
			startCode = startCode + 1
			app.SetRoot(StartCapture(app), true)
			return nil
		} else if startCode >= startMaxCode {
			service.SaveData()
			switch event.Key() {
			case tcell.KeyF1:
				app.SetRoot(HelpCapture(app), true)
			case tcell.KeyF2:
				app.SetRoot(BagCapture(app), true)
			case tcell.KeyF4:
				app.SetRoot(QuestsCapture(app), true)
			}
			return nil
		}
		return event
	})
	return scene1

}
