package request

import "net/http"

func HandleBody[T any](r *http.Request) (*T, error) {
	body, err := Decoder[T](r.Body)
	if err != nil {
		return nil, err
	}

	err = Valodator[T](body)
	if err != nil {
		return nil, err
	}

	return &body, nil
}