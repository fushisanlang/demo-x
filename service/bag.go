package service

import (
	"demo-x/model"
	"fmt"
)

var Bags model.BagsStruct

func initBags() {

	nullBags := [64]model.CaseStruct{}
	Bags = model.BagsStruct{
		UseId: 0,
		Bag:   nullBags,
	}
}

func AddGoods(goodsId, goodsCont int) {
	bagUseId := Bags.UseId
	bagId := bagUseId

	Bags.Bag[bagId].GoodID = goodsId
	Bags.Bag[bagId].Count = goodsCont
	Bags.UseId = Bags.UseId + 1
	fmt.Println()
}
