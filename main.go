package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	//metrics does nothing, just to show how different loggs or metrics can be wrapped
	svc := NewLoggingService(NewMetricsService(&priceGetter{}))

	price, err := svc.GetPrice(context.Background(), "ETH")
	if err != nil {
		fmt.Println("err")
		log.Fatal(err)
	}

	fmt.Println(price)
}
