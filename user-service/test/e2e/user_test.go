package e2e

import (
	"context"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"google.golang.org/grpc"
	"gotest.tools/assert"

	api "github.com/fredoliveira-ca/products-golang-java/user-service/app/grpc/proto/userpb"
	"github.com/fredoliveira-ca/products-golang-java/user-service/app/grpc/server"
	"github.com/fredoliveira-ca/products-golang-java/user-service/test/helper"
)

const (
	userAddress = "0.0.0.0:50053"
	createTable = "CREATE TABLE \"user\" (user_id text PRIMARY KEY, first_name text, last_name text, date_of_birth date NOT NULL);"
	insertTable = "INSERT INTO \"user\"(user_id, first_name, last_name, date_of_birth) VALUES ('41597637-8c33-409f-a869-a2090e87ec78', 'John', 'Generated', '1988-02-19');"
)

var (
	userClient api.UserServiceClient
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

	_, err = conn.ExecContext(ctx, createTable)
	if err != nil {
		log.Fatalf("did not execute: %v", err)
	}

	insertion, err := conn.Prepare(insertTable)
	if err != nil {
		log.Fatalf("did not insert: %v", err)
	}
	insertion.Exec()

	go server.RegisterServer()

	connGrpc, errGrpc := grpc.Dial(userAddress, grpc.WithInsecure())
	if errGrpc != nil {
		log.Fatalf("did not connect: %v", errGrpc)
	}
	defer connGrpc.Close()
	userClient = api.NewUserServiceClient(connGrpc)

	os.Exit(m.Run())
}

func TestFetchOne(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := userClient.FetchOne(ctx, &api.UserRequest{UserId: "41597637-8c33-409f-a869-a2090e87ec78"})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	assert.Equal(t, r.User.FirstName, "John", "The first name should be John! We got: "+r.User.FirstName)
}
