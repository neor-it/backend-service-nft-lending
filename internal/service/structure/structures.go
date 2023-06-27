package structure

import (
	"database/sql"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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
	FromAddress     string `json:"from"`
	ToAddress       string `json:"to"`
	TokenAddress    string `json:"tokenAddress"`
	TokenId         string `json:"tokenId"`
	TransactionHash string `json:"transactionHash"`
	BlockNumber     uint64 `json:"blockNumber"`
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

type DBConnection struct {
	DB *sql.DB
}

type EthClientInfo struct {
	Client          *ethclient.Client
	ContractAddress common.Address
}

var SubscriptionMap = make(map[common.Address]bool)
