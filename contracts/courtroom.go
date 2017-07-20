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

// OwnedABI is the input ABI used to generate the binding from.
const OwnedABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"}]"

// OwnedBin is the compiled bytecode used for deploying new contracts.
const OwnedBin = `0x6060604052341561000f57600080fd5b5b60008054600160a060020a03191633600160a060020a03161790555b5b61015c8061003c6000396000f300606060405263ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b8114610048578063a6f9dae114610084575b600080fd5b341561005357600080fd5b61005b6100b2565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561008f57600080fd5b6100b073ffffffffffffffffffffffffffffffffffffffff600435166100ce565b005b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6000543373ffffffffffffffffffffffffffffffffffffffff9081169116146100f657600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff83161790555b5b505600a165627a7a72305820c315c10434405183bfa44fb90a6dbb5f4ba839b9ff43887b4ff585950b4eb6070029`

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
const SampleTokenBin = `0x6060604052341561000f57600080fd5b60405160208061069d833981016040528080519150505b5b60038054600160a060020a03191633600160a060020a03161790555b600160a060020a03331660009081526001602052604081208290558190555b505b61062a806100736000396000f300606060405236156100965763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461009b57806318160ddd146100d157806323b872dd146100f657806370a08231146101325780638da5cb5b14610163578063a3aff92314610192578063a6f9dae1146101b9578063a9059cbb146101da578063dd62ed3e14610210575b600080fd5b34156100a657600080fd5b6100bd600160a060020a0360043516602435610247565b604051901515815260200160405180910390f35b34156100dc57600080fd5b6100e46102b4565b60405190815260200160405180910390f35b341561010157600080fd5b6100bd600160a060020a03600435811690602435166044356102ba565b604051901515815260200160405180910390f35b341561013d57600080fd5b6100e4600160a060020a03600435166102fd565b60405190815260200160405180910390f35b341561016e57600080fd5b61017661031c565b604051600160a060020a03909116815260200160405180910390f35b341561019d57600080fd5b6101b7600160a060020a036004351660243560443561032b565b005b34156101c457600080fd5b6101b7600160a060020a03600435166103e8565b005b34156101e557600080fd5b6100bd600160a060020a0360043516602435610430565b604051901515815260200160405180910390f35b341561021b57600080fd5b6100e4600160a060020a03600435811690602435166104d7565b60405190815260200160405180910390f35b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a35060015b92915050565b60005481565b600160a060020a0380841660009081526002602090815260408083203390941683529290529081208054830190556102f3848484610504565b90505b9392505050565b600160a060020a0381166000908152600160205260409020545b919050565b600354600160a060020a031681565b60035433600160a060020a0390811691161461034657600080fd5b600160a060020a038316600081815260016020526040808220805486019055815485019091558291907fdfb81fb379557413b0a951b4d7bf7a9df393801d8c539d5e201d6a8daeb913b99085905190815260200160405180910390a382600160a060020a031660007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a35b5b505050565b60035433600160a060020a0390811691161461040357600080fd5b6003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b50565b600160a060020a0333166000908152600160205260408120548290108015906104595750600082115b1561009657600160a060020a033381166000818152600160205260408082208054879003905592861680825290839020805486019055917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060016102ae565b600080fd5b5b92915050565b600160a060020a038083166000908152600260209081526040808320938516835292905220545b92915050565b600160a060020a0383166000908152600160205260408120548290108015906105545750600160a060020a0380851660009081526002602090815260408083203390941683529290522054829010155b80156105605750600082115b1561009657600160a060020a03808416600081815260016020908152604080832080548801905588851680845281842080548990039055600283528184203390961684529490915290819020805486900390559091907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060016102f6565b600080fd5b5b93925050505600a165627a7a723058206c439f13ad6f6e6f1dfb80d1f8376937ff17ca98344300189358cec1203c21a60029`

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
const StandardTokenBin = `0x6060604052341561000f57600080fd5b5b61043b8061001f6000396000f300606060405236156100755763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461007a57806318160ddd146100b057806323b872dd146100d557806370a0823114610111578063a9059cbb14610142578063dd62ed3e14610178575b600080fd5b341561008557600080fd5b61009c600160a060020a03600435166024356101af565b604051901515815260200160405180910390f35b34156100bb57600080fd5b6100c361021c565b60405190815260200160405180910390f35b34156100e057600080fd5b61009c600160a060020a0360043581169060243516604435610222565b604051901515815260200160405180910390f35b341561011c57600080fd5b6100c3600160a060020a036004351661031c565b60405190815260200160405180910390f35b341561014d57600080fd5b61009c600160a060020a036004351660243561033b565b604051901515815260200160405180910390f35b341561018357600080fd5b6100c3600160a060020a03600435811690602435166103e2565b60405190815260200160405180910390f35b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a35060015b92915050565b60005481565b600160a060020a0383166000908152600160205260408120548290108015906102725750600160a060020a0380851660009081526002602090815260408083203390941683529290522054829010155b801561027e5750600082115b1561007557600160a060020a03808416600081815260016020908152604080832080548801905588851680845281842080548990039055600283528184203390961684529490915290819020805486900390559091907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a3506001610314565b600080fd5b5b9392505050565b600160a060020a0381166000908152600160205260409020545b919050565b600160a060020a0333166000908152600160205260408120548290108015906103645750600082115b1561007557600160a060020a033381166000818152600160205260408082208054879003905592861680825290839020805486019055917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a3506001610216565b600080fd5b5b92915050565b600160a060020a038083166000908152600260209081526040808320938516835292905220545b929150505600a165627a7a7230582048977c08230929ef6eb2126bc965691030242f511f3ecd7f177cf1b4c240f64e0029`

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
const SwearGameABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"reward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"playerCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_player\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[{\"name\":\"registered\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ids\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"trialTransistions\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"serviceId\",\"type\":\"bytes32\"}],\"name\":\"newCase\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"trial\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"players\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_trialTransistions\",\"type\":\"address\"},{\"name\":\"_reward\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"constructor\"},{\"payable\":true,\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"decide\",\"type\":\"string\"}],\"name\":\"Decision\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"depositAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"DepositStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"Compensate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"playerId\",\"type\":\"address\"}],\"name\":\"NewPlayer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"playerId\",\"type\":\"address\"}],\"name\":\"PlayerLeftGame\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"plaintiff\",\"type\":\"address\"}],\"name\":\"NewClaimOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"plaintiff\",\"type\":\"address\"}],\"name\":\"NewEvidenceSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"plaintiff\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"ClaimResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Payment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"AdditionalDepositRequired\",\"type\":\"event\"}]"

// SwearGameBin is the compiled bytecode used for deploying new contracts.
const SwearGameBin = `0x6060604052341561000f57600080fd5b60405160608061104d8339810160405280805191906020018051919060200180519150505b5b60008054600160a060020a03191633600160a060020a03161790555b60048054600160a060020a03808616600160a060020a0319928316179092556005805492851692909116919091179055600281905560006003555b5050505b610fae8061009f6000396000f300606060405236156100a95763ffffffff60e060020a600035041663228cb7338114610197578063302bcc57146101bc5780634420e486146101e15780635de28ae0146102145780638da5cb5b146102405780639f20bc921461026f578063a6f9dae1146102a3578063c41a3477146102c4578063cc345e0a146102f3578063cf343bb81461031d578063d0e30db014610347578063e2eb41ff1461036c578063fc0c546a1461039f575b6101955b600454600080543492600160a060020a03908116926323b872dd9290911690309085906040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b151561012657600080fd5b6102c65a03f1151561013757600080fd5b50505060405180519050151561014c57600080fd5b600180548201908190557f3a7e173a9698235104076a0f536c5169527adecd20d0938c7c156ddf776c859190829060405191825260208201526040908101905180910390a15b50565b005b34156101a257600080fd5b6101aa6103ce565b60405190815260200160405180910390f35b34156101c757600080fd5b6101aa6103d4565b60405190815260200160405180910390f35b34156101ec57600080fd5b610200600160a060020a03600435166103da565b604051901515815260200160405180910390f35b341561021f57600080fd5b61022a600435610507565b60405160ff909116815260200160405180910390f35b341561024b57600080fd5b610253610522565b604051600160a060020a03909116815260200160405180910390f35b341561027a57600080fd5b6101aa600160a060020a0360043516602435610531565b60405190815260200160405180910390f35b34156102ae57600080fd5b610195600160a060020a0360043516610563565b005b34156102cf57600080fd5b6102536105ab565b604051600160a060020a03909116815260200160405180910390f35b34156102fe57600080fd5b6102006004356105ba565b604051901515815260200160405180910390f35b341561032857600080fd5b61020060043561070f565b604051901515815260200160405180910390f35b341561035257600080fd5b6101aa61075c565b60405190815260200160405180910390f35b341561037757600080fd5b610200600160a060020a0360043516610762565b604051901515815260200160405180910390f35b34156103aa57600080fd5b610253610777565b604051600160a060020a03909116815260200160405180910390f35b60025481565b60035481565b6000805433600160a060020a039081169116146103f657600080fd5b600160a060020a03821660009081526007602052604090205460ff161561041c57600080fd5b600354151561043b57600254600154101561043657600080fd5b61048e565b60025460035460015481151561044d57fe5b04101561048e577fa8ff0fbf8dc82e5914523ed3a828920f55ab93494c59933e8c45d79f8bba70d760015460405190815260200160405180910390a1600080fd5b5b600160a060020a03821660009081526007602052604090819020805460ff191660019081179091556003805490910190557f52e92d4898337244a39bd42674ac561eadfd3959e947deec1c0ab82dd58b5a7590839051600160a060020a03909116815260200160405180910390a15060015b5b919050565b60008181526006602052604090206003015460ff165b919050565b600054600160a060020a031681565b60086020528160005260406000208181548110151561054c57fe5b906000526020600020900160005b91509150505481565b60005433600160a060020a0390811691161461057e57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b50565b600554600160a060020a031681565b600160a060020a033316600090815260076020526040812054819060ff1615156105e357600080fd5b6005546106b79033908590600160a060020a03166307e5ae5f6000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561063357600080fd5b6102c65a03f1151561064457600080fd5b5050506040518051600554909150600160a060020a031663ada016106000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561069757600080fd5b6102c65a03f115156106a857600080fd5b50505060405180519050610786565b90508015156106c95760009150610709565b600160a060020a03331660009081526008602052604090208054600181016106f18382610f02565b916000526020600020900160005b5082905550600191505b50919050565b600160a060020a03331660009081526007602052604081205460ff16151561073657600080fd5b61073f826108d0565b151561074a57600080fd5b610753826108ed565b5060015b919050565b60015481565b60076020526000908152604090205460ff1681565b600454600160a060020a031681565b6000610790610f2c565b858542604051600160a060020a03939093166c01000000000000000000000000028352601483019190915260348201526054016040519081900390208152600660008251815260208101919091526040016000206005015460ff16156107f957600091506108c7565b600160a060020a03861660208201526040810185905260ff84166060820152600160a08201524383016080820152806006600082518152602081019190915260400160002081518155602082015160018201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039290921691909117905560408201516002820155606082015160038201805460ff191660ff929092169190911790556080820151816004015560a0820151600591909101805460ff191660ff90921691909117905550805191505b50949350505050565b60008181526006602052604090206005015460ff1615155b919050565b6000806000806000806108ff87610507565b955061090a87610b5f565b9196509450925060005b60ff168660ff16141561092657610b55565b5b60005b60ff878116911614610b555761093f83610b92565b15610956576109508760005b610b9c565b50610b55565b600554600160a060020a031663db0615638760006040516020015260405160e060020a63ffffffff841602815260ff9091166004820152602401602060405180830381600087803b15156109a957600080fd5b6102c65a03f115156109ba57600080fd5b5050506040518051925050600160a060020a03821615156109e4576109dd610be1565b9050610b55565b81600160a060020a031663c4b4528388868860006040516020015260405160e060020a63ffffffff861602815260048101939093526024830191909152600160a060020a03166044820152606401602060405180830381600087803b1515610a4b57600080fd5b6102c65a03f11515610a5c57600080fd5b50505060405180519150505b60025b816002811115610a7757fe5b1415610a8257610b55565b600554600160a060020a031663d39eb301826002811115610a9f57fe5b8860006040516020015260405160e060020a63ffffffff851602815260ff928316600482015291166024820152604401602060405180830381600087803b1515610ae857600080fd5b6102c65a03f11515610af957600080fd5b505050604051805190509550610b0f8787610b9c565b5060015b60ff168660ff161480610b2d575060025b60ff168660ff16145b15610b4f57610b3d878787610be7565b5060005b9550610b4d8787610b9c565b505b5b610926565b5b50505050505050565b600081815260066020526040902060018101546002820154600490920154600160a060020a0390911691905b9193909250565b438190115b919050565b6000805433600160a060020a0390811691161415610bb957600080fd5b506000828152600660205260409020600301805460ff191660ff831617905560015b92915050565b60025b90565b60008060025b60ff168460ff161415610c5e577f2f18444e939f85d10c4f7ad5f7de36bc69f5d99ff0e6614a6accbb2f0118e4778584600087604051938452600160a060020a03909216602084015260408084019190915260ff90911660608301526080909101905180910390a160009150610cdb565b610c6783610ce3565b9050610c7285610dd6565b610c7b83610e70565b7f2f18444e939f85d10c4f7ad5f7de36bc69f5d99ff0e6614a6accbb2f0118e477858460025487604051938452600160a060020a03909216602084015260408084019190915260ff90911660608301526080909101905180910390a18091505b509392505050565b600454600254600091600160a060020a0316906323b872dd9030908590856040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b1515610d5757600080fd5b6102c65a03f11515610d6857600080fd5b5050506040518051915050801515610d7f57600080fd5b6002546001805482900390557f92c2a7173158b7618078365b4ad89fd1f774ae4aa04f39e10b966b47f469d34b908390604051600160a060020a03909216825260208201526040908101905180910390a15b919050565b6000818152600660205260409020541580610e03575060008181526006602052604090206003015460ff16155b15610e0d57600080fd5b600081815260066020526040812081815560018101805473ffffffffffffffffffffffffffffffffffffffff19169055600501805460ff191690555b6000828152600660205260409020600301805460ff191660ff929092169190911790555b50565b600160a060020a03811660009081526007602052604090205460ff161515610e9757600080fd5b7f3def0aea61e344e47b98a0523a5825c0919e88f7b33d0ac7f889a7654880f9a581604051600160a060020a03909116815260200160405180910390a1600160a060020a0381166000908152600760205260409020805460ff19169055600380546000190190555b50565b815481835581811511610f2657600083815260209020610f26918101908301610f61565b5b505050565b60c06040519081016040908152600080835260208301819052908201819052606082018190526080820181905260a082015290565b610be491905b80821115610f7b5760008155600101610f67565b5090565b905600a165627a7a723058205903db143d81442f78346c84390b8cfff2eb9e8357fa3d1291addd77a8b6664e0029`

// DeploySwearGame deploys a new Ethereum contract, binding an instance of SwearGame to it.
func DeploySwearGame(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address, _trialTransistions common.Address, _reward *big.Int) (common.Address, *types.Transaction, *SwearGame, error) {
	parsed, err := abi.JSON(strings.NewReader(SwearGameABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SwearGameBin), backend, _token, _trialTransistions, _reward)
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

// Trial is a paid mutator transaction binding the contract method 0xcf343bb8.
//
// Solidity: function trial(id bytes32) returns(bool)
func (_SwearGame *SwearGameTransactor) Trial(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _SwearGame.contract.Transact(opts, "trial", id)
}

// Trial is a paid mutator transaction binding the contract method 0xcf343bb8.
//
// Solidity: function trial(id bytes32) returns(bool)
func (_SwearGame *SwearGameSession) Trial(id [32]byte) (*types.Transaction, error) {
	return _SwearGame.Contract.Trial(&_SwearGame.TransactOpts, id)
}

// Trial is a paid mutator transaction binding the contract method 0xcf343bb8.
//
// Solidity: function trial(id bytes32) returns(bool)
func (_SwearGame *SwearGameTransactorSession) Trial(id [32]byte) (*types.Transaction, error) {
	return _SwearGame.Contract.Trial(&_SwearGame.TransactOpts, id)
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
const WitnessAbstractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"caseId\",\"type\":\"bytes32\"},{\"name\":\"serviceId\",\"type\":\"bytes32\"},{\"name\":\"clientAddress\",\"type\":\"address\"}],\"name\":\"testimonyFor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"}]"

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
