package service

import (
	"demo-x/model"
	"encoding/gob"
	"os"
)

var NewGoodsSlice []model.GoodsInfo

func initData() {
	file, err := os.Open("data/goods.dat")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 创建一个解码器
	decoder := gob.NewDecoder(file)

	// 解码数据并存储到新的切片

	err = decoder.Decode(&NewGoodsSlice)
	if err != nil {
		panic(err)
	}

	// 打印新的切片
	//fmt.Println(newGoodsSlice[0])
}