package data

import "demo-x/model"

var GoodsMap map[int]model.GoodsInfoStruct

func initGoods() {
	GoodsMap = map[int]model.GoodsInfoStruct{
		1:   {GoodsId: 1, GoodsType: "心法", GoodsName: "《纳元诀》", GoodsInfo: "入门心法"},
		999: {GoodsId: 999, GoodsType: "丹药", GoodsName: "仙灵丹", GoodsInfo: "增加经验"},
	}
}
