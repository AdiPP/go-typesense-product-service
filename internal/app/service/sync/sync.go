package sync

import (
	"fmt"
	"github.com/AdiPP/go-typesense-product-service/internal/app/database/pgsql"
	"github.com/AdiPP/go-typesense-product-service/internal/app/database/typesense"
	"github.com/AdiPP/go-typesense-product-service/internal/app/service"
	"github.com/google/uuid"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Service struct {
	pgsqlRepo     *pgsql.Repository
	typesenseRepo *typesense.Repository
}

func NewService(pgsqlRepo *pgsql.Repository, tpRepo *typesense.Repository) *Service {
	return &Service{pgsqlRepo: pgsqlRepo, typesenseRepo: tpRepo}
}

type BatchProductsParam struct {
	ProductIDs []int64
}

func (p *Service) SyncBatchAsync(param *BatchProductsParam) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered from panic in SyncBatchAsync: %v", r)
			}
		}()

		if err := p.SyncBatch(param); err != nil {
			log.Printf("Error in SyncBatchAsync: %v", err)
		}
	}()
}

func (p *Service) SyncBatch(param *BatchProductsParam) (err error) {
	productsResp, err := p.pgsqlRepo.FindAllProductByIds(param.ProductIDs)
	if err != nil {
		return
	}

	fmt.Println(len(productsResp.Products))

	productSkusResp, err := p.pgsqlRepo.FindAllProductSkusByProductIds(param.ProductIDs)
	if err != nil {
		return
	}

	productDiscountsResp, err := p.pgsqlRepo.FindAllProductDiscountsByProductIds(param.ProductIDs)
	if err != nil {
		return
	}

	productSKUVariantsResp, err := p.pgsqlRepo.FindAllProductSkuVariantsByProductIds(param.ProductIDs)
	if err != nil {
		return
	}

	productImagesResp, err := p.pgsqlRepo.FindAllProductImagesByProductIds(param.ProductIDs)
	if err != nil {
		return err
	}

	productSoldResp, err := p.pgsqlRepo.FindAllProductSkuSoldsByProductIds(param.ProductIDs)
	if err != nil {
		return err
	}

	productStockResp, err := p.pgsqlRepo.FindAllProductSkuStocksByProductIds(param.ProductIDs)
	if err != nil {
		return err
	}

	productReviewResp, err := p.pgsqlRepo.FindAllProductSkuReviewsByProductIds(param.ProductIDs)
	if err != nil {
		return err
	}

	storeLocationResp, err := p.pgsqlRepo.FindAllStoreLocationsByStoreIds(productsResp.Products.GetStoreIDs())
	if err != nil {
		return err
	}

	batchID := uuid.NewString()
	batchTimestamp := time.Now()
	products := productsResp.Products

	if len(products) <= 0 {
		return
	}

	storeLocations := storeLocationResp.StoreLocations
	productSKUs := productSkusResp.ProductSkus
	productDiscounts := productDiscountsResp.ProductDiscounts
	productSKUVariants := productSKUVariantsResp.ProductSKUVariants
	productImages := productImagesResp.ProductImages
	productSolds := productSoldResp.ProductSolds
	productStocks := productStockResp.ProductStocks
	productReviews := productReviewResp.ProductReviews

	for _, productSKU := range productSKUs {
		product := products.FindByProductID(productSKU.ProductID)
		productDiscount := productDiscounts.FindByProductSkuId(productSKU.ProductSKUID)
		productSKUVariants := productSKUVariants.FindAllByProductSKUID(productSKU.ProductSKUID)
		productImages := productImages.FindAllByProductID(productSKU.ProductID)
		productSkuImage := productImages.FindByProductSKUID(productSKU.ProductSKUID)
		productSold := productSolds.FindByProductSkuId(productSKU.ProductSKUID)
		productStock := productStocks.FindByProductSkuId(productSKU.ProductSKUID)
		productReview := productReviews.FindByProductSkuId(productSKU.ProductID)
		storeLocation := storeLocations.FindByStoreID(product.StoreID)

		sellingPrince := productSKU.ProductPrice

		if productDiscount.DiscountPrice != 0 {
			sellingPrince = productDiscount.DiscountPrice
		}

		productSKUName := fmt.Sprintf("%s - %s", product.ProductName, strings.Join(productSKUVariants.GetVariantValues(), " | "))

		productImage := productSkuImage

		if productImage.ProductImage == "" {
			productImage = productImages.GetMainImage()
		}

		productVolumeWeight := float64(product.ProductLength*product.ProductWidth*product.ProductHeight) / 6000

		_, err = p.typesenseRepo.UpsertProductSku(
			&typesense.UpsertProductSkuRequest{
				Document: typesense.ProductSkuDocument{
					ID:                          strconv.FormatInt(productSKU.ProductSKUID, 10),
					BatchID:                     batchID,
					ProductSKUID:                productSKU.ProductSKUID,
					CategoryCode:                product.CategoryCode,
					CategoryID:                  product.CategoryID,
					CategoryName:                product.CategoryName,
					CategorySlug:                product.CategorySlug,
					CategoryURL:                 fmt.Sprintf("%s/%s", os.Getenv("BUYER_URL"), product.CategorySlug),
					DiscountPercentage:          productDiscount.DiscountPercentage,
					DiscountTypeID:              productDiscount.DiscountTypeID,
					DiscountTypeName:            productDiscount.DiscountTypeName,
					DiscountValue:               productDiscount.DiscountValue,
					FormattedDiscountPercentage: service.FormattedDiscountPercentage(productDiscount.DiscountPercentage),
					FormattedDiscountValue:      service.FormattedDiscountValue(productDiscount.DiscountValue, productDiscount.DiscountTypeID),
					FormattedProductPrice:       service.FormatRupiahFromInt64(productSKU.ProductPrice),
					FormattedSellingPrice:       service.FormatRupiahFromInt64(sellingPrince),
					IsDefault:                   productSKU.IsDefault,
					ProductID:                   productSKU.ProductID,
					ProductImage:                productImage.ProductImage,
					ProductImageNoBackground:    service.StringFromPtr(productImage.ProductImageNoBackground),
					ProductImageNoBackgroundURL: fmt.Sprintf("%s/%s", os.Getenv("CDN_IMAGE_URL"), service.StringFromPtr(productImage.ProductImageNoBackground)),
					ProductImageURL:             fmt.Sprintf("%s/%s", os.Getenv("CDN_IMAGE_URL"), productImage.ProductImage),
					ProductName:                 product.ProductName,
					ProductPrice:                productSKU.ProductPrice,
					ProductSkuMPN:               service.StringFromPtr(productSKU.ProductSKUMPN),
					SellerSkuNumber:             service.StringFromPtr(productSKU.ProductSKUMPN),
					ProductSkuName:              productSKUName,
					ProductSkuNumber:            productSKU.ProductSKUNumber,
					ProductSold:                 productSold.Total,
					ProductStatusID:             productSKU.ProductStatusID,
					ProductStatusName:           productSKU.ProductStatusName,
					ProductStock:                productStock.ProductStock,
					ProductWeight:               product.ProductWeight,
					SellingPrice:                sellingPrince,
					ProductStatusRecord:         product.StatusRecord,
					UOMAbbreviation:             product.UOMAbbreviation,
					UOMID:                       product.UOMID,
					UOMName:                     product.UOMName,
					IsInsurance:                 product.IsInsurance,
					IsPreorder:                  product.IsPreorder,
					PreorderDay:                 service.Int64FromPtr(product.PreorderDay),
					PreorderTypeID:              service.Int64FromPtr(product.ProductPreorderTypeID),
					PreorderTypeName:            service.StringFromPtr(product.ProductPreorderTypeName),
					FormattedPreorderLabel:      fmt.Sprintf("%v %s", service.Int64FromPtr(product.PreorderDay), product.ProductPreorderTypeName),
					ProductConditionID:          product.ProductConditionID,
					ProductConditionName:        product.ProductConditionName,
					ProductDescription:          product.ProductDescription,
					ProductHeight:               product.ProductHeight,
					ProductLength:               product.ProductLength,
					ProductMetaDescription:      product.ProductMetaDescription,
					ProductMetaTitle:            product.ProductMetaTitle,
					ProductSlug:                 product.ProductSlug,
					ProductURL:                  fmt.Sprintf("%s/p/%s/%s", os.Getenv("BUYER_URL"), product.StoreSlug, product.ProductSlug),
					ProductVolumeWeight:         productVolumeWeight,
					ProductWidth:                product.ProductWidth,
					ReviewMessage:               productReview.ScoreMessage,
					ReviewRating:                productReview.AverageScore,
					ReviewsCount:                productReview.ReviewsCount,
					StoreID:                     product.StoreID,
					StoreLocationName:           storeLocation.CityName,
					StoreName:                   product.StoreName,
					StoreProductSlug:            fmt.Sprintf("%s/%s", product.StoreSlug, product.ProductSlug),
					StoreSlug:                   product.StoreSlug,
					StoreStatusID:               product.StoreStatusID,
					StoreStatusName:             product.StoreStatusName,
					StoreURL:                    fmt.Sprintf("%s/p/%s", os.Getenv("BUYER_URL"), product.StoreSlug),
					Timestamp:                   batchTimestamp.Unix(),
					TotalRating:                 service.FormatShortNumber(float64(productReview.ReviewsCount), 1),
					UploadDate:                  product.UploadDate.Unix(),
					Categories:                  typesense.ProductSKUCategoryDocuments{},
					ProductImages:               typesense.ProductSKUProductImageDocuments{},
					ProductCourier:              typesense.ProductSKUProductCourierDocuments{},
				},
			},
		)

		if err != nil {
			log.Printf("Sync Product SKU %v Failed. Error: %s", productSKU.ProductID, err.Error())
			continue
		}

		log.Printf("Sync Product SKU %v Succeed.", productSKU.ProductID)
	}

	return
}
