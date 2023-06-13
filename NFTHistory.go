package main

import (
	"database/sql"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/lib/pq"
)

func normalizeAddress(address string) string {
	for len(address) != 42 {
		if strings.HasPrefix(address, "0x0") {
			address = "0x" + address[3:]
		} else {
			break
		}

	}

	return string(common.HexToAddress(strings.ToLower(address)).Hex())
}

func getNFTHistory(client *ethclient.Client, contractAddress common.Address, tokenAddress common.Address, tokenId *big.Int, walletAddress common.Address, db *sql.DB) []Event {
	var history []Event

	signatures := [][]byte{
		[]byte("NFTAdded"),
		[]byte("NFTWithdrawn"),
		[]byte("NFTBorrowed"),
		[]byte("NFTReturned"),
		[]byte("NFTCanceled"),
	}

	for _, signature := range signatures {
		nftEvents, err := trackEvents(client, contractAddress, signature, db)
		if err != nil {
			log.Fatal(err)
		}

		for _, event := range nftEvents {
			event.Borrower = normalizeAddress(event.Borrower)
			event.Lender = normalizeAddress(event.Lender)

			walletAddress = common.HexToAddress(strings.ToLower(walletAddress.String()))

			if ((event.TokenAddress == tokenAddress.Hex() && event.TokenId == tokenId.String()) && walletAddress.Hex() == "0x0000000000000000000000000000000000000000") ||
				((event.Borrower == walletAddress.Hex() || event.Lender == walletAddress.Hex()) && walletAddress.Hex() != "0x0000000000000000000000000000000000000000") {
				history = append(history, event)
			}
		}
	}

	return history
}
