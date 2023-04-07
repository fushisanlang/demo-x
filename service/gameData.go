package service

import (
	"demo-x/conf"
	"demo-x/model"
	"os"
)

var GameData model.GameDataStruct

func VeriftSave() bool {
	statusCode := true
	file, err := os.Open(conf.SaveFile)
	if err != nil {
		statusCode = false
	}
	defer file.Close()
	return statusCode
}

func InitData() {

	GameData = model.GameDataStruct{
		UserName:  "丹阳子",
		UserLevel: 0,
		UserEx:    0,
		UserStatus: model.StatusStruct{
			UserJin:     0,
			UserMu:      0,
			UserShui:    0,
			UserHuo:     0,
			UserTu:      0,
			UserYinYang: 0,
		},
		UserBag: model.BagsStruct{
			UseId: 0,
			Bag:   [64]model.CaseStruct{},
		},
		UserSaveId:     0,
		UserQuestDoing: map[int]model.QuestsInfoStruct{},
		UserQuestsDone: map[int]model.QuestsInfoStruct{},
		UserMateriel: model.MaterielsStruct{
			HeartSkill: model.SkillsStruct{
				SkillsId:         0,
				SkillsName:       "",
				SkillsInfo:       "",
				SkillType:        "",
				SkillTypeId:      0,
				SkillsFunc:       nil,
				SkillsUseActive:  0,
				SkillsUseByOther: 0,
			},
			BodySkill: model.SkillsStruct{
				SkillsId:         0,
				SkillsName:       "",
				SkillsInfo:       "",
				SkillType:        "",
				SkillTypeId:      0,
				SkillsFunc:       nil,
				SkillsUseActive:  0,
				SkillsUseByOther: 0,
			},
			ArmFirst: model.ArmsStruct{
				ArmsId:   0,
				ArmsName: "",
				ArmsInfo: "",
				ArmsStatus: model.StatusStruct{
					UserJin:     0,
					UserMu:      0,
					UserShui:    0,
					UserHuo:     0,
					UserTu:      0,
					UserYinYang: 0,
				},
				ArmsSkill: model.SkillsStruct{
					SkillsId:         0,
					SkillsName:       "",
					SkillsInfo:       "",
					SkillType:        "",
					SkillTypeId:      0,
					SkillsFunc:       nil,
					SkillsUseActive:  0,
					SkillsUseByOther: 0,
				},
			},
			ArmSecond: model.ArmsStruct{
				ArmsId:   0,
				ArmsName: "",
				ArmsInfo: "",
				ArmsStatus: model.StatusStruct{
					UserJin:     0,
					UserMu:      0,
					UserShui:    0,
					UserHuo:     0,
					UserTu:      0,
					UserYinYang: 0,
				},
				ArmsSkill: model.SkillsStruct{
					SkillsId:         0,
					SkillsName:       "",
					SkillsInfo:       "",
					SkillType:        "",
					SkillTypeId:      0,
					SkillsFunc:       nil,
					SkillsUseActive:  0,
					SkillsUseByOther: 0,
				},
			},
			ArmThird: model.ArmsStruct{
				ArmsId:   0,
				ArmsName: "",
				ArmsInfo: "",
				ArmsStatus: model.StatusStruct{
					UserJin:     0,
					UserMu:      0,
					UserShui:    0,
					UserHuo:     0,
					UserTu:      0,
					UserYinYang: 0,
				},
				ArmsSkill: model.SkillsStruct{
					SkillsId:         0,
					SkillsName:       "",
					SkillsInfo:       "",
					SkillType:        "",
					SkillTypeId:      0,
					SkillsFunc:       nil,
					SkillsUseActive:  0,
					SkillsUseByOther: 0,
				},
			},
		},
	}

}
func LoadSave() {

}
