package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"

	"database/sql"

	_ "github.com/lib/pq"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	sqlStatement = `INSERT INTO events (lender, borrower, tokenAddress, tokenId, transactionHash, blockNumber, signature)
	VALUES ($1, $2, $3, $4, $5, $6, $7) `
)

type Event struct {
	Lender          string `json:"lender"`
	Borrower        string `json:"borrower"`
	TokenAddress    string `json:"nftContract"`
	TokenId         string `json:"tokenId"`
	TransactionHash string `json:"transactionHash"`
	BlockNumber     uint64 `json:"blockNumber"`
	Signature       string `json:"signature"`
}

func readAbi(fileName string) abi.ABI {
	fileAbi, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(bytes.NewReader(fileAbi))
	if err != nil {
		log.Fatal(err)
	}

	return contractAbi
}

func handleNFTAdded(vLog types.Log, eventData map[string]interface{}, db *sql.DB) {
	fmt.Println("NFTAdded")
	txHash := vLog.TxHash.Hex()
	blockNumber := vLog.BlockNumber
	signature := vLog.Topics[0].Hex()

	tokenAddress := common.HexToAddress(vLog.Topics[2].Hex())
	tokenId := fmt.Sprintf("%v", eventData["tokenId"])
	lender := vLog.Topics[1].Hex()

	event := Event{
		Lender:          lender,
		Borrower:        "",
		TransactionHash: txHash,
		BlockNumber:     blockNumber,
		TokenAddress:    tokenAddress.Hex(),
		TokenId:         tokenId,
		Signature:       signature,
	}
	fmt.Println("EVENT:", event)

	_, err := db.Exec(sqlStatement, lender, "", tokenAddress.Hex(), tokenId, txHash, blockNumber, "NFTAdded")
	if err != nil {
		panic(err)
	}
}

func handleNFTReceived(vLog types.Log, eventData map[string]interface{}, db *sql.DB) {
	fmt.Println("NFTReceived")
	txHash := vLog.TxHash.Hex()
	blockNumber := vLog.BlockNumber
	signature := vLog.Topics[0].Hex()

	lender := fmt.Sprintf("%s", eventData["from"])
	tokenId := fmt.Sprintf("%v", eventData["tokenId"])

	event := Event{
		Lender:          lender,
		Borrower:        "",
		TransactionHash: txHash,
		BlockNumber:     blockNumber,
		TokenAddress:    "",
		TokenId:         tokenId,
		Signature:       signature,
	}
	fmt.Println("EVENT:", event)

	_, err := db.Exec(sqlStatement, lender, "", "", tokenId, txHash, blockNumber, "NFTReceived")
	if err != nil {
		panic(err)
	}
}

func handleNFTCancelled(vLog types.Log, eventData map[string]interface{}, db *sql.DB) {
	fmt.Println("NFTCancelled")

	txHash := vLog.TxHash.Hex()
	blockNumber := vLog.BlockNumber
	signature := vLog.Topics[0].Hex()

	lender := vLog.Topics[1].Hex()
	tokenId := fmt.Sprintf("%v", eventData["tokenId"])
	tokenAddress := common.HexToAddress(vLog.Topics[2].Hex())

	event := Event{
		Lender:          lender,
		Borrower:        "",
		TransactionHash: txHash,
		BlockNumber:     blockNumber,
		TokenAddress:    tokenAddress.Hex(),
		TokenId:         tokenId,
		Signature:       signature,
	}
	fmt.Println("EVENT:", event)

	_, err := db.Exec(sqlStatement, lender, "", tokenAddress.Hex(), tokenId, txHash, blockNumber, "NFTReceived")
	if err != nil {
		panic(err)
	}
}

func handleNFTBorrowed(vLog types.Log, eventData map[string]interface{}, db *sql.DB) {
	fmt.Println("NFTBorrowed")

	txHash := vLog.TxHash.Hex()
	blockNumber := vLog.BlockNumber
	signature := vLog.Topics[0].Hex()

	tokenId := fmt.Sprintf("%v", eventData["tokenId"])
	tokenAddress := common.HexToAddress(vLog.Topics[3].Hex())
	borrower := vLog.Topics[1].Hex()
	lender := vLog.Topics[2].Hex()

	fmt.Println("Event received:")
	fmt.Println("  Transaction Hash:", txHash)
	fmt.Println("  Block Number:", blockNumber)
	fmt.Println("  Event Signature:", signature)
	fmt.Println("  Token Address:", tokenAddress.Hex())

	event := Event{
		Lender:          lender,
		Borrower:        borrower,
		TransactionHash: txHash,
		BlockNumber:     blockNumber,
		TokenAddress:    tokenAddress.Hex(),
		TokenId:         tokenId,
		Signature:       signature,
	}
	fmt.Println("EVENT:", event)

	_, err := db.Exec(sqlStatement, lender, borrower, tokenAddress.Hex(), tokenId, txHash, blockNumber, "NFTReceived")
	if err != nil {
		panic(err)
	}
}

func handleNFTReturned(vLog types.Log, eventData map[string]interface{}, db *sql.DB) {
	fmt.Println("NFTReturned")

	txHash := vLog.TxHash.Hex()
	blockNumber := vLog.BlockNumber
	signature := vLog.Topics[0].Hex()
	tokenId := fmt.Sprintf("%v", eventData["tokenId"])

	tokenAddress := common.HexToAddress(vLog.Topics[3].Hex())
	borrower := vLog.Topics[1].Hex()
	lender := vLog.Topics[2].Hex()

	event := Event{
		Lender:          lender,
		Borrower:        borrower,
		TransactionHash: txHash,
		BlockNumber:     blockNumber,
		TokenAddress:    tokenAddress.Hex(),
		TokenId:         tokenId,
		Signature:       signature,
	}
	fmt.Println("EVENT:", event)

	_, err := db.Exec(sqlStatement, lender, borrower, tokenAddress.Hex(), tokenId, txHash, blockNumber, "NFTReceived")
	if err != nil {
		panic(err)
	}
}

func handleNFTWithdrawn(vLog types.Log, eventData map[string]interface{}, db *sql.DB) {
	fmt.Println("NFTWithdrawn")

	txHash := vLog.TxHash.Hex()
	blockNumber := vLog.BlockNumber
	signature := vLog.Topics[0].Hex()

	tokenId := fmt.Sprintf("%v", eventData["tokenId"])
	tokenAddress := common.HexToAddress(vLog.Topics[3].Hex())
	lender := vLog.Topics[1].Hex()

	event := Event{
		Lender:          lender,
		Borrower:        "",
		TransactionHash: txHash,
		BlockNumber:     blockNumber,
		TokenAddress:    tokenAddress.Hex(),
		TokenId:         tokenId,
		Signature:       signature,
	}
	fmt.Println("EVENT:", event)

	_, err := db.Exec(sqlStatement, lender, "", tokenAddress.Hex(), tokenId, txHash, blockNumber, "NFTReceived")
	if err != nil {
		panic(err)
	}
}

func handleTransfer(vLog types.Log, eventData map[string]interface{}, db *sql.DB) {
	fmt.Println("Transfer")

	fmt.Println("vLog:", vLog)
	fmt.Println("eventData:", eventData)
	fmt.Println("vLog Topics:", vLog.Topics)
}

func handleEvents(client *ethclient.Client, db *sql.DB, contractAddress common.Address, contractAbi abi.ABI) {
	// Create an event filter
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

			switch vLog.Topics[0].Hex() {
			case "0x75ee001158df8b77347acda2c33d52e5d6facd0c4331fd0910a4b5eb3993369a": // NFTAdded
				handleNFTAdded(vLog, eventData, db)
			case "0x1fc64ec2285c9891a5f6513865cacc14559dbec84a1dd36395483a27fd06324d": // NFTCancelled
				handleNFTCancelled(vLog, eventData, db)
			case "0x018587ee2904cbc9583f54aa70a102990f241760f13f7a24ae0de20693487d7b": // NFTBorrowed
				handleNFTBorrowed(vLog, eventData, db)
			case "0xc0bd3c824ba6fcb0d28b5548f84a231d7252efb9252c44196bf4e4ee7323ef33": // NFTReturned
				handleNFTReturned(vLog, eventData, db)
			case "0xbbde41973f9ce4890f7ad9762c23d8191f261fd643bdf13ed8bbc10549b49fcb": // NFTWithdrawn
				handleNFTWithdrawn(vLog, eventData, db)
			case "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef": // Transfer
				handleTransfer(vLog, eventData, db)
			}
		}
	}
}

func main() {
	// Sepolia testnet endpoint
	sepoliaEndpoint := "wss://sepolia.infura.io/ws/v3/bad30ec837404a23978e17118a2c860b"

	psqlInfo := "host=localhost user=postgres password=postgres dbname=testdb sslmode=disable"

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	// Ethereum client initialization
	client, err := ethclient.Dial(sepoliaEndpoint)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Contract address and ABI
	contractAddress := common.HexToAddress("0x7ed82e52689d7c542c3f8ca255cd921c6fc24e27")

	contractABI := readAbi("D:\\GoProjects\\GoBackend\\abi.json")

	handleEvents(client, db, contractAddress, contractABI)
}
