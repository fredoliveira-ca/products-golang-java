package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/fredoliveira-ca/products-golang-java/product-service/app/grpc/server"
	"github.com/fredoliveira-ca/products-golang-java/product-service/app/web/route"
)

func startServerGRPC() {
	server.Start()
}

func main() {
	loadEnv()

	go startServerGRPC()

	route.LoadRoutes()
	http.ListenAndServe(":8001", nil)
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
