package scenes

import (
	"demo-x/conf"
	"demo-x/model"
	"demo-x/service"
	"github.com/gogf/gf/util/gconv"
	"github.com/rivo/tview"
)

func QuestsScene() *tview.Flex {

	formatBox := aFormatBox
	footBox := aFootBox
	statusBox := tview.NewBox()
	questsBox := tview.NewBox()
	mapBox := tview.NewBox()
	helpBox := tview.NewBox()
	newsBox := tview.NewBox()
	gameBox := tview.NewBox()

	questsList := service.GameData.UserQuestDoing
	lineList := []model.TextPrintStruct{}
	for i := 1; i <= len(questsList) && i <= 5; i++ {
		lineList = append(lineList,
			model.TextPrintStruct{Line: "任务名称：" + questsList[i].QuestsName, Align: tview.AlignLeft, Color: conf.BrightColor},
			model.TextPrintStruct{Line: "任务详情：" + questsList[i].QuestsInfo, Align: tview.AlignLeft, Color: conf.NormalColor},
		)
		for j := 0; j < len(questsList[i].QuestsCondition); j++ {
			lineStr := "达成条件：" + questsList[i].QuestsCondition[j].QuestsCellInfo
			if questsList[i].QuestsCondition[j].QuestsCellCount != 1 {

				lineStr = lineStr + " * " + gconv.String(questsList[i].QuestsCondition[j].QuestsCellCount)
			}
			lineList = append(lineList,
				model.TextPrintStruct{Line: lineStr, Align: tview.AlignLeft, Color: conf.SecondColor})
		}
		lineList = append(lineList, model.TextPrintStruct{Line: "", Align: tview.AlignLeft, Color: conf.SecondColor})
	}
	gameBox = createBox("", lineList)

	gameFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(formatBox, 0, 1, false).
		AddItem(gameBox, 30, 0, false).
		AddItem(formatBox, 0, 1, false).
		AddItem(footBox, 5, 0, false)
	flex := baseScene(statusBox, questsBox, mapBox, helpBox, newsBox, gameFlex)

	return flex

}
