package contracts

//go:generate abigen --sol ./courtroom.sol --pkg contracts --out ./courtroom.go
//go:generate abigen --sol ./mirror.sol --pkg mirrorens --out ./mirror/mirrorens/mirror.go
//go:generate abigen --sol ./promisevalidator.sol --pkg promisevalidator --out ./mirror/promisevalidator/promisevalidator.go
//go:generate abigen --sol ./mirrortransitions.sol --pkg mirrortransition --out ./mirror/mirrortransition/mirrortransitions.go

import (
	"crypto/ecdsa"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/ens/contract"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	mirrorens "github.com/jaakmusic/swap-swear-and-swindle/contracts/mirror/mirrorens"
	mirrortransition "github.com/jaakmusic/swap-swear-and-swindle/contracts/mirror/mirrortransition"
	promisevalidator "github.com/jaakmusic/swap-swear-and-swindle/contracts/mirror/promisevalidator"
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
	PromiseTillNextBlocks = 5
	serviceDeposit        = int64(50)
	witnessesNumber       = int64(2)
	serviceId             = [32]byte{1, 2, 3, 4}
)

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

func deployCaseContract(prvKey *ecdsa.PrivateKey, amount *big.Int, backend *backends.SimulatedBackend) (common.Address, error) {
	deployTransactor := bind.NewKeyedTransactor(prvKey)
	deployTransactor.Value = amount

	addr, _, _, err := DeployCaseContract(deployTransactor, backend)

	if err != nil {
		return common.Address{}, err
	}
	commit(backend)
	return addr, nil
}

func deployMirror(prvKey *ecdsa.PrivateKey, amount *big.Int, backend *backends.SimulatedBackend) (common.Address, *mirrorens.Mirror, error) {
	deployTransactor := bind.NewKeyedTransactor(prvKey)
	deployTransactor.Value = amount

	addr, _, mirror, err := mirrorens.DeployMirror(deployTransactor, backend)
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

func deployMirrorTransitions(prvKey *ecdsa.PrivateKey, amount *big.Int, backend *backends.SimulatedBackend, paymentValidatorContract common.Address, ENSMirrotValidatorContract common.Address) (common.Address, error) {
	deployTransactor := bind.NewKeyedTransactor(prvKey)
	deployTransactor.Value = amount

	addr, _, _, err := mirrortransition.DeployMirrorTransistions(deployTransactor, backend, paymentValidatorContract, ENSMirrotValidatorContract)
	if err != nil {
		return common.Address{}, err
	}
	commit(backend)
	return addr, nil
}

func deploySwearGame(prvKey *ecdsa.PrivateKey, amount *big.Int, backend *backends.SimulatedBackend, caseContractAddr common.Address, tokenContractAddr common.Address, mirrorTransistions common.Address, rewordCompansation *big.Int) (common.Address, *SwearGame, error) {
	deployTransactor := bind.NewKeyedTransactor(prvKey)
	deployTransactor.Value = amount

	addr, _, swearGame, err := DeploySwearGame(deployTransactor, backend, caseContractAddr, tokenContractAddr, mirrorTransistions, rewordCompansation)
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

func deposit(t *testing.T, backend *backends.SimulatedBackend, sampleToken *SampleToken, swearGame *SwearGame) {
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
	_, err = swearGame.SwearGameTransactor.contract.Transfer(opts)
	if err != nil {
		t.Fatalf("depost tokens to contract: expected no error, got %v", err)
	}
	commit(backend)
}
func openClaimForMirrorGame(t *testing.T, clientContent string, serviceContent string, numberOfBlocksToWait int, pendingTest bool) {
	backend := newTestBackend()

	_, swearGame, sampleToken, promiseValidator, promiseValidatorAddress, _ := deployTheGame(t, backend)

	deposit(t, backend, sampleToken, swearGame)

	opts := bind.NewKeyedTransactor(serviceKey)
	_, err := swearGame.Register(opts, clientAddr)
	if err != nil {
		t.Fatalf("Register: expected no error, got %v", err)
	}
	commit(backend)

	ensAddress, resolverAddr, ens, publicResolver, err := deployENS(ethKey, big.NewInt(0), backend, ethAddr)
	t.Log("ensAddress", ensAddress.String(), resolverAddr.String())
	if err != nil {
		t.Fatalf("deployENS: expected no error, got %v", err)
	}

	err = registerENSRecord(t, backend, "service.game", serviceContent, ens, publicResolver, resolverAddr)
	if err != nil {

		t.Fatal("registerENSRecord service.game fail")
	}
	err = registerENSRecord(t, backend, "client.game", clientContent, ens, publicResolver, resolverAddr)
	if err != nil {
		t.Fatal("registerENSRecord", "client.game", "fail")
	}
	depositBefore, err := swearGame.Deposit(&bind.CallOpts{})
	if err != nil {
		t.Fatalf("AmountStaked : expected no error, got %v", err)
	}
	t.Log("deposit", depositBefore)
	balanceOfClientBefore, err := sampleToken.BalanceOf(&bind.CallOpts{}, clientAddr)
	if err != nil {
		t.Fatalf("balanceOfClientBefore : expected no error, got %v", err)
	}
	opts = bind.NewKeyedTransactor(clientKey)

	if !pendingTest {
		promise, err := issuePromise(serviceKey, clientAddr, big.NewInt(int64(blockNumber+PromiseTillNextBlocks)), promiseValidatorAddress)
		if err != nil {
			t.Fatalf("NewCase: issuePromise expected no error, got %v", err)
		}
		v, r, s := sig2vrs(promise.sig)

		promiseValidator.SubmitPromise(opts, serviceId, promise.beneficiary, promise.blockNumber, v, r, s)
	}

	for i := 0; i < numberOfBlocksToWait; i++ {
		commit(backend)
	}
	_, err = swearGame.NewCase(opts, serviceId)

	if err != nil {
		t.Fatalf("NewCase: expected no error, got %v", err)
	}
	commit(backend)

	claimid, err := swearGame.Ids(&bind.CallOpts{}, clientAddr, big.NewInt(0))
	t.Log("claim", claimid)

	status, err := swearGame.GetStatus(&bind.CallOpts{}, claimid)
	if err != nil {
		t.Fatalf("NewCase: expected no error, got %v", err)
	}
	t.Log("status", status)
	if pendingTest && status == 3 {
		//submit promise and restart from current state
		promise, err := issuePromise(serviceKey, clientAddr, big.NewInt(int64(blockNumber+PromiseTillNextBlocks)), promiseValidatorAddress)
		if err != nil {
			t.Fatalf("NewCase: issuePromise expected no error, got %v", err)
		}
		v, r, s := sig2vrs(promise.sig)

		promiseValidator.SubmitPromise(opts, serviceId, promise.beneficiary, promise.blockNumber, v, r, s)
		_, err = swearGame.ResumeCase(opts, claimid)

		if err != nil {
			t.Fatalf("NewCase: expected no error, got %v", err)
		}
		commit(backend)
	}

	depositAfter, err := swearGame.Deposit(&bind.CallOpts{})
	if err != nil {
		t.Fatalf("AmountStaked : expected no error, got %v", err)
	}
	balanceOfClientAfter, err := sampleToken.BalanceOf(&bind.CallOpts{}, clientAddr)
	if err != nil {
		t.Fatalf("balanceOfClientBefore : expected no error, got %v", err)
	}
	//If the client submit a claim a valid claim (ens are not equal) and the claim is within the time period the service promised
	//so it is a valid claim
	if (clientContent != serviceContent) && (numberOfBlocksToWait <= PromiseTillNextBlocks) {
		if depositBefore.Int64()-depositAfter.Int64() != int64(compensationAmount) {
			t.Fatalf("After a valid claim proccess : deposit at the contract should reduce by %d  ( got %d)", compensationAmount, (depositBefore.Int64() - depositAfter.Int64()))
		}
		if balanceOfClientAfter.Int64()-balanceOfClientBefore.Int64() != int64(compensationAmount) {
			t.Fatalf("After a valid claim proccess : the balance of client should increase by %d   ( got %d)", compensationAmount, (balanceOfClientAfter.Int64() - balanceOfClientBefore.Int64()))
		}
	} else {
		if depositBefore.Int64() != depositAfter.Int64() {
			t.Fatalf("non valid claim proccess : deposit at the contract should be the same as before  ( got %d)", (depositBefore.Int64() - depositAfter.Int64()))
		}
		if balanceOfClientAfter.Int64() != balanceOfClientBefore.Int64() {
			t.Fatalf("non valid  claim proccess should end by no change to the client balance   ( got %d)", (balanceOfClientAfter.Int64() - balanceOfClientBefore.Int64()))
		}
	}

}
func TestPromiseOk(t *testing.T) {
	//client and service content are diffrent  but number of block to wait is 6
	// the service promise to serve client for the next PromiseTillNextBlocks(5) blocks.
	//This test will fail if client will  get compensated for its claim.
	openClaimForMirrorGame(t, "1234", "4567", 6, false)
}

func TestOpenValidClaimPending(t *testing.T) {
	//client and service content are diffrent and number of block to wait is 0
	//This test will fail if client will not get compensated for its claim.
	openClaimForMirrorGame(t, "1234", "4567", 0, true)
}

func TestOpenValidClaim(t *testing.T) {
	//client and service content are diffrent and number of block to wait is 0
	//This test will fail if client will not get compensated for its claim.
	openClaimForMirrorGame(t, "1234", "4567", 0, false)
}

func TestOpenNoneValidClaim(t *testing.T) {
	//client and service content are the same and number of blocks to wait is 0
	//This test will fail if client will get compensated for its claim.
	openClaimForMirrorGame(t, "1234", "1234", 0, false)
}

func TestRegisterAndNewCase(t *testing.T) {
	backend := newTestBackend()

	_, swearGame, sampleToken, _, _, mirrorEns := deployTheGame(t, backend)

	deposit(t, backend, sampleToken, swearGame)

	opts := bind.NewKeyedTransactor(serviceKey)
	_, err := swearGame.Register(opts, clientAddr)
	if err != nil {
		t.Fatalf("Register: expected no error, got %v", err)
	}
	commit(backend)

	ensAddress, resolverAddr, ens, publicResolver, err := deployENS(ethKey, big.NewInt(0), backend, ethAddr)
	t.Log("ensAddress", ensAddress.String(), resolverAddr.String())
	if err != nil {
		t.Fatalf("deployENS: expected no error, got %v", err)
	}

	err = registerENSRecord(t, backend, "service.game", "123", ens, publicResolver, resolverAddr)
	if err != nil {

		t.Fatal("registerENSRecord service.game fail")
	}
	err = registerENSRecord(t, backend, "client.game", "13", ens, publicResolver, resolverAddr)
	if err != nil {
		t.Fatal("registerENSRecord", "client.game", "fail")
	}

	_, err = mirrorEns.SetENSAddress(opts, ensAddress)
	if err != nil {
		t.Fatal("SetENSAddress fail")
	}

	opts = bind.NewKeyedTransactor(clientKey)
	// promise, err := issuePromise(serviceKey, clientAddr, big.NewInt(int64(blockNumber+5)), swearGameContractAddress)
	// if err != nil {
	// 	t.Fatalf("NewCase: issuePromise expected no error, got %v", err)
	// }
	// v, r, s := sig2vrs(promise.Sig)
	_, err = swearGame.NewCase(opts, serviceId)

	if err != nil {
		t.Fatalf("Register: expected no error, got %v", err)
	}
	commit(backend)

	claimid, err := swearGame.Ids(&bind.CallOpts{}, clientAddr, big.NewInt(0))
	t.Log("claim", claimid)
	counter := 0
	for i := 0; i < 32; i++ {
		counter++
		if claimid[i] != 0 {
			break
		}
	}
	if counter == 32 {
		t.Fatal("expected claim id not zero")
	}
}

func TestNewCaseNotRegister(t *testing.T) {
	backend := newTestBackend()

	_, swearGame, _, _, _, _ := deployTheGame(t, backend)

	_, _, _, _, err := deployENS(ethKey, big.NewInt(0), backend, ethAddr)

	if err != nil {
		t.Fatalf("deployENS: expected no error, got %v", err)
	}
	commit(backend)

	opts := bind.NewKeyedTransactor(clientKey)
	// promise, err := issuePromise(serviceKey, clientAddr, big.NewInt(int64(blockNumber+5)), swearGameContractAddress)
	// if err != nil {
	// 	t.Fatalf("NewCase: issuePromise expected no error, got %v", err)
	// }
	// v, r, s := sig2vrs(promise.Sig)
	_, err = swearGame.NewCase(opts, serviceId)

	if err != nil {
		t.Fatalf("Register: expected no error, got %v", err)
	}
	claimid, err := swearGame.Ids(&bind.CallOpts{}, clientAddr, big.NewInt(0))
	for i := 0; i < 32; i++ {
		if claimid[i] != 0 {
			t.Fatalf("ClaimId: expected 0, got %v", claimid)
		}
	}
}
func TestDoubleRegistration(t *testing.T) {
	backend := newTestBackend()

	_, swearGame, sampleToken, _, _, _ := deployTheGame(t, backend)

	deposit(t, backend, sampleToken, swearGame)

	opts := bind.NewKeyedTransactor(serviceKey)
	_, err := swearGame.Register(opts, clientAddr)
	if err != nil {
		t.Fatalf("Register: expected no error, got %v", err)
	}
	commit(backend)
	opts = bind.NewKeyedTransactor(serviceKey)
	_, err = swearGame.Register(opts, clientAddr)
	if err != nil {
		t.Fatalf("Register: expected no error, got %v", err)
	}
	commit(backend)
	registerPlayerCounter, err := swearGame.PlayerCount(&bind.CallOpts{})
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

	_, swearGame, sampleToken, _, _, _ := deployTheGame(t, backend)

	deposit(t, backend, sampleToken, swearGame)

	opts := bind.NewKeyedTransactor(ethKey)
	_, err := swearGame.Register(opts, clientAddr)
	if err != nil {
		t.Fatalf("Register: expected no error, got %v", err)
	}
	commit(backend)

	PlayerCount, err := swearGame.PlayerCount(&bind.CallOpts{})
	if err != nil {
		t.Fatalf("PlayerCount : expected no error, got %v", err)
	}

	if PlayerCount.Int64() == 1 {
		t.Fatalf("PlayerCount : should not be 1 because client cannot register ")
	}
}

func TestDeposit(t *testing.T) {
	backend := newTestBackend()
	_, swearGame, sampleToken, _, _, _ := deployTheGame(t, backend)
	//trasfer 100 tokens to the Service
	opts := bind.NewKeyedTransactor(ethKey)
	_, err := sampleToken.Transfer(opts, serviceAddr, big.NewInt(100))
	if err != nil {
		t.Fatalf("trasfer tokens to service: expected no error, got %v", err)
	}
	commit(backend)
	opts = bind.NewKeyedTransactor(serviceKey)
	opts.Value = big.NewInt(serviceDeposit)
	_, err = swearGame.SwearGameTransactor.contract.Transfer(opts)

	if err != nil {
		t.Fatalf("depost tokens to contract: expected no error, got %v", err)
	}
	commit(backend)

	deposit, err := swearGame.Deposit(&bind.CallOpts{})
	if err != nil {
		t.Fatalf("AmountStaked : expected no error, got %v", err)
	}

	if deposit.Int64() != big.NewInt(serviceDeposit).Int64() {
		t.Fatalf("AmountStaked ", deposit.Int64(), "is not equal to the deposit amount", big.NewInt(serviceDeposit))
	}

	opts.Value = big.NewInt(serviceDeposit)
	_, err = swearGame.SwearGameTransactor.contract.Transfer(opts)
	if err != nil {
		t.Fatalf("depost tokens to contract: expected no error, got %v", err)
	}
	commit(backend)
	deposit, err = swearGame.Deposit(&bind.CallOpts{})
	if err != nil {
		t.Fatalf("AmountStaked : expected no error, got %v", err)
	}
	//test it accumulate the deposits
	if deposit.Int64() != big.NewInt(serviceDeposit).Int64()*2 {
		t.Fatalf("AmountStaked ", deposit.Int64(), "is not equal to the deposit amount", big.NewInt(serviceDeposit))
	}
	t.Log("amountstaked", deposit)

}

func deployTheGame(t *testing.T, backend *backends.SimulatedBackend) (
	swearGameContractAddress common.Address,
	swearGame *SwearGame,
	sampleToken *SampleToken,
	promiseValidator *promisevalidator.PromiseValidator,
	promiseValidatorAddress common.Address,
	mirrorEns *mirrorens.Mirror) {

	sampleTokenAddress, sampleToken, err := deploySampleTokenContract(ethKey, big.NewInt(0), backend, big.NewInt(1000))
	if err != nil {
		t.Fatalf("deploy contract: expected no error, got %v", err)
	}

	caseContractAddress, err := deployCaseContract(serviceKey, big.NewInt(0), backend)
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
	mirrorFlowContractAddress, err := deployMirrorTransitions(serviceKey, big.NewInt(0), backend, promiseValidatorContractAddress, mirrorContractAddress)
	if err != nil {
		t.Fatalf("deploy contract: expected no error, got %v", err)
	}

	swearGameContractAddress, swearGame, err = deploySwearGame(serviceKey, big.NewInt(0), backend, caseContractAddress, sampleTokenAddress, mirrorFlowContractAddress, big.NewInt(compensationAmount))
	if err != nil {
		t.Fatalf("deploy contract: expected no error, got %v", err)
	}
	t.Log("address", sampleTokenAddress, caseContractAddress)
	return swearGameContractAddress, swearGame, sampleToken, promiseValidator, promiseValidatorContractAddress, mirrorEns
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
//this signed promise could later be submited by the client of the service as an evident that the service promise to serve it
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
