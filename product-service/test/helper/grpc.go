package helper

import (
	context "context"
	fmt "fmt"
	"log"
	"net"

	grpc "google.golang.org/grpc"

	"github.com/fredoliveira-ca/products-golang-java/product-service/app/grpc/proto/discountpb"
)

// MockServer represents the mocked gRPC server
type MockServer struct {
}

// NewMockServer creates a new Discounts gRPC server
func NewMockServer(port int) {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	discountpb.RegisterDiscountServiceServer(grpcServer, &MockServer{})
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

// Calculate returns invoice and error
func (m *MockServer) Calculate(ctx context.Context, req *discountpb.DiscountRequest) (*discountpb.DiscountResponse, error) {
	response := &discountpb.DiscountResponse{Pct: 0, ValueInCents: 0}

	if req.UserId == "u23r-b1r7hday" {
		response = &discountpb.DiscountResponse{
			Pct:          0.05,
			ValueInCents: 0,
		}
	}

	return response, nil
}
