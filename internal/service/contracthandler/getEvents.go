package contracthandler

import (
	"database/sql"
	"math/big"
	"strings"

	"GethBackServ/internal/service/database"
	"GethBackServ/internal/service/structure"

	"github.com/ethereum/go-ethereum/common"
	_ "github.com/lib/pq"
)

func GetNFTHistory(contractAddress common.Address, tokenAddress common.Address, tokenId *big.Int, walletAddress common.Address, db *sql.DB) ([]structure.Event, error) {
	var history []structure.Event

	signatures := [][]byte{
		[]byte("NFTAdded"),
		[]byte("NFTWithdrawn"),
		[]byte("NFTBorrowed"),
		[]byte("NFTReturned"),
		[]byte("NFTCanceled"),
	}

	allEvents := make([]structure.Event, 0)

	for _, signature := range signatures {
		nftEvents, err := database.TrackEvents(contractAddress, signature, db)
		if err != nil {
			return nil, err
		}
		allEvents = append(allEvents, nftEvents...) // unpack slice of slices into one slice of events
	}

	for _, event := range allEvents {
		event.Borrower = structure.NormalizeAddress(event.Borrower)
		event.Lender = structure.NormalizeAddress(event.Lender)

		walletAddress = common.HexToAddress(strings.ToLower(walletAddress.String()))

		if ((event.TokenAddress == tokenAddress.Hex() && event.TokenId == tokenId.String()) && walletAddress.Hex() == "0x0000000000000000000000000000000000000000") ||
			((event.Borrower == walletAddress.Hex() || event.Lender == walletAddress.Hex()) && walletAddress.Hex() != "0x0000000000000000000000000000000000000000") {
			history = append(history, event)
		}
	}

	return history, nil
}

func GetTransfersByAddress(tokenAddress common.Address, tokenId int64) ([]structure.Transfers, error) {
	var transfers []structure.Transfers

	db, err := database.GetConnection()
	if err != nil {
		return transfers, err
	}

	transfers, err = database.TrackTransfers(tokenAddress, tokenId, db.DB)
	if err != nil {
		return transfers, err
	}

	return transfers, nil
}
