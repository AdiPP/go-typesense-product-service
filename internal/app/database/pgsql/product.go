package pgsql

import (
	"github.com/doug-martin/goqu/v9"
	"time"
)

type FindAllProductsByProductIdsResultProduct struct {
	ProductID               int64     `db:"product_id"`
	ProductName             string    `db:"product_name"`
	ProductWeight           float64   `db:"product_weight"`
	UOMID                   int64     `db:"uom_id"`
	UOMName                 string    `db:"uom_name"`
	UOMAbbreviation         string    `db:"uom_abbreviation"`
	IsInsurance             bool      `db:"is_insurance"`
	IsPreorder              bool      `db:"is_preorder"`
	PreorderDay             *int64    `db:"preorder_day"`
	ProductPreorderTypeID   *int64    `db:"product_preorder_type_id"`
	ProductPreorderTypeName *string   `db:"product_preorder_type_name"`
	ProductConditionID      int64     `db:"product_condition_id"`
	ProductConditionName    string    `db:"product_condition_name"`
	ProductDescription      string    `db:"product_description"`
	StoreID                 int64     `db:"store_id"`
	StoreName               string    `db:"store_name"`
	StoreSlug               string    `db:"slug" `
	StoreStatusID           int64     `db:"store_status_id"`
	StoreStatusName         string    `db:"store_status_name"`
	UploadDate              time.Time `db:"upload_date"`
	CategoryID              int64     `db:"category_id"`
	CategoryCode            string    `db:"category_code"`
	CategoryName            string    `db:"category_name"`
	CategorySlug            string    `db:"category_slug"`
	ProductHeight           float64   `db:"product_height"`
	ProductLength           float64   `db:"product_length"`
	ProductWidth            float64   `db:"product_width"`
	ProductMetaTitle        string    `db:"product_meta_title"`
	ProductMetaDescription  string    `db:"product_meta_description"`
	ProductSlug             string    `db:"product_slug"`
	StatusRecord            string    `db:"status_record"`
}

type FindAllProductsByProductIdsResultProducts []FindAllProductsByProductIdsResultProduct

func (p FindAllProductsByProductIdsResultProducts) FindByProductID(productID int64) FindAllProductsByProductIdsResultProduct {
	for _, v := range p {
		if v.ProductID == productID {
			return v
		}
	}

	return FindAllProductsByProductIdsResultProduct{}
}

func (p FindAllProductsByProductIdsResultProducts) GetStoreIDs() []int64 {
	var results []int64

	for _, v := range p {
		results = append(results, v.StoreID)
	}

	return results
}

type FindAllProductsByProductIdsResult struct {
	Products FindAllProductsByProductIdsResultProducts
}

func (r *Repository) FindAllProductByIds(productIds []int64) (result FindAllProductsByProductIdsResult, err error) {
	result = FindAllProductsByProductIdsResult{
		Products: FindAllProductsByProductIdsResultProducts{},
	}

	if len(productIds) <= 0 {
		return
	}

	ds := r.database.
		From(goqu.L(findAllProductsByProductIdsQuery).As("d")).
		Where(goqu.L("d.product_id in ?", productIds))

	var products FindAllProductsByProductIdsResultProducts

	err = ds.Executor().ScanStructs(&products)
	if err != nil {
		return
	}

	result = FindAllProductsByProductIdsResult{
		Products: products,
	}
	return
}
