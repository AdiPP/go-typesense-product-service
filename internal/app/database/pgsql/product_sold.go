package pgsql

import (
	"fmt"
	"github.com/doug-martin/goqu/v9"
)

type FindAllProductSkuSoldsByProductIdsResponseProductSold struct {
	ProductID    int64 `json:"product_id"`
	ProductSKUID int64 `json:"product_sku_id"`
	Total        int64 `json:"total"`
}

type FindAllProductSkuSoldsByProductIdsResponseProductSolds []FindAllProductSkuSoldsByProductIdsResponseProductSold

func (d FindAllProductSkuSoldsByProductIdsResponseProductSolds) FindByProductSkuId(productSkuId int64) FindAllProductSkuSoldsByProductIdsResponseProductSold {
	for _, v := range d {
		if v.ProductSKUID == productSkuId {
			return v
		}
	}

	return FindAllProductSkuSoldsByProductIdsResponseProductSold{}
}

type FindAllProductSkuSoldsByProductIdsResponse struct {
	ProductSolds FindAllProductSkuSoldsByProductIdsResponseProductSolds
}

func (r *Repository) FindAllProductSkuSoldsByProductIds(productIds []int64) (resp FindAllProductSkuSoldsByProductIdsResponse, err error) {
	resp = FindAllProductSkuSoldsByProductIdsResponse{
		ProductSolds: FindAllProductSkuSoldsByProductIdsResponseProductSolds{},
	}

	if len(productIds) == 0 {
		return
	}

	ds := r.database.
		From(goqu.L(fmt.Sprintf("(%s)", findAllProductSkuSoldsByProductIdsQuery)).As("d")).
		Where(goqu.L("d.product_id IN ?", productIds))

	var productSolds FindAllProductSkuSoldsByProductIdsResponseProductSolds
	err = ds.Executor().ScanStructs(&productSolds)
	if err != nil {
		return
	}

	resp.ProductSolds = productSolds
	return
}
