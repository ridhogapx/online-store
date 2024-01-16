package repository

type ProductRepository interface {
	CreateProduct() error
	FindProductByCategory()
}
