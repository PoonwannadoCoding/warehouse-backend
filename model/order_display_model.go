package model

type OrderRead struct {
	ID           string `json:"idtransaction"`
	ProductName  string `json:"product_name"`
	CustomerName string `json:"name"`
	Amount       int64  `json:"amount"`
	OrderDate    string `json:"orderdate"`
}
