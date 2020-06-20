package grpc

import (
	"context"
	"log"

	"google.golang.org/grpc"

	api "github.com/fredoliveira-ca/products-golang-java/product-service/app/grpc/proto/discountpb"
)

const (
	discountAddress = "0.0.0.0:50052"
)

// CalculateDiscount is ...
func CalculateDiscount(productID, userID string) *api.DiscountResponse {
	conn, err := grpc.Dial(discountAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	discountClient := api.NewDiscountServiceClient(conn)

	return doUnary(discountClient, productID, userID)
}

func doUnary(c api.DiscountServiceClient, productID, userID string) *api.DiscountResponse {
	req := &api.DiscountRequest{
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
