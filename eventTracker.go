package main

import (
	"context"
	"database/sql"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/lib/pq"
)

func trackEvents(client *ethclient.Client, contractAddress common.Address, eventSignature []byte, db *sql.DB) ([]Event, error) {
	var events []Event

	// get from db all events with eventSignature
	rows, err := db.Query("SELECT * FROM events WHERE signature = $1", eventSignature)
	if err != nil {
		log.Fatal(err)
	}

	var id int

	for rows.Next() {
		var event Event
		err = rows.Scan(&id, &event.Lender, &event.Borrower, &event.TokenAddress, &event.TokenId, &event.TransactionHash, &event.BlockNumber, &event.Signature)
		if err != nil {
			log.Fatal(err)
		}
		events = append(events, event)
	}

	return events, nil
}

func getTransfersByAddress(client *ethclient.Client, filterAddress common.Address, tokenAddress common.Address, tokenId *big.Int) []Transfers {
	transferEventSignature := crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))

	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			tokenAddress,
		},
		Topics: [][]common.Hash{
			{
				transferEventSignature,
			},
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)

	if err != nil {
		log.Fatalf("Failed to filter logs: %v", err)
	}

	var transfers []Transfers

	for _, log := range logs {
		fromAddress := common.BytesToAddress(log.Topics[1].Bytes())
		toAddress := common.BytesToAddress(log.Topics[2].Bytes())
		tokenIdLog := new(big.Int).SetBytes(log.Topics[3].Bytes())

		if (fromAddress == filterAddress || toAddress == filterAddress) && tokenIdLog.Cmp(tokenId) == 0 {
			transfers = append(transfers, Transfers{
				FromAddress: fromAddress.Hex(),
				ToAddress:   toAddress.Hex(),
				TokenID:     int(tokenIdLog.Int64()),
			})
		}
	}

	return transfers
}
