package subscribe

import (
	"fmt"
	"net/http"
	"price-tracker/pkg/request"
	"price-tracker/pkg/response"
)

type SubscribeHandler struct {}

func NewSubscribeHandler(router *http.ServeMux) {
	subscribeHandler := &SubscribeHandler{}
	router.HandleFunc("POST /suscribe", subscribeHandler.Subscribe())
}

func (handler *SubscribeHandler) Subscribe() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[SubscribeRequest](r)
		if err != nil {
			http.Error(w, "Invalide data", http.StatusBadRequest)
		}

		fmt.Println(body)

		response.Json(w, nil, http.StatusOK)
	}
}