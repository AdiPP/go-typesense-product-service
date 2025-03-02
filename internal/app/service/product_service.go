package service

import (
	"github.com/AdiPP/go-typesense-product-service/internal/app/database/typesense"
	"github.com/AdiPP/go-typesense-product-service/internal/app/entity"
)

type ProductService struct {
	client *typesense.Repository
}

type UpsertProductParam struct {
	ProductID   int64
	ProductName string
}

func (s *ProductService) Upsert(param *UpsertProductParam) (product entity.ProductSku, err error) {
	// product, err = s.client.UpsertProductSku(entity.ProductSku{
	// 	ProductID:   param.ProductID,
	// 	ProductName: param.ProductName,
	// })
	return
}

type DeleteProductParam struct {
	ProductID int64
}

func (s *ProductService) Delete(param *DeleteProductParam) (product entity.ProductSku, err error) {
	product, err = s.client.FindProductSku(param.ProductID)
	if err != nil {
		return
	}

	product, err = s.client.DeleteProductSku(product)
	if err != nil {
		return
	}
	return
}

func NewProductService(client *typesense.Repository) *ProductService {
	o := new(ProductService)
	o.client = client
	return o
}
