package model

type Order struct {
	ID           string `json:"idtransaction"`
	ProductName  string `json:"product_name"`
	CustomerName string `json:"name"`
	Amount       int64  `json:"amount"`
}
