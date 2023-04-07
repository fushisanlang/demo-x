package model

type SkillsStruct struct {
	SkillsId         int
	SkillsName       string
	SkillsInfo       string
	SkillType        string
	SkillTypeId      int //1。心法 2.功法 3.其他技能
	SkillsFunc       func()
	SkillsUseActive  int //1.主动技能 0。被动技能
	SkillsUseByOther int // 1.可以被其他人使用 0.只能主角使用
}
