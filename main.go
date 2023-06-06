package main

import (
	"bytes"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
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
	godotenv.Load()

	API_KEY := os.Getenv("API_KEY")
	log.Println("API_KEY:", API_KEY)
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/" + API_KEY)
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
