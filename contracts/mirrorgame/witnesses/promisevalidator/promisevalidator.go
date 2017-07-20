// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package promisevalidator

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
const OwnedBin = `0x6060604052341561000f57600080fd5b5b60008054600160a060020a03191633600160a060020a03161790555b5b61015c8061003c6000396000f300606060405263ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b8114610048578063a6f9dae114610084575b600080fd5b341561005357600080fd5b61005b6100b2565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561008f57600080fd5b6100b073ffffffffffffffffffffffffffffffffffffffff600435166100ce565b005b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6000543373ffffffffffffffffffffffffffffffffffffffff9081169116146100f657600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff83161790555b5b505600a165627a7a72305820e6722f1c05aff07afcd7e7d2fa4566661cb0bec3fd7ebfb0fb9186206656d03d0029`

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

// PromiseValidatorABI is the input ABI used to generate the binding from.
const PromiseValidatorABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"caseId\",\"type\":\"bytes32\"},{\"name\":\"serviceId\",\"type\":\"bytes32\"},{\"name\":\"beneficiary\",\"type\":\"address\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"name\":\"sig_v\",\"type\":\"uint8\"},{\"name\":\"sig_r\",\"type\":\"bytes32\"},{\"name\":\"sig_s\",\"type\":\"bytes32\"}],\"name\":\"submitPromise\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"caseId\",\"type\":\"bytes32\"},{\"name\":\"serviceId\",\"type\":\"bytes32\"},{\"name\":\"clientAddress\",\"type\":\"address\"}],\"name\":\"testimonyFor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"}]"

// PromiseValidatorBin is the compiled bytecode used for deploying new contracts.
const PromiseValidatorBin = `0x60606040526023600155341561001457600080fd5b5b5b60008054600160a060020a03191633600160a060020a03161790555b5b5b610505806100436000396000f300606060405263ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b811461005e578063a6f9dae11461008d578063b31e6b91146100ae578063c4b45283146100f6575b600080fd5b341561006957600080fd5b61007161013f565b604051600160a060020a03909116815260200160405180910390f35b341561009857600080fd5b6100ac600160a060020a036004351661014e565b005b34156100b957600080fd5b6100e2600435602435600160a060020a036044351660643560ff6084351660a43560c435610196565b604051901515815260200160405180910390f35b341561010157600080fd5b61011b600435602435600160a060020a0360443516610294565b6040518082600281111561012b57fe5b60ff16815260200191505060405180910390f35b600054600160a060020a031681565b60005433600160a060020a0390811691161461016957600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b50565b60006101a06104a4565b6101a989610374565b156101b75760009150610288565b600160a060020a038716808252602080830188905260ff87166040808501919091526060840187905260808401869052600160a085015260008b81526002835281812093815292909152902081908151815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039190911617815560208201518160010155604082015160028201805460ff191660ff92909216919091179055606082015160038201556080820151600482015560a0820151600591909101805460ff191691151591909117905550600191505b50979650505050505050565b6000828152600260209081526040808320600160a060020a038516845290915281206005015460ff16151561030d5760008481526003602052604090205415156102ee576000848152600360205260409020439055610304565b6102f784610374565b156103045750600161036d565b5b50600261036d565b6000838152600260208181526040808420600160a060020a03878116865292529092208054600182015492820154600383015460049093015461035b9592909216939260ff909116916103b4565b15156103695750600161036d565b5060005b9392505050565b600081815260036020526040812054156103aa57600154600083815260036020526040902054430311156103aa575060016103af565b5b5060005b919050565b600080438690106103c8576000915061049a565b3087876040516c01000000000000000000000000600160a060020a03948516810282529290931690910260148301526028820152604801604051809103902090506001818686866040518060005260200160405260006040516020015260405193845260ff90921660208085019190915260408085019290925260608401929092526080909201915160208103908084039060008661646e5a03f1151561046e57600080fd5b505060206040510351600054600160a060020a03908116911614610495576000915061049a565b600191505b5095945050505050565b60c06040519081016040908152600080835260208301819052908201819052606082018190526080820181905260a0820152905600a165627a7a72305820ce17601908c845e332ddb2db55059c9c442c155c8203d2fd9aabe6a4e31a7c3e0029`

// DeployPromiseValidator deploys a new Ethereum contract, binding an instance of PromiseValidator to it.
func DeployPromiseValidator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PromiseValidator, error) {
	parsed, err := abi.JSON(strings.NewReader(PromiseValidatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PromiseValidatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PromiseValidator{PromiseValidatorCaller: PromiseValidatorCaller{contract: contract}, PromiseValidatorTransactor: PromiseValidatorTransactor{contract: contract}}, nil
}

// PromiseValidator is an auto generated Go binding around an Ethereum contract.
type PromiseValidator struct {
	PromiseValidatorCaller     // Read-only binding to the contract
	PromiseValidatorTransactor // Write-only binding to the contract
}

// PromiseValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type PromiseValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PromiseValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PromiseValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PromiseValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PromiseValidatorSession struct {
	Contract     *PromiseValidator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PromiseValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PromiseValidatorCallerSession struct {
	Contract *PromiseValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// PromiseValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PromiseValidatorTransactorSession struct {
	Contract     *PromiseValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// PromiseValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type PromiseValidatorRaw struct {
	Contract *PromiseValidator // Generic contract binding to access the raw methods on
}

// PromiseValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PromiseValidatorCallerRaw struct {
	Contract *PromiseValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// PromiseValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PromiseValidatorTransactorRaw struct {
	Contract *PromiseValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPromiseValidator creates a new instance of PromiseValidator, bound to a specific deployed contract.
func NewPromiseValidator(address common.Address, backend bind.ContractBackend) (*PromiseValidator, error) {
	contract, err := bindPromiseValidator(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PromiseValidator{PromiseValidatorCaller: PromiseValidatorCaller{contract: contract}, PromiseValidatorTransactor: PromiseValidatorTransactor{contract: contract}}, nil
}

// NewPromiseValidatorCaller creates a new read-only instance of PromiseValidator, bound to a specific deployed contract.
func NewPromiseValidatorCaller(address common.Address, caller bind.ContractCaller) (*PromiseValidatorCaller, error) {
	contract, err := bindPromiseValidator(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &PromiseValidatorCaller{contract: contract}, nil
}

// NewPromiseValidatorTransactor creates a new write-only instance of PromiseValidator, bound to a specific deployed contract.
func NewPromiseValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*PromiseValidatorTransactor, error) {
	contract, err := bindPromiseValidator(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &PromiseValidatorTransactor{contract: contract}, nil
}

// bindPromiseValidator binds a generic wrapper to an already deployed contract.
func bindPromiseValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PromiseValidatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PromiseValidator *PromiseValidatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PromiseValidator.Contract.PromiseValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PromiseValidator *PromiseValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PromiseValidator.Contract.PromiseValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PromiseValidator *PromiseValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PromiseValidator.Contract.PromiseValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PromiseValidator *PromiseValidatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PromiseValidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PromiseValidator *PromiseValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PromiseValidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PromiseValidator *PromiseValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PromiseValidator.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PromiseValidator *PromiseValidatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PromiseValidator.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PromiseValidator *PromiseValidatorSession) Owner() (common.Address, error) {
	return _PromiseValidator.Contract.Owner(&_PromiseValidator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PromiseValidator *PromiseValidatorCallerSession) Owner() (common.Address, error) {
	return _PromiseValidator.Contract.Owner(&_PromiseValidator.CallOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_PromiseValidator *PromiseValidatorTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _PromiseValidator.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_PromiseValidator *PromiseValidatorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _PromiseValidator.Contract.ChangeOwner(&_PromiseValidator.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_PromiseValidator *PromiseValidatorTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _PromiseValidator.Contract.ChangeOwner(&_PromiseValidator.TransactOpts, _newOwner)
}

// SubmitPromise is a paid mutator transaction binding the contract method 0xb31e6b91.
//
// Solidity: function submitPromise(caseId bytes32, serviceId bytes32, beneficiary address, blockNumber uint256, sig_v uint8, sig_r bytes32, sig_s bytes32) returns(bool)
func (_PromiseValidator *PromiseValidatorTransactor) SubmitPromise(opts *bind.TransactOpts, caseId [32]byte, serviceId [32]byte, beneficiary common.Address, blockNumber *big.Int, sig_v uint8, sig_r [32]byte, sig_s [32]byte) (*types.Transaction, error) {
	return _PromiseValidator.contract.Transact(opts, "submitPromise", caseId, serviceId, beneficiary, blockNumber, sig_v, sig_r, sig_s)
}

// SubmitPromise is a paid mutator transaction binding the contract method 0xb31e6b91.
//
// Solidity: function submitPromise(caseId bytes32, serviceId bytes32, beneficiary address, blockNumber uint256, sig_v uint8, sig_r bytes32, sig_s bytes32) returns(bool)
func (_PromiseValidator *PromiseValidatorSession) SubmitPromise(caseId [32]byte, serviceId [32]byte, beneficiary common.Address, blockNumber *big.Int, sig_v uint8, sig_r [32]byte, sig_s [32]byte) (*types.Transaction, error) {
	return _PromiseValidator.Contract.SubmitPromise(&_PromiseValidator.TransactOpts, caseId, serviceId, beneficiary, blockNumber, sig_v, sig_r, sig_s)
}

// SubmitPromise is a paid mutator transaction binding the contract method 0xb31e6b91.
//
// Solidity: function submitPromise(caseId bytes32, serviceId bytes32, beneficiary address, blockNumber uint256, sig_v uint8, sig_r bytes32, sig_s bytes32) returns(bool)
func (_PromiseValidator *PromiseValidatorTransactorSession) SubmitPromise(caseId [32]byte, serviceId [32]byte, beneficiary common.Address, blockNumber *big.Int, sig_v uint8, sig_r [32]byte, sig_s [32]byte) (*types.Transaction, error) {
	return _PromiseValidator.Contract.SubmitPromise(&_PromiseValidator.TransactOpts, caseId, serviceId, beneficiary, blockNumber, sig_v, sig_r, sig_s)
}

// TestimonyFor is a paid mutator transaction binding the contract method 0xc4b45283.
//
// Solidity: function testimonyFor(caseId bytes32, serviceId bytes32, clientAddress address) returns(uint8)
func (_PromiseValidator *PromiseValidatorTransactor) TestimonyFor(opts *bind.TransactOpts, caseId [32]byte, serviceId [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _PromiseValidator.contract.Transact(opts, "testimonyFor", caseId, serviceId, clientAddress)
}

// TestimonyFor is a paid mutator transaction binding the contract method 0xc4b45283.
//
// Solidity: function testimonyFor(caseId bytes32, serviceId bytes32, clientAddress address) returns(uint8)
func (_PromiseValidator *PromiseValidatorSession) TestimonyFor(caseId [32]byte, serviceId [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _PromiseValidator.Contract.TestimonyFor(&_PromiseValidator.TransactOpts, caseId, serviceId, clientAddress)
}

// TestimonyFor is a paid mutator transaction binding the contract method 0xc4b45283.
//
// Solidity: function testimonyFor(caseId bytes32, serviceId bytes32, clientAddress address) returns(uint8)
func (_PromiseValidator *PromiseValidatorTransactorSession) TestimonyFor(caseId [32]byte, serviceId [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _PromiseValidator.Contract.TestimonyFor(&_PromiseValidator.TransactOpts, caseId, serviceId, clientAddress)
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
