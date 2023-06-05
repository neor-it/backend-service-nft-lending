package main

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getNFTHistory(client *ethclient.Client, contractAddress common.Address, tokenAddress common.Address, tokenId *big.Int, walletAddress common.Address) []Event {
	var history []Event

	signatures := [][]byte{
		[]byte("NFTAdded(address,address,uint256)"),
		[]byte("NFTWithdrawn(address,address,uint256)"),
		[]byte("NFTBorrowed(address,address,address,uint256)"),
		[]byte("NFTReturned(address,address,address,uint256)"),
		[]byte("NFTCanceled(address,address,uint256)"),
	}

	for _, signature := range signatures {
		nftEvents, err := trackEvents(client, contractAddress, signature)
		if err != nil {
			log.Fatal(err)
		}

		for _, event := range nftEvents {
			if ((event.TokenAddress == tokenAddress.Hex() && event.TokenId == tokenId.String()) && walletAddress.Hex() == "0x0000000000000000000000000000000000000000") ||
				((event.Owner == walletAddress.Hex() || event.Borrower == walletAddress.Hex() || event.Lender == walletAddress.Hex()) && walletAddress.Hex() != "0x0000000000000000000000000000000000000000") {
				history = append(history, event)
			}
		}
	}

	return history
}
