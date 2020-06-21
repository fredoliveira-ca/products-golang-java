package main

import (
	"net/http"

	"github.com/fredoliveira-ca/products-golang-java/product-service/app/grpc/server"
	"github.com/fredoliveira-ca/products-golang-java/product-service/app/web/route"
)

func main() {
	go server.RegisterServer()

	route.LoadRoutes()

	http.ListenAndServe(":8001", nil)
}
