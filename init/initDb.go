package main

import (
	"demo-x/model"
	"demo-x/tools"
	"encoding/csv"
	"encoding/gob"
	"io"
	"os"
	"strconv"
)

func main() {}
func init() {
	initGoogsDb()
	initGoodsType()
}

func initGoodsType() {
	file, err := os.Open("../baseData/goodstype.csv")
	tools.CheckErr(err)
	defer file.Close()
	csvReader := csv.NewReader(file)

	goodsTypeSlice := []model.GoodsType{}

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		tools.CheckErr(err)

		goodsType := model.GoodsType{}
		goodsType.GoodsTypeId, _ = strconv.Atoi(record[0])
		goodsType.GoodsType = record[1]

		goodsTypeSlice = append(goodsTypeSlice, goodsType)
	}
	os.Remove("../data/goodstype.dat")
	file2, err := os.Create("../data/goodstype.dat")
	tools.CheckErr(err)
	defer file2.Close()
	encoder := gob.NewEncoder(file2)
	err = encoder.Encode(goodsTypeSlice)
	tools.CheckErr(err)
}
func initGoogsDb() {
	file, err := os.Open("../baseData/goodsdb.csv")
	tools.CheckErr(err)
	defer file.Close()

	csvReader := csv.NewReader(file)

	goodsSlice := []model.GoodsInfo{}

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		tools.CheckErr(err)

		goodsInfo := model.GoodsInfo{}
		goodsInfo.GoodsId, _ = strconv.Atoi(record[0])
		goodsInfo.GoodsName = record[1]
		goodsInfo.GoodsBuy, _ = strconv.Atoi(record[2])
		goodsInfo.GoodsSale, _ = strconv.Atoi(record[3])
		goodsInfo.GoodsLevel, _ = strconv.Atoi(record[4])
		goodsInfo.GoodsTypeId, _ = strconv.Atoi(record[5])

		goodsSlice = append(goodsSlice, goodsInfo)
	}
	os.Remove("../data/goods.dat")
	file2, err := os.Create("../data/goods.dat")
	tools.CheckErr(err)
	defer file2.Close()
	encoder := gob.NewEncoder(file2)
	err = encoder.Encode(goodsSlice)
	tools.CheckErr(err)
	//file, err = os.Open("data/goods.dat")
	//if err != nil {
	//		panic(err)/
	//	}
	//	defer file.Close()
	//
	//	// 创建一个解码器
	//	decoder := gob.NewDecoder(file)
	//
	//	// 解码数据并存储到新的切片
	//	var newGoodsSlice []model.GoodsInfo
	//	err = decoder.Decode(&newGoodsSlice)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	// 打印新的切片
	//	fmt.Println(newGoodsSlice[0])
}
