package main

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func trackEvents(client *ethclient.Client, contractAddress common.Address, eventSignature []byte) ([]Event, error) {
	eventSignatureHash := crypto.Keccak256Hash(eventSignature)

	query := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   nil,
		Addresses: []common.Address{contractAddress},
		Topics: [][]common.Hash{
			{eventSignatureHash},
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, err
	}

	var events []Event

	for _, log := range logs {
		if eventSignatureHash.Hex() == log.Topics[0].Hex() {
			var event Event
			owner := common.BytesToAddress(log.Topics[1].Bytes())
			tokenId := new(big.Int).SetBytes(log.Data)
			transactionHash := log.TxHash.Hex()
			blockNumber := log.BlockNumber

			switch eventSignatureHash.Hex() {

			case crypto.Keccak256Hash([]byte("NFTAdded(address,address,uint256)")).Hex():
				contractAddr := common.BytesToAddress(log.Address.Bytes())
				tokenAddress := common.BytesToAddress(log.Topics[2].Bytes())

				event = Event{
					Owner:           owner.Hex(),
					TokenId:         tokenId.String(),
					TokenAddress:    tokenAddress.Hex(),
					Contract:        contractAddr.Hex(),
					TransactionHash: transactionHash,
					BlockNumber:     blockNumber,
					Borrower:        "",
					Lender:          "",
					Signature:       "NFTAdded",
				}

			case crypto.Keccak256Hash([]byte("NFTWithdrawn(address,address,uint256)")).Hex():
				tokenAddress := common.BytesToAddress(log.Topics[2].Bytes())

				event = Event{
					Owner:           owner.Hex(),
					TokenId:         tokenId.String(),
					TokenAddress:    tokenAddress.Hex(),
					Contract:        contractAddress.Hex(),
					TransactionHash: transactionHash,
					BlockNumber:     blockNumber,
					Borrower:        "",
					Lender:          "",
					Signature:       "NFTWithdrawn",
				}

			case crypto.Keccak256Hash([]byte("NFTBorrowed(address,address,address,uint256)")).Hex():
				borrower := common.BytesToAddress(log.Topics[1].Bytes())
				lender := common.BytesToAddress(log.Topics[2].Bytes())
				tokenAddress := common.BytesToAddress(log.Topics[3].Bytes())

				event = Event{
					Owner:           lender.Hex(),
					TokenId:         tokenId.String(),
					TokenAddress:    tokenAddress.Hex(),
					Contract:        contractAddress.Hex(),
					TransactionHash: transactionHash,
					BlockNumber:     blockNumber,
					Borrower:        borrower.Hex(),
					Lender:          lender.Hex(),
					Signature:       "NFTBorrowed",
				}

			case crypto.Keccak256Hash([]byte("NFTReturned(address,address,address,uint256)")).Hex():
				borrower := common.BytesToAddress(log.Topics[1].Bytes())
				lender := common.BytesToAddress(log.Topics[2].Bytes())
				tokenAddress := common.BytesToAddress(log.Topics[3].Bytes())

				event = Event{
					Owner:           lender.Hex(),
					TokenId:         tokenId.String(),
					TokenAddress:    tokenAddress.Hex(),
					Contract:        contractAddress.Hex(),
					TransactionHash: transactionHash,
					BlockNumber:     blockNumber,
					Borrower:        borrower.Hex(),
					Lender:          lender.Hex(),
					Signature:       "NFTReturned",
				}

			case crypto.Keccak256Hash([]byte("NFTCanceled(address,address,uint256)")).Hex():
				tokenAddress := common.BytesToAddress(log.Topics[2].Bytes())

				event = Event{
					Owner:           owner.Hex(),
					TokenId:         tokenId.String(),
					TokenAddress:    tokenAddress.Hex(),
					Contract:        contractAddress.Hex(),
					TransactionHash: transactionHash,
					BlockNumber:     blockNumber,
					Borrower:        "",
					Lender:          "",
					Signature:       "NFTCanceled",
				}
			}
			events = append(events, event)
		}
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
