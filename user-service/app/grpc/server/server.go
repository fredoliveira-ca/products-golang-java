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
	host     = "0.0.0.0"
	port     = "50053"
)

type server struct{}

// Start is ...
func Start() {
	listen, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	newServer := grpc.NewServer()
	
	userpb.RegisterUserServiceServer(newServer, &server{})
	if err := newServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (*server) FetchOne(ctx context.Context, req *userpb.UserRequest) (*userpb.UserResponse, error) {
	user := repository.FindOne(req.GetUserId())

	return &userpb.UserResponse{
		User: mapper.ToProto(user),
	}, nil
}
