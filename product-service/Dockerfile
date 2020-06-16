FROM golang:alpine as build-env

ENV GO111MODULE=on

RUN apk update && apk add bash git gcc g++ libc-dev

RUN mkdir /product-service
RUN mkdir -p /product-service/proto 

WORKDIR /product-service

COPY ./app/grpc/proto/productpb/product.pb.go /product-service/proto
COPY ./app/grpc/proto/discountpb/discount.pb.go /product-service/proto
COPY ./app/app.go /product-service

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go build -o product-service .

CMD ./product-service