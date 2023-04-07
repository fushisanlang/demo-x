package model

type MaterielsStruct struct {
	HeartSkill SkillsStruct //内功
	BodySkill  SkillsStruct //外功
	ArmFirst   ArmsStruct   //法宝1
	ArmSecond  ArmsStruct   //法宝2
	ArmThird   ArmsStruct   //法宝3
}

type ArmsStruct struct {
	ArmsId     int
	ArmsName   string
	ArmsInfo   string
	ArmsStatus StatusStruct
	ArmsSkill  SkillsStruct
}
