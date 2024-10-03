package main

import (
	"context"
	"fmt"
)

// interface to get a price
type PriceGetter interface {
	GetPrice(context.Context, string) (float64, error)
}

// priceGetter implementing the PriceGetter interface
type priceGetter struct {
}

func (s *priceGetter) GetPrice(ctx context.Context, ticker string) (float64, error) {
	//business logic... important not to use types/ do it somewhere else
	price, err := MockPriceApiCall(ctx, ticker)

	if err != nil {
		return 0, fmt.Errorf("failed to fetch price for ticker %s: %v", ticker, err)
	}

	return price, nil
}

var prices = map[string]float64{
	"BTC": 200000,
	"ETH": 3000,
}

func MockPriceApiCall(ctx context.Context, ticker string) (float64, error) {
	price, token := prices[ticker]

	if !token {
		return price, fmt.Errorf("the given coin/ticker: (%s) does not exist/is not supported", ticker)
	}

	return price, nil
}
