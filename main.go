package main

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	sqlStatement = `INSERT INTO events (lender, borrower, tokenAddress, tokenId, transactionHash, blockNumber, signature)
	VALUES ($1, $2, $3, $4, $5, $6, $7) `
)

type NFTInfo struct {
	OwnerAddress    string `json:"owner"`
	NewOwnerAddress string `json:"newOwner"`
	TokenID         int    `json:"tokenId"`
	AddressNFT      string `json:"nftContract"`
	AmountInWei     int    `json:"nftValue"`
	AmountInUSDT    int    `json:"usdtValue"`
	UseTime         int    `json:"useTime"`
	Timestamp       int    `json:"timestamp"`
	IsAvailable     bool   `json:"isAvailable"`
}

type Transfers struct {
	FromAddress string `json:"from"`
	ToAddress   string `json:"to"`
	TokenID     int    `json:"tokenId"`
}

type Event struct {
	Lender          string `json:"lender"`
	Borrower        string `json:"borrower"`
	TokenId         string `json:"tokenId"`
	TokenAddress    string `json:"nftContract"`
	TransactionHash string `json:"transactionHash"`
	BlockNumber     uint64 `json:"blockNumber"`
	Signature       string `json:"signature"`
}

func getEthClientAndAddress() (*ethclient.Client, common.Address, error) {
	godotenv.Load()

	API_KEY := os.Getenv("API_KEY")
	client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/" + API_KEY)
	if err != nil {
		return nil, common.Address{}, err
	}

	contractAddress := common.HexToAddress("0x7ed82e52689d7c542c3f8ca255cd921c6fc24e27")

	return client, contractAddress, nil
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

func getTransfersHandler(c echo.Context) error {
	client, contractAddress, err := getEthClientAndAddress()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	tokenAddress := common.HexToAddress(c.Param("tokenAddress"))
	tokenId, err := strconv.ParseInt(c.Param("tokenId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	transfers := getTransfersByAddress(client, contractAddress, tokenAddress, big.NewInt(tokenId))

	return c.JSON(http.StatusOK, transfers)
}

func getNFTHistoryHandler(c echo.Context) error {
	db := connectToDB()
	defer db.Close()

	client, contractAddress, err := getEthClientAndAddress()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	tokenAddress := common.HexToAddress(c.Param("tokenAddress"))
	tokenId, err := strconv.ParseInt(c.Param("tokenId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	history := getNFTHistory(client, contractAddress, tokenAddress, big.NewInt(tokenId), common.HexToAddress("0x0000000000000000000000000000000000000000"), db)

	return c.JSON(http.StatusOK, history)
}

func getNFTsHandler(c echo.Context) error {
	client, contractAddress, err := getEthClientAndAddress()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	contractAbi := readAbi("abi.json")

	nftInfoList, err := getNFTs(client, contractAbi, contractAddress)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "NFTs not found",
		})
	}

	return c.JSON(http.StatusOK, nftInfoList)

}

func getNFTHistoryByWalletAddressHandler(c echo.Context) error {
	db := connectToDB()
	defer db.Close()

	client, contractAddress, err := getEthClientAndAddress()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	walletAddress := common.HexToAddress(c.Param("walletAddress"))
	history := getNFTHistory(client, contractAddress, common.HexToAddress("0x0000000000000000000000000000000000000000"), big.NewInt(-1), walletAddress, db)

	return c.JSON(http.StatusOK, history)
}

func handleNFTAdded(vLog types.Log, eventData map[string]interface{}, db *sql.DB) {
	txHash := vLog.TxHash.Hex()
	blockNumber := vLog.BlockNumber

	tokenAddress := common.HexToAddress(vLog.Topics[2].Hex())
	tokenId := fmt.Sprintf("%v", eventData["tokenId"])
	lender := vLog.Topics[1].Hex()

	_, err := db.Exec(sqlStatement, lender, "", tokenAddress.Hex(), tokenId, txHash, blockNumber, "NFTAdded")
	if err != nil {
		panic(err)
	}
}

func handleNFTCanceled(vLog types.Log, eventData map[string]interface{}, db *sql.DB) {

	txHash := vLog.TxHash.Hex()
	blockNumber := vLog.BlockNumber

	lender := vLog.Topics[1].Hex()
	tokenId := fmt.Sprintf("%v", eventData["tokenId"])
	tokenAddress := common.HexToAddress(vLog.Topics[2].Hex())

	_, err := db.Exec(sqlStatement, lender, "", tokenAddress.Hex(), tokenId, txHash, blockNumber, "NFTCanceled")
	if err != nil {
		panic(err)
	}
}

func handleNFTBorrowed(vLog types.Log, eventData map[string]interface{}, db *sql.DB) {
	txHash := vLog.TxHash.Hex()
	blockNumber := vLog.BlockNumber

	tokenId := fmt.Sprintf("%v", eventData["tokenId"])
	tokenAddress := common.HexToAddress(vLog.Topics[3].Hex())
	borrower := vLog.Topics[1].Hex()
	lender := vLog.Topics[2].Hex()

	_, err := db.Exec(sqlStatement, lender, borrower, tokenAddress.Hex(), tokenId, txHash, blockNumber, "NFTBorrowed")
	if err != nil {
		panic(err)
	}
}

func handleNFTReturned(vLog types.Log, eventData map[string]interface{}, db *sql.DB) {
	txHash := vLog.TxHash.Hex()
	blockNumber := vLog.BlockNumber
	tokenId := fmt.Sprintf("%v", eventData["tokenId"])

	tokenAddress := common.HexToAddress(vLog.Topics[3].Hex())
	borrower := vLog.Topics[1].Hex()
	lender := vLog.Topics[2].Hex()

	_, err := db.Exec(sqlStatement, lender, borrower, tokenAddress.Hex(), tokenId, txHash, blockNumber, "NFTReturned")
	if err != nil {
		panic(err)
	}
}

func handleNFTWithdrawn(vLog types.Log, eventData map[string]interface{}, db *sql.DB) {
	txHash := vLog.TxHash.Hex()
	blockNumber := vLog.BlockNumber

	tokenId := fmt.Sprintf("%v", eventData["tokenId"])
	tokenAddress := common.HexToAddress(vLog.Topics[3].Hex())
	lender := vLog.Topics[1].Hex()

	_, err := db.Exec(sqlStatement, lender, "", tokenAddress.Hex(), tokenId, txHash, blockNumber, "NFTWithdrawn")
	if err != nil {
		panic(err)
	}
}

func handleEvent(vLog types.Log, eventData map[string]interface{}, db *sql.DB) {
	switch vLog.Topics[0].Hex() {
	case crypto.Keccak256Hash([]byte("NFTAdded(address,address,uint256)")).Hex():
		handleNFTAdded(vLog, eventData, db)
	case crypto.Keccak256Hash([]byte("NFTCanceled(address,address,uint256)")).Hex():
		handleNFTCanceled(vLog, eventData, db)
	case crypto.Keccak256Hash([]byte("NFTBorrowed(address,address,address,uint256)")).Hex():
		handleNFTBorrowed(vLog, eventData, db)
	case crypto.Keccak256Hash([]byte("NFTReturned(address,address,address,uint256)")).Hex():
		handleNFTReturned(vLog, eventData, db)
	case crypto.Keccak256Hash([]byte("NFTWithdrawn(address,address,uint256)")).Hex():
		handleNFTWithdrawn(vLog, eventData, db)
	}
}

func handleEvents(client *ethclient.Client, db *sql.DB, contractAddress common.Address, contractAbi abi.ABI) {
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

			handleEvent(vLog, eventData, db)
		}
	}
}

func getLastProcessedBlockNumber(db *sql.DB) int64 {
	var blockNumber int64
	err := db.QueryRow("SELECT blocknumber FROM events ORDER BY blocknumber DESC LIMIT 1").Scan(&blockNumber)

	if err != nil {
		if err == sql.ErrNoRows {
			// table is empty
			blockNumber = 0
		} else { // unexpected error
			panic(err)
		}
	}

	return blockNumber
}

func handlePastEvents(client *ethclient.Client, db *sql.DB, contractAddress common.Address, contractAbi abi.ABI) {
	blockNumber := getLastProcessedBlockNumber(db)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			contractAddress,
		},
		FromBlock: big.NewInt(blockNumber),
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatalf("Failed to get past events: %v", err)
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

		handleEvent(vLog, eventData, db)
	}
}

func createTable(db *sql.DB, tableName string) {
	// check if table exists

	if _, err := db.Exec("SELECT 1 FROM " + tableName + " LIMIT 1"); err != nil {
		log.Printf("Table %s doesn't exist, creating...", tableName)
		// create table
		sqlStatement := `
		CREATE TABLE events (
			id SERIAL PRIMARY KEY,
			lender TEXT,
			borrower TEXT,
			tokenAddress TEXT,
			tokenId TEXT,
			transactionHash TEXT,
			blockNumber INTEGER,
			signature TEXT
		);`

		_, err = db.Exec(sqlStatement)
		if err != nil {
			panic(err)
		}

		log.Println("Successfully created table!")
		return
	}

	log.Printf("Table %s exists!", tableName)
}

func connectToDB() *sql.DB {
	godotenv.Load()

	POSTGRES_HOST := os.Getenv("POSTGRES_HOST")
	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_DB := os.Getenv("POSTGRES_DB")

	sqlData := "host=" + POSTGRES_HOST + " user=" + POSTGRES_USER + " password=" + POSTGRES_PASSWORD + " dbname=" + POSTGRES_DB + " sslmode=disable"

	db, err := sql.Open("postgres", sqlData)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Successfully connected!")
	return db
}

func main() {
	client, contractAddress, err := getEthClientAndAddress()
	if err != nil {
		log.Fatal("Error:", err)
	}

	db := connectToDB()
	defer db.Close()

	createTable(db, "events")

	contractAbi := readAbi("abi.json")

	// Handle missed events
	handlePastEvents(client, db, contractAddress, contractAbi)

	// Start listening to events in real time
	go handleEvents(client, db, contractAddress, contractAbi)

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Static
	e.Static("/", "static")

	// Routes
	e.GET("/nftTransfers/:tokenAddress/:tokenId", getTransfersHandler)
	e.GET("/nftHistory/:tokenAddress/:tokenId", getNFTHistoryHandler)
	e.GET("/nftHistoryByWalletAddress/:walletAddress", getNFTHistoryByWalletAddressHandler)
	e.GET("/nfts", getNFTsHandler)

	// Start server
	e.Start(":8080")
}
