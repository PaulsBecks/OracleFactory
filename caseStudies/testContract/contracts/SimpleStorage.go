// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StoreABI is the input ABI used to generate the binding from.
const StoreABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"storedInteger\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"storedBool\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"storedInteger8\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"storedAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"storedBytes32\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"storedBytes\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"storedString\",\"type\":\"string\"}],\"name\":\"StoredAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"storedInteger\",\"type\":\"uint256\"}],\"name\":\"StoredInteger\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"integer\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"boolean\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"integer8\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_bytes32\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_bytes\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"_string\",\"type\":\"string\"}],\"name\":\"setAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"integer\",\"type\":\"uint256\"}],\"name\":\"setInteger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// SetAll is a paid mutator transaction binding the contract method 0x67c157a1.
//
// Solidity: function setAll(uint256 integer, bool boolean, uint8 integer8, address _address, bytes32 _bytes32, bytes _bytes, string _string) returns()
func (_Store *StoreTransactor) SetAll(opts *bind.TransactOpts, integer *big.Int, boolean bool, integer8 uint8, _address common.Address, _bytes32 [32]byte, _bytes []byte, _string string) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setAll", integer, boolean, integer8, _address, _bytes32, _bytes, _string)
}

// SetAll is a paid mutator transaction binding the contract method 0x67c157a1.
//
// Solidity: function setAll(uint256 integer, bool boolean, uint8 integer8, address _address, bytes32 _bytes32, bytes _bytes, string _string) returns()
func (_Store *StoreSession) SetAll(integer *big.Int, boolean bool, integer8 uint8, _address common.Address, _bytes32 [32]byte, _bytes []byte, _string string) (*types.Transaction, error) {
	return _Store.Contract.SetAll(&_Store.TransactOpts, integer, boolean, integer8, _address, _bytes32, _bytes, _string)
}

// SetAll is a paid mutator transaction binding the contract method 0x67c157a1.
//
// Solidity: function setAll(uint256 integer, bool boolean, uint8 integer8, address _address, bytes32 _bytes32, bytes _bytes, string _string) returns()
func (_Store *StoreTransactorSession) SetAll(integer *big.Int, boolean bool, integer8 uint8, _address common.Address, _bytes32 [32]byte, _bytes []byte, _string string) (*types.Transaction, error) {
	return _Store.Contract.SetAll(&_Store.TransactOpts, integer, boolean, integer8, _address, _bytes32, _bytes, _string)
}

// SetInteger is a paid mutator transaction binding the contract method 0xac588675.
//
// Solidity: function setInteger(uint256 integer) returns()
func (_Store *StoreTransactor) SetInteger(opts *bind.TransactOpts, integer *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setInteger", integer)
}

// SetInteger is a paid mutator transaction binding the contract method 0xac588675.
//
// Solidity: function setInteger(uint256 integer) returns()
func (_Store *StoreSession) SetInteger(integer *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SetInteger(&_Store.TransactOpts, integer)
}

// SetInteger is a paid mutator transaction binding the contract method 0xac588675.
//
// Solidity: function setInteger(uint256 integer) returns()
func (_Store *StoreTransactorSession) SetInteger(integer *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SetInteger(&_Store.TransactOpts, integer)
}

// StoreStoredAllIterator is returned from FilterStoredAll and is used to iterate over the raw logs and unpacked data for StoredAll events raised by the Store contract.
type StoreStoredAllIterator struct {
	Event *StoreStoredAll // Event containing the contract specifics and raw log

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
func (it *StoreStoredAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreStoredAll)
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
		it.Event = new(StoreStoredAll)
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
func (it *StoreStoredAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreStoredAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreStoredAll represents a StoredAll event raised by the Store contract.
type StoreStoredAll struct {
	StoredInteger  *big.Int
	StoredBool     bool
	StoredInteger8 uint8
	StoredAddress  common.Address
	StoredBytes32  [32]byte
	StoredBytes    []byte
	StoredString   string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterStoredAll is a free log retrieval operation binding the contract event 0xb64568d9b0a86ee0722bd1773d93497c73572db44ac1c1c0baeb14cd8e02e837.
//
// Solidity: event StoredAll(uint256 storedInteger, bool storedBool, uint8 storedInteger8, address storedAddress, bytes32 storedBytes32, bytes storedBytes, string storedString)
func (_Store *StoreFilterer) FilterStoredAll(opts *bind.FilterOpts) (*StoreStoredAllIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "StoredAll")
	if err != nil {
		return nil, err
	}
	return &StoreStoredAllIterator{contract: _Store.contract, event: "StoredAll", logs: logs, sub: sub}, nil
}

// WatchStoredAll is a free log subscription operation binding the contract event 0xb64568d9b0a86ee0722bd1773d93497c73572db44ac1c1c0baeb14cd8e02e837.
//
// Solidity: event StoredAll(uint256 storedInteger, bool storedBool, uint8 storedInteger8, address storedAddress, bytes32 storedBytes32, bytes storedBytes, string storedString)
func (_Store *StoreFilterer) WatchStoredAll(opts *bind.WatchOpts, sink chan<- *StoreStoredAll) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "StoredAll")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreStoredAll)
				if err := _Store.contract.UnpackLog(event, "StoredAll", log); err != nil {
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

// ParseStoredAll is a log parse operation binding the contract event 0xb64568d9b0a86ee0722bd1773d93497c73572db44ac1c1c0baeb14cd8e02e837.
//
// Solidity: event StoredAll(uint256 storedInteger, bool storedBool, uint8 storedInteger8, address storedAddress, bytes32 storedBytes32, bytes storedBytes, string storedString)
func (_Store *StoreFilterer) ParseStoredAll(log types.Log) (*StoreStoredAll, error) {
	event := new(StoreStoredAll)
	if err := _Store.contract.UnpackLog(event, "StoredAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreStoredIntegerIterator is returned from FilterStoredInteger and is used to iterate over the raw logs and unpacked data for StoredInteger events raised by the Store contract.
type StoreStoredIntegerIterator struct {
	Event *StoreStoredInteger // Event containing the contract specifics and raw log

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
func (it *StoreStoredIntegerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreStoredInteger)
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
		it.Event = new(StoreStoredInteger)
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
func (it *StoreStoredIntegerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreStoredIntegerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreStoredInteger represents a StoredInteger event raised by the Store contract.
type StoreStoredInteger struct {
	StoredInteger *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterStoredInteger is a free log retrieval operation binding the contract event 0xe6452fda282d738b0dafdab845143683b1ddb1739ed11699e10bea1e11aed170.
//
// Solidity: event StoredInteger(uint256 storedInteger)
func (_Store *StoreFilterer) FilterStoredInteger(opts *bind.FilterOpts) (*StoreStoredIntegerIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "StoredInteger")
	if err != nil {
		return nil, err
	}
	return &StoreStoredIntegerIterator{contract: _Store.contract, event: "StoredInteger", logs: logs, sub: sub}, nil
}

// WatchStoredInteger is a free log subscription operation binding the contract event 0xe6452fda282d738b0dafdab845143683b1ddb1739ed11699e10bea1e11aed170.
//
// Solidity: event StoredInteger(uint256 storedInteger)
func (_Store *StoreFilterer) WatchStoredInteger(opts *bind.WatchOpts, sink chan<- *StoreStoredInteger) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "StoredInteger")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreStoredInteger)
				if err := _Store.contract.UnpackLog(event, "StoredInteger", log); err != nil {
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

// ParseStoredInteger is a log parse operation binding the contract event 0xe6452fda282d738b0dafdab845143683b1ddb1739ed11699e10bea1e11aed170.
//
// Solidity: event StoredInteger(uint256 storedInteger)
func (_Store *StoreFilterer) ParseStoredInteger(log types.Log) (*StoreStoredInteger, error) {
	event := new(StoreStoredInteger)
	if err := _Store.contract.UnpackLog(event, "StoredInteger", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
