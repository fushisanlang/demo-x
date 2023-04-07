package model

type GoodsInfoStruct struct {
	GoodsId    int
	GoodsType  string
	GoodsName  string
	GoodsBuy   int
	GoodsSale  int
	GoodsLevel int
	GoodsInfo  string
}
type GoodsTypeStruct struct {
	GoodsTypeId int
	GoodsType   string
}

type ManyGoodsStruct struct {
	GoodsId    GoodsInfoStruct
	GoodsCount int
}
