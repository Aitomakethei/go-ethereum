// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethofs

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
)

// EthoFSControllerHostingObject is an auto generated low-level Go binding around an user-defined struct.
type EthoFSControllerHostingObject struct {
	HostingReplicationFactor uint8
	DeployedDaysLength       uint16
	ContractStorageUsed      uint32
	CreationTimestamp        uint32
	MainContentHash          string
	HostingContractName      string
	OwnerAddress             common.Address
}

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"contract_id\",\"type\":\"uint32\"}],\"name\":\"OnChangedOwnership\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"contract_id\",\"type\":\"uint32\"}],\"name\":\"OnContractExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"contract_id\",\"type\":\"uint32\"}],\"name\":\"OnCreateContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"contract_id\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"extendedDaysLength\",\"type\":\"uint32\"}],\"name\":\"OnExtendContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"OnFundCollection\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newHost\",\"type\":\"address\"}],\"name\":\"OnNewContractHost\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"contract_id\",\"type\":\"uint32\"}],\"name\":\"OnRemoveContract\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ExpirationDict\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"ExpirationDict_Indexer\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userID\",\"type\":\"address\"}],\"name\":\"GetActiveUserContracts\",\"outputs\":[{\"internalType\":\"uint32[]\",\"name\":\"\",\"type\":\"uint32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"bound_a\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"bound_b\",\"type\":\"uint32\"}],\"name\":\"GetAllHostingContracts\",\"outputs\":[{\"internalType\":\"uint32[]\",\"name\":\"\",\"type\":\"uint32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"fileID\",\"type\":\"uint32\"}],\"name\":\"GetContract\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"hostingReplicationFactor\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"deployedDaysLength\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"contractStorageUsed\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"creationTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"mainContentHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hostingContractName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"}],\"internalType\":\"structEthoFSController.HostingObject\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"fileID\",\"type\":\"uint32\"}],\"name\":\"GetContractDurationLeft\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"timeleft\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"ContractSize\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"ContractDuration\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"ReplicationFactor\",\"type\":\"uint8\"}],\"name\":\"GetHostingCost\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"total_cost\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userID\",\"type\":\"address\"}],\"name\":\"GetUserID\",\"outputs\":[{\"internalType\":\"uint160\",\"name\":\"userID\",\"type\":\"uint160\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userID\",\"type\":\"address\"}],\"name\":\"GetUserStats\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"numHostingObjects\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"totalHostingSize\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MaximumDeletionCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MaximumStorageTime\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MinimumReplicationFactor\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"set\",\"type\":\"address\"}],\"name\":\"SetAccountCollectionAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"set\",\"type\":\"uint128\"}],\"name\":\"SetHostingContractCost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"set\",\"type\":\"uint8\"}],\"name\":\"SetMaximumDeletionCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"set\",\"type\":\"uint16\"}],\"name\":\"SetMaximumStorageTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"set\",\"type\":\"uint8\"}],\"name\":\"SetMinimumReplicationFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newoperator\",\"type\":\"address\"}],\"name\":\"SetOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"CID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ContractName\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"ReplicationFactor\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"ContractSize\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"ContractDuration\",\"type\":\"uint16\"}],\"name\":\"addContract\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"fileID\",\"type\":\"uint32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"CID_LIST\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"NAME_LIST\",\"type\":\"string[]\"},{\"internalType\":\"uint8[]\",\"name\":\"REPLICATION_LIST\",\"type\":\"uint8[]\"},{\"internalType\":\"uint32[]\",\"name\":\"SIZE_LIST\",\"type\":\"uint32[]\"},{\"internalType\":\"uint16[]\",\"name\":\"DURATION_LIST\",\"type\":\"uint16[]\"}],\"name\":\"addContracts\",\"outputs\":[{\"internalType\":\"uint32[]\",\"name\":\"\",\"type\":\"uint32[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allHostingObjectsIndex\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"fileID\",\"type\":\"uint32\"}],\"name\":\"changeContractOwnership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collection_address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractsTotalCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractsTotalSize\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controllerCreationTime\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"day_cycle\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"fileID\",\"type\":\"uint32\"}],\"name\":\"deleteContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"fileID\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"extendDaysLength\",\"type\":\"uint16\"}],\"name\":\"extendContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fake_day_cycle\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"fileID_to_index_allActiveHostingIndices\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"force_collect\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hostingCost_perMB_perYEAR\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint160\",\"name\":\"\",\"type\":\"uint160\"}],\"name\":\"users\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"numHostingObjects\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"totalHostingSize\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"maxDelete\",\"type\":\"uint16\"}],\"name\":\"wipeContracts\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"del_count\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StoreABI is the input ABI used to generate the binding from.
// Deprecated: Use StoreMetaData.ABI instead.
var StoreABI = StoreMetaData.ABI

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// ExpirationDict is a free data retrieval call binding the contract method 0x06b0b641.
//
// Solidity: function ExpirationDict(uint16 , uint256 ) view returns(uint32)
func (_Store *StoreCaller) ExpirationDict(opts *bind.CallOpts, arg0 uint16, arg1 *big.Int) (uint32, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "ExpirationDict", arg0, arg1)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ExpirationDict is a free data retrieval call binding the contract method 0x06b0b641.
//
// Solidity: function ExpirationDict(uint16 , uint256 ) view returns(uint32)
func (_Store *StoreSession) ExpirationDict(arg0 uint16, arg1 *big.Int) (uint32, error) {
	return _Store.Contract.ExpirationDict(&_Store.CallOpts, arg0, arg1)
}

// ExpirationDict is a free data retrieval call binding the contract method 0x06b0b641.
//
// Solidity: function ExpirationDict(uint16 , uint256 ) view returns(uint32)
func (_Store *StoreCallerSession) ExpirationDict(arg0 uint16, arg1 *big.Int) (uint32, error) {
	return _Store.Contract.ExpirationDict(&_Store.CallOpts, arg0, arg1)
}

// ExpirationDictIndexer is a free data retrieval call binding the contract method 0x751f74e2.
//
// Solidity: function ExpirationDict_Indexer(uint16 ) view returns(uint32)
func (_Store *StoreCaller) ExpirationDictIndexer(opts *bind.CallOpts, arg0 uint16) (uint32, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "ExpirationDict_Indexer", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ExpirationDictIndexer is a free data retrieval call binding the contract method 0x751f74e2.
//
// Solidity: function ExpirationDict_Indexer(uint16 ) view returns(uint32)
func (_Store *StoreSession) ExpirationDictIndexer(arg0 uint16) (uint32, error) {
	return _Store.Contract.ExpirationDictIndexer(&_Store.CallOpts, arg0)
}

// ExpirationDictIndexer is a free data retrieval call binding the contract method 0x751f74e2.
//
// Solidity: function ExpirationDict_Indexer(uint16 ) view returns(uint32)
func (_Store *StoreCallerSession) ExpirationDictIndexer(arg0 uint16) (uint32, error) {
	return _Store.Contract.ExpirationDictIndexer(&_Store.CallOpts, arg0)
}

// GetActiveUserContracts is a free data retrieval call binding the contract method 0x76b0424c.
//
// Solidity: function GetActiveUserContracts(address userID) view returns(uint32[])
func (_Store *StoreCaller) GetActiveUserContracts(opts *bind.CallOpts, userID common.Address) ([]uint32, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetActiveUserContracts", userID)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetActiveUserContracts is a free data retrieval call binding the contract method 0x76b0424c.
//
// Solidity: function GetActiveUserContracts(address userID) view returns(uint32[])
func (_Store *StoreSession) GetActiveUserContracts(userID common.Address) ([]uint32, error) {
	return _Store.Contract.GetActiveUserContracts(&_Store.CallOpts, userID)
}

// GetActiveUserContracts is a free data retrieval call binding the contract method 0x76b0424c.
//
// Solidity: function GetActiveUserContracts(address userID) view returns(uint32[])
func (_Store *StoreCallerSession) GetActiveUserContracts(userID common.Address) ([]uint32, error) {
	return _Store.Contract.GetActiveUserContracts(&_Store.CallOpts, userID)
}

// GetAllHostingContracts is a free data retrieval call binding the contract method 0xd62486ee.
//
// Solidity: function GetAllHostingContracts(uint32 bound_a, uint32 bound_b) view returns(uint32[])
func (_Store *StoreCaller) GetAllHostingContracts(opts *bind.CallOpts, bound_a uint32, bound_b uint32) ([]uint32, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetAllHostingContracts", bound_a, bound_b)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetAllHostingContracts is a free data retrieval call binding the contract method 0xd62486ee.
//
// Solidity: function GetAllHostingContracts(uint32 bound_a, uint32 bound_b) view returns(uint32[])
func (_Store *StoreSession) GetAllHostingContracts(bound_a uint32, bound_b uint32) ([]uint32, error) {
	return _Store.Contract.GetAllHostingContracts(&_Store.CallOpts, bound_a, bound_b)
}

// GetAllHostingContracts is a free data retrieval call binding the contract method 0xd62486ee.
//
// Solidity: function GetAllHostingContracts(uint32 bound_a, uint32 bound_b) view returns(uint32[])
func (_Store *StoreCallerSession) GetAllHostingContracts(bound_a uint32, bound_b uint32) ([]uint32, error) {
	return _Store.Contract.GetAllHostingContracts(&_Store.CallOpts, bound_a, bound_b)
}

// GetContract is a free data retrieval call binding the contract method 0x089bb2d0.
//
// Solidity: function GetContract(uint32 fileID) view returns((uint8,uint16,uint32,uint32,string,string,address))
func (_Store *StoreCaller) GetContract(opts *bind.CallOpts, fileID uint32) (EthoFSControllerHostingObject, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetContract", fileID)

	if err != nil {
		return *new(EthoFSControllerHostingObject), err
	}

	out0 := *abi.ConvertType(out[0], new(EthoFSControllerHostingObject)).(*EthoFSControllerHostingObject)

	return out0, err

}

// GetContract is a free data retrieval call binding the contract method 0x089bb2d0.
//
// Solidity: function GetContract(uint32 fileID) view returns((uint8,uint16,uint32,uint32,string,string,address))
func (_Store *StoreSession) GetContract(fileID uint32) (EthoFSControllerHostingObject, error) {
	return _Store.Contract.GetContract(&_Store.CallOpts, fileID)
}

// GetContract is a free data retrieval call binding the contract method 0x089bb2d0.
//
// Solidity: function GetContract(uint32 fileID) view returns((uint8,uint16,uint32,uint32,string,string,address))
func (_Store *StoreCallerSession) GetContract(fileID uint32) (EthoFSControllerHostingObject, error) {
	return _Store.Contract.GetContract(&_Store.CallOpts, fileID)
}

// GetContractDurationLeft is a free data retrieval call binding the contract method 0xbd7c8cfd.
//
// Solidity: function GetContractDurationLeft(uint32 fileID) view returns(uint32 timeleft)
func (_Store *StoreCaller) GetContractDurationLeft(opts *bind.CallOpts, fileID uint32) (uint32, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetContractDurationLeft", fileID)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetContractDurationLeft is a free data retrieval call binding the contract method 0xbd7c8cfd.
//
// Solidity: function GetContractDurationLeft(uint32 fileID) view returns(uint32 timeleft)
func (_Store *StoreSession) GetContractDurationLeft(fileID uint32) (uint32, error) {
	return _Store.Contract.GetContractDurationLeft(&_Store.CallOpts, fileID)
}

// GetContractDurationLeft is a free data retrieval call binding the contract method 0xbd7c8cfd.
//
// Solidity: function GetContractDurationLeft(uint32 fileID) view returns(uint32 timeleft)
func (_Store *StoreCallerSession) GetContractDurationLeft(fileID uint32) (uint32, error) {
	return _Store.Contract.GetContractDurationLeft(&_Store.CallOpts, fileID)
}

// GetHostingCost is a free data retrieval call binding the contract method 0x1f56990f.
//
// Solidity: function GetHostingCost(uint32 ContractSize, uint16 ContractDuration, uint8 ReplicationFactor) view returns(uint128 total_cost)
func (_Store *StoreCaller) GetHostingCost(opts *bind.CallOpts, ContractSize uint32, ContractDuration uint16, ReplicationFactor uint8) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetHostingCost", ContractSize, ContractDuration, ReplicationFactor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetHostingCost is a free data retrieval call binding the contract method 0x1f56990f.
//
// Solidity: function GetHostingCost(uint32 ContractSize, uint16 ContractDuration, uint8 ReplicationFactor) view returns(uint128 total_cost)
func (_Store *StoreSession) GetHostingCost(ContractSize uint32, ContractDuration uint16, ReplicationFactor uint8) (*big.Int, error) {
	return _Store.Contract.GetHostingCost(&_Store.CallOpts, ContractSize, ContractDuration, ReplicationFactor)
}

// GetHostingCost is a free data retrieval call binding the contract method 0x1f56990f.
//
// Solidity: function GetHostingCost(uint32 ContractSize, uint16 ContractDuration, uint8 ReplicationFactor) view returns(uint128 total_cost)
func (_Store *StoreCallerSession) GetHostingCost(ContractSize uint32, ContractDuration uint16, ReplicationFactor uint8) (*big.Int, error) {
	return _Store.Contract.GetHostingCost(&_Store.CallOpts, ContractSize, ContractDuration, ReplicationFactor)
}

// GetUserID is a free data retrieval call binding the contract method 0xd7c0da4a.
//
// Solidity: function GetUserID(address _userID) pure returns(uint160 userID)
func (_Store *StoreCaller) GetUserID(opts *bind.CallOpts, _userID common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetUserID", _userID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserID is a free data retrieval call binding the contract method 0xd7c0da4a.
//
// Solidity: function GetUserID(address _userID) pure returns(uint160 userID)
func (_Store *StoreSession) GetUserID(_userID common.Address) (*big.Int, error) {
	return _Store.Contract.GetUserID(&_Store.CallOpts, _userID)
}

// GetUserID is a free data retrieval call binding the contract method 0xd7c0da4a.
//
// Solidity: function GetUserID(address _userID) pure returns(uint160 userID)
func (_Store *StoreCallerSession) GetUserID(_userID common.Address) (*big.Int, error) {
	return _Store.Contract.GetUserID(&_Store.CallOpts, _userID)
}

// GetUserStats is a free data retrieval call binding the contract method 0xb0370f26.
//
// Solidity: function GetUserStats(address userID) view returns(uint32 numHostingObjects, uint64 totalHostingSize)
func (_Store *StoreCaller) GetUserStats(opts *bind.CallOpts, userID common.Address) (struct {
	NumHostingObjects uint32
	TotalHostingSize  uint64
}, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetUserStats", userID)

	outstruct := new(struct {
		NumHostingObjects uint32
		TotalHostingSize  uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NumHostingObjects = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.TotalHostingSize = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetUserStats is a free data retrieval call binding the contract method 0xb0370f26.
//
// Solidity: function GetUserStats(address userID) view returns(uint32 numHostingObjects, uint64 totalHostingSize)
func (_Store *StoreSession) GetUserStats(userID common.Address) (struct {
	NumHostingObjects uint32
	TotalHostingSize  uint64
}, error) {
	return _Store.Contract.GetUserStats(&_Store.CallOpts, userID)
}

// GetUserStats is a free data retrieval call binding the contract method 0xb0370f26.
//
// Solidity: function GetUserStats(address userID) view returns(uint32 numHostingObjects, uint64 totalHostingSize)
func (_Store *StoreCallerSession) GetUserStats(userID common.Address) (struct {
	NumHostingObjects uint32
	TotalHostingSize  uint64
}, error) {
	return _Store.Contract.GetUserStats(&_Store.CallOpts, userID)
}

// MaximumDeletionCount is a free data retrieval call binding the contract method 0xc5fbca1f.
//
// Solidity: function MaximumDeletionCount() view returns(uint8)
func (_Store *StoreCaller) MaximumDeletionCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "MaximumDeletionCount")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MaximumDeletionCount is a free data retrieval call binding the contract method 0xc5fbca1f.
//
// Solidity: function MaximumDeletionCount() view returns(uint8)
func (_Store *StoreSession) MaximumDeletionCount() (uint8, error) {
	return _Store.Contract.MaximumDeletionCount(&_Store.CallOpts)
}

// MaximumDeletionCount is a free data retrieval call binding the contract method 0xc5fbca1f.
//
// Solidity: function MaximumDeletionCount() view returns(uint8)
func (_Store *StoreCallerSession) MaximumDeletionCount() (uint8, error) {
	return _Store.Contract.MaximumDeletionCount(&_Store.CallOpts)
}

// MaximumStorageTime is a free data retrieval call binding the contract method 0x6280fd0b.
//
// Solidity: function MaximumStorageTime() view returns(uint16)
func (_Store *StoreCaller) MaximumStorageTime(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "MaximumStorageTime")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MaximumStorageTime is a free data retrieval call binding the contract method 0x6280fd0b.
//
// Solidity: function MaximumStorageTime() view returns(uint16)
func (_Store *StoreSession) MaximumStorageTime() (uint16, error) {
	return _Store.Contract.MaximumStorageTime(&_Store.CallOpts)
}

// MaximumStorageTime is a free data retrieval call binding the contract method 0x6280fd0b.
//
// Solidity: function MaximumStorageTime() view returns(uint16)
func (_Store *StoreCallerSession) MaximumStorageTime() (uint16, error) {
	return _Store.Contract.MaximumStorageTime(&_Store.CallOpts)
}

// MinimumReplicationFactor is a free data retrieval call binding the contract method 0x6f0c6e49.
//
// Solidity: function MinimumReplicationFactor() view returns(uint8)
func (_Store *StoreCaller) MinimumReplicationFactor(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "MinimumReplicationFactor")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MinimumReplicationFactor is a free data retrieval call binding the contract method 0x6f0c6e49.
//
// Solidity: function MinimumReplicationFactor() view returns(uint8)
func (_Store *StoreSession) MinimumReplicationFactor() (uint8, error) {
	return _Store.Contract.MinimumReplicationFactor(&_Store.CallOpts)
}

// MinimumReplicationFactor is a free data retrieval call binding the contract method 0x6f0c6e49.
//
// Solidity: function MinimumReplicationFactor() view returns(uint8)
func (_Store *StoreCallerSession) MinimumReplicationFactor() (uint8, error) {
	return _Store.Contract.MinimumReplicationFactor(&_Store.CallOpts)
}

// AllHostingObjectsIndex is a free data retrieval call binding the contract method 0x0d26a2b8.
//
// Solidity: function allHostingObjectsIndex() view returns(uint32)
func (_Store *StoreCaller) AllHostingObjectsIndex(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "allHostingObjectsIndex")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// AllHostingObjectsIndex is a free data retrieval call binding the contract method 0x0d26a2b8.
//
// Solidity: function allHostingObjectsIndex() view returns(uint32)
func (_Store *StoreSession) AllHostingObjectsIndex() (uint32, error) {
	return _Store.Contract.AllHostingObjectsIndex(&_Store.CallOpts)
}

// AllHostingObjectsIndex is a free data retrieval call binding the contract method 0x0d26a2b8.
//
// Solidity: function allHostingObjectsIndex() view returns(uint32)
func (_Store *StoreCallerSession) AllHostingObjectsIndex() (uint32, error) {
	return _Store.Contract.AllHostingObjectsIndex(&_Store.CallOpts)
}

// CollectionAddress is a free data retrieval call binding the contract method 0xb2438978.
//
// Solidity: function collection_address() view returns(address)
func (_Store *StoreCaller) CollectionAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "collection_address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollectionAddress is a free data retrieval call binding the contract method 0xb2438978.
//
// Solidity: function collection_address() view returns(address)
func (_Store *StoreSession) CollectionAddress() (common.Address, error) {
	return _Store.Contract.CollectionAddress(&_Store.CallOpts)
}

// CollectionAddress is a free data retrieval call binding the contract method 0xb2438978.
//
// Solidity: function collection_address() view returns(address)
func (_Store *StoreCallerSession) CollectionAddress() (common.Address, error) {
	return _Store.Contract.CollectionAddress(&_Store.CallOpts)
}

// ContractsTotalCount is a free data retrieval call binding the contract method 0xd2df07d8.
//
// Solidity: function contractsTotalCount() view returns(uint32)
func (_Store *StoreCaller) ContractsTotalCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "contractsTotalCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ContractsTotalCount is a free data retrieval call binding the contract method 0xd2df07d8.
//
// Solidity: function contractsTotalCount() view returns(uint32)
func (_Store *StoreSession) ContractsTotalCount() (uint32, error) {
	return _Store.Contract.ContractsTotalCount(&_Store.CallOpts)
}

// ContractsTotalCount is a free data retrieval call binding the contract method 0xd2df07d8.
//
// Solidity: function contractsTotalCount() view returns(uint32)
func (_Store *StoreCallerSession) ContractsTotalCount() (uint32, error) {
	return _Store.Contract.ContractsTotalCount(&_Store.CallOpts)
}

// ContractsTotalSize is a free data retrieval call binding the contract method 0xc7f8140f.
//
// Solidity: function contractsTotalSize() view returns(uint64)
func (_Store *StoreCaller) ContractsTotalSize(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "contractsTotalSize")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ContractsTotalSize is a free data retrieval call binding the contract method 0xc7f8140f.
//
// Solidity: function contractsTotalSize() view returns(uint64)
func (_Store *StoreSession) ContractsTotalSize() (uint64, error) {
	return _Store.Contract.ContractsTotalSize(&_Store.CallOpts)
}

// ContractsTotalSize is a free data retrieval call binding the contract method 0xc7f8140f.
//
// Solidity: function contractsTotalSize() view returns(uint64)
func (_Store *StoreCallerSession) ContractsTotalSize() (uint64, error) {
	return _Store.Contract.ContractsTotalSize(&_Store.CallOpts)
}

// ControllerCreationTime is a free data retrieval call binding the contract method 0x1008219e.
//
// Solidity: function controllerCreationTime() view returns(uint32)
func (_Store *StoreCaller) ControllerCreationTime(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "controllerCreationTime")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ControllerCreationTime is a free data retrieval call binding the contract method 0x1008219e.
//
// Solidity: function controllerCreationTime() view returns(uint32)
func (_Store *StoreSession) ControllerCreationTime() (uint32, error) {
	return _Store.Contract.ControllerCreationTime(&_Store.CallOpts)
}

// ControllerCreationTime is a free data retrieval call binding the contract method 0x1008219e.
//
// Solidity: function controllerCreationTime() view returns(uint32)
func (_Store *StoreCallerSession) ControllerCreationTime() (uint32, error) {
	return _Store.Contract.ControllerCreationTime(&_Store.CallOpts)
}

// DayCycle is a free data retrieval call binding the contract method 0x8483e801.
//
// Solidity: function day_cycle() view returns(uint32)
func (_Store *StoreCaller) DayCycle(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "day_cycle")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// DayCycle is a free data retrieval call binding the contract method 0x8483e801.
//
// Solidity: function day_cycle() view returns(uint32)
func (_Store *StoreSession) DayCycle() (uint32, error) {
	return _Store.Contract.DayCycle(&_Store.CallOpts)
}

// DayCycle is a free data retrieval call binding the contract method 0x8483e801.
//
// Solidity: function day_cycle() view returns(uint32)
func (_Store *StoreCallerSession) DayCycle() (uint32, error) {
	return _Store.Contract.DayCycle(&_Store.CallOpts)
}

// FakeDayCycle is a free data retrieval call binding the contract method 0x0898f9ba.
//
// Solidity: function fake_day_cycle() view returns(uint32)
func (_Store *StoreCaller) FakeDayCycle(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "fake_day_cycle")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// FakeDayCycle is a free data retrieval call binding the contract method 0x0898f9ba.
//
// Solidity: function fake_day_cycle() view returns(uint32)
func (_Store *StoreSession) FakeDayCycle() (uint32, error) {
	return _Store.Contract.FakeDayCycle(&_Store.CallOpts)
}

// FakeDayCycle is a free data retrieval call binding the contract method 0x0898f9ba.
//
// Solidity: function fake_day_cycle() view returns(uint32)
func (_Store *StoreCallerSession) FakeDayCycle() (uint32, error) {
	return _Store.Contract.FakeDayCycle(&_Store.CallOpts)
}

// FileIDToIndexAllActiveHostingIndices is a free data retrieval call binding the contract method 0xcb89a170.
//
// Solidity: function fileID_to_index_allActiveHostingIndices(uint32 ) view returns(uint32)
func (_Store *StoreCaller) FileIDToIndexAllActiveHostingIndices(opts *bind.CallOpts, arg0 uint32) (uint32, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "fileID_to_index_allActiveHostingIndices", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// FileIDToIndexAllActiveHostingIndices is a free data retrieval call binding the contract method 0xcb89a170.
//
// Solidity: function fileID_to_index_allActiveHostingIndices(uint32 ) view returns(uint32)
func (_Store *StoreSession) FileIDToIndexAllActiveHostingIndices(arg0 uint32) (uint32, error) {
	return _Store.Contract.FileIDToIndexAllActiveHostingIndices(&_Store.CallOpts, arg0)
}

// FileIDToIndexAllActiveHostingIndices is a free data retrieval call binding the contract method 0xcb89a170.
//
// Solidity: function fileID_to_index_allActiveHostingIndices(uint32 ) view returns(uint32)
func (_Store *StoreCallerSession) FileIDToIndexAllActiveHostingIndices(arg0 uint32) (uint32, error) {
	return _Store.Contract.FileIDToIndexAllActiveHostingIndices(&_Store.CallOpts, arg0)
}

// HostingCostPerMBPerYEAR is a free data retrieval call binding the contract method 0x83094857.
//
// Solidity: function hostingCost_perMB_perYEAR() view returns(uint128)
func (_Store *StoreCaller) HostingCostPerMBPerYEAR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "hostingCost_perMB_perYEAR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HostingCostPerMBPerYEAR is a free data retrieval call binding the contract method 0x83094857.
//
// Solidity: function hostingCost_perMB_perYEAR() view returns(uint128)
func (_Store *StoreSession) HostingCostPerMBPerYEAR() (*big.Int, error) {
	return _Store.Contract.HostingCostPerMBPerYEAR(&_Store.CallOpts)
}

// HostingCostPerMBPerYEAR is a free data retrieval call binding the contract method 0x83094857.
//
// Solidity: function hostingCost_perMB_perYEAR() view returns(uint128)
func (_Store *StoreCallerSession) HostingCostPerMBPerYEAR() (*big.Int, error) {
	return _Store.Contract.HostingCostPerMBPerYEAR(&_Store.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Store *StoreCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Store *StoreSession) Operator() (common.Address, error) {
	return _Store.Contract.Operator(&_Store.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Store *StoreCallerSession) Operator() (common.Address, error) {
	return _Store.Contract.Operator(&_Store.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Store *StoreCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Store *StoreSession) Owner() (common.Address, error) {
	return _Store.Contract.Owner(&_Store.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Store *StoreCallerSession) Owner() (common.Address, error) {
	return _Store.Contract.Owner(&_Store.CallOpts)
}

// Users is a free data retrieval call binding the contract method 0xc2e71e9c.
//
// Solidity: function users(uint160 ) view returns(uint32 numHostingObjects, uint64 totalHostingSize)
func (_Store *StoreCaller) Users(opts *bind.CallOpts, arg0 *big.Int) (struct {
	NumHostingObjects uint32
	TotalHostingSize  uint64
}, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "users", arg0)

	outstruct := new(struct {
		NumHostingObjects uint32
		TotalHostingSize  uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NumHostingObjects = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.TotalHostingSize = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// Users is a free data retrieval call binding the contract method 0xc2e71e9c.
//
// Solidity: function users(uint160 ) view returns(uint32 numHostingObjects, uint64 totalHostingSize)
func (_Store *StoreSession) Users(arg0 *big.Int) (struct {
	NumHostingObjects uint32
	TotalHostingSize  uint64
}, error) {
	return _Store.Contract.Users(&_Store.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0xc2e71e9c.
//
// Solidity: function users(uint160 ) view returns(uint32 numHostingObjects, uint64 totalHostingSize)
func (_Store *StoreCallerSession) Users(arg0 *big.Int) (struct {
	NumHostingObjects uint32
	TotalHostingSize  uint64
}, error) {
	return _Store.Contract.Users(&_Store.CallOpts, arg0)
}

// SetAccountCollectionAddress is a paid mutator transaction binding the contract method 0xc8ac5fe8.
//
// Solidity: function SetAccountCollectionAddress(address set) returns()
func (_Store *StoreTransactor) SetAccountCollectionAddress(opts *bind.TransactOpts, set common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "SetAccountCollectionAddress", set)
}

// SetAccountCollectionAddress is a paid mutator transaction binding the contract method 0xc8ac5fe8.
//
// Solidity: function SetAccountCollectionAddress(address set) returns()
func (_Store *StoreSession) SetAccountCollectionAddress(set common.Address) (*types.Transaction, error) {
	return _Store.Contract.SetAccountCollectionAddress(&_Store.TransactOpts, set)
}

// SetAccountCollectionAddress is a paid mutator transaction binding the contract method 0xc8ac5fe8.
//
// Solidity: function SetAccountCollectionAddress(address set) returns()
func (_Store *StoreTransactorSession) SetAccountCollectionAddress(set common.Address) (*types.Transaction, error) {
	return _Store.Contract.SetAccountCollectionAddress(&_Store.TransactOpts, set)
}

// SetHostingContractCost is a paid mutator transaction binding the contract method 0x2029eb37.
//
// Solidity: function SetHostingContractCost(uint128 set) returns()
func (_Store *StoreTransactor) SetHostingContractCost(opts *bind.TransactOpts, set *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "SetHostingContractCost", set)
}

// SetHostingContractCost is a paid mutator transaction binding the contract method 0x2029eb37.
//
// Solidity: function SetHostingContractCost(uint128 set) returns()
func (_Store *StoreSession) SetHostingContractCost(set *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SetHostingContractCost(&_Store.TransactOpts, set)
}

// SetHostingContractCost is a paid mutator transaction binding the contract method 0x2029eb37.
//
// Solidity: function SetHostingContractCost(uint128 set) returns()
func (_Store *StoreTransactorSession) SetHostingContractCost(set *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SetHostingContractCost(&_Store.TransactOpts, set)
}

// SetMaximumDeletionCount is a paid mutator transaction binding the contract method 0xa882272e.
//
// Solidity: function SetMaximumDeletionCount(uint8 set) returns()
func (_Store *StoreTransactor) SetMaximumDeletionCount(opts *bind.TransactOpts, set uint8) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "SetMaximumDeletionCount", set)
}

// SetMaximumDeletionCount is a paid mutator transaction binding the contract method 0xa882272e.
//
// Solidity: function SetMaximumDeletionCount(uint8 set) returns()
func (_Store *StoreSession) SetMaximumDeletionCount(set uint8) (*types.Transaction, error) {
	return _Store.Contract.SetMaximumDeletionCount(&_Store.TransactOpts, set)
}

// SetMaximumDeletionCount is a paid mutator transaction binding the contract method 0xa882272e.
//
// Solidity: function SetMaximumDeletionCount(uint8 set) returns()
func (_Store *StoreTransactorSession) SetMaximumDeletionCount(set uint8) (*types.Transaction, error) {
	return _Store.Contract.SetMaximumDeletionCount(&_Store.TransactOpts, set)
}

// SetMaximumStorageTime is a paid mutator transaction binding the contract method 0x63604819.
//
// Solidity: function SetMaximumStorageTime(uint16 set) returns()
func (_Store *StoreTransactor) SetMaximumStorageTime(opts *bind.TransactOpts, set uint16) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "SetMaximumStorageTime", set)
}

// SetMaximumStorageTime is a paid mutator transaction binding the contract method 0x63604819.
//
// Solidity: function SetMaximumStorageTime(uint16 set) returns()
func (_Store *StoreSession) SetMaximumStorageTime(set uint16) (*types.Transaction, error) {
	return _Store.Contract.SetMaximumStorageTime(&_Store.TransactOpts, set)
}

// SetMaximumStorageTime is a paid mutator transaction binding the contract method 0x63604819.
//
// Solidity: function SetMaximumStorageTime(uint16 set) returns()
func (_Store *StoreTransactorSession) SetMaximumStorageTime(set uint16) (*types.Transaction, error) {
	return _Store.Contract.SetMaximumStorageTime(&_Store.TransactOpts, set)
}

// SetMinimumReplicationFactor is a paid mutator transaction binding the contract method 0x96f8e237.
//
// Solidity: function SetMinimumReplicationFactor(uint8 set) returns()
func (_Store *StoreTransactor) SetMinimumReplicationFactor(opts *bind.TransactOpts, set uint8) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "SetMinimumReplicationFactor", set)
}

// SetMinimumReplicationFactor is a paid mutator transaction binding the contract method 0x96f8e237.
//
// Solidity: function SetMinimumReplicationFactor(uint8 set) returns()
func (_Store *StoreSession) SetMinimumReplicationFactor(set uint8) (*types.Transaction, error) {
	return _Store.Contract.SetMinimumReplicationFactor(&_Store.TransactOpts, set)
}

// SetMinimumReplicationFactor is a paid mutator transaction binding the contract method 0x96f8e237.
//
// Solidity: function SetMinimumReplicationFactor(uint8 set) returns()
func (_Store *StoreTransactorSession) SetMinimumReplicationFactor(set uint8) (*types.Transaction, error) {
	return _Store.Contract.SetMinimumReplicationFactor(&_Store.TransactOpts, set)
}

// SetOperator is a paid mutator transaction binding the contract method 0xdbebfba6.
//
// Solidity: function SetOperator(address newoperator) returns()
func (_Store *StoreTransactor) SetOperator(opts *bind.TransactOpts, newoperator common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "SetOperator", newoperator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xdbebfba6.
//
// Solidity: function SetOperator(address newoperator) returns()
func (_Store *StoreSession) SetOperator(newoperator common.Address) (*types.Transaction, error) {
	return _Store.Contract.SetOperator(&_Store.TransactOpts, newoperator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xdbebfba6.
//
// Solidity: function SetOperator(address newoperator) returns()
func (_Store *StoreTransactorSession) SetOperator(newoperator common.Address) (*types.Transaction, error) {
	return _Store.Contract.SetOperator(&_Store.TransactOpts, newoperator)
}

// AddContract is a paid mutator transaction binding the contract method 0xba851e45.
//
// Solidity: function addContract(string CID, string ContractName, uint8 ReplicationFactor, uint32 ContractSize, uint16 ContractDuration) payable returns(uint32 fileID)
func (_Store *StoreTransactor) AddContract(opts *bind.TransactOpts, CID string, ContractName string, ReplicationFactor uint8, ContractSize uint32, ContractDuration uint16) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "addContract", CID, ContractName, ReplicationFactor, ContractSize, ContractDuration)
}

// AddContract is a paid mutator transaction binding the contract method 0xba851e45.
//
// Solidity: function addContract(string CID, string ContractName, uint8 ReplicationFactor, uint32 ContractSize, uint16 ContractDuration) payable returns(uint32 fileID)
func (_Store *StoreSession) AddContract(CID string, ContractName string, ReplicationFactor uint8, ContractSize uint32, ContractDuration uint16) (*types.Transaction, error) {
	return _Store.Contract.AddContract(&_Store.TransactOpts, CID, ContractName, ReplicationFactor, ContractSize, ContractDuration)
}

// AddContract is a paid mutator transaction binding the contract method 0xba851e45.
//
// Solidity: function addContract(string CID, string ContractName, uint8 ReplicationFactor, uint32 ContractSize, uint16 ContractDuration) payable returns(uint32 fileID)
func (_Store *StoreTransactorSession) AddContract(CID string, ContractName string, ReplicationFactor uint8, ContractSize uint32, ContractDuration uint16) (*types.Transaction, error) {
	return _Store.Contract.AddContract(&_Store.TransactOpts, CID, ContractName, ReplicationFactor, ContractSize, ContractDuration)
}

// AddContracts is a paid mutator transaction binding the contract method 0xd9e11dda.
//
// Solidity: function addContracts(string[] CID_LIST, string[] NAME_LIST, uint8[] REPLICATION_LIST, uint32[] SIZE_LIST, uint16[] DURATION_LIST) payable returns(uint32[])
func (_Store *StoreTransactor) AddContracts(opts *bind.TransactOpts, CID_LIST []string, NAME_LIST []string, REPLICATION_LIST []uint8, SIZE_LIST []uint32, DURATION_LIST []uint16) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "addContracts", CID_LIST, NAME_LIST, REPLICATION_LIST, SIZE_LIST, DURATION_LIST)
}

// AddContracts is a paid mutator transaction binding the contract method 0xd9e11dda.
//
// Solidity: function addContracts(string[] CID_LIST, string[] NAME_LIST, uint8[] REPLICATION_LIST, uint32[] SIZE_LIST, uint16[] DURATION_LIST) payable returns(uint32[])
func (_Store *StoreSession) AddContracts(CID_LIST []string, NAME_LIST []string, REPLICATION_LIST []uint8, SIZE_LIST []uint32, DURATION_LIST []uint16) (*types.Transaction, error) {
	return _Store.Contract.AddContracts(&_Store.TransactOpts, CID_LIST, NAME_LIST, REPLICATION_LIST, SIZE_LIST, DURATION_LIST)
}

// AddContracts is a paid mutator transaction binding the contract method 0xd9e11dda.
//
// Solidity: function addContracts(string[] CID_LIST, string[] NAME_LIST, uint8[] REPLICATION_LIST, uint32[] SIZE_LIST, uint16[] DURATION_LIST) payable returns(uint32[])
func (_Store *StoreTransactorSession) AddContracts(CID_LIST []string, NAME_LIST []string, REPLICATION_LIST []uint8, SIZE_LIST []uint32, DURATION_LIST []uint16) (*types.Transaction, error) {
	return _Store.Contract.AddContracts(&_Store.TransactOpts, CID_LIST, NAME_LIST, REPLICATION_LIST, SIZE_LIST, DURATION_LIST)
}

// ChangeContractOwnership is a paid mutator transaction binding the contract method 0xd10237e4.
//
// Solidity: function changeContractOwnership(address newOwner, uint32 fileID) returns(bool success)
func (_Store *StoreTransactor) ChangeContractOwnership(opts *bind.TransactOpts, newOwner common.Address, fileID uint32) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "changeContractOwnership", newOwner, fileID)
}

// ChangeContractOwnership is a paid mutator transaction binding the contract method 0xd10237e4.
//
// Solidity: function changeContractOwnership(address newOwner, uint32 fileID) returns(bool success)
func (_Store *StoreSession) ChangeContractOwnership(newOwner common.Address, fileID uint32) (*types.Transaction, error) {
	return _Store.Contract.ChangeContractOwnership(&_Store.TransactOpts, newOwner, fileID)
}

// ChangeContractOwnership is a paid mutator transaction binding the contract method 0xd10237e4.
//
// Solidity: function changeContractOwnership(address newOwner, uint32 fileID) returns(bool success)
func (_Store *StoreTransactorSession) ChangeContractOwnership(newOwner common.Address, fileID uint32) (*types.Transaction, error) {
	return _Store.Contract.ChangeContractOwnership(&_Store.TransactOpts, newOwner, fileID)
}

// DeleteContract is a paid mutator transaction binding the contract method 0x9fb25ce8.
//
// Solidity: function deleteContract(uint32 fileID) returns(bool success)
func (_Store *StoreTransactor) DeleteContract(opts *bind.TransactOpts, fileID uint32) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "deleteContract", fileID)
}

// DeleteContract is a paid mutator transaction binding the contract method 0x9fb25ce8.
//
// Solidity: function deleteContract(uint32 fileID) returns(bool success)
func (_Store *StoreSession) DeleteContract(fileID uint32) (*types.Transaction, error) {
	return _Store.Contract.DeleteContract(&_Store.TransactOpts, fileID)
}

// DeleteContract is a paid mutator transaction binding the contract method 0x9fb25ce8.
//
// Solidity: function deleteContract(uint32 fileID) returns(bool success)
func (_Store *StoreTransactorSession) DeleteContract(fileID uint32) (*types.Transaction, error) {
	return _Store.Contract.DeleteContract(&_Store.TransactOpts, fileID)
}

// ExtendContract is a paid mutator transaction binding the contract method 0xd070620c.
//
// Solidity: function extendContract(uint32 fileID, uint16 extendDaysLength) payable returns(bool success)
func (_Store *StoreTransactor) ExtendContract(opts *bind.TransactOpts, fileID uint32, extendDaysLength uint16) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "extendContract", fileID, extendDaysLength)
}

// ExtendContract is a paid mutator transaction binding the contract method 0xd070620c.
//
// Solidity: function extendContract(uint32 fileID, uint16 extendDaysLength) payable returns(bool success)
func (_Store *StoreSession) ExtendContract(fileID uint32, extendDaysLength uint16) (*types.Transaction, error) {
	return _Store.Contract.ExtendContract(&_Store.TransactOpts, fileID, extendDaysLength)
}

// ExtendContract is a paid mutator transaction binding the contract method 0xd070620c.
//
// Solidity: function extendContract(uint32 fileID, uint16 extendDaysLength) payable returns(bool success)
func (_Store *StoreTransactorSession) ExtendContract(fileID uint32, extendDaysLength uint16) (*types.Transaction, error) {
	return _Store.Contract.ExtendContract(&_Store.TransactOpts, fileID, extendDaysLength)
}

// ForceCollect is a paid mutator transaction binding the contract method 0xa751a4ff.
//
// Solidity: function force_collect() returns(bool success)
func (_Store *StoreTransactor) ForceCollect(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "force_collect")
}

// ForceCollect is a paid mutator transaction binding the contract method 0xa751a4ff.
//
// Solidity: function force_collect() returns(bool success)
func (_Store *StoreSession) ForceCollect() (*types.Transaction, error) {
	return _Store.Contract.ForceCollect(&_Store.TransactOpts)
}

// ForceCollect is a paid mutator transaction binding the contract method 0xa751a4ff.
//
// Solidity: function force_collect() returns(bool success)
func (_Store *StoreTransactorSession) ForceCollect() (*types.Transaction, error) {
	return _Store.Contract.ForceCollect(&_Store.TransactOpts)
}

// WipeContracts is a paid mutator transaction binding the contract method 0xa65bb905.
//
// Solidity: function wipeContracts(uint16 maxDelete) returns(uint32 del_count)
func (_Store *StoreTransactor) WipeContracts(opts *bind.TransactOpts, maxDelete uint16) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "wipeContracts", maxDelete)
}

// WipeContracts is a paid mutator transaction binding the contract method 0xa65bb905.
//
// Solidity: function wipeContracts(uint16 maxDelete) returns(uint32 del_count)
func (_Store *StoreSession) WipeContracts(maxDelete uint16) (*types.Transaction, error) {
	return _Store.Contract.WipeContracts(&_Store.TransactOpts, maxDelete)
}

// WipeContracts is a paid mutator transaction binding the contract method 0xa65bb905.
//
// Solidity: function wipeContracts(uint16 maxDelete) returns(uint32 del_count)
func (_Store *StoreTransactorSession) WipeContracts(maxDelete uint16) (*types.Transaction, error) {
	return _Store.Contract.WipeContracts(&_Store.TransactOpts, maxDelete)
}

// StoreOnChangedOwnershipIterator is returned from FilterOnChangedOwnership and is used to iterate over the raw logs and unpacked data for OnChangedOwnership events raised by the Store contract.
type StoreOnChangedOwnershipIterator struct {
	Event *StoreOnChangedOwnership // Event containing the contract specifics and raw log

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
func (it *StoreOnChangedOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreOnChangedOwnership)
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
		it.Event = new(StoreOnChangedOwnership)
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
func (it *StoreOnChangedOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreOnChangedOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreOnChangedOwnership represents a OnChangedOwnership event raised by the Store contract.
type StoreOnChangedOwnership struct {
	ContractId uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOnChangedOwnership is a free log retrieval operation binding the contract event 0xbd1a15a90e202ff87d1b2a018c573dc50c60c6378aeac4cf5b7723b09983b9be.
//
// Solidity: event OnChangedOwnership(uint32 contract_id)
func (_Store *StoreFilterer) FilterOnChangedOwnership(opts *bind.FilterOpts) (*StoreOnChangedOwnershipIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "OnChangedOwnership")
	if err != nil {
		return nil, err
	}
	return &StoreOnChangedOwnershipIterator{contract: _Store.contract, event: "OnChangedOwnership", logs: logs, sub: sub}, nil
}

// WatchOnChangedOwnership is a free log subscription operation binding the contract event 0xbd1a15a90e202ff87d1b2a018c573dc50c60c6378aeac4cf5b7723b09983b9be.
//
// Solidity: event OnChangedOwnership(uint32 contract_id)
func (_Store *StoreFilterer) WatchOnChangedOwnership(opts *bind.WatchOpts, sink chan<- *StoreOnChangedOwnership) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "OnChangedOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreOnChangedOwnership)
				if err := _Store.contract.UnpackLog(event, "OnChangedOwnership", log); err != nil {
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

// ParseOnChangedOwnership is a log parse operation binding the contract event 0xbd1a15a90e202ff87d1b2a018c573dc50c60c6378aeac4cf5b7723b09983b9be.
//
// Solidity: event OnChangedOwnership(uint32 contract_id)
func (_Store *StoreFilterer) ParseOnChangedOwnership(log types.Log) (*StoreOnChangedOwnership, error) {
	event := new(StoreOnChangedOwnership)
	if err := _Store.contract.UnpackLog(event, "OnChangedOwnership", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreOnContractExpiredIterator is returned from FilterOnContractExpired and is used to iterate over the raw logs and unpacked data for OnContractExpired events raised by the Store contract.
type StoreOnContractExpiredIterator struct {
	Event *StoreOnContractExpired // Event containing the contract specifics and raw log

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
func (it *StoreOnContractExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreOnContractExpired)
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
		it.Event = new(StoreOnContractExpired)
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
func (it *StoreOnContractExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreOnContractExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreOnContractExpired represents a OnContractExpired event raised by the Store contract.
type StoreOnContractExpired struct {
	ContractId uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOnContractExpired is a free log retrieval operation binding the contract event 0x5396ce02249cbdff77bb92d5df88038b01c5fa1daa4099192457757686da15cc.
//
// Solidity: event OnContractExpired(uint32 contract_id)
func (_Store *StoreFilterer) FilterOnContractExpired(opts *bind.FilterOpts) (*StoreOnContractExpiredIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "OnContractExpired")
	if err != nil {
		return nil, err
	}
	return &StoreOnContractExpiredIterator{contract: _Store.contract, event: "OnContractExpired", logs: logs, sub: sub}, nil
}

// WatchOnContractExpired is a free log subscription operation binding the contract event 0x5396ce02249cbdff77bb92d5df88038b01c5fa1daa4099192457757686da15cc.
//
// Solidity: event OnContractExpired(uint32 contract_id)
func (_Store *StoreFilterer) WatchOnContractExpired(opts *bind.WatchOpts, sink chan<- *StoreOnContractExpired) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "OnContractExpired")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreOnContractExpired)
				if err := _Store.contract.UnpackLog(event, "OnContractExpired", log); err != nil {
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

// ParseOnContractExpired is a log parse operation binding the contract event 0x5396ce02249cbdff77bb92d5df88038b01c5fa1daa4099192457757686da15cc.
//
// Solidity: event OnContractExpired(uint32 contract_id)
func (_Store *StoreFilterer) ParseOnContractExpired(log types.Log) (*StoreOnContractExpired, error) {
	event := new(StoreOnContractExpired)
	if err := _Store.contract.UnpackLog(event, "OnContractExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreOnCreateContractIterator is returned from FilterOnCreateContract and is used to iterate over the raw logs and unpacked data for OnCreateContract events raised by the Store contract.
type StoreOnCreateContractIterator struct {
	Event *StoreOnCreateContract // Event containing the contract specifics and raw log

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
func (it *StoreOnCreateContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreOnCreateContract)
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
		it.Event = new(StoreOnCreateContract)
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
func (it *StoreOnCreateContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreOnCreateContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreOnCreateContract represents a OnCreateContract event raised by the Store contract.
type StoreOnCreateContract struct {
	ContractId uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOnCreateContract is a free log retrieval operation binding the contract event 0xd50673a2cf6c23c1da46f94532518d1bb1f72ab9c74f799772cb39a28543f626.
//
// Solidity: event OnCreateContract(uint32 contract_id)
func (_Store *StoreFilterer) FilterOnCreateContract(opts *bind.FilterOpts) (*StoreOnCreateContractIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "OnCreateContract")
	if err != nil {
		return nil, err
	}
	return &StoreOnCreateContractIterator{contract: _Store.contract, event: "OnCreateContract", logs: logs, sub: sub}, nil
}

// WatchOnCreateContract is a free log subscription operation binding the contract event 0xd50673a2cf6c23c1da46f94532518d1bb1f72ab9c74f799772cb39a28543f626.
//
// Solidity: event OnCreateContract(uint32 contract_id)
func (_Store *StoreFilterer) WatchOnCreateContract(opts *bind.WatchOpts, sink chan<- *StoreOnCreateContract) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "OnCreateContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreOnCreateContract)
				if err := _Store.contract.UnpackLog(event, "OnCreateContract", log); err != nil {
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

// ParseOnCreateContract is a log parse operation binding the contract event 0xd50673a2cf6c23c1da46f94532518d1bb1f72ab9c74f799772cb39a28543f626.
//
// Solidity: event OnCreateContract(uint32 contract_id)
func (_Store *StoreFilterer) ParseOnCreateContract(log types.Log) (*StoreOnCreateContract, error) {
	event := new(StoreOnCreateContract)
	if err := _Store.contract.UnpackLog(event, "OnCreateContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreOnExtendContractIterator is returned from FilterOnExtendContract and is used to iterate over the raw logs and unpacked data for OnExtendContract events raised by the Store contract.
type StoreOnExtendContractIterator struct {
	Event *StoreOnExtendContract // Event containing the contract specifics and raw log

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
func (it *StoreOnExtendContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreOnExtendContract)
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
		it.Event = new(StoreOnExtendContract)
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
func (it *StoreOnExtendContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreOnExtendContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreOnExtendContract represents a OnExtendContract event raised by the Store contract.
type StoreOnExtendContract struct {
	ContractId         uint32
	ExtendedDaysLength uint32
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOnExtendContract is a free log retrieval operation binding the contract event 0xf5f6c7fc4deecd877f3e356b9bf27737fdcd4e512ba00c24af77601ea58e4a2b.
//
// Solidity: event OnExtendContract(uint32 contract_id, uint32 extendedDaysLength)
func (_Store *StoreFilterer) FilterOnExtendContract(opts *bind.FilterOpts) (*StoreOnExtendContractIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "OnExtendContract")
	if err != nil {
		return nil, err
	}
	return &StoreOnExtendContractIterator{contract: _Store.contract, event: "OnExtendContract", logs: logs, sub: sub}, nil
}

// WatchOnExtendContract is a free log subscription operation binding the contract event 0xf5f6c7fc4deecd877f3e356b9bf27737fdcd4e512ba00c24af77601ea58e4a2b.
//
// Solidity: event OnExtendContract(uint32 contract_id, uint32 extendedDaysLength)
func (_Store *StoreFilterer) WatchOnExtendContract(opts *bind.WatchOpts, sink chan<- *StoreOnExtendContract) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "OnExtendContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreOnExtendContract)
				if err := _Store.contract.UnpackLog(event, "OnExtendContract", log); err != nil {
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

// ParseOnExtendContract is a log parse operation binding the contract event 0xf5f6c7fc4deecd877f3e356b9bf27737fdcd4e512ba00c24af77601ea58e4a2b.
//
// Solidity: event OnExtendContract(uint32 contract_id, uint32 extendedDaysLength)
func (_Store *StoreFilterer) ParseOnExtendContract(log types.Log) (*StoreOnExtendContract, error) {
	event := new(StoreOnExtendContract)
	if err := _Store.contract.UnpackLog(event, "OnExtendContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreOnFundCollectionIterator is returned from FilterOnFundCollection and is used to iterate over the raw logs and unpacked data for OnFundCollection events raised by the Store contract.
type StoreOnFundCollectionIterator struct {
	Event *StoreOnFundCollection // Event containing the contract specifics and raw log

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
func (it *StoreOnFundCollectionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreOnFundCollection)
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
		it.Event = new(StoreOnFundCollection)
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
func (it *StoreOnFundCollectionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreOnFundCollectionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreOnFundCollection represents a OnFundCollection event raised by the Store contract.
type StoreOnFundCollection struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOnFundCollection is a free log retrieval operation binding the contract event 0xa1a7c0f7be600e1955c313722acf3b7ff5906e7c283225182e4cda5c4d1c1d0f.
//
// Solidity: event OnFundCollection(address indexed sender, uint256 amount)
func (_Store *StoreFilterer) FilterOnFundCollection(opts *bind.FilterOpts, sender []common.Address) (*StoreOnFundCollectionIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "OnFundCollection", senderRule)
	if err != nil {
		return nil, err
	}
	return &StoreOnFundCollectionIterator{contract: _Store.contract, event: "OnFundCollection", logs: logs, sub: sub}, nil
}

// WatchOnFundCollection is a free log subscription operation binding the contract event 0xa1a7c0f7be600e1955c313722acf3b7ff5906e7c283225182e4cda5c4d1c1d0f.
//
// Solidity: event OnFundCollection(address indexed sender, uint256 amount)
func (_Store *StoreFilterer) WatchOnFundCollection(opts *bind.WatchOpts, sink chan<- *StoreOnFundCollection, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "OnFundCollection", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreOnFundCollection)
				if err := _Store.contract.UnpackLog(event, "OnFundCollection", log); err != nil {
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

// ParseOnFundCollection is a log parse operation binding the contract event 0xa1a7c0f7be600e1955c313722acf3b7ff5906e7c283225182e4cda5c4d1c1d0f.
//
// Solidity: event OnFundCollection(address indexed sender, uint256 amount)
func (_Store *StoreFilterer) ParseOnFundCollection(log types.Log) (*StoreOnFundCollection, error) {
	event := new(StoreOnFundCollection)
	if err := _Store.contract.UnpackLog(event, "OnFundCollection", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreOnNewContractHostIterator is returned from FilterOnNewContractHost and is used to iterate over the raw logs and unpacked data for OnNewContractHost events raised by the Store contract.
type StoreOnNewContractHostIterator struct {
	Event *StoreOnNewContractHost // Event containing the contract specifics and raw log

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
func (it *StoreOnNewContractHostIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreOnNewContractHost)
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
		it.Event = new(StoreOnNewContractHost)
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
func (it *StoreOnNewContractHostIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreOnNewContractHostIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreOnNewContractHost represents a OnNewContractHost event raised by the Store contract.
type StoreOnNewContractHost struct {
	NewHost common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOnNewContractHost is a free log retrieval operation binding the contract event 0x1be6c1d17de7f202460ec62b35c0d1359ca723cd988099134bce274dad1620f9.
//
// Solidity: event OnNewContractHost(address newHost)
func (_Store *StoreFilterer) FilterOnNewContractHost(opts *bind.FilterOpts) (*StoreOnNewContractHostIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "OnNewContractHost")
	if err != nil {
		return nil, err
	}
	return &StoreOnNewContractHostIterator{contract: _Store.contract, event: "OnNewContractHost", logs: logs, sub: sub}, nil
}

// WatchOnNewContractHost is a free log subscription operation binding the contract event 0x1be6c1d17de7f202460ec62b35c0d1359ca723cd988099134bce274dad1620f9.
//
// Solidity: event OnNewContractHost(address newHost)
func (_Store *StoreFilterer) WatchOnNewContractHost(opts *bind.WatchOpts, sink chan<- *StoreOnNewContractHost) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "OnNewContractHost")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreOnNewContractHost)
				if err := _Store.contract.UnpackLog(event, "OnNewContractHost", log); err != nil {
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

// ParseOnNewContractHost is a log parse operation binding the contract event 0x1be6c1d17de7f202460ec62b35c0d1359ca723cd988099134bce274dad1620f9.
//
// Solidity: event OnNewContractHost(address newHost)
func (_Store *StoreFilterer) ParseOnNewContractHost(log types.Log) (*StoreOnNewContractHost, error) {
	event := new(StoreOnNewContractHost)
	if err := _Store.contract.UnpackLog(event, "OnNewContractHost", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreOnRemoveContractIterator is returned from FilterOnRemoveContract and is used to iterate over the raw logs and unpacked data for OnRemoveContract events raised by the Store contract.
type StoreOnRemoveContractIterator struct {
	Event *StoreOnRemoveContract // Event containing the contract specifics and raw log

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
func (it *StoreOnRemoveContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreOnRemoveContract)
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
		it.Event = new(StoreOnRemoveContract)
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
func (it *StoreOnRemoveContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreOnRemoveContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreOnRemoveContract represents a OnRemoveContract event raised by the Store contract.
type StoreOnRemoveContract struct {
	ContractId uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOnRemoveContract is a free log retrieval operation binding the contract event 0xf022cab3a0f887ef921b0185e2311568a6b5fd57da766da5445435cf0cee110a.
//
// Solidity: event OnRemoveContract(uint32 contract_id)
func (_Store *StoreFilterer) FilterOnRemoveContract(opts *bind.FilterOpts) (*StoreOnRemoveContractIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "OnRemoveContract")
	if err != nil {
		return nil, err
	}
	return &StoreOnRemoveContractIterator{contract: _Store.contract, event: "OnRemoveContract", logs: logs, sub: sub}, nil
}

// WatchOnRemoveContract is a free log subscription operation binding the contract event 0xf022cab3a0f887ef921b0185e2311568a6b5fd57da766da5445435cf0cee110a.
//
// Solidity: event OnRemoveContract(uint32 contract_id)
func (_Store *StoreFilterer) WatchOnRemoveContract(opts *bind.WatchOpts, sink chan<- *StoreOnRemoveContract) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "OnRemoveContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreOnRemoveContract)
				if err := _Store.contract.UnpackLog(event, "OnRemoveContract", log); err != nil {
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

// ParseOnRemoveContract is a log parse operation binding the contract event 0xf022cab3a0f887ef921b0185e2311568a6b5fd57da766da5445435cf0cee110a.
//
// Solidity: event OnRemoveContract(uint32 contract_id)
func (_Store *StoreFilterer) ParseOnRemoveContract(log types.Log) (*StoreOnRemoveContract, error) {
	event := new(StoreOnRemoveContract)
	if err := _Store.contract.UnpackLog(event, "OnRemoveContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
