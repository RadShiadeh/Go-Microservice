package main

import (
	"flag"
)

func main() {
	//metrics does nothing, just to show how different loggs or metrics can be wrapped
	// client := client.NewClient("http://localhost:3000")

	// price, err := client.GetPrice(context.Background(), "BTC")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%+v\n", price)

	// return

	listenAddr := flag.String("listenaddr", ":3000", "serivce is running")
	svc := NewLoggingService(NewMetricsService(&priceGetter{}))

	server := NewJSONAPIServer(*listenAddr, svc)
	server.Run()

}
