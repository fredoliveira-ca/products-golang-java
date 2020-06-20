package test

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"google.golang.org/grpc"

	api "github.com/fredoliveira-ca/products-golang-java/product-service/app/grpc/proto/discountpb"
)

const (
	discountAddress = "0.0.0.0:50052"
)

var (
	productClient api.DiscountServiceClient
)

// Fazer para produtos

func TestMain(m *testing.M) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(discountAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	productClient = api.NewDiscountServiceClient(conn)

	// Run tests
	os.Exit(m.Run())
}

func TestPayment(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &api.DiscountRequest{
		ProductId: "1",
		UserId:    "1",
	}

	r, err := productClient.Calculate(ctx, req)
	if err != nil {
		log.Fatalf("Payment error: %v", err)
	}
	if r.Pct == 0 {
		log.Fatal("Discount failed")
	}
	log.Println(r.Pct)
}
