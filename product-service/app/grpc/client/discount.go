package grpc

import (
	"context"
	"log"

	"google.golang.org/grpc"

	discountpb "github.com/fredoliveira-ca/products-golang-java/product-service/app/grpc/proto/discountpb"
)

const (
	host = "0.0.0.0"
	port = "50052"
)

// CalculateDiscount is ...
func CalculateDiscount(productID, userID string) *discountpb.DiscountResponse {
	connection, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer connection.Close()

	cc := discountpb.NewDiscountServiceClient(connection)

	return doUnary(cc, productID, userID)
}

func doUnary(c discountpb.DiscountServiceClient, productID, userID string) *discountpb.DiscountResponse {
	req := &discountpb.DiscountRequest{
		ProductId: productID,
		UserId:    userID,
	}

	res, err := c.Calculate(context.Background(), req)

	defer func() {
		if res := recover(); res != nil {
			err = res.(error)
			log.Fatalf("Could not connect: %v", err)
		}
	}()
	return res
}
