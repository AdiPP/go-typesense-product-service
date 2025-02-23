package typesense

type collectionName string

func (c collectionName) string() string {
	return string(c)
}

var (
	productCollectionName collectionName = "products"
)
