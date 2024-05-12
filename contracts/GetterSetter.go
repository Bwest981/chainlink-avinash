// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// GetterSetterMetaData contains all meta data concerning the GetterSetter contract.
var GetterSetterMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"b32\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"u256\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"b322\",\"type\":\"bytes32\"}],\"name\":\"Output\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"SetBytes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"SetBytes32\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SetUint256\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBytes32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUint256\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"}],\"name\":\"requestedBytes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_value\",\"type\":\"bytes32\"}],\"name\":\"requestedBytes32\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"requestedUint256\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"}],\"name\":\"setBytes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_value\",\"type\":\"bytes32\"}],\"name\":\"setBytes32\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setUint256\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// GetterSetterABI is the input ABI used to generate the binding from.
// Deprecated: Use GetterSetterMetaData.ABI instead.
var GetterSetterABI = GetterSetterMetaData.ABI

// GetterSetter is an auto generated Go binding around an Ethereum contract.
type GetterSetter struct {
	GetterSetterCaller     // Read-only binding to the contract
	GetterSetterTransactor // Write-only binding to the contract
	GetterSetterFilterer   // Log filterer for contract events
}

// GetterSetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type GetterSetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GetterSetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GetterSetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GetterSetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GetterSetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GetterSetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GetterSetterSession struct {
	Contract     *GetterSetter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GetterSetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GetterSetterCallerSession struct {
	Contract *GetterSetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// GetterSetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GetterSetterTransactorSession struct {
	Contract     *GetterSetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// GetterSetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type GetterSetterRaw struct {
	Contract *GetterSetter // Generic contract binding to access the raw methods on
}

// GetterSetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GetterSetterCallerRaw struct {
	Contract *GetterSetterCaller // Generic read-only contract binding to access the raw methods on
}

// GetterSetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GetterSetterTransactorRaw struct {
	Contract *GetterSetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGetterSetter creates a new instance of GetterSetter, bound to a specific deployed contract.
func NewGetterSetter(address common.Address, backend bind.ContractBackend) (*GetterSetter, error) {
	contract, err := bindGetterSetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GetterSetter{GetterSetterCaller: GetterSetterCaller{contract: contract}, GetterSetterTransactor: GetterSetterTransactor{contract: contract}, GetterSetterFilterer: GetterSetterFilterer{contract: contract}}, nil
}

// NewGetterSetterCaller creates a new read-only instance of GetterSetter, bound to a specific deployed contract.
func NewGetterSetterCaller(address common.Address, caller bind.ContractCaller) (*GetterSetterCaller, error) {
	contract, err := bindGetterSetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GetterSetterCaller{contract: contract}, nil
}

// NewGetterSetterTransactor creates a new write-only instance of GetterSetter, bound to a specific deployed contract.
func NewGetterSetterTransactor(address common.Address, transactor bind.ContractTransactor) (*GetterSetterTransactor, error) {
	contract, err := bindGetterSetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GetterSetterTransactor{contract: contract}, nil
}

// NewGetterSetterFilterer creates a new log filterer instance of GetterSetter, bound to a specific deployed contract.
func NewGetterSetterFilterer(address common.Address, filterer bind.ContractFilterer) (*GetterSetterFilterer, error) {
	contract, err := bindGetterSetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GetterSetterFilterer{contract: contract}, nil
}

// bindGetterSetter binds a generic wrapper to an already deployed contract.
func bindGetterSetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GetterSetterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GetterSetter *GetterSetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GetterSetter.Contract.GetterSetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GetterSetter *GetterSetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GetterSetter.Contract.GetterSetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GetterSetter *GetterSetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GetterSetter.Contract.GetterSetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GetterSetter *GetterSetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GetterSetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GetterSetter *GetterSetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GetterSetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GetterSetter *GetterSetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GetterSetter.Contract.contract.Transact(opts, method, params...)
}

// GetBytes is a free data retrieval call binding the contract method 0x0bcd3b33.
//
// Solidity: function getBytes() view returns(bytes)
func (_GetterSetter *GetterSetterCaller) GetBytes(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _GetterSetter.contract.Call(opts, &out, "getBytes")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetBytes is a free data retrieval call binding the contract method 0x0bcd3b33.
//
// Solidity: function getBytes() view returns(bytes)
func (_GetterSetter *GetterSetterSession) GetBytes() ([]byte, error) {
	return _GetterSetter.Contract.GetBytes(&_GetterSetter.CallOpts)
}

// GetBytes is a free data retrieval call binding the contract method 0x0bcd3b33.
//
// Solidity: function getBytes() view returns(bytes)
func (_GetterSetter *GetterSetterCallerSession) GetBytes() ([]byte, error) {
	return _GetterSetter.Contract.GetBytes(&_GetterSetter.CallOpts)
}

// GetBytes32 is a free data retrieval call binding the contract method 0x1f903037.
//
// Solidity: function getBytes32() view returns(bytes32)
func (_GetterSetter *GetterSetterCaller) GetBytes32(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GetterSetter.contract.Call(opts, &out, "getBytes32")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBytes32 is a free data retrieval call binding the contract method 0x1f903037.
//
// Solidity: function getBytes32() view returns(bytes32)
func (_GetterSetter *GetterSetterSession) GetBytes32() ([32]byte, error) {
	return _GetterSetter.Contract.GetBytes32(&_GetterSetter.CallOpts)
}

// GetBytes32 is a free data retrieval call binding the contract method 0x1f903037.
//
// Solidity: function getBytes32() view returns(bytes32)
func (_GetterSetter *GetterSetterCallerSession) GetBytes32() ([32]byte, error) {
	return _GetterSetter.Contract.GetBytes32(&_GetterSetter.CallOpts)
}

// GetUint256 is a free data retrieval call binding the contract method 0x68895979.
//
// Solidity: function getUint256() view returns(uint256)
func (_GetterSetter *GetterSetterCaller) GetUint256(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GetterSetter.contract.Call(opts, &out, "getUint256")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUint256 is a free data retrieval call binding the contract method 0x68895979.
//
// Solidity: function getUint256() view returns(uint256)
func (_GetterSetter *GetterSetterSession) GetUint256() (*big.Int, error) {
	return _GetterSetter.Contract.GetUint256(&_GetterSetter.CallOpts)
}

// GetUint256 is a free data retrieval call binding the contract method 0x68895979.
//
// Solidity: function getUint256() view returns(uint256)
func (_GetterSetter *GetterSetterCallerSession) GetUint256() (*big.Int, error) {
	return _GetterSetter.Contract.GetUint256(&_GetterSetter.CallOpts)
}

// RequestId is a free data retrieval call binding the contract method 0x006d6cae.
//
// Solidity: function requestId() view returns(bytes32)
func (_GetterSetter *GetterSetterCaller) RequestId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GetterSetter.contract.Call(opts, &out, "requestId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RequestId is a free data retrieval call binding the contract method 0x006d6cae.
//
// Solidity: function requestId() view returns(bytes32)
func (_GetterSetter *GetterSetterSession) RequestId() ([32]byte, error) {
	return _GetterSetter.Contract.RequestId(&_GetterSetter.CallOpts)
}

// RequestId is a free data retrieval call binding the contract method 0x006d6cae.
//
// Solidity: function requestId() view returns(bytes32)
func (_GetterSetter *GetterSetterCallerSession) RequestId() ([32]byte, error) {
	return _GetterSetter.Contract.RequestId(&_GetterSetter.CallOpts)
}

// RequestedBytes is a paid mutator transaction binding the contract method 0x46ddd1ff.
//
// Solidity: function requestedBytes(bytes32 _requestId, bytes _value) returns()
func (_GetterSetter *GetterSetterTransactor) RequestedBytes(opts *bind.TransactOpts, _requestId [32]byte, _value []byte) (*types.Transaction, error) {
	return _GetterSetter.contract.Transact(opts, "requestedBytes", _requestId, _value)
}

// RequestedBytes is a paid mutator transaction binding the contract method 0x46ddd1ff.
//
// Solidity: function requestedBytes(bytes32 _requestId, bytes _value) returns()
func (_GetterSetter *GetterSetterSession) RequestedBytes(_requestId [32]byte, _value []byte) (*types.Transaction, error) {
	return _GetterSetter.Contract.RequestedBytes(&_GetterSetter.TransactOpts, _requestId, _value)
}

// RequestedBytes is a paid mutator transaction binding the contract method 0x46ddd1ff.
//
// Solidity: function requestedBytes(bytes32 _requestId, bytes _value) returns()
func (_GetterSetter *GetterSetterTransactorSession) RequestedBytes(_requestId [32]byte, _value []byte) (*types.Transaction, error) {
	return _GetterSetter.Contract.RequestedBytes(&_GetterSetter.TransactOpts, _requestId, _value)
}

// RequestedBytes32 is a paid mutator transaction binding the contract method 0xed53e511.
//
// Solidity: function requestedBytes32(bytes32 _requestId, bytes32 _value) returns()
func (_GetterSetter *GetterSetterTransactor) RequestedBytes32(opts *bind.TransactOpts, _requestId [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _GetterSetter.contract.Transact(opts, "requestedBytes32", _requestId, _value)
}

// RequestedBytes32 is a paid mutator transaction binding the contract method 0xed53e511.
//
// Solidity: function requestedBytes32(bytes32 _requestId, bytes32 _value) returns()
func (_GetterSetter *GetterSetterSession) RequestedBytes32(_requestId [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _GetterSetter.Contract.RequestedBytes32(&_GetterSetter.TransactOpts, _requestId, _value)
}

// RequestedBytes32 is a paid mutator transaction binding the contract method 0xed53e511.
//
// Solidity: function requestedBytes32(bytes32 _requestId, bytes32 _value) returns()
func (_GetterSetter *GetterSetterTransactorSession) RequestedBytes32(_requestId [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _GetterSetter.Contract.RequestedBytes32(&_GetterSetter.TransactOpts, _requestId, _value)
}

// RequestedUint256 is a paid mutator transaction binding the contract method 0x3345b4d0.
//
// Solidity: function requestedUint256(bytes32 _requestId, uint256 _value) returns()
func (_GetterSetter *GetterSetterTransactor) RequestedUint256(opts *bind.TransactOpts, _requestId [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _GetterSetter.contract.Transact(opts, "requestedUint256", _requestId, _value)
}

// RequestedUint256 is a paid mutator transaction binding the contract method 0x3345b4d0.
//
// Solidity: function requestedUint256(bytes32 _requestId, uint256 _value) returns()
func (_GetterSetter *GetterSetterSession) RequestedUint256(_requestId [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _GetterSetter.Contract.RequestedUint256(&_GetterSetter.TransactOpts, _requestId, _value)
}

// RequestedUint256 is a paid mutator transaction binding the contract method 0x3345b4d0.
//
// Solidity: function requestedUint256(bytes32 _requestId, uint256 _value) returns()
func (_GetterSetter *GetterSetterTransactorSession) RequestedUint256(_requestId [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _GetterSetter.Contract.RequestedUint256(&_GetterSetter.TransactOpts, _requestId, _value)
}

// SetBytes is a paid mutator transaction binding the contract method 0xda359dc8.
//
// Solidity: function setBytes(bytes _value) returns()
func (_GetterSetter *GetterSetterTransactor) SetBytes(opts *bind.TransactOpts, _value []byte) (*types.Transaction, error) {
	return _GetterSetter.contract.Transact(opts, "setBytes", _value)
}

// SetBytes is a paid mutator transaction binding the contract method 0xda359dc8.
//
// Solidity: function setBytes(bytes _value) returns()
func (_GetterSetter *GetterSetterSession) SetBytes(_value []byte) (*types.Transaction, error) {
	return _GetterSetter.Contract.SetBytes(&_GetterSetter.TransactOpts, _value)
}

// SetBytes is a paid mutator transaction binding the contract method 0xda359dc8.
//
// Solidity: function setBytes(bytes _value) returns()
func (_GetterSetter *GetterSetterTransactorSession) SetBytes(_value []byte) (*types.Transaction, error) {
	return _GetterSetter.Contract.SetBytes(&_GetterSetter.TransactOpts, _value)
}

// SetBytes32 is a paid mutator transaction binding the contract method 0xc2b12a73.
//
// Solidity: function setBytes32(bytes32 _value) returns()
func (_GetterSetter *GetterSetterTransactor) SetBytes32(opts *bind.TransactOpts, _value [32]byte) (*types.Transaction, error) {
	return _GetterSetter.contract.Transact(opts, "setBytes32", _value)
}

// SetBytes32 is a paid mutator transaction binding the contract method 0xc2b12a73.
//
// Solidity: function setBytes32(bytes32 _value) returns()
func (_GetterSetter *GetterSetterSession) SetBytes32(_value [32]byte) (*types.Transaction, error) {
	return _GetterSetter.Contract.SetBytes32(&_GetterSetter.TransactOpts, _value)
}

// SetBytes32 is a paid mutator transaction binding the contract method 0xc2b12a73.
//
// Solidity: function setBytes32(bytes32 _value) returns()
func (_GetterSetter *GetterSetterTransactorSession) SetBytes32(_value [32]byte) (*types.Transaction, error) {
	return _GetterSetter.Contract.SetBytes32(&_GetterSetter.TransactOpts, _value)
}

// SetUint256 is a paid mutator transaction binding the contract method 0xd2282dc5.
//
// Solidity: function setUint256(uint256 _value) returns()
func (_GetterSetter *GetterSetterTransactor) SetUint256(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _GetterSetter.contract.Transact(opts, "setUint256", _value)
}

// SetUint256 is a paid mutator transaction binding the contract method 0xd2282dc5.
//
// Solidity: function setUint256(uint256 _value) returns()
func (_GetterSetter *GetterSetterSession) SetUint256(_value *big.Int) (*types.Transaction, error) {
	return _GetterSetter.Contract.SetUint256(&_GetterSetter.TransactOpts, _value)
}

// SetUint256 is a paid mutator transaction binding the contract method 0xd2282dc5.
//
// Solidity: function setUint256(uint256 _value) returns()
func (_GetterSetter *GetterSetterTransactorSession) SetUint256(_value *big.Int) (*types.Transaction, error) {
	return _GetterSetter.Contract.SetUint256(&_GetterSetter.TransactOpts, _value)
}

// GetterSetterOutputIterator is returned from FilterOutput and is used to iterate over the raw logs and unpacked data for Output events raised by the GetterSetter contract.
type GetterSetterOutputIterator struct {
	Event *GetterSetterOutput // Event containing the contract specifics and raw log

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
func (it *GetterSetterOutputIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GetterSetterOutput)
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
		it.Event = new(GetterSetterOutput)
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
func (it *GetterSetterOutputIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GetterSetterOutputIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GetterSetterOutput represents a Output event raised by the GetterSetter contract.
type GetterSetterOutput struct {
	B32  [32]byte
	U256 *big.Int
	B322 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOutput is a free log retrieval operation binding the contract event 0x98353d8a928d8ef0ed244d763914357f71b8a4f66644fe62ea0aaf01c113a355.
//
// Solidity: event Output(bytes32 b32, uint256 u256, bytes32 b322)
func (_GetterSetter *GetterSetterFilterer) FilterOutput(opts *bind.FilterOpts) (*GetterSetterOutputIterator, error) {

	logs, sub, err := _GetterSetter.contract.FilterLogs(opts, "Output")
	if err != nil {
		return nil, err
	}
	return &GetterSetterOutputIterator{contract: _GetterSetter.contract, event: "Output", logs: logs, sub: sub}, nil
}

// WatchOutput is a free log subscription operation binding the contract event 0x98353d8a928d8ef0ed244d763914357f71b8a4f66644fe62ea0aaf01c113a355.
//
// Solidity: event Output(bytes32 b32, uint256 u256, bytes32 b322)
func (_GetterSetter *GetterSetterFilterer) WatchOutput(opts *bind.WatchOpts, sink chan<- *GetterSetterOutput) (event.Subscription, error) {

	logs, sub, err := _GetterSetter.contract.WatchLogs(opts, "Output")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GetterSetterOutput)
				if err := _GetterSetter.contract.UnpackLog(event, "Output", log); err != nil {
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

// ParseOutput is a log parse operation binding the contract event 0x98353d8a928d8ef0ed244d763914357f71b8a4f66644fe62ea0aaf01c113a355.
//
// Solidity: event Output(bytes32 b32, uint256 u256, bytes32 b322)
func (_GetterSetter *GetterSetterFilterer) ParseOutput(log types.Log) (*GetterSetterOutput, error) {
	event := new(GetterSetterOutput)
	if err := _GetterSetter.contract.UnpackLog(event, "Output", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GetterSetterSetBytesIterator is returned from FilterSetBytes and is used to iterate over the raw logs and unpacked data for SetBytes events raised by the GetterSetter contract.
type GetterSetterSetBytesIterator struct {
	Event *GetterSetterSetBytes // Event containing the contract specifics and raw log

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
func (it *GetterSetterSetBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GetterSetterSetBytes)
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
		it.Event = new(GetterSetterSetBytes)
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
func (it *GetterSetterSetBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GetterSetterSetBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GetterSetterSetBytes represents a SetBytes event raised by the GetterSetter contract.
type GetterSetterSetBytes struct {
	From  common.Address
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetBytes is a free log retrieval operation binding the contract event 0xf22a519d38e59bc517532f666f8da532fdd5356e68d617191e82a8fdcc8abdcf.
//
// Solidity: event SetBytes(address indexed from, bytes value)
func (_GetterSetter *GetterSetterFilterer) FilterSetBytes(opts *bind.FilterOpts, from []common.Address) (*GetterSetterSetBytesIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _GetterSetter.contract.FilterLogs(opts, "SetBytes", fromRule)
	if err != nil {
		return nil, err
	}
	return &GetterSetterSetBytesIterator{contract: _GetterSetter.contract, event: "SetBytes", logs: logs, sub: sub}, nil
}

// WatchSetBytes is a free log subscription operation binding the contract event 0xf22a519d38e59bc517532f666f8da532fdd5356e68d617191e82a8fdcc8abdcf.
//
// Solidity: event SetBytes(address indexed from, bytes value)
func (_GetterSetter *GetterSetterFilterer) WatchSetBytes(opts *bind.WatchOpts, sink chan<- *GetterSetterSetBytes, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _GetterSetter.contract.WatchLogs(opts, "SetBytes", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GetterSetterSetBytes)
				if err := _GetterSetter.contract.UnpackLog(event, "SetBytes", log); err != nil {
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

// ParseSetBytes is a log parse operation binding the contract event 0xf22a519d38e59bc517532f666f8da532fdd5356e68d617191e82a8fdcc8abdcf.
//
// Solidity: event SetBytes(address indexed from, bytes value)
func (_GetterSetter *GetterSetterFilterer) ParseSetBytes(log types.Log) (*GetterSetterSetBytes, error) {
	event := new(GetterSetterSetBytes)
	if err := _GetterSetter.contract.UnpackLog(event, "SetBytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GetterSetterSetBytes32Iterator is returned from FilterSetBytes32 and is used to iterate over the raw logs and unpacked data for SetBytes32 events raised by the GetterSetter contract.
type GetterSetterSetBytes32Iterator struct {
	Event *GetterSetterSetBytes32 // Event containing the contract specifics and raw log

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
func (it *GetterSetterSetBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GetterSetterSetBytes32)
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
		it.Event = new(GetterSetterSetBytes32)
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
func (it *GetterSetterSetBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GetterSetterSetBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GetterSetterSetBytes32 represents a SetBytes32 event raised by the GetterSetter contract.
type GetterSetterSetBytes32 struct {
	From  common.Address
	Value [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetBytes32 is a free log retrieval operation binding the contract event 0xdc73ee99832252105ed74a404690c2f10ad1b294cbbeb0ff5cded48ef2aa437d.
//
// Solidity: event SetBytes32(address indexed from, bytes32 indexed value)
func (_GetterSetter *GetterSetterFilterer) FilterSetBytes32(opts *bind.FilterOpts, from []common.Address, value [][32]byte) (*GetterSetterSetBytes32Iterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _GetterSetter.contract.FilterLogs(opts, "SetBytes32", fromRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &GetterSetterSetBytes32Iterator{contract: _GetterSetter.contract, event: "SetBytes32", logs: logs, sub: sub}, nil
}

// WatchSetBytes32 is a free log subscription operation binding the contract event 0xdc73ee99832252105ed74a404690c2f10ad1b294cbbeb0ff5cded48ef2aa437d.
//
// Solidity: event SetBytes32(address indexed from, bytes32 indexed value)
func (_GetterSetter *GetterSetterFilterer) WatchSetBytes32(opts *bind.WatchOpts, sink chan<- *GetterSetterSetBytes32, from []common.Address, value [][32]byte) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _GetterSetter.contract.WatchLogs(opts, "SetBytes32", fromRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GetterSetterSetBytes32)
				if err := _GetterSetter.contract.UnpackLog(event, "SetBytes32", log); err != nil {
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

// ParseSetBytes32 is a log parse operation binding the contract event 0xdc73ee99832252105ed74a404690c2f10ad1b294cbbeb0ff5cded48ef2aa437d.
//
// Solidity: event SetBytes32(address indexed from, bytes32 indexed value)
func (_GetterSetter *GetterSetterFilterer) ParseSetBytes32(log types.Log) (*GetterSetterSetBytes32, error) {
	event := new(GetterSetterSetBytes32)
	if err := _GetterSetter.contract.UnpackLog(event, "SetBytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GetterSetterSetUint256Iterator is returned from FilterSetUint256 and is used to iterate over the raw logs and unpacked data for SetUint256 events raised by the GetterSetter contract.
type GetterSetterSetUint256Iterator struct {
	Event *GetterSetterSetUint256 // Event containing the contract specifics and raw log

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
func (it *GetterSetterSetUint256Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GetterSetterSetUint256)
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
		it.Event = new(GetterSetterSetUint256)
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
func (it *GetterSetterSetUint256Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GetterSetterSetUint256Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GetterSetterSetUint256 represents a SetUint256 event raised by the GetterSetter contract.
type GetterSetterSetUint256 struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetUint256 is a free log retrieval operation binding the contract event 0xd943f063acdb1c6f206cf6a3f6d1ba39687bcc07feb7f44019bdbd4773c9c28d.
//
// Solidity: event SetUint256(address indexed from, uint256 indexed value)
func (_GetterSetter *GetterSetterFilterer) FilterSetUint256(opts *bind.FilterOpts, from []common.Address, value []*big.Int) (*GetterSetterSetUint256Iterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _GetterSetter.contract.FilterLogs(opts, "SetUint256", fromRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &GetterSetterSetUint256Iterator{contract: _GetterSetter.contract, event: "SetUint256", logs: logs, sub: sub}, nil
}

// WatchSetUint256 is a free log subscription operation binding the contract event 0xd943f063acdb1c6f206cf6a3f6d1ba39687bcc07feb7f44019bdbd4773c9c28d.
//
// Solidity: event SetUint256(address indexed from, uint256 indexed value)
func (_GetterSetter *GetterSetterFilterer) WatchSetUint256(opts *bind.WatchOpts, sink chan<- *GetterSetterSetUint256, from []common.Address, value []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _GetterSetter.contract.WatchLogs(opts, "SetUint256", fromRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GetterSetterSetUint256)
				if err := _GetterSetter.contract.UnpackLog(event, "SetUint256", log); err != nil {
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

// ParseSetUint256 is a log parse operation binding the contract event 0xd943f063acdb1c6f206cf6a3f6d1ba39687bcc07feb7f44019bdbd4773c9c28d.
//
// Solidity: event SetUint256(address indexed from, uint256 indexed value)
func (_GetterSetter *GetterSetterFilterer) ParseSetUint256(log types.Log) (*GetterSetterSetUint256, error) {
	event := new(GetterSetterSetUint256)
	if err := _GetterSetter.contract.UnpackLog(event, "SetUint256", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
