package request

type CardRequest struct {
	CardNumber      string `json:"card_number" validate:"required"`
	ExpirationMonth int    `json:"expiration_month" validate:"required"`
	ExpirationYear  int    `json:"expiration_year" validate:"required,gte=2021,lte=9999"`
}
