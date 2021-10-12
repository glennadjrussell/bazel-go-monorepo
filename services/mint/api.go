package main

import (
	"context"
	"encoding/json"
	"net/http"
	"log"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type MintService interface {
	Mintit([]byte) (string, error)
}

type mintService struct{}

func (mintService) Mintit(payload []byte) (string, error) {
	return "look at me, Im a payload", nil
}

type mintitRequest struct {
	Payload []byte `json:"payload"`
}

type mintitResponse struct {
	Result string `json:"result"`
}

func makeMintitEndpoint(svc MintService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(mintitRequest)
		s, err := svc.Mintit(req.Payload)
		if err != nil {
			return mintitResponse{"Derp"}, nil
		}
		return mintitResponse{s}, nil
	}
}

func main() {
	svc := mintService{}

	mintHandler := httptransport.NewServer(
		makeMintitEndpoint(svc),
		decodeMintitRequest,
		encodeResponse,
	)

	http.Handle("/mint", mintHandler)
	log.Fatal(http.ListenAndServe(":8180", nil))
}

func decodeMintitRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request mintitRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
