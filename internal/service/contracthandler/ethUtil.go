package contracthandler

import (
	"bytes"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func GetEthClientAndAddress() (*ethclient.Client, common.Address, error) {
	godotenv.Load()

	API_KEY := os.Getenv("API_KEY")
	client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/" + API_KEY)
	if err != nil {
		return nil, common.Address{}, err
	}

	contractAddress := common.HexToAddress("0x7ed82e52689d7c542c3f8ca255cd921c6fc24e27")

	return client, contractAddress, nil
}

func ReadAbi(fileName string) abi.ABI {
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
