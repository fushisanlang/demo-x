package captures

import (
	"demo-x/scenes"

	"github.com/rivo/tview"
)

func StartCapture(app *tview.Application) *tview.Flex {
	scene1 := scenes.StartScene()

	return scene1

}
