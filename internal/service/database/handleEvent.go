package database

import (
	"GethBackServ/internal/endpoint/abigencontract"
	"database/sql"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

const (
	sqlStatement = `INSERT INTO events (lender, borrower, tokenAddress, tokenId, transactionHash, blockNumber, signature)
	VALUES ($1, $2, $3, $4, $5, $6, $7) `
)

func HandleNFTAdded(event *abigencontract.MainNFTAdded, db *sql.DB) error {
	txHash := event.Raw.TxHash.Hex()

	isExists, err := GetTransactionByHash(txHash, db)
	if err != nil {
		return err
	}

	if isExists { // if transaction already exists in database
		return nil
	}

	blockNumber := event.Raw.BlockNumber

	tokenAddress := common.HexToAddress(event.NFTAddress.Hex())
	tokenId := fmt.Sprintf("%v", event.TokenId)
	lender := event.Owner.Hex()

	_, err = db.Exec(sqlStatement, lender, "", tokenAddress.Hex(), tokenId, txHash, blockNumber, "NFTAdded")
	if err != nil {
		return err
	}

	return nil
}

func HandleNFTCanceled(event *abigencontract.MainNFTCanceled, db *sql.DB) error {
	txHash := event.Raw.TxHash.Hex()

	isExists, err := GetTransactionByHash(txHash, db)
	if err != nil {
		return err
	}

	if isExists { // if transaction already exists in database
		return nil
	}

	blockNumber := event.Raw.BlockNumber

	tokenAddress := common.HexToAddress(event.NFTAddress.Hex())
	tokenId := fmt.Sprintf("%v", event.TokenId)
	lender := event.Owner.Hex()

	_, err = db.Exec(sqlStatement, lender, "", tokenAddress.Hex(), tokenId, txHash, blockNumber, "NFTCanceled")
	if err != nil {
		return err
	}

	return nil
}

func HandleNFTBorrowed(event *abigencontract.MainNFTBorrowed, db *sql.DB) error {
	txHash := event.Raw.TxHash.Hex()

	isExists, err := GetTransactionByHash(txHash, db)
	if err != nil {
		return err
	}

	if isExists { // if transaction already exists in database
		return nil
	}

	blockNumber := event.Raw.BlockNumber

	tokenAddress := common.HexToAddress(event.NFTAddress.Hex())
	tokenId := fmt.Sprintf("%v", event.TokenId)
	lender := event.Lender.Hex()
	borrower := event.Borrower.Hex()

	_, err = db.Exec(sqlStatement, lender, borrower, tokenAddress.Hex(), tokenId, txHash, blockNumber, "NFTBorrowed")
	if err != nil {
		return err
	}

	return nil
}

func HandleNFTReturned(event *abigencontract.MainNFTReturned, db *sql.DB) error {
	txHash := event.Raw.TxHash.Hex()

	isExists, err := GetTransactionByHash(txHash, db)
	if err != nil {
		return err
	}

	if isExists { // if transaction already exists in database
		return nil
	}

	blockNumber := event.Raw.BlockNumber

	tokenAddress := common.HexToAddress(event.NFTAddress.Hex())
	tokenId := fmt.Sprintf("%v", event.TokenId)
	lender := event.Lender.Hex()
	borrower := event.Borrower.Hex()

	_, err = db.Exec(sqlStatement, lender, borrower, tokenAddress.Hex(), tokenId, txHash, blockNumber, "NFTReturned")
	if err != nil {
		return err
	}

	return nil
}

func HandleNFTWithdrawn(event *abigencontract.MainNFTWithdrawn, db *sql.DB) error {
	txHash := event.Raw.TxHash.Hex()

	isExists, err := GetTransactionByHash(txHash, db)
	if err != nil {
		return err
	}

	if isExists { // if transaction already exists in database
		return nil
	}

	blockNumber := event.Raw.BlockNumber

	tokenAddress := common.HexToAddress(event.NFTAddress.Hex())
	tokenId := fmt.Sprintf("%v", event.TokenId)
	lender := event.Owner.Hex()

	_, err = db.Exec(sqlStatement, lender, "", tokenAddress.Hex(), tokenId, txHash, blockNumber, "NFTWithdrawn")
	if err != nil {
		return err
	}

	return nil
}

func HandleTransfer(contractAddress common.Address, event *abigencontract.Erc721Transfer, db *sql.DB) error {
	tokenAddress := event.Raw.Address.Hex()
	from := event.From.Hex()
	to := event.To.Hex()
	tokenId := event.Raw.Topics[3].Big().Uint64()
	blockNumber := event.Raw.BlockNumber
	txHash := event.Raw.TxHash.Hex()

	if contractAddress.Hex() == from || contractAddress.Hex() == to {
		sqlStatement := `INSERT INTO transfers (fromAddress, toAddress, tokenAddress, tokenId, transactionHash, blockNumber)
		VALUES ($1, $2, $3, $4, $5, $6)`

		_, err := db.Exec(sqlStatement, from, to, tokenAddress, tokenId, txHash, blockNumber)
		if err != nil {
			return err
		}
	}

	return nil
}
