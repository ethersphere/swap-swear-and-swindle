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
const OwnedBin = `0x6060604052341561000f57600080fd5b5b60008054600160a060020a03191633600160a060020a03161790555b5b61015c8061003c6000396000f300606060405263ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b8114610048578063a6f9dae114610084575b600080fd5b341561005357600080fd5b61005b6100b2565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561008f57600080fd5b6100b073ffffffffffffffffffffffffffffffffffffffff600435166100ce565b005b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6000543373ffffffffffffffffffffffffffffffffffffffff9081169116146100f657600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff83161790555b5b505600a165627a7a72305820e6878ca472b720e2171b0a310985784b387f9601d878c6f5e89f0fde37feb1320029`

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
const PromiseValidatorABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"caseId\",\"type\":\"bytes32\"},{\"name\":\"serviceId\",\"type\":\"bytes32\"},{\"name\":\"clientAddress\",\"type\":\"address\"}],\"name\":\"isEvidentSubmited\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"caseId\",\"type\":\"bytes32\"},{\"name\":\"serviceId\",\"type\":\"bytes32\"},{\"name\":\"beneficiary\",\"type\":\"address\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"name\":\"sig_v\",\"type\":\"uint8\"},{\"name\":\"sig_r\",\"type\":\"bytes32\"},{\"name\":\"sig_s\",\"type\":\"bytes32\"}],\"name\":\"submitPromise\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"caseId\",\"type\":\"bytes32\"},{\"name\":\"serviceId\",\"type\":\"bytes32\"},{\"name\":\"clientAddress\",\"type\":\"address\"}],\"name\":\"testimonyFor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"}]"

// PromiseValidatorBin is the compiled bytecode used for deploying new contracts.
const PromiseValidatorBin = `0x6060604052341561000f57600080fd5b5b5b60008054600160a060020a03191633600160a060020a03161790555b5b5b6104c48061003e6000396000f300606060405263ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166370e39acd81146100695780638da5cb5b146100a2578063a6f9dae1146100d1578063b31e6b91146100f2578063c4b452831461013a575b600080fd5b341561007457600080fd5b61008e600435602435600160a060020a0360443516610183565b604051901515815260200160405180910390f35b34156100ad57600080fd5b6100b56101bc565b604051600160a060020a03909116815260200160405180910390f35b34156100dc57600080fd5b6100f0600160a060020a03600435166101cb565b005b34156100fd57600080fd5b61008e600435602435600160a060020a036044351660643560ff6084351660a43560c435610213565b604051901515815260200160405180910390f35b341561014557600080fd5b61015f600435602435600160a060020a036044351661033a565b6040518082600281111561016f57fe5b60ff16815260200191505060405180910390f35b60008381526001602090815260408083208584528252808320600160a060020a038516845290915290206005015460ff165b9392505050565b600054600160a060020a031681565b60005433600160a060020a039081169116146101e657600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b50565b60008781526001602090815260408083208984528252808320600160a060020a038916845290915281206005015460ff16156102515750600061032f565b60c06040519081016040908152600160a060020a038816808352602080840189905260ff8816838501526060840187905260808401869052600160a0850181905260008d81529082528381208c82528252838120928152919052208151815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039190911617815560208201518160010155604082015160028201805460ff191660ff92909216919091179055606082015160038201556080820151600482015560a0820151600591909101805460ff191691151591909117905550600190505b979650505050505050565b60008381526001602081815260408084208685528252808420600160a060020a038681168652925283208054928101546002820154600383015460049093015461038f9590941693919260ff909116916103a8565b151561039d575060016101b5565b5060005b9392505050565b600080438690106103bc576000915061048e565b3087876040516c01000000000000000000000000600160a060020a03948516810282529290931690910260148301526028820152604801604051809103902090506001818686866040518060005260200160405260006040516020015260405193845260ff90921660208085019190915260408085019290925260608401929092526080909201915160208103908084039060008661646e5a03f1151561046257600080fd5b505060206040510351600054600160a060020a03908116911614610489576000915061048e565b600191505b50959450505050505600a165627a7a72305820dfff739e677d955de497444f0f6fa738e2877672cd910bd5b014b918dc29b2310029`

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

// IsEvidentSubmited is a paid mutator transaction binding the contract method 0x70e39acd.
//
// Solidity: function isEvidentSubmited(caseId bytes32, serviceId bytes32, clientAddress address) returns(bool)
func (_PromiseValidator *PromiseValidatorTransactor) IsEvidentSubmited(opts *bind.TransactOpts, caseId [32]byte, serviceId [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _PromiseValidator.contract.Transact(opts, "isEvidentSubmited", caseId, serviceId, clientAddress)
}

// IsEvidentSubmited is a paid mutator transaction binding the contract method 0x70e39acd.
//
// Solidity: function isEvidentSubmited(caseId bytes32, serviceId bytes32, clientAddress address) returns(bool)
func (_PromiseValidator *PromiseValidatorSession) IsEvidentSubmited(caseId [32]byte, serviceId [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _PromiseValidator.Contract.IsEvidentSubmited(&_PromiseValidator.TransactOpts, caseId, serviceId, clientAddress)
}

// IsEvidentSubmited is a paid mutator transaction binding the contract method 0x70e39acd.
//
// Solidity: function isEvidentSubmited(caseId bytes32, serviceId bytes32, clientAddress address) returns(bool)
func (_PromiseValidator *PromiseValidatorTransactorSession) IsEvidentSubmited(caseId [32]byte, serviceId [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _PromiseValidator.Contract.IsEvidentSubmited(&_PromiseValidator.TransactOpts, caseId, serviceId, clientAddress)
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

// TrialRulesAbstractABI is the input ABI used to generate the binding from.
const TrialRulesAbstractABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"getInitialStatus\",\"outputs\":[{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReward\",\"outputs\":[{\"name\":\"reward\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"caseId\",\"type\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"startGracePeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"witnessStatus\",\"type\":\"uint8\"},{\"name\":\"trialStatus\",\"type\":\"uint8\"}],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"trialStatus\",\"type\":\"uint8\"}],\"name\":\"getWitness\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"caseId\",\"type\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"expired\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"}]"

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
const WitnessAbstractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"caseId\",\"type\":\"bytes32\"},{\"name\":\"serviceId\",\"type\":\"bytes32\"},{\"name\":\"clientAddress\",\"type\":\"address\"}],\"name\":\"isEvidentSubmited\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"caseId\",\"type\":\"bytes32\"},{\"name\":\"serviceId\",\"type\":\"bytes32\"},{\"name\":\"clientAddress\",\"type\":\"address\"}],\"name\":\"testimonyFor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"}]"

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

// IsEvidentSubmited is a paid mutator transaction binding the contract method 0x70e39acd.
//
// Solidity: function isEvidentSubmited(caseId bytes32, serviceId bytes32, clientAddress address) returns(bool)
func (_WitnessAbstract *WitnessAbstractTransactor) IsEvidentSubmited(opts *bind.TransactOpts, caseId [32]byte, serviceId [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _WitnessAbstract.contract.Transact(opts, "isEvidentSubmited", caseId, serviceId, clientAddress)
}

// IsEvidentSubmited is a paid mutator transaction binding the contract method 0x70e39acd.
//
// Solidity: function isEvidentSubmited(caseId bytes32, serviceId bytes32, clientAddress address) returns(bool)
func (_WitnessAbstract *WitnessAbstractSession) IsEvidentSubmited(caseId [32]byte, serviceId [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _WitnessAbstract.Contract.IsEvidentSubmited(&_WitnessAbstract.TransactOpts, caseId, serviceId, clientAddress)
}

// IsEvidentSubmited is a paid mutator transaction binding the contract method 0x70e39acd.
//
// Solidity: function isEvidentSubmited(caseId bytes32, serviceId bytes32, clientAddress address) returns(bool)
func (_WitnessAbstract *WitnessAbstractTransactorSession) IsEvidentSubmited(caseId [32]byte, serviceId [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _WitnessAbstract.Contract.IsEvidentSubmited(&_WitnessAbstract.TransactOpts, caseId, serviceId, clientAddress)
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
