package scenes

import (
	"demo-x/conf"
	"demo-x/model"
	"demo-x/service"
	"github.com/gogf/gf/util/gconv"
	"github.com/rivo/tview"
)

func QuestsScene(pageId int) *tview.Flex {

	formatBox := aFormatBox
	footBox := aFootBoxESC
	statusBox := tview.NewBox()
	questsBox := tview.NewBox()
	mapBox := tview.NewBox()
	helpBox := tview.NewBox()
	newsBox := tview.NewBox()
	gameBox := tview.NewBox()

	questsList := service.GameData.UserQuestDoing
	lineList := []model.TextPrintStruct{}

	for i := 1 + (pageId-1)*5; i <= len(questsList) && i <= 5*pageId; i++ {
		lineList = append(lineList,
			model.TextPrintStruct{Line: "任务" + gconv.String(i) + " ：" + questsList[i].QuestsName, Align: tview.AlignLeft, Color: conf.BrightColor},
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
	gameBox = createBox("第"+gconv.String(pageId)+"页", lineList)
	pageBox := tview.NewBox()
	switch pageId {
	case 1:
		pageBox = createBoxWithoutSider([]model.TextPrintStruct{{"点击 \"->\" 键向右翻页", tview.AlignCenter, conf.GuideColor}})
	case 2:
		pageBox = createBoxWithoutSider([]model.TextPrintStruct{{"点击 \"<-/->\" 键向左/右翻页", tview.AlignCenter, conf.GuideColor}})
	case 3:
		pageBox = createBoxWithoutSider([]model.TextPrintStruct{{"点击 \"<-\" 键 向左翻页", tview.AlignCenter, conf.GuideColor}})

	}

	gameFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(formatBox, 0, 1, false).
		AddItem(gameBox, 30, 0, false).
		AddItem(formatBox, 0, 1, false).
		AddItem(pageBox, 2, 0, false).
		AddItem(footBox, 5, 0, false)
	flex := baseScene(statusBox, questsBox, mapBox, helpBox, newsBox, gameFlex)

	return flex

}
