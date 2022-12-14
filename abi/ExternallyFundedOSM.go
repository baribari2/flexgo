// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// ExternallyFundedMetaData contains all meta data concerning the ExternallyFunded contract.
var ExternallyFundedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"priceSource_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddAuthorization\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"delay\",\"type\":\"uint16\"}],\"name\":\"ChangeDelay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"priceSource\",\"type\":\"address\"}],\"name\":\"ChangePriceSource\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wrapper\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"FailRenumerateCaller\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"parameter\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"ModifyParameters\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"parameter\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"ModifyParameters\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RemoveAuthorization\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"RestartValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Start\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Stop\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMedian\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastUpdateTime\",\"type\":\"uint256\"}],\"name\":\"UpdateResult\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addAuthorization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorizedAccounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"delay\",\"type\":\"uint16\"}],\"name\":\"changeDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"priceSource_\",\"type\":\"address\"}],\"name\":\"changePriceSource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fsmWrapper\",\"outputs\":[{\"internalType\":\"contractFSMWrapperLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextResultWithValidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getResultWithValidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"parameter\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"modifyParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"passedDelay\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"ok\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceSource\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"read\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeAuthorization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"restartValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateDelay\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ExternallyFundedABI is the input ABI used to generate the binding from.
// Deprecated: Use ExternallyFundedMetaData.ABI instead.
var ExternallyFundedABI = ExternallyFundedMetaData.ABI

// ExternallyFunded is an auto generated Go binding around an Ethereum contract.
type ExternallyFunded struct {
	ExternallyFundedCaller     // Read-only binding to the contract
	ExternallyFundedTransactor // Write-only binding to the contract
	ExternallyFundedFilterer   // Log filterer for contract events
}

// ExternallyFundedCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExternallyFundedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExternallyFundedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExternallyFundedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExternallyFundedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExternallyFundedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExternallyFundedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExternallyFundedSession struct {
	Contract     *ExternallyFunded // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExternallyFundedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExternallyFundedCallerSession struct {
	Contract *ExternallyFundedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ExternallyFundedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExternallyFundedTransactorSession struct {
	Contract     *ExternallyFundedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ExternallyFundedRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExternallyFundedRaw struct {
	Contract *ExternallyFunded // Generic contract binding to access the raw methods on
}

// ExternallyFundedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExternallyFundedCallerRaw struct {
	Contract *ExternallyFundedCaller // Generic read-only contract binding to access the raw methods on
}

// ExternallyFundedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExternallyFundedTransactorRaw struct {
	Contract *ExternallyFundedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExternallyFunded creates a new instance of ExternallyFunded, bound to a specific deployed contract.
func NewExternallyFunded(address common.Address, backend bind.ContractBackend) (*ExternallyFunded, error) {
	contract, err := bindExternallyFunded(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExternallyFunded{ExternallyFundedCaller: ExternallyFundedCaller{contract: contract}, ExternallyFundedTransactor: ExternallyFundedTransactor{contract: contract}, ExternallyFundedFilterer: ExternallyFundedFilterer{contract: contract}}, nil
}

// NewExternallyFundedCaller creates a new read-only instance of ExternallyFunded, bound to a specific deployed contract.
func NewExternallyFundedCaller(address common.Address, caller bind.ContractCaller) (*ExternallyFundedCaller, error) {
	contract, err := bindExternallyFunded(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExternallyFundedCaller{contract: contract}, nil
}

// NewExternallyFundedTransactor creates a new write-only instance of ExternallyFunded, bound to a specific deployed contract.
func NewExternallyFundedTransactor(address common.Address, transactor bind.ContractTransactor) (*ExternallyFundedTransactor, error) {
	contract, err := bindExternallyFunded(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExternallyFundedTransactor{contract: contract}, nil
}

// NewExternallyFundedFilterer creates a new log filterer instance of ExternallyFunded, bound to a specific deployed contract.
func NewExternallyFundedFilterer(address common.Address, filterer bind.ContractFilterer) (*ExternallyFundedFilterer, error) {
	contract, err := bindExternallyFunded(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExternallyFundedFilterer{contract: contract}, nil
}

// bindExternallyFunded binds a generic wrapper to an already deployed contract.
func bindExternallyFunded(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExternallyFundedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExternallyFunded *ExternallyFundedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExternallyFunded.Contract.ExternallyFundedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExternallyFunded *ExternallyFundedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExternallyFunded.Contract.ExternallyFundedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExternallyFunded *ExternallyFundedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExternallyFunded.Contract.ExternallyFundedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExternallyFunded *ExternallyFundedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExternallyFunded.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExternallyFunded *ExternallyFundedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExternallyFunded.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExternallyFunded *ExternallyFundedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExternallyFunded.Contract.contract.Transact(opts, method, params...)
}

// AuthorizedAccounts is a free data retrieval call binding the contract method 0x24ba5884.
//
// Solidity: function authorizedAccounts(address ) view returns(uint256)
func (_ExternallyFunded *ExternallyFundedCaller) AuthorizedAccounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ExternallyFunded.contract.Call(opts, &out, "authorizedAccounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuthorizedAccounts is a free data retrieval call binding the contract method 0x24ba5884.
//
// Solidity: function authorizedAccounts(address ) view returns(uint256)
func (_ExternallyFunded *ExternallyFundedSession) AuthorizedAccounts(arg0 common.Address) (*big.Int, error) {
	return _ExternallyFunded.Contract.AuthorizedAccounts(&_ExternallyFunded.CallOpts, arg0)
}

// AuthorizedAccounts is a free data retrieval call binding the contract method 0x24ba5884.
//
// Solidity: function authorizedAccounts(address ) view returns(uint256)
func (_ExternallyFunded *ExternallyFundedCallerSession) AuthorizedAccounts(arg0 common.Address) (*big.Int, error) {
	return _ExternallyFunded.Contract.AuthorizedAccounts(&_ExternallyFunded.CallOpts, arg0)
}

// FsmWrapper is a free data retrieval call binding the contract method 0x81db5917.
//
// Solidity: function fsmWrapper() view returns(address)
func (_ExternallyFunded *ExternallyFundedCaller) FsmWrapper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExternallyFunded.contract.Call(opts, &out, "fsmWrapper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FsmWrapper is a free data retrieval call binding the contract method 0x81db5917.
//
// Solidity: function fsmWrapper() view returns(address)
func (_ExternallyFunded *ExternallyFundedSession) FsmWrapper() (common.Address, error) {
	return _ExternallyFunded.Contract.FsmWrapper(&_ExternallyFunded.CallOpts)
}

// FsmWrapper is a free data retrieval call binding the contract method 0x81db5917.
//
// Solidity: function fsmWrapper() view returns(address)
func (_ExternallyFunded *ExternallyFundedCallerSession) FsmWrapper() (common.Address, error) {
	return _ExternallyFunded.Contract.FsmWrapper(&_ExternallyFunded.CallOpts)
}

// GetNextResultWithValidity is a free data retrieval call binding the contract method 0x4f0a32de.
//
// Solidity: function getNextResultWithValidity() view returns(uint256, bool)
func (_ExternallyFunded *ExternallyFundedCaller) GetNextResultWithValidity(opts *bind.CallOpts) (*big.Int, bool, error) {
	var out []interface{}
	err := _ExternallyFunded.contract.Call(opts, &out, "getNextResultWithValidity")

	if err != nil {
		return *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetNextResultWithValidity is a free data retrieval call binding the contract method 0x4f0a32de.
//
// Solidity: function getNextResultWithValidity() view returns(uint256, bool)
func (_ExternallyFunded *ExternallyFundedSession) GetNextResultWithValidity() (*big.Int, bool, error) {
	return _ExternallyFunded.Contract.GetNextResultWithValidity(&_ExternallyFunded.CallOpts)
}

// GetNextResultWithValidity is a free data retrieval call binding the contract method 0x4f0a32de.
//
// Solidity: function getNextResultWithValidity() view returns(uint256, bool)
func (_ExternallyFunded *ExternallyFundedCallerSession) GetNextResultWithValidity() (*big.Int, bool, error) {
	return _ExternallyFunded.Contract.GetNextResultWithValidity(&_ExternallyFunded.CallOpts)
}

// GetResultWithValidity is a free data retrieval call binding the contract method 0x4fd0ada8.
//
// Solidity: function getResultWithValidity() view returns(uint256, bool)
func (_ExternallyFunded *ExternallyFundedCaller) GetResultWithValidity(opts *bind.CallOpts) (*big.Int, bool, error) {
	var out []interface{}
	err := _ExternallyFunded.contract.Call(opts, &out, "getResultWithValidity")

	if err != nil {
		return *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetResultWithValidity is a free data retrieval call binding the contract method 0x4fd0ada8.
//
// Solidity: function getResultWithValidity() view returns(uint256, bool)
func (_ExternallyFunded *ExternallyFundedSession) GetResultWithValidity() (*big.Int, bool, error) {
	return _ExternallyFunded.Contract.GetResultWithValidity(&_ExternallyFunded.CallOpts)
}

// GetResultWithValidity is a free data retrieval call binding the contract method 0x4fd0ada8.
//
// Solidity: function getResultWithValidity() view returns(uint256, bool)
func (_ExternallyFunded *ExternallyFundedCallerSession) GetResultWithValidity() (*big.Int, bool, error) {
	return _ExternallyFunded.Contract.GetResultWithValidity(&_ExternallyFunded.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint64)
func (_ExternallyFunded *ExternallyFundedCaller) LastUpdateTime(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ExternallyFunded.contract.Call(opts, &out, "lastUpdateTime")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint64)
func (_ExternallyFunded *ExternallyFundedSession) LastUpdateTime() (uint64, error) {
	return _ExternallyFunded.Contract.LastUpdateTime(&_ExternallyFunded.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint64)
func (_ExternallyFunded *ExternallyFundedCallerSession) LastUpdateTime() (uint64, error) {
	return _ExternallyFunded.Contract.LastUpdateTime(&_ExternallyFunded.CallOpts)
}

// PassedDelay is a free data retrieval call binding the contract method 0x63bfa88b.
//
// Solidity: function passedDelay() view returns(bool ok)
func (_ExternallyFunded *ExternallyFundedCaller) PassedDelay(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ExternallyFunded.contract.Call(opts, &out, "passedDelay")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PassedDelay is a free data retrieval call binding the contract method 0x63bfa88b.
//
// Solidity: function passedDelay() view returns(bool ok)
func (_ExternallyFunded *ExternallyFundedSession) PassedDelay() (bool, error) {
	return _ExternallyFunded.Contract.PassedDelay(&_ExternallyFunded.CallOpts)
}

// PassedDelay is a free data retrieval call binding the contract method 0x63bfa88b.
//
// Solidity: function passedDelay() view returns(bool ok)
func (_ExternallyFunded *ExternallyFundedCallerSession) PassedDelay() (bool, error) {
	return _ExternallyFunded.Contract.PassedDelay(&_ExternallyFunded.CallOpts)
}

// PriceSource is a free data retrieval call binding the contract method 0x20531bc9.
//
// Solidity: function priceSource() view returns(address)
func (_ExternallyFunded *ExternallyFundedCaller) PriceSource(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExternallyFunded.contract.Call(opts, &out, "priceSource")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceSource is a free data retrieval call binding the contract method 0x20531bc9.
//
// Solidity: function priceSource() view returns(address)
func (_ExternallyFunded *ExternallyFundedSession) PriceSource() (common.Address, error) {
	return _ExternallyFunded.Contract.PriceSource(&_ExternallyFunded.CallOpts)
}

// PriceSource is a free data retrieval call binding the contract method 0x20531bc9.
//
// Solidity: function priceSource() view returns(address)
func (_ExternallyFunded *ExternallyFundedCallerSession) PriceSource() (common.Address, error) {
	return _ExternallyFunded.Contract.PriceSource(&_ExternallyFunded.CallOpts)
}

// Read is a free data retrieval call binding the contract method 0x57de26a4.
//
// Solidity: function read() view returns(uint256)
func (_ExternallyFunded *ExternallyFundedCaller) Read(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExternallyFunded.contract.Call(opts, &out, "read")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Read is a free data retrieval call binding the contract method 0x57de26a4.
//
// Solidity: function read() view returns(uint256)
func (_ExternallyFunded *ExternallyFundedSession) Read() (*big.Int, error) {
	return _ExternallyFunded.Contract.Read(&_ExternallyFunded.CallOpts)
}

// Read is a free data retrieval call binding the contract method 0x57de26a4.
//
// Solidity: function read() view returns(uint256)
func (_ExternallyFunded *ExternallyFundedCallerSession) Read() (*big.Int, error) {
	return _ExternallyFunded.Contract.Read(&_ExternallyFunded.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() view returns(uint256)
func (_ExternallyFunded *ExternallyFundedCaller) Stopped(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExternallyFunded.contract.Call(opts, &out, "stopped")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() view returns(uint256)
func (_ExternallyFunded *ExternallyFundedSession) Stopped() (*big.Int, error) {
	return _ExternallyFunded.Contract.Stopped(&_ExternallyFunded.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() view returns(uint256)
func (_ExternallyFunded *ExternallyFundedCallerSession) Stopped() (*big.Int, error) {
	return _ExternallyFunded.Contract.Stopped(&_ExternallyFunded.CallOpts)
}

// UpdateDelay is a free data retrieval call binding the contract method 0x554f94db.
//
// Solidity: function updateDelay() view returns(uint16)
func (_ExternallyFunded *ExternallyFundedCaller) UpdateDelay(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ExternallyFunded.contract.Call(opts, &out, "updateDelay")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// UpdateDelay is a free data retrieval call binding the contract method 0x554f94db.
//
// Solidity: function updateDelay() view returns(uint16)
func (_ExternallyFunded *ExternallyFundedSession) UpdateDelay() (uint16, error) {
	return _ExternallyFunded.Contract.UpdateDelay(&_ExternallyFunded.CallOpts)
}

// UpdateDelay is a free data retrieval call binding the contract method 0x554f94db.
//
// Solidity: function updateDelay() view returns(uint16)
func (_ExternallyFunded *ExternallyFundedCallerSession) UpdateDelay() (uint16, error) {
	return _ExternallyFunded.Contract.UpdateDelay(&_ExternallyFunded.CallOpts)
}

// AddAuthorization is a paid mutator transaction binding the contract method 0x35b28153.
//
// Solidity: function addAuthorization(address account) returns()
func (_ExternallyFunded *ExternallyFundedTransactor) AddAuthorization(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ExternallyFunded.contract.Transact(opts, "addAuthorization", account)
}

// AddAuthorization is a paid mutator transaction binding the contract method 0x35b28153.
//
// Solidity: function addAuthorization(address account) returns()
func (_ExternallyFunded *ExternallyFundedSession) AddAuthorization(account common.Address) (*types.Transaction, error) {
	return _ExternallyFunded.Contract.AddAuthorization(&_ExternallyFunded.TransactOpts, account)
}

// AddAuthorization is a paid mutator transaction binding the contract method 0x35b28153.
//
// Solidity: function addAuthorization(address account) returns()
func (_ExternallyFunded *ExternallyFundedTransactorSession) AddAuthorization(account common.Address) (*types.Transaction, error) {
	return _ExternallyFunded.Contract.AddAuthorization(&_ExternallyFunded.TransactOpts, account)
}

// ChangeDelay is a paid mutator transaction binding the contract method 0x2b9c7ee3.
//
// Solidity: function changeDelay(uint16 delay) returns()
func (_ExternallyFunded *ExternallyFundedTransactor) ChangeDelay(opts *bind.TransactOpts, delay uint16) (*types.Transaction, error) {
	return _ExternallyFunded.contract.Transact(opts, "changeDelay", delay)
}

// ChangeDelay is a paid mutator transaction binding the contract method 0x2b9c7ee3.
//
// Solidity: function changeDelay(uint16 delay) returns()
func (_ExternallyFunded *ExternallyFundedSession) ChangeDelay(delay uint16) (*types.Transaction, error) {
	return _ExternallyFunded.Contract.ChangeDelay(&_ExternallyFunded.TransactOpts, delay)
}

// ChangeDelay is a paid mutator transaction binding the contract method 0x2b9c7ee3.
//
// Solidity: function changeDelay(uint16 delay) returns()
func (_ExternallyFunded *ExternallyFundedTransactorSession) ChangeDelay(delay uint16) (*types.Transaction, error) {
	return _ExternallyFunded.Contract.ChangeDelay(&_ExternallyFunded.TransactOpts, delay)
}

// ChangePriceSource is a paid mutator transaction binding the contract method 0x3fcdfe26.
//
// Solidity: function changePriceSource(address priceSource_) returns()
func (_ExternallyFunded *ExternallyFundedTransactor) ChangePriceSource(opts *bind.TransactOpts, priceSource_ common.Address) (*types.Transaction, error) {
	return _ExternallyFunded.contract.Transact(opts, "changePriceSource", priceSource_)
}

// ChangePriceSource is a paid mutator transaction binding the contract method 0x3fcdfe26.
//
// Solidity: function changePriceSource(address priceSource_) returns()
func (_ExternallyFunded *ExternallyFundedSession) ChangePriceSource(priceSource_ common.Address) (*types.Transaction, error) {
	return _ExternallyFunded.Contract.ChangePriceSource(&_ExternallyFunded.TransactOpts, priceSource_)
}

// ChangePriceSource is a paid mutator transaction binding the contract method 0x3fcdfe26.
//
// Solidity: function changePriceSource(address priceSource_) returns()
func (_ExternallyFunded *ExternallyFundedTransactorSession) ChangePriceSource(priceSource_ common.Address) (*types.Transaction, error) {
	return _ExternallyFunded.Contract.ChangePriceSource(&_ExternallyFunded.TransactOpts, priceSource_)
}

// ModifyParameters is a paid mutator transaction binding the contract method 0x6614f010.
//
// Solidity: function modifyParameters(bytes32 parameter, address val) returns()
func (_ExternallyFunded *ExternallyFundedTransactor) ModifyParameters(opts *bind.TransactOpts, parameter [32]byte, val common.Address) (*types.Transaction, error) {
	return _ExternallyFunded.contract.Transact(opts, "modifyParameters", parameter, val)
}

// ModifyParameters is a paid mutator transaction binding the contract method 0x6614f010.
//
// Solidity: function modifyParameters(bytes32 parameter, address val) returns()
func (_ExternallyFunded *ExternallyFundedSession) ModifyParameters(parameter [32]byte, val common.Address) (*types.Transaction, error) {
	return _ExternallyFunded.Contract.ModifyParameters(&_ExternallyFunded.TransactOpts, parameter, val)
}

// ModifyParameters is a paid mutator transaction binding the contract method 0x6614f010.
//
// Solidity: function modifyParameters(bytes32 parameter, address val) returns()
func (_ExternallyFunded *ExternallyFundedTransactorSession) ModifyParameters(parameter [32]byte, val common.Address) (*types.Transaction, error) {
	return _ExternallyFunded.Contract.ModifyParameters(&_ExternallyFunded.TransactOpts, parameter, val)
}

// RemoveAuthorization is a paid mutator transaction binding the contract method 0x94f3f81d.
//
// Solidity: function removeAuthorization(address account) returns()
func (_ExternallyFunded *ExternallyFundedTransactor) RemoveAuthorization(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ExternallyFunded.contract.Transact(opts, "removeAuthorization", account)
}

// RemoveAuthorization is a paid mutator transaction binding the contract method 0x94f3f81d.
//
// Solidity: function removeAuthorization(address account) returns()
func (_ExternallyFunded *ExternallyFundedSession) RemoveAuthorization(account common.Address) (*types.Transaction, error) {
	return _ExternallyFunded.Contract.RemoveAuthorization(&_ExternallyFunded.TransactOpts, account)
}

// RemoveAuthorization is a paid mutator transaction binding the contract method 0x94f3f81d.
//
// Solidity: function removeAuthorization(address account) returns()
func (_ExternallyFunded *ExternallyFundedTransactorSession) RemoveAuthorization(account common.Address) (*types.Transaction, error) {
	return _ExternallyFunded.Contract.RemoveAuthorization(&_ExternallyFunded.TransactOpts, account)
}

// RestartValue is a paid mutator transaction binding the contract method 0xceedd63d.
//
// Solidity: function restartValue() returns()
func (_ExternallyFunded *ExternallyFundedTransactor) RestartValue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExternallyFunded.contract.Transact(opts, "restartValue")
}

// RestartValue is a paid mutator transaction binding the contract method 0xceedd63d.
//
// Solidity: function restartValue() returns()
func (_ExternallyFunded *ExternallyFundedSession) RestartValue() (*types.Transaction, error) {
	return _ExternallyFunded.Contract.RestartValue(&_ExternallyFunded.TransactOpts)
}

// RestartValue is a paid mutator transaction binding the contract method 0xceedd63d.
//
// Solidity: function restartValue() returns()
func (_ExternallyFunded *ExternallyFundedTransactorSession) RestartValue() (*types.Transaction, error) {
	return _ExternallyFunded.Contract.RestartValue(&_ExternallyFunded.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_ExternallyFunded *ExternallyFundedTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExternallyFunded.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_ExternallyFunded *ExternallyFundedSession) Start() (*types.Transaction, error) {
	return _ExternallyFunded.Contract.Start(&_ExternallyFunded.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_ExternallyFunded *ExternallyFundedTransactorSession) Start() (*types.Transaction, error) {
	return _ExternallyFunded.Contract.Start(&_ExternallyFunded.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_ExternallyFunded *ExternallyFundedTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExternallyFunded.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_ExternallyFunded *ExternallyFundedSession) Stop() (*types.Transaction, error) {
	return _ExternallyFunded.Contract.Stop(&_ExternallyFunded.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_ExternallyFunded *ExternallyFundedTransactorSession) Stop() (*types.Transaction, error) {
	return _ExternallyFunded.Contract.Stop(&_ExternallyFunded.TransactOpts)
}

// UpdateResult is a paid mutator transaction binding the contract method 0x80ebb08e.
//
// Solidity: function updateResult() returns()
func (_ExternallyFunded *ExternallyFundedTransactor) UpdateResult(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExternallyFunded.contract.Transact(opts, "updateResult")
}

// UpdateResult is a paid mutator transaction binding the contract method 0x80ebb08e.
//
// Solidity: function updateResult() returns()
func (_ExternallyFunded *ExternallyFundedSession) UpdateResult() (*types.Transaction, error) {
	return _ExternallyFunded.Contract.UpdateResult(&_ExternallyFunded.TransactOpts)
}

// UpdateResult is a paid mutator transaction binding the contract method 0x80ebb08e.
//
// Solidity: function updateResult() returns()
func (_ExternallyFunded *ExternallyFundedTransactorSession) UpdateResult() (*types.Transaction, error) {
	return _ExternallyFunded.Contract.UpdateResult(&_ExternallyFunded.TransactOpts)
}

// ExternallyFundedAddAuthorizationIterator is returned from FilterAddAuthorization and is used to iterate over the raw logs and unpacked data for AddAuthorization events raised by the ExternallyFunded contract.
type ExternallyFundedAddAuthorizationIterator struct {
	Event *ExternallyFundedAddAuthorization // Event containing the contract specifics and raw log

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
func (it *ExternallyFundedAddAuthorizationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExternallyFundedAddAuthorization)
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
		it.Event = new(ExternallyFundedAddAuthorization)
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
func (it *ExternallyFundedAddAuthorizationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExternallyFundedAddAuthorizationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExternallyFundedAddAuthorization represents a AddAuthorization event raised by the ExternallyFunded contract.
type ExternallyFundedAddAuthorization struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddAuthorization is a free log retrieval operation binding the contract event 0x599a298163e1678bb1c676052a8930bf0b8a1261ed6e01b8a2391e55f7000102.
//
// Solidity: event AddAuthorization(address account)
func (_ExternallyFunded *ExternallyFundedFilterer) FilterAddAuthorization(opts *bind.FilterOpts) (*ExternallyFundedAddAuthorizationIterator, error) {

	logs, sub, err := _ExternallyFunded.contract.FilterLogs(opts, "AddAuthorization")
	if err != nil {
		return nil, err
	}
	return &ExternallyFundedAddAuthorizationIterator{contract: _ExternallyFunded.contract, event: "AddAuthorization", logs: logs, sub: sub}, nil
}

// WatchAddAuthorization is a free log subscription operation binding the contract event 0x599a298163e1678bb1c676052a8930bf0b8a1261ed6e01b8a2391e55f7000102.
//
// Solidity: event AddAuthorization(address account)
func (_ExternallyFunded *ExternallyFundedFilterer) WatchAddAuthorization(opts *bind.WatchOpts, sink chan<- *ExternallyFundedAddAuthorization) (event.Subscription, error) {

	logs, sub, err := _ExternallyFunded.contract.WatchLogs(opts, "AddAuthorization")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExternallyFundedAddAuthorization)
				if err := _ExternallyFunded.contract.UnpackLog(event, "AddAuthorization", log); err != nil {
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

// ParseAddAuthorization is a log parse operation binding the contract event 0x599a298163e1678bb1c676052a8930bf0b8a1261ed6e01b8a2391e55f7000102.
//
// Solidity: event AddAuthorization(address account)
func (_ExternallyFunded *ExternallyFundedFilterer) ParseAddAuthorization(log types.Log) (*ExternallyFundedAddAuthorization, error) {
	event := new(ExternallyFundedAddAuthorization)
	if err := _ExternallyFunded.contract.UnpackLog(event, "AddAuthorization", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExternallyFundedChangeDelayIterator is returned from FilterChangeDelay and is used to iterate over the raw logs and unpacked data for ChangeDelay events raised by the ExternallyFunded contract.
type ExternallyFundedChangeDelayIterator struct {
	Event *ExternallyFundedChangeDelay // Event containing the contract specifics and raw log

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
func (it *ExternallyFundedChangeDelayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExternallyFundedChangeDelay)
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
		it.Event = new(ExternallyFundedChangeDelay)
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
func (it *ExternallyFundedChangeDelayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExternallyFundedChangeDelayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExternallyFundedChangeDelay represents a ChangeDelay event raised by the ExternallyFunded contract.
type ExternallyFundedChangeDelay struct {
	Delay uint16
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterChangeDelay is a free log retrieval operation binding the contract event 0xe5e55d7a00a8da3c2d2be2af849b5680612dc299cf26e12a879220d5f42ea231.
//
// Solidity: event ChangeDelay(uint16 delay)
func (_ExternallyFunded *ExternallyFundedFilterer) FilterChangeDelay(opts *bind.FilterOpts) (*ExternallyFundedChangeDelayIterator, error) {

	logs, sub, err := _ExternallyFunded.contract.FilterLogs(opts, "ChangeDelay")
	if err != nil {
		return nil, err
	}
	return &ExternallyFundedChangeDelayIterator{contract: _ExternallyFunded.contract, event: "ChangeDelay", logs: logs, sub: sub}, nil
}

// WatchChangeDelay is a free log subscription operation binding the contract event 0xe5e55d7a00a8da3c2d2be2af849b5680612dc299cf26e12a879220d5f42ea231.
//
// Solidity: event ChangeDelay(uint16 delay)
func (_ExternallyFunded *ExternallyFundedFilterer) WatchChangeDelay(opts *bind.WatchOpts, sink chan<- *ExternallyFundedChangeDelay) (event.Subscription, error) {

	logs, sub, err := _ExternallyFunded.contract.WatchLogs(opts, "ChangeDelay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExternallyFundedChangeDelay)
				if err := _ExternallyFunded.contract.UnpackLog(event, "ChangeDelay", log); err != nil {
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

// ParseChangeDelay is a log parse operation binding the contract event 0xe5e55d7a00a8da3c2d2be2af849b5680612dc299cf26e12a879220d5f42ea231.
//
// Solidity: event ChangeDelay(uint16 delay)
func (_ExternallyFunded *ExternallyFundedFilterer) ParseChangeDelay(log types.Log) (*ExternallyFundedChangeDelay, error) {
	event := new(ExternallyFundedChangeDelay)
	if err := _ExternallyFunded.contract.UnpackLog(event, "ChangeDelay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExternallyFundedChangePriceSourceIterator is returned from FilterChangePriceSource and is used to iterate over the raw logs and unpacked data for ChangePriceSource events raised by the ExternallyFunded contract.
type ExternallyFundedChangePriceSourceIterator struct {
	Event *ExternallyFundedChangePriceSource // Event containing the contract specifics and raw log

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
func (it *ExternallyFundedChangePriceSourceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExternallyFundedChangePriceSource)
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
		it.Event = new(ExternallyFundedChangePriceSource)
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
func (it *ExternallyFundedChangePriceSourceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExternallyFundedChangePriceSourceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExternallyFundedChangePriceSource represents a ChangePriceSource event raised by the ExternallyFunded contract.
type ExternallyFundedChangePriceSource struct {
	PriceSource common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterChangePriceSource is a free log retrieval operation binding the contract event 0xe64aaa084ff9fb9651f314c5d25888701090b34db4f4809b517e863c62bd8508.
//
// Solidity: event ChangePriceSource(address priceSource)
func (_ExternallyFunded *ExternallyFundedFilterer) FilterChangePriceSource(opts *bind.FilterOpts) (*ExternallyFundedChangePriceSourceIterator, error) {

	logs, sub, err := _ExternallyFunded.contract.FilterLogs(opts, "ChangePriceSource")
	if err != nil {
		return nil, err
	}
	return &ExternallyFundedChangePriceSourceIterator{contract: _ExternallyFunded.contract, event: "ChangePriceSource", logs: logs, sub: sub}, nil
}

// WatchChangePriceSource is a free log subscription operation binding the contract event 0xe64aaa084ff9fb9651f314c5d25888701090b34db4f4809b517e863c62bd8508.
//
// Solidity: event ChangePriceSource(address priceSource)
func (_ExternallyFunded *ExternallyFundedFilterer) WatchChangePriceSource(opts *bind.WatchOpts, sink chan<- *ExternallyFundedChangePriceSource) (event.Subscription, error) {

	logs, sub, err := _ExternallyFunded.contract.WatchLogs(opts, "ChangePriceSource")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExternallyFundedChangePriceSource)
				if err := _ExternallyFunded.contract.UnpackLog(event, "ChangePriceSource", log); err != nil {
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

// ParseChangePriceSource is a log parse operation binding the contract event 0xe64aaa084ff9fb9651f314c5d25888701090b34db4f4809b517e863c62bd8508.
//
// Solidity: event ChangePriceSource(address priceSource)
func (_ExternallyFunded *ExternallyFundedFilterer) ParseChangePriceSource(log types.Log) (*ExternallyFundedChangePriceSource, error) {
	event := new(ExternallyFundedChangePriceSource)
	if err := _ExternallyFunded.contract.UnpackLog(event, "ChangePriceSource", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExternallyFundedFailRenumerateCallerIterator is returned from FilterFailRenumerateCaller and is used to iterate over the raw logs and unpacked data for FailRenumerateCaller events raised by the ExternallyFunded contract.
type ExternallyFundedFailRenumerateCallerIterator struct {
	Event *ExternallyFundedFailRenumerateCaller // Event containing the contract specifics and raw log

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
func (it *ExternallyFundedFailRenumerateCallerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExternallyFundedFailRenumerateCaller)
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
		it.Event = new(ExternallyFundedFailRenumerateCaller)
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
func (it *ExternallyFundedFailRenumerateCallerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExternallyFundedFailRenumerateCallerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExternallyFundedFailRenumerateCaller represents a FailRenumerateCaller event raised by the ExternallyFunded contract.
type ExternallyFundedFailRenumerateCaller struct {
	Wrapper common.Address
	Caller  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFailRenumerateCaller is a free log retrieval operation binding the contract event 0x96957874bf643b6d75552ddf6e2f265e2c481209f5b12355e64caede7201b740.
//
// Solidity: event FailRenumerateCaller(address wrapper, address caller)
func (_ExternallyFunded *ExternallyFundedFilterer) FilterFailRenumerateCaller(opts *bind.FilterOpts) (*ExternallyFundedFailRenumerateCallerIterator, error) {

	logs, sub, err := _ExternallyFunded.contract.FilterLogs(opts, "FailRenumerateCaller")
	if err != nil {
		return nil, err
	}
	return &ExternallyFundedFailRenumerateCallerIterator{contract: _ExternallyFunded.contract, event: "FailRenumerateCaller", logs: logs, sub: sub}, nil
}

// WatchFailRenumerateCaller is a free log subscription operation binding the contract event 0x96957874bf643b6d75552ddf6e2f265e2c481209f5b12355e64caede7201b740.
//
// Solidity: event FailRenumerateCaller(address wrapper, address caller)
func (_ExternallyFunded *ExternallyFundedFilterer) WatchFailRenumerateCaller(opts *bind.WatchOpts, sink chan<- *ExternallyFundedFailRenumerateCaller) (event.Subscription, error) {

	logs, sub, err := _ExternallyFunded.contract.WatchLogs(opts, "FailRenumerateCaller")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExternallyFundedFailRenumerateCaller)
				if err := _ExternallyFunded.contract.UnpackLog(event, "FailRenumerateCaller", log); err != nil {
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

// ParseFailRenumerateCaller is a log parse operation binding the contract event 0x96957874bf643b6d75552ddf6e2f265e2c481209f5b12355e64caede7201b740.
//
// Solidity: event FailRenumerateCaller(address wrapper, address caller)
func (_ExternallyFunded *ExternallyFundedFilterer) ParseFailRenumerateCaller(log types.Log) (*ExternallyFundedFailRenumerateCaller, error) {
	event := new(ExternallyFundedFailRenumerateCaller)
	if err := _ExternallyFunded.contract.UnpackLog(event, "FailRenumerateCaller", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExternallyFundedModifyParametersIterator is returned from FilterModifyParameters and is used to iterate over the raw logs and unpacked data for ModifyParameters events raised by the ExternallyFunded contract.
type ExternallyFundedModifyParametersIterator struct {
	Event *ExternallyFundedModifyParameters // Event containing the contract specifics and raw log

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
func (it *ExternallyFundedModifyParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExternallyFundedModifyParameters)
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
		it.Event = new(ExternallyFundedModifyParameters)
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
func (it *ExternallyFundedModifyParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExternallyFundedModifyParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExternallyFundedModifyParameters represents a ModifyParameters event raised by the ExternallyFunded contract.
type ExternallyFundedModifyParameters struct {
	Parameter [32]byte
	Val       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterModifyParameters is a free log retrieval operation binding the contract event 0xac7c5c1afaef770ec56ac6268cd3f2fbb1035858ead2601d6553157c33036c3a.
//
// Solidity: event ModifyParameters(bytes32 parameter, uint256 val)
func (_ExternallyFunded *ExternallyFundedFilterer) FilterModifyParameters(opts *bind.FilterOpts) (*ExternallyFundedModifyParametersIterator, error) {

	logs, sub, err := _ExternallyFunded.contract.FilterLogs(opts, "ModifyParameters")
	if err != nil {
		return nil, err
	}
	return &ExternallyFundedModifyParametersIterator{contract: _ExternallyFunded.contract, event: "ModifyParameters", logs: logs, sub: sub}, nil
}

// WatchModifyParameters is a free log subscription operation binding the contract event 0xac7c5c1afaef770ec56ac6268cd3f2fbb1035858ead2601d6553157c33036c3a.
//
// Solidity: event ModifyParameters(bytes32 parameter, uint256 val)
func (_ExternallyFunded *ExternallyFundedFilterer) WatchModifyParameters(opts *bind.WatchOpts, sink chan<- *ExternallyFundedModifyParameters) (event.Subscription, error) {

	logs, sub, err := _ExternallyFunded.contract.WatchLogs(opts, "ModifyParameters")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExternallyFundedModifyParameters)
				if err := _ExternallyFunded.contract.UnpackLog(event, "ModifyParameters", log); err != nil {
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

// ParseModifyParameters is a log parse operation binding the contract event 0xac7c5c1afaef770ec56ac6268cd3f2fbb1035858ead2601d6553157c33036c3a.
//
// Solidity: event ModifyParameters(bytes32 parameter, uint256 val)
func (_ExternallyFunded *ExternallyFundedFilterer) ParseModifyParameters(log types.Log) (*ExternallyFundedModifyParameters, error) {
	event := new(ExternallyFundedModifyParameters)
	if err := _ExternallyFunded.contract.UnpackLog(event, "ModifyParameters", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExternallyFundedModifyParameters0Iterator is returned from FilterModifyParameters0 and is used to iterate over the raw logs and unpacked data for ModifyParameters0 events raised by the ExternallyFunded contract.
type ExternallyFundedModifyParameters0Iterator struct {
	Event *ExternallyFundedModifyParameters0 // Event containing the contract specifics and raw log

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
func (it *ExternallyFundedModifyParameters0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExternallyFundedModifyParameters0)
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
		it.Event = new(ExternallyFundedModifyParameters0)
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
func (it *ExternallyFundedModifyParameters0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExternallyFundedModifyParameters0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExternallyFundedModifyParameters0 represents a ModifyParameters0 event raised by the ExternallyFunded contract.
type ExternallyFundedModifyParameters0 struct {
	Parameter [32]byte
	Val       common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterModifyParameters0 is a free log retrieval operation binding the contract event 0xd91f38cf03346b5dc15fb60f9076f866295231ad3c3841a1051f8443f25170d1.
//
// Solidity: event ModifyParameters(bytes32 parameter, address val)
func (_ExternallyFunded *ExternallyFundedFilterer) FilterModifyParameters0(opts *bind.FilterOpts) (*ExternallyFundedModifyParameters0Iterator, error) {

	logs, sub, err := _ExternallyFunded.contract.FilterLogs(opts, "ModifyParameters0")
	if err != nil {
		return nil, err
	}
	return &ExternallyFundedModifyParameters0Iterator{contract: _ExternallyFunded.contract, event: "ModifyParameters0", logs: logs, sub: sub}, nil
}

// WatchModifyParameters0 is a free log subscription operation binding the contract event 0xd91f38cf03346b5dc15fb60f9076f866295231ad3c3841a1051f8443f25170d1.
//
// Solidity: event ModifyParameters(bytes32 parameter, address val)
func (_ExternallyFunded *ExternallyFundedFilterer) WatchModifyParameters0(opts *bind.WatchOpts, sink chan<- *ExternallyFundedModifyParameters0) (event.Subscription, error) {

	logs, sub, err := _ExternallyFunded.contract.WatchLogs(opts, "ModifyParameters0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExternallyFundedModifyParameters0)
				if err := _ExternallyFunded.contract.UnpackLog(event, "ModifyParameters0", log); err != nil {
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

// ParseModifyParameters0 is a log parse operation binding the contract event 0xd91f38cf03346b5dc15fb60f9076f866295231ad3c3841a1051f8443f25170d1.
//
// Solidity: event ModifyParameters(bytes32 parameter, address val)
func (_ExternallyFunded *ExternallyFundedFilterer) ParseModifyParameters0(log types.Log) (*ExternallyFundedModifyParameters0, error) {
	event := new(ExternallyFundedModifyParameters0)
	if err := _ExternallyFunded.contract.UnpackLog(event, "ModifyParameters0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExternallyFundedRemoveAuthorizationIterator is returned from FilterRemoveAuthorization and is used to iterate over the raw logs and unpacked data for RemoveAuthorization events raised by the ExternallyFunded contract.
type ExternallyFundedRemoveAuthorizationIterator struct {
	Event *ExternallyFundedRemoveAuthorization // Event containing the contract specifics and raw log

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
func (it *ExternallyFundedRemoveAuthorizationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExternallyFundedRemoveAuthorization)
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
		it.Event = new(ExternallyFundedRemoveAuthorization)
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
func (it *ExternallyFundedRemoveAuthorizationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExternallyFundedRemoveAuthorizationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExternallyFundedRemoveAuthorization represents a RemoveAuthorization event raised by the ExternallyFunded contract.
type ExternallyFundedRemoveAuthorization struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRemoveAuthorization is a free log retrieval operation binding the contract event 0x8834a87e641e9716be4f34527af5d23e11624f1ddeefede6ad75a9acfc31b903.
//
// Solidity: event RemoveAuthorization(address account)
func (_ExternallyFunded *ExternallyFundedFilterer) FilterRemoveAuthorization(opts *bind.FilterOpts) (*ExternallyFundedRemoveAuthorizationIterator, error) {

	logs, sub, err := _ExternallyFunded.contract.FilterLogs(opts, "RemoveAuthorization")
	if err != nil {
		return nil, err
	}
	return &ExternallyFundedRemoveAuthorizationIterator{contract: _ExternallyFunded.contract, event: "RemoveAuthorization", logs: logs, sub: sub}, nil
}

// WatchRemoveAuthorization is a free log subscription operation binding the contract event 0x8834a87e641e9716be4f34527af5d23e11624f1ddeefede6ad75a9acfc31b903.
//
// Solidity: event RemoveAuthorization(address account)
func (_ExternallyFunded *ExternallyFundedFilterer) WatchRemoveAuthorization(opts *bind.WatchOpts, sink chan<- *ExternallyFundedRemoveAuthorization) (event.Subscription, error) {

	logs, sub, err := _ExternallyFunded.contract.WatchLogs(opts, "RemoveAuthorization")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExternallyFundedRemoveAuthorization)
				if err := _ExternallyFunded.contract.UnpackLog(event, "RemoveAuthorization", log); err != nil {
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

// ParseRemoveAuthorization is a log parse operation binding the contract event 0x8834a87e641e9716be4f34527af5d23e11624f1ddeefede6ad75a9acfc31b903.
//
// Solidity: event RemoveAuthorization(address account)
func (_ExternallyFunded *ExternallyFundedFilterer) ParseRemoveAuthorization(log types.Log) (*ExternallyFundedRemoveAuthorization, error) {
	event := new(ExternallyFundedRemoveAuthorization)
	if err := _ExternallyFunded.contract.UnpackLog(event, "RemoveAuthorization", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExternallyFundedRestartValueIterator is returned from FilterRestartValue and is used to iterate over the raw logs and unpacked data for RestartValue events raised by the ExternallyFunded contract.
type ExternallyFundedRestartValueIterator struct {
	Event *ExternallyFundedRestartValue // Event containing the contract specifics and raw log

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
func (it *ExternallyFundedRestartValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExternallyFundedRestartValue)
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
		it.Event = new(ExternallyFundedRestartValue)
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
func (it *ExternallyFundedRestartValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExternallyFundedRestartValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExternallyFundedRestartValue represents a RestartValue event raised by the ExternallyFunded contract.
type ExternallyFundedRestartValue struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRestartValue is a free log retrieval operation binding the contract event 0x34cd470f8814066090f532fe3cb8158c5d974d7a01f4ada846dc9869e7e9d59a.
//
// Solidity: event RestartValue()
func (_ExternallyFunded *ExternallyFundedFilterer) FilterRestartValue(opts *bind.FilterOpts) (*ExternallyFundedRestartValueIterator, error) {

	logs, sub, err := _ExternallyFunded.contract.FilterLogs(opts, "RestartValue")
	if err != nil {
		return nil, err
	}
	return &ExternallyFundedRestartValueIterator{contract: _ExternallyFunded.contract, event: "RestartValue", logs: logs, sub: sub}, nil
}

// WatchRestartValue is a free log subscription operation binding the contract event 0x34cd470f8814066090f532fe3cb8158c5d974d7a01f4ada846dc9869e7e9d59a.
//
// Solidity: event RestartValue()
func (_ExternallyFunded *ExternallyFundedFilterer) WatchRestartValue(opts *bind.WatchOpts, sink chan<- *ExternallyFundedRestartValue) (event.Subscription, error) {

	logs, sub, err := _ExternallyFunded.contract.WatchLogs(opts, "RestartValue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExternallyFundedRestartValue)
				if err := _ExternallyFunded.contract.UnpackLog(event, "RestartValue", log); err != nil {
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

// ParseRestartValue is a log parse operation binding the contract event 0x34cd470f8814066090f532fe3cb8158c5d974d7a01f4ada846dc9869e7e9d59a.
//
// Solidity: event RestartValue()
func (_ExternallyFunded *ExternallyFundedFilterer) ParseRestartValue(log types.Log) (*ExternallyFundedRestartValue, error) {
	event := new(ExternallyFundedRestartValue)
	if err := _ExternallyFunded.contract.UnpackLog(event, "RestartValue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExternallyFundedStartIterator is returned from FilterStart and is used to iterate over the raw logs and unpacked data for Start events raised by the ExternallyFunded contract.
type ExternallyFundedStartIterator struct {
	Event *ExternallyFundedStart // Event containing the contract specifics and raw log

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
func (it *ExternallyFundedStartIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExternallyFundedStart)
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
		it.Event = new(ExternallyFundedStart)
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
func (it *ExternallyFundedStartIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExternallyFundedStartIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExternallyFundedStart represents a Start event raised by the ExternallyFunded contract.
type ExternallyFundedStart struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStart is a free log retrieval operation binding the contract event 0x1b55ba3aa851a46be3b365aee5b5c140edd620d578922f3e8466d2cbd96f954b.
//
// Solidity: event Start()
func (_ExternallyFunded *ExternallyFundedFilterer) FilterStart(opts *bind.FilterOpts) (*ExternallyFundedStartIterator, error) {

	logs, sub, err := _ExternallyFunded.contract.FilterLogs(opts, "Start")
	if err != nil {
		return nil, err
	}
	return &ExternallyFundedStartIterator{contract: _ExternallyFunded.contract, event: "Start", logs: logs, sub: sub}, nil
}

// WatchStart is a free log subscription operation binding the contract event 0x1b55ba3aa851a46be3b365aee5b5c140edd620d578922f3e8466d2cbd96f954b.
//
// Solidity: event Start()
func (_ExternallyFunded *ExternallyFundedFilterer) WatchStart(opts *bind.WatchOpts, sink chan<- *ExternallyFundedStart) (event.Subscription, error) {

	logs, sub, err := _ExternallyFunded.contract.WatchLogs(opts, "Start")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExternallyFundedStart)
				if err := _ExternallyFunded.contract.UnpackLog(event, "Start", log); err != nil {
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

// ParseStart is a log parse operation binding the contract event 0x1b55ba3aa851a46be3b365aee5b5c140edd620d578922f3e8466d2cbd96f954b.
//
// Solidity: event Start()
func (_ExternallyFunded *ExternallyFundedFilterer) ParseStart(log types.Log) (*ExternallyFundedStart, error) {
	event := new(ExternallyFundedStart)
	if err := _ExternallyFunded.contract.UnpackLog(event, "Start", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExternallyFundedStopIterator is returned from FilterStop and is used to iterate over the raw logs and unpacked data for Stop events raised by the ExternallyFunded contract.
type ExternallyFundedStopIterator struct {
	Event *ExternallyFundedStop // Event containing the contract specifics and raw log

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
func (it *ExternallyFundedStopIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExternallyFundedStop)
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
		it.Event = new(ExternallyFundedStop)
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
func (it *ExternallyFundedStopIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExternallyFundedStopIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExternallyFundedStop represents a Stop event raised by the ExternallyFunded contract.
type ExternallyFundedStop struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStop is a free log retrieval operation binding the contract event 0xbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b.
//
// Solidity: event Stop()
func (_ExternallyFunded *ExternallyFundedFilterer) FilterStop(opts *bind.FilterOpts) (*ExternallyFundedStopIterator, error) {

	logs, sub, err := _ExternallyFunded.contract.FilterLogs(opts, "Stop")
	if err != nil {
		return nil, err
	}
	return &ExternallyFundedStopIterator{contract: _ExternallyFunded.contract, event: "Stop", logs: logs, sub: sub}, nil
}

// WatchStop is a free log subscription operation binding the contract event 0xbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b.
//
// Solidity: event Stop()
func (_ExternallyFunded *ExternallyFundedFilterer) WatchStop(opts *bind.WatchOpts, sink chan<- *ExternallyFundedStop) (event.Subscription, error) {

	logs, sub, err := _ExternallyFunded.contract.WatchLogs(opts, "Stop")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExternallyFundedStop)
				if err := _ExternallyFunded.contract.UnpackLog(event, "Stop", log); err != nil {
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

// ParseStop is a log parse operation binding the contract event 0xbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b.
//
// Solidity: event Stop()
func (_ExternallyFunded *ExternallyFundedFilterer) ParseStop(log types.Log) (*ExternallyFundedStop, error) {
	event := new(ExternallyFundedStop)
	if err := _ExternallyFunded.contract.UnpackLog(event, "Stop", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExternallyFundedUpdateResultIterator is returned from FilterUpdateResult and is used to iterate over the raw logs and unpacked data for UpdateResult events raised by the ExternallyFunded contract.
type ExternallyFundedUpdateResultIterator struct {
	Event *ExternallyFundedUpdateResult // Event containing the contract specifics and raw log

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
func (it *ExternallyFundedUpdateResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExternallyFundedUpdateResult)
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
		it.Event = new(ExternallyFundedUpdateResult)
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
func (it *ExternallyFundedUpdateResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExternallyFundedUpdateResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExternallyFundedUpdateResult represents a UpdateResult event raised by the ExternallyFunded contract.
type ExternallyFundedUpdateResult struct {
	NewMedian      *big.Int
	LastUpdateTime *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpdateResult is a free log retrieval operation binding the contract event 0xf85e44c6c3597d176b8d59bfbf500dfdb2badfc8cf91e6d960b16583a5807e48.
//
// Solidity: event UpdateResult(uint256 newMedian, uint256 lastUpdateTime)
func (_ExternallyFunded *ExternallyFundedFilterer) FilterUpdateResult(opts *bind.FilterOpts) (*ExternallyFundedUpdateResultIterator, error) {

	logs, sub, err := _ExternallyFunded.contract.FilterLogs(opts, "UpdateResult")
	if err != nil {
		return nil, err
	}
	return &ExternallyFundedUpdateResultIterator{contract: _ExternallyFunded.contract, event: "UpdateResult", logs: logs, sub: sub}, nil
}

// WatchUpdateResult is a free log subscription operation binding the contract event 0xf85e44c6c3597d176b8d59bfbf500dfdb2badfc8cf91e6d960b16583a5807e48.
//
// Solidity: event UpdateResult(uint256 newMedian, uint256 lastUpdateTime)
func (_ExternallyFunded *ExternallyFundedFilterer) WatchUpdateResult(opts *bind.WatchOpts, sink chan<- *ExternallyFundedUpdateResult) (event.Subscription, error) {

	logs, sub, err := _ExternallyFunded.contract.WatchLogs(opts, "UpdateResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExternallyFundedUpdateResult)
				if err := _ExternallyFunded.contract.UnpackLog(event, "UpdateResult", log); err != nil {
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

// ParseUpdateResult is a log parse operation binding the contract event 0xf85e44c6c3597d176b8d59bfbf500dfdb2badfc8cf91e6d960b16583a5807e48.
//
// Solidity: event UpdateResult(uint256 newMedian, uint256 lastUpdateTime)
func (_ExternallyFunded *ExternallyFundedFilterer) ParseUpdateResult(log types.Log) (*ExternallyFundedUpdateResult, error) {
	event := new(ExternallyFundedUpdateResult)
	if err := _ExternallyFunded.contract.UnpackLog(event, "UpdateResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
