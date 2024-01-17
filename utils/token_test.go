package utils

import (
	"fmt"
	"testing"
)

func TestGenToken(t *testing.T) {
	payload := Payload{
		CustomerID:   "foo-id-xx",
		CustomerName: "Foo",
		Email:        "foo@mail.com",
		Secret:       []byte("very secret"),
	}

	token, err := GenerateToken(payload)

	if err != nil {
		t.Error(err)
	}

	fmt.Println("Token:", token)
}
