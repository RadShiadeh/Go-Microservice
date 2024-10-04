
FROM golang:1.20-alpine AS builder

WORKDIR /Go-Microservice

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /priceGetter

EXPOSE 3000

CMD [ "/priceGetter" ]