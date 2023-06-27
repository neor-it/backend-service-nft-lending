package database

import (
	"database/sql"

	"github.com/ethereum/go-ethereum/core/types"
	_ "github.com/lib/pq"
)

func GetLastProcessedBlockNumber(db *sql.DB) (int64, error) {
	var blockNumber int64
	err := db.QueryRow("SELECT blocknumber FROM events ORDER BY blocknumber DESC LIMIT 1").Scan(&blockNumber)

	if err != nil {
		if err == sql.ErrNoRows {
			// table is empty
			blockNumber = 0
		} else { // unexpected error
			return 0, err
		}
	}

	return blockNumber, nil
}

func GetTransferEvent(log types.Log, db *sql.DB) (bool, error) {

	txHash := log.TxHash.Hex()
	rows, err := db.Query("SELECT * FROM transfers WHERE transactionhash = $1", txHash)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	if rows.Next() {
		return true, nil
	}

	return false, nil
}
