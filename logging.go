package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type loggingService struct {
	next PriceGetter
}

func NewLoggingService(next PriceGetter) PriceGetter {
	return &loggingService{
		next: next,
	}
}

func (s *loggingService) GetPrice(ctx context.Context, key string) (price float64, err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"requestID": ctx.Value("requestID"),
			"time":      time.Since(start),
			"err":       err,
			"price":     price,
		}).Info("GetPrice")
	}(time.Now())

	return s.next.GetPrice(ctx, key)
}
