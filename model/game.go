package model

type GameDataStruct struct {
	UserName       string
	UserLevel      int
	UserEx         int
	UserStatus     StatusStruct
	UserBag        BagsStruct
	UserMateriel   MaterielsStruct
	UserSaveId     int                      //场景保存，用于记录线性剧情
	UserQuestsDone map[int]QuestsInfoStruct //已完成任务列表
	UserQuestDoing map[int]QuestsInfoStruct
}
type StatusStruct struct {
	//阴阳五行
	UserJin     int
	UserMu      int
	UserShui    int
	UserHuo     int
	UserTu      int
	UserYinYang int
}
