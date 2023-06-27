package database

import (
	"GethBackServ/internal/service/structure"
	"database/sql"

	"github.com/ethereum/go-ethereum/common"
	_ "github.com/lib/pq"
)

func TrackEvents(contractAddress common.Address, eventSignature []byte, db *sql.DB) ([]structure.Event, error) {
	var events []structure.Event

	rows, err := db.Query("SELECT * FROM events WHERE signature = $1", eventSignature)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var id int

	for rows.Next() {
		var event structure.Event
		err = rows.Scan(&id, &event.Lender, &event.Borrower, &event.TokenAddress, &event.TokenId, &event.TransactionHash, &event.BlockNumber, &event.Signature)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func TrackTransfers(filterAddress common.Address, tokenId int64, db *sql.DB) ([]structure.Transfers, error) {
	rows, err := db.Query("SELECT * FROM transfers WHERE tokenAddress = $1 AND tokenId = $2", filterAddress.Hex(), tokenId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transfers []structure.Transfers

	var id int

	for rows.Next() {
		var transfer structure.Transfers
		err = rows.Scan(&id, &transfer.FromAddress, &transfer.ToAddress, &transfer.TokenAddress, &transfer.TokenId, &transfer.TransactionHash, &transfer.BlockNumber)
		if err != nil {
			return nil, err
		}
		transfers = append(transfers, transfer)
	}

	return transfers, nil
}
