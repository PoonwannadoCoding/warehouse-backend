package model

type OrderInput struct {
	ID         string `json:"idtransaction"`
	ProductId  string `json:"productid"`
	CustomerId string `json:"customerid"`
	Amount     int64  `json:"amount"`
}
