package service

import (
	"fmt"

	"github.com/AdiPP/go-typesense-product-service/internal/app/database/pgsql"
)

type ProductSynchorizerService struct {
	pgsqlRepo *pgsql.Repository
}

type SyncProductBatchParam struct {
	ProductIDs []int64
}

func (p *ProductSynchorizerService) SyncBatch(param *SyncProductBatchParam) (err error) {
	findAllProductByIdsResult, err := p.pgsqlRepo.FindAllProductByIds(param.ProductIDs)
	if err != nil {
		return
	}

	fmt.Println(findAllProductByIdsResult)

	return
}

func NewProductSynchorizerService(pgsqlDatabase *pgsql.Repository) *ProductSynchorizerService {
	return &ProductSynchorizerService{pgsqlRepo: pgsqlDatabase}
}
