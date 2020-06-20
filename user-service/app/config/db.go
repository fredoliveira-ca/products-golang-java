package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

const (
	defaultHost = "localhost"
	defaultPort = 5432
	user        = "postgres"
	password    = "admin"
	dbname      = "productdb"
	driver      = "postgres"
)

// ConnectDataBase is a method to provide a connection with the database.
// If the environment variable is not available, it must assume the default value.
func ConnectDataBase() *sql.DB {
	host := defaultHost
	if os.Getenv("DB_HOST") != "" {
		host = os.Getenv("DB_HOST")
	}
	port := defaultPort
	if os.Getenv("DB_PORT") != "" {
		port, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	}

	conn := fmt.Sprintf(
		"host=%s port=%d user=%s  password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open(driver, conn)
	if err != nil {
		panic(err.Error())
	}

	return db
}
