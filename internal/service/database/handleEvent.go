package database

import (
	"GethBackServ/internal/endpoint/abigencontract"
	"GethBackServ/internal/service/structure"
	"database/sql"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	sqlStatement = `INSERT INTO events (lender, borrower, tokenAddress, tokenId, transactionHash, blockNumber, signature)
	VALUES ($1, $2, $3, $4, $5, $6, $7) `
)

func HandleNFTAdded(event *abigencontract.MainNFTAdded, db *sql.DB) error {
	txHash := event.Raw.TxHash.Hex()
	blockNumber := event.Raw.BlockNumber

	tokenAddress := common.HexToAddress(event.NFTAddress.Hex())
	tokenId := fmt.Sprintf("%v", event.TokenId)
	lender := event.Owner.Hex()

	_, err := db.Exec(sqlStatement, lender, "", tokenAddress.Hex(), tokenId, txHash, blockNumber, "NFTAdded")
	if err != nil {
		return err
	}

	return nil
}

func HandleNFTCanceled(event *abigencontract.MainNFTCanceled, db *sql.DB) error {
	txHash := event.Raw.TxHash.Hex()
	blockNumber := event.Raw.BlockNumber

	tokenAddress := common.HexToAddress(event.NFTAddress.Hex())
	tokenId := fmt.Sprintf("%v", event.TokenId)
	lender := event.Owner.Hex()

	_, err := db.Exec(sqlStatement, lender, "", tokenAddress.Hex(), tokenId, txHash, blockNumber, "NFTCanceled")
	if err != nil {
		return err
	}

	return nil
}

func HandleNFTBorrowed(event *abigencontract.MainNFTBorrowed, db *sql.DB) error {
	txHash := event.Raw.TxHash.Hex()
	blockNumber := event.Raw.BlockNumber

	tokenAddress := common.HexToAddress(event.NFTAddress.Hex())
	tokenId := fmt.Sprintf("%v", event.TokenId)
	lender := event.Lender.Hex()
	borrower := event.Borrower.Hex()

	_, err := db.Exec(sqlStatement, lender, borrower, tokenAddress.Hex(), tokenId, txHash, blockNumber, "NFTBorrowed")
	if err != nil {
		return err
	}

	return nil
}

func HandleNFTReturned(event *abigencontract.MainNFTReturned, db *sql.DB) error {
	txHash := event.Raw.TxHash.Hex()
	blockNumber := event.Raw.BlockNumber

	tokenAddress := common.HexToAddress(event.NFTAddress.Hex())
	tokenId := fmt.Sprintf("%v", event.TokenId)
	lender := event.Lender.Hex()
	borrower := event.Borrower.Hex()

	_, err := db.Exec(sqlStatement, lender, borrower, tokenAddress.Hex(), tokenId, txHash, blockNumber, "NFTReturned")
	if err != nil {
		return err
	}

	return nil
}

func HandleNFTWithdrawn(event *abigencontract.MainNFTWithdrawn, db *sql.DB) error {

	txHash := event.Raw.TxHash.Hex()
	blockNumber := event.Raw.BlockNumber

	tokenAddress := common.HexToAddress(event.NFTAddress.Hex())
	tokenId := fmt.Sprintf("%v", event.TokenId)
	lender := event.Owner.Hex()

	_, err := db.Exec(sqlStatement, lender, "", tokenAddress.Hex(), tokenId, txHash, blockNumber, "NFTWithdrawn")
	if err != nil {
		return err
	}

	return nil
}

func HandleTransfer(contractAddress common.Address, event types.Log, db *sql.DB) error {
	tokenAddress := event.Address.Hex()
	from := event.Topics[1].Hex()
	to := event.Topics[2].Hex()
	tokenId := event.Topics[3].Big().Uint64()
	blockNumber := event.BlockNumber
	txHash := event.TxHash.Hex()

	nContractAddress := structure.NormalizeAddress(contractAddress.Hex())
	nFrom := structure.NormalizeAddress(from)
	nTo := structure.NormalizeAddress(to)

	if nContractAddress == nFrom || nContractAddress == nTo {
		sqlStatement := `INSERT INTO transfers (fromAddress, toAddress, tokenAddress, tokenId, transactionHash, blockNumber)
		VALUES ($1, $2, $3, $4, $5, $6)`

		_, err := db.Exec(sqlStatement, from, to, tokenAddress, tokenId, txHash, blockNumber)
		if err != nil {
			return err
		}
	}

	return nil
}
