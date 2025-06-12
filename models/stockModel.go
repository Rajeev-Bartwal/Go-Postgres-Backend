package models

type Stock struct{
	StockId int `json:"id"`
	Name string `json:"name"`
	Price int  `json:"price"`
	Company string  `json:"company"`
}