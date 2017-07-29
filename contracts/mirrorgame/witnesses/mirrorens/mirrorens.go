// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package mirrorens

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

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

// MirrorENSABI is the input ABI used to generate the binding from.
const MirrorENSABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_ensAddr\",\"type\":\"address\"}],\"name\":\"setENSAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"caseId\",\"type\":\"bytes32\"},{\"name\":\"serviceId\",\"type\":\"bytes32\"},{\"name\":\"clientNameHash\",\"type\":\"bytes32\"},{\"name\":\"serviceNameHash\",\"type\":\"bytes32\"}],\"name\":\"submitNameHashes\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"caseId\",\"type\":\"bytes32\"},{\"name\":\"serviceId\",\"type\":\"bytes32\"},{\"name\":\"clientAddress\",\"type\":\"address\"}],\"name\":\"isEvidenceSubmitted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ensAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"caseId\",\"type\":\"bytes32\"},{\"name\":\"serviceId\",\"type\":\"bytes32\"},{\"name\":\"clientAddress\",\"type\":\"address\"}],\"name\":\"testimonyFor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"}]"

// MirrorENSBin is the compiled bytecode used for deploying new contracts.
const MirrorENSBin = `0x606060405260028054600160a060020a031916738163bc885c2b14478b75f178ca76f31581dc967f179055341561003557600080fd5b5b5b5b6103fb806100476000396000f300606060405263ffffffff60e060020a6000350416634cad1ce981146100505780636fc34567146100835780637aba9a67146100b6578063a6b694f2146100ef578063c4b452831461011e575b600080fd5b341561005b57600080fd5b61006f600160a060020a0360043516610167565b604051901515815260200160405180910390f35b341561008e57600080fd5b61006f600435602435604435606435610197565b604051901515815260200160405180910390f35b34156100c157600080fd5b61006f600435602435600160a060020a0360443516610200565b604051901515815260200160405180910390f35b34156100fa57600080fd5b610102610222565b604051600160a060020a03909116815260200160405180910390f35b341561012957600080fd5b610143600435602435600160a060020a0360443516610231565b6040518082600281111561015357fe5b60ff16815260200191505060405180910390f35b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03831617905560015b919050565b600084815260208181526040808320868452909152812054156101bc575060006101f8565b60408051908101604090815284825260208083018590526000888152808252828120888252909152208151815560208201516001918201559150505b949350505050565b60008381526020818152604080832085845290915290205415155b9392505050565b600254600160a060020a031681565b600061023d8484610255565b1561024a5750600061021b565b5060015b9392505050565b6002546001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039092169190911781556000838152602081815260408083208584529091528120909101546102a8906102e0565b6000848152602081815260408083208684529091529020546102c9906102e0565b146102d6575060016102da565b5060005b92915050565b600154600090819081908190600160a060020a0316630178b8bf86836040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561033957600080fd5b6102c65a03f1151561034a57600080fd5b5050506040518051935083925050600160a060020a038216632dff69418660006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156103a757600080fd5b6102c65a03f115156103b857600080fd5b50505060405180519450849150505b5050509190505600a165627a7a723058209b7176bd1e113aae4e6de6ab51b74e47056d7b7c85002217ccc04b9cde6b451c0029`

// DeployMirrorENS deploys a new Ethereum contract, binding an instance of MirrorENS to it.
func DeployMirrorENS(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MirrorENS, error) {
	parsed, err := abi.JSON(strings.NewReader(MirrorENSABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MirrorENSBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MirrorENS{MirrorENSCaller: MirrorENSCaller{contract: contract}, MirrorENSTransactor: MirrorENSTransactor{contract: contract}}, nil
}

// MirrorENS is an auto generated Go binding around an Ethereum contract.
type MirrorENS struct {
	MirrorENSCaller     // Read-only binding to the contract
	MirrorENSTransactor // Write-only binding to the contract
}

// MirrorENSCaller is an auto generated read-only Go binding around an Ethereum contract.
type MirrorENSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MirrorENSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MirrorENSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MirrorENSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MirrorENSSession struct {
	Contract     *MirrorENS        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MirrorENSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MirrorENSCallerSession struct {
	Contract *MirrorENSCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MirrorENSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MirrorENSTransactorSession struct {
	Contract     *MirrorENSTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MirrorENSRaw is an auto generated low-level Go binding around an Ethereum contract.
type MirrorENSRaw struct {
	Contract *MirrorENS // Generic contract binding to access the raw methods on
}

// MirrorENSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MirrorENSCallerRaw struct {
	Contract *MirrorENSCaller // Generic read-only contract binding to access the raw methods on
}

// MirrorENSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MirrorENSTransactorRaw struct {
	Contract *MirrorENSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMirrorENS creates a new instance of MirrorENS, bound to a specific deployed contract.
func NewMirrorENS(address common.Address, backend bind.ContractBackend) (*MirrorENS, error) {
	contract, err := bindMirrorENS(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MirrorENS{MirrorENSCaller: MirrorENSCaller{contract: contract}, MirrorENSTransactor: MirrorENSTransactor{contract: contract}}, nil
}

// NewMirrorENSCaller creates a new read-only instance of MirrorENS, bound to a specific deployed contract.
func NewMirrorENSCaller(address common.Address, caller bind.ContractCaller) (*MirrorENSCaller, error) {
	contract, err := bindMirrorENS(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MirrorENSCaller{contract: contract}, nil
}

// NewMirrorENSTransactor creates a new write-only instance of MirrorENS, bound to a specific deployed contract.
func NewMirrorENSTransactor(address common.Address, transactor bind.ContractTransactor) (*MirrorENSTransactor, error) {
	contract, err := bindMirrorENS(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MirrorENSTransactor{contract: contract}, nil
}

// bindMirrorENS binds a generic wrapper to an already deployed contract.
func bindMirrorENS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MirrorENSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MirrorENS *MirrorENSRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MirrorENS.Contract.MirrorENSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MirrorENS *MirrorENSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MirrorENS.Contract.MirrorENSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MirrorENS *MirrorENSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MirrorENS.Contract.MirrorENSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MirrorENS *MirrorENSCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MirrorENS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MirrorENS *MirrorENSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MirrorENS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MirrorENS *MirrorENSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MirrorENS.Contract.contract.Transact(opts, method, params...)
}

// EnsAddress is a free data retrieval call binding the contract method 0xa6b694f2.
//
// Solidity: function ensAddress() constant returns(address)
func (_MirrorENS *MirrorENSCaller) EnsAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MirrorENS.contract.Call(opts, out, "ensAddress")
	return *ret0, err
}

// EnsAddress is a free data retrieval call binding the contract method 0xa6b694f2.
//
// Solidity: function ensAddress() constant returns(address)
func (_MirrorENS *MirrorENSSession) EnsAddress() (common.Address, error) {
	return _MirrorENS.Contract.EnsAddress(&_MirrorENS.CallOpts)
}

// EnsAddress is a free data retrieval call binding the contract method 0xa6b694f2.
//
// Solidity: function ensAddress() constant returns(address)
func (_MirrorENS *MirrorENSCallerSession) EnsAddress() (common.Address, error) {
	return _MirrorENS.Contract.EnsAddress(&_MirrorENS.CallOpts)
}

// IsEvidenceSubmitted is a paid mutator transaction binding the contract method 0x7aba9a67.
//
// Solidity: function isEvidenceSubmitted(caseId bytes32, serviceId bytes32, clientAddress address) returns(bool)
func (_MirrorENS *MirrorENSTransactor) IsEvidenceSubmitted(opts *bind.TransactOpts, caseId [32]byte, serviceId [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _MirrorENS.contract.Transact(opts, "isEvidenceSubmitted", caseId, serviceId, clientAddress)
}

// IsEvidenceSubmitted is a paid mutator transaction binding the contract method 0x7aba9a67.
//
// Solidity: function isEvidenceSubmitted(caseId bytes32, serviceId bytes32, clientAddress address) returns(bool)
func (_MirrorENS *MirrorENSSession) IsEvidenceSubmitted(caseId [32]byte, serviceId [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _MirrorENS.Contract.IsEvidenceSubmitted(&_MirrorENS.TransactOpts, caseId, serviceId, clientAddress)
}

// IsEvidenceSubmitted is a paid mutator transaction binding the contract method 0x7aba9a67.
//
// Solidity: function isEvidenceSubmitted(caseId bytes32, serviceId bytes32, clientAddress address) returns(bool)
func (_MirrorENS *MirrorENSTransactorSession) IsEvidenceSubmitted(caseId [32]byte, serviceId [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _MirrorENS.Contract.IsEvidenceSubmitted(&_MirrorENS.TransactOpts, caseId, serviceId, clientAddress)
}

// SetENSAddress is a paid mutator transaction binding the contract method 0x4cad1ce9.
//
// Solidity: function setENSAddress(_ensAddr address) returns(bool)
func (_MirrorENS *MirrorENSTransactor) SetENSAddress(opts *bind.TransactOpts, _ensAddr common.Address) (*types.Transaction, error) {
	return _MirrorENS.contract.Transact(opts, "setENSAddress", _ensAddr)
}

// SetENSAddress is a paid mutator transaction binding the contract method 0x4cad1ce9.
//
// Solidity: function setENSAddress(_ensAddr address) returns(bool)
func (_MirrorENS *MirrorENSSession) SetENSAddress(_ensAddr common.Address) (*types.Transaction, error) {
	return _MirrorENS.Contract.SetENSAddress(&_MirrorENS.TransactOpts, _ensAddr)
}

// SetENSAddress is a paid mutator transaction binding the contract method 0x4cad1ce9.
//
// Solidity: function setENSAddress(_ensAddr address) returns(bool)
func (_MirrorENS *MirrorENSTransactorSession) SetENSAddress(_ensAddr common.Address) (*types.Transaction, error) {
	return _MirrorENS.Contract.SetENSAddress(&_MirrorENS.TransactOpts, _ensAddr)
}

// SubmitNameHashes is a paid mutator transaction binding the contract method 0x6fc34567.
//
// Solidity: function submitNameHashes(caseId bytes32, serviceId bytes32, clientNameHash bytes32, serviceNameHash bytes32) returns(bool)
func (_MirrorENS *MirrorENSTransactor) SubmitNameHashes(opts *bind.TransactOpts, caseId [32]byte, serviceId [32]byte, clientNameHash [32]byte, serviceNameHash [32]byte) (*types.Transaction, error) {
	return _MirrorENS.contract.Transact(opts, "submitNameHashes", caseId, serviceId, clientNameHash, serviceNameHash)
}

// SubmitNameHashes is a paid mutator transaction binding the contract method 0x6fc34567.
//
// Solidity: function submitNameHashes(caseId bytes32, serviceId bytes32, clientNameHash bytes32, serviceNameHash bytes32) returns(bool)
func (_MirrorENS *MirrorENSSession) SubmitNameHashes(caseId [32]byte, serviceId [32]byte, clientNameHash [32]byte, serviceNameHash [32]byte) (*types.Transaction, error) {
	return _MirrorENS.Contract.SubmitNameHashes(&_MirrorENS.TransactOpts, caseId, serviceId, clientNameHash, serviceNameHash)
}

// SubmitNameHashes is a paid mutator transaction binding the contract method 0x6fc34567.
//
// Solidity: function submitNameHashes(caseId bytes32, serviceId bytes32, clientNameHash bytes32, serviceNameHash bytes32) returns(bool)
func (_MirrorENS *MirrorENSTransactorSession) SubmitNameHashes(caseId [32]byte, serviceId [32]byte, clientNameHash [32]byte, serviceNameHash [32]byte) (*types.Transaction, error) {
	return _MirrorENS.Contract.SubmitNameHashes(&_MirrorENS.TransactOpts, caseId, serviceId, clientNameHash, serviceNameHash)
}

// TestimonyFor is a paid mutator transaction binding the contract method 0xc4b45283.
//
// Solidity: function testimonyFor(caseId bytes32, serviceId bytes32, clientAddress address) returns(uint8)
func (_MirrorENS *MirrorENSTransactor) TestimonyFor(opts *bind.TransactOpts, caseId [32]byte, serviceId [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _MirrorENS.contract.Transact(opts, "testimonyFor", caseId, serviceId, clientAddress)
}

// TestimonyFor is a paid mutator transaction binding the contract method 0xc4b45283.
//
// Solidity: function testimonyFor(caseId bytes32, serviceId bytes32, clientAddress address) returns(uint8)
func (_MirrorENS *MirrorENSSession) TestimonyFor(caseId [32]byte, serviceId [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _MirrorENS.Contract.TestimonyFor(&_MirrorENS.TransactOpts, caseId, serviceId, clientAddress)
}

// TestimonyFor is a paid mutator transaction binding the contract method 0xc4b45283.
//
// Solidity: function testimonyFor(caseId bytes32, serviceId bytes32, clientAddress address) returns(uint8)
func (_MirrorENS *MirrorENSTransactorSession) TestimonyFor(caseId [32]byte, serviceId [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _MirrorENS.Contract.TestimonyFor(&_MirrorENS.TransactOpts, caseId, serviceId, clientAddress)
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
