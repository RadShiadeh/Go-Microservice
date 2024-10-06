package main

import (
	"flag"
)

func main() {
	var (
		JSONAPIPort = flag.String("JSONAPIPort", ":3000", "json service running on port 3000")
		GRPCPort    = flag.String("GRPCPort", ":50051", "grpc running on port 50051")
		//ctx         = context.Background()
	)

	flag.Parse()

	svc := NewLoggingService(NewMetricsService(&priceGetter{}))

	go CreateGRPCServerAndRun(*GRPCPort, svc)

	JSONServer := NewJSONAPIServer(*JSONAPIPort, svc)
	JSONServer.Run()
}

//client code and how it would interact with an already running Json_API
// client := client.NewClient("http://localhost:3000")

// price, err := client.GetPrice(context.Background(), "bitcoin", "gbp")

// if err != nil {
// 	log.Fatal(err)
// }

// fmt.Printf("%+v\n", price)

// return

//grpc client
// grpcClient, err := client.NewGrpcClient(*GRPCPort)
// if err != nil {
// 	fmt.Printf("err here")
// 	log.Fatal(err)
// }

// go func() {
// 	time.Sleep(2 * time.Second)
// 	resp, err := grpcClient.GetPrice(ctx, &proto.PriceRequest{Key: "bitcoin", Currency: "eur"})
// 	if err != nil {
// 		fmt.Printf("err here 2")
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("%+v\n", resp)
// }()
