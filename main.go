package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	Owner           string `json:"owner"`
	TokenId         string `json:"tokenId"`
	TokenAddress    string `json:"nftContract"`
	Contract        string `json:"contract"`
	TransactionHash string `json:"transactionHash"`
	BlockNumber     uint64 `json:"blockNumber"`
	Borrower        string `json:"borrower"`
	Lender          string `json:"lender"`
	Signature       string `json:"signature"`
}

func getEthClientAndAddress() (*ethclient.Client, common.Address, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	API_KEY := os.Getenv("API_KEY")
	log.Println("API_KEY:", API_KEY)
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/" + API_KEY)
	if err != nil {
		return nil, common.Address{}, err
	}

	contractAddress := common.HexToAddress("0x22b63f333dB05DC4ead6c781349893378ed77F70")

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

func callGetAllNFTs(client *ethclient.Client, contractAbi abi.ABI, contractAddress common.Address) ([]interface{}, error) {
	getAllNFTs, err := contractAbi.Pack("getAllNFTs")
	if err != nil {
		return nil, err
	}

	result, err := client.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contractAddress,
		Data: getAllNFTs,
	}, nil)

	if err != nil {
		return nil, err
	}

	return contractAbi.Unpack("getAllNFTs", result)
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

func getNFTs(client *ethclient.Client, contractAbi abi.ABI, contractAddress common.Address) ([]NFTInfo, error) {
	data, err := callGetAllNFTs(client, contractAbi, contractAddress)

	if err != nil {
		log.Fatal(err)
	}

	var nftInfoList []NFTInfo

	for i := range data {
		jsonData, err := json.Marshal(data[i])
		if err != nil {
			log.Println("Ошибка при преобразовании в JSON:", err)
			return nil, err
		}

		jsonData = jsonData[1 : len(jsonData)-1] // delete [ and ] from string

		var nftInfo NFTInfo

		err = json.Unmarshal(jsonData, &nftInfo)

		if err != nil {
			log.Println("Ошибка при разборе JSON:", err)
			return nil, err
		}

		nftInfoList = append(nftInfoList, nftInfo)
	}

	return nftInfoList, nil
}

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

	history := getNFTHistory(client, contractAddress, tokenAddress, big.NewInt(tokenId), common.HexToAddress("0x0000000000000000000000000000000000000000"))

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
	client, contractAddress, err := getEthClientAndAddress()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	walletAddress := common.HexToAddress(c.Param("walletAddress"))

	history := getNFTHistory(client, contractAddress, common.HexToAddress("0x0000000000000000000000000000000000000000"), big.NewInt(-1), walletAddress)

	return c.JSON(http.StatusOK, history)
}

func main() {
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
