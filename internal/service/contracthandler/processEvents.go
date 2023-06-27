package contracthandler

import (
	"GethBackServ/internal/endpoint/abigencontract"
	"GethBackServ/internal/service/database"
	"database/sql"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func processNFTAddedEvents(contractAbigen *abigencontract.MainFilterer, filterQuery *bind.FilterOpts, db *sql.DB) error {
	addedLogs, err := contractAbigen.FilterNFTAdded(filterQuery, nil, nil)
	if err != nil {
		return err
	}
	defer addedLogs.Close()

	for addedLogs.Next() {
		event := addedLogs.Event
		database.HandleNFTAdded(event, db)
	}

	if err := addedLogs.Error(); err != nil {
		return err
	}

	return nil
}

func processNFTCanceledEvents(contractAbigen *abigencontract.MainFilterer, filterQuery *bind.FilterOpts, db *sql.DB) error {
	canceledLogs, err := contractAbigen.FilterNFTCanceled(filterQuery, nil, nil)
	if err != nil {
		return err
	}
	defer canceledLogs.Close()

	for canceledLogs.Next() {
		event := canceledLogs.Event
		database.HandleNFTCanceled(event, db)
	}

	if err := canceledLogs.Error(); err != nil {
		return err
	}

	return nil
}

func processNFTReturnedEvents(contractAbigen *abigencontract.MainFilterer, filterQuery *bind.FilterOpts, db *sql.DB) error {
	returnedLogs, err := contractAbigen.FilterNFTReturned(filterQuery, nil, nil, nil)
	if err != nil {
		return err
	}
	defer returnedLogs.Close()

	for returnedLogs.Next() {
		event := returnedLogs.Event
		database.HandleNFTReturned(event, db)
	}

	if err := returnedLogs.Error(); err != nil {
		return err
	}

	return nil
}

func processNFTWithdrawnEvents(contractAbigen *abigencontract.MainFilterer, filterQuery *bind.FilterOpts, db *sql.DB) error {
	withdrawnLogs, err := contractAbigen.FilterNFTWithdrawn(filterQuery, nil, nil)
	if err != nil {
		return err
	}
	defer withdrawnLogs.Close()

	for withdrawnLogs.Next() {
		event := withdrawnLogs.Event
		database.HandleNFTWithdrawn(event, db)
	}

	if err := withdrawnLogs.Error(); err != nil {
		return err
	}

	return nil
}

func processNFTBorrowedEvents(contractAbigen *abigencontract.MainFilterer, filterQuery *bind.FilterOpts, db *sql.DB) error {
	borrowedLogs, err := contractAbigen.FilterNFTBorrowed(filterQuery, nil, nil, nil)
	if err != nil {
		return err
	}
	defer borrowedLogs.Close()

	for borrowedLogs.Next() {
		event := borrowedLogs.Event
		database.HandleNFTBorrowed(event, db)
	}

	if err := borrowedLogs.Error(); err != nil {
		return err
	}

	return nil
}
