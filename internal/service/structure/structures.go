package structure

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
