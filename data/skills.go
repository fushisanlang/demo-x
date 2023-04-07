package data

import (
	"demo-x/model"
	"fmt"
)

var SkillsMap map[int]model.SkillsStruct

func AddEx() {
	fmt.Println("add ex")
}
func initSkills() {

	SkillsMap = map[int]model.SkillsStruct{
		1: model.SkillsStruct{
			SkillsId:         1,
			SkillsName:       "纳元诀",
			SkillsInfo:       "三才宗入门心法，调用后可以缓慢增加修为。",
			SkillType:        "心法",
			SkillTypeId:      1,
			SkillsFunc:       AddEx,
			SkillsUseActive:  0,
			SkillsUseByOther: 1,
		},
	}
}
