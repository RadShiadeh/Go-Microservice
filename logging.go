package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type loggingService struct {
	next priceGetter
}

func (s *loggingService) LogGetPrice(ctx context.Context, key string) (price float64, err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"time":  time.Since(start),
			"err":   err,
			"price": price,
		}).Info("GetPrice")
	}(time.Now())

	return s.next.GetPrice(ctx, key)
}
