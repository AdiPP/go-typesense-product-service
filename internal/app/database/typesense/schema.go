package typesense

import (
	"github.com/typesense/typesense-go/v3/typesense/api"
	"github.com/typesense/typesense-go/v3/typesense/api/pointer"
)

type schemaGetter struct {
}

func newSchemaGetter() *schemaGetter {
	return &schemaGetter{}
}

func (g *schemaGetter) getProductSkuSchema() *api.CollectionSchema {
	return &api.CollectionSchema{
		Name: productSkuCollectionName.string(),
		Fields: []api.Field{
			{Name: "id", Type: "string"},
			{Name: "batch_id", Type: "string"},
			{Name: "category_codes", Type: "string[]"},
			{Name: "category_ids", Type: "int64[]"},
			{Name: "category_names", Type: "string[]"},
			{Name: "category_slugs", Type: "string[]"},
			{Name: "category_urls", Type: "string[]"},
			{Name: "discount_percentage", Type: "int64"},
			{Name: "discount_type_id", Type: "int64"},
			{Name: "discount_type_name", Type: "string"},
			{Name: "discount_value", Type: "int64"},
			{Name: "formatted_discount_percentage", Type: "string"},
			{Name: "formatted_discount_value", Type: "string"},
			{Name: "formatted_product_price", Type: "string"},
			{Name: "formatted_selling_price", Type: "string"},
			{Name: "is_default", Type: "bool"},
			{Name: "product_id", Type: "int64"},
			{Name: "product_image_urls", Type: "string[]"},
			{Name: "product_name", Type: "string"},
			{Name: "product_price", Type: "int64"},
			{Name: "product_sku_id", Type: "int64"},
			{Name: "product_sku_number", Type: "string"},
			{Name: "product_sold", Type: "int64"},
			{Name: "product_status_id", Type: "int64"},
			{Name: "product_stock", Type: "int64"},
			{Name: "product_weight", Type: "float"},
			{Name: "selling_price", Type: "int64"},
			{Name: "selling_price_end", Type: "int64"},
			{Name: "selling_price_start", Type: "int64"},
			{Name: "uom_abbreviation", Type: "string"},
			{Name: "uom_id", Type: "int64"},
			{Name: "uom_name", Type: "string"},
			{Name: "is_insurance", Type: "bool"},
			{Name: "is_preorder", Type: "bool"},
			{Name: "preorder_day", Type: "int64"},
			{Name: "preorder_type_id", Type: "int64"},
			{Name: "preorder_type_name", Type: "string"},
			{Name: "product_condition_id", Type: "int64"},
			{Name: "product_condition_name", Type: "string"},
			{Name: "product_description", Type: "string"},
			{Name: "product_dimensions", Type: "int64[]"},
			{Name: "review_rating", Type: "int64"},
			{Name: "reviews_count", Type: "int64"},
			{Name: "store_id", Type: "int64"},
			{Name: "store_location_name", Type: "string"},
			{Name: "store_name", Type: "string"},
			{Name: "store_slug", Type: "string"},
			{Name: "store_status_id", Type: "int64"},
			{Name: "store_status_name", Type: "string"},
			{Name: "store_url", Type: "string"},
			{Name: "tags", Type: "string[]"},
			{Name: "timestamp", Type: "int64"},
			{Name: "total_rating", Type: "string"},
			{Name: "upload_date", Type: "int64"},
		},
		DefaultSortingField: pointer.String("selling_price"),
	}
}
