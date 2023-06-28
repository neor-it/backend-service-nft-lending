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

// GetAllTokenAddressesAndIds - returns all unique token addresses and token ids from the events table
func GetAllTokenAddressesAndIds(db *sql.DB) ([]string, []int64, error) {
	rows, err := db.Query("SELECT DISTINCT tokenAddress, tokenId FROM events")
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var tokenAddresses []string
	var tokenIds []int64

	for rows.Next() {
		var tokenAddress string
		var tokenId int64
		err = rows.Scan(&tokenAddress, &tokenId)
		if err != nil {
			return nil, nil, err
		}
		tokenAddresses = append(tokenAddresses, tokenAddress)
		tokenIds = append(tokenIds, tokenId)
	}

	return tokenAddresses, tokenIds, nil
}
