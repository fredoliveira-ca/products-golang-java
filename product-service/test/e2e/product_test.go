package e2e

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"gotest.tools/assert"

	api "github.com/fredoliveira-ca/products-golang-java/product-service/app/grpc/proto/discountpb"
	"github.com/fredoliveira-ca/products-golang-java/product-service/app/web/route"
	domain "github.com/fredoliveira-ca/products-golang-java/product-service/domain"
	"github.com/fredoliveira-ca/products-golang-java/product-service/test/helper"
)

const (
	discountAddress = "0.0.0.0:50052"
	createTable     = "CREATE TABLE product (product_id text PRIMARY KEY, price_in_cents integer, title text, description text)"
	insertTable     = "INSERT INTO product(product_id, price_in_cents, title, description) VALUES ($1, $2, $3, $4)"
)

var (
	productClient api.DiscountServiceClient
)

// Fazer para produtos
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

	route.LoadRoutes()

	go http.ListenAndServe(":8001", nil)

	// Run tests
	os.Exit(m.Run())
}

func TestListAllProductWithoutDiscount(t *testing.T) {
	req, err := http.Get("http://localhost:8001/product")
	if err != nil {
		log.Fatalf("did not execute: %v", err)
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("did not execute: %v", err)
	}

	var products []domain.Product
	json.Unmarshal([]byte(body), &products)

	expected := map[string]struct {
		ID          string
		title       string
		description string
		price       int64
	}{
		"12345":                                {"12345", "Title test", "Description test", 5000},
		"1":                                    {"1", "0", "-1", 0},
		"9325817d-f543-4718-9621-6d42d93d73f4": {"9325817d-f543-4718-9621-6d42d93d73f4", "product", "cool", 1002003004},
	}

	for _, product := range products {
		v, found := expected[product.ID]

		strPrice := strconv.FormatInt(product.PriceInCents, 10)
		strPriceTest := strconv.FormatInt(v.price, 10)
		assert.Equal(t, true, found, "Assertion failure! We expected found the product: "+product.ID)
		assert.Equal(t, v.title, product.Title, "Assertion failure! We've got: "+product.Title+" instead of "+strPriceTest+" for the product: "+product.ID)
		assert.Equal(t, v.description, product.Description, "Assertion failure! We've got: "+product.Description+" instead of "+strPriceTest+" for the product: "+product.ID)
		assert.Equal(t, int64(v.price), product.PriceInCents, "Assertion failure! We've got: "+strPrice+" instead of "+strPriceTest+" for the product: "+product.ID)

	}

}

func TestListAllProductWithBirthdayDiscount(t *testing.T) {

}
