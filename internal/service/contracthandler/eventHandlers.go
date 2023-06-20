package contracthandler

import (
	"GethBackServ/internal/service/database"
	"context"
	"database/sql"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// HandleEvents handles all events from contract and saves them to database using HandleEvent function
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

			database.HandleEvent(vLog, eventData, db)
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

		database.HandleEvent(vLog, eventData, db)
	}
}
