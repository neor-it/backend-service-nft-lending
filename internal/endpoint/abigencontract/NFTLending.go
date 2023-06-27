// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abigencontract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// Nft is an auto generated low-level Go binding around an user-defined struct.
type Nft struct {
	Owner       common.Address
	NewOwner    common.Address
	TokenId     *big.Int
	NftContract common.Address
	NftValue    *big.Int
	UsdtValue   *big.Int
	UseTime     *big.Int
	Timestamp   *big.Int
	IsAvailable bool
}

// MainMetaData contains all meta data concerning the Main contract.
var MainMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"NFTAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"NFTAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"NFTAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"NFTBorrowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"NFTAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"NFTCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"NFTReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"NFTAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"NFTReturned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"NFTAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"NFTWithdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelPurposeNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fakeUSDT\",\"outputs\":[{\"internalType\":\"contractFakeUSDT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllNFTs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usdtValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"useTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isAvailable\",\"type\":\"bool\"}],\"internalType\":\"structNft[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nfts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usdtValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"useTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isAvailable\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"purchaseNFT\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"purchaseNFTWithUSDT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_useTime\",\"type\":\"uint256\"}],\"name\":\"purposeNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_useTime\",\"type\":\"uint256\"}],\"name\":\"purposeNFTWithUSDT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"returnNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fakeUSDTAddress\",\"type\":\"address\"}],\"name\":\"setFakeUSDTContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MainABI is the input ABI used to generate the binding from.
// Deprecated: Use MainMetaData.ABI instead.
var MainABI = MainMetaData.ABI

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// FakeUSDT is a free data retrieval call binding the contract method 0x04f7f39b.
//
// Solidity: function fakeUSDT() view returns(address)
func (_Main *MainCaller) FakeUSDT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "fakeUSDT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FakeUSDT is a free data retrieval call binding the contract method 0x04f7f39b.
//
// Solidity: function fakeUSDT() view returns(address)
func (_Main *MainSession) FakeUSDT() (common.Address, error) {
	return _Main.Contract.FakeUSDT(&_Main.CallOpts)
}

// FakeUSDT is a free data retrieval call binding the contract method 0x04f7f39b.
//
// Solidity: function fakeUSDT() view returns(address)
func (_Main *MainCallerSession) FakeUSDT() (common.Address, error) {
	return _Main.Contract.FakeUSDT(&_Main.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Main *MainCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Main *MainSession) Fee() (*big.Int, error) {
	return _Main.Contract.Fee(&_Main.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Main *MainCallerSession) Fee() (*big.Int, error) {
	return _Main.Contract.Fee(&_Main.CallOpts)
}

// GetAllNFTs is a free data retrieval call binding the contract method 0xe0391b09.
//
// Solidity: function getAllNFTs() view returns((address,address,uint256,address,uint256,uint256,uint256,uint256,bool)[])
func (_Main *MainCaller) GetAllNFTs(opts *bind.CallOpts) ([]Nft, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getAllNFTs")

	if err != nil {
		return *new([]Nft), err
	}

	out0 := *abi.ConvertType(out[0], new([]Nft)).(*[]Nft)

	return out0, err

}

// GetAllNFTs is a free data retrieval call binding the contract method 0xe0391b09.
//
// Solidity: function getAllNFTs() view returns((address,address,uint256,address,uint256,uint256,uint256,uint256,bool)[])
func (_Main *MainSession) GetAllNFTs() ([]Nft, error) {
	return _Main.Contract.GetAllNFTs(&_Main.CallOpts)
}

// GetAllNFTs is a free data retrieval call binding the contract method 0xe0391b09.
//
// Solidity: function getAllNFTs() view returns((address,address,uint256,address,uint256,uint256,uint256,uint256,bool)[])
func (_Main *MainCallerSession) GetAllNFTs() ([]Nft, error) {
	return _Main.Contract.GetAllNFTs(&_Main.CallOpts)
}

// Nfts is a free data retrieval call binding the contract method 0x43af356b.
//
// Solidity: function nfts(address , uint256 ) view returns(address owner, address newOwner, uint256 tokenId, address nftContract, uint256 nftValue, uint256 usdtValue, uint256 useTime, uint256 timestamp, bool isAvailable)
func (_Main *MainCaller) Nfts(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Owner       common.Address
	NewOwner    common.Address
	TokenId     *big.Int
	NftContract common.Address
	NftValue    *big.Int
	UsdtValue   *big.Int
	UseTime     *big.Int
	Timestamp   *big.Int
	IsAvailable bool
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "nfts", arg0, arg1)

	outstruct := new(struct {
		Owner       common.Address
		NewOwner    common.Address
		TokenId     *big.Int
		NftContract common.Address
		NftValue    *big.Int
		UsdtValue   *big.Int
		UseTime     *big.Int
		Timestamp   *big.Int
		IsAvailable bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NewOwner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.NftContract = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.NftValue = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.UsdtValue = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.UseTime = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.IsAvailable = *abi.ConvertType(out[8], new(bool)).(*bool)

	return *outstruct, err

}

// Nfts is a free data retrieval call binding the contract method 0x43af356b.
//
// Solidity: function nfts(address , uint256 ) view returns(address owner, address newOwner, uint256 tokenId, address nftContract, uint256 nftValue, uint256 usdtValue, uint256 useTime, uint256 timestamp, bool isAvailable)
func (_Main *MainSession) Nfts(arg0 common.Address, arg1 *big.Int) (struct {
	Owner       common.Address
	NewOwner    common.Address
	TokenId     *big.Int
	NftContract common.Address
	NftValue    *big.Int
	UsdtValue   *big.Int
	UseTime     *big.Int
	Timestamp   *big.Int
	IsAvailable bool
}, error) {
	return _Main.Contract.Nfts(&_Main.CallOpts, arg0, arg1)
}

// Nfts is a free data retrieval call binding the contract method 0x43af356b.
//
// Solidity: function nfts(address , uint256 ) view returns(address owner, address newOwner, uint256 tokenId, address nftContract, uint256 nftValue, uint256 usdtValue, uint256 useTime, uint256 timestamp, bool isAvailable)
func (_Main *MainCallerSession) Nfts(arg0 common.Address, arg1 *big.Int) (struct {
	Owner       common.Address
	NewOwner    common.Address
	TokenId     *big.Int
	NftContract common.Address
	NftValue    *big.Int
	UsdtValue   *big.Int
	UseTime     *big.Int
	Timestamp   *big.Int
	IsAvailable bool
}, error) {
	return _Main.Contract.Nfts(&_Main.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Main *MainCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Main *MainSession) Owner() (common.Address, error) {
	return _Main.Contract.Owner(&_Main.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Main *MainCallerSession) Owner() (common.Address, error) {
	return _Main.Contract.Owner(&_Main.CallOpts)
}

// CancelPurposeNFT is a paid mutator transaction binding the contract method 0xe57e4e35.
//
// Solidity: function cancelPurposeNFT(address _nftContract, uint256 _tokenId) returns()
func (_Main *MainTransactor) CancelPurposeNFT(opts *bind.TransactOpts, _nftContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "cancelPurposeNFT", _nftContract, _tokenId)
}

// CancelPurposeNFT is a paid mutator transaction binding the contract method 0xe57e4e35.
//
// Solidity: function cancelPurposeNFT(address _nftContract, uint256 _tokenId) returns()
func (_Main *MainSession) CancelPurposeNFT(_nftContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Main.Contract.CancelPurposeNFT(&_Main.TransactOpts, _nftContract, _tokenId)
}

// CancelPurposeNFT is a paid mutator transaction binding the contract method 0xe57e4e35.
//
// Solidity: function cancelPurposeNFT(address _nftContract, uint256 _tokenId) returns()
func (_Main *MainTransactorSession) CancelPurposeNFT(_nftContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Main.Contract.CancelPurposeNFT(&_Main.TransactOpts, _nftContract, _tokenId)
}

// PurchaseNFT is a paid mutator transaction binding the contract method 0xa87e25ac.
//
// Solidity: function purchaseNFT(address _nftContract, uint256 _tokenId) payable returns()
func (_Main *MainTransactor) PurchaseNFT(opts *bind.TransactOpts, _nftContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "purchaseNFT", _nftContract, _tokenId)
}

// PurchaseNFT is a paid mutator transaction binding the contract method 0xa87e25ac.
//
// Solidity: function purchaseNFT(address _nftContract, uint256 _tokenId) payable returns()
func (_Main *MainSession) PurchaseNFT(_nftContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Main.Contract.PurchaseNFT(&_Main.TransactOpts, _nftContract, _tokenId)
}

// PurchaseNFT is a paid mutator transaction binding the contract method 0xa87e25ac.
//
// Solidity: function purchaseNFT(address _nftContract, uint256 _tokenId) payable returns()
func (_Main *MainTransactorSession) PurchaseNFT(_nftContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Main.Contract.PurchaseNFT(&_Main.TransactOpts, _nftContract, _tokenId)
}

// PurchaseNFTWithUSDT is a paid mutator transaction binding the contract method 0x84cb3eca.
//
// Solidity: function purchaseNFTWithUSDT(uint256 amount, address _nftContract, uint256 _tokenId) returns()
func (_Main *MainTransactor) PurchaseNFTWithUSDT(opts *bind.TransactOpts, amount *big.Int, _nftContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "purchaseNFTWithUSDT", amount, _nftContract, _tokenId)
}

// PurchaseNFTWithUSDT is a paid mutator transaction binding the contract method 0x84cb3eca.
//
// Solidity: function purchaseNFTWithUSDT(uint256 amount, address _nftContract, uint256 _tokenId) returns()
func (_Main *MainSession) PurchaseNFTWithUSDT(amount *big.Int, _nftContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Main.Contract.PurchaseNFTWithUSDT(&_Main.TransactOpts, amount, _nftContract, _tokenId)
}

// PurchaseNFTWithUSDT is a paid mutator transaction binding the contract method 0x84cb3eca.
//
// Solidity: function purchaseNFTWithUSDT(uint256 amount, address _nftContract, uint256 _tokenId) returns()
func (_Main *MainTransactorSession) PurchaseNFTWithUSDT(amount *big.Int, _nftContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Main.Contract.PurchaseNFTWithUSDT(&_Main.TransactOpts, amount, _nftContract, _tokenId)
}

// PurposeNFT is a paid mutator transaction binding the contract method 0x5891854a.
//
// Solidity: function purposeNFT(address _nftContract, uint256 _tokenId, uint256 _value, uint256 _useTime) returns()
func (_Main *MainTransactor) PurposeNFT(opts *bind.TransactOpts, _nftContract common.Address, _tokenId *big.Int, _value *big.Int, _useTime *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "purposeNFT", _nftContract, _tokenId, _value, _useTime)
}

// PurposeNFT is a paid mutator transaction binding the contract method 0x5891854a.
//
// Solidity: function purposeNFT(address _nftContract, uint256 _tokenId, uint256 _value, uint256 _useTime) returns()
func (_Main *MainSession) PurposeNFT(_nftContract common.Address, _tokenId *big.Int, _value *big.Int, _useTime *big.Int) (*types.Transaction, error) {
	return _Main.Contract.PurposeNFT(&_Main.TransactOpts, _nftContract, _tokenId, _value, _useTime)
}

// PurposeNFT is a paid mutator transaction binding the contract method 0x5891854a.
//
// Solidity: function purposeNFT(address _nftContract, uint256 _tokenId, uint256 _value, uint256 _useTime) returns()
func (_Main *MainTransactorSession) PurposeNFT(_nftContract common.Address, _tokenId *big.Int, _value *big.Int, _useTime *big.Int) (*types.Transaction, error) {
	return _Main.Contract.PurposeNFT(&_Main.TransactOpts, _nftContract, _tokenId, _value, _useTime)
}

// PurposeNFTWithUSDT is a paid mutator transaction binding the contract method 0x61169737.
//
// Solidity: function purposeNFTWithUSDT(address _nftContract, uint256 _tokenId, uint256 _value, uint256 _useTime) returns()
func (_Main *MainTransactor) PurposeNFTWithUSDT(opts *bind.TransactOpts, _nftContract common.Address, _tokenId *big.Int, _value *big.Int, _useTime *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "purposeNFTWithUSDT", _nftContract, _tokenId, _value, _useTime)
}

// PurposeNFTWithUSDT is a paid mutator transaction binding the contract method 0x61169737.
//
// Solidity: function purposeNFTWithUSDT(address _nftContract, uint256 _tokenId, uint256 _value, uint256 _useTime) returns()
func (_Main *MainSession) PurposeNFTWithUSDT(_nftContract common.Address, _tokenId *big.Int, _value *big.Int, _useTime *big.Int) (*types.Transaction, error) {
	return _Main.Contract.PurposeNFTWithUSDT(&_Main.TransactOpts, _nftContract, _tokenId, _value, _useTime)
}

// PurposeNFTWithUSDT is a paid mutator transaction binding the contract method 0x61169737.
//
// Solidity: function purposeNFTWithUSDT(address _nftContract, uint256 _tokenId, uint256 _value, uint256 _useTime) returns()
func (_Main *MainTransactorSession) PurposeNFTWithUSDT(_nftContract common.Address, _tokenId *big.Int, _value *big.Int, _useTime *big.Int) (*types.Transaction, error) {
	return _Main.Contract.PurposeNFTWithUSDT(&_Main.TransactOpts, _nftContract, _tokenId, _value, _useTime)
}

// ReturnNFT is a paid mutator transaction binding the contract method 0xccb568ec.
//
// Solidity: function returnNFT(address _nftContract, uint256 _tokenId) returns()
func (_Main *MainTransactor) ReturnNFT(opts *bind.TransactOpts, _nftContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "returnNFT", _nftContract, _tokenId)
}

// ReturnNFT is a paid mutator transaction binding the contract method 0xccb568ec.
//
// Solidity: function returnNFT(address _nftContract, uint256 _tokenId) returns()
func (_Main *MainSession) ReturnNFT(_nftContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Main.Contract.ReturnNFT(&_Main.TransactOpts, _nftContract, _tokenId)
}

// ReturnNFT is a paid mutator transaction binding the contract method 0xccb568ec.
//
// Solidity: function returnNFT(address _nftContract, uint256 _tokenId) returns()
func (_Main *MainTransactorSession) ReturnNFT(_nftContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Main.Contract.ReturnNFT(&_Main.TransactOpts, _nftContract, _tokenId)
}

// SetFakeUSDTContract is a paid mutator transaction binding the contract method 0x6f52cd6e.
//
// Solidity: function setFakeUSDTContract(address _fakeUSDTAddress) returns()
func (_Main *MainTransactor) SetFakeUSDTContract(opts *bind.TransactOpts, _fakeUSDTAddress common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setFakeUSDTContract", _fakeUSDTAddress)
}

// SetFakeUSDTContract is a paid mutator transaction binding the contract method 0x6f52cd6e.
//
// Solidity: function setFakeUSDTContract(address _fakeUSDTAddress) returns()
func (_Main *MainSession) SetFakeUSDTContract(_fakeUSDTAddress common.Address) (*types.Transaction, error) {
	return _Main.Contract.SetFakeUSDTContract(&_Main.TransactOpts, _fakeUSDTAddress)
}

// SetFakeUSDTContract is a paid mutator transaction binding the contract method 0x6f52cd6e.
//
// Solidity: function setFakeUSDTContract(address _fakeUSDTAddress) returns()
func (_Main *MainTransactorSession) SetFakeUSDTContract(_fakeUSDTAddress common.Address) (*types.Transaction, error) {
	return _Main.Contract.SetFakeUSDTContract(&_Main.TransactOpts, _fakeUSDTAddress)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns()
func (_Main *MainTransactor) SetFee(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setFee", _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns()
func (_Main *MainSession) SetFee(_fee *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetFee(&_Main.TransactOpts, _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns()
func (_Main *MainTransactorSession) SetFee(_fee *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetFee(&_Main.TransactOpts, _fee)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Main *MainTransactor) WithdrawAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "withdrawAll")
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Main *MainSession) WithdrawAll() (*types.Transaction, error) {
	return _Main.Contract.WithdrawAll(&_Main.TransactOpts)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Main *MainTransactorSession) WithdrawAll() (*types.Transaction, error) {
	return _Main.Contract.WithdrawAll(&_Main.TransactOpts)
}

// MainNFTAddedIterator is returned from FilterNFTAdded and is used to iterate over the raw logs and unpacked data for NFTAdded events raised by the Main contract.
type MainNFTAddedIterator struct {
	Event *MainNFTAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainNFTAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainNFTAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainNFTAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainNFTAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainNFTAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainNFTAdded represents a NFTAdded event raised by the Main contract.
type MainNFTAdded struct {
	Owner      common.Address
	NFTAddress common.Address
	TokenId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNFTAdded is a free log retrieval operation binding the contract event 0x75ee001158df8b77347acda2c33d52e5d6facd0c4331fd0910a4b5eb3993369a.
//
// Solidity: event NFTAdded(address indexed owner, address indexed NFTAddress, uint256 tokenId)
func (_Main *MainFilterer) FilterNFTAdded(opts *bind.FilterOpts, owner []common.Address, NFTAddress []common.Address) (*MainNFTAddedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var NFTAddressRule []interface{}
	for _, NFTAddressItem := range NFTAddress {
		NFTAddressRule = append(NFTAddressRule, NFTAddressItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "NFTAdded", ownerRule, NFTAddressRule)
	if err != nil {
		return nil, err
	}
	return &MainNFTAddedIterator{contract: _Main.contract, event: "NFTAdded", logs: logs, sub: sub}, nil
}

// WatchNFTAdded is a free log subscription operation binding the contract event 0x75ee001158df8b77347acda2c33d52e5d6facd0c4331fd0910a4b5eb3993369a.
//
// Solidity: event NFTAdded(address indexed owner, address indexed NFTAddress, uint256 tokenId)
func (_Main *MainFilterer) WatchNFTAdded(opts *bind.WatchOpts, sink chan<- *MainNFTAdded, owner []common.Address, NFTAddress []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var NFTAddressRule []interface{}
	for _, NFTAddressItem := range NFTAddress {
		NFTAddressRule = append(NFTAddressRule, NFTAddressItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "NFTAdded", ownerRule, NFTAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainNFTAdded)
				if err := _Main.contract.UnpackLog(event, "NFTAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNFTAdded is a log parse operation binding the contract event 0x75ee001158df8b77347acda2c33d52e5d6facd0c4331fd0910a4b5eb3993369a.
//
// Solidity: event NFTAdded(address indexed owner, address indexed NFTAddress, uint256 tokenId)
func (_Main *MainFilterer) ParseNFTAdded(log types.Log) (*MainNFTAdded, error) {
	event := new(MainNFTAdded)
	if err := _Main.contract.UnpackLog(event, "NFTAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainNFTBorrowedIterator is returned from FilterNFTBorrowed and is used to iterate over the raw logs and unpacked data for NFTBorrowed events raised by the Main contract.
type MainNFTBorrowedIterator struct {
	Event *MainNFTBorrowed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainNFTBorrowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainNFTBorrowed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainNFTBorrowed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainNFTBorrowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainNFTBorrowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainNFTBorrowed represents a NFTBorrowed event raised by the Main contract.
type MainNFTBorrowed struct {
	Borrower   common.Address
	Lender     common.Address
	NFTAddress common.Address
	TokenId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNFTBorrowed is a free log retrieval operation binding the contract event 0x018587ee2904cbc9583f54aa70a102990f241760f13f7a24ae0de20693487d7b.
//
// Solidity: event NFTBorrowed(address indexed borrower, address indexed lender, address indexed NFTAddress, uint256 tokenId)
func (_Main *MainFilterer) FilterNFTBorrowed(opts *bind.FilterOpts, borrower []common.Address, lender []common.Address, NFTAddress []common.Address) (*MainNFTBorrowedIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}
	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}
	var NFTAddressRule []interface{}
	for _, NFTAddressItem := range NFTAddress {
		NFTAddressRule = append(NFTAddressRule, NFTAddressItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "NFTBorrowed", borrowerRule, lenderRule, NFTAddressRule)
	if err != nil {
		return nil, err
	}
	return &MainNFTBorrowedIterator{contract: _Main.contract, event: "NFTBorrowed", logs: logs, sub: sub}, nil
}

// WatchNFTBorrowed is a free log subscription operation binding the contract event 0x018587ee2904cbc9583f54aa70a102990f241760f13f7a24ae0de20693487d7b.
//
// Solidity: event NFTBorrowed(address indexed borrower, address indexed lender, address indexed NFTAddress, uint256 tokenId)
func (_Main *MainFilterer) WatchNFTBorrowed(opts *bind.WatchOpts, sink chan<- *MainNFTBorrowed, borrower []common.Address, lender []common.Address, NFTAddress []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}
	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}
	var NFTAddressRule []interface{}
	for _, NFTAddressItem := range NFTAddress {
		NFTAddressRule = append(NFTAddressRule, NFTAddressItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "NFTBorrowed", borrowerRule, lenderRule, NFTAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainNFTBorrowed)
				if err := _Main.contract.UnpackLog(event, "NFTBorrowed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNFTBorrowed is a log parse operation binding the contract event 0x018587ee2904cbc9583f54aa70a102990f241760f13f7a24ae0de20693487d7b.
//
// Solidity: event NFTBorrowed(address indexed borrower, address indexed lender, address indexed NFTAddress, uint256 tokenId)
func (_Main *MainFilterer) ParseNFTBorrowed(log types.Log) (*MainNFTBorrowed, error) {
	event := new(MainNFTBorrowed)
	if err := _Main.contract.UnpackLog(event, "NFTBorrowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainNFTCanceledIterator is returned from FilterNFTCanceled and is used to iterate over the raw logs and unpacked data for NFTCanceled events raised by the Main contract.
type MainNFTCanceledIterator struct {
	Event *MainNFTCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainNFTCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainNFTCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainNFTCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainNFTCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainNFTCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainNFTCanceled represents a NFTCanceled event raised by the Main contract.
type MainNFTCanceled struct {
	Owner      common.Address
	NFTAddress common.Address
	TokenId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNFTCanceled is a free log retrieval operation binding the contract event 0x1fc64ec2285c9891a5f6513865cacc14559dbec84a1dd36395483a27fd06324d.
//
// Solidity: event NFTCanceled(address indexed owner, address indexed NFTAddress, uint256 tokenId)
func (_Main *MainFilterer) FilterNFTCanceled(opts *bind.FilterOpts, owner []common.Address, NFTAddress []common.Address) (*MainNFTCanceledIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var NFTAddressRule []interface{}
	for _, NFTAddressItem := range NFTAddress {
		NFTAddressRule = append(NFTAddressRule, NFTAddressItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "NFTCanceled", ownerRule, NFTAddressRule)
	if err != nil {
		return nil, err
	}
	return &MainNFTCanceledIterator{contract: _Main.contract, event: "NFTCanceled", logs: logs, sub: sub}, nil
}

// WatchNFTCanceled is a free log subscription operation binding the contract event 0x1fc64ec2285c9891a5f6513865cacc14559dbec84a1dd36395483a27fd06324d.
//
// Solidity: event NFTCanceled(address indexed owner, address indexed NFTAddress, uint256 tokenId)
func (_Main *MainFilterer) WatchNFTCanceled(opts *bind.WatchOpts, sink chan<- *MainNFTCanceled, owner []common.Address, NFTAddress []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var NFTAddressRule []interface{}
	for _, NFTAddressItem := range NFTAddress {
		NFTAddressRule = append(NFTAddressRule, NFTAddressItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "NFTCanceled", ownerRule, NFTAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainNFTCanceled)
				if err := _Main.contract.UnpackLog(event, "NFTCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNFTCanceled is a log parse operation binding the contract event 0x1fc64ec2285c9891a5f6513865cacc14559dbec84a1dd36395483a27fd06324d.
//
// Solidity: event NFTCanceled(address indexed owner, address indexed NFTAddress, uint256 tokenId)
func (_Main *MainFilterer) ParseNFTCanceled(log types.Log) (*MainNFTCanceled, error) {
	event := new(MainNFTCanceled)
	if err := _Main.contract.UnpackLog(event, "NFTCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainNFTReceivedIterator is returned from FilterNFTReceived and is used to iterate over the raw logs and unpacked data for NFTReceived events raised by the Main contract.
type MainNFTReceivedIterator struct {
	Event *MainNFTReceived // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainNFTReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainNFTReceived)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainNFTReceived)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainNFTReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainNFTReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainNFTReceived represents a NFTReceived event raised by the Main contract.
type MainNFTReceived struct {
	From    common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNFTReceived is a free log retrieval operation binding the contract event 0x126facc6a4d0c86875713c3734ddb43944c8b837ec483016d6fa4b921e0ac90b.
//
// Solidity: event NFTReceived(address from, uint256 tokenId)
func (_Main *MainFilterer) FilterNFTReceived(opts *bind.FilterOpts) (*MainNFTReceivedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "NFTReceived")
	if err != nil {
		return nil, err
	}
	return &MainNFTReceivedIterator{contract: _Main.contract, event: "NFTReceived", logs: logs, sub: sub}, nil
}

// WatchNFTReceived is a free log subscription operation binding the contract event 0x126facc6a4d0c86875713c3734ddb43944c8b837ec483016d6fa4b921e0ac90b.
//
// Solidity: event NFTReceived(address from, uint256 tokenId)
func (_Main *MainFilterer) WatchNFTReceived(opts *bind.WatchOpts, sink chan<- *MainNFTReceived) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "NFTReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainNFTReceived)
				if err := _Main.contract.UnpackLog(event, "NFTReceived", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNFTReceived is a log parse operation binding the contract event 0x126facc6a4d0c86875713c3734ddb43944c8b837ec483016d6fa4b921e0ac90b.
//
// Solidity: event NFTReceived(address from, uint256 tokenId)
func (_Main *MainFilterer) ParseNFTReceived(log types.Log) (*MainNFTReceived, error) {
	event := new(MainNFTReceived)
	if err := _Main.contract.UnpackLog(event, "NFTReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainNFTReturnedIterator is returned from FilterNFTReturned and is used to iterate over the raw logs and unpacked data for NFTReturned events raised by the Main contract.
type MainNFTReturnedIterator struct {
	Event *MainNFTReturned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainNFTReturnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainNFTReturned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainNFTReturned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainNFTReturnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainNFTReturnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainNFTReturned represents a NFTReturned event raised by the Main contract.
type MainNFTReturned struct {
	Borrower   common.Address
	Lender     common.Address
	NFTAddress common.Address
	TokenId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNFTReturned is a free log retrieval operation binding the contract event 0xc0bd3c824ba6fcb0d28b5548f84a231d7252efb9252c44196bf4e4ee7323ef33.
//
// Solidity: event NFTReturned(address indexed borrower, address indexed lender, address indexed NFTAddress, uint256 tokenId)
func (_Main *MainFilterer) FilterNFTReturned(opts *bind.FilterOpts, borrower []common.Address, lender []common.Address, NFTAddress []common.Address) (*MainNFTReturnedIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}
	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}
	var NFTAddressRule []interface{}
	for _, NFTAddressItem := range NFTAddress {
		NFTAddressRule = append(NFTAddressRule, NFTAddressItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "NFTReturned", borrowerRule, lenderRule, NFTAddressRule)
	if err != nil {
		return nil, err
	}
	return &MainNFTReturnedIterator{contract: _Main.contract, event: "NFTReturned", logs: logs, sub: sub}, nil
}

// WatchNFTReturned is a free log subscription operation binding the contract event 0xc0bd3c824ba6fcb0d28b5548f84a231d7252efb9252c44196bf4e4ee7323ef33.
//
// Solidity: event NFTReturned(address indexed borrower, address indexed lender, address indexed NFTAddress, uint256 tokenId)
func (_Main *MainFilterer) WatchNFTReturned(opts *bind.WatchOpts, sink chan<- *MainNFTReturned, borrower []common.Address, lender []common.Address, NFTAddress []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}
	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}
	var NFTAddressRule []interface{}
	for _, NFTAddressItem := range NFTAddress {
		NFTAddressRule = append(NFTAddressRule, NFTAddressItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "NFTReturned", borrowerRule, lenderRule, NFTAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainNFTReturned)
				if err := _Main.contract.UnpackLog(event, "NFTReturned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNFTReturned is a log parse operation binding the contract event 0xc0bd3c824ba6fcb0d28b5548f84a231d7252efb9252c44196bf4e4ee7323ef33.
//
// Solidity: event NFTReturned(address indexed borrower, address indexed lender, address indexed NFTAddress, uint256 tokenId)
func (_Main *MainFilterer) ParseNFTReturned(log types.Log) (*MainNFTReturned, error) {
	event := new(MainNFTReturned)
	if err := _Main.contract.UnpackLog(event, "NFTReturned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainNFTWithdrawnIterator is returned from FilterNFTWithdrawn and is used to iterate over the raw logs and unpacked data for NFTWithdrawn events raised by the Main contract.
type MainNFTWithdrawnIterator struct {
	Event *MainNFTWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MainNFTWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainNFTWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MainNFTWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MainNFTWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainNFTWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainNFTWithdrawn represents a NFTWithdrawn event raised by the Main contract.
type MainNFTWithdrawn struct {
	Owner      common.Address
	NFTAddress common.Address
	TokenId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNFTWithdrawn is a free log retrieval operation binding the contract event 0xbbde41973f9ce4890f7ad9762c23d8191f261fd643bdf13ed8bbc10549b49fcb.
//
// Solidity: event NFTWithdrawn(address indexed owner, address indexed NFTAddress, uint256 tokenId)
func (_Main *MainFilterer) FilterNFTWithdrawn(opts *bind.FilterOpts, owner []common.Address, NFTAddress []common.Address) (*MainNFTWithdrawnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var NFTAddressRule []interface{}
	for _, NFTAddressItem := range NFTAddress {
		NFTAddressRule = append(NFTAddressRule, NFTAddressItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "NFTWithdrawn", ownerRule, NFTAddressRule)
	if err != nil {
		return nil, err
	}
	return &MainNFTWithdrawnIterator{contract: _Main.contract, event: "NFTWithdrawn", logs: logs, sub: sub}, nil
}

// WatchNFTWithdrawn is a free log subscription operation binding the contract event 0xbbde41973f9ce4890f7ad9762c23d8191f261fd643bdf13ed8bbc10549b49fcb.
//
// Solidity: event NFTWithdrawn(address indexed owner, address indexed NFTAddress, uint256 tokenId)
func (_Main *MainFilterer) WatchNFTWithdrawn(opts *bind.WatchOpts, sink chan<- *MainNFTWithdrawn, owner []common.Address, NFTAddress []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var NFTAddressRule []interface{}
	for _, NFTAddressItem := range NFTAddress {
		NFTAddressRule = append(NFTAddressRule, NFTAddressItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "NFTWithdrawn", ownerRule, NFTAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainNFTWithdrawn)
				if err := _Main.contract.UnpackLog(event, "NFTWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNFTWithdrawn is a log parse operation binding the contract event 0xbbde41973f9ce4890f7ad9762c23d8191f261fd643bdf13ed8bbc10549b49fcb.
//
// Solidity: event NFTWithdrawn(address indexed owner, address indexed NFTAddress, uint256 tokenId)
func (_Main *MainFilterer) ParseNFTWithdrawn(log types.Log) (*MainNFTWithdrawn, error) {
	event := new(MainNFTWithdrawn)
	if err := _Main.contract.UnpackLog(event, "NFTWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
