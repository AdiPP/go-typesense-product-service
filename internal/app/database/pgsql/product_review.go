package pgsql

import "github.com/doug-martin/goqu/v9"

type FindAllProductSkuReviewsByProductIdsResponseProductReview struct {
	ProductID    int64  `db:"product_id" json:"product_id"`
	ProductSkuID int64  `db:"product_sku_id" json:"product_sku_id"`
	AverageScore int64  `db:"average_score" json:"average_score"`
	ReviewsCount int64  `db:"reviews_count" json:"reviews_count"`
	ScoreMessage string `db:"score_message" json:"score_message"`
}

type FindAllProductSkuReviewsByProductIdsResponseProductReviews []FindAllProductSkuReviewsByProductIdsResponseProductReview

func (d FindAllProductSkuReviewsByProductIdsResponseProductReviews) FindByProductSkuId(productSkuId int64) FindAllProductSkuReviewsByProductIdsResponseProductReview {
	for _, v := range d {
		if v.ProductSkuID == productSkuId {
			return v
		}
	}

	return FindAllProductSkuReviewsByProductIdsResponseProductReview{}
}

type FindAllProductSkuReviewsByProductIdsResponse struct {
	ProductReviews FindAllProductSkuReviewsByProductIdsResponseProductReviews
}

func (r *Repository) FindAllProductSkuReviewsByProductIds(productIds []int64) (resp FindAllProductSkuReviewsByProductIdsResponse, err error) {
	resp = FindAllProductSkuReviewsByProductIdsResponse{
		ProductReviews: FindAllProductSkuReviewsByProductIdsResponseProductReviews{},
	}

	if len(productIds) == 0 {
		return
	}

	ds := r.database.
		From(goqu.L(findAllProductSkuReviewsByProductIdsQuery).As("d")).
		Where(goqu.L("d.product_id IN ?", productIds))

	var productSkuReviews FindAllProductSkuReviewsByProductIdsResponseProductReviews
	err = ds.Executor().ScanStructs(&productSkuReviews)
	if err != nil {
		return
	}

	resp.ProductReviews = productSkuReviews
	return
}
