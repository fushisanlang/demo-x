package model

type GoodsInfo struct {
	GoodsId     int
	GoodsTypeId int
	GoodsName   string
	GoodsBuy    int
	GoodsSale   int
	GoodsLevel  int
}
type GoodsType struct {
	GoodsTypeId int
	GoodsType   string
}
