package response

type CardResponse struct {
	Valid bool       `json:"valid"`
	Error *CardError `json:"error,omitempty"`
}

type CardError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
