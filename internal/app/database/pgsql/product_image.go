package pgsql

import (
	"github.com/doug-martin/goqu/v9"
)

type FindAllProductImagesByProductIdsResponseProductImage struct {
	ProductID                    int64   `db:"product_id"`
	ProductSKUID                 int64   `db:"product_sku_id"`
	ImageTypeID                  int64   `db:"image_type_id"`
	ProductImage                 string  `db:"product_image"`
	ProductImageName             string  `db:"product_image_name"`
	ProductImageNoBackground     *string `db:"product_image_no_background"`
	ProductImageNoBackgroundName *string `db:"product_image_no_background_name"`
}

type FindAllProductImagesByProductIdsResponseProductImages []FindAllProductImagesByProductIdsResponseProductImage

func (r FindAllProductImagesByProductIdsResponseProductImages) FindAllByProductID(productID int64) FindAllProductImagesByProductIdsResponseProductImages {
	results := FindAllProductImagesByProductIdsResponseProductImages{}

	for _, v := range r {
		if v.ProductID == productID {
			results = append(results, v)
		}
	}

	return results
}

func (r FindAllProductImagesByProductIdsResponseProductImages) FindByProductSKUID(productSKUID int64) FindAllProductImagesByProductIdsResponseProductImage {
	for _, v := range r {
		if v.ProductSKUID == productSKUID {
			return v
		}
	}

	return FindAllProductImagesByProductIdsResponseProductImage{}
}

func (r FindAllProductImagesByProductIdsResponseProductImages) GetMainImage() FindAllProductImagesByProductIdsResponseProductImage {
	for _, v := range r {
		if v.ImageTypeID == 1 {
			return v
		}
	}

	return FindAllProductImagesByProductIdsResponseProductImage{}
}

type FindAllProductImagesByProductIdsResponse struct {
	ProductImages FindAllProductImagesByProductIdsResponseProductImages
}

func (r *Repository) FindAllProductImagesByProductIds(productIds []int64) (resp FindAllProductImagesByProductIdsResponse, err error) {
	resp = FindAllProductImagesByProductIdsResponse{
		ProductImages: FindAllProductImagesByProductIdsResponseProductImages{},
	}

	ds := r.database.
		From(goqu.L(findAllProductImagesByProductIdsQuery).As("d")).
		Where(goqu.L("d.product_id IN ?", productIds))

	var productImages FindAllProductImagesByProductIdsResponseProductImages
	err = ds.Executor().ScanStructs(&productImages)
	if err != nil {
		return FindAllProductImagesByProductIdsResponse{}, err
	}

	resp.ProductImages = productImages
	return
}
