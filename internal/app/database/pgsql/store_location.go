package pgsql

import "github.com/doug-martin/goqu/v9"

type FindAllStoreLocationsByStoreIdsResponseStoreLocation struct {
	StoreID  int64  `db:"store_id" json:"store_id"`
	CityName string `db:"city_name" json:"city_name"`
}

type FindAllStoreLocationsByStoreIdsResponseStoreLocations []FindAllStoreLocationsByStoreIdsResponseStoreLocation

func (s FindAllStoreLocationsByStoreIdsResponseStoreLocations) FindByStoreID(storeID int64) FindAllStoreLocationsByStoreIdsResponseStoreLocation {
	for _, v := range s {
		if v.StoreID == storeID {
			return v
		}
	}

	return FindAllStoreLocationsByStoreIdsResponseStoreLocation{}
}

type FindAllStoreLocationsByStoreIdsResponse struct {
	StoreLocations FindAllStoreLocationsByStoreIdsResponseStoreLocations
}

func (r *Repository) FindAllStoreLocationsByStoreIds(storeIds []int64) (resp FindAllStoreLocationsByStoreIdsResponse, err error) {
	resp = FindAllStoreLocationsByStoreIdsResponse{
		StoreLocations: FindAllStoreLocationsByStoreIdsResponseStoreLocations{},
	}

	if len(storeIds) == 0 {
		return
	}

	ds := r.database.
		From(goqu.L(findAllStoreLocationsByStoreIdsQuery).As("d")).
		Where(goqu.L("d.store_id IN ?", storeIds))

	var storeLocations FindAllStoreLocationsByStoreIdsResponseStoreLocations
	err = ds.Executor().ScanStructs(&storeLocations)
	if err != nil {
		return
	}

	resp.StoreLocations = storeLocations
	return
}
