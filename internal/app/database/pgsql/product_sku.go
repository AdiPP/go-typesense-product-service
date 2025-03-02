package pgsql

import (
	"github.com/doug-martin/goqu/v9"
)

type FindAllProductSkuByProductIdsResponseProductSku struct {
	IsDefault         bool    `db:"is_default"`
	ProductID         int64   `db:"product_id"`
	ProductPrice      int64   `db:"product_price"`
	ProductSKUID      int64   `db:"product_sku_id"`
	ProductSKUNumber  string  `db:"product_sku_number"`
	ProductSKUMPN     *string `db:"product_sku_mpn"`
	ProductStatusID   int64   `db:"product_status_id"`
	ProductStatusName string  `db:"product_status_name"`
	ProductWeight     float64 `db:"product_weight"`
}

type FindAllProductSkuByProductIdsResponseProductSkus []FindAllProductSkuByProductIdsResponseProductSku

type FindAllProductSkuByProductIdsResponse struct {
	ProductSkus FindAllProductSkuByProductIdsResponseProductSkus
}

func (r *Repository) FindAllProductSkusByProductIds(productIds []int64) (response *FindAllProductSkuByProductIdsResponse, err error) {
	response = &FindAllProductSkuByProductIdsResponse{
		ProductSkus: FindAllProductSkuByProductIdsResponseProductSkus{},
	}

	if len(productIds) <= 0 {
		return
	}

	ds := r.database.
		From(goqu.L(findAllProductSkusByProductIdsQuery).As("d")).
		Where(goqu.L("d.product_id").In(productIds))

	var productSkus FindAllProductSkuByProductIdsResponseProductSkus
	err = ds.Executor().ScanStructs(&productSkus)
	if err != nil {
		return nil, err
	}

	response = &FindAllProductSkuByProductIdsResponse{
		ProductSkus: productSkus,
	}
	return
}
