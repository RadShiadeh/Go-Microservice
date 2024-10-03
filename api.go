package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
)

type JSONAPIServer struct {
	svc PriceGetter
}

type PriceResponse struct {
	Key   string  `json:"key"`
	Price float64 `json:"price"`
}

type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error

func (s *JSONAPIServer) Run() {
	//http.HandleFunc("/")
}

func makeHTTPAPIFunc(apiFn APIFunc) http.HandlerFunc {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestID", rand.Intn(1000000))
	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFn(ctx, w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, map[string]any{"error": err})
		}
	}
}

func (s *JSONAPIServer) HandleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	key := r.URL.Query().Get("key")

	price, err := s.svc.GetPrice(ctx, key)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	priceResponse := PriceResponse{
		Price: price,
		Key:   key,
	}

	return WriteJSON(w, http.StatusAccepted, &priceResponse)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
