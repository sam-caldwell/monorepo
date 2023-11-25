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
	PreComputeSize = 100
	//PreComputeSize = 256 * 256 * 256 * 256 //Table Size: 256 * 256 * 256 * 256 * 20 / 1048576 MB
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
	terminate := make(chan bool, 1)

	go func() {
		log.Println("Generating hashes")
		c, _ := counters.NewByteCounter(keySpaceSize)
		for i := 0; i < PreComputeSize; i++ {
			hash := c.Sha1Bytes()
			queue <- hash[:]
			_ = c.FastIncrement()
			log.Printf("generating %d (progress %3.4f %%)", i, 100*float64(i)/float64(PreComputeSize))
		}
		log.Println()
		continueRunning = false
	}()
	go func() {
		log.Println("storing hashes...")
		stored := 0
		for continueRunning {
			valueList := "("
			for i := 0; i < 1048576; i++ {
				valueList += fmt.Sprintf("%v,", <-queue)
			}
			valueList = strings.TrimRight(valueList, ",")
			valueList += ");"
			log.Printf("storing %d (progress %3.4f %%)", stored, 100*float64(stored)/float64(PreComputeSize))
			if _, err := db.Exec("INSERT INTO hashes (h) VALUES ($1)", valueList); err != nil {
				log.Fatal(err)
			}
		}
	}()

	<-terminate
	fmt.Println("Data inserted successfully.")
}
