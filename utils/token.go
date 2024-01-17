package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	jwt.StandardClaims
	CustomerID   string `json:"customer_id"`
	CustomerName string `json:"customer_name"`
	Email        string `json:"email"`
}

type Payload struct {
	CustomerID   string
	CustomerName string
	Email        string
	Secret       []byte
}

func GenerateToken(payload Payload) (string, error) {
	exp := time.Duration(1) * time.Hour

	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "Synapsis Online Store",
			ExpiresAt: time.Now().Add(exp).Unix(),
		},
		CustomerID:   payload.CustomerID,
		CustomerName: payload.CustomerName,
		Email:        payload.Email,
	}

	signingMethod := jwt.SigningMethodHS256

	token := jwt.NewWithClaims(
		signingMethod,
		claims,
	)

	signedToken, err := token.SignedString(payload.Secret)

	if err != nil {
		return "", err
	}

	return signedToken, nil

}
