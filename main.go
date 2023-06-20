package main

import (
	"log"

	"GethBackServ/internal/endpoint/httphandler"
	"GethBackServ/internal/service/contracthandler"
	"GethBackServ/internal/service/database"

	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	client, contractAddress, err := contracthandler.GetEthClientAndAddress()
	if err != nil {
		log.Fatal("Error:", err)
	}

	db := database.ConnectToDB()
	defer db.Close()

	database.CreateTable(db, "events")

	contractAbi := contracthandler.ReadAbi("./api/abi.json")

	// Handle missed events
	contracthandler.HandleMissedEvents(client, db, contractAddress, contractAbi)

	// Start listening events in real time
	go contracthandler.HandleEvents(client, db, contractAddress, contractAbi)

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
