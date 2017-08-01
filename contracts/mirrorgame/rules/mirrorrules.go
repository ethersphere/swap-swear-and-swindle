// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package mirrorrules

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// MirrorRulesABI is the input ABI used to generate the binding from.
const MirrorRulesABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"getInitialStatus\",\"outputs\":[{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReward\",\"outputs\":[{\"name\":\"reward\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpoch\",\"outputs\":[{\"name\":\"epoch\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"transitions\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"caseId\",\"type\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"startGracePeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"witnesses\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"witnessStatus\",\"type\":\"uint8\"},{\"name\":\"trialStatus\",\"type\":\"uint8\"}],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"trialStatus\",\"type\":\"uint8\"}],\"name\":\"getWitness\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"gracePeriods\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"caseId\",\"type\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"expired\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"paymentValidatorContract\",\"type\":\"address\"},{\"name\":\"ensMirrorValidatorContract\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"}]"

// MirrorRulesBin is the compiled bytecode used for deploying new contracts.
const MirrorRulesBin = `0x6060604052341561000f57600080fd5b60405160408061069c83398101604052808051919060200180519150505b8160008060025b60010160ff1660ff16815260200190815260200160002060006101000a815481600160a060020a030219169083600160a060020a031602179055508060008060028081111561007f57fe5b600290810160ff16825260208201929092526040016000208054600160a060020a031916600160a060020a0393909316929092179091555b60020160016000805b60ff168152602081019190915260400160009081209060025b60010160ff90811682526020820192909252604001600020805460ff19169290911691909117905560025b60016000815b60ff168152602081019190915260400160009081209060025b600190810160ff908116835260208301939093526040909101600020805460ff1916939092169290921790555b60016000805b60ff168152602081019190915260400160009081209060025b600290810160ff908116835260208301939093526040909101600020805460ff1916939092169290921790555b60016000815b60ff168152602081019190915260400160009081209060025b60020160ff1660ff16815260200190815260200160002060006101000a81548160ff021916908360ff1602179055506023600260006002808111156101fc57fe5b60010160ff1660ff1681526020019081526020016000208190555060236002600060028081111561022957fe5b60020160ff1660ff168152602001908152602001600020819055505b50505b610445806102576000396000f300606060405236156100a15763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166307e5ae5f81146100a65780633d18b912146100cf578063757991a8146100f45780639c2a102914610119578063a7edcbe41461014e578063b2400f091461017e578063d39eb301146101c0578063db061563146101f5578063fa778a0f14610237578063fcc0a2ae14610262575b600080fd5b34156100b157600080fd5b6100b9610292565b60405160ff909116815260200160405180910390f35b34156100da57600080fd5b6100e26102a0565b60405190815260200160405180910390f35b34156100ff57600080fd5b6100e26102a6565b60405190815260200160405180910390f35b341561012457600080fd5b6100b960ff600435811690602435166102ac565b60405160ff909116815260200160405180910390f35b341561015957600080fd5b61016a60043560ff602435166102cc565b604051901515815260200160405180910390f35b341561018957600080fd5b61019760ff60043516610320565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34156101cb57600080fd5b6100b960ff60043581169060243516610348565b60405160ff909116815260200160405180910390f35b341561020057600080fd5b61019760ff60043516610371565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561024257600080fd5b6100e260ff600435166103a0565b60405190815260200160405180910390f35b341561026d57600080fd5b61016a60043560ff602435166103b2565b604051901515815260200160405180910390f35b600060025b60010190505b90565b60055b90565b60035b90565b600160209081526000928352604080842090915290825290205460ff1681565b600082815260036020908152604080832060ff8516845290915281205415156103165750600082815260036020908152604080832060ff851684529091529020439055600161031a565b5060005b92915050565b60006020819052908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60ff80831660009081526001602090815260408083208585168452909152902054165b92915050565b60ff811660009081526020819052604090205473ffffffffffffffffffffffffffffffffffffffff165b919050565b60026020526000908152604090205481565b600082815260036020908152604080832060ff85168452909152812054156103165760ff82166000818152600260209081526040808320548784526003835281842094845293909152902054430311156103165750600161031a565b5b5060005b929150505600a165627a7a7230582006692c1098d531bd7065709ebeeb4efc074f41f4f382035f22a4882c4b69c8f80029`

// DeployMirrorRules deploys a new Ethereum contract, binding an instance of MirrorRules to it.
func DeployMirrorRules(auth *bind.TransactOpts, backend bind.ContractBackend, paymentValidatorContract common.Address, ensMirrorValidatorContract common.Address) (common.Address, *types.Transaction, *MirrorRules, error) {
	parsed, err := abi.JSON(strings.NewReader(MirrorRulesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MirrorRulesBin), backend, paymentValidatorContract, ensMirrorValidatorContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MirrorRules{MirrorRulesCaller: MirrorRulesCaller{contract: contract}, MirrorRulesTransactor: MirrorRulesTransactor{contract: contract}}, nil
}

// MirrorRules is an auto generated Go binding around an Ethereum contract.
type MirrorRules struct {
	MirrorRulesCaller     // Read-only binding to the contract
	MirrorRulesTransactor // Write-only binding to the contract
}

// MirrorRulesCaller is an auto generated read-only Go binding around an Ethereum contract.
type MirrorRulesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MirrorRulesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MirrorRulesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MirrorRulesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MirrorRulesSession struct {
	Contract     *MirrorRules      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MirrorRulesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MirrorRulesCallerSession struct {
	Contract *MirrorRulesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MirrorRulesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MirrorRulesTransactorSession struct {
	Contract     *MirrorRulesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MirrorRulesRaw is an auto generated low-level Go binding around an Ethereum contract.
type MirrorRulesRaw struct {
	Contract *MirrorRules // Generic contract binding to access the raw methods on
}

// MirrorRulesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MirrorRulesCallerRaw struct {
	Contract *MirrorRulesCaller // Generic read-only contract binding to access the raw methods on
}

// MirrorRulesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MirrorRulesTransactorRaw struct {
	Contract *MirrorRulesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMirrorRules creates a new instance of MirrorRules, bound to a specific deployed contract.
func NewMirrorRules(address common.Address, backend bind.ContractBackend) (*MirrorRules, error) {
	contract, err := bindMirrorRules(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MirrorRules{MirrorRulesCaller: MirrorRulesCaller{contract: contract}, MirrorRulesTransactor: MirrorRulesTransactor{contract: contract}}, nil
}

// NewMirrorRulesCaller creates a new read-only instance of MirrorRules, bound to a specific deployed contract.
func NewMirrorRulesCaller(address common.Address, caller bind.ContractCaller) (*MirrorRulesCaller, error) {
	contract, err := bindMirrorRules(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MirrorRulesCaller{contract: contract}, nil
}

// NewMirrorRulesTransactor creates a new write-only instance of MirrorRules, bound to a specific deployed contract.
func NewMirrorRulesTransactor(address common.Address, transactor bind.ContractTransactor) (*MirrorRulesTransactor, error) {
	contract, err := bindMirrorRules(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MirrorRulesTransactor{contract: contract}, nil
}

// bindMirrorRules binds a generic wrapper to an already deployed contract.
func bindMirrorRules(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MirrorRulesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MirrorRules *MirrorRulesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MirrorRules.Contract.MirrorRulesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MirrorRules *MirrorRulesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MirrorRules.Contract.MirrorRulesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MirrorRules *MirrorRulesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MirrorRules.Contract.MirrorRulesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MirrorRules *MirrorRulesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MirrorRules.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MirrorRules *MirrorRulesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MirrorRules.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MirrorRules *MirrorRulesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MirrorRules.Contract.contract.Transact(opts, method, params...)
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(epoch uint256)
func (_MirrorRules *MirrorRulesCaller) GetEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MirrorRules.contract.Call(opts, out, "getEpoch")
	return *ret0, err
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(epoch uint256)
func (_MirrorRules *MirrorRulesSession) GetEpoch() (*big.Int, error) {
	return _MirrorRules.Contract.GetEpoch(&_MirrorRules.CallOpts)
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(epoch uint256)
func (_MirrorRules *MirrorRulesCallerSession) GetEpoch() (*big.Int, error) {
	return _MirrorRules.Contract.GetEpoch(&_MirrorRules.CallOpts)
}

// GetReward is a free data retrieval call binding the contract method 0x3d18b912.
//
// Solidity: function getReward() constant returns(reward uint256)
func (_MirrorRules *MirrorRulesCaller) GetReward(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MirrorRules.contract.Call(opts, out, "getReward")
	return *ret0, err
}

// GetReward is a free data retrieval call binding the contract method 0x3d18b912.
//
// Solidity: function getReward() constant returns(reward uint256)
func (_MirrorRules *MirrorRulesSession) GetReward() (*big.Int, error) {
	return _MirrorRules.Contract.GetReward(&_MirrorRules.CallOpts)
}

// GetReward is a free data retrieval call binding the contract method 0x3d18b912.
//
// Solidity: function getReward() constant returns(reward uint256)
func (_MirrorRules *MirrorRulesCallerSession) GetReward() (*big.Int, error) {
	return _MirrorRules.Contract.GetReward(&_MirrorRules.CallOpts)
}

// GracePeriods is a free data retrieval call binding the contract method 0xfa778a0f.
//
// Solidity: function gracePeriods( uint8) constant returns(uint256)
func (_MirrorRules *MirrorRulesCaller) GracePeriods(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MirrorRules.contract.Call(opts, out, "gracePeriods", arg0)
	return *ret0, err
}

// GracePeriods is a free data retrieval call binding the contract method 0xfa778a0f.
//
// Solidity: function gracePeriods( uint8) constant returns(uint256)
func (_MirrorRules *MirrorRulesSession) GracePeriods(arg0 uint8) (*big.Int, error) {
	return _MirrorRules.Contract.GracePeriods(&_MirrorRules.CallOpts, arg0)
}

// GracePeriods is a free data retrieval call binding the contract method 0xfa778a0f.
//
// Solidity: function gracePeriods( uint8) constant returns(uint256)
func (_MirrorRules *MirrorRulesCallerSession) GracePeriods(arg0 uint8) (*big.Int, error) {
	return _MirrorRules.Contract.GracePeriods(&_MirrorRules.CallOpts, arg0)
}

// Transitions is a free data retrieval call binding the contract method 0x9c2a1029.
//
// Solidity: function transitions( uint8,  uint8) constant returns(uint8)
func (_MirrorRules *MirrorRulesCaller) Transitions(opts *bind.CallOpts, arg0 uint8, arg1 uint8) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _MirrorRules.contract.Call(opts, out, "transitions", arg0, arg1)
	return *ret0, err
}

// Transitions is a free data retrieval call binding the contract method 0x9c2a1029.
//
// Solidity: function transitions( uint8,  uint8) constant returns(uint8)
func (_MirrorRules *MirrorRulesSession) Transitions(arg0 uint8, arg1 uint8) (uint8, error) {
	return _MirrorRules.Contract.Transitions(&_MirrorRules.CallOpts, arg0, arg1)
}

// Transitions is a free data retrieval call binding the contract method 0x9c2a1029.
//
// Solidity: function transitions( uint8,  uint8) constant returns(uint8)
func (_MirrorRules *MirrorRulesCallerSession) Transitions(arg0 uint8, arg1 uint8) (uint8, error) {
	return _MirrorRules.Contract.Transitions(&_MirrorRules.CallOpts, arg0, arg1)
}

// Witnesses is a free data retrieval call binding the contract method 0xb2400f09.
//
// Solidity: function witnesses( uint8) constant returns(address)
func (_MirrorRules *MirrorRulesCaller) Witnesses(opts *bind.CallOpts, arg0 uint8) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MirrorRules.contract.Call(opts, out, "witnesses", arg0)
	return *ret0, err
}

// Witnesses is a free data retrieval call binding the contract method 0xb2400f09.
//
// Solidity: function witnesses( uint8) constant returns(address)
func (_MirrorRules *MirrorRulesSession) Witnesses(arg0 uint8) (common.Address, error) {
	return _MirrorRules.Contract.Witnesses(&_MirrorRules.CallOpts, arg0)
}

// Witnesses is a free data retrieval call binding the contract method 0xb2400f09.
//
// Solidity: function witnesses( uint8) constant returns(address)
func (_MirrorRules *MirrorRulesCallerSession) Witnesses(arg0 uint8) (common.Address, error) {
	return _MirrorRules.Contract.Witnesses(&_MirrorRules.CallOpts, arg0)
}

// Expired is a paid mutator transaction binding the contract method 0xfcc0a2ae.
//
// Solidity: function expired(caseId bytes32, status uint8) returns(bool)
func (_MirrorRules *MirrorRulesTransactor) Expired(opts *bind.TransactOpts, caseId [32]byte, status uint8) (*types.Transaction, error) {
	return _MirrorRules.contract.Transact(opts, "expired", caseId, status)
}

// Expired is a paid mutator transaction binding the contract method 0xfcc0a2ae.
//
// Solidity: function expired(caseId bytes32, status uint8) returns(bool)
func (_MirrorRules *MirrorRulesSession) Expired(caseId [32]byte, status uint8) (*types.Transaction, error) {
	return _MirrorRules.Contract.Expired(&_MirrorRules.TransactOpts, caseId, status)
}

// Expired is a paid mutator transaction binding the contract method 0xfcc0a2ae.
//
// Solidity: function expired(caseId bytes32, status uint8) returns(bool)
func (_MirrorRules *MirrorRulesTransactorSession) Expired(caseId [32]byte, status uint8) (*types.Transaction, error) {
	return _MirrorRules.Contract.Expired(&_MirrorRules.TransactOpts, caseId, status)
}

// GetInitialStatus is a paid mutator transaction binding the contract method 0x07e5ae5f.
//
// Solidity: function getInitialStatus() returns(status uint8)
func (_MirrorRules *MirrorRulesTransactor) GetInitialStatus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MirrorRules.contract.Transact(opts, "getInitialStatus")
}

// GetInitialStatus is a paid mutator transaction binding the contract method 0x07e5ae5f.
//
// Solidity: function getInitialStatus() returns(status uint8)
func (_MirrorRules *MirrorRulesSession) GetInitialStatus() (*types.Transaction, error) {
	return _MirrorRules.Contract.GetInitialStatus(&_MirrorRules.TransactOpts)
}

// GetInitialStatus is a paid mutator transaction binding the contract method 0x07e5ae5f.
//
// Solidity: function getInitialStatus() returns(status uint8)
func (_MirrorRules *MirrorRulesTransactorSession) GetInitialStatus() (*types.Transaction, error) {
	return _MirrorRules.Contract.GetInitialStatus(&_MirrorRules.TransactOpts)
}

// GetStatus is a paid mutator transaction binding the contract method 0xd39eb301.
//
// Solidity: function getStatus(witnessStatus uint8, trialStatus uint8) returns(status uint8)
func (_MirrorRules *MirrorRulesTransactor) GetStatus(opts *bind.TransactOpts, witnessStatus uint8, trialStatus uint8) (*types.Transaction, error) {
	return _MirrorRules.contract.Transact(opts, "getStatus", witnessStatus, trialStatus)
}

// GetStatus is a paid mutator transaction binding the contract method 0xd39eb301.
//
// Solidity: function getStatus(witnessStatus uint8, trialStatus uint8) returns(status uint8)
func (_MirrorRules *MirrorRulesSession) GetStatus(witnessStatus uint8, trialStatus uint8) (*types.Transaction, error) {
	return _MirrorRules.Contract.GetStatus(&_MirrorRules.TransactOpts, witnessStatus, trialStatus)
}

// GetStatus is a paid mutator transaction binding the contract method 0xd39eb301.
//
// Solidity: function getStatus(witnessStatus uint8, trialStatus uint8) returns(status uint8)
func (_MirrorRules *MirrorRulesTransactorSession) GetStatus(witnessStatus uint8, trialStatus uint8) (*types.Transaction, error) {
	return _MirrorRules.Contract.GetStatus(&_MirrorRules.TransactOpts, witnessStatus, trialStatus)
}

// GetWitness is a paid mutator transaction binding the contract method 0xdb061563.
//
// Solidity: function getWitness(trialStatus uint8) returns(address)
func (_MirrorRules *MirrorRulesTransactor) GetWitness(opts *bind.TransactOpts, trialStatus uint8) (*types.Transaction, error) {
	return _MirrorRules.contract.Transact(opts, "getWitness", trialStatus)
}

// GetWitness is a paid mutator transaction binding the contract method 0xdb061563.
//
// Solidity: function getWitness(trialStatus uint8) returns(address)
func (_MirrorRules *MirrorRulesSession) GetWitness(trialStatus uint8) (*types.Transaction, error) {
	return _MirrorRules.Contract.GetWitness(&_MirrorRules.TransactOpts, trialStatus)
}

// GetWitness is a paid mutator transaction binding the contract method 0xdb061563.
//
// Solidity: function getWitness(trialStatus uint8) returns(address)
func (_MirrorRules *MirrorRulesTransactorSession) GetWitness(trialStatus uint8) (*types.Transaction, error) {
	return _MirrorRules.Contract.GetWitness(&_MirrorRules.TransactOpts, trialStatus)
}

// StartGracePeriod is a paid mutator transaction binding the contract method 0xa7edcbe4.
//
// Solidity: function startGracePeriod(caseId bytes32, status uint8) returns(bool)
func (_MirrorRules *MirrorRulesTransactor) StartGracePeriod(opts *bind.TransactOpts, caseId [32]byte, status uint8) (*types.Transaction, error) {
	return _MirrorRules.contract.Transact(opts, "startGracePeriod", caseId, status)
}

// StartGracePeriod is a paid mutator transaction binding the contract method 0xa7edcbe4.
//
// Solidity: function startGracePeriod(caseId bytes32, status uint8) returns(bool)
func (_MirrorRules *MirrorRulesSession) StartGracePeriod(caseId [32]byte, status uint8) (*types.Transaction, error) {
	return _MirrorRules.Contract.StartGracePeriod(&_MirrorRules.TransactOpts, caseId, status)
}

// StartGracePeriod is a paid mutator transaction binding the contract method 0xa7edcbe4.
//
// Solidity: function startGracePeriod(caseId bytes32, status uint8) returns(bool)
func (_MirrorRules *MirrorRulesTransactorSession) StartGracePeriod(caseId [32]byte, status uint8) (*types.Transaction, error) {
	return _MirrorRules.Contract.StartGracePeriod(&_MirrorRules.TransactOpts, caseId, status)
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
