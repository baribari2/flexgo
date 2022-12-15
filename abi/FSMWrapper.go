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

// FSMMetaData contains all meta data concerning the FSM contract.
var FSMMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fsm_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"reimburseDelay_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddAuthorization\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertReason\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FailRewardCaller\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"parameter\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"ModifyParameters\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"parameter\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"ModifyParameters\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RemoveAuthorization\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"RAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WAD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addAuthorization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"addition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"z\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorizedAccounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseUpdateCallerReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fsm\",\"outputs\":[{\"internalType\":\"contractFSMLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeOfLastUpdate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"defaultDelayBetweenCalls\",\"type\":\"uint256\"}],\"name\":\"getCallerReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextBoundedPrice\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextPriceLowerBound\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextPriceUpperBound\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextResultWithValidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getResultWithValidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastReimburseTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxRewardIncreaseDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxUpdateCallerReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"minimum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"z\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"parameter\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"modifyParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"parameter\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"modifyParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"multiply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"z\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newPriceDeviation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"passedDelay\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"perSecondCallerRewardIncrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceSource\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"rad\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"z\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"ray\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"z\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"rdivide\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"z\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"read\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reimburseDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeAuthorization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"}],\"name\":\"renumerateCaller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"rmultiply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"z\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"base\",\"type\":\"uint256\"}],\"name\":\"rpower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"z\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"subtract\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"z\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"contractStabilityFeeTreasuryLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryAllowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateDelay\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"wdivide\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"z\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"wmultiply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"z\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// FSMABI is the input ABI used to generate the binding from.
// Deprecated: Use FSMMetaData.ABI instead.
var FSMABI = FSMMetaData.ABI

// FSM is an auto generated Go binding around an Ethereum contract.
type FSM struct {
	FSMCaller     // Read-only binding to the contract
	FSMTransactor // Write-only binding to the contract
	FSMFilterer   // Log filterer for contract events
}

// FSMCaller is an auto generated read-only Go binding around an Ethereum contract.
type FSMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FSMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FSMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FSMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FSMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FSMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FSMSession struct {
	Contract     *FSM              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FSMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FSMCallerSession struct {
	Contract *FSMCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FSMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FSMTransactorSession struct {
	Contract     *FSMTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FSMRaw is an auto generated low-level Go binding around an Ethereum contract.
type FSMRaw struct {
	Contract *FSM // Generic contract binding to access the raw methods on
}

// FSMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FSMCallerRaw struct {
	Contract *FSMCaller // Generic read-only contract binding to access the raw methods on
}

// FSMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FSMTransactorRaw struct {
	Contract *FSMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFSM creates a new instance of FSM, bound to a specific deployed contract.
func NewFSM(address common.Address, backend bind.ContractBackend) (*FSM, error) {
	contract, err := bindFSM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FSM{FSMCaller: FSMCaller{contract: contract}, FSMTransactor: FSMTransactor{contract: contract}, FSMFilterer: FSMFilterer{contract: contract}}, nil
}

// NewFSMCaller creates a new read-only instance of FSM, bound to a specific deployed contract.
func NewFSMCaller(address common.Address, caller bind.ContractCaller) (*FSMCaller, error) {
	contract, err := bindFSM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FSMCaller{contract: contract}, nil
}

// NewFSMTransactor creates a new write-only instance of FSM, bound to a specific deployed contract.
func NewFSMTransactor(address common.Address, transactor bind.ContractTransactor) (*FSMTransactor, error) {
	contract, err := bindFSM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FSMTransactor{contract: contract}, nil
}

// NewFSMFilterer creates a new log filterer instance of FSM, bound to a specific deployed contract.
func NewFSMFilterer(address common.Address, filterer bind.ContractFilterer) (*FSMFilterer, error) {
	contract, err := bindFSM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FSMFilterer{contract: contract}, nil
}

// bindFSM binds a generic wrapper to an already deployed contract.
func bindFSM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FSMMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FSM *FSMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FSM.Contract.FSMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FSM *FSMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FSM.Contract.FSMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FSM *FSMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FSM.Contract.FSMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FSM *FSMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FSM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FSM *FSMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FSM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FSM *FSMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FSM.Contract.contract.Transact(opts, method, params...)
}

// RAY is a free data retrieval call binding the contract method 0x552033c4.
//
// Solidity: function RAY() view returns(uint256)
func (_FSM *FSMCaller) RAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "RAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RAY is a free data retrieval call binding the contract method 0x552033c4.
//
// Solidity: function RAY() view returns(uint256)
func (_FSM *FSMSession) RAY() (*big.Int, error) {
	return _FSM.Contract.RAY(&_FSM.CallOpts)
}

// RAY is a free data retrieval call binding the contract method 0x552033c4.
//
// Solidity: function RAY() view returns(uint256)
func (_FSM *FSMCallerSession) RAY() (*big.Int, error) {
	return _FSM.Contract.RAY(&_FSM.CallOpts)
}

// WAD is a free data retrieval call binding the contract method 0x6a146024.
//
// Solidity: function WAD() view returns(uint256)
func (_FSM *FSMCaller) WAD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "WAD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WAD is a free data retrieval call binding the contract method 0x6a146024.
//
// Solidity: function WAD() view returns(uint256)
func (_FSM *FSMSession) WAD() (*big.Int, error) {
	return _FSM.Contract.WAD(&_FSM.CallOpts)
}

// WAD is a free data retrieval call binding the contract method 0x6a146024.
//
// Solidity: function WAD() view returns(uint256)
func (_FSM *FSMCallerSession) WAD() (*big.Int, error) {
	return _FSM.Contract.WAD(&_FSM.CallOpts)
}

// Addition is a free data retrieval call binding the contract method 0x54f363a3.
//
// Solidity: function addition(uint256 x, uint256 y) pure returns(uint256 z)
func (_FSM *FSMCaller) Addition(opts *bind.CallOpts, x *big.Int, y *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "addition", x, y)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Addition is a free data retrieval call binding the contract method 0x54f363a3.
//
// Solidity: function addition(uint256 x, uint256 y) pure returns(uint256 z)
func (_FSM *FSMSession) Addition(x *big.Int, y *big.Int) (*big.Int, error) {
	return _FSM.Contract.Addition(&_FSM.CallOpts, x, y)
}

// Addition is a free data retrieval call binding the contract method 0x54f363a3.
//
// Solidity: function addition(uint256 x, uint256 y) pure returns(uint256 z)
func (_FSM *FSMCallerSession) Addition(x *big.Int, y *big.Int) (*big.Int, error) {
	return _FSM.Contract.Addition(&_FSM.CallOpts, x, y)
}

// AuthorizedAccounts is a free data retrieval call binding the contract method 0x24ba5884.
//
// Solidity: function authorizedAccounts(address ) view returns(uint256)
func (_FSM *FSMCaller) AuthorizedAccounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "authorizedAccounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuthorizedAccounts is a free data retrieval call binding the contract method 0x24ba5884.
//
// Solidity: function authorizedAccounts(address ) view returns(uint256)
func (_FSM *FSMSession) AuthorizedAccounts(arg0 common.Address) (*big.Int, error) {
	return _FSM.Contract.AuthorizedAccounts(&_FSM.CallOpts, arg0)
}

// AuthorizedAccounts is a free data retrieval call binding the contract method 0x24ba5884.
//
// Solidity: function authorizedAccounts(address ) view returns(uint256)
func (_FSM *FSMCallerSession) AuthorizedAccounts(arg0 common.Address) (*big.Int, error) {
	return _FSM.Contract.AuthorizedAccounts(&_FSM.CallOpts, arg0)
}

// BaseUpdateCallerReward is a free data retrieval call binding the contract method 0x1c1f908c.
//
// Solidity: function baseUpdateCallerReward() view returns(uint256)
func (_FSM *FSMCaller) BaseUpdateCallerReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "baseUpdateCallerReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseUpdateCallerReward is a free data retrieval call binding the contract method 0x1c1f908c.
//
// Solidity: function baseUpdateCallerReward() view returns(uint256)
func (_FSM *FSMSession) BaseUpdateCallerReward() (*big.Int, error) {
	return _FSM.Contract.BaseUpdateCallerReward(&_FSM.CallOpts)
}

// BaseUpdateCallerReward is a free data retrieval call binding the contract method 0x1c1f908c.
//
// Solidity: function baseUpdateCallerReward() view returns(uint256)
func (_FSM *FSMCallerSession) BaseUpdateCallerReward() (*big.Int, error) {
	return _FSM.Contract.BaseUpdateCallerReward(&_FSM.CallOpts)
}

// Fsm is a free data retrieval call binding the contract method 0x25b2cda1.
//
// Solidity: function fsm() view returns(address)
func (_FSM *FSMCaller) Fsm(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "fsm")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Fsm is a free data retrieval call binding the contract method 0x25b2cda1.
//
// Solidity: function fsm() view returns(address)
func (_FSM *FSMSession) Fsm() (common.Address, error) {
	return _FSM.Contract.Fsm(&_FSM.CallOpts)
}

// Fsm is a free data retrieval call binding the contract method 0x25b2cda1.
//
// Solidity: function fsm() view returns(address)
func (_FSM *FSMCallerSession) Fsm() (common.Address, error) {
	return _FSM.Contract.Fsm(&_FSM.CallOpts)
}

// GetCallerReward is a free data retrieval call binding the contract method 0xf238ffd2.
//
// Solidity: function getCallerReward(uint256 timeOfLastUpdate, uint256 defaultDelayBetweenCalls) view returns(uint256)
func (_FSM *FSMCaller) GetCallerReward(opts *bind.CallOpts, timeOfLastUpdate *big.Int, defaultDelayBetweenCalls *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "getCallerReward", timeOfLastUpdate, defaultDelayBetweenCalls)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCallerReward is a free data retrieval call binding the contract method 0xf238ffd2.
//
// Solidity: function getCallerReward(uint256 timeOfLastUpdate, uint256 defaultDelayBetweenCalls) view returns(uint256)
func (_FSM *FSMSession) GetCallerReward(timeOfLastUpdate *big.Int, defaultDelayBetweenCalls *big.Int) (*big.Int, error) {
	return _FSM.Contract.GetCallerReward(&_FSM.CallOpts, timeOfLastUpdate, defaultDelayBetweenCalls)
}

// GetCallerReward is a free data retrieval call binding the contract method 0xf238ffd2.
//
// Solidity: function getCallerReward(uint256 timeOfLastUpdate, uint256 defaultDelayBetweenCalls) view returns(uint256)
func (_FSM *FSMCallerSession) GetCallerReward(timeOfLastUpdate *big.Int, defaultDelayBetweenCalls *big.Int) (*big.Int, error) {
	return _FSM.Contract.GetCallerReward(&_FSM.CallOpts, timeOfLastUpdate, defaultDelayBetweenCalls)
}

// GetNextBoundedPrice is a free data retrieval call binding the contract method 0x3cff8d0b.
//
// Solidity: function getNextBoundedPrice() view returns(uint128)
func (_FSM *FSMCaller) GetNextBoundedPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "getNextBoundedPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextBoundedPrice is a free data retrieval call binding the contract method 0x3cff8d0b.
//
// Solidity: function getNextBoundedPrice() view returns(uint128)
func (_FSM *FSMSession) GetNextBoundedPrice() (*big.Int, error) {
	return _FSM.Contract.GetNextBoundedPrice(&_FSM.CallOpts)
}

// GetNextBoundedPrice is a free data retrieval call binding the contract method 0x3cff8d0b.
//
// Solidity: function getNextBoundedPrice() view returns(uint128)
func (_FSM *FSMCallerSession) GetNextBoundedPrice() (*big.Int, error) {
	return _FSM.Contract.GetNextBoundedPrice(&_FSM.CallOpts)
}

// GetNextPriceLowerBound is a free data retrieval call binding the contract method 0xddc2ad8f.
//
// Solidity: function getNextPriceLowerBound() view returns(uint128)
func (_FSM *FSMCaller) GetNextPriceLowerBound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "getNextPriceLowerBound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextPriceLowerBound is a free data retrieval call binding the contract method 0xddc2ad8f.
//
// Solidity: function getNextPriceLowerBound() view returns(uint128)
func (_FSM *FSMSession) GetNextPriceLowerBound() (*big.Int, error) {
	return _FSM.Contract.GetNextPriceLowerBound(&_FSM.CallOpts)
}

// GetNextPriceLowerBound is a free data retrieval call binding the contract method 0xddc2ad8f.
//
// Solidity: function getNextPriceLowerBound() view returns(uint128)
func (_FSM *FSMCallerSession) GetNextPriceLowerBound() (*big.Int, error) {
	return _FSM.Contract.GetNextPriceLowerBound(&_FSM.CallOpts)
}

// GetNextPriceUpperBound is a free data retrieval call binding the contract method 0x51330adc.
//
// Solidity: function getNextPriceUpperBound() view returns(uint128)
func (_FSM *FSMCaller) GetNextPriceUpperBound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "getNextPriceUpperBound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextPriceUpperBound is a free data retrieval call binding the contract method 0x51330adc.
//
// Solidity: function getNextPriceUpperBound() view returns(uint128)
func (_FSM *FSMSession) GetNextPriceUpperBound() (*big.Int, error) {
	return _FSM.Contract.GetNextPriceUpperBound(&_FSM.CallOpts)
}

// GetNextPriceUpperBound is a free data retrieval call binding the contract method 0x51330adc.
//
// Solidity: function getNextPriceUpperBound() view returns(uint128)
func (_FSM *FSMCallerSession) GetNextPriceUpperBound() (*big.Int, error) {
	return _FSM.Contract.GetNextPriceUpperBound(&_FSM.CallOpts)
}

// GetNextResultWithValidity is a free data retrieval call binding the contract method 0x4f0a32de.
//
// Solidity: function getNextResultWithValidity() view returns(uint256, bool)
func (_FSM *FSMCaller) GetNextResultWithValidity(opts *bind.CallOpts) (*big.Int, bool, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "getNextResultWithValidity")

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
func (_FSM *FSMSession) GetNextResultWithValidity() (*big.Int, bool, error) {
	return _FSM.Contract.GetNextResultWithValidity(&_FSM.CallOpts)
}

// GetNextResultWithValidity is a free data retrieval call binding the contract method 0x4f0a32de.
//
// Solidity: function getNextResultWithValidity() view returns(uint256, bool)
func (_FSM *FSMCallerSession) GetNextResultWithValidity() (*big.Int, bool, error) {
	return _FSM.Contract.GetNextResultWithValidity(&_FSM.CallOpts)
}

// GetResultWithValidity is a free data retrieval call binding the contract method 0x4fd0ada8.
//
// Solidity: function getResultWithValidity() view returns(uint256, bool)
func (_FSM *FSMCaller) GetResultWithValidity(opts *bind.CallOpts) (*big.Int, bool, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "getResultWithValidity")

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
func (_FSM *FSMSession) GetResultWithValidity() (*big.Int, bool, error) {
	return _FSM.Contract.GetResultWithValidity(&_FSM.CallOpts)
}

// GetResultWithValidity is a free data retrieval call binding the contract method 0x4fd0ada8.
//
// Solidity: function getResultWithValidity() view returns(uint256, bool)
func (_FSM *FSMCallerSession) GetResultWithValidity() (*big.Int, bool, error) {
	return _FSM.Contract.GetResultWithValidity(&_FSM.CallOpts)
}

// LastReimburseTime is a free data retrieval call binding the contract method 0xc0585a4e.
//
// Solidity: function lastReimburseTime() view returns(uint256)
func (_FSM *FSMCaller) LastReimburseTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "lastReimburseTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastReimburseTime is a free data retrieval call binding the contract method 0xc0585a4e.
//
// Solidity: function lastReimburseTime() view returns(uint256)
func (_FSM *FSMSession) LastReimburseTime() (*big.Int, error) {
	return _FSM.Contract.LastReimburseTime(&_FSM.CallOpts)
}

// LastReimburseTime is a free data retrieval call binding the contract method 0xc0585a4e.
//
// Solidity: function lastReimburseTime() view returns(uint256)
func (_FSM *FSMCallerSession) LastReimburseTime() (*big.Int, error) {
	return _FSM.Contract.LastReimburseTime(&_FSM.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint64)
func (_FSM *FSMCaller) LastUpdateTime(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "lastUpdateTime")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint64)
func (_FSM *FSMSession) LastUpdateTime() (uint64, error) {
	return _FSM.Contract.LastUpdateTime(&_FSM.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint64)
func (_FSM *FSMCallerSession) LastUpdateTime() (uint64, error) {
	return _FSM.Contract.LastUpdateTime(&_FSM.CallOpts)
}

// MaxRewardIncreaseDelay is a free data retrieval call binding the contract method 0x2009e568.
//
// Solidity: function maxRewardIncreaseDelay() view returns(uint256)
func (_FSM *FSMCaller) MaxRewardIncreaseDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "maxRewardIncreaseDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRewardIncreaseDelay is a free data retrieval call binding the contract method 0x2009e568.
//
// Solidity: function maxRewardIncreaseDelay() view returns(uint256)
func (_FSM *FSMSession) MaxRewardIncreaseDelay() (*big.Int, error) {
	return _FSM.Contract.MaxRewardIncreaseDelay(&_FSM.CallOpts)
}

// MaxRewardIncreaseDelay is a free data retrieval call binding the contract method 0x2009e568.
//
// Solidity: function maxRewardIncreaseDelay() view returns(uint256)
func (_FSM *FSMCallerSession) MaxRewardIncreaseDelay() (*big.Int, error) {
	return _FSM.Contract.MaxRewardIncreaseDelay(&_FSM.CallOpts)
}

// MaxUpdateCallerReward is a free data retrieval call binding the contract method 0x69dec276.
//
// Solidity: function maxUpdateCallerReward() view returns(uint256)
func (_FSM *FSMCaller) MaxUpdateCallerReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "maxUpdateCallerReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxUpdateCallerReward is a free data retrieval call binding the contract method 0x69dec276.
//
// Solidity: function maxUpdateCallerReward() view returns(uint256)
func (_FSM *FSMSession) MaxUpdateCallerReward() (*big.Int, error) {
	return _FSM.Contract.MaxUpdateCallerReward(&_FSM.CallOpts)
}

// MaxUpdateCallerReward is a free data retrieval call binding the contract method 0x69dec276.
//
// Solidity: function maxUpdateCallerReward() view returns(uint256)
func (_FSM *FSMCallerSession) MaxUpdateCallerReward() (*big.Int, error) {
	return _FSM.Contract.MaxUpdateCallerReward(&_FSM.CallOpts)
}

// Minimum is a free data retrieval call binding the contract method 0xdd2d2a12.
//
// Solidity: function minimum(uint256 x, uint256 y) pure returns(uint256 z)
func (_FSM *FSMCaller) Minimum(opts *bind.CallOpts, x *big.Int, y *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "minimum", x, y)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Minimum is a free data retrieval call binding the contract method 0xdd2d2a12.
//
// Solidity: function minimum(uint256 x, uint256 y) pure returns(uint256 z)
func (_FSM *FSMSession) Minimum(x *big.Int, y *big.Int) (*big.Int, error) {
	return _FSM.Contract.Minimum(&_FSM.CallOpts, x, y)
}

// Minimum is a free data retrieval call binding the contract method 0xdd2d2a12.
//
// Solidity: function minimum(uint256 x, uint256 y) pure returns(uint256 z)
func (_FSM *FSMCallerSession) Minimum(x *big.Int, y *big.Int) (*big.Int, error) {
	return _FSM.Contract.Minimum(&_FSM.CallOpts, x, y)
}

// Multiply is a free data retrieval call binding the contract method 0x165c4a16.
//
// Solidity: function multiply(uint256 x, uint256 y) pure returns(uint256 z)
func (_FSM *FSMCaller) Multiply(opts *bind.CallOpts, x *big.Int, y *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "multiply", x, y)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Multiply is a free data retrieval call binding the contract method 0x165c4a16.
//
// Solidity: function multiply(uint256 x, uint256 y) pure returns(uint256 z)
func (_FSM *FSMSession) Multiply(x *big.Int, y *big.Int) (*big.Int, error) {
	return _FSM.Contract.Multiply(&_FSM.CallOpts, x, y)
}

// Multiply is a free data retrieval call binding the contract method 0x165c4a16.
//
// Solidity: function multiply(uint256 x, uint256 y) pure returns(uint256 z)
func (_FSM *FSMCallerSession) Multiply(x *big.Int, y *big.Int) (*big.Int, error) {
	return _FSM.Contract.Multiply(&_FSM.CallOpts, x, y)
}

// NewPriceDeviation is a free data retrieval call binding the contract method 0x7c720555.
//
// Solidity: function newPriceDeviation() view returns(uint256)
func (_FSM *FSMCaller) NewPriceDeviation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "newPriceDeviation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NewPriceDeviation is a free data retrieval call binding the contract method 0x7c720555.
//
// Solidity: function newPriceDeviation() view returns(uint256)
func (_FSM *FSMSession) NewPriceDeviation() (*big.Int, error) {
	return _FSM.Contract.NewPriceDeviation(&_FSM.CallOpts)
}

// NewPriceDeviation is a free data retrieval call binding the contract method 0x7c720555.
//
// Solidity: function newPriceDeviation() view returns(uint256)
func (_FSM *FSMCallerSession) NewPriceDeviation() (*big.Int, error) {
	return _FSM.Contract.NewPriceDeviation(&_FSM.CallOpts)
}

// PassedDelay is a free data retrieval call binding the contract method 0x63bfa88b.
//
// Solidity: function passedDelay() view returns(bool)
func (_FSM *FSMCaller) PassedDelay(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "passedDelay")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PassedDelay is a free data retrieval call binding the contract method 0x63bfa88b.
//
// Solidity: function passedDelay() view returns(bool)
func (_FSM *FSMSession) PassedDelay() (bool, error) {
	return _FSM.Contract.PassedDelay(&_FSM.CallOpts)
}

// PassedDelay is a free data retrieval call binding the contract method 0x63bfa88b.
//
// Solidity: function passedDelay() view returns(bool)
func (_FSM *FSMCallerSession) PassedDelay() (bool, error) {
	return _FSM.Contract.PassedDelay(&_FSM.CallOpts)
}

// PerSecondCallerRewardIncrease is a free data retrieval call binding the contract method 0x43943b6b.
//
// Solidity: function perSecondCallerRewardIncrease() view returns(uint256)
func (_FSM *FSMCaller) PerSecondCallerRewardIncrease(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "perSecondCallerRewardIncrease")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PerSecondCallerRewardIncrease is a free data retrieval call binding the contract method 0x43943b6b.
//
// Solidity: function perSecondCallerRewardIncrease() view returns(uint256)
func (_FSM *FSMSession) PerSecondCallerRewardIncrease() (*big.Int, error) {
	return _FSM.Contract.PerSecondCallerRewardIncrease(&_FSM.CallOpts)
}

// PerSecondCallerRewardIncrease is a free data retrieval call binding the contract method 0x43943b6b.
//
// Solidity: function perSecondCallerRewardIncrease() view returns(uint256)
func (_FSM *FSMCallerSession) PerSecondCallerRewardIncrease() (*big.Int, error) {
	return _FSM.Contract.PerSecondCallerRewardIncrease(&_FSM.CallOpts)
}

// PriceSource is a free data retrieval call binding the contract method 0x20531bc9.
//
// Solidity: function priceSource() view returns(address)
func (_FSM *FSMCaller) PriceSource(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "priceSource")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceSource is a free data retrieval call binding the contract method 0x20531bc9.
//
// Solidity: function priceSource() view returns(address)
func (_FSM *FSMSession) PriceSource() (common.Address, error) {
	return _FSM.Contract.PriceSource(&_FSM.CallOpts)
}

// PriceSource is a free data retrieval call binding the contract method 0x20531bc9.
//
// Solidity: function priceSource() view returns(address)
func (_FSM *FSMCallerSession) PriceSource() (common.Address, error) {
	return _FSM.Contract.PriceSource(&_FSM.CallOpts)
}

// Rad is a free data retrieval call binding the contract method 0x46f3e81c.
//
// Solidity: function rad(uint256 x) pure returns(uint256 z)
func (_FSM *FSMCaller) Rad(opts *bind.CallOpts, x *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "rad", x)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rad is a free data retrieval call binding the contract method 0x46f3e81c.
//
// Solidity: function rad(uint256 x) pure returns(uint256 z)
func (_FSM *FSMSession) Rad(x *big.Int) (*big.Int, error) {
	return _FSM.Contract.Rad(&_FSM.CallOpts, x)
}

// Rad is a free data retrieval call binding the contract method 0x46f3e81c.
//
// Solidity: function rad(uint256 x) pure returns(uint256 z)
func (_FSM *FSMCallerSession) Rad(x *big.Int) (*big.Int, error) {
	return _FSM.Contract.Rad(&_FSM.CallOpts, x)
}

// Ray is a free data retrieval call binding the contract method 0x10213447.
//
// Solidity: function ray(uint256 x) pure returns(uint256 z)
func (_FSM *FSMCaller) Ray(opts *bind.CallOpts, x *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "ray", x)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Ray is a free data retrieval call binding the contract method 0x10213447.
//
// Solidity: function ray(uint256 x) pure returns(uint256 z)
func (_FSM *FSMSession) Ray(x *big.Int) (*big.Int, error) {
	return _FSM.Contract.Ray(&_FSM.CallOpts, x)
}

// Ray is a free data retrieval call binding the contract method 0x10213447.
//
// Solidity: function ray(uint256 x) pure returns(uint256 z)
func (_FSM *FSMCallerSession) Ray(x *big.Int) (*big.Int, error) {
	return _FSM.Contract.Ray(&_FSM.CallOpts, x)
}

// Rdivide is a free data retrieval call binding the contract method 0xa0871637.
//
// Solidity: function rdivide(uint256 x, uint256 y) pure returns(uint256 z)
func (_FSM *FSMCaller) Rdivide(opts *bind.CallOpts, x *big.Int, y *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "rdivide", x, y)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rdivide is a free data retrieval call binding the contract method 0xa0871637.
//
// Solidity: function rdivide(uint256 x, uint256 y) pure returns(uint256 z)
func (_FSM *FSMSession) Rdivide(x *big.Int, y *big.Int) (*big.Int, error) {
	return _FSM.Contract.Rdivide(&_FSM.CallOpts, x, y)
}

// Rdivide is a free data retrieval call binding the contract method 0xa0871637.
//
// Solidity: function rdivide(uint256 x, uint256 y) pure returns(uint256 z)
func (_FSM *FSMCallerSession) Rdivide(x *big.Int, y *big.Int) (*big.Int, error) {
	return _FSM.Contract.Rdivide(&_FSM.CallOpts, x, y)
}

// Read is a free data retrieval call binding the contract method 0x57de26a4.
//
// Solidity: function read() view returns(uint256)
func (_FSM *FSMCaller) Read(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "read")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Read is a free data retrieval call binding the contract method 0x57de26a4.
//
// Solidity: function read() view returns(uint256)
func (_FSM *FSMSession) Read() (*big.Int, error) {
	return _FSM.Contract.Read(&_FSM.CallOpts)
}

// Read is a free data retrieval call binding the contract method 0x57de26a4.
//
// Solidity: function read() view returns(uint256)
func (_FSM *FSMCallerSession) Read() (*big.Int, error) {
	return _FSM.Contract.Read(&_FSM.CallOpts)
}

// ReimburseDelay is a free data retrieval call binding the contract method 0x4a5a9608.
//
// Solidity: function reimburseDelay() view returns(uint256)
func (_FSM *FSMCaller) ReimburseDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "reimburseDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReimburseDelay is a free data retrieval call binding the contract method 0x4a5a9608.
//
// Solidity: function reimburseDelay() view returns(uint256)
func (_FSM *FSMSession) ReimburseDelay() (*big.Int, error) {
	return _FSM.Contract.ReimburseDelay(&_FSM.CallOpts)
}

// ReimburseDelay is a free data retrieval call binding the contract method 0x4a5a9608.
//
// Solidity: function reimburseDelay() view returns(uint256)
func (_FSM *FSMCallerSession) ReimburseDelay() (*big.Int, error) {
	return _FSM.Contract.ReimburseDelay(&_FSM.CallOpts)
}

// Rmultiply is a free data retrieval call binding the contract method 0x056640b7.
//
// Solidity: function rmultiply(uint256 x, uint256 y) pure returns(uint256 z)
func (_FSM *FSMCaller) Rmultiply(opts *bind.CallOpts, x *big.Int, y *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "rmultiply", x, y)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rmultiply is a free data retrieval call binding the contract method 0x056640b7.
//
// Solidity: function rmultiply(uint256 x, uint256 y) pure returns(uint256 z)
func (_FSM *FSMSession) Rmultiply(x *big.Int, y *big.Int) (*big.Int, error) {
	return _FSM.Contract.Rmultiply(&_FSM.CallOpts, x, y)
}

// Rmultiply is a free data retrieval call binding the contract method 0x056640b7.
//
// Solidity: function rmultiply(uint256 x, uint256 y) pure returns(uint256 z)
func (_FSM *FSMCallerSession) Rmultiply(x *big.Int, y *big.Int) (*big.Int, error) {
	return _FSM.Contract.Rmultiply(&_FSM.CallOpts, x, y)
}

// Rpower is a free data retrieval call binding the contract method 0xd6e882dc.
//
// Solidity: function rpower(uint256 x, uint256 n, uint256 base) pure returns(uint256 z)
func (_FSM *FSMCaller) Rpower(opts *bind.CallOpts, x *big.Int, n *big.Int, base *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "rpower", x, n, base)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rpower is a free data retrieval call binding the contract method 0xd6e882dc.
//
// Solidity: function rpower(uint256 x, uint256 n, uint256 base) pure returns(uint256 z)
func (_FSM *FSMSession) Rpower(x *big.Int, n *big.Int, base *big.Int) (*big.Int, error) {
	return _FSM.Contract.Rpower(&_FSM.CallOpts, x, n, base)
}

// Rpower is a free data retrieval call binding the contract method 0xd6e882dc.
//
// Solidity: function rpower(uint256 x, uint256 n, uint256 base) pure returns(uint256 z)
func (_FSM *FSMCallerSession) Rpower(x *big.Int, n *big.Int, base *big.Int) (*big.Int, error) {
	return _FSM.Contract.Rpower(&_FSM.CallOpts, x, n, base)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() view returns(uint256)
func (_FSM *FSMCaller) Stopped(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "stopped")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() view returns(uint256)
func (_FSM *FSMSession) Stopped() (*big.Int, error) {
	return _FSM.Contract.Stopped(&_FSM.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() view returns(uint256)
func (_FSM *FSMCallerSession) Stopped() (*big.Int, error) {
	return _FSM.Contract.Stopped(&_FSM.CallOpts)
}

// Subtract is a free data retrieval call binding the contract method 0x3ef5e445.
//
// Solidity: function subtract(uint256 x, uint256 y) pure returns(uint256 z)
func (_FSM *FSMCaller) Subtract(opts *bind.CallOpts, x *big.Int, y *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "subtract", x, y)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Subtract is a free data retrieval call binding the contract method 0x3ef5e445.
//
// Solidity: function subtract(uint256 x, uint256 y) pure returns(uint256 z)
func (_FSM *FSMSession) Subtract(x *big.Int, y *big.Int) (*big.Int, error) {
	return _FSM.Contract.Subtract(&_FSM.CallOpts, x, y)
}

// Subtract is a free data retrieval call binding the contract method 0x3ef5e445.
//
// Solidity: function subtract(uint256 x, uint256 y) pure returns(uint256 z)
func (_FSM *FSMCallerSession) Subtract(x *big.Int, y *big.Int) (*big.Int, error) {
	return _FSM.Contract.Subtract(&_FSM.CallOpts, x, y)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_FSM *FSMCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_FSM *FSMSession) Treasury() (common.Address, error) {
	return _FSM.Contract.Treasury(&_FSM.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_FSM *FSMCallerSession) Treasury() (common.Address, error) {
	return _FSM.Contract.Treasury(&_FSM.CallOpts)
}

// TreasuryAllowance is a free data retrieval call binding the contract method 0x3425677e.
//
// Solidity: function treasuryAllowance() view returns(uint256)
func (_FSM *FSMCaller) TreasuryAllowance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "treasuryAllowance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreasuryAllowance is a free data retrieval call binding the contract method 0x3425677e.
//
// Solidity: function treasuryAllowance() view returns(uint256)
func (_FSM *FSMSession) TreasuryAllowance() (*big.Int, error) {
	return _FSM.Contract.TreasuryAllowance(&_FSM.CallOpts)
}

// TreasuryAllowance is a free data retrieval call binding the contract method 0x3425677e.
//
// Solidity: function treasuryAllowance() view returns(uint256)
func (_FSM *FSMCallerSession) TreasuryAllowance() (*big.Int, error) {
	return _FSM.Contract.TreasuryAllowance(&_FSM.CallOpts)
}

// UpdateDelay is a free data retrieval call binding the contract method 0x554f94db.
//
// Solidity: function updateDelay() view returns(uint16)
func (_FSM *FSMCaller) UpdateDelay(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "updateDelay")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// UpdateDelay is a free data retrieval call binding the contract method 0x554f94db.
//
// Solidity: function updateDelay() view returns(uint16)
func (_FSM *FSMSession) UpdateDelay() (uint16, error) {
	return _FSM.Contract.UpdateDelay(&_FSM.CallOpts)
}

// UpdateDelay is a free data retrieval call binding the contract method 0x554f94db.
//
// Solidity: function updateDelay() view returns(uint16)
func (_FSM *FSMCallerSession) UpdateDelay() (uint16, error) {
	return _FSM.Contract.UpdateDelay(&_FSM.CallOpts)
}

// Wdivide is a free data retrieval call binding the contract method 0xf752fdc3.
//
// Solidity: function wdivide(uint256 x, uint256 y) pure returns(uint256 z)
func (_FSM *FSMCaller) Wdivide(opts *bind.CallOpts, x *big.Int, y *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "wdivide", x, y)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wdivide is a free data retrieval call binding the contract method 0xf752fdc3.
//
// Solidity: function wdivide(uint256 x, uint256 y) pure returns(uint256 z)
func (_FSM *FSMSession) Wdivide(x *big.Int, y *big.Int) (*big.Int, error) {
	return _FSM.Contract.Wdivide(&_FSM.CallOpts, x, y)
}

// Wdivide is a free data retrieval call binding the contract method 0xf752fdc3.
//
// Solidity: function wdivide(uint256 x, uint256 y) pure returns(uint256 z)
func (_FSM *FSMCallerSession) Wdivide(x *big.Int, y *big.Int) (*big.Int, error) {
	return _FSM.Contract.Wdivide(&_FSM.CallOpts, x, y)
}

// Wmultiply is a free data retrieval call binding the contract method 0x3c8bb3e6.
//
// Solidity: function wmultiply(uint256 x, uint256 y) pure returns(uint256 z)
func (_FSM *FSMCaller) Wmultiply(opts *bind.CallOpts, x *big.Int, y *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FSM.contract.Call(opts, &out, "wmultiply", x, y)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wmultiply is a free data retrieval call binding the contract method 0x3c8bb3e6.
//
// Solidity: function wmultiply(uint256 x, uint256 y) pure returns(uint256 z)
func (_FSM *FSMSession) Wmultiply(x *big.Int, y *big.Int) (*big.Int, error) {
	return _FSM.Contract.Wmultiply(&_FSM.CallOpts, x, y)
}

// Wmultiply is a free data retrieval call binding the contract method 0x3c8bb3e6.
//
// Solidity: function wmultiply(uint256 x, uint256 y) pure returns(uint256 z)
func (_FSM *FSMCallerSession) Wmultiply(x *big.Int, y *big.Int) (*big.Int, error) {
	return _FSM.Contract.Wmultiply(&_FSM.CallOpts, x, y)
}

// AddAuthorization is a paid mutator transaction binding the contract method 0x35b28153.
//
// Solidity: function addAuthorization(address account) returns()
func (_FSM *FSMTransactor) AddAuthorization(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _FSM.contract.Transact(opts, "addAuthorization", account)
}

// AddAuthorization is a paid mutator transaction binding the contract method 0x35b28153.
//
// Solidity: function addAuthorization(address account) returns()
func (_FSM *FSMSession) AddAuthorization(account common.Address) (*types.Transaction, error) {
	return _FSM.Contract.AddAuthorization(&_FSM.TransactOpts, account)
}

// AddAuthorization is a paid mutator transaction binding the contract method 0x35b28153.
//
// Solidity: function addAuthorization(address account) returns()
func (_FSM *FSMTransactorSession) AddAuthorization(account common.Address) (*types.Transaction, error) {
	return _FSM.Contract.AddAuthorization(&_FSM.TransactOpts, account)
}

// ModifyParameters is a paid mutator transaction binding the contract method 0x6614f010.
//
// Solidity: function modifyParameters(bytes32 parameter, address addr) returns()
func (_FSM *FSMTransactor) ModifyParameters(opts *bind.TransactOpts, parameter [32]byte, addr common.Address) (*types.Transaction, error) {
	return _FSM.contract.Transact(opts, "modifyParameters", parameter, addr)
}

// ModifyParameters is a paid mutator transaction binding the contract method 0x6614f010.
//
// Solidity: function modifyParameters(bytes32 parameter, address addr) returns()
func (_FSM *FSMSession) ModifyParameters(parameter [32]byte, addr common.Address) (*types.Transaction, error) {
	return _FSM.Contract.ModifyParameters(&_FSM.TransactOpts, parameter, addr)
}

// ModifyParameters is a paid mutator transaction binding the contract method 0x6614f010.
//
// Solidity: function modifyParameters(bytes32 parameter, address addr) returns()
func (_FSM *FSMTransactorSession) ModifyParameters(parameter [32]byte, addr common.Address) (*types.Transaction, error) {
	return _FSM.Contract.ModifyParameters(&_FSM.TransactOpts, parameter, addr)
}

// ModifyParameters0 is a paid mutator transaction binding the contract method 0xfe4f5890.
//
// Solidity: function modifyParameters(bytes32 parameter, uint256 val) returns()
func (_FSM *FSMTransactor) ModifyParameters0(opts *bind.TransactOpts, parameter [32]byte, val *big.Int) (*types.Transaction, error) {
	return _FSM.contract.Transact(opts, "modifyParameters0", parameter, val)
}

// ModifyParameters0 is a paid mutator transaction binding the contract method 0xfe4f5890.
//
// Solidity: function modifyParameters(bytes32 parameter, uint256 val) returns()
func (_FSM *FSMSession) ModifyParameters0(parameter [32]byte, val *big.Int) (*types.Transaction, error) {
	return _FSM.Contract.ModifyParameters0(&_FSM.TransactOpts, parameter, val)
}

// ModifyParameters0 is a paid mutator transaction binding the contract method 0xfe4f5890.
//
// Solidity: function modifyParameters(bytes32 parameter, uint256 val) returns()
func (_FSM *FSMTransactorSession) ModifyParameters0(parameter [32]byte, val *big.Int) (*types.Transaction, error) {
	return _FSM.Contract.ModifyParameters0(&_FSM.TransactOpts, parameter, val)
}

// RemoveAuthorization is a paid mutator transaction binding the contract method 0x94f3f81d.
//
// Solidity: function removeAuthorization(address account) returns()
func (_FSM *FSMTransactor) RemoveAuthorization(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _FSM.contract.Transact(opts, "removeAuthorization", account)
}

// RemoveAuthorization is a paid mutator transaction binding the contract method 0x94f3f81d.
//
// Solidity: function removeAuthorization(address account) returns()
func (_FSM *FSMSession) RemoveAuthorization(account common.Address) (*types.Transaction, error) {
	return _FSM.Contract.RemoveAuthorization(&_FSM.TransactOpts, account)
}

// RemoveAuthorization is a paid mutator transaction binding the contract method 0x94f3f81d.
//
// Solidity: function removeAuthorization(address account) returns()
func (_FSM *FSMTransactorSession) RemoveAuthorization(account common.Address) (*types.Transaction, error) {
	return _FSM.Contract.RemoveAuthorization(&_FSM.TransactOpts, account)
}

// RenumerateCaller is a paid mutator transaction binding the contract method 0x2761f27b.
//
// Solidity: function renumerateCaller(address feeReceiver) returns()
func (_FSM *FSMTransactor) RenumerateCaller(opts *bind.TransactOpts, feeReceiver common.Address) (*types.Transaction, error) {
	return _FSM.contract.Transact(opts, "renumerateCaller", feeReceiver)
}

// RenumerateCaller is a paid mutator transaction binding the contract method 0x2761f27b.
//
// Solidity: function renumerateCaller(address feeReceiver) returns()
func (_FSM *FSMSession) RenumerateCaller(feeReceiver common.Address) (*types.Transaction, error) {
	return _FSM.Contract.RenumerateCaller(&_FSM.TransactOpts, feeReceiver)
}

// RenumerateCaller is a paid mutator transaction binding the contract method 0x2761f27b.
//
// Solidity: function renumerateCaller(address feeReceiver) returns()
func (_FSM *FSMTransactorSession) RenumerateCaller(feeReceiver common.Address) (*types.Transaction, error) {
	return _FSM.Contract.RenumerateCaller(&_FSM.TransactOpts, feeReceiver)
}

// FSMAddAuthorizationIterator is returned from FilterAddAuthorization and is used to iterate over the raw logs and unpacked data for AddAuthorization events raised by the FSM contract.
type FSMAddAuthorizationIterator struct {
	Event *FSMAddAuthorization // Event containing the contract specifics and raw log

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
func (it *FSMAddAuthorizationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FSMAddAuthorization)
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
		it.Event = new(FSMAddAuthorization)
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
func (it *FSMAddAuthorizationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FSMAddAuthorizationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FSMAddAuthorization represents a AddAuthorization event raised by the FSM contract.
type FSMAddAuthorization struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddAuthorization is a free log retrieval operation binding the contract event 0x599a298163e1678bb1c676052a8930bf0b8a1261ed6e01b8a2391e55f7000102.
//
// Solidity: event AddAuthorization(address account)
func (_FSM *FSMFilterer) FilterAddAuthorization(opts *bind.FilterOpts) (*FSMAddAuthorizationIterator, error) {

	logs, sub, err := _FSM.contract.FilterLogs(opts, "AddAuthorization")
	if err != nil {
		return nil, err
	}
	return &FSMAddAuthorizationIterator{contract: _FSM.contract, event: "AddAuthorization", logs: logs, sub: sub}, nil
}

// WatchAddAuthorization is a free log subscription operation binding the contract event 0x599a298163e1678bb1c676052a8930bf0b8a1261ed6e01b8a2391e55f7000102.
//
// Solidity: event AddAuthorization(address account)
func (_FSM *FSMFilterer) WatchAddAuthorization(opts *bind.WatchOpts, sink chan<- *FSMAddAuthorization) (event.Subscription, error) {

	logs, sub, err := _FSM.contract.WatchLogs(opts, "AddAuthorization")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FSMAddAuthorization)
				if err := _FSM.contract.UnpackLog(event, "AddAuthorization", log); err != nil {
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
func (_FSM *FSMFilterer) ParseAddAuthorization(log types.Log) (*FSMAddAuthorization, error) {
	event := new(FSMAddAuthorization)
	if err := _FSM.contract.UnpackLog(event, "AddAuthorization", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FSMFailRewardCallerIterator is returned from FilterFailRewardCaller and is used to iterate over the raw logs and unpacked data for FailRewardCaller events raised by the FSM contract.
type FSMFailRewardCallerIterator struct {
	Event *FSMFailRewardCaller // Event containing the contract specifics and raw log

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
func (it *FSMFailRewardCallerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FSMFailRewardCaller)
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
		it.Event = new(FSMFailRewardCaller)
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
func (it *FSMFailRewardCallerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FSMFailRewardCallerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FSMFailRewardCaller represents a FailRewardCaller event raised by the FSM contract.
type FSMFailRewardCaller struct {
	RevertReason []byte
	FeeReceiver  common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFailRewardCaller is a free log retrieval operation binding the contract event 0xf7bf1f7447ce563690edb2abe40636178ff64fc766b07bf3e171b16102794a54.
//
// Solidity: event FailRewardCaller(bytes revertReason, address feeReceiver, uint256 amount)
func (_FSM *FSMFilterer) FilterFailRewardCaller(opts *bind.FilterOpts) (*FSMFailRewardCallerIterator, error) {

	logs, sub, err := _FSM.contract.FilterLogs(opts, "FailRewardCaller")
	if err != nil {
		return nil, err
	}
	return &FSMFailRewardCallerIterator{contract: _FSM.contract, event: "FailRewardCaller", logs: logs, sub: sub}, nil
}

// WatchFailRewardCaller is a free log subscription operation binding the contract event 0xf7bf1f7447ce563690edb2abe40636178ff64fc766b07bf3e171b16102794a54.
//
// Solidity: event FailRewardCaller(bytes revertReason, address feeReceiver, uint256 amount)
func (_FSM *FSMFilterer) WatchFailRewardCaller(opts *bind.WatchOpts, sink chan<- *FSMFailRewardCaller) (event.Subscription, error) {

	logs, sub, err := _FSM.contract.WatchLogs(opts, "FailRewardCaller")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FSMFailRewardCaller)
				if err := _FSM.contract.UnpackLog(event, "FailRewardCaller", log); err != nil {
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

// ParseFailRewardCaller is a log parse operation binding the contract event 0xf7bf1f7447ce563690edb2abe40636178ff64fc766b07bf3e171b16102794a54.
//
// Solidity: event FailRewardCaller(bytes revertReason, address feeReceiver, uint256 amount)
func (_FSM *FSMFilterer) ParseFailRewardCaller(log types.Log) (*FSMFailRewardCaller, error) {
	event := new(FSMFailRewardCaller)
	if err := _FSM.contract.UnpackLog(event, "FailRewardCaller", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FSMModifyParametersIterator is returned from FilterModifyParameters and is used to iterate over the raw logs and unpacked data for ModifyParameters events raised by the FSM contract.
type FSMModifyParametersIterator struct {
	Event *FSMModifyParameters // Event containing the contract specifics and raw log

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
func (it *FSMModifyParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FSMModifyParameters)
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
		it.Event = new(FSMModifyParameters)
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
func (it *FSMModifyParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FSMModifyParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FSMModifyParameters represents a ModifyParameters event raised by the FSM contract.
type FSMModifyParameters struct {
	Parameter [32]byte
	Addr      common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterModifyParameters is a free log retrieval operation binding the contract event 0xd91f38cf03346b5dc15fb60f9076f866295231ad3c3841a1051f8443f25170d1.
//
// Solidity: event ModifyParameters(bytes32 parameter, address addr)
func (_FSM *FSMFilterer) FilterModifyParameters(opts *bind.FilterOpts) (*FSMModifyParametersIterator, error) {

	logs, sub, err := _FSM.contract.FilterLogs(opts, "ModifyParameters")
	if err != nil {
		return nil, err
	}
	return &FSMModifyParametersIterator{contract: _FSM.contract, event: "ModifyParameters", logs: logs, sub: sub}, nil
}

// WatchModifyParameters is a free log subscription operation binding the contract event 0xd91f38cf03346b5dc15fb60f9076f866295231ad3c3841a1051f8443f25170d1.
//
// Solidity: event ModifyParameters(bytes32 parameter, address addr)
func (_FSM *FSMFilterer) WatchModifyParameters(opts *bind.WatchOpts, sink chan<- *FSMModifyParameters) (event.Subscription, error) {

	logs, sub, err := _FSM.contract.WatchLogs(opts, "ModifyParameters")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FSMModifyParameters)
				if err := _FSM.contract.UnpackLog(event, "ModifyParameters", log); err != nil {
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

// ParseModifyParameters is a log parse operation binding the contract event 0xd91f38cf03346b5dc15fb60f9076f866295231ad3c3841a1051f8443f25170d1.
//
// Solidity: event ModifyParameters(bytes32 parameter, address addr)
func (_FSM *FSMFilterer) ParseModifyParameters(log types.Log) (*FSMModifyParameters, error) {
	event := new(FSMModifyParameters)
	if err := _FSM.contract.UnpackLog(event, "ModifyParameters", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FSMModifyParameters0Iterator is returned from FilterModifyParameters0 and is used to iterate over the raw logs and unpacked data for ModifyParameters0 events raised by the FSM contract.
type FSMModifyParameters0Iterator struct {
	Event *FSMModifyParameters0 // Event containing the contract specifics and raw log

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
func (it *FSMModifyParameters0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FSMModifyParameters0)
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
		it.Event = new(FSMModifyParameters0)
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
func (it *FSMModifyParameters0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FSMModifyParameters0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FSMModifyParameters0 represents a ModifyParameters0 event raised by the FSM contract.
type FSMModifyParameters0 struct {
	Parameter [32]byte
	Val       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterModifyParameters0 is a free log retrieval operation binding the contract event 0xac7c5c1afaef770ec56ac6268cd3f2fbb1035858ead2601d6553157c33036c3a.
//
// Solidity: event ModifyParameters(bytes32 parameter, uint256 val)
func (_FSM *FSMFilterer) FilterModifyParameters0(opts *bind.FilterOpts) (*FSMModifyParameters0Iterator, error) {

	logs, sub, err := _FSM.contract.FilterLogs(opts, "ModifyParameters0")
	if err != nil {
		return nil, err
	}
	return &FSMModifyParameters0Iterator{contract: _FSM.contract, event: "ModifyParameters0", logs: logs, sub: sub}, nil
}

// WatchModifyParameters0 is a free log subscription operation binding the contract event 0xac7c5c1afaef770ec56ac6268cd3f2fbb1035858ead2601d6553157c33036c3a.
//
// Solidity: event ModifyParameters(bytes32 parameter, uint256 val)
func (_FSM *FSMFilterer) WatchModifyParameters0(opts *bind.WatchOpts, sink chan<- *FSMModifyParameters0) (event.Subscription, error) {

	logs, sub, err := _FSM.contract.WatchLogs(opts, "ModifyParameters0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FSMModifyParameters0)
				if err := _FSM.contract.UnpackLog(event, "ModifyParameters0", log); err != nil {
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

// ParseModifyParameters0 is a log parse operation binding the contract event 0xac7c5c1afaef770ec56ac6268cd3f2fbb1035858ead2601d6553157c33036c3a.
//
// Solidity: event ModifyParameters(bytes32 parameter, uint256 val)
func (_FSM *FSMFilterer) ParseModifyParameters0(log types.Log) (*FSMModifyParameters0, error) {
	event := new(FSMModifyParameters0)
	if err := _FSM.contract.UnpackLog(event, "ModifyParameters0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FSMRemoveAuthorizationIterator is returned from FilterRemoveAuthorization and is used to iterate over the raw logs and unpacked data for RemoveAuthorization events raised by the FSM contract.
type FSMRemoveAuthorizationIterator struct {
	Event *FSMRemoveAuthorization // Event containing the contract specifics and raw log

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
func (it *FSMRemoveAuthorizationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FSMRemoveAuthorization)
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
		it.Event = new(FSMRemoveAuthorization)
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
func (it *FSMRemoveAuthorizationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FSMRemoveAuthorizationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FSMRemoveAuthorization represents a RemoveAuthorization event raised by the FSM contract.
type FSMRemoveAuthorization struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRemoveAuthorization is a free log retrieval operation binding the contract event 0x8834a87e641e9716be4f34527af5d23e11624f1ddeefede6ad75a9acfc31b903.
//
// Solidity: event RemoveAuthorization(address account)
func (_FSM *FSMFilterer) FilterRemoveAuthorization(opts *bind.FilterOpts) (*FSMRemoveAuthorizationIterator, error) {

	logs, sub, err := _FSM.contract.FilterLogs(opts, "RemoveAuthorization")
	if err != nil {
		return nil, err
	}
	return &FSMRemoveAuthorizationIterator{contract: _FSM.contract, event: "RemoveAuthorization", logs: logs, sub: sub}, nil
}

// WatchRemoveAuthorization is a free log subscription operation binding the contract event 0x8834a87e641e9716be4f34527af5d23e11624f1ddeefede6ad75a9acfc31b903.
//
// Solidity: event RemoveAuthorization(address account)
func (_FSM *FSMFilterer) WatchRemoveAuthorization(opts *bind.WatchOpts, sink chan<- *FSMRemoveAuthorization) (event.Subscription, error) {

	logs, sub, err := _FSM.contract.WatchLogs(opts, "RemoveAuthorization")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FSMRemoveAuthorization)
				if err := _FSM.contract.UnpackLog(event, "RemoveAuthorization", log); err != nil {
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
func (_FSM *FSMFilterer) ParseRemoveAuthorization(log types.Log) (*FSMRemoveAuthorization, error) {
	event := new(FSMRemoveAuthorization)
	if err := _FSM.contract.UnpackLog(event, "RemoveAuthorization", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
