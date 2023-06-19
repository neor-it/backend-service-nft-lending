package contracthandler

import (
	"context"
	"log"
	"math/big"

	"GethBackServ/structure"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/lib/pq"
)

func GetTransfersByAddress(client *ethclient.Client, filterAddress common.Address, tokenAddress common.Address, tokenId *big.Int) []structure.Transfers {
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

	var transfers []structure.Transfers

	for _, log := range logs {
		fromAddress := common.BytesToAddress(log.Topics[1].Bytes())
		toAddress := common.BytesToAddress(log.Topics[2].Bytes())
		tokenIdLog := new(big.Int).SetBytes(log.Topics[3].Bytes())

		if (fromAddress == filterAddress || toAddress == filterAddress) && tokenIdLog.Cmp(tokenId) == 0 {
			transfers = append(transfers, structure.Transfers{
				FromAddress: fromAddress.Hex(),
				ToAddress:   toAddress.Hex(),
				TokenID:     int(tokenIdLog.Int64()),
			})
		}
	}

	return transfers
}
