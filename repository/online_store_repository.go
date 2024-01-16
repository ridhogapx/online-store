package repository

import "github.com/RageNeko26/online-store/entity"

type ProductRepository interface {
	CreateProduct(*entity.ProductEntity) error
	FindProductByCategory(id string) (*entity.ProductEntity, error)
}
