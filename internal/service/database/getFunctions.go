package database

import (
	"GethBackServ/internal/endpoint/abigencontract"
	"database/sql"

	_ "github.com/lib/pq"
)

func GetLastProcessedBlockNumberInEvents(db *sql.DB) (int64, error) {
	var blockNumber int64
	err := db.QueryRow("SELECT blocknumber FROM events ORDER BY blocknumber DESC LIMIT 1").Scan(&blockNumber)

	if err != nil {
		if err != sql.ErrNoRows {
			return 0, err // error other than no rows
		}
		blockNumber = 0 // table is empty
	}

	return blockNumber, nil
}

func GetLastProcessedBlockNumberInTransfers(db *sql.DB) (int64, error) {
	var blockNumber int64
	err := db.QueryRow("SELECT blocknumber FROM transfers ORDER BY blocknumber DESC LIMIT 1").Scan(&blockNumber)

	if err != nil {
		if err != sql.ErrNoRows {
			return 0, err // error other than no rows
		}
		blockNumber = 0 // table is empty
	}

	return blockNumber, nil
}

func GetTransferEvent(log *abigencontract.Erc721Transfer, db *sql.DB) (bool, error) {
	txHash := log.Raw.TxHash.Hex()

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

// GetTransactionByHash - returns true if the transaction hash is already in the database
func GetTransactionByHash(txHash string, db *sql.DB) (bool, error) {
	rows, err := db.Query("SELECT * FROM events WHERE transactionhash = $1", txHash)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	if rows.Next() {
		return true, nil
	}

	return false, nil
}
