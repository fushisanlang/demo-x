package model

type GoodsInfo struct {
	GoodsId    int
	GoodsType  string
	GoodsName  string
	GoodsBuy   int
	GoodsSale  int
	GoodsLevel int
	GoodsInfo  string
}
type GoodsType struct {
	GoodsTypeId int
	GoodsType   string
}

type ManyGoods struct {
	GoodsId    GoodsInfo
	GoodsCount int
}
