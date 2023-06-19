package contracthandler

import (
	"context"
	"encoding/json"
	"log"

	"GethBackServ/structure"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

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

func GetNFTs(client *ethclient.Client, contractAbi abi.ABI, contractAddress common.Address) ([]structure.NFTInfo, error) {
	data, err := callGetAllNFTs(client, contractAbi, contractAddress)

	if err != nil {
		log.Fatal(err)
	}

	var nftInfoList []structure.NFTInfo

	for i := range data {
		jsonData, err := json.Marshal(data[i])
		if err != nil {
			log.Println("Error marshaling data:", err)
			return nil, err
		}

		jsonData = jsonData[1 : len(jsonData)-1] // delete [ and ] from string

		var nftInfo structure.NFTInfo

		err = json.Unmarshal(jsonData, &nftInfo)

		if err != nil {
			log.Println("Error unmarshaling data:", err)
			return nil, err
		}

		nftInfoList = append(nftInfoList, nftInfo)
	}

	return nftInfoList, nil
}
