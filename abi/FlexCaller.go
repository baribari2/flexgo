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
	_ = abi.ConvertType
)

// FlexCallerMetaData contains all meta data concerning the FlexCaller contract.
var FlexCallerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"flexCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FlexCallerABI is the input ABI used to generate the binding from.
// Deprecated: Use FlexCallerMetaData.ABI instead.
var FlexCallerABI = FlexCallerMetaData.ABI

// FlexCaller is an auto generated Go binding around an Ethereum contract.
type FlexCaller struct {
	FlexCallerCaller     // Read-only binding to the contract
	FlexCallerTransactor // Write-only binding to the contract
	FlexCallerFilterer   // Log filterer for contract events
}

// FlexCallerCaller is an auto generated read-only Go binding around an Ethereum contract.
type FlexCallerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlexCallerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FlexCallerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlexCallerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FlexCallerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlexCallerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FlexCallerSession struct {
	Contract     *FlexCaller       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FlexCallerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FlexCallerCallerSession struct {
	Contract *FlexCallerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// FlexCallerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FlexCallerTransactorSession struct {
	Contract     *FlexCallerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FlexCallerRaw is an auto generated low-level Go binding around an Ethereum contract.
type FlexCallerRaw struct {
	Contract *FlexCaller // Generic contract binding to access the raw methods on
}

// FlexCallerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FlexCallerCallerRaw struct {
	Contract *FlexCallerCaller // Generic read-only contract binding to access the raw methods on
}

// FlexCallerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FlexCallerTransactorRaw struct {
	Contract *FlexCallerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFlexCaller creates a new instance of FlexCaller, bound to a specific deployed contract.
func NewFlexCaller(address common.Address, backend bind.ContractBackend) (*FlexCaller, error) {
	contract, err := bindFlexCaller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FlexCaller{FlexCallerCaller: FlexCallerCaller{contract: contract}, FlexCallerTransactor: FlexCallerTransactor{contract: contract}, FlexCallerFilterer: FlexCallerFilterer{contract: contract}}, nil
}

// NewFlexCallerCaller creates a new read-only instance of FlexCaller, bound to a specific deployed contract.
func NewFlexCallerCaller(address common.Address, caller bind.ContractCaller) (*FlexCallerCaller, error) {
	contract, err := bindFlexCaller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FlexCallerCaller{contract: contract}, nil
}

// NewFlexCallerTransactor creates a new write-only instance of FlexCaller, bound to a specific deployed contract.
func NewFlexCallerTransactor(address common.Address, transactor bind.ContractTransactor) (*FlexCallerTransactor, error) {
	contract, err := bindFlexCaller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FlexCallerTransactor{contract: contract}, nil
}

// NewFlexCallerFilterer creates a new log filterer instance of FlexCaller, bound to a specific deployed contract.
func NewFlexCallerFilterer(address common.Address, filterer bind.ContractFilterer) (*FlexCallerFilterer, error) {
	contract, err := bindFlexCaller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FlexCallerFilterer{contract: contract}, nil
}

// bindFlexCaller binds a generic wrapper to an already deployed contract.
func bindFlexCaller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FlexCallerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlexCaller *FlexCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlexCaller.Contract.FlexCallerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlexCaller *FlexCallerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlexCaller.Contract.FlexCallerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlexCaller *FlexCallerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlexCaller.Contract.FlexCallerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlexCaller *FlexCallerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlexCaller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlexCaller *FlexCallerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlexCaller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlexCaller *FlexCallerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlexCaller.Contract.contract.Transact(opts, method, params...)
}

// FlexCall is a paid mutator transaction binding the contract method 0x9942176c.
//
// Solidity: function flexCall() returns()
func (_FlexCaller *FlexCallerTransactor) FlexCall(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlexCaller.contract.Transact(opts, "flexCall")
}

// FlexCall is a paid mutator transaction binding the contract method 0x9942176c.
//
// Solidity: function flexCall() returns()
func (_FlexCaller *FlexCallerSession) FlexCall() (*types.Transaction, error) {
	return _FlexCaller.Contract.FlexCall(&_FlexCaller.TransactOpts)
}

// FlexCall is a paid mutator transaction binding the contract method 0x9942176c.
//
// Solidity: function flexCall() returns()
func (_FlexCaller *FlexCallerTransactorSession) FlexCall() (*types.Transaction, error) {
	return _FlexCaller.Contract.FlexCall(&_FlexCaller.TransactOpts)
}
