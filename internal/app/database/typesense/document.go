package typesense

type ProductSkuDocument struct {
	ID                          string                            `json:"id"`
	BatchID                     string                            `json:"batch_id"`
	ProductSKUID                int64                             `json:"product_sku_id"`
	CategoryCode                string                            `json:"category_code"`
	CategoryID                  int64                             `json:"category_id"`
	CategoryName                string                            `json:"category_name"`
	CategorySlug                string                            `json:"category_slug"`
	CategoryURL                 string                            `json:"category_url"`
	DiscountPercentage          int64                             `json:"discount_percentage"`
	DiscountTypeID              int64                             `json:"discount_type_id"`
	DiscountTypeName            string                            `json:"discount_type_name"`
	DiscountValue               int64                             `json:"discount_value"`
	FormattedDiscountPercentage string                            `json:"formatted_discount_percentage"`
	FormattedDiscountValue      string                            `json:"formatted_discount_value"`
	FormattedProductPrice       string                            `json:"formatted_product_price"`
	FormattedSellingPrice       string                            `json:"formatted_selling_price"`
	IsDefault                   bool                              `json:"is_default"`
	ProductID                   int64                             `json:"product_id"`
	ProductImage                string                            `json:"product_image"`
	ProductImageNoBackground    string                            `json:"product_image_no_background"`
	ProductImageNoBackgroundURL string                            `json:"product_image_no_background_url"`
	ProductImageURL             string                            `json:"product_image_url"`
	ProductName                 string                            `json:"product_name"`
	ProductPrice                int64                             `json:"product_price"`
	ProductSkuMPN               string                            `json:"product_sku_mpn"`
	SellerSkuNumber             string                            `json:"seller_sku_number"`
	ProductSkuName              string                            `json:"product_sku_name"`
	ProductSkuNumber            string                            `json:"product_sku_number"`
	ProductSold                 int64                             `json:"product_sold"`
	ProductStatusID             int64                             `json:"product_status_id"`
	ProductStatusName           string                            `json:"product_status_name"`
	ProductStock                int64                             `json:"product_stock"`
	ProductWeight               float64                           `json:"product_weight"`
	SellingPrice                int64                             `json:"selling_price"`
	ProductStatusRecord         string                            `json:"product_status_record"`
	UOMAbbreviation             string                            `json:"uom_abbreviation"`
	UOMID                       int64                             `json:"uom_id"`
	UOMName                     string                            `json:"uom_name"`
	IsInsurance                 bool                              `json:"is_insurance"`
	IsPreorder                  bool                              `json:"is_preorder"`
	PreorderDay                 int64                             `json:"preorder_day"`
	PreorderTypeID              int64                             `json:"preorder_type_id"`
	PreorderTypeName            string                            `json:"preorder_type_name"`
	FormattedPreorderLabel      string                            `json:"formatted_preorder_label"`
	ProductConditionID          int64                             `json:"product_condition_id"`
	ProductConditionName        string                            `json:"product_condition_name"`
	ProductDescription          string                            `json:"product_description"`
	ProductHeight               float64                           `json:"product_height"`
	ProductLength               float64                           `json:"product_length"`
	ProductMetaDescription      string                            `json:"product_meta_description"`
	ProductMetaTitle            string                            `json:"product_meta_title"`
	ProductSlug                 string                            `json:"product_slug"`
	ProductURL                  string                            `json:"product_url"`
	ProductVolumeWeight         float64                           `json:"product_volume_weight"`
	ProductWidth                float64                           `json:"product_width"`
	ReviewMessage               string                            `json:"review_message"`
	ReviewRating                int64                             `json:"review_rating"`
	ReviewsCount                int64                             `json:"reviews_count"`
	StoreID                     int64                             `json:"store_id"`
	StoreLocationName           string                            `json:"store_location_name"`
	StoreName                   string                            `json:"store_name"`
	StoreProductSlug            string                            `json:"store_product_slug"`
	StoreSlug                   string                            `json:"store_slug"`
	StoreStatusID               int64                             `json:"store_status_id"`
	StoreStatusName             string                            `json:"store_status_name"`
	StoreURL                    string                            `json:"store_url"`
	Timestamp                   int64                             `json:"timestamp"`
	TotalRating                 string                            `json:"total_rating"`
	UploadDate                  int64                             `json:"upload_date"`
	Categories                  ProductSKUCategoryDocuments       `json:"categories"`
	ProductImages               ProductSKUProductImageDocuments   `json:"product_images"`
	ProductCourier              ProductSKUProductCourierDocuments `json:"product_couriers"`
}

type ProductSKUCategoryDocument struct {
	CategoryId              int64  `json:"category_id"`
	ParentId                int64  `json:"parent_id"`
	CategoryCode            string `json:"category_code"`
	CategoryName            string `json:"category_name"`
	CategorySlug            string `json:"category_slug"`
	CategoryOrder           int64  `json:"category_order"`
	CategoryImage           string `json:"category_image"`
	CategoryMetaTitle       string `json:"category_meta_title"`
	CategoryMetaDescription string `json:"category_meta_description"`
	CategoryUrl             string `json:"category_url"`
}

type ProductSKUCategoryDocuments []ProductSKUCategoryDocument

type ProductSKUProductImageDocument struct {
	ProductImageId    int64  `json:"product_image_id"`
	ProductId         int64  `json:"product_id"`
	ProductSkuId      int64  `json:"product_sku_id"`
	ProductImageName  string `json:"product_image_name"`
	ProductImage      string `json:"product_image"`
	ProductImageOrder int64  `json:"product_image_order"`
	ImageTypeId       int64  `json:"image_type_id"`
	ProductImageUrl   string `json:"product_image_url"`
}

type ProductSKUProductImageDocuments []ProductSKUProductImageDocument

type ProductSKUStoreDocument struct {
	CreatedDate     string                           `json:"created_date"`
	UpdatedDate     string                           `json:"updated_date"`
	CreatedBy       string                           `json:"created_by"`
	UpdatedBy       string                           `json:"updated_by"`
	Name            string                           `json:"name"`
	ImagePath       string                           `json:"image_path"`
	ImageUrl        string                           `json:"image_url"`
	BannerImagePath string                           `json:"banner_image_path"`
	BannerImageUrl  string                           `json:"banner_image_url"`
	UserId          int64                            `json:"user_id"`
	Description     string                           `json:"description"`
	Slug            string                           `json:"slug"`
	Active          bool                             `json:"active"`
	ActiveInt       int64                            `json:"active_int"`
	IsOpenToggle    bool                             `json:"is_open_toggle"`
	Id              int64                            `json:"id"`
	Locations       ProductSKUStoreLocationDocuments `json:"locations"`
	IsOpen          struct {
		Status       bool   `json:"status"`
		Case         string `json:"case"`
		Reason       string `json:"reason"`
		DateOfStatus string `json:"date_of_status"`
	} `json:"is_open"`
	HasProduct      bool   `json:"has_product"`
	IsOpenInt       int64  `json:"is_open_int"`
	StoreProcessAvg string `json:"store_process_avg"`
	ReviewsCount    string `json:"reviews_count"`
	RatingAvg       string `json:"rating_avg"`
	ProductSold     string `json:"product_sold"`
	Domicile        string `json:"domicile"`
	UserOnline      bool   `json:"user_online"`
	UserLastSeen    string `json:"user_last_seen"`
	StoreStatuses   string `json:"store_statuses"`
}

type ProductSKUStoreLocationDocument struct {
	CreatedDate       string `json:"created_date"`
	UpdatedDate       string `json:"updated_date"`
	CreatedBy         string `json:"created_by"`
	UpdatedBy         string `json:"updated_by"`
	Name              string `json:"name"`
	ContactPerson     string `json:"contact_person"`
	Description       string `json:"description"`
	StoreId           int64  `json:"store_id"`
	Address           string `json:"address"`
	ProvinceId        int64  `json:"province_id"`
	ProvinceName      string `json:"province_name"`
	CityId            int64  `json:"city_id"`
	CityName          string `json:"city_name"`
	DistrictId        int64  `json:"district_id"`
	DistrictName      string `json:"district_name"`
	LowerDistrictId   int64  `json:"lower_district_id"`
	LowerDistrictName string `json:"lower_district_name"`
	PostalCode        string `json:"postal_code"`
	ContactNumber     string `json:"contact_number"`
	IsPrimary         bool   `json:"is_primary"`
	Coordinates       struct {
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	} `json:"coordinates"`
	IsPrimaryInt int64  `json:"is_primary_int"`
	PointAddress string `json:"point_address"`
	Id           int64  `json:"id"`
}

type ProductSKUStoreLocationDocuments []ProductSKUStoreLocationDocument

type ProductSKUProductCourierDocument struct {
	CourierId     int64                                    `json:"courier_id"`
	CourierName   string                                   `json:"courier_name"`
	IsSellerFleet bool                                     `json:"is_seller_fleet"`
	Services      ProductSKUProductCourierServiceDocuments `json:"services"`
}

type ProductSKUProductCourierDocuments []ProductSKUProductCourierDocument

type ProductSKUProductCourierServiceDocument struct {
	CourierServiceId   int64  `json:"courier_service_id"`
	CourierServiceName string `json:"courier_service_name"`
}

type ProductSKUProductCourierServiceDocuments []ProductSKUProductCourierServiceDocument
