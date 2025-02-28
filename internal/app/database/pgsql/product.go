package pgsql

import (
	"errors"
	"time"

	"github.com/AdiPP/go-typesense-product-service/internal/app/entity"
	"github.com/doug-martin/goqu/v9"
)

type FindProductByIdResult struct {
	ProductID              int64      `json:"product_id" db:"product_id"`
	ProductConditionID     int        `json:"product_condition_id" db:"product_condition_id"`
	ShippingTypeID         *int       `json:"shipping_type_id,omitempty" db:"shipping_type_id"`
	CategoryID             int        `json:"category_id" db:"category_id"`
	UomID                  int        `json:"uom_id" db:"uom_id"`
	StoreID                int        `json:"store_id" db:"store_id"`
	ProductName            string     `json:"product_name" db:"product_name"`
	ProductWeight          float64    `json:"product_weight" db:"product_weight"`
	ProductLength          *float64   `json:"product_length,omitempty" db:"product_length"`
	ProductWidth           *float64   `json:"product_width,omitempty" db:"product_width"`
	ProductHeight          *float64   `json:"product_height,omitempty" db:"product_height"`
	ProductDescription     *string    `json:"product_description,omitempty" db:"product_description"`
	ProductSlug            *string    `json:"product_slug,omitempty" db:"product_slug"`
	ProductMetaTitle       *string    `json:"product_meta_title,omitempty" db:"product_meta_title"`
	ProductMetaDescription *string    `json:"product_meta_description,omitempty" db:"product_meta_description"`
	ProductVideoURL        *string    `json:"product_video_url,omitempty" db:"product_video_url"`
	IsInsurance            int        `json:"is_insurance" db:"is_insurance"`
	IsPreorder             *int       `json:"is_preorder,omitempty" db:"is_preorder"`
	DateIn                 time.Time  `json:"date_in" db:"date_in"`
	UserIn                 string     `json:"user_in" db:"user_in"`
	DateUp                 *time.Time `json:"date_up,omitempty" db:"date_up"`
	UserUp                 *string    `json:"user_up,omitempty" db:"user_up"`
	StatusRecord           string     `json:"status_record" db:"status_record"`
	PreorderDay            *int       `json:"preorder_day,omitempty" db:"preorder_day"`
	ProductUploadFileID    *int64     `json:"product_upload_file_id,omitempty" db:"product_upload_file_id"`
	ProductSource          *string    `json:"product_source,omitempty" db:"product_source"`
	ProductPreorderTypeID  *int64     `json:"product_preorder_type_id,omitempty" db:"product_preorder_type_id"`
	PreorderTypeID         *int       `json:"preorder_type_id,omitempty" db:"preorder_type_id"`
	PreorderTypeName       *string    `json:"preorder_type_name,omitempty" db:"preorder_type_name"`
}

func (r *Repository) FindProductById(productId int64) (result FindProductByIdResult, err error) {
	result = FindProductByIdResult{}

	ds := r.database.
		Select(goqu.L("rp.product_id")).
		From(goqu.T(productTable).As("rp")).
		Where(goqu.L("rp.product_id").Eq(productId)).
		Where(goqu.L("rp.status_record").Neq("D"))

	found, err := ds.ScanStruct(&result)
	if err != nil {
		return
	}

	if !found {
		err = errors.New(entity.ErrProductNotFound)
		return
	}
	return
}
