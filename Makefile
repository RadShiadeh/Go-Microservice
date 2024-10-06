build:
	go build -o priceGetter

run: build
	./priceGetter

proto:
	protoc --proto_path=proto \
  		--go_out=. \
  		--go-grpc_out=. \
  		proto/service.proto

.PHONY: proto
