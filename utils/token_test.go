package utils

import (
	"fmt"
	"testing"
)

var testSecret = []byte("very secret")

func TestGenToken(t *testing.T) {
	payload := Payload{
		CustomerID:   "foo-id-xx",
		CustomerName: "Foo",
		Email:        "foo@mail.com",
		Secret:       testSecret,
	}

	token, err := GenerateToken(payload)

	if err != nil {
		t.Error(err)
	}

	fmt.Println("Token:", token)
}

func TestDecode(t *testing.T) {
	_, err := DecodeToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDU0NzY5NTAsImlzcyI6IlN5bmFwc2lzIE9ubGluZSBTdG9yZSIsImN1c3RvbWVyX2lkIjoiZm9vLWlkLXh4IiwiY3VzdG9tZXJfbmFtZSI6IkZvbyIsImVtYWlsIjoiZm9vQG1haWwuY29tIn0.2bE5pNMBkMMgvqWERPpJT4R3WM89x9debJte1kmd4xE", testSecret)

	if err != nil {
		t.Error(err)
	}
}
