package httphandler

import (
	"math/big"
	"net/http"
	"strconv"

	"GethBackServ/internal/service/contracthandler"
	"GethBackServ/internal/service/database"

	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
)

func GetTransfersHandler(c echo.Context) error {
	tokenAddress := common.HexToAddress(c.Param("tokenAddress"))
	tokenId, err := strconv.ParseInt(c.Param("tokenId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	transfers, err := contracthandler.GetTransfersByAddress(tokenAddress, tokenId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, transfers)
}

func GetNFTHistoryHandler(c echo.Context) error {
	db, err := database.GetConnection()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	ethInfo, err := contracthandler.GetEthClientInfo()
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

	history, err := contracthandler.GetNFTHistory(ethInfo.ContractAddress, tokenAddress, big.NewInt(tokenId), common.HexToAddress("0x0000000000000000000000000000000000000000"), db.DB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, history)
}

func GetNFTsHandler(c echo.Context) error {
	ethInfo, err := contracthandler.GetEthClientInfo()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	nftInfoList, err := contracthandler.GetNFTs(ethInfo.Client, ethInfo.ContractAddress)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "NFTs not found",
		})
	}

	return c.JSON(http.StatusOK, nftInfoList)
}

func GetNFTHistoryByWalletAddressHandler(c echo.Context) error {
	db, err := database.GetConnection()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	ethInfo, err := contracthandler.GetEthClientInfo()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	walletAddress := common.HexToAddress(c.Param("walletAddress"))
	history, err := contracthandler.GetNFTHistory(ethInfo.ContractAddress, common.HexToAddress("0x0000000000000000000000000000000000000000"), big.NewInt(-1), walletAddress, db.DB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, history)
}
