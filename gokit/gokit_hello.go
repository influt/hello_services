package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-kit/kit/endpoint"
	http_t7t "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
)

type DigitCharService interface {
	Chars() []rune
}

type digitCharService struct{}

func (digitCharService) Chars() []rune {
	var ret []rune = []rune{}
	for i := 65; i < 91; i++ {
		ret = append(ret, rune(i))
	}
	return ret
}

type charsRequest struct{}

type charsResponse struct {
	Chars string `json:"chars"`
}

func makeCharsEndpoint(svc DigitCharService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		chars := svc.Chars()

		return charsResponse{string(chars)}, nil
	}
}

func main() {

	logger := log.NewLogfmtLogger(os.Stderr)
	svc := digitCharService{}

	charsHandler := http_t7t.NewServer(
		makeCharsEndpoint(svc),
		decodeCharsRequest,
		encodeResponse,
	)

	http.Handle("/chars", charsHandler)
	logger.Log("err", http.ListenAndServe(":8080", nil))
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
