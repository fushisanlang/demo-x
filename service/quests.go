package service

import "demo-x/model"

func AddQuests(infoStruct model.QuestsInfoStruct) {
	questId := len(GameData.UserQuestDoing) + 1
	GameData.UserQuestDoing[questId] = infoStruct
}
