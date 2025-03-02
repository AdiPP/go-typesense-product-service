package entity

type ProductSku struct {
	ProductID   int64  `json:"product_id"`
	ProductName string `json:"product_name"`
}

type ProductSkus []ProductSku
