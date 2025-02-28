package service

import (
	"github.com/AdiPP/go-typesense-product-service/internal/app/database/typesense"
	"github.com/AdiPP/go-typesense-product-service/internal/app/entity"
)

type ProductService struct {
	client *typesense.Client
}

type UpsertProductParam struct {
	ProductID   int64
	ProductName string
}

func (s *ProductService) Upsert(param *UpsertProductParam) (product entity.Product, err error) {
	product, err = s.client.UpsertProduct(entity.Product{
		ProductID:   param.ProductID,
		ProductName: param.ProductName,
	})
	return
}

type DeleteProductParam struct {
	ProductID int64
}

func (s *ProductService) Delete(param *DeleteProductParam) (product entity.Product, err error) {
	product, err = s.client.FindProduct(param.ProductID)
	if err != nil {
		return
	}

	product, err = s.client.DeleteProduct(product)
	if err != nil {
		return
	}
	return
}

func NewProductService(client *typesense.Client) *ProductService {
	o := new(ProductService)
	o.client = client
	return o
}
