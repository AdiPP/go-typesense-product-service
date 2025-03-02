package pgsql

import (
	"fmt"
	"github.com/doug-martin/goqu/v9"
)

type FindAllProductDiscountsByProductIdsResultProductDiscount struct {
	ProductID          int64  `db:"product_id"`
	ProductSKUID       int64  `db:"product_sku_id"`
	DiscountPercentage int64  `db:"discount_percentage"`
	DiscountPrice      int64  `db:"discount_price"`
	DiscountAmount     int64  `db:"discount_amount"`
	DiscountTypeID     int64  `db:"discount_type_id"`
	DiscountTypeName   string `db:"discount_type_name"`
	DiscountValue      int64  `db:"discount_value"`
}

type FindAllProductDiscountsByProductIdsResultProductDiscounts []FindAllProductDiscountsByProductIdsResultProductDiscount

func (d FindAllProductDiscountsByProductIdsResultProductDiscounts) FindByProductSkuId(productSkuId int64) FindAllProductDiscountsByProductIdsResultProductDiscount {
	for _, v := range d {
		if v.ProductSKUID == productSkuId {
			return v
		}
	}

	return FindAllProductDiscountsByProductIdsResultProductDiscount{}
}

type FindAllProductDiscountsByProductIdsResult struct {
	ProductDiscounts FindAllProductDiscountsByProductIdsResultProductDiscounts
}

func (r *Repository) FindAllProductDiscountsByProductIds(productIDs []int64) (result FindAllProductDiscountsByProductIdsResult, err error) {
	result = FindAllProductDiscountsByProductIdsResult{
		ProductDiscounts: FindAllProductDiscountsByProductIdsResultProductDiscounts{},
	}

	if len(productIDs) == 0 {
		return
	}

	ds := r.database.
		From(goqu.L(fmt.Sprintf("(%s)", findAllProductDiscountsByProductIdsQuery)).As("d")).
		Where(goqu.L("d.product_id").In(productIDs))

	var productDiscounts FindAllProductDiscountsByProductIdsResultProductDiscounts
	err = ds.Executor().ScanStructs(&productDiscounts)
	if err != nil {
		return FindAllProductDiscountsByProductIdsResult{}, err
	}

	result.ProductDiscounts = productDiscounts
	return
}
