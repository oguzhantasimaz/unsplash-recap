package utils

import (
	"unsplash-recap/unsplash"
)

type Response struct {
	StatusCode      int               `json:"statusCode"`
	Message         string            `json:"message"`
	Body            *unsplash.Recap   `json:"body"`
	Headers         map[string]string `json:"headers"`
	IsBase64Encoded bool              `json:"isBase64Encoded"`
}
