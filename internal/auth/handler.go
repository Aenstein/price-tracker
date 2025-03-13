package auth

import (
	"net/http"
	"price-tracker/configs"
	"price-tracker/pkg/jwt"
	"price-tracker/pkg/request"
	"price-tracker/pkg/response"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

type AuthHandler struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (a *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[LoginRequest](r)
		if err != nil {
			return
		}
		email, err := a.AuthService.Login(body.Email, body.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		token, err := jwt.NewJWT(a.Config.Auth.Secret).Create(jwt.JWTData{Email: email})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := LoginResponse{
			Token: token,
		}

		response.Json(w, data, http.StatusOK)
	}
}

func (a *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[RegisterRequest](r)
		if err != nil {
			return
		}
		email, err := a.AuthService.Register(body.Email, body.Password, body.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		token, err := jwt.NewJWT(a.Config.Auth.Secret).Create(jwt.JWTData{Email: email})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := RegisterResponse{
			Token: token,
		}

		response.Json(w, data, http.StatusCreated)
	}
}
