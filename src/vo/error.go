package vo

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
}
