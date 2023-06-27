package contracthandler

import (
	"encoding/json"

	"GethBackServ/internal/endpoint/abigencontract"
	"GethBackServ/internal/service/structure"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func callGetAllNFTs(client *ethclient.Client, contractInstance *abigencontract.Main) ([]abigencontract.Nft, error) {
	result, err := contractInstance.GetAllNFTs(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetNFTs(client *ethclient.Client, contractAddress common.Address) ([]structure.NFTInfo, error) {
	contractInstance, err := abigencontract.NewMain(contractAddress, client)
	if err != nil {
		return nil, err
	}

	data, err := callGetAllNFTs(client, contractInstance)
	if err != nil {
		return nil, err
	}

	var nftInfoList []structure.NFTInfo

	for _, nft := range data {
		jsonData, err := json.Marshal(nft)
		if err != nil {
			return nil, err
		}

		var nftInfo structure.NFTInfo

		err = json.Unmarshal(jsonData, &nftInfo)
		if err != nil {
			return nil, err
		}

		nftInfoList = append(nftInfoList, nftInfo)
	}

	return nftInfoList, nil
}
