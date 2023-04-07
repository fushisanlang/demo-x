package service

/*
背包逻辑：

背包有  useid 和 bag 两个数据
useid 是现在使用的id，用于快速存放
bag是 通过 64 个 model.BagsStruct 建立了一个列表，代表64个背包，每个都是单独的格子
每个格子都是一个 model.BagsStruct 结构体
*/

func AddGoods(goodsId, goodsCont int) {
	bagUseId := GameData.UserBag.UseId

	GameData.UserBag.Bag[bagUseId].GoodID = goodsId
	GameData.UserBag.Bag[bagUseId].Count = goodsCont
	GameData.UserBag.UseId = GameData.UserBag.UseId + 1
}
