package contracts

//go:generate abigen --sol ./courtroom.sol --pkg contracts --out ./courtroom.go
//go:generate abigen --sol ./mirrorens.sol --pkg mirrorens --out ./mirrorgame/witnesses/mirrorens/mirrorens.go
//go:generate abigen --sol ./promisevalidator.sol --pkg promisevalidator --out ./mirrorgame/witnesses/promisevalidator/promisevalidator.go
//go:generate abigen --sol ./mirrorrules.sol --pkg mirrorrules --out ./mirrorgame/rules/mirrorrules.go
//go:generate abigen --sol ./mirrorregistrar.sol --pkg mirrorregistrar --out ./mirrorgame/registrar/mirrorregistrar.go

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/ens/contract"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	mirrorregistrar "github.com/jaakmusic/swap-swear-and-swindle/contracts/mirrorgame/registrar"
	mirrorrules "github.com/jaakmusic/swap-swear-and-swindle/contracts/mirrorgame/rules"
	mirrorens "github.com/jaakmusic/swap-swear-and-swindle/contracts/mirrorgame/witnesses/mirrorens"
	promisevalidator "github.com/jaakmusic/swap-swear-and-swindle/contracts/mirrorgame/witnesses/promisevalidator"
)

var (
	serviceKey, _         = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291") //service
	ethKey, _             = crypto.HexToECDSA("8a1f9a8f95be41cd7ccb6168179afb4504aefe388d1e14474d32c45c72ce7b7a")
	clientKey, _          = crypto.HexToECDSA("49a7b37aa6f6645917e7b807e9d1c00d4fa71f18343b0d4122a4d2df64dd6fee")
	serviceAddr           = crypto.PubkeyToAddress(serviceKey.PublicKey)
	ethAddr               = crypto.PubkeyToAddress(ethKey.PublicKey)
	clientAddr            = crypto.PubkeyToAddress(clientKey.PublicKey)
	blockNumber           = 0
	compensationAmount    = int64(5)
	PromiseTillNextBlocks = 51
	serviceDeposit        = int64(50)
	witnessesNumber       = int64(2)
	serviceId             = [32]byte{1, 2, 3, 4}
	clientENSName         = "client.game"
	serviceENSName        = "service.game"
	depositVestingPeriod  = int64(100) //100 blocks
)

type TheGame struct {
	swearGameContractAddress common.Address
	swearGame                *Swear
	sampleToken              *SampleToken
	promiseValidator         *promisevalidator.PromiseValidator
	promiseValidatorAddress  common.Address
	mirrorEns                *mirrorens.MirrorENS
	mirrorRules              *mirrorrules.MirrorRules
	mirrorRegistrar          *mirrorregistrar.MirrorRegistrar
}

func newTestBackend() *backends.SimulatedBackend {
	return backends.NewSimulatedBackend(core.GenesisAlloc{
		serviceAddr: {Balance: big.NewInt(1000000000)},
		ethAddr:     {Balance: big.NewInt(1000000000)},
		clientAddr:  {Balance: big.NewInt(1000000000)},
	})
}

//commit new block to the backends.SimulatedBackend
//the blockcount is increment on each commit.
func commit(b *backends.SimulatedBackend) {
	b.Commit()
	blockNumber++
}

func deployMirror(prvKey *ecdsa.PrivateKey, amount *big.Int, backend *backends.SimulatedBackend) (common.Address, *mirrorens.MirrorENS, error) {
	deployTransactor := bind.NewKeyedTransactor(prvKey)
	deployTransactor.Value = amount

	addr, _, mirror, err := mirrorens.DeployMirrorENS(deployTransactor, backend)
	if err != nil {
		return common.Address{}, nil, err
	}
	commit(backend)
	return addr, mirror, nil
}

func deployPromiseValidator(prvKey *ecdsa.PrivateKey, amount *big.Int, backend *backends.SimulatedBackend) (common.Address, *promisevalidator.PromiseValidator, error) {
	deployTransactor := bind.NewKeyedTransactor(prvKey)
	deployTransactor.Value = amount

	addr, _, promiseValidator, err := promisevalidator.DeployPromiseValidator(deployTransactor, backend)
	if err != nil {
		return common.Address{}, nil, err
	}
	commit(backend)
	return addr, promiseValidator, nil
}

func deployMirrorRules(prvKey *ecdsa.PrivateKey, amount *big.Int, backend *backends.SimulatedBackend, paymentValidatorContract common.Address, ENSMirrotValidatorContract common.Address) (common.Address, *mirrorrules.MirrorRules, error) {
	deployTransactor := bind.NewKeyedTransactor(prvKey)
	deployTransactor.Value = amount

	addr, _, mirrorRules, err := mirrorrules.DeployMirrorRules(deployTransactor, backend, paymentValidatorContract, ENSMirrotValidatorContract)
	if err != nil {
		return common.Address{}, nil, err
	}
	commit(backend)
	return addr, mirrorRules, nil
}

func deployMirrorRegistrar(prvKey *ecdsa.PrivateKey, amount *big.Int, backend *backends.SimulatedBackend, trialRules common.Address, token common.Address) (common.Address, *mirrorregistrar.MirrorRegistrar, error) {
	deployTransactor := bind.NewKeyedTransactor(prvKey)
	deployTransactor.Value = amount

	addr, _, mirrorRegistrar, err := mirrorregistrar.DeployMirrorRegistrar(deployTransactor, backend, trialRules, token)
	if err != nil {
		return common.Address{}, nil, err
	}
	commit(backend)
	return addr, mirrorRegistrar, nil
}

func deploySwear(prvKey *ecdsa.PrivateKey, amount *big.Int, backend *backends.SimulatedBackend, tokenContractAddr common.Address, mirrorTransistions common.Address, rewordCompansation *big.Int) (common.Address, *Swear, error) {
	deployTransactor := bind.NewKeyedTransactor(prvKey)
	deployTransactor.Value = amount

	addr, _, swearGame, err := DeploySwear(deployTransactor, backend, tokenContractAddr, mirrorTransistions)
	if err != nil {
		return common.Address{}, nil, err
	}
	commit(backend)
	return addr, swearGame, nil
}

func deploySampleTokenContract(prvKey *ecdsa.PrivateKey, amount *big.Int, backend *backends.SimulatedBackend, initialSupply *big.Int) (common.Address, *SampleToken, error) {
	deployTransactor := bind.NewKeyedTransactor(prvKey)
	deployTransactor.Value = amount

	addr, _, sampleToken, err := DeploySampleToken(deployTransactor, backend, initialSupply)

	if err != nil {
		return common.Address{}, nil, err
	}
	commit(backend)
	return addr, sampleToken, nil
}

func deployENS(prvKey *ecdsa.PrivateKey, amount *big.Int, backend *backends.SimulatedBackend, owner common.Address) (common.Address, common.Address, *contract.ENS, *contract.PublicResolver, error) {
	deployTransactor := bind.NewKeyedTransactor(prvKey)
	deployTransactor.Value = amount

	ensAddress, _, ens, err := contract.DeployENS(deployTransactor, backend, owner)

	if err != nil {
		return common.Address{}, common.Address{}, nil, nil, err
	}
	commit(backend)

	addr, _, publicResolver, err := contract.DeployPublicResolver(deployTransactor, backend, ensAddress)
	if err != nil {
		return common.Address{}, common.Address{}, nil, nil, err
	}
	commit(backend)

	//setting up .game domain

	_, err = ens.SetOwner(deployTransactor, common.Hash{}, ethAddr)

	if err != nil {
		return common.Address{}, common.Address{}, nil, nil, err
	}
	commit(backend)
	_, err = ens.SetSubnodeOwner(deployTransactor, common.Hash{}, SHA3("game"), ethAddr)

	if err != nil {
		return common.Address{}, common.Address{}, nil, nil, err
	}
	commit(backend)

	return ensAddress, addr, ens, publicResolver, nil
}

func deposit(t *testing.T, backend *backends.SimulatedBackend, sampleToken *SampleToken, registrar *mirrorregistrar.MirrorRegistrar, vestingPeriod int64) {
	//trasfer 100 tokens to the Service
	opts := bind.NewKeyedTransactor(ethKey)
	_, err := sampleToken.Transfer(opts, serviceAddr, big.NewInt(100))
	if err != nil {
		t.Fatalf("trasfer tokens to service: expected no error, got %v", err)
	}
	commit(backend)

	//deposit 50 token to the contract
	opts = bind.NewKeyedTransactor(serviceKey)
	opts.Value = big.NewInt(serviceDeposit)
	_, err = registrar.Deposit(opts, big.NewInt(vestingPeriod))
	if err != nil {
		t.Fatalf("depost tokens to contract: expected no error, got %v", err)
	}
	commit(backend)
}

func openCaseForMirrorGame(t *testing.T, clientContent string, serviceContent string, numberOfBlocksToWait int, pendingTest bool) {
	backend := newTestBackend()

	theGame := deployTheGame(t, backend)

	opts := bind.NewKeyedTransactor(serviceKey)

	deposit(t, backend, theGame.sampleToken, theGame.mirrorRegistrar, depositVestingPeriod)

	_, err := theGame.mirrorRegistrar.Register(opts, clientAddr)
	if err != nil {
		t.Fatalf("Register: expected no error, got %v", err)
	}
	commit(backend)

	ensAddress, resolverAddr, ens, publicResolver, err := deployENS(ethKey, big.NewInt(0), backend, ethAddr)
	t.Log("ensAddress", ensAddress.String(), resolverAddr.String())
	if err != nil {
		t.Fatalf("deployENS: expected no error, got %v", err)
	}

	err = registerENSRecord(t, backend, serviceENSName, serviceContent, ens, publicResolver, resolverAddr)
	if err != nil {

		t.Fatal("registerENSRecord service.game fail")
	}
	err = registerENSRecord(t, backend, clientENSName, clientContent, ens, publicResolver, resolverAddr)
	if err != nil {
		t.Fatal("registerENSRecord", clientENSName, "fail")
	}
	depositBefore, err := theGame.mirrorRegistrar.Deposits(&bind.CallOpts{}, serviceAddr)
	if err != nil {
		t.Fatalf("AmountStaked : expected no error, got %v", err)
	}
	t.Log("deposit", depositBefore)
	balanceOfClientBefore, err := theGame.sampleToken.BalanceOf(&bind.CallOpts{}, clientAddr)
	if err != nil {
		t.Fatalf("balanceOfClientBefore : expected no error, got %v", err)
	}
	opts = bind.NewKeyedTransactor(clientKey)
	fmt.Println("add", theGame.swearGameContractAddress.String())
	_, err = theGame.swearGame.NewCase(opts, serviceId)

	if err != nil {
		t.Fatalf("NewCase: expected no error, got %v", err)
	}
	commit(backend)

	caseId, err := theGame.swearGame.Ids(&bind.CallOpts{}, clientAddr, big.NewInt(0))
	t.Log("case", caseId)

	_, err = theGame.mirrorEns.SubmitNameHashes(opts, caseId, serviceId, Namehash(clientENSName), Namehash(serviceENSName))
	if err != nil {
		t.Fatalf("SubmitNameHashes:  expected no error, got %v", err)
	}
	commit(backend)

	if !pendingTest {
		promise, err := issuePromise(serviceKey, clientAddr, big.NewInt(int64(blockNumber+PromiseTillNextBlocks)), theGame.promiseValidatorAddress)
		if err != nil {
			t.Fatalf("NewCase: issuePromise expected no error, got %v", err)
		}
		v, r, s := sig2vrs(promise.sig)

		theGame.promiseValidator.SubmitPromise(opts, caseId, serviceId, promise.beneficiary, promise.blockNumber, v, r, s)
	}

	for i := 0; i < numberOfBlocksToWait; i++ {
		commit(backend)
	}

	_, err = theGame.swearGame.Trial(opts, caseId)
	if err != nil {
		t.Fatalf("Trial: expected no error, got %v", err)
	}

	commit(backend)

	status, err := theGame.swearGame.GetStatus(&bind.CallOpts{}, caseId)
	if err != nil {
		t.Fatalf("NewCase: expected no error, got %v", err)
	}
	t.Log("status", status)

	if pendingTest && status == 3 {
		//submit promise and restart from current state
		promise, err := issuePromise(serviceKey, clientAddr, big.NewInt(int64(blockNumber+PromiseTillNextBlocks)), theGame.promiseValidatorAddress)
		if err != nil {
			t.Fatalf("NewCase: issuePromise expected no error, got %v", err)
		}
		v, r, s := sig2vrs(promise.sig)

		theGame.promiseValidator.SubmitPromise(opts, caseId, serviceId, promise.beneficiary, promise.blockNumber, v, r, s)

		_, err = theGame.swearGame.Trial(opts, caseId)

		if err != nil {
			t.Fatalf("NewCase: expected no error, got %v", err)
		}
		commit(backend)
	}

	depositAfter, err := theGame.mirrorRegistrar.Deposits(&bind.CallOpts{}, serviceAddr)
	if err != nil {
		t.Fatalf("AmountStaked : expected no error, got %v", err)
	}
	balanceOfClientAfter, err := theGame.sampleToken.BalanceOf(&bind.CallOpts{}, clientAddr)
	if err != nil {
		t.Fatalf("balanceOfClientBefore : expected no error, got %v", err)
	}
	//If the client submit a case a valid case (ens are not equal) and the case is within the time period the service promised
	//so it is a valid case
	if (clientContent != serviceContent) && (numberOfBlocksToWait <= PromiseTillNextBlocks) {
		if depositBefore.DepositedAmount.Int64()-depositAfter.DepositedAmount.Int64() != int64(compensationAmount) {
			t.Fatalf("After a valid case proccess : deposit at the contract should reduce by %d  ( got %d)", compensationAmount, (depositBefore.DepositedAmount.Int64() - depositAfter.DepositedAmount.Int64()))
		}
		if balanceOfClientAfter.Int64()-balanceOfClientBefore.Int64() != int64(compensationAmount) {
			t.Fatalf("After a valid case proccess : the balance of client should increase by %d   ( got %d)", compensationAmount, (balanceOfClientAfter.Int64() - balanceOfClientBefore.Int64()))
		}
	} else {
		if depositBefore.DepositedAmount.Int64() != depositAfter.DepositedAmount.Int64() {
			t.Fatalf("non valid case proccess : deposit at the contract should be the same as before  ( got %d)", (depositBefore.DepositedAmount.Int64() - depositAfter.DepositedAmount.Int64()))
		}
		if balanceOfClientAfter.Int64() != balanceOfClientBefore.Int64() {
			t.Fatalf("non valid  case proccess should end by no change to the client balance   ( got %d)", (balanceOfClientAfter.Int64() - balanceOfClientBefore.Int64()))
		}
	}

}
func TestPromiseOk(t *testing.T) {
	//client and service content are diffrent  but number of block to wait is 6
	// the service promise to serve client for the next PromiseTillNextBlocks(5) blocks.
	//This test will fail if client will  get compensated for its case.
	openCaseForMirrorGame(t, "1234", "4567", 6, false)
}

func TestOpenValidCasePending(t *testing.T) {
	//client and service content are diffrent and number of block to wait is 0
	//This test will fail if client will not get compensated for its case.
	//This test also submit a newCase ...without submiting enough evidence ...
	//check the status submit the missing evidence and resume the trial on the same case.
	openCaseForMirrorGame(t, "1234", "4567", 0, true)
}

func TestOpenValidCase(t *testing.T) {
	//client and service content are diffrent and number of block to wait is 0
	//This test will fail if client will not get compensated for its case.
	openCaseForMirrorGame(t, "1234", "4567", 0, false)
}

func TestOpenNoneValidCase(t *testing.T) {
	//client and service content are the same and number of blocks to wait is 0
	//This test will fail if client will get compensated for its case.
	openCaseForMirrorGame(t, "1234", "1234", 0, false)
}

func TestCollectDepositDuringTrial(t *testing.T) {

	backend := newTestBackend()

	theGame := deployTheGame(t, backend)

	deposit(t, backend, theGame.sampleToken, theGame.mirrorRegistrar, depositVestingPeriod)

	opts := bind.NewKeyedTransactor(serviceKey)
	_, err := theGame.mirrorRegistrar.Register(opts, clientAddr)
	if err != nil {
		t.Fatalf("Register: expected no error, got %v", err)
	}
	commit(backend)

	opts = bind.NewKeyedTransactor(clientKey)

	_, err = theGame.swearGame.NewCase(opts, serviceId)

	if err != nil {
		t.Fatalf("Register: expected no error, got %v", err)
	}
	commit(backend)

	balanceBefore, err := theGame.sampleToken.BalanceOf(&bind.CallOpts{}, serviceAddr)
	if err != nil {
		t.Fatalf("BalanceOf : expected no error, got %v", err)
	}

	mirrorEpoch, err := theGame.mirrorRules.GetEpoch(&bind.CallOpts{})
	if err != nil {
		t.Fatalf("mirrorRules.GetEpoch expected no error got ", err)
	}

	var i int64 = 0

	for i = 0; i < (depositVestingPeriod*mirrorEpoch.Int64() + 1); i++ {
		commit(backend)
	}

	_, err = theGame.mirrorRegistrar.CollectDeposit(opts)

	if err != nil {
		t.Fatalf("CollectDeposit: expected no error, got %v", err)
	}
	commit(backend)

	balanceAfter, err := theGame.sampleToken.BalanceOf(&bind.CallOpts{}, serviceAddr)
	if err != nil {
		t.Fatalf("BalanceOf : expected no error, got %v", err)
	}

	if balanceBefore.Int64() != balanceAfter.Int64() {
		t.Fatalf("balance should  be equal because we are during trial and service cannot collect deposit")
	}

}
func TestCollectDeposit(t *testing.T) {
	backend := newTestBackend()

	theGame := deployTheGame(t, backend)

	deposit(t, backend, theGame.sampleToken, theGame.mirrorRegistrar, depositVestingPeriod)
	commit(backend)

	opts := bind.NewKeyedTransactor(serviceKey)

	balanceBefore, err := theGame.sampleToken.BalanceOf(&bind.CallOpts{}, serviceAddr)
	if err != nil {
		t.Fatalf("BalanceOf : expected no error, got %v", err)
	}

	_, err = theGame.mirrorRegistrar.CollectDeposit(opts)
	if err != nil {
		t.Fatalf("CollectDeposit: expected no error, got %v", err)
	}
	commit(backend)

	balanceAfter, err := theGame.sampleToken.BalanceOf(&bind.CallOpts{}, serviceAddr)
	if err != nil {
		t.Fatalf("BalanceOf : expected no error, got %v", err)
	}

	if balanceBefore.Int64() != balanceAfter.Int64() {
		t.Fatalf("balance should not be changed because vesting period not passed")
	}

	mirrorEpoch, err := theGame.mirrorRules.GetEpoch(&bind.CallOpts{})
	if err != nil {
		t.Fatalf("mirrorRules.GetEpoch expected no error got ", err)
	}

	var i int64 = 0

	for i = 0; i < (depositVestingPeriod*mirrorEpoch.Int64() + 1); i++ {
		commit(backend)
	}

	_, err = theGame.mirrorRegistrar.CollectDeposit(opts)

	if err != nil {
		t.Fatalf("CollectDeposit: expected no error, got %v", err)
	}
	commit(backend)

	balanceAfter, err = theGame.sampleToken.BalanceOf(&bind.CallOpts{}, serviceAddr)
	if err != nil {
		t.Fatalf("BalanceOf : expected no error, got %v", err)
	}

	if balanceBefore.Int64() == balanceAfter.Int64() {
		t.Fatalf("balance should change because vesting period passed")
	}

}

func TestRegisterAndNewCase(t *testing.T) {
	backend := newTestBackend()

	theGame := deployTheGame(t, backend)

	deposit(t, backend, theGame.sampleToken, theGame.mirrorRegistrar, depositVestingPeriod)

	opts := bind.NewKeyedTransactor(serviceKey)
	_, err := theGame.mirrorRegistrar.Register(opts, clientAddr)
	if err != nil {
		t.Fatalf("Register: expected no error, got %v", err)
	}
	commit(backend)

	ensAddress, resolverAddr, ens, publicResolver, err := deployENS(ethKey, big.NewInt(0), backend, ethAddr)
	t.Log("ensAddress", ensAddress.String(), resolverAddr.String())
	if err != nil {
		t.Fatalf("deployENS: expected no error, got %v", err)
	}

	err = registerENSRecord(t, backend, serviceENSName, "123", ens, publicResolver, resolverAddr)
	if err != nil {

		t.Fatal("registerENSRecord service.game fail")
	}
	err = registerENSRecord(t, backend, clientENSName, "13", ens, publicResolver, resolverAddr)
	if err != nil {
		t.Fatal("registerENSRecord", clientENSName, "fail")
	}

	_, err = theGame.mirrorEns.SetENSAddress(opts, ensAddress)
	if err != nil {
		t.Fatal("SetENSAddress fail")
	}

	opts = bind.NewKeyedTransactor(clientKey)

	_, err = theGame.swearGame.NewCase(opts, serviceId)

	if err != nil {
		t.Fatalf("Register: expected no error, got %v", err)
	}
	commit(backend)

	caseid, err := theGame.swearGame.Ids(&bind.CallOpts{}, clientAddr, big.NewInt(0))
	t.Log("case", caseid)
	counter := 0
	for i := 0; i < 32; i++ {
		counter++
		if caseid[i] != 0 {
			break
		}
	}
	if counter == 32 {
		t.Fatal("expected case id not zero")
	}
}

func TestNewCaseNotRegister(t *testing.T) {
	backend := newTestBackend()

	theGame := deployTheGame(t, backend)

	_, _, _, _, err := deployENS(ethKey, big.NewInt(0), backend, ethAddr)

	if err != nil {
		t.Fatalf("deployENS: expected no error, got %v", err)
	}
	commit(backend)

	opts := bind.NewKeyedTransactor(clientKey)

	_, err = theGame.swearGame.NewCase(opts, serviceId)

	if err != nil {
		t.Fatalf("Register: expected no error, got %v", err)
	}
	caseid, err := theGame.swearGame.Ids(&bind.CallOpts{}, clientAddr, big.NewInt(0))
	for i := 0; i < 32; i++ {
		if caseid[i] != 0 {
			t.Fatalf("CaseId: expected 0, got %v", caseid)
		}
	}
}
func TestDoubleRegistration(t *testing.T) {
	backend := newTestBackend()

	theGame := deployTheGame(t, backend)

	deposit(t, backend, theGame.sampleToken, theGame.mirrorRegistrar, depositVestingPeriod)

	opts := bind.NewKeyedTransactor(serviceKey)
	_, err := theGame.mirrorRegistrar.Register(opts, clientAddr)
	if err != nil {
		t.Fatalf("Register: expected no error, got %v", err)
	}
	commit(backend)
	opts = bind.NewKeyedTransactor(serviceKey)
	_, err = theGame.mirrorRegistrar.Register(opts, clientAddr)
	if err != nil {
		t.Fatalf("Register: expected no error, got %v", err)
	}
	commit(backend)
	registerPlayerCounter, err := theGame.mirrorRegistrar.PlayerCount(&bind.CallOpts{})
	if err != nil {
		t.Fatalf("PlayerCount : expected no error, got %v", err)
	}
	t.Log("registerPlayerCounter", registerPlayerCounter)
	if registerPlayerCounter.Int64() == 2 {
		t.Fatalf("PlayerCount : should not be 2 because cannot allowed multiple registeration for the same client ")
	}

}

//client (not the owner of the contract) try to register
func TestRegisterFromClient(t *testing.T) {
	backend := newTestBackend()

	theGame := deployTheGame(t, backend)

	deposit(t, backend, theGame.sampleToken, theGame.mirrorRegistrar, depositVestingPeriod)

	opts := bind.NewKeyedTransactor(ethKey)
	_, err := theGame.mirrorRegistrar.Register(opts, clientAddr)
	if err != nil {
		t.Fatalf("Register: expected no error, got %v", err)
	}
	commit(backend)

	PlayerCount, err := theGame.mirrorRegistrar.PlayerCount(&bind.CallOpts{})
	if err != nil {
		t.Fatalf("PlayerCount : expected no error, got %v", err)
	}

	if PlayerCount.Int64() == 1 {
		t.Fatalf("PlayerCount : should not be 1 because client cannot register ")
	}
}

func TestDeposit(t *testing.T) {
	backend := newTestBackend()
	theGame := deployTheGame(t, backend)
	//trasfer 100 tokens to the Service
	opts := bind.NewKeyedTransactor(ethKey)
	_, err := theGame.sampleToken.Transfer(opts, serviceAddr, big.NewInt(100))
	if err != nil {
		t.Fatalf("trasfer tokens to service: expected no error, got %v", err)
	}
	commit(backend)
	opts = bind.NewKeyedTransactor(serviceKey)
	opts.Value = big.NewInt(serviceDeposit)
	_, err = theGame.mirrorRegistrar.Deposit(opts, big.NewInt(100))

	if err != nil {
		t.Fatalf("depost tokens to contract: expected no error, got %v", err)
	}
	commit(backend)

	deposit, err := theGame.mirrorRegistrar.Deposits(&bind.CallOpts{}, serviceAddr)
	if err != nil {
		t.Fatalf("AmountStaked : expected no error, got %v", err)
	}

	if deposit.DepositedAmount.Int64() != big.NewInt(serviceDeposit).Int64() {
		t.Fatalf("AmountStaked ", deposit.DepositedAmount.Int64(), "is not equal to the deposit amount", big.NewInt(serviceDeposit))
	}

	opts.Value = big.NewInt(serviceDeposit)
	_, err = theGame.mirrorRegistrar.Deposit(opts, big.NewInt(100))
	if err != nil {
		t.Fatalf("depost tokens to contract: expected no error, got %v", err)
	}
	commit(backend)
	deposit, err = theGame.mirrorRegistrar.Deposits(&bind.CallOpts{}, serviceAddr)
	if err != nil {
		t.Fatalf("AmountStaked : expected no error, got %v", err)
	}
	//test it accumulate the deposits
	if deposit.DepositedAmount.Int64() != big.NewInt(serviceDeposit).Int64()*2 {
		t.Fatalf("AmountStaked ", deposit.DepositedAmount.Int64(), "is not equal to the deposit amount", big.NewInt(serviceDeposit))
	}
}

func deployTheGame(t *testing.T, backend *backends.SimulatedBackend) *TheGame {

	sampleTokenAddress, sampleToken, err := deploySampleTokenContract(ethKey, big.NewInt(0), backend, big.NewInt(1000))
	if err != nil {
		t.Fatalf("deploy contract: expected no error, got %v", err)
	}

	mirrorContractAddress, mirrorEns, err := deployMirror(serviceKey, big.NewInt(0), backend)
	if err != nil {
		t.Fatalf("deploy contract: expected no error, got %v", err)
	}
	promiseValidatorContractAddress, promiseValidator, err := deployPromiseValidator(serviceKey, big.NewInt(0), backend)
	if err != nil {
		t.Fatalf("deploy contract: expected no error, got %v", err)
	}
	mirrorRulesContractAddress, mirrorRules, err := deployMirrorRules(serviceKey, big.NewInt(0), backend, promiseValidatorContractAddress, mirrorContractAddress)
	if err != nil {
		t.Fatalf("deploy contract: expected no error, got %v", err)
	}

	mirrorRegistrarContractAddress, mirrorRegistrar, err := deployMirrorRegistrar(serviceKey, big.NewInt(0), backend, mirrorRulesContractAddress, sampleTokenAddress)
	if err != nil {
		t.Fatalf("deploy contract: expected no error, got %v", err)
	}

	swearGameContractAddress, swearGame, err := deploySwear(serviceKey, big.NewInt(0), backend, mirrorRegistrarContractAddress, mirrorRulesContractAddress, big.NewInt(compensationAmount))
	if err != nil {
		t.Fatalf("deploy contract: expected no error, got %v", err)
	}
	theGame := TheGame{}
	theGame.mirrorEns = mirrorEns
	theGame.mirrorRules = mirrorRules
	theGame.swearGameContractAddress = swearGameContractAddress
	theGame.sampleToken = sampleToken
	theGame.promiseValidator = promiseValidator
	theGame.promiseValidatorAddress = promiseValidatorContractAddress
	theGame.swearGame = swearGame
	theGame.mirrorRegistrar = mirrorRegistrar

	return &theGame
}

func registerENSRecord(t *testing.T, backend *backends.SimulatedBackend, name, contentHash string, ens *contract.ENS, resolver *contract.PublicResolver, resolverAddr common.Address) error {

	label := strings.TrimSuffix(name, ".game")
	opts := bind.NewKeyedTransactor(ethKey)
	_, err := ens.SetSubnodeOwner(opts, Namehash("game"), SHA3(label), ethAddr)
	if err != nil {
		return err
	}
	commit(backend)
	_, err = ens.SetResolver(opts, Namehash(name), resolverAddr)
	if err != nil {
		return err
	}
	commit(backend)
	_, err = resolver.SetContent(opts, Namehash(name), common.HexToHash(contentHash))
	if err != nil {
		return err
	}
	commit(backend)

	return nil
}

func SHA3(s string) common.Hash {
	return crypto.Keccak256Hash([]byte(s))
}

func Namehash(name string) (node common.Hash) {
	if name != "" {
		parts := strings.Split(name, ".")
		for i := len(parts) - 1; i >= 0; i-- {
			label := SHA3(parts[i])
			node = crypto.Keccak256Hash(append(node[:], label[:]...))
		}
	}
	return
}

// Promise represents a promise from the service to serve the client up to a certain blocknumber.
type Promise struct {
	contract    common.Address // address of swearGame contract, needed to avoid cross-contract submission
	beneficiary common.Address // address of Beneficiary (client)
	blockNumber *big.Int       // Until which block number this promise is valid
	sig         []byte         // signature Sign(Keccak256(contract, beneficiary, blocknumber), prvKey)
}

//issuePromise creates a promise signed by the serive's private key .
//this signed promise could later be submitted by the client of the service as an evidence that the service promise to serve it
func issuePromise(prvKey *ecdsa.PrivateKey, beneficiary common.Address, blockNumber *big.Int, contractAddress common.Address) (promise *Promise, err error) {

	sig, err := crypto.Sign(sigHash(contractAddress, beneficiary, blockNumber), prvKey)
	if err == nil {
		promise = &Promise{
			contract:    contractAddress,
			beneficiary: beneficiary,
			blockNumber: blockNumber,
			sig:         sig,
		}
	}
	return promise, err
}

// data to sign: contract address, beneficiary, cumulative amount of funds ever sent
func sigHash(beneficiary common.Address, contractAddress common.Address, blockNumber *big.Int) []byte {
	blocknumber := blockNumber.Bytes()
	if len(blocknumber) > 32 {
		return nil
	}
	var blocknumber32 [32]byte
	copy(blocknumber32[32-len(blocknumber):32], blocknumber)
	input := append(beneficiary.Bytes(), contractAddress.Bytes()...)
	input = append(input, blocknumber32[:]...)
	return crypto.Keccak256(input)
}

// v/r/s representation of signature
func sig2vrs(sig []byte) (v byte, r, s [32]byte) {
	v = sig[64] + 27
	copy(r[:], sig[:32])
	copy(s[:], sig[32:64])
	return
}
