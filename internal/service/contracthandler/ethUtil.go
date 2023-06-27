package contracthandler

import (
	"GethBackServ/internal/service/structure"
	"bytes"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

var instance *structure.EthClientInfo

func GetEthClientInfo() (*structure.EthClientInfo, error) {
	if instance == nil {
		godotenv.Load()

		API_KEY := os.Getenv("API_KEY")
		client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/" + API_KEY)
		if err != nil {
			return nil, err
		}
		contractAddress := common.HexToAddress("0x7ed82e52689d7c542c3f8ca255cd921c6fc24e27")

		instance = &structure.EthClientInfo{
			Client:          client,
			ContractAddress: contractAddress,
		}

	}
	return instance, nil
}
