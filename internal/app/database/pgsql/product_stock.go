package pgsql

import (
	"github.com/doug-martin/goqu/v9"
)

type FindAllProductSkuStocksByProductIdsResponseProductStock struct {
	ProductID    int64 `db:"product_id" json:"product_id"`
	ProductSkuID int64 `db:"product_sku_id" json:"product_sku_id"`
	ProductStock int64 `db:"product_stock" json:"product_stock"`
}

type FindAllProductSkuStocksByProductIdsResponseProductStocks []FindAllProductSkuStocksByProductIdsResponseProductStock

func (d FindAllProductSkuStocksByProductIdsResponseProductStocks) FindByProductSkuId(productSkuId int64) FindAllProductSkuStocksByProductIdsResponseProductStock {
	for _, v := range d {
		if v.ProductSkuID == productSkuId {
			return v
		}
	}

	return FindAllProductSkuStocksByProductIdsResponseProductStock{}
}

type FindAllProductSkuStocksByProductIdsResponse struct {
	ProductStocks FindAllProductSkuStocksByProductIdsResponseProductStocks
}

func (r *Repository) FindAllProductSkuStocksByProductIds(productIds []int64) (resp FindAllProductSkuStocksByProductIdsResponse, err error) {
	resp = FindAllProductSkuStocksByProductIdsResponse{
		ProductStocks: FindAllProductSkuStocksByProductIdsResponseProductStocks{},
	}

	if len(productIds) == 0 {
		return
	}

	ds := r.database.
		From(goqu.L(findAllProductSkuStocksByProductIdsQuery).As("d")).
		Where(goqu.L("d.product_id IN ?", productIds))

	var productStocks FindAllProductSkuStocksByProductIdsResponseProductStocks
	err = ds.Executor().ScanStructs(&productStocks)
	if err != nil {
		return
	}

	resp.ProductStocks = productStocks
	return
}
