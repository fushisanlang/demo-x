package model

type GoodsInfo struct {
	GoodsId     int
	GoodsTypeId int
	GoodsName   string
	GoodsBuy    int
	GoodsSale   int
	GoodsLevel  int
	GoodsInfo   string
}
type GoodsType struct {
	GoodsTypeId int
	GoodsType   string
}

type ManyGoods struct {
	GoodsId    int
	GoodsCount int
}
