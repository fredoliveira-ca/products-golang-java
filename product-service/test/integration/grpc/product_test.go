package integration

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"google.golang.org/grpc"
	"gotest.tools/assert"

	api "github.com/fredoliveira-ca/products-golang-java/product-service/app/grpc/proto/productpb"
	"github.com/fredoliveira-ca/products-golang-java/product-service/app/grpc/server"
	"github.com/fredoliveira-ca/products-golang-java/product-service/test/helper"
)

const (
	productAddress = "0.0.0.0:50051"
	createTable    = "CREATE TABLE product (product_id text PRIMARY KEY, price_in_cents integer, title text, description text)"
	insertTable    = "INSERT INTO product(product_id, price_in_cents, title, description) VALUES ($1, $2, $3, $4)"
)

var (
	productClient api.ProductPriceServiceClient
)

func TestMain(m *testing.M) {
	ctx := context.Background()

	c, err := helper.NewPostgreSQLContainer(ctx, helper.PostgreSQLContainerRequest{
		GenericContainerRequest: testcontainers.GenericContainerRequest{
			Started: true,
		},
	})
	if err != nil {
		log.Fatalf("did not get postgres container: %v", err)
	}
	defer c.Container.Terminate(ctx)

	conn, port, host, err := c.GetDriver(ctx)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	os.Setenv("DB_PORT", strings.Split(string(*port), "/")[0])
	os.Setenv("DB_HOST", host)

	time.Sleep(time.Second * 10)

	_, err = conn.ExecContext(ctx, createTable)
	if err != nil {
		log.Fatalf("did not execute: %v", err)
	}

	insertion, err := conn.Prepare(insertTable)
	insertion.Exec("12345", 5000, "Title test", "Description test")
	insertion2, err := conn.Prepare(insertTable)
	insertion2.Exec("1", 0, "0", "-1")
	insertion3, err := conn.Prepare(insertTable)
	insertion3.Exec("9325817d-f543-4718-9621-6d42d93d73f4", 1002003004, "product", "cool")

	go server.RegisterServer()

	connGrpc, errGrpc := grpc.Dial(productAddress, grpc.WithInsecure())
	if errGrpc != nil {
		log.Fatalf("did not connect: %v", errGrpc)
	}
	defer connGrpc.Close()
	productClient = api.NewProductPriceServiceClient(connGrpc)

	os.Exit(m.Run())
}

func TestFetchOne(t *testing.T) {
	tests := []struct {
		ID    string
		price int64
	}{
		{"12345", 5000},
		{"1", 0},
		{"9325817d-f543-4718-9621-6d42d93d73f4", 1002003004},
		{"nonexistent", 0},
	}

	for _, test := range tests {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		product, err := productClient.FetchOne(ctx, &api.ProductPriceRequest{ProductId: test.ID})
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		strPrice := strconv.FormatInt(product.ValueInCents, 10)
		strPriceTest := strconv.FormatInt(test.price, 10)
		assert.Equal(t, int64(product.ValueInCents), test.price, "Assertion failure! We've got the price: "+strPrice+" instead of "+strPriceTest+" for the product: "+test.ID)
	}

}
