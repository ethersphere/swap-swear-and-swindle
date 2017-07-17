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
const CaseContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"},{\"name\":\"_evidence\",\"type\":\"bytes32\"}],\"name\":\"submitEvidence\",\"outputs\":[{\"name\":\"status\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"status\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"resolveClaim\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getClaim\",\"outputs\":[{\"name\":\"plaintiff\",\"type\":\"address\"},{\"name\":\"beneficiary\",\"type\":\"address\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"name\":\"sig_v\",\"type\":\"uint8\"},{\"name\":\"sig_r\",\"type\":\"bytes32\"},{\"name\":\"sig_s\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_plaintiff\",\"type\":\"address\"},{\"name\":\"_evidence\",\"type\":\"bytes32\"}],\"name\":\"newClaim\",\"outputs\":[{\"name\":\"id\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"},{\"name\":\"beneficiary\",\"type\":\"address\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"name\":\"sig_v\",\"type\":\"uint8\"},{\"name\":\"sig_r\",\"type\":\"bytes32\"},{\"name\":\"sig_s\",\"type\":\"bytes32\"}],\"name\":\"submitPromise\",\"outputs\":[{\"name\":\"status\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"}]"

// CaseContractBin is the compiled bytecode used for deploying new contracts.
const CaseContractBin = `0x6060604052341561000f57600080fd5b5b5b60008054600160a060020a03191633600160a060020a03161790555b5b5b6108468061003e6000396000f3006060604052361561008b5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166317dfa44c81146100905780635de28ae0146100bb5780638da5cb5b146100e3578063a6f9dae114610112578063c1203cb814610133578063c9100bcb1461014b578063e391c4c0146101a5578063fb9e6e77146101d9575b600080fd5b341561009b57600080fd5b6100a960043560243561021c565b60405190815260200160405180910390f35b34156100c657600080fd5b6100a9600435610296565b60405190815260200160405180910390f35b34156100ee57600080fd5b6100f66102ae565b604051600160a060020a03909116815260200160405180910390f35b341561011d57600080fd5b610131600160a060020a03600435166102bd565b005b341561013e57600080fd5b610131600435610305565b005b341561015657600080fd5b6101616004356103c1565b604051600160a060020a03968716815294909516602085015260408085019390935260ff9091166060840152608083015260a082019290925260c001905180910390f35b34156101b057600080fd5b6100a9600160a060020a0360043516602435610413565b60405190815260200160405180910390f35b34156101e457600080fd5b6100a9600435600160a060020a036024351660443560ff6064351660843560a4356105dc565b60405190815260200160405180910390f35b6000828152600160205260408120600201546004811061023b57600080fd5b600084815260016020526040902060030180548491908390811061025b57fe5b906000526020600020900160005b505560008481526001602081905260409091206002810180549092019091556009015491505b5092915050565b6000818152600160205260409020600901545b919050565b600054600160a060020a031681565b60005433600160a060020a039081169116146102d857600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b50565b600081815260016020526040902054158061032f5750600081815260016020526040902060090154155b1561033957600080fd5b6000818152600160208190526040808320838155909101805473ffffffffffffffffffffffffffffffffffffffff19169055518059106103765750595b908082528060200260200182016040525b5060008281526001602052604090206003019080516103aa9291602001906106ab565b506000818152600160205260408120600901555b50565b6000818152600160208190526040909120908101546004820154600583015460068401546007850154600890950154600160a060020a03948516959390941693919260ff909116915b91939550919395565b600061041d6106f9565b838342604051600160a060020a03939093166c01000000000000000000000000028352601483019190915260348201526054016040519081900390208152600160008251815260208101919091526040016000206009015415610483576000915061028f565b600160a060020a038416602082015260046040518059106104a15750595b908082528060200260200182016040525b50606082019081528390516000815181106104c957fe5b6020908102909101015260016040820181905260a082018190528190600082518152602081019190915260400160002081518155602082015160018201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055604082015181600201556060820151816003019080516105559291602001906106ab565b506080820151600482018151815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039190911617815560208201518160010155604082015160028201805460ff191660ff929092169190911790556060820151600382015560808201516004909101555060a082015160099091015550805191505b5092915050565b60006105e661078b565b600160a060020a0387168152602080820187905260ff8616604080840191909152606083018690526080830185905260008a81526001909252902081906004018151815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039190911617815560208201518160010155604082015160028201805460ff191660ff929092169190911790556060820151600382015560808201516004909101555060008881526001602052604090206009015491505b509695505050505050565b8280548282559060005260206000209081019282156106e8579160200282015b828111156106e857825182556020909201916001909101906106cb565b5b506106f59291506107b9565b5090565b610140604051908101604090815260008083526020830181905290820152606081016107236107da565b815260200161073061078b565b8152602001600081525090565b8280548282559060005260206000209081019282156106e8579160200282015b828111156106e857825182556020909201916001909101906106cb565b5b506106f59291506107b9565b5090565b60a0604051908101604090815260008083526020830181905290820181905260608201819052608082015290565b6107d791905b808211156106f557600081556001016107bf565b5090565b90565b60206040519081016040526000815290565b60a06040519081016040908152600080835260208301819052908201819052606082018190526080820152905600a165627a7a72305820481198ad1d6f04e7db79cfed7a194abcae90768fc4c85fd71c266b85a110055c0029`

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
// Solidity: function getStatus(id bytes32) constant returns(status uint256)
func (_CaseContract *CaseContractCaller) GetStatus(opts *bind.CallOpts, id [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CaseContract.contract.Call(opts, out, "getStatus", id)
	return *ret0, err
}

// GetStatus is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(id bytes32) constant returns(status uint256)
func (_CaseContract *CaseContractSession) GetStatus(id [32]byte) (*big.Int, error) {
	return _CaseContract.Contract.GetStatus(&_CaseContract.CallOpts, id)
}

// GetStatus is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(id bytes32) constant returns(status uint256)
func (_CaseContract *CaseContractCallerSession) GetStatus(id [32]byte) (*big.Int, error) {
	return _CaseContract.Contract.GetStatus(&_CaseContract.CallOpts, id)
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
// Solidity: function getClaim(_id bytes32) returns(plaintiff address, beneficiary address, blockNumber uint256, sig_v uint8, sig_r bytes32, sig_s bytes32)
func (_CaseContract *CaseContractTransactor) GetClaim(opts *bind.TransactOpts, _id [32]byte) (*types.Transaction, error) {
	return _CaseContract.contract.Transact(opts, "getClaim", _id)
}

// GetClaim is a paid mutator transaction binding the contract method 0xc9100bcb.
//
// Solidity: function getClaim(_id bytes32) returns(plaintiff address, beneficiary address, blockNumber uint256, sig_v uint8, sig_r bytes32, sig_s bytes32)
func (_CaseContract *CaseContractSession) GetClaim(_id [32]byte) (*types.Transaction, error) {
	return _CaseContract.Contract.GetClaim(&_CaseContract.TransactOpts, _id)
}

// GetClaim is a paid mutator transaction binding the contract method 0xc9100bcb.
//
// Solidity: function getClaim(_id bytes32) returns(plaintiff address, beneficiary address, blockNumber uint256, sig_v uint8, sig_r bytes32, sig_s bytes32)
func (_CaseContract *CaseContractTransactorSession) GetClaim(_id [32]byte) (*types.Transaction, error) {
	return _CaseContract.Contract.GetClaim(&_CaseContract.TransactOpts, _id)
}

// NewClaim is a paid mutator transaction binding the contract method 0xe391c4c0.
//
// Solidity: function newClaim(_plaintiff address, _evidence bytes32) returns(id bytes32)
func (_CaseContract *CaseContractTransactor) NewClaim(opts *bind.TransactOpts, _plaintiff common.Address, _evidence [32]byte) (*types.Transaction, error) {
	return _CaseContract.contract.Transact(opts, "newClaim", _plaintiff, _evidence)
}

// NewClaim is a paid mutator transaction binding the contract method 0xe391c4c0.
//
// Solidity: function newClaim(_plaintiff address, _evidence bytes32) returns(id bytes32)
func (_CaseContract *CaseContractSession) NewClaim(_plaintiff common.Address, _evidence [32]byte) (*types.Transaction, error) {
	return _CaseContract.Contract.NewClaim(&_CaseContract.TransactOpts, _plaintiff, _evidence)
}

// NewClaim is a paid mutator transaction binding the contract method 0xe391c4c0.
//
// Solidity: function newClaim(_plaintiff address, _evidence bytes32) returns(id bytes32)
func (_CaseContract *CaseContractTransactorSession) NewClaim(_plaintiff common.Address, _evidence [32]byte) (*types.Transaction, error) {
	return _CaseContract.Contract.NewClaim(&_CaseContract.TransactOpts, _plaintiff, _evidence)
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

// SubmitEvidence is a paid mutator transaction binding the contract method 0x17dfa44c.
//
// Solidity: function submitEvidence(_id bytes32, _evidence bytes32) returns(status uint256)
func (_CaseContract *CaseContractTransactor) SubmitEvidence(opts *bind.TransactOpts, _id [32]byte, _evidence [32]byte) (*types.Transaction, error) {
	return _CaseContract.contract.Transact(opts, "submitEvidence", _id, _evidence)
}

// SubmitEvidence is a paid mutator transaction binding the contract method 0x17dfa44c.
//
// Solidity: function submitEvidence(_id bytes32, _evidence bytes32) returns(status uint256)
func (_CaseContract *CaseContractSession) SubmitEvidence(_id [32]byte, _evidence [32]byte) (*types.Transaction, error) {
	return _CaseContract.Contract.SubmitEvidence(&_CaseContract.TransactOpts, _id, _evidence)
}

// SubmitEvidence is a paid mutator transaction binding the contract method 0x17dfa44c.
//
// Solidity: function submitEvidence(_id bytes32, _evidence bytes32) returns(status uint256)
func (_CaseContract *CaseContractTransactorSession) SubmitEvidence(_id [32]byte, _evidence [32]byte) (*types.Transaction, error) {
	return _CaseContract.Contract.SubmitEvidence(&_CaseContract.TransactOpts, _id, _evidence)
}

// SubmitPromise is a paid mutator transaction binding the contract method 0xfb9e6e77.
//
// Solidity: function submitPromise(_id bytes32, beneficiary address, blockNumber uint256, sig_v uint8, sig_r bytes32, sig_s bytes32) returns(status uint256)
func (_CaseContract *CaseContractTransactor) SubmitPromise(opts *bind.TransactOpts, _id [32]byte, beneficiary common.Address, blockNumber *big.Int, sig_v uint8, sig_r [32]byte, sig_s [32]byte) (*types.Transaction, error) {
	return _CaseContract.contract.Transact(opts, "submitPromise", _id, beneficiary, blockNumber, sig_v, sig_r, sig_s)
}

// SubmitPromise is a paid mutator transaction binding the contract method 0xfb9e6e77.
//
// Solidity: function submitPromise(_id bytes32, beneficiary address, blockNumber uint256, sig_v uint8, sig_r bytes32, sig_s bytes32) returns(status uint256)
func (_CaseContract *CaseContractSession) SubmitPromise(_id [32]byte, beneficiary common.Address, blockNumber *big.Int, sig_v uint8, sig_r [32]byte, sig_s [32]byte) (*types.Transaction, error) {
	return _CaseContract.Contract.SubmitPromise(&_CaseContract.TransactOpts, _id, beneficiary, blockNumber, sig_v, sig_r, sig_s)
}

// SubmitPromise is a paid mutator transaction binding the contract method 0xfb9e6e77.
//
// Solidity: function submitPromise(_id bytes32, beneficiary address, blockNumber uint256, sig_v uint8, sig_r bytes32, sig_s bytes32) returns(status uint256)
func (_CaseContract *CaseContractTransactorSession) SubmitPromise(_id [32]byte, beneficiary common.Address, blockNumber *big.Int, sig_v uint8, sig_r [32]byte, sig_s [32]byte) (*types.Transaction, error) {
	return _CaseContract.Contract.SubmitPromise(&_CaseContract.TransactOpts, _id, beneficiary, blockNumber, sig_v, sig_r, sig_s)
}

// ENSAbstractABI is the input ABI used to generate the binding from.
const ENSAbstractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"resolver\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"label\",\"type\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setSubnodeOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"setTTL\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"ttl\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"setResolver\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"}]"

// ENSAbstractBin is the compiled bytecode used for deploying new contracts.
const ENSAbstractBin = `0x`

// DeployENSAbstract deploys a new Ethereum contract, binding an instance of ENSAbstract to it.
func DeployENSAbstract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ENSAbstract, error) {
	parsed, err := abi.JSON(strings.NewReader(ENSAbstractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ENSAbstractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ENSAbstract{ENSAbstractCaller: ENSAbstractCaller{contract: contract}, ENSAbstractTransactor: ENSAbstractTransactor{contract: contract}}, nil
}

// ENSAbstract is an auto generated Go binding around an Ethereum contract.
type ENSAbstract struct {
	ENSAbstractCaller     // Read-only binding to the contract
	ENSAbstractTransactor // Write-only binding to the contract
}

// ENSAbstractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ENSAbstractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENSAbstractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ENSAbstractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENSAbstractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ENSAbstractSession struct {
	Contract     *ENSAbstract      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ENSAbstractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ENSAbstractCallerSession struct {
	Contract *ENSAbstractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ENSAbstractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ENSAbstractTransactorSession struct {
	Contract     *ENSAbstractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ENSAbstractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ENSAbstractRaw struct {
	Contract *ENSAbstract // Generic contract binding to access the raw methods on
}

// ENSAbstractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ENSAbstractCallerRaw struct {
	Contract *ENSAbstractCaller // Generic read-only contract binding to access the raw methods on
}

// ENSAbstractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ENSAbstractTransactorRaw struct {
	Contract *ENSAbstractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewENSAbstract creates a new instance of ENSAbstract, bound to a specific deployed contract.
func NewENSAbstract(address common.Address, backend bind.ContractBackend) (*ENSAbstract, error) {
	contract, err := bindENSAbstract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ENSAbstract{ENSAbstractCaller: ENSAbstractCaller{contract: contract}, ENSAbstractTransactor: ENSAbstractTransactor{contract: contract}}, nil
}

// NewENSAbstractCaller creates a new read-only instance of ENSAbstract, bound to a specific deployed contract.
func NewENSAbstractCaller(address common.Address, caller bind.ContractCaller) (*ENSAbstractCaller, error) {
	contract, err := bindENSAbstract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ENSAbstractCaller{contract: contract}, nil
}

// NewENSAbstractTransactor creates a new write-only instance of ENSAbstract, bound to a specific deployed contract.
func NewENSAbstractTransactor(address common.Address, transactor bind.ContractTransactor) (*ENSAbstractTransactor, error) {
	contract, err := bindENSAbstract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ENSAbstractTransactor{contract: contract}, nil
}

// bindENSAbstract binds a generic wrapper to an already deployed contract.
func bindENSAbstract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ENSAbstractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ENSAbstract *ENSAbstractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ENSAbstract.Contract.ENSAbstractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ENSAbstract *ENSAbstractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ENSAbstract.Contract.ENSAbstractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ENSAbstract *ENSAbstractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ENSAbstract.Contract.ENSAbstractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ENSAbstract *ENSAbstractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ENSAbstract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ENSAbstract *ENSAbstractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ENSAbstract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ENSAbstract *ENSAbstractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ENSAbstract.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(node bytes32) constant returns(address)
func (_ENSAbstract *ENSAbstractCaller) Owner(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ENSAbstract.contract.Call(opts, out, "owner", node)
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(node bytes32) constant returns(address)
func (_ENSAbstract *ENSAbstractSession) Owner(node [32]byte) (common.Address, error) {
	return _ENSAbstract.Contract.Owner(&_ENSAbstract.CallOpts, node)
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(node bytes32) constant returns(address)
func (_ENSAbstract *ENSAbstractCallerSession) Owner(node [32]byte) (common.Address, error) {
	return _ENSAbstract.Contract.Owner(&_ENSAbstract.CallOpts, node)
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(node bytes32) constant returns(address)
func (_ENSAbstract *ENSAbstractCaller) Resolver(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ENSAbstract.contract.Call(opts, out, "resolver", node)
	return *ret0, err
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(node bytes32) constant returns(address)
func (_ENSAbstract *ENSAbstractSession) Resolver(node [32]byte) (common.Address, error) {
	return _ENSAbstract.Contract.Resolver(&_ENSAbstract.CallOpts, node)
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(node bytes32) constant returns(address)
func (_ENSAbstract *ENSAbstractCallerSession) Resolver(node [32]byte) (common.Address, error) {
	return _ENSAbstract.Contract.Resolver(&_ENSAbstract.CallOpts, node)
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(node bytes32) constant returns(uint64)
func (_ENSAbstract *ENSAbstractCaller) Ttl(opts *bind.CallOpts, node [32]byte) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _ENSAbstract.contract.Call(opts, out, "ttl", node)
	return *ret0, err
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(node bytes32) constant returns(uint64)
func (_ENSAbstract *ENSAbstractSession) Ttl(node [32]byte) (uint64, error) {
	return _ENSAbstract.Contract.Ttl(&_ENSAbstract.CallOpts, node)
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(node bytes32) constant returns(uint64)
func (_ENSAbstract *ENSAbstractCallerSession) Ttl(node [32]byte) (uint64, error) {
	return _ENSAbstract.Contract.Ttl(&_ENSAbstract.CallOpts, node)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(node bytes32, owner address) returns()
func (_ENSAbstract *ENSAbstractTransactor) SetOwner(opts *bind.TransactOpts, node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENSAbstract.contract.Transact(opts, "setOwner", node, owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(node bytes32, owner address) returns()
func (_ENSAbstract *ENSAbstractSession) SetOwner(node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENSAbstract.Contract.SetOwner(&_ENSAbstract.TransactOpts, node, owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(node bytes32, owner address) returns()
func (_ENSAbstract *ENSAbstractTransactorSession) SetOwner(node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENSAbstract.Contract.SetOwner(&_ENSAbstract.TransactOpts, node, owner)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(node bytes32, resolver address) returns()
func (_ENSAbstract *ENSAbstractTransactor) SetResolver(opts *bind.TransactOpts, node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _ENSAbstract.contract.Transact(opts, "setResolver", node, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(node bytes32, resolver address) returns()
func (_ENSAbstract *ENSAbstractSession) SetResolver(node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _ENSAbstract.Contract.SetResolver(&_ENSAbstract.TransactOpts, node, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(node bytes32, resolver address) returns()
func (_ENSAbstract *ENSAbstractTransactorSession) SetResolver(node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _ENSAbstract.Contract.SetResolver(&_ENSAbstract.TransactOpts, node, resolver)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(node bytes32, label bytes32, owner address) returns()
func (_ENSAbstract *ENSAbstractTransactor) SetSubnodeOwner(opts *bind.TransactOpts, node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENSAbstract.contract.Transact(opts, "setSubnodeOwner", node, label, owner)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(node bytes32, label bytes32, owner address) returns()
func (_ENSAbstract *ENSAbstractSession) SetSubnodeOwner(node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENSAbstract.Contract.SetSubnodeOwner(&_ENSAbstract.TransactOpts, node, label, owner)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(node bytes32, label bytes32, owner address) returns()
func (_ENSAbstract *ENSAbstractTransactorSession) SetSubnodeOwner(node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENSAbstract.Contract.SetSubnodeOwner(&_ENSAbstract.TransactOpts, node, label, owner)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(node bytes32, ttl uint64) returns()
func (_ENSAbstract *ENSAbstractTransactor) SetTTL(opts *bind.TransactOpts, node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _ENSAbstract.contract.Transact(opts, "setTTL", node, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(node bytes32, ttl uint64) returns()
func (_ENSAbstract *ENSAbstractSession) SetTTL(node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _ENSAbstract.Contract.SetTTL(&_ENSAbstract.TransactOpts, node, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(node bytes32, ttl uint64) returns()
func (_ENSAbstract *ENSAbstractTransactorSession) SetTTL(node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _ENSAbstract.Contract.SetTTL(&_ENSAbstract.TransactOpts, node, ttl)
}

// OwnedABI is the input ABI used to generate the binding from.
const OwnedABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"}]"

// OwnedBin is the compiled bytecode used for deploying new contracts.
const OwnedBin = `0x6060604052341561000f57600080fd5b5b60008054600160a060020a03191633600160a060020a03161790555b5b61015c8061003c6000396000f300606060405263ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b8114610048578063a6f9dae114610084575b600080fd5b341561005357600080fd5b61005b6100b2565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561008f57600080fd5b6100b073ffffffffffffffffffffffffffffffffffffffff600435166100ce565b005b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6000543373ffffffffffffffffffffffffffffffffffffffff9081169116146100f657600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff83161790555b5b505600a165627a7a723058200c617cfff026d930f454a0614fa2b8e85871dcb768dc277b2284ec659c67f7540029`

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

// ResolverAbstractABI is the input ABI used to generate the binding from.
const ResolverAbstractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"content\",\"outputs\":[{\"name\":\"content\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"}]"

// ResolverAbstractBin is the compiled bytecode used for deploying new contracts.
const ResolverAbstractBin = `0x`

// DeployResolverAbstract deploys a new Ethereum contract, binding an instance of ResolverAbstract to it.
func DeployResolverAbstract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ResolverAbstract, error) {
	parsed, err := abi.JSON(strings.NewReader(ResolverAbstractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ResolverAbstractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ResolverAbstract{ResolverAbstractCaller: ResolverAbstractCaller{contract: contract}, ResolverAbstractTransactor: ResolverAbstractTransactor{contract: contract}}, nil
}

// ResolverAbstract is an auto generated Go binding around an Ethereum contract.
type ResolverAbstract struct {
	ResolverAbstractCaller     // Read-only binding to the contract
	ResolverAbstractTransactor // Write-only binding to the contract
}

// ResolverAbstractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ResolverAbstractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverAbstractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ResolverAbstractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverAbstractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ResolverAbstractSession struct {
	Contract     *ResolverAbstract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ResolverAbstractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ResolverAbstractCallerSession struct {
	Contract *ResolverAbstractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ResolverAbstractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ResolverAbstractTransactorSession struct {
	Contract     *ResolverAbstractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ResolverAbstractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ResolverAbstractRaw struct {
	Contract *ResolverAbstract // Generic contract binding to access the raw methods on
}

// ResolverAbstractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ResolverAbstractCallerRaw struct {
	Contract *ResolverAbstractCaller // Generic read-only contract binding to access the raw methods on
}

// ResolverAbstractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ResolverAbstractTransactorRaw struct {
	Contract *ResolverAbstractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewResolverAbstract creates a new instance of ResolverAbstract, bound to a specific deployed contract.
func NewResolverAbstract(address common.Address, backend bind.ContractBackend) (*ResolverAbstract, error) {
	contract, err := bindResolverAbstract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ResolverAbstract{ResolverAbstractCaller: ResolverAbstractCaller{contract: contract}, ResolverAbstractTransactor: ResolverAbstractTransactor{contract: contract}}, nil
}

// NewResolverAbstractCaller creates a new read-only instance of ResolverAbstract, bound to a specific deployed contract.
func NewResolverAbstractCaller(address common.Address, caller bind.ContractCaller) (*ResolverAbstractCaller, error) {
	contract, err := bindResolverAbstract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ResolverAbstractCaller{contract: contract}, nil
}

// NewResolverAbstractTransactor creates a new write-only instance of ResolverAbstract, bound to a specific deployed contract.
func NewResolverAbstractTransactor(address common.Address, transactor bind.ContractTransactor) (*ResolverAbstractTransactor, error) {
	contract, err := bindResolverAbstract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ResolverAbstractTransactor{contract: contract}, nil
}

// bindResolverAbstract binds a generic wrapper to an already deployed contract.
func bindResolverAbstract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ResolverAbstractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ResolverAbstract *ResolverAbstractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ResolverAbstract.Contract.ResolverAbstractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ResolverAbstract *ResolverAbstractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ResolverAbstract.Contract.ResolverAbstractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ResolverAbstract *ResolverAbstractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ResolverAbstract.Contract.ResolverAbstractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ResolverAbstract *ResolverAbstractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ResolverAbstract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ResolverAbstract *ResolverAbstractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ResolverAbstract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ResolverAbstract *ResolverAbstractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ResolverAbstract.Contract.contract.Transact(opts, method, params...)
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(node bytes32) constant returns(content bytes32)
func (_ResolverAbstract *ResolverAbstractCaller) Content(opts *bind.CallOpts, node [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ResolverAbstract.contract.Call(opts, out, "content", node)
	return *ret0, err
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(node bytes32) constant returns(content bytes32)
func (_ResolverAbstract *ResolverAbstractSession) Content(node [32]byte) ([32]byte, error) {
	return _ResolverAbstract.Contract.Content(&_ResolverAbstract.CallOpts, node)
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(node bytes32) constant returns(content bytes32)
func (_ResolverAbstract *ResolverAbstractCallerSession) Content(node [32]byte) ([32]byte, error) {
	return _ResolverAbstract.Contract.Content(&_ResolverAbstract.CallOpts, node)
}

// SampleTokenABI is the input ABI used to generate the binding from.
const SampleTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"beneficiary\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"ref\",\"type\":\"bytes32\"}],\"name\":\"createTokens\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"ref\",\"type\":\"bytes32\"}],\"name\":\"TokenMined\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// SampleTokenBin is the compiled bytecode used for deploying new contracts.
const SampleTokenBin = `0x6060604052341561000f57600080fd5b60405160208061069d833981016040528080519150505b5b60038054600160a060020a03191633600160a060020a03161790555b600160a060020a03331660009081526001602052604081208290558190555b505b61062a806100736000396000f300606060405236156100965763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461009b57806318160ddd146100d157806323b872dd146100f657806370a08231146101325780638da5cb5b14610163578063a3aff92314610192578063a6f9dae1146101b9578063a9059cbb146101da578063dd62ed3e14610210575b600080fd5b34156100a657600080fd5b6100bd600160a060020a0360043516602435610247565b604051901515815260200160405180910390f35b34156100dc57600080fd5b6100e46102b4565b60405190815260200160405180910390f35b341561010157600080fd5b6100bd600160a060020a03600435811690602435166044356102ba565b604051901515815260200160405180910390f35b341561013d57600080fd5b6100e4600160a060020a03600435166102fd565b60405190815260200160405180910390f35b341561016e57600080fd5b61017661031c565b604051600160a060020a03909116815260200160405180910390f35b341561019d57600080fd5b6101b7600160a060020a036004351660243560443561032b565b005b34156101c457600080fd5b6101b7600160a060020a03600435166103e8565b005b34156101e557600080fd5b6100bd600160a060020a0360043516602435610430565b604051901515815260200160405180910390f35b341561021b57600080fd5b6100e4600160a060020a03600435811690602435166104d7565b60405190815260200160405180910390f35b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a35060015b92915050565b60005481565b600160a060020a0380841660009081526002602090815260408083203390941683529290529081208054830190556102f3848484610504565b90505b9392505050565b600160a060020a0381166000908152600160205260409020545b919050565b600354600160a060020a031681565b60035433600160a060020a0390811691161461034657600080fd5b600160a060020a038316600081815260016020526040808220805486019055815485019091558291907fdfb81fb379557413b0a951b4d7bf7a9df393801d8c539d5e201d6a8daeb913b99085905190815260200160405180910390a382600160a060020a031660007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a35b5b505050565b60035433600160a060020a0390811691161461040357600080fd5b6003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b50565b600160a060020a0333166000908152600160205260408120548290108015906104595750600082115b1561009657600160a060020a033381166000818152600160205260408082208054879003905592861680825290839020805486019055917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060016102ae565b600080fd5b5b92915050565b600160a060020a038083166000908152600260209081526040808320938516835292905220545b92915050565b600160a060020a0383166000908152600160205260408120548290108015906105545750600160a060020a0380851660009081526002602090815260408083203390941683529290522054829010155b80156105605750600082115b1561009657600160a060020a03808416600081815260016020908152604080832080548801905588851680845281842080548990039055600283528184203390961684529490915290819020805486900390559091907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060016102f6565b600080fd5b5b93925050505600a165627a7a72305820555a6635d7fe875400026b215357fdb2a338e25a5e0538a9213c0c31e33a17770029`

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
const StandardTokenBin = `0x6060604052341561000f57600080fd5b5b61043b8061001f6000396000f300606060405236156100755763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461007a57806318160ddd146100b057806323b872dd146100d557806370a0823114610111578063a9059cbb14610142578063dd62ed3e14610178575b600080fd5b341561008557600080fd5b61009c600160a060020a03600435166024356101af565b604051901515815260200160405180910390f35b34156100bb57600080fd5b6100c361021c565b60405190815260200160405180910390f35b34156100e057600080fd5b61009c600160a060020a0360043581169060243516604435610222565b604051901515815260200160405180910390f35b341561011c57600080fd5b6100c3600160a060020a036004351661031c565b60405190815260200160405180910390f35b341561014d57600080fd5b61009c600160a060020a036004351660243561033b565b604051901515815260200160405180910390f35b341561018357600080fd5b6100c3600160a060020a03600435811690602435166103e2565b60405190815260200160405180910390f35b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a35060015b92915050565b60005481565b600160a060020a0383166000908152600160205260408120548290108015906102725750600160a060020a0380851660009081526002602090815260408083203390941683529290522054829010155b801561027e5750600082115b1561007557600160a060020a03808416600081815260016020908152604080832080548801905588851680845281842080548990039055600283528184203390961684529490915290819020805486900390559091907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a3506001610314565b600080fd5b5b9392505050565b600160a060020a0381166000908152600160205260409020545b919050565b600160a060020a0333166000908152600160205260408120548290108015906103645750600082115b1561007557600160a060020a033381166000818152600160205260408082208054879003905592861680825290839020805486019055917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a3506001610216565b600080fd5b5b92915050565b600160a060020a038083166000908152600260209081526040808320938516835292905220545b929150505600a165627a7a72305820bc32d0c5346dc4bb7892c9cb74e14e5c1d6890c5c08f3578c93871b9e76af17e0029`

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
const SwearGameABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"beneficiary\",\"type\":\"address\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"name\":\"sig_v\",\"type\":\"uint8\"},{\"name\":\"sig_r\",\"type\":\"bytes32\"},{\"name\":\"sig_s\",\"type\":\"bytes32\"},{\"name\":\"_evidence\",\"type\":\"bytes32\"}],\"name\":\"newCase\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"playerCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_player\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[{\"name\":\"registered\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_ensAddr\",\"type\":\"address\"}],\"name\":\"setENSAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ids\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ensAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"id\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"caseContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"players\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_CaseContract\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_reward\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"constructor\"},{\"payable\":true,\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"decide\",\"type\":\"string\"}],\"name\":\"Decision\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"depositAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"DepositStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"Compensate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"playerId\",\"type\":\"address\"}],\"name\":\"NewPlayer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"playerId\",\"type\":\"address\"}],\"name\":\"PlayerLeftGame\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"plaintiff\",\"type\":\"address\"}],\"name\":\"NewClaimOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"plaintiff\",\"type\":\"address\"}],\"name\":\"NewEvidenceSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"plaintiff\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"ClaimResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Payment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"AdditionalDepositRequired\",\"type\":\"event\"}]"

// SwearGameBin is the compiled bytecode used for deploying new contracts.
const SwearGameBin = `0x6060604052600a8054600160a060020a031916738163bc885c2b14478b75f178ca76f31581dc967f179055341561003557600080fd5b604051606080610ff18339810160405280805191906020018051919060200180519150505b5b60008054600160a060020a03191633600160a060020a03161790555b60068054600160a060020a03808616600160a060020a0319928316179092556007805492851692909116919091179055600281905560006003555b5050505b610f2c806100c56000396000f300606060405236156100b45763ffffffff60e060020a6000350416631845c55881146101a2578063228cb733146101e7578063302bcc571461020c5780634420e486146102315780634cad1ce9146102645780638da5cb5b146102975780639f20bc92146102c6578063a6b694f2146102fa578063a6f9dae114610329578063af640d0f1461034a578063d0e30db01461036f578063d556e8e614610394578063e2eb41ff146103c3578063fc0c546a146103f6575b6101a05b600754600080543492600160a060020a03908116926323b872dd9290911690309085906040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b151561013157600080fd5b6102c65a03f1151561014257600080fd5b50505060405180519050151561015757600080fd5b600180548201908190557f3a7e173a9698235104076a0f536c5169527adecd20d0938c7c156ddf776c859190829060405191825260208201526040908101905180910390a15b50565b005b34156101ad57600080fd5b6101d3600160a060020a036004351660243560ff6044351660643560843560a435610425565b604051901515815260200160405180910390f35b34156101f257600080fd5b6101fa6105d2565b60405190815260200160405180910390f35b341561021757600080fd5b6101fa6105d8565b60405190815260200160405180910390f35b341561023c57600080fd5b6101d3600160a060020a03600435166105de565b604051901515815260200160405180910390f35b341561026f57600080fd5b6101d3600160a060020a036004351661070b565b604051901515815260200160405180910390f35b34156102a257600080fd5b6102aa61073b565b604051600160a060020a03909116815260200160405180910390f35b34156102d157600080fd5b6101fa600160a060020a036004351660243561074a565b60405190815260200160405180910390f35b341561030557600080fd5b6102aa61077c565b604051600160a060020a03909116815260200160405180910390f35b341561033457600080fd5b6101a0600160a060020a036004351661078b565b005b341561035557600080fd5b6101fa6107d3565b60405190815260200160405180910390f35b341561037a57600080fd5b6101fa6107d9565b60405190815260200160405180910390f35b341561039f57600080fd5b6102aa6107df565b604051600160a060020a03909116815260200160405180910390f35b34156103ce57600080fd5b6101d3600160a060020a03600435166107ee565b604051901515815260200160405180910390f35b341561040157600080fd5b6102aa610803565b604051600160a060020a03909116815260200160405180910390f35b600160a060020a03331660009081526004602052604081205460ff16151561044c57600080fd5b600654600160a060020a031663e391c4c0338460006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b15156104ab57600080fd5b6102c65a03f115156104bc57600080fd5b50505060405180516008819055151590506104d9575060006105c8565b600654600854600160a060020a039091169063fb9e6e7790898989898960006040516020015260405160e060020a63ffffffff89160281526004810196909652600160a060020a039094166024860152604485019290925260ff166064840152608483015260a482015260c401602060405180830381600087803b151561055f57600080fd5b6102c65a03f1151561057057600080fd5b50505060405180515050600160a060020a03331660009081526005602052604090208054600181016105a28382610eb5565b916000526020600020900160005b50600854908190556105c29150610812565b50600190505b9695505050505050565b60025481565b60035481565b6000805433600160a060020a039081169116146105fa57600080fd5b600160a060020a03821660009081526004602052604090205460ff161561062057600080fd5b600354151561063f57600254600154101561063a57600080fd5b610692565b60025460035460015481151561065157fe5b041015610692577fa8ff0fbf8dc82e5914523ed3a828920f55ab93494c59933e8c45d79f8bba70d760015460405190815260200160405180910390a1600080fd5b5b600160a060020a03821660009081526004602052604090819020805460ff191660019081179091556003805490910190557f52e92d4898337244a39bd42674ac561eadfd3959e947deec1c0ab82dd58b5a7590839051600160a060020a03909116815260200160405180910390a15060015b5b919050565b600a805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03831617905560015b919050565b600054600160a060020a031681565b60056020528160005260406000208181548110151561076557fe5b906000526020600020900160005b91509150505481565b600a54600160a060020a031681565b60005433600160a060020a039081169116146107a657600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b50565b60085481565b60015481565b600654600160a060020a031681565b60046020526000908152604090205460ff1681565b600754600160a060020a031681565b60065460009081908190819081908190819081908190600160a060020a0316635de28ae08b836040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561087557600080fd5b6102c65a03f1151561088657600080fd5b50505060405180511515905061089b57600080fd5b600654600160a060020a031663c9100bcb8b600060405160c0015260405160e060020a63ffffffff8416028152600481019190915260240160c060405180830381600087803b15156108ec57600080fd5b6102c65a03f115156108fd57600080fd5b5050506040518051906020018051906020018051906020018051906020018051906020018051905095509550955095509550955061093e8585858585610a2f565b801561094d575061094d610b1f565b5b9650600187151514156109cc5761096486610c85565b600654909850600160a060020a031663c1203cb88b60405160e060020a63ffffffff84160281526004810191909152602401600060405180830381600087803b15156109af57600080fd5b6102c65a03f115156109c057600080fd5b5050506109cc86610d34565b5b7f62c32f975d56ca09b1255432b0bbe6b55f0cea57bf9efab6714409e47a5bf6838a87600254604051928352600160a060020a0390911660208301526040808301919091526060909101905180910390a18798505b5050505050505050919050565b60008043869010610a435760009150610b15565b3087876040516c01000000000000000000000000600160a060020a03948516810282529290931690910260148301526028820152604801604051809103902090506001818686866040518060005260200160405260006040516020015260405193845260ff90921660208085019190915260408085019290925260608401929092526080909201915160208103908084039060008661646e5a03f11515610ae957600080fd5b505060206040510351600054600160a060020a03908116911614610b105760009150610b15565b600191505b5095945050505050565b600a546009805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039092169190911790556000610b797facd7f5ed7d93b1526477b93e6c7def60c40420a868e7f694a7671413d89bb9a5610dc6565b610ba27f94c4860d894e91f2df683b61455630d721209c6265d2e80c86a1f92cab14b370610dc6565b14610c11577f3cf8334d363b8556cf92ddf8f803d5074c4f884fc6b92484e8f2b22fe3d5fc8d60405160208082526006908201527f6775696c747900000000000000000000000000000000000000000000000000006040808301919091526060909101905180910390a1610c7e565b7f3cf8334d363b8556cf92ddf8f803d5074c4f884fc6b92484e8f2b22fe3d5fc8d6040516020808252600a908201527f6e6f74206775696c7479000000000000000000000000000000000000000000006040808301919091526060909101905180910390a1506000610c82565b5060015b90565b600754600254600091600160a060020a0316906323b872dd9030908590856040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b1515610cf957600080fd5b6102c65a03f11515610d0a57600080fd5b5050506040518051915050801515610d2157600080fd5b600254600180549190910390555b919050565b600160a060020a03811660009081526004602052604090205460ff161515610d5b57600080fd5b7f3def0aea61e344e47b98a0523a5825c0919e88f7b33d0ac7f889a7654880f9a581604051600160a060020a03909116815260200160405180910390a1600160a060020a0381166000908152600460205260409020805460ff19169055600380546000190190555b50565b600954600090819081908190600160a060020a0316630178b8bf86836040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515610e1f57600080fd5b6102c65a03f11515610e3057600080fd5b5050506040518051935083925050600160a060020a038216632dff69418660006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515610e8d57600080fd5b6102c65a03f11515610e9e57600080fd5b50505060405180519450849150505b505050919050565b815481835581811511610ed957600083815260209020610ed9918101908301610edf565b5b505050565b610c8291905b80821115610ef95760008155600101610ee5565b5090565b905600a165627a7a723058202c5398aa928d2603cf8240a62d22f36ef7be6534c5e6b5a9fc63509631f04bd70029`

// DeploySwearGame deploys a new Ethereum contract, binding an instance of SwearGame to it.
func DeploySwearGame(auth *bind.TransactOpts, backend bind.ContractBackend, _CaseContract common.Address, _token common.Address, _reward *big.Int) (common.Address, *types.Transaction, *SwearGame, error) {
	parsed, err := abi.JSON(strings.NewReader(SwearGameABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SwearGameBin), backend, _CaseContract, _token, _reward)
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

// EnsAddress is a free data retrieval call binding the contract method 0xa6b694f2.
//
// Solidity: function ensAddress() constant returns(address)
func (_SwearGame *SwearGameCaller) EnsAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SwearGame.contract.Call(opts, out, "ensAddress")
	return *ret0, err
}

// EnsAddress is a free data retrieval call binding the contract method 0xa6b694f2.
//
// Solidity: function ensAddress() constant returns(address)
func (_SwearGame *SwearGameSession) EnsAddress() (common.Address, error) {
	return _SwearGame.Contract.EnsAddress(&_SwearGame.CallOpts)
}

// EnsAddress is a free data retrieval call binding the contract method 0xa6b694f2.
//
// Solidity: function ensAddress() constant returns(address)
func (_SwearGame *SwearGameCallerSession) EnsAddress() (common.Address, error) {
	return _SwearGame.Contract.EnsAddress(&_SwearGame.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() constant returns(bytes32)
func (_SwearGame *SwearGameCaller) Id(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _SwearGame.contract.Call(opts, out, "id")
	return *ret0, err
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() constant returns(bytes32)
func (_SwearGame *SwearGameSession) Id() ([32]byte, error) {
	return _SwearGame.Contract.Id(&_SwearGame.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() constant returns(bytes32)
func (_SwearGame *SwearGameCallerSession) Id() ([32]byte, error) {
	return _SwearGame.Contract.Id(&_SwearGame.CallOpts)
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

// NewCase is a paid mutator transaction binding the contract method 0x1845c558.
//
// Solidity: function newCase(beneficiary address, blockNumber uint256, sig_v uint8, sig_r bytes32, sig_s bytes32, _evidence bytes32) returns(bool)
func (_SwearGame *SwearGameTransactor) NewCase(opts *bind.TransactOpts, beneficiary common.Address, blockNumber *big.Int, sig_v uint8, sig_r [32]byte, sig_s [32]byte, _evidence [32]byte) (*types.Transaction, error) {
	return _SwearGame.contract.Transact(opts, "newCase", beneficiary, blockNumber, sig_v, sig_r, sig_s, _evidence)
}

// NewCase is a paid mutator transaction binding the contract method 0x1845c558.
//
// Solidity: function newCase(beneficiary address, blockNumber uint256, sig_v uint8, sig_r bytes32, sig_s bytes32, _evidence bytes32) returns(bool)
func (_SwearGame *SwearGameSession) NewCase(beneficiary common.Address, blockNumber *big.Int, sig_v uint8, sig_r [32]byte, sig_s [32]byte, _evidence [32]byte) (*types.Transaction, error) {
	return _SwearGame.Contract.NewCase(&_SwearGame.TransactOpts, beneficiary, blockNumber, sig_v, sig_r, sig_s, _evidence)
}

// NewCase is a paid mutator transaction binding the contract method 0x1845c558.
//
// Solidity: function newCase(beneficiary address, blockNumber uint256, sig_v uint8, sig_r bytes32, sig_s bytes32, _evidence bytes32) returns(bool)
func (_SwearGame *SwearGameTransactorSession) NewCase(beneficiary common.Address, blockNumber *big.Int, sig_v uint8, sig_r [32]byte, sig_s [32]byte, _evidence [32]byte) (*types.Transaction, error) {
	return _SwearGame.Contract.NewCase(&_SwearGame.TransactOpts, beneficiary, blockNumber, sig_v, sig_r, sig_s, _evidence)
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

// SetENSAddress is a paid mutator transaction binding the contract method 0x4cad1ce9.
//
// Solidity: function setENSAddress(_ensAddr address) returns(bool)
func (_SwearGame *SwearGameTransactor) SetENSAddress(opts *bind.TransactOpts, _ensAddr common.Address) (*types.Transaction, error) {
	return _SwearGame.contract.Transact(opts, "setENSAddress", _ensAddr)
}

// SetENSAddress is a paid mutator transaction binding the contract method 0x4cad1ce9.
//
// Solidity: function setENSAddress(_ensAddr address) returns(bool)
func (_SwearGame *SwearGameSession) SetENSAddress(_ensAddr common.Address) (*types.Transaction, error) {
	return _SwearGame.Contract.SetENSAddress(&_SwearGame.TransactOpts, _ensAddr)
}

// SetENSAddress is a paid mutator transaction binding the contract method 0x4cad1ce9.
//
// Solidity: function setENSAddress(_ensAddr address) returns(bool)
func (_SwearGame *SwearGameTransactorSession) SetENSAddress(_ensAddr common.Address) (*types.Transaction, error) {
	return _SwearGame.Contract.SetENSAddress(&_SwearGame.TransactOpts, _ensAddr)
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
