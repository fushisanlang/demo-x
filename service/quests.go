package service

import (
	"demo-x/conf"
	"demo-x/model"
	"fmt"
	"github.com/rivo/tview"
)

var Bags model.BagsStruct

/*
任务逻辑：

任务分为主线和支线两种。
通过QuestsMaster区分 0支线，1主线

任务结构体中有如下数据：

	QuestsId  总表中的id
	QuestsName string
	QuestsInfo string
	QuestsCondition 任务完成条件，由n个子条件组成list构成  []QuestsCell
	QuestsLevel 需求等级 int
	QuestsReward  任务奖励，由manygoods组成的列表构成。 ManyGoods 则是 由物品id和个数组成的结构体，用于表示多个物体
	QuestsMaster  主线与否 int
	QuestsBefore 前置任务

系统内置的任务存放在data中
维护一张人物任务表
接取任务后，加载该任务的信息到任务列表。
通过任务详情界面查看任务信息，切换任务追踪
完成的任务记作已完成，不会再接收
//任务间没有前后置，没有连环任务。
但是任务只能保留3主10支。
*/
func initAllQuests() {

	a := []model.QuestsInfo{
		{QuestsId: 1,
			QuestsName: "运转《纳元诀》",
			QuestsInfo: "运转《纳元诀》",
			QuestsCondition: []model.QuestsCell{

				{
					QuestsCellInfo:  "运转《纳元诀》",
					QuestsCellCount: 1,
				},
			},
			QuestsLevel: 0,
			QuestsReward: []model.ManyGoods{
				{
					GoodsId:    999, //经验值
					GoodsCount: 1,
				},
			},
			QuestsBefore: 0,
			QuestsMaster: 1},

		{QuestsId: 2,
			QuestsName:      "",
			QuestsInfo:      "",
			QuestsCondition: nil,
			QuestsLevel:     0,
			QuestsReward:    nil,
			QuestsBefore:    0,
			QuestsMaster:    0},
	}
}
