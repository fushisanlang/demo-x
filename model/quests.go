package model

type QuestsInfoStruct struct {
	QuestsId        int
	QuestsName      string
	QuestsInfo      string
	QuestsCondition []QuestsCellStruct
	QuestsLevel     int
	QuestsReward    []ManyGoodsStruct
	QuestsBefore    int
	QuestsMaster    int
}

type QuestsCellStruct struct {
	QuestsCellInfo  string
	QuestsCellCount int
}
