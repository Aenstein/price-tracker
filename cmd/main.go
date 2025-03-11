package main

import (
	"fmt"
	"net/http"
	"price-tracker/internal/auth"
)

func main() {
	router := http.NewServeMux()

	auth.NewAuthHandler(router)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Printf("Server is listening on port 8080")
	server.ListenAndServe()
}
