package repository

import (
	db "github.com/RageNeko26/online-store/db/sqlc"
	"github.com/RageNeko26/online-store/entity"
)

type onlineStoreImpl struct {
	Q *db.Queries
}

func NewOnlineStoreRepository(q *db.Queries) OnlineStoreRepository {
	return &onlineStoreImpl{
		Q: q,
	}
}

func (repos *onlineStoreImpl) CreateProduct(product *entity.ProductEntity) error {

}

func (repos *onlineStoreImpl) FindProductByCategory(category_id string) (*entity.ProductEntity, error) {

}
