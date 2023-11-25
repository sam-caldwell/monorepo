package main

import (
	"database/sql"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/counters"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	keySpaceSize   = 1024
	PreComputeSize = 256 * 256 * 256 * 256 //Table Size: 256 * 256 * 256 * 256 * 20 / 1048576 MB
)

func main() {
	// Connect to the PostgreSQL database
	DbHost := os.Getenv("db_host")
	DbPort := os.Getenv("db_port")
	DbUser := os.Getenv("db_user")
	DbPassword := os.Getenv("db_pass")
	DbName := os.Getenv("db_name")

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DbHost, DbPort, DbUser, DbPassword, DbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = db.Close() }()

	// Create the "hashes" table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS hashes (
			h BYTEA PRIMARY KEY
		);
	`)
	if err != nil {
		log.Fatal(err)
	}

	c, _ := counters.NewByteCounter(keySpaceSize)
	for i := 0; i < PreComputeSize; i++ {
		hash := c.Sha1Bytes()
		_, err := db.Exec("INSERT INTO hashes (h) VALUES ($1)", hash[:])
		if err != nil {
			log.Fatal(err)
		}
		_ = c.FastIncrement()
	}
	fmt.Println("Data inserted successfully.")
}
