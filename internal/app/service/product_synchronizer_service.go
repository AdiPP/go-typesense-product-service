package service

import (
	"github.com/AdiPP/go-typesense-product-service/internal/app/database/pgsql"
)

type ProductSynchorizerService struct {
	pgsqlDatabase *pgsql.Repository
}

type SyncProductParam struct {
	ProductID int64
}

func (p *ProductSynchorizerService) Sync(param *SyncProductParam) (err error) {

	return
}

func NewProductSynchorizerService(pgsqlDatabase *pgsql.Repository) *ProductSynchorizerService {
	return &ProductSynchorizerService{pgsqlDatabase: pgsqlDatabase}
}
