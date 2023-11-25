package main

import (
	"database/sql"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/counters"
	"log"
	"os"
	"strings"

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

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DbHost, DbPort, DbUser, DbPassword, DbName)

	log.Println("connecting to database")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = db.Close() }()

	log.Println("Database is connected...confirming")

	// Create the "hashes" table
	_, err = db.Exec(`
		select count(*) from hashes;
	`)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection confirmed")
	continueRunning := true
	queue := make(chan []byte, 1048576)

	go func() {
		log.Println("Generating hashes")
		c, _ := counters.NewByteCounter(keySpaceSize)
		for i := 0; i < PreComputeSize; i++ {
			hash := c.Sha1Bytes()
			queue <- hash[:]
			_ = c.FastIncrement()
		}
		continueRunning = false
	}()

	log.Println("storing hashes...")
	for continueRunning {
		valueList := "("
		for i := 0; i < 1048576; i++ {
			valueList += fmt.Sprintf("%v,", <-queue)
		}
		valueList = strings.TrimRight(valueList, ",")
		valueList += ");"
		_, err := db.Exec("INSERT INTO hashes (h) VALUES ($1)", valueList)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Data inserted successfully.")
}
