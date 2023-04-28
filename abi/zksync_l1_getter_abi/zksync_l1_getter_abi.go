// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zksync_l1_getter_abi

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

// IGettersMetaData contains all meta data concerning the IGetters contract.
var IGettersMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"facetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"facet\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facetAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"facets\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_facet\",\"type\":\"address\"}],\"name\":\"facetFunctionSelectors\",\"outputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"\",\"type\":\"bytes4[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentProposalId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFirstUnprocessedPriorityTx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGovernor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getL2BootloaderBytecodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getL2DefaultAccountBytecodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPendingGovernor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriorityQueueSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriorityTxMaxGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProposedUpgradeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProposedUpgradeTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSecurityCouncil\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalBlocksCommitted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalBlocksExecuted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalBlocksVerified\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalPriorityTxs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isApprovedBySecurityCouncil\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isDiamondStorageFrozen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"}],\"name\":\"isEthWithdrawalFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_facet\",\"type\":\"address\"}],\"name\":\"isFacetFreezable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isFreezable\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"isFunctionFreezable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"l2LogsRootHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"storedBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"cdffacc6": "facetAddress(bytes4)",
		"52ef6b2c": "facetAddresses()",
		"adfca15e": "facetFunctionSelectors(address)",
		"a7cd63b7": "getAllowList()",
		"fe10226d": "getCurrentProposalId()",
		"79823c9a": "getFirstUnprocessedPriorityTx()",
		"4fc07d75": "getGovernor()",
		"d86970d8": "getL2BootloaderBytecodeHash()",
		"fd791f3c": "getL2DefaultAccountBytecodeHash()",
		"8665b150": "getPendingGovernor()",
		"631f4bac": "getPriorityQueueSize()",
		"0ec6b0b7": "getPriorityTxMaxGasLimit()",
		"1b60e626": "getProposedUpgradeHash()",
		"e39d3bff": "getProposedUpgradeTimestamp()",
		"0ef240a0": "getSecurityCouncil()",
		"fe26699e": "getTotalBlocksCommitted()",
		"39607382": "getTotalBlocksExecuted()",
		"af6a2dcd": "getTotalBlocksVerified()",
		"a1954fc5": "getTotalPriorityTxs()",
		"46657fe9": "getVerifier()",
		"3db920ce": "isApprovedBySecurityCouncil()",
		"29b98c67": "isDiamondStorageFrozen()",
		"bd7c5412": "isEthWithdrawalFinalized(uint256,uint256)",
		"c3bbd2d7": "isFacetFreezable(address)",
		"e81e0ba1": "isFunctionFreezable(bytes4)",
		"facd743b": "isValidator(address)",
		"9cd939e4": "l2LogsRootHash(uint256)",
		"74f4d30d": "storedBlockHash(uint256)",
	},
}

// IGettersABI is the input ABI used to generate the binding from.
// Deprecated: Use IGettersMetaData.ABI instead.
var IGettersABI = IGettersMetaData.ABI

// Deprecated: Use IGettersMetaData.Sigs instead.
// IGettersFuncSigs maps the 4-byte function signature to its string representation.
var IGettersFuncSigs = IGettersMetaData.Sigs

// IGetters is an auto generated Go binding around an Ethereum contract.
type IGetters struct {
	IGettersCaller     // Read-only binding to the contract
	IGettersTransactor // Write-only binding to the contract
	IGettersFilterer   // Log filterer for contract events
}

// IGettersCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGettersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGettersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGettersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGettersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGettersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGettersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGettersSession struct {
	Contract     *IGetters         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGettersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGettersCallerSession struct {
	Contract *IGettersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IGettersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGettersTransactorSession struct {
	Contract     *IGettersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IGettersRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGettersRaw struct {
	Contract *IGetters // Generic contract binding to access the raw methods on
}

// IGettersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGettersCallerRaw struct {
	Contract *IGettersCaller // Generic read-only contract binding to access the raw methods on
}

// IGettersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGettersTransactorRaw struct {
	Contract *IGettersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGetters creates a new instance of IGetters, bound to a specific deployed contract.
func NewIGetters(address common.Address, backend bind.ContractBackend) (*IGetters, error) {
	contract, err := bindIGetters(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGetters{IGettersCaller: IGettersCaller{contract: contract}, IGettersTransactor: IGettersTransactor{contract: contract}, IGettersFilterer: IGettersFilterer{contract: contract}}, nil
}

// NewIGettersCaller creates a new read-only instance of IGetters, bound to a specific deployed contract.
func NewIGettersCaller(address common.Address, caller bind.ContractCaller) (*IGettersCaller, error) {
	contract, err := bindIGetters(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGettersCaller{contract: contract}, nil
}

// NewIGettersTransactor creates a new write-only instance of IGetters, bound to a specific deployed contract.
func NewIGettersTransactor(address common.Address, transactor bind.ContractTransactor) (*IGettersTransactor, error) {
	contract, err := bindIGetters(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGettersTransactor{contract: contract}, nil
}

// NewIGettersFilterer creates a new log filterer instance of IGetters, bound to a specific deployed contract.
func NewIGettersFilterer(address common.Address, filterer bind.ContractFilterer) (*IGettersFilterer, error) {
	contract, err := bindIGetters(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGettersFilterer{contract: contract}, nil
}

// bindIGetters binds a generic wrapper to an already deployed contract.
func bindIGetters(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IGettersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGetters *IGettersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGetters.Contract.IGettersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGetters *IGettersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGetters.Contract.IGettersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGetters *IGettersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGetters.Contract.IGettersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGetters *IGettersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGetters.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGetters *IGettersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGetters.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGetters *IGettersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGetters.Contract.contract.Transact(opts, method, params...)
}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _selector) view returns(address facet)
func (_IGetters *IGettersCaller) FacetAddress(opts *bind.CallOpts, _selector [4]byte) (common.Address, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "facetAddress", _selector)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _selector) view returns(address facet)
func (_IGetters *IGettersSession) FacetAddress(_selector [4]byte) (common.Address, error) {
	return _IGetters.Contract.FacetAddress(&_IGetters.CallOpts, _selector)
}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _selector) view returns(address facet)
func (_IGetters *IGettersCallerSession) FacetAddress(_selector [4]byte) (common.Address, error) {
	return _IGetters.Contract.FacetAddress(&_IGetters.CallOpts, _selector)
}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facets)
func (_IGetters *IGettersCaller) FacetAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "facetAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facets)
func (_IGetters *IGettersSession) FacetAddresses() ([]common.Address, error) {
	return _IGetters.Contract.FacetAddresses(&_IGetters.CallOpts)
}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facets)
func (_IGetters *IGettersCallerSession) FacetAddresses() ([]common.Address, error) {
	return _IGetters.Contract.FacetAddresses(&_IGetters.CallOpts)
}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[])
func (_IGetters *IGettersCaller) FacetFunctionSelectors(opts *bind.CallOpts, _facet common.Address) ([][4]byte, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "facetFunctionSelectors", _facet)

	if err != nil {
		return *new([][4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][4]byte)).(*[][4]byte)

	return out0, err

}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[])
func (_IGetters *IGettersSession) FacetFunctionSelectors(_facet common.Address) ([][4]byte, error) {
	return _IGetters.Contract.FacetFunctionSelectors(&_IGetters.CallOpts, _facet)
}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[])
func (_IGetters *IGettersCallerSession) FacetFunctionSelectors(_facet common.Address) ([][4]byte, error) {
	return _IGetters.Contract.FacetFunctionSelectors(&_IGetters.CallOpts, _facet)
}

// GetAllowList is a free data retrieval call binding the contract method 0xa7cd63b7.
//
// Solidity: function getAllowList() view returns(address)
func (_IGetters *IGettersCaller) GetAllowList(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAllowList is a free data retrieval call binding the contract method 0xa7cd63b7.
//
// Solidity: function getAllowList() view returns(address)
func (_IGetters *IGettersSession) GetAllowList() (common.Address, error) {
	return _IGetters.Contract.GetAllowList(&_IGetters.CallOpts)
}

// GetAllowList is a free data retrieval call binding the contract method 0xa7cd63b7.
//
// Solidity: function getAllowList() view returns(address)
func (_IGetters *IGettersCallerSession) GetAllowList() (common.Address, error) {
	return _IGetters.Contract.GetAllowList(&_IGetters.CallOpts)
}

// GetCurrentProposalId is a free data retrieval call binding the contract method 0xfe10226d.
//
// Solidity: function getCurrentProposalId() view returns(uint256)
func (_IGetters *IGettersCaller) GetCurrentProposalId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "getCurrentProposalId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentProposalId is a free data retrieval call binding the contract method 0xfe10226d.
//
// Solidity: function getCurrentProposalId() view returns(uint256)
func (_IGetters *IGettersSession) GetCurrentProposalId() (*big.Int, error) {
	return _IGetters.Contract.GetCurrentProposalId(&_IGetters.CallOpts)
}

// GetCurrentProposalId is a free data retrieval call binding the contract method 0xfe10226d.
//
// Solidity: function getCurrentProposalId() view returns(uint256)
func (_IGetters *IGettersCallerSession) GetCurrentProposalId() (*big.Int, error) {
	return _IGetters.Contract.GetCurrentProposalId(&_IGetters.CallOpts)
}

// GetFirstUnprocessedPriorityTx is a free data retrieval call binding the contract method 0x79823c9a.
//
// Solidity: function getFirstUnprocessedPriorityTx() view returns(uint256)
func (_IGetters *IGettersCaller) GetFirstUnprocessedPriorityTx(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "getFirstUnprocessedPriorityTx")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFirstUnprocessedPriorityTx is a free data retrieval call binding the contract method 0x79823c9a.
//
// Solidity: function getFirstUnprocessedPriorityTx() view returns(uint256)
func (_IGetters *IGettersSession) GetFirstUnprocessedPriorityTx() (*big.Int, error) {
	return _IGetters.Contract.GetFirstUnprocessedPriorityTx(&_IGetters.CallOpts)
}

// GetFirstUnprocessedPriorityTx is a free data retrieval call binding the contract method 0x79823c9a.
//
// Solidity: function getFirstUnprocessedPriorityTx() view returns(uint256)
func (_IGetters *IGettersCallerSession) GetFirstUnprocessedPriorityTx() (*big.Int, error) {
	return _IGetters.Contract.GetFirstUnprocessedPriorityTx(&_IGetters.CallOpts)
}

// GetGovernor is a free data retrieval call binding the contract method 0x4fc07d75.
//
// Solidity: function getGovernor() view returns(address)
func (_IGetters *IGettersCaller) GetGovernor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "getGovernor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGovernor is a free data retrieval call binding the contract method 0x4fc07d75.
//
// Solidity: function getGovernor() view returns(address)
func (_IGetters *IGettersSession) GetGovernor() (common.Address, error) {
	return _IGetters.Contract.GetGovernor(&_IGetters.CallOpts)
}

// GetGovernor is a free data retrieval call binding the contract method 0x4fc07d75.
//
// Solidity: function getGovernor() view returns(address)
func (_IGetters *IGettersCallerSession) GetGovernor() (common.Address, error) {
	return _IGetters.Contract.GetGovernor(&_IGetters.CallOpts)
}

// GetL2BootloaderBytecodeHash is a free data retrieval call binding the contract method 0xd86970d8.
//
// Solidity: function getL2BootloaderBytecodeHash() view returns(bytes32)
func (_IGetters *IGettersCaller) GetL2BootloaderBytecodeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "getL2BootloaderBytecodeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetL2BootloaderBytecodeHash is a free data retrieval call binding the contract method 0xd86970d8.
//
// Solidity: function getL2BootloaderBytecodeHash() view returns(bytes32)
func (_IGetters *IGettersSession) GetL2BootloaderBytecodeHash() ([32]byte, error) {
	return _IGetters.Contract.GetL2BootloaderBytecodeHash(&_IGetters.CallOpts)
}

// GetL2BootloaderBytecodeHash is a free data retrieval call binding the contract method 0xd86970d8.
//
// Solidity: function getL2BootloaderBytecodeHash() view returns(bytes32)
func (_IGetters *IGettersCallerSession) GetL2BootloaderBytecodeHash() ([32]byte, error) {
	return _IGetters.Contract.GetL2BootloaderBytecodeHash(&_IGetters.CallOpts)
}

// GetL2DefaultAccountBytecodeHash is a free data retrieval call binding the contract method 0xfd791f3c.
//
// Solidity: function getL2DefaultAccountBytecodeHash() view returns(bytes32)
func (_IGetters *IGettersCaller) GetL2DefaultAccountBytecodeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "getL2DefaultAccountBytecodeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetL2DefaultAccountBytecodeHash is a free data retrieval call binding the contract method 0xfd791f3c.
//
// Solidity: function getL2DefaultAccountBytecodeHash() view returns(bytes32)
func (_IGetters *IGettersSession) GetL2DefaultAccountBytecodeHash() ([32]byte, error) {
	return _IGetters.Contract.GetL2DefaultAccountBytecodeHash(&_IGetters.CallOpts)
}

// GetL2DefaultAccountBytecodeHash is a free data retrieval call binding the contract method 0xfd791f3c.
//
// Solidity: function getL2DefaultAccountBytecodeHash() view returns(bytes32)
func (_IGetters *IGettersCallerSession) GetL2DefaultAccountBytecodeHash() ([32]byte, error) {
	return _IGetters.Contract.GetL2DefaultAccountBytecodeHash(&_IGetters.CallOpts)
}

// GetPendingGovernor is a free data retrieval call binding the contract method 0x8665b150.
//
// Solidity: function getPendingGovernor() view returns(address)
func (_IGetters *IGettersCaller) GetPendingGovernor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "getPendingGovernor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPendingGovernor is a free data retrieval call binding the contract method 0x8665b150.
//
// Solidity: function getPendingGovernor() view returns(address)
func (_IGetters *IGettersSession) GetPendingGovernor() (common.Address, error) {
	return _IGetters.Contract.GetPendingGovernor(&_IGetters.CallOpts)
}

// GetPendingGovernor is a free data retrieval call binding the contract method 0x8665b150.
//
// Solidity: function getPendingGovernor() view returns(address)
func (_IGetters *IGettersCallerSession) GetPendingGovernor() (common.Address, error) {
	return _IGetters.Contract.GetPendingGovernor(&_IGetters.CallOpts)
}

// GetPriorityQueueSize is a free data retrieval call binding the contract method 0x631f4bac.
//
// Solidity: function getPriorityQueueSize() view returns(uint256)
func (_IGetters *IGettersCaller) GetPriorityQueueSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "getPriorityQueueSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriorityQueueSize is a free data retrieval call binding the contract method 0x631f4bac.
//
// Solidity: function getPriorityQueueSize() view returns(uint256)
func (_IGetters *IGettersSession) GetPriorityQueueSize() (*big.Int, error) {
	return _IGetters.Contract.GetPriorityQueueSize(&_IGetters.CallOpts)
}

// GetPriorityQueueSize is a free data retrieval call binding the contract method 0x631f4bac.
//
// Solidity: function getPriorityQueueSize() view returns(uint256)
func (_IGetters *IGettersCallerSession) GetPriorityQueueSize() (*big.Int, error) {
	return _IGetters.Contract.GetPriorityQueueSize(&_IGetters.CallOpts)
}

// GetPriorityTxMaxGasLimit is a free data retrieval call binding the contract method 0x0ec6b0b7.
//
// Solidity: function getPriorityTxMaxGasLimit() view returns(uint256)
func (_IGetters *IGettersCaller) GetPriorityTxMaxGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "getPriorityTxMaxGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriorityTxMaxGasLimit is a free data retrieval call binding the contract method 0x0ec6b0b7.
//
// Solidity: function getPriorityTxMaxGasLimit() view returns(uint256)
func (_IGetters *IGettersSession) GetPriorityTxMaxGasLimit() (*big.Int, error) {
	return _IGetters.Contract.GetPriorityTxMaxGasLimit(&_IGetters.CallOpts)
}

// GetPriorityTxMaxGasLimit is a free data retrieval call binding the contract method 0x0ec6b0b7.
//
// Solidity: function getPriorityTxMaxGasLimit() view returns(uint256)
func (_IGetters *IGettersCallerSession) GetPriorityTxMaxGasLimit() (*big.Int, error) {
	return _IGetters.Contract.GetPriorityTxMaxGasLimit(&_IGetters.CallOpts)
}

// GetProposedUpgradeHash is a free data retrieval call binding the contract method 0x1b60e626.
//
// Solidity: function getProposedUpgradeHash() view returns(bytes32)
func (_IGetters *IGettersCaller) GetProposedUpgradeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "getProposedUpgradeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetProposedUpgradeHash is a free data retrieval call binding the contract method 0x1b60e626.
//
// Solidity: function getProposedUpgradeHash() view returns(bytes32)
func (_IGetters *IGettersSession) GetProposedUpgradeHash() ([32]byte, error) {
	return _IGetters.Contract.GetProposedUpgradeHash(&_IGetters.CallOpts)
}

// GetProposedUpgradeHash is a free data retrieval call binding the contract method 0x1b60e626.
//
// Solidity: function getProposedUpgradeHash() view returns(bytes32)
func (_IGetters *IGettersCallerSession) GetProposedUpgradeHash() ([32]byte, error) {
	return _IGetters.Contract.GetProposedUpgradeHash(&_IGetters.CallOpts)
}

// GetProposedUpgradeTimestamp is a free data retrieval call binding the contract method 0xe39d3bff.
//
// Solidity: function getProposedUpgradeTimestamp() view returns(uint256)
func (_IGetters *IGettersCaller) GetProposedUpgradeTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "getProposedUpgradeTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProposedUpgradeTimestamp is a free data retrieval call binding the contract method 0xe39d3bff.
//
// Solidity: function getProposedUpgradeTimestamp() view returns(uint256)
func (_IGetters *IGettersSession) GetProposedUpgradeTimestamp() (*big.Int, error) {
	return _IGetters.Contract.GetProposedUpgradeTimestamp(&_IGetters.CallOpts)
}

// GetProposedUpgradeTimestamp is a free data retrieval call binding the contract method 0xe39d3bff.
//
// Solidity: function getProposedUpgradeTimestamp() view returns(uint256)
func (_IGetters *IGettersCallerSession) GetProposedUpgradeTimestamp() (*big.Int, error) {
	return _IGetters.Contract.GetProposedUpgradeTimestamp(&_IGetters.CallOpts)
}

// GetSecurityCouncil is a free data retrieval call binding the contract method 0x0ef240a0.
//
// Solidity: function getSecurityCouncil() view returns(address)
func (_IGetters *IGettersCaller) GetSecurityCouncil(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "getSecurityCouncil")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSecurityCouncil is a free data retrieval call binding the contract method 0x0ef240a0.
//
// Solidity: function getSecurityCouncil() view returns(address)
func (_IGetters *IGettersSession) GetSecurityCouncil() (common.Address, error) {
	return _IGetters.Contract.GetSecurityCouncil(&_IGetters.CallOpts)
}

// GetSecurityCouncil is a free data retrieval call binding the contract method 0x0ef240a0.
//
// Solidity: function getSecurityCouncil() view returns(address)
func (_IGetters *IGettersCallerSession) GetSecurityCouncil() (common.Address, error) {
	return _IGetters.Contract.GetSecurityCouncil(&_IGetters.CallOpts)
}

// GetTotalBlocksCommitted is a free data retrieval call binding the contract method 0xfe26699e.
//
// Solidity: function getTotalBlocksCommitted() view returns(uint256)
func (_IGetters *IGettersCaller) GetTotalBlocksCommitted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "getTotalBlocksCommitted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalBlocksCommitted is a free data retrieval call binding the contract method 0xfe26699e.
//
// Solidity: function getTotalBlocksCommitted() view returns(uint256)
func (_IGetters *IGettersSession) GetTotalBlocksCommitted() (*big.Int, error) {
	return _IGetters.Contract.GetTotalBlocksCommitted(&_IGetters.CallOpts)
}

// GetTotalBlocksCommitted is a free data retrieval call binding the contract method 0xfe26699e.
//
// Solidity: function getTotalBlocksCommitted() view returns(uint256)
func (_IGetters *IGettersCallerSession) GetTotalBlocksCommitted() (*big.Int, error) {
	return _IGetters.Contract.GetTotalBlocksCommitted(&_IGetters.CallOpts)
}

// GetTotalBlocksExecuted is a free data retrieval call binding the contract method 0x39607382.
//
// Solidity: function getTotalBlocksExecuted() view returns(uint256)
func (_IGetters *IGettersCaller) GetTotalBlocksExecuted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "getTotalBlocksExecuted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalBlocksExecuted is a free data retrieval call binding the contract method 0x39607382.
//
// Solidity: function getTotalBlocksExecuted() view returns(uint256)
func (_IGetters *IGettersSession) GetTotalBlocksExecuted() (*big.Int, error) {
	return _IGetters.Contract.GetTotalBlocksExecuted(&_IGetters.CallOpts)
}

// GetTotalBlocksExecuted is a free data retrieval call binding the contract method 0x39607382.
//
// Solidity: function getTotalBlocksExecuted() view returns(uint256)
func (_IGetters *IGettersCallerSession) GetTotalBlocksExecuted() (*big.Int, error) {
	return _IGetters.Contract.GetTotalBlocksExecuted(&_IGetters.CallOpts)
}

// GetTotalBlocksVerified is a free data retrieval call binding the contract method 0xaf6a2dcd.
//
// Solidity: function getTotalBlocksVerified() view returns(uint256)
func (_IGetters *IGettersCaller) GetTotalBlocksVerified(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "getTotalBlocksVerified")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalBlocksVerified is a free data retrieval call binding the contract method 0xaf6a2dcd.
//
// Solidity: function getTotalBlocksVerified() view returns(uint256)
func (_IGetters *IGettersSession) GetTotalBlocksVerified() (*big.Int, error) {
	return _IGetters.Contract.GetTotalBlocksVerified(&_IGetters.CallOpts)
}

// GetTotalBlocksVerified is a free data retrieval call binding the contract method 0xaf6a2dcd.
//
// Solidity: function getTotalBlocksVerified() view returns(uint256)
func (_IGetters *IGettersCallerSession) GetTotalBlocksVerified() (*big.Int, error) {
	return _IGetters.Contract.GetTotalBlocksVerified(&_IGetters.CallOpts)
}

// GetTotalPriorityTxs is a free data retrieval call binding the contract method 0xa1954fc5.
//
// Solidity: function getTotalPriorityTxs() view returns(uint256)
func (_IGetters *IGettersCaller) GetTotalPriorityTxs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "getTotalPriorityTxs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPriorityTxs is a free data retrieval call binding the contract method 0xa1954fc5.
//
// Solidity: function getTotalPriorityTxs() view returns(uint256)
func (_IGetters *IGettersSession) GetTotalPriorityTxs() (*big.Int, error) {
	return _IGetters.Contract.GetTotalPriorityTxs(&_IGetters.CallOpts)
}

// GetTotalPriorityTxs is a free data retrieval call binding the contract method 0xa1954fc5.
//
// Solidity: function getTotalPriorityTxs() view returns(uint256)
func (_IGetters *IGettersCallerSession) GetTotalPriorityTxs() (*big.Int, error) {
	return _IGetters.Contract.GetTotalPriorityTxs(&_IGetters.CallOpts)
}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_IGetters *IGettersCaller) GetVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "getVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_IGetters *IGettersSession) GetVerifier() (common.Address, error) {
	return _IGetters.Contract.GetVerifier(&_IGetters.CallOpts)
}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_IGetters *IGettersCallerSession) GetVerifier() (common.Address, error) {
	return _IGetters.Contract.GetVerifier(&_IGetters.CallOpts)
}

// IsApprovedBySecurityCouncil is a free data retrieval call binding the contract method 0x3db920ce.
//
// Solidity: function isApprovedBySecurityCouncil() view returns(bool)
func (_IGetters *IGettersCaller) IsApprovedBySecurityCouncil(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "isApprovedBySecurityCouncil")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedBySecurityCouncil is a free data retrieval call binding the contract method 0x3db920ce.
//
// Solidity: function isApprovedBySecurityCouncil() view returns(bool)
func (_IGetters *IGettersSession) IsApprovedBySecurityCouncil() (bool, error) {
	return _IGetters.Contract.IsApprovedBySecurityCouncil(&_IGetters.CallOpts)
}

// IsApprovedBySecurityCouncil is a free data retrieval call binding the contract method 0x3db920ce.
//
// Solidity: function isApprovedBySecurityCouncil() view returns(bool)
func (_IGetters *IGettersCallerSession) IsApprovedBySecurityCouncil() (bool, error) {
	return _IGetters.Contract.IsApprovedBySecurityCouncil(&_IGetters.CallOpts)
}

// IsDiamondStorageFrozen is a free data retrieval call binding the contract method 0x29b98c67.
//
// Solidity: function isDiamondStorageFrozen() view returns(bool)
func (_IGetters *IGettersCaller) IsDiamondStorageFrozen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "isDiamondStorageFrozen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDiamondStorageFrozen is a free data retrieval call binding the contract method 0x29b98c67.
//
// Solidity: function isDiamondStorageFrozen() view returns(bool)
func (_IGetters *IGettersSession) IsDiamondStorageFrozen() (bool, error) {
	return _IGetters.Contract.IsDiamondStorageFrozen(&_IGetters.CallOpts)
}

// IsDiamondStorageFrozen is a free data retrieval call binding the contract method 0x29b98c67.
//
// Solidity: function isDiamondStorageFrozen() view returns(bool)
func (_IGetters *IGettersCallerSession) IsDiamondStorageFrozen() (bool, error) {
	return _IGetters.Contract.IsDiamondStorageFrozen(&_IGetters.CallOpts)
}

// IsEthWithdrawalFinalized is a free data retrieval call binding the contract method 0xbd7c5412.
//
// Solidity: function isEthWithdrawalFinalized(uint256 _l2BlockNumber, uint256 _l2MessageIndex) view returns(bool)
func (_IGetters *IGettersCaller) IsEthWithdrawalFinalized(opts *bind.CallOpts, _l2BlockNumber *big.Int, _l2MessageIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "isEthWithdrawalFinalized", _l2BlockNumber, _l2MessageIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEthWithdrawalFinalized is a free data retrieval call binding the contract method 0xbd7c5412.
//
// Solidity: function isEthWithdrawalFinalized(uint256 _l2BlockNumber, uint256 _l2MessageIndex) view returns(bool)
func (_IGetters *IGettersSession) IsEthWithdrawalFinalized(_l2BlockNumber *big.Int, _l2MessageIndex *big.Int) (bool, error) {
	return _IGetters.Contract.IsEthWithdrawalFinalized(&_IGetters.CallOpts, _l2BlockNumber, _l2MessageIndex)
}

// IsEthWithdrawalFinalized is a free data retrieval call binding the contract method 0xbd7c5412.
//
// Solidity: function isEthWithdrawalFinalized(uint256 _l2BlockNumber, uint256 _l2MessageIndex) view returns(bool)
func (_IGetters *IGettersCallerSession) IsEthWithdrawalFinalized(_l2BlockNumber *big.Int, _l2MessageIndex *big.Int) (bool, error) {
	return _IGetters.Contract.IsEthWithdrawalFinalized(&_IGetters.CallOpts, _l2BlockNumber, _l2MessageIndex)
}

// IsFacetFreezable is a free data retrieval call binding the contract method 0xc3bbd2d7.
//
// Solidity: function isFacetFreezable(address _facet) view returns(bool isFreezable)
func (_IGetters *IGettersCaller) IsFacetFreezable(opts *bind.CallOpts, _facet common.Address) (bool, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "isFacetFreezable", _facet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFacetFreezable is a free data retrieval call binding the contract method 0xc3bbd2d7.
//
// Solidity: function isFacetFreezable(address _facet) view returns(bool isFreezable)
func (_IGetters *IGettersSession) IsFacetFreezable(_facet common.Address) (bool, error) {
	return _IGetters.Contract.IsFacetFreezable(&_IGetters.CallOpts, _facet)
}

// IsFacetFreezable is a free data retrieval call binding the contract method 0xc3bbd2d7.
//
// Solidity: function isFacetFreezable(address _facet) view returns(bool isFreezable)
func (_IGetters *IGettersCallerSession) IsFacetFreezable(_facet common.Address) (bool, error) {
	return _IGetters.Contract.IsFacetFreezable(&_IGetters.CallOpts, _facet)
}

// IsFunctionFreezable is a free data retrieval call binding the contract method 0xe81e0ba1.
//
// Solidity: function isFunctionFreezable(bytes4 _selector) view returns(bool)
func (_IGetters *IGettersCaller) IsFunctionFreezable(opts *bind.CallOpts, _selector [4]byte) (bool, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "isFunctionFreezable", _selector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFunctionFreezable is a free data retrieval call binding the contract method 0xe81e0ba1.
//
// Solidity: function isFunctionFreezable(bytes4 _selector) view returns(bool)
func (_IGetters *IGettersSession) IsFunctionFreezable(_selector [4]byte) (bool, error) {
	return _IGetters.Contract.IsFunctionFreezable(&_IGetters.CallOpts, _selector)
}

// IsFunctionFreezable is a free data retrieval call binding the contract method 0xe81e0ba1.
//
// Solidity: function isFunctionFreezable(bytes4 _selector) view returns(bool)
func (_IGetters *IGettersCallerSession) IsFunctionFreezable(_selector [4]byte) (bool, error) {
	return _IGetters.Contract.IsFunctionFreezable(&_IGetters.CallOpts, _selector)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _address) view returns(bool)
func (_IGetters *IGettersCaller) IsValidator(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "isValidator", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _address) view returns(bool)
func (_IGetters *IGettersSession) IsValidator(_address common.Address) (bool, error) {
	return _IGetters.Contract.IsValidator(&_IGetters.CallOpts, _address)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _address) view returns(bool)
func (_IGetters *IGettersCallerSession) IsValidator(_address common.Address) (bool, error) {
	return _IGetters.Contract.IsValidator(&_IGetters.CallOpts, _address)
}

// L2LogsRootHash is a free data retrieval call binding the contract method 0x9cd939e4.
//
// Solidity: function l2LogsRootHash(uint256 _blockNumber) view returns(bytes32 hash)
func (_IGetters *IGettersCaller) L2LogsRootHash(opts *bind.CallOpts, _blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "l2LogsRootHash", _blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L2LogsRootHash is a free data retrieval call binding the contract method 0x9cd939e4.
//
// Solidity: function l2LogsRootHash(uint256 _blockNumber) view returns(bytes32 hash)
func (_IGetters *IGettersSession) L2LogsRootHash(_blockNumber *big.Int) ([32]byte, error) {
	return _IGetters.Contract.L2LogsRootHash(&_IGetters.CallOpts, _blockNumber)
}

// L2LogsRootHash is a free data retrieval call binding the contract method 0x9cd939e4.
//
// Solidity: function l2LogsRootHash(uint256 _blockNumber) view returns(bytes32 hash)
func (_IGetters *IGettersCallerSession) L2LogsRootHash(_blockNumber *big.Int) ([32]byte, error) {
	return _IGetters.Contract.L2LogsRootHash(&_IGetters.CallOpts, _blockNumber)
}

// StoredBlockHash is a free data retrieval call binding the contract method 0x74f4d30d.
//
// Solidity: function storedBlockHash(uint256 _blockNumber) view returns(bytes32)
func (_IGetters *IGettersCaller) StoredBlockHash(opts *bind.CallOpts, _blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "storedBlockHash", _blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StoredBlockHash is a free data retrieval call binding the contract method 0x74f4d30d.
//
// Solidity: function storedBlockHash(uint256 _blockNumber) view returns(bytes32)
func (_IGetters *IGettersSession) StoredBlockHash(_blockNumber *big.Int) ([32]byte, error) {
	return _IGetters.Contract.StoredBlockHash(&_IGetters.CallOpts, _blockNumber)
}

// StoredBlockHash is a free data retrieval call binding the contract method 0x74f4d30d.
//
// Solidity: function storedBlockHash(uint256 _blockNumber) view returns(bytes32)
func (_IGetters *IGettersCallerSession) StoredBlockHash(_blockNumber *big.Int) ([32]byte, error) {
	return _IGetters.Contract.StoredBlockHash(&_IGetters.CallOpts, _blockNumber)
}
