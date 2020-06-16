package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/fredoliveira-ca/products-golang-java/user-service/app/grpc/server"
)

func main() {
	loadEnv()
	server.Start()
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
