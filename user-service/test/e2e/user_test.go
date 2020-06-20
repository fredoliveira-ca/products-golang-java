package e2e

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"google.golang.org/grpc"
	"gotest.tools/assert"

	api "github.com/fredoliveira-ca/products-golang-java/user-service/app/grpc/proto/userpb"
	"github.com/fredoliveira-ca/products-golang-java/user-service/app/grpc/server"
)

const (
	userAddress = "0.0.0.0:50058"
)

var (
	userClient api.UserServiceClient
)

func TestMain(m *testing.M) {
	conn, err := grpc.Dial(userAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	userClient = api.NewUserServiceClient(conn)

	go server.RegisterServer()

	os.Exit(m.Run())
}

func TestFetchOne(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := userClient.FetchOne(ctx, &api.UserRequest{UserId: "41597637-8c33-409f-a869-a2090e87ec78"})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	assert.Equal(t, r.User.FirstName, "John", "The two words should be the same.")

}
