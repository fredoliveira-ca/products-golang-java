FROM golang:alpine as build-env

ENV GO111MODULE=on

RUN apk update && apk add bash git gcc g++ libc-dev

RUN mkdir /user-service
RUN mkdir -p /user-service/proto 

WORKDIR /user-service

COPY ./app/grpc/proto/userpb/user.pb.go /user-service/proto
COPY ./app/app.go /user-service

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go build -o user-service .

CMD ./user-service