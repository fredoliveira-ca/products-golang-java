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
	"gotest.tools/assert"

	"github.com/fredoliveira-ca/products-golang-java/product-service/data/repository"
	"github.com/fredoliveira-ca/products-golang-java/product-service/test/helper"
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

	time.Sleep(time.Second * 3)

	_, err = conn.ExecContext(ctx, "CREATE TABLE product (product_id text PRIMARY KEY, price_in_cents integer, title text, description text)")
	if err != nil {
		log.Fatalf("did not execute: %v", err)
	}

	insertion, err := conn.Prepare("INSERT INTO product(product_id, price_in_cents, title, description) VALUES ('12345', 5000, 'Title test', 'Description test');")
	insertion2, err := conn.Prepare("INSERT INTO product(product_id, price_in_cents, title, description) VALUES ('1', 0, '0', '-1');")
	insertion3, err := conn.Prepare("INSERT INTO product(product_id, price_in_cents, title, description) VALUES ('9325817d-f543-4718-9621-6d42d93d73f4', 1002003004, 'product', 'cool');")
	insertion.Exec()
	insertion2.Exec()
	insertion3.Exec()

	os.Exit(m.Run())
}
func TestFindOne(t *testing.T) {
	tests := []struct {
		ID          string
		title       string
		description string
		price       int64
	}{
		{"12345", "Title test", "Description test", 5000},
		{"1", "0", "-1", 0},
		{"9325817d-f543-4718-9621-6d42d93d73f4", "product", "cool", 1002003004},
		{"nonexistent", "", "", 0},
	}

	for _, test := range tests {
		product := repository.FindOne(test.ID)

		strPrice := strconv.FormatInt(product.PriceInCents, 10)
		assert.Equal(t, product.Title, test.title, "Assertion failure! We've got: "+product.Title+" for the product: "+test.ID)
		assert.Equal(t, product.Description, test.description, "Assertion failure! We've got: "+product.Description+" for the product: "+test.ID)
		assert.Equal(t, int64(product.PriceInCents), test.price, "Assertion failure! We've got: "+strPrice+" for the product: "+test.ID)
	}
}

func TestFindAll(t *testing.T) {
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

	products := repository.FindAll("")

	for _, product := range products {
		v, found := expected[product.ID]

		strPrice := strconv.FormatInt(product.PriceInCents, 10)
		assert.Equal(t, true, found, "Assertion failure! We expected found the product: "+product.ID)
		assert.Equal(t, v.title, product.Title, "Assertion failure! We've got: "+product.Title+" for the product: "+product.ID)
		assert.Equal(t, v.description, product.Description, "Assertion failure! We've got: "+product.Description+" for the product: "+product.ID)
		assert.Equal(t, int64(v.price), product.PriceInCents, "Assertion failure! We've got: "+strPrice+" for the product: "+product.ID)

	}
}
