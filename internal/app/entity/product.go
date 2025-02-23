package entity

type Product struct {
	ProductID   int64  `json:"product_id"`
	ProductName string `json:"product_name"`
}

type Products []Product
