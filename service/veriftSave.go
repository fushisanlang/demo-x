package service

import (
	"os"
)

func VeriftSave() bool {
	//判断存档
	statusCode := true
	statusCode = false
	//dao.VeriftSave()
	//saveFileList := conf.SaveFileList
	//for i := 0; i < len(saveFileList); i++ {
	//	if !veriftSave(saveFileList[i]) {
	//		statusCode = false
	//		break
	//	}
	//}
	//return statusCode
	return statusCode
}
func veriftSave(saveFile string) bool {
	statusCode := true
	file, err := os.Open("person.bin")
	if err != nil {
		statusCode = false
	}
	defer file.Close()
	return statusCode
}
func LoadSave() {

}
