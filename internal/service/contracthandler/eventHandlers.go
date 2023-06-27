package contracthandler

import (
	"GethBackServ/internal/endpoint/abigencontract"
	"GethBackServ/internal/service/database"
	"GethBackServ/internal/service/structure"
	"context"
	"database/sql"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// HandleEvents - subscribe on events, handle them and write to database
func HandleEvents(client *ethclient.Client, db *sql.DB, contractAddress common.Address, contractAbigen *abigencontract.MainFilterer) error {
	// Subscribe on event NFTAdded
	addedEvents, addedSubscription, err := subscriptionOnNFTAdded(client, contractAddress, contractAbigen)
	if err != nil {
		return err
	}
	defer addedSubscription.Unsubscribe()

	// Subscribe on event NFTCanceled
	canceledEvents, canceledSubscription, err := subscriptionOnNFTCanceled(client, contractAddress, contractAbigen)
	if err != nil {
		return err
	}
	defer canceledSubscription.Unsubscribe()

	// Subscribe on event NFTReturned
	returnedEvents, returnedSubscription, err := subscriptionOnNFTReturned(client, contractAddress, contractAbigen)
	if err != nil {
		return err
	}
	defer returnedSubscription.Unsubscribe()

	// Subscribe on event NFTWithdrawn
	withdrawnEvents, withdrawnSubscription, err := subscriptionOnNFTWithdrawn(client, contractAddress, contractAbigen)
	if err != nil {
		return err
	}
	defer withdrawnSubscription.Unsubscribe()

	// Subscribe on event NFTBorrowed
	borrowedEvents, borrowedSubscription, err := subscriptionOnNFTBorrowed(client, contractAddress, contractAbigen)
	if err != nil {
		return err
	}
	defer borrowedSubscription.Unsubscribe()

	for {
		select {
		case event := <-addedEvents:
			database.HandleNFTAdded(event, db)
			go HandleTransfers(client, db, event)

		case event := <-canceledEvents:
			database.HandleNFTCanceled(event, db)

		case event := <-returnedEvents:
			database.HandleNFTReturned(event, db)

		case event := <-withdrawnEvents:
			database.HandleNFTWithdrawn(event, db)

		case event := <-borrowedEvents:
			database.HandleNFTBorrowed(event, db)
		}
	}
}

func HandleTransfers(client *ethclient.Client, db *sql.DB, event *abigencontract.MainNFTAdded) error {
	contractAddress := event.Raw.Address
	tokenAddress := event.NFTAddress

	if structure.SubscriptionMap[tokenAddress] { // if subscription on this token already exists - return
		return nil
	}

	sub, logs, err := subscribeOnTransferEvent(client, event)
	if err != nil {
		return err
	}

	for {
		select {
		case err := <-sub.Err():
			return err
		case event := <-logs:
			log.Println("Transfer event", event)
			transferEventExists, err := database.GetTransferEvent(event, db)
			if err != nil {
				return err
			}

			if !transferEventExists {
				database.HandleTransfer(contractAddress, event, db)
			}
		}
	}
}

func HandleTransferEvents(client *ethclient.Client, tokenAddress common.Address, tokenId *big.Int) error {
	ethInfo, _ := GetEthClientInfo()
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

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
		return err
	}

	for _, log := range logs {
		tokenIdLog := log.Topics[3].Big()

		transferEventExists, err := database.GetTransferEvent(log, db.DB)
		if err != nil {
			return err
		}

		if transferEventExists {
			continue
		}

		if tokenIdLog.Cmp(tokenId) == 0 {
			database.HandleTransfer(ethInfo.ContractAddress, log, db.DB)
		}
	}

	return nil
}

func HandleMissedEvents(client *ethclient.Client, db *sql.DB, contractAddress common.Address, contractAbigen *abigencontract.MainFilterer) error {
	blockNumber, err := database.GetLastProcessedBlockNumber(db)
	if err != nil {
		return err
	}

	filterQuery := &bind.FilterOpts{
		Start:   uint64(blockNumber),
		End:     nil,
		Context: context.Background(),
	}

	err = processNFTAddedEvents(contractAbigen, filterQuery, db)
	if err != nil {
		return err
	}

	err = processNFTCanceledEvents(contractAbigen, filterQuery, db)
	if err != nil {
		return err
	}

	err = processNFTReturnedEvents(contractAbigen, filterQuery, db)
	if err != nil {
		return err
	}

	err = processNFTWithdrawnEvents(contractAbigen, filterQuery, db)
	if err != nil {
		return err
	}

	err = processNFTBorrowedEvents(contractAbigen, filterQuery, db)
	if err != nil {
		return err
	}

	return nil
}
