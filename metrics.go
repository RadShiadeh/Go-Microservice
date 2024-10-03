package main

import (
	"context"
	"fmt"
)

type metricsService struct {
	next PriceGetter
}

func NewMetricsService(next PriceGetter) PriceGetter {
	return &metricsService{
		next: next,
	}
}

func (s *metricsService) GetPrice(ctx context.Context, key string) (price float64, err error) {
	fmt.Println("wrapper with some metrics, testing only")
	return s.next.GetPrice(ctx, key)
}
