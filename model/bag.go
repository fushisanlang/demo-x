package model

type CaseStruct struct {
	ID     int
	GoodID int
	Count  int
}

type BagsStruct struct {
	Bag [64]CaseStruct

	UseId int
}
