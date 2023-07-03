package contracthandler

import (
	"GethBackServ/internal/endpoint/abigencontract"
	"GethBackServ/internal/service/structure"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
)

// subscriptionOnNFTAdded - subscribe on NFTAdded event
func subscriptionOnNFTAdded(client *ethclient.Client, contractAddress common.Address, contractAbigen *abigencontract.MainFilterer) (chan *abigencontract.MainNFTAdded, event.Subscription, error) {
	addedEvents := make(chan *abigencontract.MainNFTAdded)
	addedSubscription, err := contractAbigen.WatchNFTAdded(nil, addedEvents, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	return addedEvents, addedSubscription, nil
}

// subscriptionOnNFTCanceled - subscribe on NFTCanceled event
func subscriptionOnNFTCanceled(client *ethclient.Client, contractAddress common.Address, contractAbigen *abigencontract.MainFilterer) (chan *abigencontract.MainNFTCanceled, event.Subscription, error) {
	canceledEvents := make(chan *abigencontract.MainNFTCanceled)
	canceledSubscription, err := contractAbigen.WatchNFTCanceled(nil, canceledEvents, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	return canceledEvents, canceledSubscription, nil
}

// subscriptionOnNFTReturned - subscribe on NFTReturned event
func subscriptionOnNFTReturned(client *ethclient.Client, contractAddress common.Address, contractAbigen *abigencontract.MainFilterer) (chan *abigencontract.MainNFTReturned, event.Subscription, error) {
	returnedEvents := make(chan *abigencontract.MainNFTReturned)
	returnedSubscription, err := contractAbigen.WatchNFTReturned(nil, returnedEvents, nil, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	return returnedEvents, returnedSubscription, nil
}

// subscriptionOnNFTWithdrawn - subscribe on NFTWithdrawn event
func subscriptionOnNFTWithdrawn(client *ethclient.Client, contractAddress common.Address, contractAbigen *abigencontract.MainFilterer) (chan *abigencontract.MainNFTWithdrawn, event.Subscription, error) {
	withdrawnEvents := make(chan *abigencontract.MainNFTWithdrawn)
	withdrawnSubscription, err := contractAbigen.WatchNFTWithdrawn(nil, withdrawnEvents, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	return withdrawnEvents, withdrawnSubscription, nil
}

// subscriptionOnNFTBorrowed - subscribe on NFTBorrowed event
func subscriptionOnNFTBorrowed(client *ethclient.Client, contractAddress common.Address, contract *abigencontract.MainFilterer) (chan *abigencontract.MainNFTBorrowed, event.Subscription, error) {
	borrowedEvents := make(chan *abigencontract.MainNFTBorrowed)
	borrowedSubscription, err := contract.WatchNFTBorrowed(nil, borrowedEvents, nil, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	return borrowedEvents, borrowedSubscription, nil
}

// subscribeOnTransferEvent - subscribe on Transfer event for token with address tokenAddress
func subscribeOnTransferEvent(client *ethclient.Client, contractAddress common.Address, contract *abigencontract.Erc721Filterer) (ethereum.Subscription, chan *abigencontract.Erc721Transfer, error) {
	transferEvents := make(chan *abigencontract.Erc721Transfer)
	transferSubscription, err := contract.WatchTransfer(nil, transferEvents, nil, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	structure.SubscriptionMap[contractAddress] = true

	return transferSubscription, transferEvents, nil
}
