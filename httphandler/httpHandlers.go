package httphandler

import (
	"math/big"
	"net/http"
	"strconv"

	"GethBackServ/contracthandler"
	"GethBackServ/database"

	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
)

func GetTransfersHandler(c echo.Context) error {
	client, contractAddress, err := contracthandler.GetEthClientAndAddress()
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

	transfers := contracthandler.GetTransfersByAddress(client, contractAddress, tokenAddress, big.NewInt(tokenId))

	return c.JSON(http.StatusOK, transfers)
}

func GetNFTHistoryHandler(c echo.Context) error {
	db := database.ConnectToDB()
	defer db.Close()

	client, contractAddress, err := contracthandler.GetEthClientAndAddress()
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

	history := contracthandler.GetNFTHistory(client, contractAddress, tokenAddress, big.NewInt(tokenId), common.HexToAddress("0x0000000000000000000000000000000000000000"), db)

	return c.JSON(http.StatusOK, history)
}

func GetNFTsHandler(c echo.Context) error {
	client, contractAddress, err := contracthandler.GetEthClientAndAddress()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	contractAbi := contracthandler.ReadAbi("abi.json")

	nftInfoList, err := contracthandler.GetNFTs(client, contractAbi, contractAddress)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "NFTs not found",
		})
	}

	return c.JSON(http.StatusOK, nftInfoList)

}

func GetNFTHistoryByWalletAddressHandler(c echo.Context) error {
	db := database.ConnectToDB()
	defer db.Close()

	client, contractAddress, err := contracthandler.GetEthClientAndAddress()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	walletAddress := common.HexToAddress(c.Param("walletAddress"))
	history := contracthandler.GetNFTHistory(client, contractAddress, common.HexToAddress("0x0000000000000000000000000000000000000000"), big.NewInt(-1), walletAddress, db)

	return c.JSON(http.StatusOK, history)
}
