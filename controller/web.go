package controller

import db "github.com/RageNeko26/online-store/db/sqlc"

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type CreateCartRequest struct {
	ProductID string `json:"product_id"`
}

type CartResponse struct {
	TotalQuantity int64            `json:"total_quantity"`
	Products      []db.FindCartRow `json:"products"`
}

type CreateCategoryRequest struct {
	CategoryName string `json:"category_name"`
}
