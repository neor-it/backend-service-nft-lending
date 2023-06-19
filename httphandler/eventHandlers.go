package httphandler

import (
	"GethBackServ/database"
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	sqlStatement = `INSERT INTO events (lender, borrower, tokenAddress, tokenId, transactionHash, blockNumber, signature)
	VALUES ($1, $2, $3, $4, $5, $6, $7) `
)

func handleNFTAdded(vLog types.Log, eventData map[string]interface{}, db *sql.DB) {
	txHash := vLog.TxHash.Hex()
	blockNumber := vLog.BlockNumber

	tokenAddress := common.HexToAddress(vLog.Topics[2].Hex())
	tokenId := fmt.Sprintf("%v", eventData["tokenId"])
	lender := vLog.Topics[1].Hex()

	_, err := db.Exec(sqlStatement, lender, "", tokenAddress.Hex(), tokenId, txHash, blockNumber, "NFTAdded")
	if err != nil {
		panic(err)
	}
}

func handleNFTCanceled(vLog types.Log, eventData map[string]interface{}, db *sql.DB) {

	txHash := vLog.TxHash.Hex()
	blockNumber := vLog.BlockNumber

	lender := vLog.Topics[1].Hex()
	tokenId := fmt.Sprintf("%v", eventData["tokenId"])
	tokenAddress := common.HexToAddress(vLog.Topics[2].Hex())

	_, err := db.Exec(sqlStatement, lender, "", tokenAddress.Hex(), tokenId, txHash, blockNumber, "NFTCanceled")
	if err != nil {
		panic(err)
	}
}

func handleNFTBorrowed(vLog types.Log, eventData map[string]interface{}, db *sql.DB) {
	txHash := vLog.TxHash.Hex()
	blockNumber := vLog.BlockNumber

	tokenId := fmt.Sprintf("%v", eventData["tokenId"])
	tokenAddress := common.HexToAddress(vLog.Topics[3].Hex())
	borrower := vLog.Topics[1].Hex()
	lender := vLog.Topics[2].Hex()

	_, err := db.Exec(sqlStatement, lender, borrower, tokenAddress.Hex(), tokenId, txHash, blockNumber, "NFTBorrowed")
	if err != nil {
		panic(err)
	}
}

func handleNFTReturned(vLog types.Log, eventData map[string]interface{}, db *sql.DB) {
	txHash := vLog.TxHash.Hex()
	blockNumber := vLog.BlockNumber
	tokenId := fmt.Sprintf("%v", eventData["tokenId"])

	tokenAddress := common.HexToAddress(vLog.Topics[3].Hex())
	borrower := vLog.Topics[1].Hex()
	lender := vLog.Topics[2].Hex()

	_, err := db.Exec(sqlStatement, lender, borrower, tokenAddress.Hex(), tokenId, txHash, blockNumber, "NFTReturned")
	if err != nil {
		panic(err)
	}
}

func handleNFTWithdrawn(vLog types.Log, eventData map[string]interface{}, db *sql.DB) {
	txHash := vLog.TxHash.Hex()
	blockNumber := vLog.BlockNumber

	tokenId := fmt.Sprintf("%v", eventData["tokenId"])
	tokenAddress := common.HexToAddress(vLog.Topics[2].Hex())
	lender := vLog.Topics[1].Hex()

	_, err := db.Exec(sqlStatement, lender, "", tokenAddress.Hex(), tokenId, txHash, blockNumber, "NFTWithdrawn")
	if err != nil {
		panic(err)
	}
}

func handleEvent(vLog types.Log, eventData map[string]interface{}, db *sql.DB) {
	switch vLog.Topics[0].Hex() {
	case crypto.Keccak256Hash([]byte("NFTAdded(address,address,uint256)")).Hex():
		handleNFTAdded(vLog, eventData, db)
	case crypto.Keccak256Hash([]byte("NFTCanceled(address,address,uint256)")).Hex():
		handleNFTCanceled(vLog, eventData, db)
	case crypto.Keccak256Hash([]byte("NFTBorrowed(address,address,address,uint256)")).Hex():
		handleNFTBorrowed(vLog, eventData, db)
	case crypto.Keccak256Hash([]byte("NFTReturned(address,address,address,uint256)")).Hex():
		handleNFTReturned(vLog, eventData, db)
	case crypto.Keccak256Hash([]byte("NFTWithdrawn(address,address,uint256)")).Hex():
		handleNFTWithdrawn(vLog, eventData, db)
	}
}

func HandleEvents(client *ethclient.Client, db *sql.DB, contractAddress common.Address, contractAbi abi.ABI) {
	// event filter
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("Failed to subscribe to event logs: %v", err)
	}

	// Listening to event logs
	for {
		select {
		case err := <-sub.Err():
			log.Fatalf("Subscription error: %v", err)
		case vLog := <-logs:
			eventName, err := contractAbi.EventByID(vLog.Topics[0])
			if err != nil {
				log.Printf("Failed to retrieve event name: %v", err)
				continue
			}

			eventData := make(map[string]interface{})
			err = eventName.Inputs.UnpackIntoMap(eventData, vLog.Data)
			if err != nil {
				log.Printf("Failed to unmarshal event data: %v", err)
				continue
			}

			handleEvent(vLog, eventData, db)
		}
	}
}

func HandleMissedEvents(client *ethclient.Client, db *sql.DB, contractAddress common.Address, contractAbi abi.ABI) {
	blockNumber := database.GetLastProcessedBlockNumber(db)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			contractAddress,
		},
		FromBlock: big.NewInt(blockNumber),
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatalf("Failed to get missed events: %v", err)
	}

	for _, vLog := range logs {
		eventName, err := contractAbi.EventByID(vLog.Topics[0])
		if err != nil {
			log.Printf("Failed to retrieve event name: %v", err)
			continue
		}

		eventData := make(map[string]interface{})
		err = eventName.Inputs.UnpackIntoMap(eventData, vLog.Data)
		if err != nil {
			log.Printf("Failed to unmarshal event data: %v", err)
			continue
		}

		handleEvent(vLog, eventData, db)
	}
}
