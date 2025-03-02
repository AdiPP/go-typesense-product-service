package pgsql

import (
	"github.com/doug-martin/goqu/v9"
)

type FindAllProductSkuVariantsByProductIdsResultProductSkuVariant struct {
	ProductID    int64  `db:"product_id"`
	ProductSKUID int64  `db:"product_sku_id"`
	VariantName  string `db:"variant_name"`
	VariantValue string `db:"variant_value"`
}

type FindAllProductSkuVariantsByProductIdsResultProductSkuVariants []FindAllProductSkuVariantsByProductIdsResultProductSkuVariant

func (p FindAllProductSkuVariantsByProductIdsResultProductSkuVariants) FindAllByProductSKUID(productSKUID int64) FindAllProductSkuVariantsByProductIdsResultProductSkuVariants {
	results := FindAllProductSkuVariantsByProductIdsResultProductSkuVariants{}

	for _, v := range p {
		if v.ProductSKUID == productSKUID {
			results = append(results, v)
		}
	}

	return results
}

func (p FindAllProductSkuVariantsByProductIdsResultProductSkuVariants) GetVariantValues() []string {
	var results []string

	for _, v := range p {
		results = append(results, v.VariantValue)
	}

	return results
}

type FindAllProductSkuVariantsByProductIdsResult struct {
	ProductSKUVariants FindAllProductSkuVariantsByProductIdsResultProductSkuVariants
}

func (r *Repository) FindAllProductSkuVariantsByProductIds(productIds []int64) (result FindAllProductSkuVariantsByProductIdsResult, err error) {
	result = FindAllProductSkuVariantsByProductIdsResult{
		ProductSKUVariants: FindAllProductSkuVariantsByProductIdsResultProductSkuVariants{},
	}

	if len(productIds) <= 0 {
		return
	}

	ds := r.database.
		From(goqu.L(findAllProductSkuVariantsByProductIdsQuery).As("d")).
		Where(goqu.L("d.product_id IN ?", productIds))

	var productSKUVariants FindAllProductSkuVariantsByProductIdsResultProductSkuVariants
	err = ds.Executor().ScanStructs(&productSKUVariants)
	if err != nil {
		return
	}

	result.ProductSKUVariants = productSKUVariants
	return
}
