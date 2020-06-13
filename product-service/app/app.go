package main

import (
	"net/http"

	_ "github.com/lib/pq"

	"github.com/fredoliveira-ca/products-golang-java/product-service/app/grpc/server"
	"github.com/fredoliveira-ca/products-golang-java/product-service/app/web/route"
)

func startServerGRPC() {
	server.Start()
}

func main() {
	go startServerGRPC()

	route.LoadRoutes()
	http.ListenAndServe(":8001", nil)
}
