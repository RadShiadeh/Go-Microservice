package main

import (
	"context"
	"encoding/json"
	"get-price/types"
	"math/rand"
	"net/http"
)

type JSONAPIServer struct {
	listenAddr string
	svc        PriceGetter
}

type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error

func (s *JSONAPIServer) Run() {
	http.HandleFunc("/", makeHTTPAPIFunc(s.HandleFetchPrice))
	http.ListenAndServe(s.listenAddr, nil)
}

func NewJSONAPIServer(listenAddr string, svc PriceGetter) *JSONAPIServer {
	return &JSONAPIServer{
		listenAddr: listenAddr,
		svc:        svc,
	}
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
		return nil
	}

	priceResponse := types.PriceResponse{
		Price: price,
		Key:   key,
	}

	return WriteJSON(w, http.StatusAccepted, &priceResponse)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
