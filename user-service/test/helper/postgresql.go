package helper

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/docker/go-connections/nat"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	testcontainers "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	postgresUser       = "postgres"
	postgresPassword   = "admin"
	postgresDatabase   = "productdb"
	postgresImage      = "postgres"
	postgresDefaultTag = "alpine"
	postgresPort       = "5432/tcp"
)

// PostgreSQLContainerRequest completes GenericContainerRequest
//  with PostgreSQL specific parameters.
type PostgreSQLContainerRequest struct {
	testcontainers.GenericContainerRequest
	User     string
	Password string
	Database string
}

// PostgreSQLContainer should always be created via NewPostgreSQLContainer.
type PostgreSQLContainer struct {
	Container testcontainers.Container
	db        *sql.DB
	req       PostgreSQLContainerRequest
}

// GetDriver returns a sql.DB connecting to the previously started Postgres DB.
// All the parameters are taken from the previous PostgreSQLContainerRequest.
func (c *PostgreSQLContainer) GetDriver(ctx context.Context) (*sql.DB, *nat.Port, string, error) {

	host, err := c.Container.Host(ctx)
	if err != nil {
		return nil, nil, "", err
	}

	mappedPort, err := c.Container.MappedPort(ctx, postgresPort)
	if err != nil {
		return nil, nil, "", err
	}

	db, err := sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		mappedPort.Int(),
		c.req.User,
		c.req.Password,
		c.req.Database,
	))
	if err != nil {
		return nil, nil, "", err
	}

	log.Println(fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		mappedPort.Int(),
		c.req.User,
		c.req.Password,
		c.req.Database,
	))
	log.Println(db)

	return db, &mappedPort, host, nil
}

// NewPostgreSQLContainer creates and starts a Postgres database.
// If autostarted, the function returns only after a successful execution of a query
// (confirming that the database is ready)
func NewPostgreSQLContainer(ctx context.Context, req PostgreSQLContainerRequest) (*PostgreSQLContainer, error) {
	provider, err := req.ProviderType.GetProvider()
	if err != nil {
		return nil, err
	}

	req.ExposedPorts = []string{postgresPort}

	if req.Env == nil {
		req.Env = map[string]string{}
	}

	if req.Image == "" && req.FromDockerfile.Context == "" {
		req.Image = fmt.Sprintf("%s:%s", postgresImage, postgresDefaultTag)
	}

	if req.User == "" {
		req.User = postgresUser
	}

	if req.Password == "" {
		req.Password = postgresPassword
	}

	if req.Database == "" {
		req.Database = postgresDatabase
	}

	req.Env["POSTGRES_USER"] = req.User
	req.Env["POSTGRES_PASSWORD"] = req.Password
	req.Env["POSTGRES_DB"] = req.Database

	req.WaitingFor = wait.ForLog("database system is ready to accept connections").WithStartupTimeout(time.Minute * 1)

	c, err := provider.CreateContainer(ctx, req.ContainerRequest)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create container")
	}

	postgresC := &PostgreSQLContainer{
		Container: c,
		req:       req,
	}

	if req.Started {
		if err := c.Start(ctx); err != nil {
			return postgresC, errors.Wrap(err, "failed to start container")
		}
	}

	return postgresC, nil
}
