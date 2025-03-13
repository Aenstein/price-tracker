package main

import (
	"fmt"
	"net/http"
	"price-tracker/configs"
	"price-tracker/internal/auth"
	"price-tracker/internal/subscribe"
	"price-tracker/internal/user"
	"price-tracker/pkg/db"
)

func main() {
	conf := configs.LoadConfigs()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	userRepository := user.NewUserRepository(db)

	authService := auth.NewAuthService(*userRepository)
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
		AuthService: authService,
	})
	subscribe.NewSubscribeHandler(router)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Printf("Server is listening on port 8080")
	server.ListenAndServe()
}
