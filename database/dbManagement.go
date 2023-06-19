package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ConnectToDB() *sql.DB {
	godotenv.Load()

	POSTGRES_HOST := os.Getenv("POSTGRES_HOST")
	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_DB := os.Getenv("POSTGRES_DB")

	sqlData := "host=" + POSTGRES_HOST + " user=" + POSTGRES_USER + " password=" + POSTGRES_PASSWORD + " dbname=" + POSTGRES_DB + " sslmode=disable"

	db, err := sql.Open("postgres", sqlData)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Successfully connected!")
	return db
}

func CreateTable(db *sql.DB, tableName string) {
	// check if table exists
	if _, err := db.Exec("SELECT 1 FROM " + tableName + " LIMIT 1"); err != nil {
		log.Printf("Table %s doesn't exist, creating...", tableName)
		// create table
		sqlStatement := `
		CREATE TABLE events (
			id SERIAL PRIMARY KEY,
			lender TEXT,
			borrower TEXT,
			tokenAddress TEXT,
			tokenId TEXT,
			transactionHash TEXT,
			blockNumber INTEGER,
			signature TEXT
		);`

		_, err = db.Exec(sqlStatement)
		if err != nil {
			panic(err)
		}

		log.Println("Successfully created table!")
		return
	}

	log.Printf("Table %s exists!", tableName)
}
