package server

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	userpb "github.com/fredoliveira-ca/products-golang-java/user-service/app/grpc/proto/userpb"
	mapper "github.com/fredoliveira-ca/products-golang-java/user-service/data/mapper"
	repository "github.com/fredoliveira-ca/products-golang-java/user-service/data/repository"
)

const (
	userAddress = "0.0.0.0:50053"
)

type server struct {
}

// RegisterServer sets up the connection server.
func RegisterServer() {
	lis, err := net.Listen("tcp", userAddress)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	userpb.RegisterUserServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// FetchOne goes to the repository and return a user based on the informed identifier.
func (*server) FetchOne(ctx context.Context, req *userpb.UserRequest) (*userpb.UserResponse, error) {
	log.Println("Starting user server...")

	user := repository.FindOne(req.GetUserId())

	return &userpb.UserResponse{
		User: mapper.ToProto(user),
	}, nil
}
