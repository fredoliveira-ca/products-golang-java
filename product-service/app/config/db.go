package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	defaultHost = "localhost"
	user        = "postgres"
	password    = "admin"
	dbname      = "productdb"
	driver      = "postgres"
)

// ConnectDataBase is a way to open a connection with the database.
// If the environment variable is not available, it must assume the default value.
func ConnectDataBase() *sql.DB {
	host := defaultHost
	if os.Getenv("DB_HOST") != "" {
		host = os.Getenv("DB_HOST")
	}

	connection := fmt.Sprintf(
		"host=%s port=%d user=%s  password=%s dbname=%s sslmode=disable",
		host, 5432, user, password, dbname,
	)

	log.Println("Connecting database", connection)
	db, err := sql.Open(driver, connection)
	if err != nil {
		panic(err.Error())
	}

	return db
}
