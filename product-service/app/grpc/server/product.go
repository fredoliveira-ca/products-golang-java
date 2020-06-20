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
	productAddress = "0.0.0.0:50051"
)

// Server represents the gRPC server
type Server struct {
}

// RegisterServer sets up the gRPC server and waits for a connection.
func RegisterServer() {
	listen, err := net.Listen("tcp", productAddress)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	newServer := grpc.NewServer()

	productpb.RegisterProductPriceServiceServer(newServer, &Server{})

	if err := newServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// FetchOne returns product details and error
func (*Server) FetchOne(ctx context.Context, req *productpb.ProductPriceRequest) (*productpb.ProductPriceResponse, error) {
	productID := req.ProductId

	product := repository.FindOne(productID)

	return &productpb.ProductPriceResponse{
		ValueInCents: product.PriceInCents,
	}, nil
}
