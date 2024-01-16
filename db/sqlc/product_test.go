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
