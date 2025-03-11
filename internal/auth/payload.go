package auth

type LoginRequest struct {
	Login string `json:"login" validate:"required, email"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Login string `json:"login" validate:"required, email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}