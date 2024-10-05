package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// interface to get a price
type PriceGetter interface {
	GetPrice(context.Context, string, string) (float64, error)
}

// priceGetter implementing the PriceGetter interface
type priceGetter struct{}

func (s *priceGetter) GetPrice(ctx context.Context, key string, currency string) (float64, error) {
	//business logic... important not to use types/ do it somewhere else
	price, err := API(ctx, key, currency)

	if err != nil {
		return 0, fmt.Errorf("failed to fetch price for key %s: %v", key, err)
	}

	return price, nil
}

func API(ctx context.Context, key string, curr string) (float64, error) {

	price, err := MakeAPICall(key, curr)
	if err != nil {
		return 0, fmt.Errorf("sth went wrong making the api call: %s", err)
	}

	return price, nil
}

type PriceResponse map[string]map[string]float64

func MakeAPICall(key string, curr string) (float64, error) {

	url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=%s", key, curr)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return 0.0, err
	}
	defer resp.Body.Close()

	var priceResponse PriceResponse
	if err := json.NewDecoder(resp.Body).Decode(&priceResponse); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return 0.0, err
	}

	if price, exist := priceResponse[key][curr]; exist {
		return price, nil
	}

	return 0.0, fmt.Errorf("key error, %s does not exist or %s does not exist in the %s currency", key, key, curr)
}
