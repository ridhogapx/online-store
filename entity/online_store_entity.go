package entity

type ProductEntity struct {
	ProductID   string
	CategoryID  string
	ProductName string
	Price       int64
}

type CategoryEntity struct {
	CategoryID string
	Name       string
}

type CustomerEntity struct {
	CustomerID   string
	CustomerName string
	Address      string
	Email        string
	Password     string
}

type ShoppingCart struct {
	CartID     string
	CustomerID string
	ProductID  string
}

type TransactionReport struct {
	TransactionID string
	CartID        string
}
