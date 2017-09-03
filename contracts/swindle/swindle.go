// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package swindle

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// OwnedABI is the input ABI used to generate the binding from.
const OwnedABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"}]"

// OwnedBin is the compiled bytecode used for deploying new contracts.
const OwnedBin = `0x6060604052341561000f57600080fd5b5b60008054600160a060020a03191633600160a060020a03161790555b5b61015c8061003c6000396000f300606060405263ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b8114610048578063a6f9dae114610084575b600080fd5b341561005357600080fd5b61005b6100b2565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561008f57600080fd5b6100b073ffffffffffffffffffffffffffffffffffffffff600435166100ce565b005b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6000543373ffffffffffffffffffffffffffffffffffffffff9081169116146100f657600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff83161790555b5b505600a165627a7a723058203c8bfd43707cccaf486bb55f9a5581ad9b902d48809ddab1fdb4242744e7464c0029`

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
const SampleTokenBin = `0x6060604052341561000f57600080fd5b60405160208061069d833981016040528080519150505b5b60038054600160a060020a03191633600160a060020a03161790555b600160a060020a03331660009081526001602052604081208290558190555b505b61062a806100736000396000f300606060405236156100965763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461009b57806318160ddd146100d157806323b872dd146100f657806370a08231146101325780638da5cb5b14610163578063a3aff92314610192578063a6f9dae1146101b9578063a9059cbb146101da578063dd62ed3e14610210575b600080fd5b34156100a657600080fd5b6100bd600160a060020a0360043516602435610247565b604051901515815260200160405180910390f35b34156100dc57600080fd5b6100e46102b4565b60405190815260200160405180910390f35b341561010157600080fd5b6100bd600160a060020a03600435811690602435166044356102ba565b604051901515815260200160405180910390f35b341561013d57600080fd5b6100e4600160a060020a03600435166102fd565b60405190815260200160405180910390f35b341561016e57600080fd5b61017661031c565b604051600160a060020a03909116815260200160405180910390f35b341561019d57600080fd5b6101b7600160a060020a036004351660243560443561032b565b005b34156101c457600080fd5b6101b7600160a060020a03600435166103e8565b005b34156101e557600080fd5b6100bd600160a060020a0360043516602435610430565b604051901515815260200160405180910390f35b341561021b57600080fd5b6100e4600160a060020a03600435811690602435166104d7565b60405190815260200160405180910390f35b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a35060015b92915050565b60005481565b600160a060020a0380841660009081526002602090815260408083203390941683529290529081208054830190556102f3848484610504565b90505b9392505050565b600160a060020a0381166000908152600160205260409020545b919050565b600354600160a060020a031681565b60035433600160a060020a0390811691161461034657600080fd5b600160a060020a038316600081815260016020526040808220805486019055815485019091558291907fdfb81fb379557413b0a951b4d7bf7a9df393801d8c539d5e201d6a8daeb913b99085905190815260200160405180910390a382600160a060020a031660007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a35b5b505050565b60035433600160a060020a0390811691161461040357600080fd5b6003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b50565b600160a060020a0333166000908152600160205260408120548290108015906104595750600082115b1561009657600160a060020a033381166000818152600160205260408082208054879003905592861680825290839020805486019055917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060016102ae565b600080fd5b5b92915050565b600160a060020a038083166000908152600260209081526040808320938516835292905220545b92915050565b600160a060020a0383166000908152600160205260408120548290108015906105545750600160a060020a0380851660009081526002602090815260408083203390941683529290522054829010155b80156105605750600082115b1561009657600160a060020a03808416600081815260016020908152604080832080548801905588851680845281842080548990039055600283528184203390961684529490915290819020805486900390559091907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060016102f6565b600080fd5b5b93925050505600a165627a7a72305820055d55bf16b89ad6bb8e496f491bdb9fd1e3a3c7dd49f5163d7d8785f2d7045c0029`

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
const StandardTokenBin = `0x6060604052341561000f57600080fd5b5b61043b8061001f6000396000f300606060405236156100755763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461007a57806318160ddd146100b057806323b872dd146100d557806370a0823114610111578063a9059cbb14610142578063dd62ed3e14610178575b600080fd5b341561008557600080fd5b61009c600160a060020a03600435166024356101af565b604051901515815260200160405180910390f35b34156100bb57600080fd5b6100c361021c565b60405190815260200160405180910390f35b34156100e057600080fd5b61009c600160a060020a0360043581169060243516604435610222565b604051901515815260200160405180910390f35b341561011c57600080fd5b6100c3600160a060020a036004351661031c565b60405190815260200160405180910390f35b341561014d57600080fd5b61009c600160a060020a036004351660243561033b565b604051901515815260200160405180910390f35b341561018357600080fd5b6100c3600160a060020a03600435811690602435166103e2565b60405190815260200160405180910390f35b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a35060015b92915050565b60005481565b600160a060020a0383166000908152600160205260408120548290108015906102725750600160a060020a0380851660009081526002602090815260408083203390941683529290522054829010155b801561027e5750600082115b1561007557600160a060020a03808416600081815260016020908152604080832080548801905588851680845281842080548990039055600283528184203390961684529490915290819020805486900390559091907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a3506001610314565b600080fd5b5b9392505050565b600160a060020a0381166000908152600160205260409020545b919050565b600160a060020a0333166000908152600160205260408120548290108015906103645750600082115b1561007557600160a060020a033381166000818152600160205260408082208054879003905592861680825290839020805486019055917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a3506001610216565b600080fd5b5b92915050565b600160a060020a038083166000908152600260209081526040808320938516835292905220545b929150505600a165627a7a72305820345a8443fac98d6f1d77c09fbd611c6f92c6e07163d332fdb7ce9803a17176f70029`

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

// SwearAbstractABI is the input ABI used to generate the binding from.
const SwearAbstractABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"collectDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_player\",\"type\":\"address\"}],\"name\":\"unRegister\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"incrementOpenCases\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_player\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_swindle\",\"type\":\"address\"}],\"name\":\"setSwindleContractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"},{\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"compensate\",\"outputs\":[{\"name\":\"compensated\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"decrementOpenCases\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"epochs\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"player\",\"type\":\"address\"}],\"name\":\"isRegistered\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"}]"

// SwearAbstractBin is the compiled bytecode used for deploying new contracts.
const SwearAbstractBin = `0x`

// DeploySwearAbstract deploys a new Ethereum contract, binding an instance of SwearAbstract to it.
func DeploySwearAbstract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SwearAbstract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwearAbstractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SwearAbstractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SwearAbstract{SwearAbstractCaller: SwearAbstractCaller{contract: contract}, SwearAbstractTransactor: SwearAbstractTransactor{contract: contract}}, nil
}

// SwearAbstract is an auto generated Go binding around an Ethereum contract.
type SwearAbstract struct {
	SwearAbstractCaller     // Read-only binding to the contract
	SwearAbstractTransactor // Write-only binding to the contract
}

// SwearAbstractCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwearAbstractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwearAbstractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwearAbstractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwearAbstractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwearAbstractSession struct {
	Contract     *SwearAbstract    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwearAbstractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwearAbstractCallerSession struct {
	Contract *SwearAbstractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SwearAbstractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwearAbstractTransactorSession struct {
	Contract     *SwearAbstractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SwearAbstractRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwearAbstractRaw struct {
	Contract *SwearAbstract // Generic contract binding to access the raw methods on
}

// SwearAbstractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwearAbstractCallerRaw struct {
	Contract *SwearAbstractCaller // Generic read-only contract binding to access the raw methods on
}

// SwearAbstractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwearAbstractTransactorRaw struct {
	Contract *SwearAbstractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwearAbstract creates a new instance of SwearAbstract, bound to a specific deployed contract.
func NewSwearAbstract(address common.Address, backend bind.ContractBackend) (*SwearAbstract, error) {
	contract, err := bindSwearAbstract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwearAbstract{SwearAbstractCaller: SwearAbstractCaller{contract: contract}, SwearAbstractTransactor: SwearAbstractTransactor{contract: contract}}, nil
}

// NewSwearAbstractCaller creates a new read-only instance of SwearAbstract, bound to a specific deployed contract.
func NewSwearAbstractCaller(address common.Address, caller bind.ContractCaller) (*SwearAbstractCaller, error) {
	contract, err := bindSwearAbstract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SwearAbstractCaller{contract: contract}, nil
}

// NewSwearAbstractTransactor creates a new write-only instance of SwearAbstract, bound to a specific deployed contract.
func NewSwearAbstractTransactor(address common.Address, transactor bind.ContractTransactor) (*SwearAbstractTransactor, error) {
	contract, err := bindSwearAbstract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SwearAbstractTransactor{contract: contract}, nil
}

// bindSwearAbstract binds a generic wrapper to an already deployed contract.
func bindSwearAbstract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwearAbstractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwearAbstract *SwearAbstractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SwearAbstract.Contract.SwearAbstractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwearAbstract *SwearAbstractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwearAbstract.Contract.SwearAbstractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwearAbstract *SwearAbstractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwearAbstract.Contract.SwearAbstractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwearAbstract *SwearAbstractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SwearAbstract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwearAbstract *SwearAbstractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwearAbstract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwearAbstract *SwearAbstractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwearAbstract.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SwearAbstract *SwearAbstractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SwearAbstract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SwearAbstract *SwearAbstractSession) Owner() (common.Address, error) {
	return _SwearAbstract.Contract.Owner(&_SwearAbstract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SwearAbstract *SwearAbstractCallerSession) Owner() (common.Address, error) {
	return _SwearAbstract.Contract.Owner(&_SwearAbstract.CallOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_SwearAbstract *SwearAbstractTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _SwearAbstract.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_SwearAbstract *SwearAbstractSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _SwearAbstract.Contract.ChangeOwner(&_SwearAbstract.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_SwearAbstract *SwearAbstractTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _SwearAbstract.Contract.ChangeOwner(&_SwearAbstract.TransactOpts, _newOwner)
}

// CollectDeposit is a paid mutator transaction binding the contract method 0x02336f92.
//
// Solidity: function collectDeposit() returns(bool)
func (_SwearAbstract *SwearAbstractTransactor) CollectDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwearAbstract.contract.Transact(opts, "collectDeposit")
}

// CollectDeposit is a paid mutator transaction binding the contract method 0x02336f92.
//
// Solidity: function collectDeposit() returns(bool)
func (_SwearAbstract *SwearAbstractSession) CollectDeposit() (*types.Transaction, error) {
	return _SwearAbstract.Contract.CollectDeposit(&_SwearAbstract.TransactOpts)
}

// CollectDeposit is a paid mutator transaction binding the contract method 0x02336f92.
//
// Solidity: function collectDeposit() returns(bool)
func (_SwearAbstract *SwearAbstractTransactorSession) CollectDeposit() (*types.Transaction, error) {
	return _SwearAbstract.Contract.CollectDeposit(&_SwearAbstract.TransactOpts)
}

// Compensate is a paid mutator transaction binding the contract method 0x8a1bd7c0.
//
// Solidity: function compensate(_beneficiary address, reward uint256) returns(compensated bool)
func (_SwearAbstract *SwearAbstractTransactor) Compensate(opts *bind.TransactOpts, _beneficiary common.Address, reward *big.Int) (*types.Transaction, error) {
	return _SwearAbstract.contract.Transact(opts, "compensate", _beneficiary, reward)
}

// Compensate is a paid mutator transaction binding the contract method 0x8a1bd7c0.
//
// Solidity: function compensate(_beneficiary address, reward uint256) returns(compensated bool)
func (_SwearAbstract *SwearAbstractSession) Compensate(_beneficiary common.Address, reward *big.Int) (*types.Transaction, error) {
	return _SwearAbstract.Contract.Compensate(&_SwearAbstract.TransactOpts, _beneficiary, reward)
}

// Compensate is a paid mutator transaction binding the contract method 0x8a1bd7c0.
//
// Solidity: function compensate(_beneficiary address, reward uint256) returns(compensated bool)
func (_SwearAbstract *SwearAbstractTransactorSession) Compensate(_beneficiary common.Address, reward *big.Int) (*types.Transaction, error) {
	return _SwearAbstract.Contract.Compensate(&_SwearAbstract.TransactOpts, _beneficiary, reward)
}

// DecrementOpenCases is a paid mutator transaction binding the contract method 0x8f6a8b03.
//
// Solidity: function decrementOpenCases(_address address) returns()
func (_SwearAbstract *SwearAbstractTransactor) DecrementOpenCases(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _SwearAbstract.contract.Transact(opts, "decrementOpenCases", _address)
}

// DecrementOpenCases is a paid mutator transaction binding the contract method 0x8f6a8b03.
//
// Solidity: function decrementOpenCases(_address address) returns()
func (_SwearAbstract *SwearAbstractSession) DecrementOpenCases(_address common.Address) (*types.Transaction, error) {
	return _SwearAbstract.Contract.DecrementOpenCases(&_SwearAbstract.TransactOpts, _address)
}

// DecrementOpenCases is a paid mutator transaction binding the contract method 0x8f6a8b03.
//
// Solidity: function decrementOpenCases(_address address) returns()
func (_SwearAbstract *SwearAbstractTransactorSession) DecrementOpenCases(_address common.Address) (*types.Transaction, error) {
	return _SwearAbstract.Contract.DecrementOpenCases(&_SwearAbstract.TransactOpts, _address)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(epochs uint256) returns(bool)
func (_SwearAbstract *SwearAbstractTransactor) Deposit(opts *bind.TransactOpts, epochs *big.Int) (*types.Transaction, error) {
	return _SwearAbstract.contract.Transact(opts, "deposit", epochs)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(epochs uint256) returns(bool)
func (_SwearAbstract *SwearAbstractSession) Deposit(epochs *big.Int) (*types.Transaction, error) {
	return _SwearAbstract.Contract.Deposit(&_SwearAbstract.TransactOpts, epochs)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(epochs uint256) returns(bool)
func (_SwearAbstract *SwearAbstractTransactorSession) Deposit(epochs *big.Int) (*types.Transaction, error) {
	return _SwearAbstract.Contract.Deposit(&_SwearAbstract.TransactOpts, epochs)
}

// IncrementOpenCases is a paid mutator transaction binding the contract method 0x27f730ee.
//
// Solidity: function incrementOpenCases(_address address) returns()
func (_SwearAbstract *SwearAbstractTransactor) IncrementOpenCases(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _SwearAbstract.contract.Transact(opts, "incrementOpenCases", _address)
}

// IncrementOpenCases is a paid mutator transaction binding the contract method 0x27f730ee.
//
// Solidity: function incrementOpenCases(_address address) returns()
func (_SwearAbstract *SwearAbstractSession) IncrementOpenCases(_address common.Address) (*types.Transaction, error) {
	return _SwearAbstract.Contract.IncrementOpenCases(&_SwearAbstract.TransactOpts, _address)
}

// IncrementOpenCases is a paid mutator transaction binding the contract method 0x27f730ee.
//
// Solidity: function incrementOpenCases(_address address) returns()
func (_SwearAbstract *SwearAbstractTransactorSession) IncrementOpenCases(_address common.Address) (*types.Transaction, error) {
	return _SwearAbstract.Contract.IncrementOpenCases(&_SwearAbstract.TransactOpts, _address)
}

// IsRegistered is a paid mutator transaction binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(player address) returns(bool)
func (_SwearAbstract *SwearAbstractTransactor) IsRegistered(opts *bind.TransactOpts, player common.Address) (*types.Transaction, error) {
	return _SwearAbstract.contract.Transact(opts, "isRegistered", player)
}

// IsRegistered is a paid mutator transaction binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(player address) returns(bool)
func (_SwearAbstract *SwearAbstractSession) IsRegistered(player common.Address) (*types.Transaction, error) {
	return _SwearAbstract.Contract.IsRegistered(&_SwearAbstract.TransactOpts, player)
}

// IsRegistered is a paid mutator transaction binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(player address) returns(bool)
func (_SwearAbstract *SwearAbstractTransactorSession) IsRegistered(player common.Address) (*types.Transaction, error) {
	return _SwearAbstract.Contract.IsRegistered(&_SwearAbstract.TransactOpts, player)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(_player address) returns(bool)
func (_SwearAbstract *SwearAbstractTransactor) Register(opts *bind.TransactOpts, _player common.Address) (*types.Transaction, error) {
	return _SwearAbstract.contract.Transact(opts, "register", _player)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(_player address) returns(bool)
func (_SwearAbstract *SwearAbstractSession) Register(_player common.Address) (*types.Transaction, error) {
	return _SwearAbstract.Contract.Register(&_SwearAbstract.TransactOpts, _player)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(_player address) returns(bool)
func (_SwearAbstract *SwearAbstractTransactorSession) Register(_player common.Address) (*types.Transaction, error) {
	return _SwearAbstract.Contract.Register(&_SwearAbstract.TransactOpts, _player)
}

// SetSwindleContractAddress is a paid mutator transaction binding the contract method 0x59f50e73.
//
// Solidity: function setSwindleContractAddress(_swindle address) returns(bool)
func (_SwearAbstract *SwearAbstractTransactor) SetSwindleContractAddress(opts *bind.TransactOpts, _swindle common.Address) (*types.Transaction, error) {
	return _SwearAbstract.contract.Transact(opts, "setSwindleContractAddress", _swindle)
}

// SetSwindleContractAddress is a paid mutator transaction binding the contract method 0x59f50e73.
//
// Solidity: function setSwindleContractAddress(_swindle address) returns(bool)
func (_SwearAbstract *SwearAbstractSession) SetSwindleContractAddress(_swindle common.Address) (*types.Transaction, error) {
	return _SwearAbstract.Contract.SetSwindleContractAddress(&_SwearAbstract.TransactOpts, _swindle)
}

// SetSwindleContractAddress is a paid mutator transaction binding the contract method 0x59f50e73.
//
// Solidity: function setSwindleContractAddress(_swindle address) returns(bool)
func (_SwearAbstract *SwearAbstractTransactorSession) SetSwindleContractAddress(_swindle common.Address) (*types.Transaction, error) {
	return _SwearAbstract.Contract.SetSwindleContractAddress(&_SwearAbstract.TransactOpts, _swindle)
}

// UnRegister is a paid mutator transaction binding the contract method 0x27c95625.
//
// Solidity: function unRegister(_player address) returns()
func (_SwearAbstract *SwearAbstractTransactor) UnRegister(opts *bind.TransactOpts, _player common.Address) (*types.Transaction, error) {
	return _SwearAbstract.contract.Transact(opts, "unRegister", _player)
}

// UnRegister is a paid mutator transaction binding the contract method 0x27c95625.
//
// Solidity: function unRegister(_player address) returns()
func (_SwearAbstract *SwearAbstractSession) UnRegister(_player common.Address) (*types.Transaction, error) {
	return _SwearAbstract.Contract.UnRegister(&_SwearAbstract.TransactOpts, _player)
}

// UnRegister is a paid mutator transaction binding the contract method 0x27c95625.
//
// Solidity: function unRegister(_player address) returns()
func (_SwearAbstract *SwearAbstractTransactorSession) UnRegister(_player common.Address) (*types.Transaction, error) {
	return _SwearAbstract.Contract.UnRegister(&_SwearAbstract.TransactOpts, _player)
}

// SwindleABI is the input ABI used to generate the binding from.
const SwindleABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"trialRules\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ids\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"serviceId\",\"type\":\"bytes32\"}],\"name\":\"newCase\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"trial\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"swear\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_swear\",\"type\":\"address\"},{\"name\":\"_trialRules\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"decide\",\"type\":\"string\"}],\"name\":\"Decision\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"plaintiff\",\"type\":\"address\"}],\"name\":\"NewCaseOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"plaintiff\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"CaseResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"plaintiff\",\"type\":\"address\"}],\"name\":\"NewEvidenceSubmitted\",\"type\":\"event\"}]"

// SwindleBin is the compiled bytecode used for deploying new contracts.
const SwindleBin = `0x6060604052341561000f57600080fd5b60405160408061100683398101604052808051919060200180519150505b5b60008054600160a060020a03191633600160a060020a03161790555b60028054600160a060020a031916600160a060020a038481169190911791829055166359f50e73306000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b15156100d257600080fd5b6102c65a03f115156100e357600080fd5b5050506040518051905015156100f857600080fd5b60018054600160a060020a031916600160a060020a0383161790555b50505b610ee0806101266000396000f300606060405236156100725763ffffffff60e060020a6000350416635de28ae081146100775780638a29d969146100a35780638da5cb5b146100d25780639f20bc9214610101578063a6f9dae114610135578063cc345e0a14610156578063cf343bb814610180578063f20f1f58146101aa575b600080fd5b341561008257600080fd5b61008d6004356101d9565b60405160ff909116815260200160405180910390f35b34156100ae57600080fd5b6100b66101f4565b604051600160a060020a03909116815260200160405180910390f35b34156100dd57600080fd5b6100b6610203565b604051600160a060020a03909116815260200160405180910390f35b341561010c57600080fd5b610123600160a060020a0360043516602435610212565b60405190815260200160405180910390f35b341561014057600080fd5b610154600160a060020a0360043516610244565b005b341561016157600080fd5b61016c60043561028c565b604051901515815260200160405180910390f35b341561018b57600080fd5b61016c6004356104a6565b604051901515815260200160405180910390f35b34156101b557600080fd5b6100b661054d565b604051600160a060020a03909116815260200160405180910390f35b60008181526003602052604090206002015460ff165b919050565b600154600160a060020a031681565b600054600160a060020a031681565b60046020528160005260406000208181548110151561022d57fe5b906000526020600020900160005b91509150505481565b60005433600160a060020a0390811691161461025f57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b50565b6002546000908190600160a060020a031663c3c5a54733836040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b15156102e957600080fd5b6102c65a03f115156102fa57600080fd5b50505060405180519050151561030f57600080fd5b60015461037f9033908590600160a060020a03166307e5ae5f6000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561035f57600080fd5b6102c65a03f1151561037057600080fd5b5050506040518051905061055c565b905080151561039157600091506104a0565b600254600160a060020a03166327f730ee3360405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401600060405180830381600087803b15156103e157600080fd5b6102c65a03f115156103f257600080fd5b5050600254600054600160a060020a0391821692506327f730ee911660405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401600060405180830381600087803b151561044c57600080fd5b6102c65a03f1151561045d57600080fd5b505050600160a060020a03331660009081526004602052604090208054600181016104888382610e69565b916000526020600020900160005b5082905550600191505b50919050565b600254600090600160a060020a031663c3c5a54733836040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561050157600080fd5b6102c65a03f1151561051257600080fd5b50505060405180519050151561052757600080fd5b6105308261067a565b151561053b57600080fd5b6105448261069c565b5060015b919050565b600254600160a060020a031681565b6000838342604051600160a060020a03939093166c0100000000000000000000000002835260148301919091526034820152605401604051908190039020600081815260036020526040902060020154909150610100900460ff16156105c457506000610673565b60806040519081016040908152600160a060020a0386168252602080830186905260ff8516828401526001606084015260008481526003909152208151815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039190911617815560208201516001820155604082015160028201805460ff191660ff9290921691909117905560608201516002909101805460ff929092166101000261ff0019909216919091179055505b9392505050565b600081815260036020526040902060020154610100900460ff1615155b919050565b6000806000806000806106ae876101d9565b95506106b987610b44565b945094505b60005b60ff878116911614610b3a57600154600160a060020a031663db0615638760006040516020015260405160e060020a63ffffffff841602815260ff9091166004820152602401602060405180830381600087803b151561072057600080fd5b6102c65a03f1151561073157600080fd5b5050506040518051935050600160a060020a038316151561075b57610754610b6c565b9150610b3a565b600154600160a060020a031663fcc0a2ae888860006040516020015260405160e060020a63ffffffff8516028152600481019290925260ff166024820152604401602060405180830381600087803b15156107b557600080fd5b6102c65a03f115156107c657600080fd5b5050506040518051915050600160a060020a038316637aba9a6788868860006040516020015260405160e060020a63ffffffff861602815260048101939093526024830191909152600160a060020a03166044820152606401602060405180830381600087803b151561083857600080fd5b6102c65a03f1151561084957600080fd5b50505060405180519050801561085d575080155b156108eb5782600160a060020a031663c4b4528388868860006040516020015260405160e060020a63ffffffff861602815260048101939093526024830191909152600160a060020a03166044820152606401602060405180830381600087803b15156108c957600080fd5b6102c65a03f115156108da57600080fd5b50505060405180519050915061097c565b600154600160a060020a031663a7edcbe4888860006040516020015260405160e060020a63ffffffff8516028152600481019290925260ff166024820152604401602060405180830381600087803b151561094557600080fd5b6102c65a03f1151561095657600080fd5b5050506040518051905080610969575080155b15610977576002915061097c565b600191505b5b5b60025b82600281111561098d57fe5b141561099857610b3a565b600154600160a060020a031663d39eb3018360028111156109b557fe5b8860006040516020015260405160e060020a63ffffffff851602815260ff928316600482015291166024820152604401602060405180830381600087803b15156109fe57600080fd5b6102c65a03f11515610a0f57600080fd5b505050604051805190509550610a258787610b72565b5060015b60ff168660ff161480610a43575060025b60ff168660ff16145b15610b3557610a53878787610bb7565b5060005b600254909650600160a060020a0316638f6a8b038660405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401600060405180830381600087803b1515610aaa57600080fd5b6102c65a03f11515610abb57600080fd5b5050600254600054600160a060020a039182169250638f6a8b03911660405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401600060405180830381600087803b1515610b1557600080fd5b6102c65a03f11515610b2657600080fd5b505050610b338787610b72565b505b6106be565b5b50505050505050565b60008181526003602052604090208054600190910154600160a060020a03909116905b915091565b60025b90565b6000805433600160a060020a0390811691161415610b8f57600080fd5b506000828152600360205260409020600201805460ff191660ff831617905560015b92915050565b6000808060025b60ff168560ff161415610c2f577f3f379bb1eae8d950e4e7920751b198b0471a8c95bee23d7f7336e4d5ec70f3398685600088604051938452600160a060020a03909216602084015260408084019190915260ff90911660608301526080909101905180910390a160009250610ddc565b600154600160a060020a0316633d18b9126000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610c7757600080fd5b6102c65a03f11515610c8857600080fd5b5050506040518051600254909350600160a060020a03169050638a1bd7c0858460006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b1515610cf457600080fd5b6102c65a03f11515610d0557600080fd5b505050604051805190509050610d1a86610de5565b600254600160a060020a03166327c956258560405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401600060405180830381600087803b1515610d6a57600080fd5b6102c65a03f11515610d7b57600080fd5b5050507f3f379bb1eae8d950e4e7920751b198b0471a8c95bee23d7f7336e4d5ec70f33986858488604051938452600160a060020a03909216602084015260408084019190915260ff90911660608301526080909101905180910390a18092505b50509392505050565b60005b60008281526003602052604090206002015460ff90811691161415610e0c57600080fd5b6000818152600360205260408120805473ffffffffffffffffffffffffffffffffffffffff19168155600201805461ff00191690555b6000828152600360205260409020600201805460ff191660ff929092169190911790555b50565b815481835581811511610e8d57600083815260209020610e8d918101908301610e93565b5b505050565b610b6f91905b80821115610ead5760008155600101610e99565b5090565b905600a165627a7a72305820796ff660e7d711a5666d9b4607946619b555f99988dfcab4160c43df5e95eb720029`

// DeploySwindle deploys a new Ethereum contract, binding an instance of Swindle to it.
func DeploySwindle(auth *bind.TransactOpts, backend bind.ContractBackend, _swear common.Address, _trialRules common.Address) (common.Address, *types.Transaction, *Swindle, error) {
	parsed, err := abi.JSON(strings.NewReader(SwindleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SwindleBin), backend, _swear, _trialRules)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Swindle{SwindleCaller: SwindleCaller{contract: contract}, SwindleTransactor: SwindleTransactor{contract: contract}}, nil
}

// Swindle is an auto generated Go binding around an Ethereum contract.
type Swindle struct {
	SwindleCaller     // Read-only binding to the contract
	SwindleTransactor // Write-only binding to the contract
}

// SwindleCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwindleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwindleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwindleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwindleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwindleSession struct {
	Contract     *Swindle          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwindleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwindleCallerSession struct {
	Contract *SwindleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SwindleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwindleTransactorSession struct {
	Contract     *SwindleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SwindleRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwindleRaw struct {
	Contract *Swindle // Generic contract binding to access the raw methods on
}

// SwindleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwindleCallerRaw struct {
	Contract *SwindleCaller // Generic read-only contract binding to access the raw methods on
}

// SwindleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwindleTransactorRaw struct {
	Contract *SwindleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwindle creates a new instance of Swindle, bound to a specific deployed contract.
func NewSwindle(address common.Address, backend bind.ContractBackend) (*Swindle, error) {
	contract, err := bindSwindle(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Swindle{SwindleCaller: SwindleCaller{contract: contract}, SwindleTransactor: SwindleTransactor{contract: contract}}, nil
}

// NewSwindleCaller creates a new read-only instance of Swindle, bound to a specific deployed contract.
func NewSwindleCaller(address common.Address, caller bind.ContractCaller) (*SwindleCaller, error) {
	contract, err := bindSwindle(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SwindleCaller{contract: contract}, nil
}

// NewSwindleTransactor creates a new write-only instance of Swindle, bound to a specific deployed contract.
func NewSwindleTransactor(address common.Address, transactor bind.ContractTransactor) (*SwindleTransactor, error) {
	contract, err := bindSwindle(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SwindleTransactor{contract: contract}, nil
}

// bindSwindle binds a generic wrapper to an already deployed contract.
func bindSwindle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwindleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swindle *SwindleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Swindle.Contract.SwindleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swindle *SwindleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swindle.Contract.SwindleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swindle *SwindleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swindle.Contract.SwindleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swindle *SwindleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Swindle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swindle *SwindleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swindle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swindle *SwindleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swindle.Contract.contract.Transact(opts, method, params...)
}

// GetStatus is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(id bytes32) constant returns(status uint8)
func (_Swindle *SwindleCaller) GetStatus(opts *bind.CallOpts, id [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Swindle.contract.Call(opts, out, "getStatus", id)
	return *ret0, err
}

// GetStatus is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(id bytes32) constant returns(status uint8)
func (_Swindle *SwindleSession) GetStatus(id [32]byte) (uint8, error) {
	return _Swindle.Contract.GetStatus(&_Swindle.CallOpts, id)
}

// GetStatus is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(id bytes32) constant returns(status uint8)
func (_Swindle *SwindleCallerSession) GetStatus(id [32]byte) (uint8, error) {
	return _Swindle.Contract.GetStatus(&_Swindle.CallOpts, id)
}

// Ids is a free data retrieval call binding the contract method 0x9f20bc92.
//
// Solidity: function ids( address,  uint256) constant returns(bytes32)
func (_Swindle *SwindleCaller) Ids(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Swindle.contract.Call(opts, out, "ids", arg0, arg1)
	return *ret0, err
}

// Ids is a free data retrieval call binding the contract method 0x9f20bc92.
//
// Solidity: function ids( address,  uint256) constant returns(bytes32)
func (_Swindle *SwindleSession) Ids(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _Swindle.Contract.Ids(&_Swindle.CallOpts, arg0, arg1)
}

// Ids is a free data retrieval call binding the contract method 0x9f20bc92.
//
// Solidity: function ids( address,  uint256) constant returns(bytes32)
func (_Swindle *SwindleCallerSession) Ids(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _Swindle.Contract.Ids(&_Swindle.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Swindle *SwindleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Swindle.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Swindle *SwindleSession) Owner() (common.Address, error) {
	return _Swindle.Contract.Owner(&_Swindle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Swindle *SwindleCallerSession) Owner() (common.Address, error) {
	return _Swindle.Contract.Owner(&_Swindle.CallOpts)
}

// Swear is a free data retrieval call binding the contract method 0xf20f1f58.
//
// Solidity: function swear() constant returns(address)
func (_Swindle *SwindleCaller) Swear(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Swindle.contract.Call(opts, out, "swear")
	return *ret0, err
}

// Swear is a free data retrieval call binding the contract method 0xf20f1f58.
//
// Solidity: function swear() constant returns(address)
func (_Swindle *SwindleSession) Swear() (common.Address, error) {
	return _Swindle.Contract.Swear(&_Swindle.CallOpts)
}

// Swear is a free data retrieval call binding the contract method 0xf20f1f58.
//
// Solidity: function swear() constant returns(address)
func (_Swindle *SwindleCallerSession) Swear() (common.Address, error) {
	return _Swindle.Contract.Swear(&_Swindle.CallOpts)
}

// TrialRules is a free data retrieval call binding the contract method 0x8a29d969.
//
// Solidity: function trialRules() constant returns(address)
func (_Swindle *SwindleCaller) TrialRules(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Swindle.contract.Call(opts, out, "trialRules")
	return *ret0, err
}

// TrialRules is a free data retrieval call binding the contract method 0x8a29d969.
//
// Solidity: function trialRules() constant returns(address)
func (_Swindle *SwindleSession) TrialRules() (common.Address, error) {
	return _Swindle.Contract.TrialRules(&_Swindle.CallOpts)
}

// TrialRules is a free data retrieval call binding the contract method 0x8a29d969.
//
// Solidity: function trialRules() constant returns(address)
func (_Swindle *SwindleCallerSession) TrialRules() (common.Address, error) {
	return _Swindle.Contract.TrialRules(&_Swindle.CallOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_Swindle *SwindleTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Swindle.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_Swindle *SwindleSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Swindle.Contract.ChangeOwner(&_Swindle.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_Swindle *SwindleTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Swindle.Contract.ChangeOwner(&_Swindle.TransactOpts, _newOwner)
}

// NewCase is a paid mutator transaction binding the contract method 0xcc345e0a.
//
// Solidity: function newCase(serviceId bytes32) returns(bool)
func (_Swindle *SwindleTransactor) NewCase(opts *bind.TransactOpts, serviceId [32]byte) (*types.Transaction, error) {
	return _Swindle.contract.Transact(opts, "newCase", serviceId)
}

// NewCase is a paid mutator transaction binding the contract method 0xcc345e0a.
//
// Solidity: function newCase(serviceId bytes32) returns(bool)
func (_Swindle *SwindleSession) NewCase(serviceId [32]byte) (*types.Transaction, error) {
	return _Swindle.Contract.NewCase(&_Swindle.TransactOpts, serviceId)
}

// NewCase is a paid mutator transaction binding the contract method 0xcc345e0a.
//
// Solidity: function newCase(serviceId bytes32) returns(bool)
func (_Swindle *SwindleTransactorSession) NewCase(serviceId [32]byte) (*types.Transaction, error) {
	return _Swindle.Contract.NewCase(&_Swindle.TransactOpts, serviceId)
}

// Trial is a paid mutator transaction binding the contract method 0xcf343bb8.
//
// Solidity: function trial(id bytes32) returns(bool)
func (_Swindle *SwindleTransactor) Trial(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _Swindle.contract.Transact(opts, "trial", id)
}

// Trial is a paid mutator transaction binding the contract method 0xcf343bb8.
//
// Solidity: function trial(id bytes32) returns(bool)
func (_Swindle *SwindleSession) Trial(id [32]byte) (*types.Transaction, error) {
	return _Swindle.Contract.Trial(&_Swindle.TransactOpts, id)
}

// Trial is a paid mutator transaction binding the contract method 0xcf343bb8.
//
// Solidity: function trial(id bytes32) returns(bool)
func (_Swindle *SwindleTransactorSession) Trial(id [32]byte) (*types.Transaction, error) {
	return _Swindle.Contract.Trial(&_Swindle.TransactOpts, id)
}

// SwindleAbstractABI is the input ABI used to generate the binding from.
const SwindleAbstractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"serviceId\",\"type\":\"bytes32\"}],\"name\":\"newCase\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"trial\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"decide\",\"type\":\"string\"}],\"name\":\"Decision\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"plaintiff\",\"type\":\"address\"}],\"name\":\"NewCaseOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"plaintiff\",\"type\":\"address\"}],\"name\":\"NewEvidenceSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"plaintiff\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"CaseResolved\",\"type\":\"event\"}]"

// SwindleAbstractBin is the compiled bytecode used for deploying new contracts.
const SwindleAbstractBin = `0x`

// DeploySwindleAbstract deploys a new Ethereum contract, binding an instance of SwindleAbstract to it.
func DeploySwindleAbstract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SwindleAbstract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwindleAbstractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SwindleAbstractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SwindleAbstract{SwindleAbstractCaller: SwindleAbstractCaller{contract: contract}, SwindleAbstractTransactor: SwindleAbstractTransactor{contract: contract}}, nil
}

// SwindleAbstract is an auto generated Go binding around an Ethereum contract.
type SwindleAbstract struct {
	SwindleAbstractCaller     // Read-only binding to the contract
	SwindleAbstractTransactor // Write-only binding to the contract
}

// SwindleAbstractCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwindleAbstractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwindleAbstractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwindleAbstractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwindleAbstractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwindleAbstractSession struct {
	Contract     *SwindleAbstract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwindleAbstractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwindleAbstractCallerSession struct {
	Contract *SwindleAbstractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SwindleAbstractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwindleAbstractTransactorSession struct {
	Contract     *SwindleAbstractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SwindleAbstractRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwindleAbstractRaw struct {
	Contract *SwindleAbstract // Generic contract binding to access the raw methods on
}

// SwindleAbstractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwindleAbstractCallerRaw struct {
	Contract *SwindleAbstractCaller // Generic read-only contract binding to access the raw methods on
}

// SwindleAbstractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwindleAbstractTransactorRaw struct {
	Contract *SwindleAbstractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwindleAbstract creates a new instance of SwindleAbstract, bound to a specific deployed contract.
func NewSwindleAbstract(address common.Address, backend bind.ContractBackend) (*SwindleAbstract, error) {
	contract, err := bindSwindleAbstract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwindleAbstract{SwindleAbstractCaller: SwindleAbstractCaller{contract: contract}, SwindleAbstractTransactor: SwindleAbstractTransactor{contract: contract}}, nil
}

// NewSwindleAbstractCaller creates a new read-only instance of SwindleAbstract, bound to a specific deployed contract.
func NewSwindleAbstractCaller(address common.Address, caller bind.ContractCaller) (*SwindleAbstractCaller, error) {
	contract, err := bindSwindleAbstract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SwindleAbstractCaller{contract: contract}, nil
}

// NewSwindleAbstractTransactor creates a new write-only instance of SwindleAbstract, bound to a specific deployed contract.
func NewSwindleAbstractTransactor(address common.Address, transactor bind.ContractTransactor) (*SwindleAbstractTransactor, error) {
	contract, err := bindSwindleAbstract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SwindleAbstractTransactor{contract: contract}, nil
}

// bindSwindleAbstract binds a generic wrapper to an already deployed contract.
func bindSwindleAbstract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwindleAbstractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwindleAbstract *SwindleAbstractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SwindleAbstract.Contract.SwindleAbstractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwindleAbstract *SwindleAbstractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwindleAbstract.Contract.SwindleAbstractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwindleAbstract *SwindleAbstractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwindleAbstract.Contract.SwindleAbstractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwindleAbstract *SwindleAbstractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SwindleAbstract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwindleAbstract *SwindleAbstractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwindleAbstract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwindleAbstract *SwindleAbstractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwindleAbstract.Contract.contract.Transact(opts, method, params...)
}

// GetStatus is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(id bytes32) constant returns(status uint8)
func (_SwindleAbstract *SwindleAbstractCaller) GetStatus(opts *bind.CallOpts, id [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _SwindleAbstract.contract.Call(opts, out, "getStatus", id)
	return *ret0, err
}

// GetStatus is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(id bytes32) constant returns(status uint8)
func (_SwindleAbstract *SwindleAbstractSession) GetStatus(id [32]byte) (uint8, error) {
	return _SwindleAbstract.Contract.GetStatus(&_SwindleAbstract.CallOpts, id)
}

// GetStatus is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(id bytes32) constant returns(status uint8)
func (_SwindleAbstract *SwindleAbstractCallerSession) GetStatus(id [32]byte) (uint8, error) {
	return _SwindleAbstract.Contract.GetStatus(&_SwindleAbstract.CallOpts, id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SwindleAbstract *SwindleAbstractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SwindleAbstract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SwindleAbstract *SwindleAbstractSession) Owner() (common.Address, error) {
	return _SwindleAbstract.Contract.Owner(&_SwindleAbstract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SwindleAbstract *SwindleAbstractCallerSession) Owner() (common.Address, error) {
	return _SwindleAbstract.Contract.Owner(&_SwindleAbstract.CallOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_SwindleAbstract *SwindleAbstractTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _SwindleAbstract.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_SwindleAbstract *SwindleAbstractSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _SwindleAbstract.Contract.ChangeOwner(&_SwindleAbstract.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_SwindleAbstract *SwindleAbstractTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _SwindleAbstract.Contract.ChangeOwner(&_SwindleAbstract.TransactOpts, _newOwner)
}

// NewCase is a paid mutator transaction binding the contract method 0xcc345e0a.
//
// Solidity: function newCase(serviceId bytes32) returns(bool)
func (_SwindleAbstract *SwindleAbstractTransactor) NewCase(opts *bind.TransactOpts, serviceId [32]byte) (*types.Transaction, error) {
	return _SwindleAbstract.contract.Transact(opts, "newCase", serviceId)
}

// NewCase is a paid mutator transaction binding the contract method 0xcc345e0a.
//
// Solidity: function newCase(serviceId bytes32) returns(bool)
func (_SwindleAbstract *SwindleAbstractSession) NewCase(serviceId [32]byte) (*types.Transaction, error) {
	return _SwindleAbstract.Contract.NewCase(&_SwindleAbstract.TransactOpts, serviceId)
}

// NewCase is a paid mutator transaction binding the contract method 0xcc345e0a.
//
// Solidity: function newCase(serviceId bytes32) returns(bool)
func (_SwindleAbstract *SwindleAbstractTransactorSession) NewCase(serviceId [32]byte) (*types.Transaction, error) {
	return _SwindleAbstract.Contract.NewCase(&_SwindleAbstract.TransactOpts, serviceId)
}

// Trial is a paid mutator transaction binding the contract method 0xcf343bb8.
//
// Solidity: function trial(id bytes32) returns(bool)
func (_SwindleAbstract *SwindleAbstractTransactor) Trial(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _SwindleAbstract.contract.Transact(opts, "trial", id)
}

// Trial is a paid mutator transaction binding the contract method 0xcf343bb8.
//
// Solidity: function trial(id bytes32) returns(bool)
func (_SwindleAbstract *SwindleAbstractSession) Trial(id [32]byte) (*types.Transaction, error) {
	return _SwindleAbstract.Contract.Trial(&_SwindleAbstract.TransactOpts, id)
}

// Trial is a paid mutator transaction binding the contract method 0xcf343bb8.
//
// Solidity: function trial(id bytes32) returns(bool)
func (_SwindleAbstract *SwindleAbstractTransactorSession) Trial(id [32]byte) (*types.Transaction, error) {
	return _SwindleAbstract.Contract.Trial(&_SwindleAbstract.TransactOpts, id)
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

// TrialRulesAbstractABI is the input ABI used to generate the binding from.
const TrialRulesAbstractABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"getInitialStatus\",\"outputs\":[{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReward\",\"outputs\":[{\"name\":\"reward\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpoch\",\"outputs\":[{\"name\":\"epoch\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"caseId\",\"type\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"startGracePeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"witnessStatus\",\"type\":\"uint8\"},{\"name\":\"trialStatus\",\"type\":\"uint8\"}],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"trialStatus\",\"type\":\"uint8\"}],\"name\":\"getWitness\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"caseId\",\"type\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"expired\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"}]"

// TrialRulesAbstractBin is the compiled bytecode used for deploying new contracts.
const TrialRulesAbstractBin = `0x`

// DeployTrialRulesAbstract deploys a new Ethereum contract, binding an instance of TrialRulesAbstract to it.
func DeployTrialRulesAbstract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TrialRulesAbstract, error) {
	parsed, err := abi.JSON(strings.NewReader(TrialRulesAbstractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TrialRulesAbstractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TrialRulesAbstract{TrialRulesAbstractCaller: TrialRulesAbstractCaller{contract: contract}, TrialRulesAbstractTransactor: TrialRulesAbstractTransactor{contract: contract}}, nil
}

// TrialRulesAbstract is an auto generated Go binding around an Ethereum contract.
type TrialRulesAbstract struct {
	TrialRulesAbstractCaller     // Read-only binding to the contract
	TrialRulesAbstractTransactor // Write-only binding to the contract
}

// TrialRulesAbstractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TrialRulesAbstractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrialRulesAbstractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TrialRulesAbstractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrialRulesAbstractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TrialRulesAbstractSession struct {
	Contract     *TrialRulesAbstract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TrialRulesAbstractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TrialRulesAbstractCallerSession struct {
	Contract *TrialRulesAbstractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// TrialRulesAbstractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TrialRulesAbstractTransactorSession struct {
	Contract     *TrialRulesAbstractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// TrialRulesAbstractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TrialRulesAbstractRaw struct {
	Contract *TrialRulesAbstract // Generic contract binding to access the raw methods on
}

// TrialRulesAbstractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TrialRulesAbstractCallerRaw struct {
	Contract *TrialRulesAbstractCaller // Generic read-only contract binding to access the raw methods on
}

// TrialRulesAbstractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TrialRulesAbstractTransactorRaw struct {
	Contract *TrialRulesAbstractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrialRulesAbstract creates a new instance of TrialRulesAbstract, bound to a specific deployed contract.
func NewTrialRulesAbstract(address common.Address, backend bind.ContractBackend) (*TrialRulesAbstract, error) {
	contract, err := bindTrialRulesAbstract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TrialRulesAbstract{TrialRulesAbstractCaller: TrialRulesAbstractCaller{contract: contract}, TrialRulesAbstractTransactor: TrialRulesAbstractTransactor{contract: contract}}, nil
}

// NewTrialRulesAbstractCaller creates a new read-only instance of TrialRulesAbstract, bound to a specific deployed contract.
func NewTrialRulesAbstractCaller(address common.Address, caller bind.ContractCaller) (*TrialRulesAbstractCaller, error) {
	contract, err := bindTrialRulesAbstract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TrialRulesAbstractCaller{contract: contract}, nil
}

// NewTrialRulesAbstractTransactor creates a new write-only instance of TrialRulesAbstract, bound to a specific deployed contract.
func NewTrialRulesAbstractTransactor(address common.Address, transactor bind.ContractTransactor) (*TrialRulesAbstractTransactor, error) {
	contract, err := bindTrialRulesAbstract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TrialRulesAbstractTransactor{contract: contract}, nil
}

// bindTrialRulesAbstract binds a generic wrapper to an already deployed contract.
func bindTrialRulesAbstract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TrialRulesAbstractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TrialRulesAbstract *TrialRulesAbstractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TrialRulesAbstract.Contract.TrialRulesAbstractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TrialRulesAbstract *TrialRulesAbstractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrialRulesAbstract.Contract.TrialRulesAbstractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TrialRulesAbstract *TrialRulesAbstractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TrialRulesAbstract.Contract.TrialRulesAbstractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TrialRulesAbstract *TrialRulesAbstractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TrialRulesAbstract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TrialRulesAbstract *TrialRulesAbstractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrialRulesAbstract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TrialRulesAbstract *TrialRulesAbstractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TrialRulesAbstract.Contract.contract.Transact(opts, method, params...)
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(epoch uint256)
func (_TrialRulesAbstract *TrialRulesAbstractCaller) GetEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TrialRulesAbstract.contract.Call(opts, out, "getEpoch")
	return *ret0, err
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(epoch uint256)
func (_TrialRulesAbstract *TrialRulesAbstractSession) GetEpoch() (*big.Int, error) {
	return _TrialRulesAbstract.Contract.GetEpoch(&_TrialRulesAbstract.CallOpts)
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(epoch uint256)
func (_TrialRulesAbstract *TrialRulesAbstractCallerSession) GetEpoch() (*big.Int, error) {
	return _TrialRulesAbstract.Contract.GetEpoch(&_TrialRulesAbstract.CallOpts)
}

// GetReward is a free data retrieval call binding the contract method 0x3d18b912.
//
// Solidity: function getReward() constant returns(reward uint256)
func (_TrialRulesAbstract *TrialRulesAbstractCaller) GetReward(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TrialRulesAbstract.contract.Call(opts, out, "getReward")
	return *ret0, err
}

// GetReward is a free data retrieval call binding the contract method 0x3d18b912.
//
// Solidity: function getReward() constant returns(reward uint256)
func (_TrialRulesAbstract *TrialRulesAbstractSession) GetReward() (*big.Int, error) {
	return _TrialRulesAbstract.Contract.GetReward(&_TrialRulesAbstract.CallOpts)
}

// GetReward is a free data retrieval call binding the contract method 0x3d18b912.
//
// Solidity: function getReward() constant returns(reward uint256)
func (_TrialRulesAbstract *TrialRulesAbstractCallerSession) GetReward() (*big.Int, error) {
	return _TrialRulesAbstract.Contract.GetReward(&_TrialRulesAbstract.CallOpts)
}

// Expired is a paid mutator transaction binding the contract method 0xfcc0a2ae.
//
// Solidity: function expired(caseId bytes32, status uint8) returns(bool)
func (_TrialRulesAbstract *TrialRulesAbstractTransactor) Expired(opts *bind.TransactOpts, caseId [32]byte, status uint8) (*types.Transaction, error) {
	return _TrialRulesAbstract.contract.Transact(opts, "expired", caseId, status)
}

// Expired is a paid mutator transaction binding the contract method 0xfcc0a2ae.
//
// Solidity: function expired(caseId bytes32, status uint8) returns(bool)
func (_TrialRulesAbstract *TrialRulesAbstractSession) Expired(caseId [32]byte, status uint8) (*types.Transaction, error) {
	return _TrialRulesAbstract.Contract.Expired(&_TrialRulesAbstract.TransactOpts, caseId, status)
}

// Expired is a paid mutator transaction binding the contract method 0xfcc0a2ae.
//
// Solidity: function expired(caseId bytes32, status uint8) returns(bool)
func (_TrialRulesAbstract *TrialRulesAbstractTransactorSession) Expired(caseId [32]byte, status uint8) (*types.Transaction, error) {
	return _TrialRulesAbstract.Contract.Expired(&_TrialRulesAbstract.TransactOpts, caseId, status)
}

// GetInitialStatus is a paid mutator transaction binding the contract method 0x07e5ae5f.
//
// Solidity: function getInitialStatus() returns(status uint8)
func (_TrialRulesAbstract *TrialRulesAbstractTransactor) GetInitialStatus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrialRulesAbstract.contract.Transact(opts, "getInitialStatus")
}

// GetInitialStatus is a paid mutator transaction binding the contract method 0x07e5ae5f.
//
// Solidity: function getInitialStatus() returns(status uint8)
func (_TrialRulesAbstract *TrialRulesAbstractSession) GetInitialStatus() (*types.Transaction, error) {
	return _TrialRulesAbstract.Contract.GetInitialStatus(&_TrialRulesAbstract.TransactOpts)
}

// GetInitialStatus is a paid mutator transaction binding the contract method 0x07e5ae5f.
//
// Solidity: function getInitialStatus() returns(status uint8)
func (_TrialRulesAbstract *TrialRulesAbstractTransactorSession) GetInitialStatus() (*types.Transaction, error) {
	return _TrialRulesAbstract.Contract.GetInitialStatus(&_TrialRulesAbstract.TransactOpts)
}

// GetStatus is a paid mutator transaction binding the contract method 0xd39eb301.
//
// Solidity: function getStatus(witnessStatus uint8, trialStatus uint8) returns(status uint8)
func (_TrialRulesAbstract *TrialRulesAbstractTransactor) GetStatus(opts *bind.TransactOpts, witnessStatus uint8, trialStatus uint8) (*types.Transaction, error) {
	return _TrialRulesAbstract.contract.Transact(opts, "getStatus", witnessStatus, trialStatus)
}

// GetStatus is a paid mutator transaction binding the contract method 0xd39eb301.
//
// Solidity: function getStatus(witnessStatus uint8, trialStatus uint8) returns(status uint8)
func (_TrialRulesAbstract *TrialRulesAbstractSession) GetStatus(witnessStatus uint8, trialStatus uint8) (*types.Transaction, error) {
	return _TrialRulesAbstract.Contract.GetStatus(&_TrialRulesAbstract.TransactOpts, witnessStatus, trialStatus)
}

// GetStatus is a paid mutator transaction binding the contract method 0xd39eb301.
//
// Solidity: function getStatus(witnessStatus uint8, trialStatus uint8) returns(status uint8)
func (_TrialRulesAbstract *TrialRulesAbstractTransactorSession) GetStatus(witnessStatus uint8, trialStatus uint8) (*types.Transaction, error) {
	return _TrialRulesAbstract.Contract.GetStatus(&_TrialRulesAbstract.TransactOpts, witnessStatus, trialStatus)
}

// GetWitness is a paid mutator transaction binding the contract method 0xdb061563.
//
// Solidity: function getWitness(trialStatus uint8) returns(address)
func (_TrialRulesAbstract *TrialRulesAbstractTransactor) GetWitness(opts *bind.TransactOpts, trialStatus uint8) (*types.Transaction, error) {
	return _TrialRulesAbstract.contract.Transact(opts, "getWitness", trialStatus)
}

// GetWitness is a paid mutator transaction binding the contract method 0xdb061563.
//
// Solidity: function getWitness(trialStatus uint8) returns(address)
func (_TrialRulesAbstract *TrialRulesAbstractSession) GetWitness(trialStatus uint8) (*types.Transaction, error) {
	return _TrialRulesAbstract.Contract.GetWitness(&_TrialRulesAbstract.TransactOpts, trialStatus)
}

// GetWitness is a paid mutator transaction binding the contract method 0xdb061563.
//
// Solidity: function getWitness(trialStatus uint8) returns(address)
func (_TrialRulesAbstract *TrialRulesAbstractTransactorSession) GetWitness(trialStatus uint8) (*types.Transaction, error) {
	return _TrialRulesAbstract.Contract.GetWitness(&_TrialRulesAbstract.TransactOpts, trialStatus)
}

// StartGracePeriod is a paid mutator transaction binding the contract method 0xa7edcbe4.
//
// Solidity: function startGracePeriod(caseId bytes32, status uint8) returns(bool)
func (_TrialRulesAbstract *TrialRulesAbstractTransactor) StartGracePeriod(opts *bind.TransactOpts, caseId [32]byte, status uint8) (*types.Transaction, error) {
	return _TrialRulesAbstract.contract.Transact(opts, "startGracePeriod", caseId, status)
}

// StartGracePeriod is a paid mutator transaction binding the contract method 0xa7edcbe4.
//
// Solidity: function startGracePeriod(caseId bytes32, status uint8) returns(bool)
func (_TrialRulesAbstract *TrialRulesAbstractSession) StartGracePeriod(caseId [32]byte, status uint8) (*types.Transaction, error) {
	return _TrialRulesAbstract.Contract.StartGracePeriod(&_TrialRulesAbstract.TransactOpts, caseId, status)
}

// StartGracePeriod is a paid mutator transaction binding the contract method 0xa7edcbe4.
//
// Solidity: function startGracePeriod(caseId bytes32, status uint8) returns(bool)
func (_TrialRulesAbstract *TrialRulesAbstractTransactorSession) StartGracePeriod(caseId [32]byte, status uint8) (*types.Transaction, error) {
	return _TrialRulesAbstract.Contract.StartGracePeriod(&_TrialRulesAbstract.TransactOpts, caseId, status)
}

// WitnessAbstractABI is the input ABI used to generate the binding from.
const WitnessAbstractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"caseId\",\"type\":\"bytes32\"},{\"name\":\"serviceId\",\"type\":\"bytes32\"},{\"name\":\"clientAddress\",\"type\":\"address\"}],\"name\":\"isEvidenceSubmitted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"caseId\",\"type\":\"bytes32\"},{\"name\":\"serviceId\",\"type\":\"bytes32\"},{\"name\":\"clientAddress\",\"type\":\"address\"}],\"name\":\"testimonyFor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"}]"

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

// IsEvidenceSubmitted is a paid mutator transaction binding the contract method 0x7aba9a67.
//
// Solidity: function isEvidenceSubmitted(caseId bytes32, serviceId bytes32, clientAddress address) returns(bool)
func (_WitnessAbstract *WitnessAbstractTransactor) IsEvidenceSubmitted(opts *bind.TransactOpts, caseId [32]byte, serviceId [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _WitnessAbstract.contract.Transact(opts, "isEvidenceSubmitted", caseId, serviceId, clientAddress)
}

// IsEvidenceSubmitted is a paid mutator transaction binding the contract method 0x7aba9a67.
//
// Solidity: function isEvidenceSubmitted(caseId bytes32, serviceId bytes32, clientAddress address) returns(bool)
func (_WitnessAbstract *WitnessAbstractSession) IsEvidenceSubmitted(caseId [32]byte, serviceId [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _WitnessAbstract.Contract.IsEvidenceSubmitted(&_WitnessAbstract.TransactOpts, caseId, serviceId, clientAddress)
}

// IsEvidenceSubmitted is a paid mutator transaction binding the contract method 0x7aba9a67.
//
// Solidity: function isEvidenceSubmitted(caseId bytes32, serviceId bytes32, clientAddress address) returns(bool)
func (_WitnessAbstract *WitnessAbstractTransactorSession) IsEvidenceSubmitted(caseId [32]byte, serviceId [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _WitnessAbstract.Contract.IsEvidenceSubmitted(&_WitnessAbstract.TransactOpts, caseId, serviceId, clientAddress)
}

// TestimonyFor is a paid mutator transaction binding the contract method 0xc4b45283.
//
// Solidity: function testimonyFor(caseId bytes32, serviceId bytes32, clientAddress address) returns(uint8)
func (_WitnessAbstract *WitnessAbstractTransactor) TestimonyFor(opts *bind.TransactOpts, caseId [32]byte, serviceId [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _WitnessAbstract.contract.Transact(opts, "testimonyFor", caseId, serviceId, clientAddress)
}

// TestimonyFor is a paid mutator transaction binding the contract method 0xc4b45283.
//
// Solidity: function testimonyFor(caseId bytes32, serviceId bytes32, clientAddress address) returns(uint8)
func (_WitnessAbstract *WitnessAbstractSession) TestimonyFor(caseId [32]byte, serviceId [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _WitnessAbstract.Contract.TestimonyFor(&_WitnessAbstract.TransactOpts, caseId, serviceId, clientAddress)
}

// TestimonyFor is a paid mutator transaction binding the contract method 0xc4b45283.
//
// Solidity: function testimonyFor(caseId bytes32, serviceId bytes32, clientAddress address) returns(uint8)
func (_WitnessAbstract *WitnessAbstractTransactorSession) TestimonyFor(caseId [32]byte, serviceId [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _WitnessAbstract.Contract.TestimonyFor(&_WitnessAbstract.TransactOpts, caseId, serviceId, clientAddress)
}
