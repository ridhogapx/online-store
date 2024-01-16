package repository

import "github.com/RageNeko26/online-store/entity"

type OnlineStoreRepository interface {
	CreateProduct(*entity.ProductEntity) error
	FindProductByCategory(id string) (*entity.ProductEntity, error)
}
