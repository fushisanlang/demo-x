package service

import (
	"demo-x/model"
	"errors"
)

func AddQuests(infoStruct model.QuestsInfoStruct) error {
	questId := len(GameData.UserQuestDoing) + 1
	if questId >= 16 {
		return errors.New("无法新增，队列满")
	} else {
		GameData.UserQuestDoing[questId] = infoStruct
		return nil
	}
}
