package main

import "flag"

func main() {
	//metrics does nothing, just to show how different loggs or metrics can be wrapped
	listenAddr := flag.String("listenaddr", ":3000", "serivce is running")
	svc := NewLoggingService(NewMetricsService(&priceGetter{}))

	server := NewJSONAPIServer(*listenAddr, svc)
	server.Run()

}
