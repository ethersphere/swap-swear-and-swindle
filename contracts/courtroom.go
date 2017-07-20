// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package contracts

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// CaseContractABI is the input ABI used to generate the binding from.
const CaseContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_plaintiff\",\"type\":\"address\"},{\"name\":\"_serviceId\",\"type\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\"},{\"name\":\"expiery\",\"type\":\"uint256\"}],\"name\":\"newClaim\",\"outputs\":[{\"name\":\"id\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"isValid\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"resolveClaim\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getClaim\",\"outputs\":[{\"name\":\"plaintiff\",\"type\":\"address\"},{\"name\":\"serviceId\",\"type\":\"bytes32\"},{\"name\":\"expiery\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"}]"

// CaseContractBin is the compiled bytecode used for deploying new contracts.
const CaseContractBin = `0x6060604052341561000f57600080fd5b5b5b60008054600160a060020a03191633600160a060020a03161790555b5b5b6105328061003e6000396000f3006060604052361561008b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631f0b2d5781146100905780635de28ae0146100cd5780636a938567146100f95780638da5cb5b146101235780638de654ba14610152578063a6f9dae114610182578063c1203cb8146101a3578063c9100bcb146101bb575b600080fd5b341561009b57600080fd5b6100bb600160a060020a036004351660243560ff604435166064356101ff565b60405190815260200160405180910390f35b34156100d857600080fd5b6100e360043561034a565b60405160ff909116815260200160405180910390f35b341561010457600080fd5b61010f600435610365565b604051901515815260200160405180910390f35b341561012e57600080fd5b610136610382565b604051600160a060020a03909116815260200160405180910390f35b341561015d57600080fd5b61010f60043560ff60243516610391565b604051901515815260200160405180910390f35b341561018d57600080fd5b6101a1600160a060020a03600435166103b9565b005b34156101ae57600080fd5b6101a1600435610401565b005b34156101c657600080fd5b6101d160043561049c565b604051600160a060020a03909316835260208301919091526040808301919091526060909101905180910390f35b60006102096104d1565b858542604051600160a060020a03939093166c01000000000000000000000000028352601483019190915260348201526054016040519081900390208152600160008251815260208101919091526040016000206005015460ff16156102725760009150610341565b600160a060020a03861660208201526040810185905260ff84166060820152600160a0820181905242840160808301528190600082518152602081019190915260400160002081518155602082015160018201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039290921691909117905560408201516002820155606082015160038201805460ff191660ff929092169190911790556080820151816004015560a0820151600591909101805460ff191660ff90921691909117905550805191505b50949350505050565b60008181526001602052604090206003015460ff165b919050565b60008181526001602052604090206005015460ff1615155b919050565b600054600160a060020a031681565b6000828152600160208190526040909120600301805460ff191660ff84161790555b92915050565b60005433600160a060020a039081169116146103d457600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b50565b600081815260016020526040902054158061042e575060008181526001602052604090206003015460ff16155b1561043857600080fd5b60008181526001602081905260408220828155908101805473ffffffffffffffffffffffffffffffffffffffff19169055600501805460ff191690555b6000828152600160205260409020600301805460ff191660ff929092169190911790555b50565b6000818152600160208190526040909120908101546002820154600490920154600160a060020a0390911691905b9193909250565b60c06040519081016040908152600080835260208301819052908201819052606082018190526080820181905260a0820152905600a165627a7a72305820f725771df2d8af12a9c94b20e2be72b17a5c3830de04c218691af31f1196e6f20029`

// DeployCaseContract deploys a new Ethereum contract, binding an instance of CaseContract to it.
func DeployCaseContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CaseContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CaseContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CaseContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CaseContract{CaseContractCaller: CaseContractCaller{contract: contract}, CaseContractTransactor: CaseContractTransactor{contract: contract}}, nil
}

// CaseContract is an auto generated Go binding around an Ethereum contract.
type CaseContract struct {
	CaseContractCaller     // Read-only binding to the contract
	CaseContractTransactor // Write-only binding to the contract
}

// CaseContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type CaseContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CaseContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CaseContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CaseContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CaseContractSession struct {
	Contract     *CaseContract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CaseContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CaseContractCallerSession struct {
	Contract *CaseContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CaseContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CaseContractTransactorSession struct {
	Contract     *CaseContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CaseContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type CaseContractRaw struct {
	Contract *CaseContract // Generic contract binding to access the raw methods on
}

// CaseContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CaseContractCallerRaw struct {
	Contract *CaseContractCaller // Generic read-only contract binding to access the raw methods on
}

// CaseContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CaseContractTransactorRaw struct {
	Contract *CaseContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCaseContract creates a new instance of CaseContract, bound to a specific deployed contract.
func NewCaseContract(address common.Address, backend bind.ContractBackend) (*CaseContract, error) {
	contract, err := bindCaseContract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CaseContract{CaseContractCaller: CaseContractCaller{contract: contract}, CaseContractTransactor: CaseContractTransactor{contract: contract}}, nil
}

// NewCaseContractCaller creates a new read-only instance of CaseContract, bound to a specific deployed contract.
func NewCaseContractCaller(address common.Address, caller bind.ContractCaller) (*CaseContractCaller, error) {
	contract, err := bindCaseContract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &CaseContractCaller{contract: contract}, nil
}

// NewCaseContractTransactor creates a new write-only instance of CaseContract, bound to a specific deployed contract.
func NewCaseContractTransactor(address common.Address, transactor bind.ContractTransactor) (*CaseContractTransactor, error) {
	contract, err := bindCaseContract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &CaseContractTransactor{contract: contract}, nil
}

// bindCaseContract binds a generic wrapper to an already deployed contract.
func bindCaseContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CaseContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CaseContract *CaseContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CaseContract.Contract.CaseContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CaseContract *CaseContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CaseContract.Contract.CaseContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CaseContract *CaseContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CaseContract.Contract.CaseContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CaseContract *CaseContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CaseContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CaseContract *CaseContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CaseContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CaseContract *CaseContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CaseContract.Contract.contract.Transact(opts, method, params...)
}

// GetStatus is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(id bytes32) constant returns(status uint8)
func (_CaseContract *CaseContractCaller) GetStatus(opts *bind.CallOpts, id [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _CaseContract.contract.Call(opts, out, "getStatus", id)
	return *ret0, err
}

// GetStatus is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(id bytes32) constant returns(status uint8)
func (_CaseContract *CaseContractSession) GetStatus(id [32]byte) (uint8, error) {
	return _CaseContract.Contract.GetStatus(&_CaseContract.CallOpts, id)
}

// GetStatus is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(id bytes32) constant returns(status uint8)
func (_CaseContract *CaseContractCallerSession) GetStatus(id [32]byte) (uint8, error) {
	return _CaseContract.Contract.GetStatus(&_CaseContract.CallOpts, id)
}

// IsValid is a free data retrieval call binding the contract method 0x6a938567.
//
// Solidity: function isValid(id bytes32) constant returns(bool)
func (_CaseContract *CaseContractCaller) IsValid(opts *bind.CallOpts, id [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CaseContract.contract.Call(opts, out, "isValid", id)
	return *ret0, err
}

// IsValid is a free data retrieval call binding the contract method 0x6a938567.
//
// Solidity: function isValid(id bytes32) constant returns(bool)
func (_CaseContract *CaseContractSession) IsValid(id [32]byte) (bool, error) {
	return _CaseContract.Contract.IsValid(&_CaseContract.CallOpts, id)
}

// IsValid is a free data retrieval call binding the contract method 0x6a938567.
//
// Solidity: function isValid(id bytes32) constant returns(bool)
func (_CaseContract *CaseContractCallerSession) IsValid(id [32]byte) (bool, error) {
	return _CaseContract.Contract.IsValid(&_CaseContract.CallOpts, id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CaseContract *CaseContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CaseContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CaseContract *CaseContractSession) Owner() (common.Address, error) {
	return _CaseContract.Contract.Owner(&_CaseContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CaseContract *CaseContractCallerSession) Owner() (common.Address, error) {
	return _CaseContract.Contract.Owner(&_CaseContract.CallOpts)
}

// SetStatus is a free data retrieval call binding the contract method 0x8de654ba.
//
// Solidity: function setStatus(id bytes32, status uint8) constant returns(bool)
func (_CaseContract *CaseContractCaller) SetStatus(opts *bind.CallOpts, id [32]byte, status uint8) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CaseContract.contract.Call(opts, out, "setStatus", id, status)
	return *ret0, err
}

// SetStatus is a free data retrieval call binding the contract method 0x8de654ba.
//
// Solidity: function setStatus(id bytes32, status uint8) constant returns(bool)
func (_CaseContract *CaseContractSession) SetStatus(id [32]byte, status uint8) (bool, error) {
	return _CaseContract.Contract.SetStatus(&_CaseContract.CallOpts, id, status)
}

// SetStatus is a free data retrieval call binding the contract method 0x8de654ba.
//
// Solidity: function setStatus(id bytes32, status uint8) constant returns(bool)
func (_CaseContract *CaseContractCallerSession) SetStatus(id [32]byte, status uint8) (bool, error) {
	return _CaseContract.Contract.SetStatus(&_CaseContract.CallOpts, id, status)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_CaseContract *CaseContractTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _CaseContract.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_CaseContract *CaseContractSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _CaseContract.Contract.ChangeOwner(&_CaseContract.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_CaseContract *CaseContractTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _CaseContract.Contract.ChangeOwner(&_CaseContract.TransactOpts, _newOwner)
}

// GetClaim is a paid mutator transaction binding the contract method 0xc9100bcb.
//
// Solidity: function getClaim(_id bytes32) returns(plaintiff address, serviceId bytes32, expiery uint256)
func (_CaseContract *CaseContractTransactor) GetClaim(opts *bind.TransactOpts, _id [32]byte) (*types.Transaction, error) {
	return _CaseContract.contract.Transact(opts, "getClaim", _id)
}

// GetClaim is a paid mutator transaction binding the contract method 0xc9100bcb.
//
// Solidity: function getClaim(_id bytes32) returns(plaintiff address, serviceId bytes32, expiery uint256)
func (_CaseContract *CaseContractSession) GetClaim(_id [32]byte) (*types.Transaction, error) {
	return _CaseContract.Contract.GetClaim(&_CaseContract.TransactOpts, _id)
}

// GetClaim is a paid mutator transaction binding the contract method 0xc9100bcb.
//
// Solidity: function getClaim(_id bytes32) returns(plaintiff address, serviceId bytes32, expiery uint256)
func (_CaseContract *CaseContractTransactorSession) GetClaim(_id [32]byte) (*types.Transaction, error) {
	return _CaseContract.Contract.GetClaim(&_CaseContract.TransactOpts, _id)
}

// NewClaim is a paid mutator transaction binding the contract method 0x1f0b2d57.
//
// Solidity: function newClaim(_plaintiff address, _serviceId bytes32, status uint8, expiery uint256) returns(id bytes32)
func (_CaseContract *CaseContractTransactor) NewClaim(opts *bind.TransactOpts, _plaintiff common.Address, _serviceId [32]byte, status uint8, expiery *big.Int) (*types.Transaction, error) {
	return _CaseContract.contract.Transact(opts, "newClaim", _plaintiff, _serviceId, status, expiery)
}

// NewClaim is a paid mutator transaction binding the contract method 0x1f0b2d57.
//
// Solidity: function newClaim(_plaintiff address, _serviceId bytes32, status uint8, expiery uint256) returns(id bytes32)
func (_CaseContract *CaseContractSession) NewClaim(_plaintiff common.Address, _serviceId [32]byte, status uint8, expiery *big.Int) (*types.Transaction, error) {
	return _CaseContract.Contract.NewClaim(&_CaseContract.TransactOpts, _plaintiff, _serviceId, status, expiery)
}

// NewClaim is a paid mutator transaction binding the contract method 0x1f0b2d57.
//
// Solidity: function newClaim(_plaintiff address, _serviceId bytes32, status uint8, expiery uint256) returns(id bytes32)
func (_CaseContract *CaseContractTransactorSession) NewClaim(_plaintiff common.Address, _serviceId [32]byte, status uint8, expiery *big.Int) (*types.Transaction, error) {
	return _CaseContract.Contract.NewClaim(&_CaseContract.TransactOpts, _plaintiff, _serviceId, status, expiery)
}

// ResolveClaim is a paid mutator transaction binding the contract method 0xc1203cb8.
//
// Solidity: function resolveClaim(_id bytes32) returns()
func (_CaseContract *CaseContractTransactor) ResolveClaim(opts *bind.TransactOpts, _id [32]byte) (*types.Transaction, error) {
	return _CaseContract.contract.Transact(opts, "resolveClaim", _id)
}

// ResolveClaim is a paid mutator transaction binding the contract method 0xc1203cb8.
//
// Solidity: function resolveClaim(_id bytes32) returns()
func (_CaseContract *CaseContractSession) ResolveClaim(_id [32]byte) (*types.Transaction, error) {
	return _CaseContract.Contract.ResolveClaim(&_CaseContract.TransactOpts, _id)
}

// ResolveClaim is a paid mutator transaction binding the contract method 0xc1203cb8.
//
// Solidity: function resolveClaim(_id bytes32) returns()
func (_CaseContract *CaseContractTransactorSession) ResolveClaim(_id [32]byte) (*types.Transaction, error) {
	return _CaseContract.Contract.ResolveClaim(&_CaseContract.TransactOpts, _id)
}

// OwnedABI is the input ABI used to generate the binding from.
const OwnedABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"}]"

// OwnedBin is the compiled bytecode used for deploying new contracts.
const OwnedBin = `0x6060604052341561000f57600080fd5b5b60008054600160a060020a03191633600160a060020a03161790555b5b61015c8061003c6000396000f300606060405263ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b8114610048578063a6f9dae114610084575b600080fd5b341561005357600080fd5b61005b6100b2565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561008f57600080fd5b6100b073ffffffffffffffffffffffffffffffffffffffff600435166100ce565b005b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6000543373ffffffffffffffffffffffffffffffffffffffff9081169116146100f657600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff83161790555b5b505600a165627a7a72305820a3d8a5991d5f81a82353bdb05d99ca47797b209198df16a302cf47aed9049c030029`

// DeployOwned deploys a new Ethereum contract, binding an instance of Owned to it.
func DeployOwned(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Owned, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnedBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}}, nil
}

// Owned is an auto generated Go binding around an Ethereum contract.
type Owned struct {
	OwnedCaller     // Read-only binding to the contract
	OwnedTransactor // Write-only binding to the contract
}

// OwnedCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnedSession struct {
	Contract     *Owned            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnedCallerSession struct {
	Contract *OwnedCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OwnedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnedTransactorSession struct {
	Contract     *OwnedTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnedRaw struct {
	Contract *Owned // Generic contract binding to access the raw methods on
}

// OwnedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnedCallerRaw struct {
	Contract *OwnedCaller // Generic read-only contract binding to access the raw methods on
}

// OwnedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnedTransactorRaw struct {
	Contract *OwnedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwned creates a new instance of Owned, bound to a specific deployed contract.
func NewOwned(address common.Address, backend bind.ContractBackend) (*Owned, error) {
	contract, err := bindOwned(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}}, nil
}

// NewOwnedCaller creates a new read-only instance of Owned, bound to a specific deployed contract.
func NewOwnedCaller(address common.Address, caller bind.ContractCaller) (*OwnedCaller, error) {
	contract, err := bindOwned(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &OwnedCaller{contract: contract}, nil
}

// NewOwnedTransactor creates a new write-only instance of Owned, bound to a specific deployed contract.
func NewOwnedTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnedTransactor, error) {
	contract, err := bindOwned(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &OwnedTransactor{contract: contract}, nil
}

// bindOwned binds a generic wrapper to an already deployed contract.
func bindOwned(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.OwnedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Owned.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedSession) Owner() (common.Address, error) {
	return _Owned.Contract.Owner(&_Owned.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedCallerSession) Owner() (common.Address, error) {
	return _Owned.Contract.Owner(&_Owned.CallOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_Owned *OwnedTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Owned.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_Owned *OwnedSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Owned.Contract.ChangeOwner(&_Owned.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_Owned *OwnedTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Owned.Contract.ChangeOwner(&_Owned.TransactOpts, _newOwner)
}

// SampleTokenABI is the input ABI used to generate the binding from.
const SampleTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"beneficiary\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"ref\",\"type\":\"bytes32\"}],\"name\":\"createTokens\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"ref\",\"type\":\"bytes32\"}],\"name\":\"TokenMined\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// SampleTokenBin is the compiled bytecode used for deploying new contracts.
const SampleTokenBin = `0x6060604052341561000f57600080fd5b60405160208061069d833981016040528080519150505b5b60038054600160a060020a03191633600160a060020a03161790555b600160a060020a03331660009081526001602052604081208290558190555b505b61062a806100736000396000f300606060405236156100965763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461009b57806318160ddd146100d157806323b872dd146100f657806370a08231146101325780638da5cb5b14610163578063a3aff92314610192578063a6f9dae1146101b9578063a9059cbb146101da578063dd62ed3e14610210575b600080fd5b34156100a657600080fd5b6100bd600160a060020a0360043516602435610247565b604051901515815260200160405180910390f35b34156100dc57600080fd5b6100e46102b4565b60405190815260200160405180910390f35b341561010157600080fd5b6100bd600160a060020a03600435811690602435166044356102ba565b604051901515815260200160405180910390f35b341561013d57600080fd5b6100e4600160a060020a03600435166102fd565b60405190815260200160405180910390f35b341561016e57600080fd5b61017661031c565b604051600160a060020a03909116815260200160405180910390f35b341561019d57600080fd5b6101b7600160a060020a036004351660243560443561032b565b005b34156101c457600080fd5b6101b7600160a060020a03600435166103e8565b005b34156101e557600080fd5b6100bd600160a060020a0360043516602435610430565b604051901515815260200160405180910390f35b341561021b57600080fd5b6100e4600160a060020a03600435811690602435166104d7565b60405190815260200160405180910390f35b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a35060015b92915050565b60005481565b600160a060020a0380841660009081526002602090815260408083203390941683529290529081208054830190556102f3848484610504565b90505b9392505050565b600160a060020a0381166000908152600160205260409020545b919050565b600354600160a060020a031681565b60035433600160a060020a0390811691161461034657600080fd5b600160a060020a038316600081815260016020526040808220805486019055815485019091558291907fdfb81fb379557413b0a951b4d7bf7a9df393801d8c539d5e201d6a8daeb913b99085905190815260200160405180910390a382600160a060020a031660007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a35b5b505050565b60035433600160a060020a0390811691161461040357600080fd5b6003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b50565b600160a060020a0333166000908152600160205260408120548290108015906104595750600082115b1561009657600160a060020a033381166000818152600160205260408082208054879003905592861680825290839020805486019055917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060016102ae565b600080fd5b5b92915050565b600160a060020a038083166000908152600260209081526040808320938516835292905220545b92915050565b600160a060020a0383166000908152600160205260408120548290108015906105545750600160a060020a0380851660009081526002602090815260408083203390941683529290522054829010155b80156105605750600082115b1561009657600160a060020a03808416600081815260016020908152604080832080548801905588851680845281842080548990039055600283528184203390961684529490915290819020805486900390559091907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060016102f6565b600080fd5b5b93925050505600a165627a7a7230582001759bcd714022196869a0bf32ff6a0e7da2b22d5163d1d6991673cdd78eeb040029`

// DeploySampleToken deploys a new Ethereum contract, binding an instance of SampleToken to it.
func DeploySampleToken(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int) (common.Address, *types.Transaction, *SampleToken, error) {
	parsed, err := abi.JSON(strings.NewReader(SampleTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SampleTokenBin), backend, initialSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SampleToken{SampleTokenCaller: SampleTokenCaller{contract: contract}, SampleTokenTransactor: SampleTokenTransactor{contract: contract}}, nil
}

// SampleToken is an auto generated Go binding around an Ethereum contract.
type SampleToken struct {
	SampleTokenCaller     // Read-only binding to the contract
	SampleTokenTransactor // Write-only binding to the contract
}

// SampleTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type SampleTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SampleTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SampleTokenSession struct {
	Contract     *SampleToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SampleTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SampleTokenCallerSession struct {
	Contract *SampleTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SampleTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SampleTokenTransactorSession struct {
	Contract     *SampleTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SampleTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type SampleTokenRaw struct {
	Contract *SampleToken // Generic contract binding to access the raw methods on
}

// SampleTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SampleTokenCallerRaw struct {
	Contract *SampleTokenCaller // Generic read-only contract binding to access the raw methods on
}

// SampleTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SampleTokenTransactorRaw struct {
	Contract *SampleTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSampleToken creates a new instance of SampleToken, bound to a specific deployed contract.
func NewSampleToken(address common.Address, backend bind.ContractBackend) (*SampleToken, error) {
	contract, err := bindSampleToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SampleToken{SampleTokenCaller: SampleTokenCaller{contract: contract}, SampleTokenTransactor: SampleTokenTransactor{contract: contract}}, nil
}

// NewSampleTokenCaller creates a new read-only instance of SampleToken, bound to a specific deployed contract.
func NewSampleTokenCaller(address common.Address, caller bind.ContractCaller) (*SampleTokenCaller, error) {
	contract, err := bindSampleToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SampleTokenCaller{contract: contract}, nil
}

// NewSampleTokenTransactor creates a new write-only instance of SampleToken, bound to a specific deployed contract.
func NewSampleTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*SampleTokenTransactor, error) {
	contract, err := bindSampleToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SampleTokenTransactor{contract: contract}, nil
}

// bindSampleToken binds a generic wrapper to an already deployed contract.
func bindSampleToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SampleTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SampleToken *SampleTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SampleToken.Contract.SampleTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SampleToken *SampleTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleToken.Contract.SampleTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SampleToken *SampleTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SampleToken.Contract.SampleTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SampleToken *SampleTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SampleToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SampleToken *SampleTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SampleToken *SampleTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SampleToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_SampleToken *SampleTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SampleToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_SampleToken *SampleTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _SampleToken.Contract.Allowance(&_SampleToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_SampleToken *SampleTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _SampleToken.Contract.Allowance(&_SampleToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_SampleToken *SampleTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SampleToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_SampleToken *SampleTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _SampleToken.Contract.BalanceOf(&_SampleToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_SampleToken *SampleTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _SampleToken.Contract.BalanceOf(&_SampleToken.CallOpts, _owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SampleToken *SampleTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SampleToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SampleToken *SampleTokenSession) Owner() (common.Address, error) {
	return _SampleToken.Contract.Owner(&_SampleToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SampleToken *SampleTokenCallerSession) Owner() (common.Address, error) {
	return _SampleToken.Contract.Owner(&_SampleToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SampleToken *SampleTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SampleToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SampleToken *SampleTokenSession) TotalSupply() (*big.Int, error) {
	return _SampleToken.Contract.TotalSupply(&_SampleToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SampleToken *SampleTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _SampleToken.Contract.TotalSupply(&_SampleToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_SampleToken *SampleTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SampleToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_SampleToken *SampleTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SampleToken.Contract.Approve(&_SampleToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_SampleToken *SampleTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SampleToken.Contract.Approve(&_SampleToken.TransactOpts, _spender, _value)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_SampleToken *SampleTokenTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _SampleToken.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_SampleToken *SampleTokenSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _SampleToken.Contract.ChangeOwner(&_SampleToken.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_SampleToken *SampleTokenTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _SampleToken.Contract.ChangeOwner(&_SampleToken.TransactOpts, _newOwner)
}

// CreateTokens is a paid mutator transaction binding the contract method 0xa3aff923.
//
// Solidity: function createTokens(beneficiary address, amount uint256, ref bytes32) returns()
func (_SampleToken *SampleTokenTransactor) CreateTokens(opts *bind.TransactOpts, beneficiary common.Address, amount *big.Int, ref [32]byte) (*types.Transaction, error) {
	return _SampleToken.contract.Transact(opts, "createTokens", beneficiary, amount, ref)
}

// CreateTokens is a paid mutator transaction binding the contract method 0xa3aff923.
//
// Solidity: function createTokens(beneficiary address, amount uint256, ref bytes32) returns()
func (_SampleToken *SampleTokenSession) CreateTokens(beneficiary common.Address, amount *big.Int, ref [32]byte) (*types.Transaction, error) {
	return _SampleToken.Contract.CreateTokens(&_SampleToken.TransactOpts, beneficiary, amount, ref)
}

// CreateTokens is a paid mutator transaction binding the contract method 0xa3aff923.
//
// Solidity: function createTokens(beneficiary address, amount uint256, ref bytes32) returns()
func (_SampleToken *SampleTokenTransactorSession) CreateTokens(beneficiary common.Address, amount *big.Int, ref [32]byte) (*types.Transaction, error) {
	return _SampleToken.Contract.CreateTokens(&_SampleToken.TransactOpts, beneficiary, amount, ref)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_SampleToken *SampleTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SampleToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_SampleToken *SampleTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SampleToken.Contract.Transfer(&_SampleToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_SampleToken *SampleTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SampleToken.Contract.Transfer(&_SampleToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_SampleToken *SampleTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SampleToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_SampleToken *SampleTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SampleToken.Contract.TransferFrom(&_SampleToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_SampleToken *SampleTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SampleToken.Contract.TransferFrom(&_SampleToken.TransactOpts, _from, _to, _value)
}

// StandardTokenABI is the input ABI used to generate the binding from.
const StandardTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// StandardTokenBin is the compiled bytecode used for deploying new contracts.
const StandardTokenBin = `0x6060604052341561000f57600080fd5b5b61043b8061001f6000396000f300606060405236156100755763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461007a57806318160ddd146100b057806323b872dd146100d557806370a0823114610111578063a9059cbb14610142578063dd62ed3e14610178575b600080fd5b341561008557600080fd5b61009c600160a060020a03600435166024356101af565b604051901515815260200160405180910390f35b34156100bb57600080fd5b6100c361021c565b60405190815260200160405180910390f35b34156100e057600080fd5b61009c600160a060020a0360043581169060243516604435610222565b604051901515815260200160405180910390f35b341561011c57600080fd5b6100c3600160a060020a036004351661031c565b60405190815260200160405180910390f35b341561014d57600080fd5b61009c600160a060020a036004351660243561033b565b604051901515815260200160405180910390f35b341561018357600080fd5b6100c3600160a060020a03600435811690602435166103e2565b60405190815260200160405180910390f35b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a35060015b92915050565b60005481565b600160a060020a0383166000908152600160205260408120548290108015906102725750600160a060020a0380851660009081526002602090815260408083203390941683529290522054829010155b801561027e5750600082115b1561007557600160a060020a03808416600081815260016020908152604080832080548801905588851680845281842080548990039055600283528184203390961684529490915290819020805486900390559091907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a3506001610314565b600080fd5b5b9392505050565b600160a060020a0381166000908152600160205260409020545b919050565b600160a060020a0333166000908152600160205260408120548290108015906103645750600082115b1561007557600160a060020a033381166000818152600160205260408082208054879003905592861680825290839020805486019055917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a3506001610216565b600080fd5b5b92915050565b600160a060020a038083166000908152600260209081526040808320938516835292905220545b929150505600a165627a7a723058207aa3399e12bdf4ac02bb2ba917c98cf5cf42c0aedc0482b41a77141b9504a13a0029`

// DeployStandardToken deploys a new Ethereum contract, binding an instance of StandardToken to it.
func DeployStandardToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StandardToken, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StandardTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}}, nil
}

// StandardToken is an auto generated Go binding around an Ethereum contract.
type StandardToken struct {
	StandardTokenCaller     // Read-only binding to the contract
	StandardTokenTransactor // Write-only binding to the contract
}

// StandardTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type StandardTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StandardTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StandardTokenSession struct {
	Contract     *StandardToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StandardTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StandardTokenCallerSession struct {
	Contract *StandardTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StandardTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StandardTokenTransactorSession struct {
	Contract     *StandardTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StandardTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type StandardTokenRaw struct {
	Contract *StandardToken // Generic contract binding to access the raw methods on
}

// StandardTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StandardTokenCallerRaw struct {
	Contract *StandardTokenCaller // Generic read-only contract binding to access the raw methods on
}

// StandardTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StandardTokenTransactorRaw struct {
	Contract *StandardTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStandardToken creates a new instance of StandardToken, bound to a specific deployed contract.
func NewStandardToken(address common.Address, backend bind.ContractBackend) (*StandardToken, error) {
	contract, err := bindStandardToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}}, nil
}

// NewStandardTokenCaller creates a new read-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenCaller(address common.Address, caller bind.ContractCaller) (*StandardTokenCaller, error) {
	contract, err := bindStandardToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenCaller{contract: contract}, nil
}

// NewStandardTokenTransactor creates a new write-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*StandardTokenTransactor, error) {
	contract, err := bindStandardToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransactor{contract: contract}, nil
}

// bindStandardToken binds a generic wrapper to an already deployed contract.
func bindStandardToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.StandardTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_StandardToken *StandardTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_StandardToken *StandardTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_StandardToken *StandardTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// SwearGameABI is the input ABI used to generate the binding from.
const SwearGameABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"reward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"playerCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_player\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[{\"name\":\"registered\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ids\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"trialTransistions\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"serviceId\",\"type\":\"bytes32\"}],\"name\":\"newCase\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"caseContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"players\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"resumeCase\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_CaseContract\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_trialTransistions\",\"type\":\"address\"},{\"name\":\"_reward\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"constructor\"},{\"payable\":true,\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"decide\",\"type\":\"string\"}],\"name\":\"Decision\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"depositAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"DepositStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"Compensate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"playerId\",\"type\":\"address\"}],\"name\":\"NewPlayer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"playerId\",\"type\":\"address\"}],\"name\":\"PlayerLeftGame\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"plaintiff\",\"type\":\"address\"}],\"name\":\"NewClaimOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"plaintiff\",\"type\":\"address\"}],\"name\":\"NewEvidenceSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"plaintiff\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"ClaimResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Payment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"AdditionalDepositRequired\",\"type\":\"event\"}]"

// SwearGameBin is the compiled bytecode used for deploying new contracts.
const SwearGameBin = `0x6060604052341561000f57600080fd5b6040516080806111a2833981016040528080519190602001805191906020018051919060200180519150505b5b60008054600160a060020a03191633600160a060020a03161790555b60068054600160a060020a03808716600160a060020a031992831617909255600780548684169083161790556008805492851692909116919091179055600281905560006003555b505050505b6110ee806100b46000396000f300606060405236156100b45763ffffffff60e060020a600035041663228cb73381146101a2578063302bcc57146101c75780634420e486146101ec5780635de28ae01461021f5780638da5cb5b1461024b5780639f20bc921461027a578063a6f9dae1146102ae578063c41a3477146102cf578063cc345e0a146102fe578063d0e30db014610328578063d556e8e61461034d578063e2eb41ff1461037c578063e375bebd146103af578063fc0c546a146103d9575b6101a05b600754600080543492600160a060020a03908116926323b872dd9290911690309085906040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b151561013157600080fd5b6102c65a03f1151561014257600080fd5b50505060405180519050151561015757600080fd5b600180548201908190557f3a7e173a9698235104076a0f536c5169527adecd20d0938c7c156ddf776c859190829060405191825260208201526040908101905180910390a15b50565b005b34156101ad57600080fd5b6101b5610408565b60405190815260200160405180910390f35b34156101d257600080fd5b6101b561040e565b60405190815260200160405180910390f35b34156101f757600080fd5b61020b600160a060020a0360043516610414565b604051901515815260200160405180910390f35b341561022a57600080fd5b610235600435610541565b60405160ff909116815260200160405180910390f35b341561025657600080fd5b61025e6105b6565b604051600160a060020a03909116815260200160405180910390f35b341561028557600080fd5b6101b5600160a060020a03600435166024356105c5565b60405190815260200160405180910390f35b34156102b957600080fd5b6101a0600160a060020a03600435166105f7565b005b34156102da57600080fd5b61025e61063f565b604051600160a060020a03909116815260200160405180910390f35b341561030957600080fd5b61020b60043561064e565b604051901515815260200160405180910390f35b341561033357600080fd5b6101b5610826565b60405190815260200160405180910390f35b341561035857600080fd5b61025e61082c565b604051600160a060020a03909116815260200160405180910390f35b341561038757600080fd5b61020b600160a060020a036004351661083b565b604051901515815260200160405180910390f35b34156103ba57600080fd5b61020b600435610850565b604051901515815260200160405180910390f35b34156103e457600080fd5b61025e610900565b604051600160a060020a03909116815260200160405180910390f35b60025481565b60035481565b6000805433600160a060020a0390811691161461043057600080fd5b600160a060020a03821660009081526004602052604090205460ff161561045657600080fd5b600354151561047557600254600154101561047057600080fd5b6104c8565b60025460035460015481151561048757fe5b0410156104c8577fa8ff0fbf8dc82e5914523ed3a828920f55ab93494c59933e8c45d79f8bba70d760015460405190815260200160405180910390a1600080fd5b5b600160a060020a03821660009081526004602052604090819020805460ff191660019081179091556003805490910190557f52e92d4898337244a39bd42674ac561eadfd3959e947deec1c0ab82dd58b5a7590839051600160a060020a03909116815260200160405180910390a15060015b5b919050565b600654600090600160a060020a0316635de28ae083836040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561059457600080fd5b6102c65a03f115156105a557600080fd5b50505060405180519150505b919050565b600054600160a060020a031681565b6005602052816000526040600020818154811015156105e057fe5b906000526020600020900160005b91509150505481565b60005433600160a060020a0390811691161461061257600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b50565b600854600160a060020a031681565b600160a060020a033316600090815260046020526040812054819060ff16151561067757600080fd5b600654600854600160a060020a0391821691631f0b2d579133918791166307e5ae5f6000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156106d057600080fd5b6102c65a03f115156106e157600080fd5b5050506040518051600854909150600160a060020a031663ada016106000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561073457600080fd5b6102c65a03f1151561074557600080fd5b5050506040518051905060006040516020015260405160e060020a63ffffffff8716028152600160a060020a039094166004850152602484019290925260ff1660448301526064820152608401602060405180830381600087803b15156107ab57600080fd5b6102c65a03f115156107bc57600080fd5b50505060405180519150508015156107d75760009150610820565b600160a060020a03331660009081526005602052604090208054600181016107ff8382611077565b916000526020600020900160005b508290555061081b8161090f565b600191505b50919050565b60015481565b600654600160a060020a031681565b60046020526000908152604090205460ff1681565b600160a060020a03331660009081526004602052604081205460ff16151561087757600080fd5b600654600160a060020a0316636a9385678360006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156108c857600080fd5b6102c65a03f115156108d957600080fd5b5050506040518051905015156108ee57600080fd5b6108f78261090f565b5060015b919050565b600754600160a060020a031681565b60065460009081908190819081908190600160a060020a0316635de28ae088836040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561096c57600080fd5b6102c65a03f1151561097d57600080fd5b5050506040518051600654909750600160a060020a0316905063c9100bcb8860006040516060015260405160e060020a63ffffffff84160281526004810191909152602401606060405180830381600087803b15156109db57600080fd5b6102c65a03f115156109ec57600080fd5b5050506040518051906020018051906020018051929750909550909350600090505b60ff168660ff161415610a2057610d88565b5b60005b60ff878116911614610d8857610a3983610d92565b15610abb57600654600160a060020a0316638de654ba8860005b60006040516020015260405160e060020a63ffffffff8516028152600481019290925260ff166024820152604401602060405180830381600087803b1515610a9a57600080fd5b6102c65a03f11515610aab57600080fd5b5050506040518051905050610d88565b600854600160a060020a031663db0615638760006040516020015260405160e060020a63ffffffff841602815260ff9091166004820152602401602060405180830381600087803b1515610b0e57600080fd5b6102c65a03f11515610b1f57600080fd5b5050506040518051925050600160a060020a0382161515610b4957610b42610d9c565b9050610d88565b81600160a060020a031663f243aee5858760006040516020015260405160e060020a63ffffffff85160281526004810192909252600160a060020a03166024820152604401602060405180830381600087803b1515610ba757600080fd5b6102c65a03f11515610bb857600080fd5b50505060405180519150505b60025b816002811115610bd357fe5b1415610bde57610d88565b600854600160a060020a031663d39eb301826002811115610bfb57fe5b8860006040516020015260405160e060020a63ffffffff851602815260ff928316600482015291166024820152604401602060405180830381600087803b1515610c4457600080fd5b6102c65a03f11515610c5557600080fd5b5050506040518051600654909750600160a060020a03169050638de654ba888860006040516020015260405160e060020a63ffffffff8516028152600481019290925260ff166024820152604401602060405180830381600087803b1515610cbc57600080fd5b6102c65a03f11515610ccd57600080fd5b505050604051805150600190505b60ff168660ff161480610cf5575060025b60ff168660ff16145b15610d8257610d05878787610da2565b5060005b600654909650600160a060020a0316638de654ba888860006040516020015260405160e060020a63ffffffff8516028152600481019290925260ff166024820152604401602060405180830381600087803b1515610d6657600080fd5b6102c65a03f11515610d7757600080fd5b505050604051805150505b5b610a20565b5b50505050505050565b428190115b919050565b60025b90565b60008060025b60ff168460ff161415610e19577f2f18444e939f85d10c4f7ad5f7de36bc69f5d99ff0e6614a6accbb2f0118e4778584600087604051938452600160a060020a03909216602084015260408084019190915260ff90911660608301526080909101905180910390a160009150610eea565b610e2283610ef2565b600654909150600160a060020a031663c1203cb88660405160e060020a63ffffffff84160281526004810191909152602401600060405180830381600087803b1515610e6d57600080fd5b6102c65a03f11515610e7e57600080fd5b505050610e8a83610fe5565b7f2f18444e939f85d10c4f7ad5f7de36bc69f5d99ff0e6614a6accbb2f0118e477858460025487604051938452600160a060020a03909216602084015260408084019190915260ff90911660608301526080909101905180910390a18091505b509392505050565b600754600254600091600160a060020a0316906323b872dd9030908590856040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b1515610f6657600080fd5b6102c65a03f11515610f7757600080fd5b5050506040518051915050801515610f8e57600080fd5b6002546001805482900390557f92c2a7173158b7618078365b4ad89fd1f774ae4aa04f39e10b966b47f469d34b908390604051600160a060020a03909216825260208201526040908101905180910390a15b919050565b600160a060020a03811660009081526004602052604090205460ff16151561100c57600080fd5b7f3def0aea61e344e47b98a0523a5825c0919e88f7b33d0ac7f889a7654880f9a581604051600160a060020a03909116815260200160405180910390a1600160a060020a0381166000908152600460205260409020805460ff19169055600380546000190190555b50565b81548183558181151161109b5760008381526020902061109b9181019083016110a1565b5b505050565b610d9f91905b808211156110bb57600081556001016110a7565b5090565b905600a165627a7a72305820a3ef61bc7d31f713d9fcef5f983a8cdbc65055ebeac284a48297f7c3a3f663f00029`

// DeploySwearGame deploys a new Ethereum contract, binding an instance of SwearGame to it.
func DeploySwearGame(auth *bind.TransactOpts, backend bind.ContractBackend, _CaseContract common.Address, _token common.Address, _trialTransistions common.Address, _reward *big.Int) (common.Address, *types.Transaction, *SwearGame, error) {
	parsed, err := abi.JSON(strings.NewReader(SwearGameABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SwearGameBin), backend, _CaseContract, _token, _trialTransistions, _reward)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SwearGame{SwearGameCaller: SwearGameCaller{contract: contract}, SwearGameTransactor: SwearGameTransactor{contract: contract}}, nil
}

// SwearGame is an auto generated Go binding around an Ethereum contract.
type SwearGame struct {
	SwearGameCaller     // Read-only binding to the contract
	SwearGameTransactor // Write-only binding to the contract
}

// SwearGameCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwearGameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwearGameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwearGameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwearGameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwearGameSession struct {
	Contract     *SwearGame        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwearGameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwearGameCallerSession struct {
	Contract *SwearGameCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SwearGameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwearGameTransactorSession struct {
	Contract     *SwearGameTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SwearGameRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwearGameRaw struct {
	Contract *SwearGame // Generic contract binding to access the raw methods on
}

// SwearGameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwearGameCallerRaw struct {
	Contract *SwearGameCaller // Generic read-only contract binding to access the raw methods on
}

// SwearGameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwearGameTransactorRaw struct {
	Contract *SwearGameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwearGame creates a new instance of SwearGame, bound to a specific deployed contract.
func NewSwearGame(address common.Address, backend bind.ContractBackend) (*SwearGame, error) {
	contract, err := bindSwearGame(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwearGame{SwearGameCaller: SwearGameCaller{contract: contract}, SwearGameTransactor: SwearGameTransactor{contract: contract}}, nil
}

// NewSwearGameCaller creates a new read-only instance of SwearGame, bound to a specific deployed contract.
func NewSwearGameCaller(address common.Address, caller bind.ContractCaller) (*SwearGameCaller, error) {
	contract, err := bindSwearGame(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SwearGameCaller{contract: contract}, nil
}

// NewSwearGameTransactor creates a new write-only instance of SwearGame, bound to a specific deployed contract.
func NewSwearGameTransactor(address common.Address, transactor bind.ContractTransactor) (*SwearGameTransactor, error) {
	contract, err := bindSwearGame(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SwearGameTransactor{contract: contract}, nil
}

// bindSwearGame binds a generic wrapper to an already deployed contract.
func bindSwearGame(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwearGameABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwearGame *SwearGameRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SwearGame.Contract.SwearGameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwearGame *SwearGameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwearGame.Contract.SwearGameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwearGame *SwearGameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwearGame.Contract.SwearGameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwearGame *SwearGameCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SwearGame.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwearGame *SwearGameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwearGame.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwearGame *SwearGameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwearGame.Contract.contract.Transact(opts, method, params...)
}

// CaseContract is a free data retrieval call binding the contract method 0xd556e8e6.
//
// Solidity: function caseContract() constant returns(address)
func (_SwearGame *SwearGameCaller) CaseContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SwearGame.contract.Call(opts, out, "caseContract")
	return *ret0, err
}

// CaseContract is a free data retrieval call binding the contract method 0xd556e8e6.
//
// Solidity: function caseContract() constant returns(address)
func (_SwearGame *SwearGameSession) CaseContract() (common.Address, error) {
	return _SwearGame.Contract.CaseContract(&_SwearGame.CallOpts)
}

// CaseContract is a free data retrieval call binding the contract method 0xd556e8e6.
//
// Solidity: function caseContract() constant returns(address)
func (_SwearGame *SwearGameCallerSession) CaseContract() (common.Address, error) {
	return _SwearGame.Contract.CaseContract(&_SwearGame.CallOpts)
}

// Deposit is a free data retrieval call binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() constant returns(uint256)
func (_SwearGame *SwearGameCaller) Deposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SwearGame.contract.Call(opts, out, "deposit")
	return *ret0, err
}

// Deposit is a free data retrieval call binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() constant returns(uint256)
func (_SwearGame *SwearGameSession) Deposit() (*big.Int, error) {
	return _SwearGame.Contract.Deposit(&_SwearGame.CallOpts)
}

// Deposit is a free data retrieval call binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() constant returns(uint256)
func (_SwearGame *SwearGameCallerSession) Deposit() (*big.Int, error) {
	return _SwearGame.Contract.Deposit(&_SwearGame.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(id bytes32) constant returns(status uint8)
func (_SwearGame *SwearGameCaller) GetStatus(opts *bind.CallOpts, id [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _SwearGame.contract.Call(opts, out, "getStatus", id)
	return *ret0, err
}

// GetStatus is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(id bytes32) constant returns(status uint8)
func (_SwearGame *SwearGameSession) GetStatus(id [32]byte) (uint8, error) {
	return _SwearGame.Contract.GetStatus(&_SwearGame.CallOpts, id)
}

// GetStatus is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(id bytes32) constant returns(status uint8)
func (_SwearGame *SwearGameCallerSession) GetStatus(id [32]byte) (uint8, error) {
	return _SwearGame.Contract.GetStatus(&_SwearGame.CallOpts, id)
}

// Ids is a free data retrieval call binding the contract method 0x9f20bc92.
//
// Solidity: function ids( address,  uint256) constant returns(bytes32)
func (_SwearGame *SwearGameCaller) Ids(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _SwearGame.contract.Call(opts, out, "ids", arg0, arg1)
	return *ret0, err
}

// Ids is a free data retrieval call binding the contract method 0x9f20bc92.
//
// Solidity: function ids( address,  uint256) constant returns(bytes32)
func (_SwearGame *SwearGameSession) Ids(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _SwearGame.Contract.Ids(&_SwearGame.CallOpts, arg0, arg1)
}

// Ids is a free data retrieval call binding the contract method 0x9f20bc92.
//
// Solidity: function ids( address,  uint256) constant returns(bytes32)
func (_SwearGame *SwearGameCallerSession) Ids(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _SwearGame.Contract.Ids(&_SwearGame.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SwearGame *SwearGameCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SwearGame.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SwearGame *SwearGameSession) Owner() (common.Address, error) {
	return _SwearGame.Contract.Owner(&_SwearGame.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SwearGame *SwearGameCallerSession) Owner() (common.Address, error) {
	return _SwearGame.Contract.Owner(&_SwearGame.CallOpts)
}

// PlayerCount is a free data retrieval call binding the contract method 0x302bcc57.
//
// Solidity: function playerCount() constant returns(uint256)
func (_SwearGame *SwearGameCaller) PlayerCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SwearGame.contract.Call(opts, out, "playerCount")
	return *ret0, err
}

// PlayerCount is a free data retrieval call binding the contract method 0x302bcc57.
//
// Solidity: function playerCount() constant returns(uint256)
func (_SwearGame *SwearGameSession) PlayerCount() (*big.Int, error) {
	return _SwearGame.Contract.PlayerCount(&_SwearGame.CallOpts)
}

// PlayerCount is a free data retrieval call binding the contract method 0x302bcc57.
//
// Solidity: function playerCount() constant returns(uint256)
func (_SwearGame *SwearGameCallerSession) PlayerCount() (*big.Int, error) {
	return _SwearGame.Contract.PlayerCount(&_SwearGame.CallOpts)
}

// Players is a free data retrieval call binding the contract method 0xe2eb41ff.
//
// Solidity: function players( address) constant returns(bool)
func (_SwearGame *SwearGameCaller) Players(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SwearGame.contract.Call(opts, out, "players", arg0)
	return *ret0, err
}

// Players is a free data retrieval call binding the contract method 0xe2eb41ff.
//
// Solidity: function players( address) constant returns(bool)
func (_SwearGame *SwearGameSession) Players(arg0 common.Address) (bool, error) {
	return _SwearGame.Contract.Players(&_SwearGame.CallOpts, arg0)
}

// Players is a free data retrieval call binding the contract method 0xe2eb41ff.
//
// Solidity: function players( address) constant returns(bool)
func (_SwearGame *SwearGameCallerSession) Players(arg0 common.Address) (bool, error) {
	return _SwearGame.Contract.Players(&_SwearGame.CallOpts, arg0)
}

// Reward is a free data retrieval call binding the contract method 0x228cb733.
//
// Solidity: function reward() constant returns(uint256)
func (_SwearGame *SwearGameCaller) Reward(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SwearGame.contract.Call(opts, out, "reward")
	return *ret0, err
}

// Reward is a free data retrieval call binding the contract method 0x228cb733.
//
// Solidity: function reward() constant returns(uint256)
func (_SwearGame *SwearGameSession) Reward() (*big.Int, error) {
	return _SwearGame.Contract.Reward(&_SwearGame.CallOpts)
}

// Reward is a free data retrieval call binding the contract method 0x228cb733.
//
// Solidity: function reward() constant returns(uint256)
func (_SwearGame *SwearGameCallerSession) Reward() (*big.Int, error) {
	return _SwearGame.Contract.Reward(&_SwearGame.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_SwearGame *SwearGameCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SwearGame.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_SwearGame *SwearGameSession) Token() (common.Address, error) {
	return _SwearGame.Contract.Token(&_SwearGame.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_SwearGame *SwearGameCallerSession) Token() (common.Address, error) {
	return _SwearGame.Contract.Token(&_SwearGame.CallOpts)
}

// TrialTransistions is a free data retrieval call binding the contract method 0xc41a3477.
//
// Solidity: function trialTransistions() constant returns(address)
func (_SwearGame *SwearGameCaller) TrialTransistions(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SwearGame.contract.Call(opts, out, "trialTransistions")
	return *ret0, err
}

// TrialTransistions is a free data retrieval call binding the contract method 0xc41a3477.
//
// Solidity: function trialTransistions() constant returns(address)
func (_SwearGame *SwearGameSession) TrialTransistions() (common.Address, error) {
	return _SwearGame.Contract.TrialTransistions(&_SwearGame.CallOpts)
}

// TrialTransistions is a free data retrieval call binding the contract method 0xc41a3477.
//
// Solidity: function trialTransistions() constant returns(address)
func (_SwearGame *SwearGameCallerSession) TrialTransistions() (common.Address, error) {
	return _SwearGame.Contract.TrialTransistions(&_SwearGame.CallOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_SwearGame *SwearGameTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _SwearGame.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_SwearGame *SwearGameSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _SwearGame.Contract.ChangeOwner(&_SwearGame.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_SwearGame *SwearGameTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _SwearGame.Contract.ChangeOwner(&_SwearGame.TransactOpts, _newOwner)
}

// NewCase is a paid mutator transaction binding the contract method 0xcc345e0a.
//
// Solidity: function newCase(serviceId bytes32) returns(bool)
func (_SwearGame *SwearGameTransactor) NewCase(opts *bind.TransactOpts, serviceId [32]byte) (*types.Transaction, error) {
	return _SwearGame.contract.Transact(opts, "newCase", serviceId)
}

// NewCase is a paid mutator transaction binding the contract method 0xcc345e0a.
//
// Solidity: function newCase(serviceId bytes32) returns(bool)
func (_SwearGame *SwearGameSession) NewCase(serviceId [32]byte) (*types.Transaction, error) {
	return _SwearGame.Contract.NewCase(&_SwearGame.TransactOpts, serviceId)
}

// NewCase is a paid mutator transaction binding the contract method 0xcc345e0a.
//
// Solidity: function newCase(serviceId bytes32) returns(bool)
func (_SwearGame *SwearGameTransactorSession) NewCase(serviceId [32]byte) (*types.Transaction, error) {
	return _SwearGame.Contract.NewCase(&_SwearGame.TransactOpts, serviceId)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(_player address) returns(registered bool)
func (_SwearGame *SwearGameTransactor) Register(opts *bind.TransactOpts, _player common.Address) (*types.Transaction, error) {
	return _SwearGame.contract.Transact(opts, "register", _player)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(_player address) returns(registered bool)
func (_SwearGame *SwearGameSession) Register(_player common.Address) (*types.Transaction, error) {
	return _SwearGame.Contract.Register(&_SwearGame.TransactOpts, _player)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(_player address) returns(registered bool)
func (_SwearGame *SwearGameTransactorSession) Register(_player common.Address) (*types.Transaction, error) {
	return _SwearGame.Contract.Register(&_SwearGame.TransactOpts, _player)
}

// ResumeCase is a paid mutator transaction binding the contract method 0xe375bebd.
//
// Solidity: function resumeCase(id bytes32) returns(bool)
func (_SwearGame *SwearGameTransactor) ResumeCase(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _SwearGame.contract.Transact(opts, "resumeCase", id)
}

// ResumeCase is a paid mutator transaction binding the contract method 0xe375bebd.
//
// Solidity: function resumeCase(id bytes32) returns(bool)
func (_SwearGame *SwearGameSession) ResumeCase(id [32]byte) (*types.Transaction, error) {
	return _SwearGame.Contract.ResumeCase(&_SwearGame.TransactOpts, id)
}

// ResumeCase is a paid mutator transaction binding the contract method 0xe375bebd.
//
// Solidity: function resumeCase(id bytes32) returns(bool)
func (_SwearGame *SwearGameTransactorSession) ResumeCase(id [32]byte) (*types.Transaction, error) {
	return _SwearGame.Contract.ResumeCase(&_SwearGame.TransactOpts, id)
}

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// TokenBin is the compiled bytecode used for deploying new contracts.
const TokenBin = `0x`

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}}, nil
}

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Token *TokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Token *TokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Token *TokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Token *TokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Token *TokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Token *TokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _value)
}

// TrialTransistionsAbstractABI is the input ABI used to generate the binding from.
const TrialTransistionsAbstractABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"getInitialStatus\",\"outputs\":[{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getTrialExpiry\",\"outputs\":[{\"name\":\"expiery\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"witnessState\",\"type\":\"uint8\"},{\"name\":\"trialStatus\",\"type\":\"uint8\"}],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"trialStatus\",\"type\":\"uint8\"}],\"name\":\"getWitness\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"}]"

// TrialTransistionsAbstractBin is the compiled bytecode used for deploying new contracts.
const TrialTransistionsAbstractBin = `0x`

// DeployTrialTransistionsAbstract deploys a new Ethereum contract, binding an instance of TrialTransistionsAbstract to it.
func DeployTrialTransistionsAbstract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TrialTransistionsAbstract, error) {
	parsed, err := abi.JSON(strings.NewReader(TrialTransistionsAbstractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TrialTransistionsAbstractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TrialTransistionsAbstract{TrialTransistionsAbstractCaller: TrialTransistionsAbstractCaller{contract: contract}, TrialTransistionsAbstractTransactor: TrialTransistionsAbstractTransactor{contract: contract}}, nil
}

// TrialTransistionsAbstract is an auto generated Go binding around an Ethereum contract.
type TrialTransistionsAbstract struct {
	TrialTransistionsAbstractCaller     // Read-only binding to the contract
	TrialTransistionsAbstractTransactor // Write-only binding to the contract
}

// TrialTransistionsAbstractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TrialTransistionsAbstractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrialTransistionsAbstractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TrialTransistionsAbstractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrialTransistionsAbstractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TrialTransistionsAbstractSession struct {
	Contract     *TrialTransistionsAbstract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// TrialTransistionsAbstractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TrialTransistionsAbstractCallerSession struct {
	Contract *TrialTransistionsAbstractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// TrialTransistionsAbstractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TrialTransistionsAbstractTransactorSession struct {
	Contract     *TrialTransistionsAbstractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// TrialTransistionsAbstractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TrialTransistionsAbstractRaw struct {
	Contract *TrialTransistionsAbstract // Generic contract binding to access the raw methods on
}

// TrialTransistionsAbstractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TrialTransistionsAbstractCallerRaw struct {
	Contract *TrialTransistionsAbstractCaller // Generic read-only contract binding to access the raw methods on
}

// TrialTransistionsAbstractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TrialTransistionsAbstractTransactorRaw struct {
	Contract *TrialTransistionsAbstractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrialTransistionsAbstract creates a new instance of TrialTransistionsAbstract, bound to a specific deployed contract.
func NewTrialTransistionsAbstract(address common.Address, backend bind.ContractBackend) (*TrialTransistionsAbstract, error) {
	contract, err := bindTrialTransistionsAbstract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TrialTransistionsAbstract{TrialTransistionsAbstractCaller: TrialTransistionsAbstractCaller{contract: contract}, TrialTransistionsAbstractTransactor: TrialTransistionsAbstractTransactor{contract: contract}}, nil
}

// NewTrialTransistionsAbstractCaller creates a new read-only instance of TrialTransistionsAbstract, bound to a specific deployed contract.
func NewTrialTransistionsAbstractCaller(address common.Address, caller bind.ContractCaller) (*TrialTransistionsAbstractCaller, error) {
	contract, err := bindTrialTransistionsAbstract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TrialTransistionsAbstractCaller{contract: contract}, nil
}

// NewTrialTransistionsAbstractTransactor creates a new write-only instance of TrialTransistionsAbstract, bound to a specific deployed contract.
func NewTrialTransistionsAbstractTransactor(address common.Address, transactor bind.ContractTransactor) (*TrialTransistionsAbstractTransactor, error) {
	contract, err := bindTrialTransistionsAbstract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TrialTransistionsAbstractTransactor{contract: contract}, nil
}

// bindTrialTransistionsAbstract binds a generic wrapper to an already deployed contract.
func bindTrialTransistionsAbstract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TrialTransistionsAbstractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TrialTransistionsAbstract *TrialTransistionsAbstractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TrialTransistionsAbstract.Contract.TrialTransistionsAbstractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TrialTransistionsAbstract *TrialTransistionsAbstractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrialTransistionsAbstract.Contract.TrialTransistionsAbstractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TrialTransistionsAbstract *TrialTransistionsAbstractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TrialTransistionsAbstract.Contract.TrialTransistionsAbstractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TrialTransistionsAbstract *TrialTransistionsAbstractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TrialTransistionsAbstract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TrialTransistionsAbstract *TrialTransistionsAbstractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrialTransistionsAbstract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TrialTransistionsAbstract *TrialTransistionsAbstractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TrialTransistionsAbstract.Contract.contract.Transact(opts, method, params...)
}

// GetInitialStatus is a paid mutator transaction binding the contract method 0x07e5ae5f.
//
// Solidity: function getInitialStatus() returns(status uint8)
func (_TrialTransistionsAbstract *TrialTransistionsAbstractTransactor) GetInitialStatus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrialTransistionsAbstract.contract.Transact(opts, "getInitialStatus")
}

// GetInitialStatus is a paid mutator transaction binding the contract method 0x07e5ae5f.
//
// Solidity: function getInitialStatus() returns(status uint8)
func (_TrialTransistionsAbstract *TrialTransistionsAbstractSession) GetInitialStatus() (*types.Transaction, error) {
	return _TrialTransistionsAbstract.Contract.GetInitialStatus(&_TrialTransistionsAbstract.TransactOpts)
}

// GetInitialStatus is a paid mutator transaction binding the contract method 0x07e5ae5f.
//
// Solidity: function getInitialStatus() returns(status uint8)
func (_TrialTransistionsAbstract *TrialTransistionsAbstractTransactorSession) GetInitialStatus() (*types.Transaction, error) {
	return _TrialTransistionsAbstract.Contract.GetInitialStatus(&_TrialTransistionsAbstract.TransactOpts)
}

// GetStatus is a paid mutator transaction binding the contract method 0xd39eb301.
//
// Solidity: function getStatus(witnessState uint8, trialStatus uint8) returns(status uint8)
func (_TrialTransistionsAbstract *TrialTransistionsAbstractTransactor) GetStatus(opts *bind.TransactOpts, witnessState uint8, trialStatus uint8) (*types.Transaction, error) {
	return _TrialTransistionsAbstract.contract.Transact(opts, "getStatus", witnessState, trialStatus)
}

// GetStatus is a paid mutator transaction binding the contract method 0xd39eb301.
//
// Solidity: function getStatus(witnessState uint8, trialStatus uint8) returns(status uint8)
func (_TrialTransistionsAbstract *TrialTransistionsAbstractSession) GetStatus(witnessState uint8, trialStatus uint8) (*types.Transaction, error) {
	return _TrialTransistionsAbstract.Contract.GetStatus(&_TrialTransistionsAbstract.TransactOpts, witnessState, trialStatus)
}

// GetStatus is a paid mutator transaction binding the contract method 0xd39eb301.
//
// Solidity: function getStatus(witnessState uint8, trialStatus uint8) returns(status uint8)
func (_TrialTransistionsAbstract *TrialTransistionsAbstractTransactorSession) GetStatus(witnessState uint8, trialStatus uint8) (*types.Transaction, error) {
	return _TrialTransistionsAbstract.Contract.GetStatus(&_TrialTransistionsAbstract.TransactOpts, witnessState, trialStatus)
}

// GetTrialExpiry is a paid mutator transaction binding the contract method 0xada01610.
//
// Solidity: function getTrialExpiry() returns(expiery uint256)
func (_TrialTransistionsAbstract *TrialTransistionsAbstractTransactor) GetTrialExpiry(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrialTransistionsAbstract.contract.Transact(opts, "getTrialExpiry")
}

// GetTrialExpiry is a paid mutator transaction binding the contract method 0xada01610.
//
// Solidity: function getTrialExpiry() returns(expiery uint256)
func (_TrialTransistionsAbstract *TrialTransistionsAbstractSession) GetTrialExpiry() (*types.Transaction, error) {
	return _TrialTransistionsAbstract.Contract.GetTrialExpiry(&_TrialTransistionsAbstract.TransactOpts)
}

// GetTrialExpiry is a paid mutator transaction binding the contract method 0xada01610.
//
// Solidity: function getTrialExpiry() returns(expiery uint256)
func (_TrialTransistionsAbstract *TrialTransistionsAbstractTransactorSession) GetTrialExpiry() (*types.Transaction, error) {
	return _TrialTransistionsAbstract.Contract.GetTrialExpiry(&_TrialTransistionsAbstract.TransactOpts)
}

// GetWitness is a paid mutator transaction binding the contract method 0xdb061563.
//
// Solidity: function getWitness(trialStatus uint8) returns(address)
func (_TrialTransistionsAbstract *TrialTransistionsAbstractTransactor) GetWitness(opts *bind.TransactOpts, trialStatus uint8) (*types.Transaction, error) {
	return _TrialTransistionsAbstract.contract.Transact(opts, "getWitness", trialStatus)
}

// GetWitness is a paid mutator transaction binding the contract method 0xdb061563.
//
// Solidity: function getWitness(trialStatus uint8) returns(address)
func (_TrialTransistionsAbstract *TrialTransistionsAbstractSession) GetWitness(trialStatus uint8) (*types.Transaction, error) {
	return _TrialTransistionsAbstract.Contract.GetWitness(&_TrialTransistionsAbstract.TransactOpts, trialStatus)
}

// GetWitness is a paid mutator transaction binding the contract method 0xdb061563.
//
// Solidity: function getWitness(trialStatus uint8) returns(address)
func (_TrialTransistionsAbstract *TrialTransistionsAbstractTransactorSession) GetWitness(trialStatus uint8) (*types.Transaction, error) {
	return _TrialTransistionsAbstract.Contract.GetWitness(&_TrialTransistionsAbstract.TransactOpts, trialStatus)
}

// WitnessAbstractABI is the input ABI used to generate the binding from.
const WitnessAbstractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"serviceId\",\"type\":\"bytes32\"},{\"name\":\"clientAddress\",\"type\":\"address\"}],\"name\":\"testimonyFor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"}]"

// WitnessAbstractBin is the compiled bytecode used for deploying new contracts.
const WitnessAbstractBin = `0x`

// DeployWitnessAbstract deploys a new Ethereum contract, binding an instance of WitnessAbstract to it.
func DeployWitnessAbstract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WitnessAbstract, error) {
	parsed, err := abi.JSON(strings.NewReader(WitnessAbstractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WitnessAbstractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WitnessAbstract{WitnessAbstractCaller: WitnessAbstractCaller{contract: contract}, WitnessAbstractTransactor: WitnessAbstractTransactor{contract: contract}}, nil
}

// WitnessAbstract is an auto generated Go binding around an Ethereum contract.
type WitnessAbstract struct {
	WitnessAbstractCaller     // Read-only binding to the contract
	WitnessAbstractTransactor // Write-only binding to the contract
}

// WitnessAbstractCaller is an auto generated read-only Go binding around an Ethereum contract.
type WitnessAbstractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WitnessAbstractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WitnessAbstractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WitnessAbstractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WitnessAbstractSession struct {
	Contract     *WitnessAbstract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WitnessAbstractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WitnessAbstractCallerSession struct {
	Contract *WitnessAbstractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// WitnessAbstractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WitnessAbstractTransactorSession struct {
	Contract     *WitnessAbstractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// WitnessAbstractRaw is an auto generated low-level Go binding around an Ethereum contract.
type WitnessAbstractRaw struct {
	Contract *WitnessAbstract // Generic contract binding to access the raw methods on
}

// WitnessAbstractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WitnessAbstractCallerRaw struct {
	Contract *WitnessAbstractCaller // Generic read-only contract binding to access the raw methods on
}

// WitnessAbstractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WitnessAbstractTransactorRaw struct {
	Contract *WitnessAbstractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWitnessAbstract creates a new instance of WitnessAbstract, bound to a specific deployed contract.
func NewWitnessAbstract(address common.Address, backend bind.ContractBackend) (*WitnessAbstract, error) {
	contract, err := bindWitnessAbstract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WitnessAbstract{WitnessAbstractCaller: WitnessAbstractCaller{contract: contract}, WitnessAbstractTransactor: WitnessAbstractTransactor{contract: contract}}, nil
}

// NewWitnessAbstractCaller creates a new read-only instance of WitnessAbstract, bound to a specific deployed contract.
func NewWitnessAbstractCaller(address common.Address, caller bind.ContractCaller) (*WitnessAbstractCaller, error) {
	contract, err := bindWitnessAbstract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &WitnessAbstractCaller{contract: contract}, nil
}

// NewWitnessAbstractTransactor creates a new write-only instance of WitnessAbstract, bound to a specific deployed contract.
func NewWitnessAbstractTransactor(address common.Address, transactor bind.ContractTransactor) (*WitnessAbstractTransactor, error) {
	contract, err := bindWitnessAbstract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &WitnessAbstractTransactor{contract: contract}, nil
}

// bindWitnessAbstract binds a generic wrapper to an already deployed contract.
func bindWitnessAbstract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WitnessAbstractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WitnessAbstract *WitnessAbstractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WitnessAbstract.Contract.WitnessAbstractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WitnessAbstract *WitnessAbstractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WitnessAbstract.Contract.WitnessAbstractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WitnessAbstract *WitnessAbstractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WitnessAbstract.Contract.WitnessAbstractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WitnessAbstract *WitnessAbstractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WitnessAbstract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WitnessAbstract *WitnessAbstractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WitnessAbstract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WitnessAbstract *WitnessAbstractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WitnessAbstract.Contract.contract.Transact(opts, method, params...)
}

// TestimonyFor is a paid mutator transaction binding the contract method 0xf243aee5.
//
// Solidity: function testimonyFor(serviceId bytes32, clientAddress address) returns(uint8)
func (_WitnessAbstract *WitnessAbstractTransactor) TestimonyFor(opts *bind.TransactOpts, serviceId [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _WitnessAbstract.contract.Transact(opts, "testimonyFor", serviceId, clientAddress)
}

// TestimonyFor is a paid mutator transaction binding the contract method 0xf243aee5.
//
// Solidity: function testimonyFor(serviceId bytes32, clientAddress address) returns(uint8)
func (_WitnessAbstract *WitnessAbstractSession) TestimonyFor(serviceId [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _WitnessAbstract.Contract.TestimonyFor(&_WitnessAbstract.TransactOpts, serviceId, clientAddress)
}

// TestimonyFor is a paid mutator transaction binding the contract method 0xf243aee5.
//
// Solidity: function testimonyFor(serviceId bytes32, clientAddress address) returns(uint8)
func (_WitnessAbstract *WitnessAbstractTransactorSession) TestimonyFor(serviceId [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _WitnessAbstract.Contract.TestimonyFor(&_WitnessAbstract.TransactOpts, serviceId, clientAddress)
}
