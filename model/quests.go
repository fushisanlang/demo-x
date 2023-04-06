package model

type QuestsInfo struct {
	QuestsId        int
	QuestsName      string
	QuestsInfo      string
	QuestsCondition []QuestsCell
	QuestsLevel     int
	QuestsReward    []ManyGoods
	QuestsBefore    int
	QuestsMaster    int
}

type QuestsCell struct {
	QuestsCellInfo  string
	QuestsCellCount int
}
