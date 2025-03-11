package auth

import (
	"net/http"
	"price-tracker/pkg/request"
	"price-tracker/pkg/response"
)

type AuthHandler struct{}

func NewAuthHandler(router *http.ServeMux) {
	authHandler := &AuthHandler{}
	router.HandleFunc("POST /auth/login", authHandler.Login())
	router.HandleFunc("POST /auth/register", authHandler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := request.HandleBody[LoginRequest](w, r)
		if err != nil {
			http.Error(w, "Invalide data", http.StatusBadRequest)
		}

		data := LoginResponse{
			Token: "token",
		}
		response.Json(w, data, http.StatusOK)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := request.HandleBody[RegisterRequest](w, r)
		if err != nil {
			http.Error(w, "Invalide data", http.StatusBadRequest)
		}

		data := LoginResponse{
			Token: "token",
		}
		response.Json(w, data, http.StatusOK)
	}
}
