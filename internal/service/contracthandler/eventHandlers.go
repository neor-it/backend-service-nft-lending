package contracthandler

import (
	"GethBackServ/internal/endpoint/abigencontract"
	"GethBackServ/internal/service/database"
	"GethBackServ/internal/service/structure"
	"context"
	"database/sql"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

	contractAbigen, err := abigencontract.NewErc721Filterer(tokenAddress, client)
	if err != nil {
		return err
	}

	sub, logs, err := subscribeOnTransferEvent(client, tokenAddress, contractAbigen)
	if err != nil {
		return err
	}

	for {
		select {
		case err := <-sub.Err():
			log.Printf("Error in subscription on transfer event: %v", err)
		case event := <-logs:
			transferEventExists, err := database.GetTransferEvent(event, db)
			if err != nil {
				panic(err)
			}

			if !transferEventExists {
				database.HandleTransfer(contractAddress, event, db)
			}
		}
	}
}

func HandleTransferEvents(client *ethclient.Client, tokenAddress common.Address, tokenId *big.Int) error {
	ethInfo, err := GetEthClientInfo()
	if err != nil {
		return err
	}

	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	contractAbigen, err := abigencontract.NewErc721Filterer(tokenAddress, client)
	if err != nil {
		return err
	}

	blockNumber, err := database.GetLastProcessedBlockNumberInTransfers(db.DB)
	if err != nil {
		return err
	}

	filterQuery := &bind.FilterOpts{
		Start:   uint64(blockNumber),
		End:     nil,
		Context: context.Background(),
	}

	err = processTransferEvents(ethInfo.ContractAddress, contractAbigen, filterQuery, db.DB)
	if err != nil {
		return err
	}

	return nil
}

func HandleMissedEvents(client *ethclient.Client, db *sql.DB, contractAddress common.Address, contractAbigen *abigencontract.MainFilterer) error {
	blockNumber, err := database.GetLastProcessedBlockNumberInEvents(db)
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

func HandleMissedTransfers(ethInfo *structure.EthClientInfo, db *sql.DB) error {
	tokenAddresses, tokenIds, err := database.GetAllTokenAddressesAndIds(db)
	if err != nil {
		panic(err)
	}

	contractAddress := ethInfo.ContractAddress

	for i, tokenAddress := range tokenAddresses {
		HandleTransferEvents(ethInfo.Client, common.HexToAddress(tokenAddress), big.NewInt(tokenIds[i]))

		event := &abigencontract.MainNFTAdded{
			NFTAddress: common.HexToAddress(tokenAddress),
			TokenId:    big.NewInt(tokenIds[i]),
			Raw:        types.Log{Address: contractAddress},
		}
		go HandleTransfers(ethInfo.Client, db, event)

	}

	return nil
}
