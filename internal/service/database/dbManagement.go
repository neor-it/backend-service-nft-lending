package database

import (
	"GethBackServ/internal/service/structure"
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var instance *structure.DBConnection

func GetConnection() (*structure.DBConnection, error) {
	if instance == nil {
		conn, err := connectToDB()
		if err != nil {
			return nil, err
		}
		instance = &structure.DBConnection{DB: conn}
	}

	return instance, nil
}

func connectToDB() (*sql.DB, error) {
	godotenv.Load()

	POSTGRES_HOST := os.Getenv("POSTGRES_HOST")
	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_DB := os.Getenv("POSTGRES_DB")

	sqlData := "host=" + POSTGRES_HOST + " user=" + POSTGRES_USER + " password=" + POSTGRES_PASSWORD + " dbname=" + POSTGRES_DB + " sslmode=disable"

	db, err := sql.Open("postgres", sqlData)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected!")
	return db, nil
}

func CreateTable(db *sql.DB, tableName string) error {
	// check if table exists
	sqlStatement := `
		CREATE TABLE IF NOT EXISTS events (
			id SERIAL PRIMARY KEY,
			lender TEXT,
			borrower TEXT,
			tokenAddress TEXT,
			tokenId TEXT,
			transactionHash TEXT,
			blockNumber INTEGER,
			signature TEXT
		);`

	_, err := db.Exec(sqlStatement)
	if err != nil {
		return err
	}

	sqlStatement = `
		CREATE TABLE IF NOT EXISTS transfers (
			id SERIAL PRIMARY KEY,
			fromAddress TEXT,
			toAddress TEXT,
			tokenAddress TEXT,
			tokenId TEXT,
			transactionHash TEXT,
			blockNumber INTEGER
		);`

	_, err = db.Exec(sqlStatement)
	if err != nil {
		return err
	}

	return nil
}
