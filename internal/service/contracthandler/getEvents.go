package contracthandler

import (
	"database/sql"
	"math/big"

	"GethBackServ/internal/service/database"
	"GethBackServ/internal/service/structure"

	"github.com/ethereum/go-ethereum/common"
	_ "github.com/lib/pq"
)

func GetNFTHistory(contractAddress common.Address, tokenAddress common.Address, tokenId *big.Int, walletAddress common.Address, db *sql.DB) ([]structure.Event, error) {
	nftEvents, err := database.TrackEvents(contractAddress, tokenAddress, tokenId, walletAddress, db)
	if err != nil {
		return nil, err
	}

	return nftEvents, nil
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
