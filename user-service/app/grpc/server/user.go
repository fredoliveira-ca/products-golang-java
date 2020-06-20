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

// Server represents the gRPC server
type Server struct {
}

// RegisterServer sets up the gRPC server and waits for a connection.
func RegisterServer() {
	lis, err := net.Listen("tcp", userAddress)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	userpb.RegisterUserServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// FetchOne goes to the repository and returns a user based on the informed identifier.
func (*Server) FetchOne(ctx context.Context, req *userpb.UserRequest) (*userpb.UserResponse, error) {
	user := repository.FindOne(req.GetUserId())

	return &userpb.UserResponse{
		User: mapper.ToProto(user),
	}, nil
}
