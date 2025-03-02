package typesense

type collectionName string

func (c collectionName) string() string {
	return string(c)
}

var (
	productSkuCollectionName collectionName = "rns_product_skus_001"
)
