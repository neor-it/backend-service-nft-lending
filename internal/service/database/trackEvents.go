package database

import (
	"database/sql"
	"log"

	"GethBackServ/internal/service/structure"

	_ "github.com/lib/pq"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TrackEvents(client *ethclient.Client, contractAddress common.Address, eventSignature []byte, db *sql.DB) ([]structure.Event, error) {
	var events []structure.Event

	// get from db all events with eventSignature
	rows, err := db.Query("SELECT * FROM events WHERE signature = $1", eventSignature)
	if err != nil {
		log.Fatal(err)
	}

	var id int

	for rows.Next() {
		var event structure.Event
		err = rows.Scan(&id, &event.Lender, &event.Borrower, &event.TokenAddress, &event.TokenId, &event.TransactionHash, &event.BlockNumber, &event.Signature)
		if err != nil {
			log.Fatal(err)
		}
		events = append(events, event)
	}

	return events, nil
}
