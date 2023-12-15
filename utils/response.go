package utils

type Response struct {
	StatusCode      int               `json:"statusCode"`
	Message         string            `json:"message"`
	Body            interface{}       `json:"body"`
	Headers         map[string]string `json:"headers"`
	IsBase64Encoded bool              `json:"isBase64Encoded"`
}

type ErrorResponseBody struct {
	Message string `json:"message"`
}
