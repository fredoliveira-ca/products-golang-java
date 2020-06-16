package server

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	productpb "github.com/fredoliveira-ca/products-golang-java/product-service/app/grpc/proto/productpb"
	repository "github.com/fredoliveira-ca/products-golang-java/product-service/data/repository"
)

const (
	host = "0.0.0.0"
	port = "50051"
)

type server struct{}

// Start is ...
func Start() {
	log.Println("Starting product server...")

	listen, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	newServer := grpc.NewServer()
	productpb.RegisterProductPriceServiceServer(newServer, &server{})
	if err := newServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (*server) FetchOne(ctx context.Context, req *productpb.ProductPriceRequest) (*productpb.ProductPriceResponse, error) {
	productID := req.ProductId

	product := repository.FindOne(productID)

	return &productpb.ProductPriceResponse{
		ValueInCents: product.PriceInCents,
	}, nil
}
