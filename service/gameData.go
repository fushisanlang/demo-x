package service

import (
	"demo-x/conf"
	"demo-x/model"
	"demo-x/tools"
	"encoding/gob"
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
	SaveData()
}

func LoadData() {
	file, err := os.Open(conf.SaveFile)
	tools.CheckErr(err)
	defer file.Close()
	// 创建一个解码器
	decoder := gob.NewDecoder(file)
	// 解码数据并存储到新的切片
	err = decoder.Decode(&GameData)
	tools.CheckErr(err)
}

func SaveData() {
	file, err := os.Create(conf.SaveFile)
	tools.CheckErr(err)
	defer file.Close()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(GameData)
	tools.CheckErr(err)
}
