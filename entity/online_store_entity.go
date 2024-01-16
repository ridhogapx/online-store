package entity

type ProductEntity struct {
	ID         string
	CategoryID string
	Name       string
	Price      int64
}

type CategoryEntity struct {
	ID   string
	Name string
}

type CustomerEntity struct {
	ID       string
	Name     string
	Address  string
	Email    string
	Password string
}

type ShoppingCart struct {
	ID         string
	CustomerID string
	ProductID  string
}

type TransactionReport struct {
	ID     string
	CartID string
}
