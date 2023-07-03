package main

import (
	"log"

	"GethBackServ/internal/endpoint/abigencontract"
	"GethBackServ/internal/endpoint/httphandler"
	"GethBackServ/internal/service/contracthandler"
	"GethBackServ/internal/service/database"

	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	ethInfo, err := contracthandler.GetEthClientInfo()
	if err != nil {
		log.Fatal("Error:", err)
	}

	db, err := database.GetConnection()
	if err != nil {
		log.Fatal("Error:", err)
	}

	database.CreateTable(db.DB, "events")

	// Abigen
	contractAbigen, err := abigencontract.NewMainFilterer(ethInfo.ContractAddress, ethInfo.Client)
	if err != nil {
		log.Fatal("Error:", err)
	}

	// Handle missed events
	contracthandler.HandleMissedEvents(ethInfo.Client, db.DB, ethInfo.ContractAddress, contractAbigen)

	// Handle missed transfers and subscribe on transfers events
	contracthandler.HandleMissedTransfers(ethInfo, db.DB)

	// Start listening events in real time
	go contracthandler.HandleEvents(ethInfo.Client, db.DB, ethInfo.ContractAddress, contractAbigen)

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Static
	e.Static("/", "static")

	// Routes
	e.GET("/nftTransfers/:tokenAddress/:tokenId", httphandler.GetTransfersHandler)
	e.GET("/nftHistory/:tokenAddress/:tokenId", httphandler.GetNFTHistoryHandler)
	e.GET("/nftHistoryByWalletAddress/:walletAddress", httphandler.GetNFTHistoryByWalletAddressHandler)
	e.GET("/nfts", httphandler.GetNFTsHandler)

	// Start server
	e.Start(":8080")
}
