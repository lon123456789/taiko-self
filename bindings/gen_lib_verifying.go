// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// LibVerifyingMetaData contains all meta data concerning the LibVerifying contract.
var LibVerifyingMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"isConfigValid\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structTaikoData.Config\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockMaxProposals\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockRingBufferSize\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxBlocksToVerifyPerProposal\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockMaxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"blockMaxTxListBytes\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"blobExpiry\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"blobAllowedForDA\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"livenessBond\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"ethDepositRingBufferSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ethDepositMinCountPerBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"ethDepositMaxCountPerBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"ethDepositMinAmount\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"ethDepositMaxAmount\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"ethDepositGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ethDepositMaxFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"isValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"BlockVerified\",\"inputs\":[{\"name\":\"blockId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"assignedProver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"prover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"signalRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"tier\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"contestations\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CrossChainSynced\",\"inputs\":[{\"name\":\"syncedInBlock\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"blockId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"signalRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"L1_BLOCK_MISMATCH\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_INVALID_CONFIG\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_TRANSITION_ID_ZERO\",\"inputs\":[]}]",
}

// LibVerifyingABI is the input ABI used to generate the binding from.
// Deprecated: Use LibVerifyingMetaData.ABI instead.
var LibVerifyingABI = LibVerifyingMetaData.ABI

// LibVerifying is an auto generated Go binding around an Ethereum contract.
type LibVerifying struct {
	LibVerifyingCaller     // Read-only binding to the contract
	LibVerifyingTransactor // Write-only binding to the contract
	LibVerifyingFilterer   // Log filterer for contract events
}

// LibVerifyingCaller is an auto generated read-only Go binding around an Ethereum contract.
type LibVerifyingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibVerifyingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LibVerifyingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibVerifyingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LibVerifyingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibVerifyingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LibVerifyingSession struct {
	Contract     *LibVerifying     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LibVerifyingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LibVerifyingCallerSession struct {
	Contract *LibVerifyingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// LibVerifyingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LibVerifyingTransactorSession struct {
	Contract     *LibVerifyingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LibVerifyingRaw is an auto generated low-level Go binding around an Ethereum contract.
type LibVerifyingRaw struct {
	Contract *LibVerifying // Generic contract binding to access the raw methods on
}

// LibVerifyingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LibVerifyingCallerRaw struct {
	Contract *LibVerifyingCaller // Generic read-only contract binding to access the raw methods on
}

// LibVerifyingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LibVerifyingTransactorRaw struct {
	Contract *LibVerifyingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLibVerifying creates a new instance of LibVerifying, bound to a specific deployed contract.
func NewLibVerifying(address common.Address, backend bind.ContractBackend) (*LibVerifying, error) {
	contract, err := bindLibVerifying(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LibVerifying{LibVerifyingCaller: LibVerifyingCaller{contract: contract}, LibVerifyingTransactor: LibVerifyingTransactor{contract: contract}, LibVerifyingFilterer: LibVerifyingFilterer{contract: contract}}, nil
}

// NewLibVerifyingCaller creates a new read-only instance of LibVerifying, bound to a specific deployed contract.
func NewLibVerifyingCaller(address common.Address, caller bind.ContractCaller) (*LibVerifyingCaller, error) {
	contract, err := bindLibVerifying(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LibVerifyingCaller{contract: contract}, nil
}

// NewLibVerifyingTransactor creates a new write-only instance of LibVerifying, bound to a specific deployed contract.
func NewLibVerifyingTransactor(address common.Address, transactor bind.ContractTransactor) (*LibVerifyingTransactor, error) {
	contract, err := bindLibVerifying(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LibVerifyingTransactor{contract: contract}, nil
}

// NewLibVerifyingFilterer creates a new log filterer instance of LibVerifying, bound to a specific deployed contract.
func NewLibVerifyingFilterer(address common.Address, filterer bind.ContractFilterer) (*LibVerifyingFilterer, error) {
	contract, err := bindLibVerifying(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LibVerifyingFilterer{contract: contract}, nil
}

// bindLibVerifying binds a generic wrapper to an already deployed contract.
func bindLibVerifying(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LibVerifyingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibVerifying *LibVerifyingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LibVerifying.Contract.LibVerifyingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibVerifying *LibVerifyingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibVerifying.Contract.LibVerifyingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibVerifying *LibVerifyingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibVerifying.Contract.LibVerifyingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibVerifying *LibVerifyingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LibVerifying.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibVerifying *LibVerifyingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibVerifying.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibVerifying *LibVerifyingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibVerifying.Contract.contract.Transact(opts, method, params...)
}

// IsConfigValid is a free data retrieval call binding the contract method 0x3a5a6743.
//
// Solidity: function isConfigValid((uint64,uint64,uint64,uint64,uint32,uint24,uint24,bool,uint96,uint256,uint64,uint64,uint96,uint96,uint256,uint256) config) pure returns(bool isValid)
func (_LibVerifying *LibVerifyingCaller) IsConfigValid(opts *bind.CallOpts, config TaikoDataConfig) (bool, error) {
	var out []interface{}
	err := _LibVerifying.contract.Call(opts, &out, "isConfigValid", config)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsConfigValid is a free data retrieval call binding the contract method 0x3a5a6743.
//
// Solidity: function isConfigValid((uint64,uint64,uint64,uint64,uint32,uint24,uint24,bool,uint96,uint256,uint64,uint64,uint96,uint96,uint256,uint256) config) pure returns(bool isValid)
func (_LibVerifying *LibVerifyingSession) IsConfigValid(config TaikoDataConfig) (bool, error) {
	return _LibVerifying.Contract.IsConfigValid(&_LibVerifying.CallOpts, config)
}

// IsConfigValid is a free data retrieval call binding the contract method 0x3a5a6743.
//
// Solidity: function isConfigValid((uint64,uint64,uint64,uint64,uint32,uint24,uint24,bool,uint96,uint256,uint64,uint64,uint96,uint96,uint256,uint256) config) pure returns(bool isValid)
func (_LibVerifying *LibVerifyingCallerSession) IsConfigValid(config TaikoDataConfig) (bool, error) {
	return _LibVerifying.Contract.IsConfigValid(&_LibVerifying.CallOpts, config)
}

// LibVerifyingBlockVerifiedIterator is returned from FilterBlockVerified and is used to iterate over the raw logs and unpacked data for BlockVerified events raised by the LibVerifying contract.
type LibVerifyingBlockVerifiedIterator struct {
	Event *LibVerifyingBlockVerified // Event containing the contract specifics and raw log

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
func (it *LibVerifyingBlockVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LibVerifyingBlockVerified)
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
		it.Event = new(LibVerifyingBlockVerified)
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
func (it *LibVerifyingBlockVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LibVerifyingBlockVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LibVerifyingBlockVerified represents a BlockVerified event raised by the LibVerifying contract.
type LibVerifyingBlockVerified struct {
	BlockId        *big.Int
	AssignedProver common.Address
	Prover         common.Address
	BlockHash      [32]byte
	SignalRoot     [32]byte
	Tier           uint16
	Contestations  uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlockVerified is a free log retrieval operation binding the contract event 0xaeba6e73abba9419294b1017075cf8dc2e7de6f2d7fd3b336b3ba882a2acfca5.
//
// Solidity: event BlockVerified(uint256 indexed blockId, address indexed assignedProver, address indexed prover, bytes32 blockHash, bytes32 signalRoot, uint16 tier, uint8 contestations)
func (_LibVerifying *LibVerifyingFilterer) FilterBlockVerified(opts *bind.FilterOpts, blockId []*big.Int, assignedProver []common.Address, prover []common.Address) (*LibVerifyingBlockVerifiedIterator, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}
	var assignedProverRule []interface{}
	for _, assignedProverItem := range assignedProver {
		assignedProverRule = append(assignedProverRule, assignedProverItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _LibVerifying.contract.FilterLogs(opts, "BlockVerified", blockIdRule, assignedProverRule, proverRule)
	if err != nil {
		return nil, err
	}
	return &LibVerifyingBlockVerifiedIterator{contract: _LibVerifying.contract, event: "BlockVerified", logs: logs, sub: sub}, nil
}

// WatchBlockVerified is a free log subscription operation binding the contract event 0xaeba6e73abba9419294b1017075cf8dc2e7de6f2d7fd3b336b3ba882a2acfca5.
//
// Solidity: event BlockVerified(uint256 indexed blockId, address indexed assignedProver, address indexed prover, bytes32 blockHash, bytes32 signalRoot, uint16 tier, uint8 contestations)
func (_LibVerifying *LibVerifyingFilterer) WatchBlockVerified(opts *bind.WatchOpts, sink chan<- *LibVerifyingBlockVerified, blockId []*big.Int, assignedProver []common.Address, prover []common.Address) (event.Subscription, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}
	var assignedProverRule []interface{}
	for _, assignedProverItem := range assignedProver {
		assignedProverRule = append(assignedProverRule, assignedProverItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _LibVerifying.contract.WatchLogs(opts, "BlockVerified", blockIdRule, assignedProverRule, proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LibVerifyingBlockVerified)
				if err := _LibVerifying.contract.UnpackLog(event, "BlockVerified", log); err != nil {
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

// ParseBlockVerified is a log parse operation binding the contract event 0xaeba6e73abba9419294b1017075cf8dc2e7de6f2d7fd3b336b3ba882a2acfca5.
//
// Solidity: event BlockVerified(uint256 indexed blockId, address indexed assignedProver, address indexed prover, bytes32 blockHash, bytes32 signalRoot, uint16 tier, uint8 contestations)
func (_LibVerifying *LibVerifyingFilterer) ParseBlockVerified(log types.Log) (*LibVerifyingBlockVerified, error) {
	event := new(LibVerifyingBlockVerified)
	if err := _LibVerifying.contract.UnpackLog(event, "BlockVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LibVerifyingCrossChainSyncedIterator is returned from FilterCrossChainSynced and is used to iterate over the raw logs and unpacked data for CrossChainSynced events raised by the LibVerifying contract.
type LibVerifyingCrossChainSyncedIterator struct {
	Event *LibVerifyingCrossChainSynced // Event containing the contract specifics and raw log

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
func (it *LibVerifyingCrossChainSyncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LibVerifyingCrossChainSynced)
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
		it.Event = new(LibVerifyingCrossChainSynced)
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
func (it *LibVerifyingCrossChainSyncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LibVerifyingCrossChainSyncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LibVerifyingCrossChainSynced represents a CrossChainSynced event raised by the LibVerifying contract.
type LibVerifyingCrossChainSynced struct {
	SyncedInBlock uint64
	BlockId       uint64
	BlockHash     [32]byte
	SignalRoot    [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCrossChainSynced is a free log retrieval operation binding the contract event 0xf35ec3b262cf74881db1b8051c635496bccb1497a1e776dacb463d0e0e2b0f51.
//
// Solidity: event CrossChainSynced(uint64 indexed syncedInBlock, uint64 indexed blockId, bytes32 blockHash, bytes32 signalRoot)
func (_LibVerifying *LibVerifyingFilterer) FilterCrossChainSynced(opts *bind.FilterOpts, syncedInBlock []uint64, blockId []uint64) (*LibVerifyingCrossChainSyncedIterator, error) {

	var syncedInBlockRule []interface{}
	for _, syncedInBlockItem := range syncedInBlock {
		syncedInBlockRule = append(syncedInBlockRule, syncedInBlockItem)
	}
	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}

	logs, sub, err := _LibVerifying.contract.FilterLogs(opts, "CrossChainSynced", syncedInBlockRule, blockIdRule)
	if err != nil {
		return nil, err
	}
	return &LibVerifyingCrossChainSyncedIterator{contract: _LibVerifying.contract, event: "CrossChainSynced", logs: logs, sub: sub}, nil
}

// WatchCrossChainSynced is a free log subscription operation binding the contract event 0xf35ec3b262cf74881db1b8051c635496bccb1497a1e776dacb463d0e0e2b0f51.
//
// Solidity: event CrossChainSynced(uint64 indexed syncedInBlock, uint64 indexed blockId, bytes32 blockHash, bytes32 signalRoot)
func (_LibVerifying *LibVerifyingFilterer) WatchCrossChainSynced(opts *bind.WatchOpts, sink chan<- *LibVerifyingCrossChainSynced, syncedInBlock []uint64, blockId []uint64) (event.Subscription, error) {

	var syncedInBlockRule []interface{}
	for _, syncedInBlockItem := range syncedInBlock {
		syncedInBlockRule = append(syncedInBlockRule, syncedInBlockItem)
	}
	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}

	logs, sub, err := _LibVerifying.contract.WatchLogs(opts, "CrossChainSynced", syncedInBlockRule, blockIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LibVerifyingCrossChainSynced)
				if err := _LibVerifying.contract.UnpackLog(event, "CrossChainSynced", log); err != nil {
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

// ParseCrossChainSynced is a log parse operation binding the contract event 0xf35ec3b262cf74881db1b8051c635496bccb1497a1e776dacb463d0e0e2b0f51.
//
// Solidity: event CrossChainSynced(uint64 indexed syncedInBlock, uint64 indexed blockId, bytes32 blockHash, bytes32 signalRoot)
func (_LibVerifying *LibVerifyingFilterer) ParseCrossChainSynced(log types.Log) (*LibVerifyingCrossChainSynced, error) {
	event := new(LibVerifyingCrossChainSynced)
	if err := _LibVerifying.contract.UnpackLog(event, "CrossChainSynced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
