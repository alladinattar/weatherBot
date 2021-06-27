package main

import (
	"context"
	"net/http"
)

type Response struct {
	StatusCode int         `json:"statusCode"`
	Body       interface{} `json:"body"`
}

func Handler(ctx context.Context) (request *http.Request) {

	return
}
