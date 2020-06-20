package integration

import (
	"context"
	"testing"
	"time"

	_ "github.com/lib/pq"
	testcontainers "github.com/testcontainers/testcontainers-go"
	"gotest.tools/assert"

	"github.com/fredoliveira-ca/products-golang-java/user-service/test/helper"
)

func TestWriteIntoDB(t *testing.T) {
	ctx := context.Background()

	c, err := helper.NewPostgreSQLContainer(ctx, helper.PostgreSQLContainerRequest{
		GenericContainerRequest: testcontainers.GenericContainerRequest{
			Started: true,
		},
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	defer c.Container.Terminate(ctx)

	conn, _, _, err := c.GetDriver(ctx)
	if err != nil {
		t.Fatal(err.Error())
	}

	time.Sleep(time.Second * 3)

	_, err = conn.ExecContext(ctx, "CREATE TABLE testing ( id integer, data varchar(32) )")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestInteractWithDB(t *testing.T) {
	ctx := context.Background()

	c, err := helper.NewPostgreSQLContainer(ctx, helper.PostgreSQLContainerRequest{
		GenericContainerRequest: testcontainers.GenericContainerRequest{
			Started: true,
		},
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	defer c.Container.Terminate(ctx)

	conn, _, _, err := c.GetDriver(ctx)
	if err != nil {
		t.Fatal(err.Error())
	}

	time.Sleep(time.Second * 3)

	_, err = conn.ExecContext(ctx, "CREATE TABLE \"user\" ( user_id integer, name varchar(32) )")
	if err != nil {
		t.Fatal(err.Error())
	}

	insertion, err := conn.Prepare("INSERT INTO \"user\"(user_id, name) VALUES($1, $2)")
	if err != nil {
		t.Fatal(err.Error())
	}
	insertion.Exec(1, "John")

	selection, err := conn.Query("SELECT * FROM \"user\" WHERE user_id=$1;", 1)
	if err != nil {
		t.Fatal(err.Error())
	}

	for selection.Next() {
		var (
			id   int
			name string
		)

		err = selection.Scan(&id, &name)
		if err != nil {
			t.Fatal(err.Error())
		}

		assert.Equal(t, 1, id, "The property 'ID' should be equals: 1.")
		assert.Equal(t, "John", name, "The property 'name' should be: John.")
	}

}
