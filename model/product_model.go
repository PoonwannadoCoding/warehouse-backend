package model

type Product struct {
	ID       string  `json:"idproduct"`
	Title    string  `json:"product_name"`
	Category string  `json:"category"`
	Price    float32 `json:"price"`
	Amount   int64   `json:"amount"`
}
