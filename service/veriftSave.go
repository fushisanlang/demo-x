package service

import (
	"demo-x/conf"
	"os"
)

func VeriftSave() bool {
	statusCode := true
	saveFileList := conf.SaveFileList
	for i := 0; i < len(saveFileList); i++ {
		if !veriftSave(saveFileList[i]) {
			statusCode = false
			break
		}
	}
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
