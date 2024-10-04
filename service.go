package main

import (
	"context"
	"fmt"
	"time"
)

// interface to get a price
type PriceGetter interface {
	GetPrice(context.Context, string) (float64, error)
}

// priceGetter implementing the PriceGetter interface
type priceGetter struct {
}

func (s *priceGetter) GetPrice(ctx context.Context, key string) (float64, error) {
	//business logic... important not to use types/ do it somewhere else
	price, err := MockPriceApiCall(ctx, key)

	if err != nil {
		return 0, fmt.Errorf("failed to fetch price for key %s: %v", key, err)
	}

	return price, nil
}

var prices = map[string]float64{
	"BTC": 200000,
	"ETH": 3000,
} //in mem data, need to replace with an actual API

func MockPriceApiCall(ctx context.Context, key string) (float64, error) {
	//mimic the http roundtrip, needs to be replaced with gRPC and actual APIs
	time.Sleep(100 * time.Millisecond)

	price, token := prices[key]

	if !token {
		return price, fmt.Errorf("the given coin/key: (%s) does not exist/is not supported", key)
	}

	return price, nil
}
