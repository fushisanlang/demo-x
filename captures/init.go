package captures

import "github.com/rivo/tview"

var startCode, startMaxCode, questCode int

func init() {
	startCode = 0
	startMaxCode = 6
	questCode = 1 //任务页翻面
	initCaptures()
}

var CaptureMap map[int]func(app *tview.Application) *tview.Flex

func initCaptures() {
	CaptureMap = map[int]func(app *tview.Application) *tview.Flex{
		1: StartCapture,
	}
}
