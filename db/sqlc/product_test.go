package db

import (
	"context"
	"fmt"
	"testing"
)

func TestCreateCategory(t *testing.T) {
	res, err := testQueries.CreateCategory(context.Background(), "electronic")

	if err != nil {
		t.Error(err)
	}

	fmt.Println("Result:", res)
}
