package main

import (
	"demo-x/model"
	"encoding/csv"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("baseData/goodsdb.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	goodsSlice := []model.GoodsInfo{}

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		goodsInfo := model.GoodsInfo{}
		goodsInfo.GoodsId, _ = strconv.Atoi(record[0])
		goodsInfo.GoodsName = record[1]
		goodsInfo.GoodsBuy, _ = strconv.Atoi(record[2])
		goodsInfo.GoodsSale, _ = strconv.Atoi(record[3])
		goodsInfo.GoodsLevel, _ = strconv.Atoi(record[4])

		goodsSlice = append(goodsSlice, goodsInfo)
	}

	//file2, err := os.Create("data/goods.dat")
	//if err != nil {
	//	panic(err)
	//}
	//defer file2.Close()
	//encoder := gob.NewEncoder(file2)
	//err = encoder.Encode(goodsSlice)
	//if err != nil {
	//	panic(err)
	//}
	file, err = os.Open("data/goods.dat")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 创建一个解码器
	decoder := gob.NewDecoder(file)

	// 解码数据并存储到新的切片
	var newGoodsSlice []model.GoodsInfo
	err = decoder.Decode(&newGoodsSlice)
	if err != nil {
		panic(err)
	}

	// 打印新的切片
	fmt.Println(newGoodsSlice[0])
}
