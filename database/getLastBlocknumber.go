package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func GetLastProcessedBlockNumber(db *sql.DB) int64 {
	var blockNumber int64
	err := db.QueryRow("SELECT blocknumber FROM events ORDER BY blocknumber DESC LIMIT 1").Scan(&blockNumber)

	if err != nil {
		if err == sql.ErrNoRows {
			// table is empty
			blockNumber = 0
		} else { // unexpected error
			panic(err)
		}
	}

	return blockNumber
}
