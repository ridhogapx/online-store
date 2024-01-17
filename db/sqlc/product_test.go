package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestCreateCategory(t *testing.T) {
	res, err := testQueries.CreateCategory(context.Background(), "electronic")

	if err != nil {
		t.Error(err)
	}

	fmt.Println("Result:", res)
}

func TestCreateProduct(t *testing.T) {
	res, err := testQueries.CreateProduct(context.Background(), CreateProductParams{
		ProductID:   uuid.New().String(),
		CategoryID:  2,
		ProductName: "Iphone 8",
		Price:       8000000,
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Println("Result:", res)
}

func TestFindingProduct(t *testing.T) {
	// Finding product within id category 2
	res, err := testQueries.FindProductByCategory(context.Background(), 2)

	if err != nil {
		t.Error(err)
	}

	for _, items := range res {
		fmt.Println("------------------------------------------------------")
		fmt.Println("Product ID:", items.ProductID)
		fmt.Println("Product Name:", items.ProductName)
		fmt.Println("Category:", items.CategoryName)
		fmt.Println("Price:", items.Price)
		fmt.Println("------------------------------------------------------")
	}
}

func TestCreateCustomer(t *testing.T) {
	res, err := testQueries.CreateCustomer(context.Background(), CreateCustomerParams{
		// Dummy data for testing queries
		CustomerID:      uuid.NewString(),
		CustomerName:    "Jne",
		CustomerAddress: "Jl. Soda 53",
		Email:           "jane@does.com",
		// Just plain text for testing query
		Password: "myverysecret",
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Println("Result:", res.CustomerName)
}

func TestCreateCart(t *testing.T) {
	res, err := testQueries.CreateCart(context.Background(), CreateCartParams{
		CartID:     uuid.NewString(),
		CustomerID: "1e354f60-168c-4036-8643-e32699215aaa",
		ProductID:  "ee1407bd-254d-4dcf-9903-01ee3542d77a",
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Println("Result:", res.CartID)
}

func TestFindCart(t *testing.T) {
	res, err := testQueries.FindCart(context.Background(), "foo-id-xxx")

	if err != nil {
		t.Error(err)
	}

	for _, items := range res {
		fmt.Println("-------------------------------")
		fmt.Println("Customer Name:", items.CustomerName)
		fmt.Println("Product:", items.ProductName)
	}
}
