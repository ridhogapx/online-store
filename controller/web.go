package controller

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
