package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sam-caldwell/monorepo/go/counters"
	"log"
	"os"
	"strings"
	"time"
)

const (
	keySpaceSize = 1024
	//PreComputeSize = 100
	PreComputeSize = 256 * 256 * 256 * 256 //Table Size: 256 * 256 * 256 * 256 * 20 / 1048576 MB
)

func main() {
	DbHost := os.Getenv("db_host")
	DbPort := os.Getenv("db_port")
	DbUser := os.Getenv("db_user")
	DbPassword := os.Getenv("db_pass")
	DbName := os.Getenv("db_name")

	continueRunning := true
	queue := make(chan []byte, 1048576)
	terminate := make(chan bool, 1)
	genCount := 0
	storeCount := 0

	go func() {
		t := time.NewTicker(1 * time.Second)
		defer t.Stop()
		for continueRunning {
			select {
			case <-t.C:
				gProgress := 100 * float64(genCount) / float64(PreComputeSize)
				sProgress := 100 * float64(storeCount) / float64(len(queue))
				log.Printf("t: %d generator progres: %d/%d (3.4%f %%) storage: %d/%d (3.4%f %%)",
					genCount, PreComputeSize, gProgress, storeCount, len(queue), sProgress)
			}
		}
	}()

	// Connect to the PostgreSQL database
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DbHost, DbPort, DbUser, DbPassword, DbName)

	log.Println("connecting to database")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = db.Close() }()

	log.Println("Database is connected...confirming")
	if _, err = db.Exec(`select count(*) from hashes;`); err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection confirmed")

	go func() {
		log.Println("Generating hashes")
		c, _ := counters.NewByteCounter(keySpaceSize)
		for i := 0; i < PreComputeSize; i++ {
			hash := c.Sha1Bytes()
			queue <- hash[:]
			_ = c.FastIncrement()
			genCount = i
		}
		log.Println()
		continueRunning = false
	}()
	go func() {
		log.Println("storing hashes...")
		for continueRunning {
			valueList := "("
			for i := 0; i < 1048576; i++ {
				valueList += fmt.Sprintf("%v,", <-queue)
			}
			valueList = strings.TrimRight(valueList, ",")
			valueList += ");"
			storeCount++
			if _, err := db.Exec("INSERT INTO hashes (h) VALUES ($1)", valueList); err != nil {
				log.Fatal(err)
			}
		}
		log.Println("storage routine terminating")
		terminate <- true
	}()

	<-terminate
	log.Println("Data inserted successfully.")
}
