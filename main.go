package main

import (
	"flag"
)

func main() {
	var (
		JSONAPIPort = flag.String("JSONAPIPort", ":3000", "json service running on port 3000")
		GRPCPort    = flag.String("GRPCPort", ":4000", "grpc running on port 4000")
	)

	flag.Parse()

	svc := NewLoggingService(NewMetricsService(&priceGetter{}))

	JSONServer := NewJSONAPIServer(*JSONAPIPort, svc)
	JSONServer.Run()

	go CreateGRPCServerAndRun(*GRPCPort, svc)
}

//client code and how it would interact with an already running Json_API
// client := client.NewClient("http://localhost:3000")

// price, err := client.GetPrice(context.Background(), "bitcoin", "gbp")

// if err != nil {
// 	log.Fatal(err)
// }

// fmt.Printf("%+v\n", price)

// return
