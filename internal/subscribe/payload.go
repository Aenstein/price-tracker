package subscribe

type SubscribeRequest struct {
	Link string `json:"link" validate:"required,url"`
}