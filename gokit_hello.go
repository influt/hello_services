package main

import (
	"context"
	"net/http"
	"encoding/json"

	"github.com/go-kit/kit/endpoint"
	http_t7t "github.com/go-kit/kit/transport/http"
)

type DigitCharService interface {
	Chars() string
}

type digitCharService struct{}

func (digitCharService) Chars() string {
	var ret string = ""
	for i:=65; i<91; i++ {
		ret += string(i)
	}
	return ret
}

type charsRequest struct{}

type charsResponse struct {
	Chars string `json:"chars"`
}

func makeCharsEndpoint(svc DigitCharService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		chars := svc.Chars()
		return charsResponse{chars}, nil
	}
}

func main() {
	svc := digitCharService{}

	charsHandler := http_t7t.NewServer(
		makeCharsEndpoint(svc),
		decodeCharsRequest,
		encodeResponse,
	)

	http.Handle("/chars", charsHandler)
	http.ListenAndServe(":8080", nil)
}

func decodeCharsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request charsRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
