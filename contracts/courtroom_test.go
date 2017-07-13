// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package contracts

//go:generate abigen --sol ./courtroom.sol --pkg contracts --out ./courtroom.go

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
)

var (
	serviceKey, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291") //service
	ethKey, _     = crypto.HexToECDSA("8a1f9a8f95be41cd7ccb6168179afb4504aefe388d1e14474d32c45c72ce7b7a")
	clientKey, _  = crypto.HexToECDSA("49a7b37aa6f6645917e7b807e9d1c00d4fa71f18343b0d4122a4d2df64dd6fee")
	serviceAddr   = crypto.PubkeyToAddress(serviceKey.PublicKey)
	ethAddr       = crypto.PubkeyToAddress(ethKey.PublicKey)
	clientAddr    = crypto.PubkeyToAddress(clientKey.PublicKey)
)

func newTestBackend() *backends.SimulatedBackend {
	return backends.NewSimulatedBackend(core.GenesisAlloc{
		serviceAddr: {Balance: big.NewInt(1000000000)},
		ethAddr:     {Balance: big.NewInt(1000000000)},
		clientAddr:  {Balance: big.NewInt(1000000000)},
	})
}

func deployCaseContract(prvKey *ecdsa.PrivateKey, amount *big.Int, backend *backends.SimulatedBackend) (common.Address, error) {
	deployTransactor := bind.NewKeyedTransactor(prvKey)
	deployTransactor.Value = amount

	addr, _, _, err := DeployCaseContract(deployTransactor, backend)

	if err != nil {
		return common.Address{}, err
	}
	backend.Commit()
	return addr, nil
}

func deploySwearGame(prvKey *ecdsa.PrivateKey, amount *big.Int, backend *backends.SimulatedBackend, caseContractAddr common.Address, tokenContractAddr common.Address, rewordCompansation *big.Int) (common.Address, *SwearGame, error) {
	deployTransactor := bind.NewKeyedTransactor(prvKey)
	deployTransactor.Value = amount

	addr, _, swearGame, err := DeploySwearGame(deployTransactor, backend, caseContractAddr, tokenContractAddr, rewordCompansation)

	if err != nil {
		return common.Address{}, nil, err
	}
	backend.Commit()
	return addr, swearGame, nil
}

func deploySampleTokenContract(prvKey *ecdsa.PrivateKey, amount *big.Int, backend *backends.SimulatedBackend, initialSupply *big.Int) (common.Address, *SampleToken, error) {
	deployTransactor := bind.NewKeyedTransactor(prvKey)
	deployTransactor.Value = amount

	addr, _, sampleToken, err := DeploySampleToken(deployTransactor, backend, initialSupply)

	if err != nil {
		return common.Address{}, nil, err
	}
	backend.Commit()
	return addr, sampleToken, nil
}

func deployENS(prvKey *ecdsa.PrivateKey, amount *big.Int, backend *backends.SimulatedBackend, owner common.Address) (common.Address, common.Address, *contract.ENS, *contract.PublicResolver, error) {
	deployTransactor := bind.NewKeyedTransactor(prvKey)
	deployTransactor.Value = amount

	ensAddress, _, ens, err := contract.DeployENS(deployTransactor, backend, owner)

	if err != nil {
		return common.Address{}, common.Address{}, nil, nil, err
	}
	backend.Commit()

	addr, _, publicResolver, err := contract.DeployPublicResolver(deployTransactor, backend, ensAddress)
	if err != nil {
		return common.Address{}, common.Address{}, nil, nil, err
	}
	backend.Commit()

	//setting up .game domain

	_, err = ens.SetOwner(deployTransactor, common.Hash{}, ethAddr)

	if err != nil {
		return common.Address{}, common.Address{}, nil, nil, err
	}
	backend.Commit()
	_, err = ens.SetSubnodeOwner(deployTransactor, common.Hash{}, SHA3("game"), ethAddr)

	if err != nil {
		return common.Address{}, common.Address{}, nil, nil, err
	}
	backend.Commit()

	return ensAddress, addr, ens, publicResolver, nil
}

func deposit(t *testing.T, backend *backends.SimulatedBackend, sampleToken *SampleToken, swearGame *SwearGame) {
	//trasfer 100 tokens to the Service
	opts := bind.NewKeyedTransactor(ethKey)
	_, err := sampleToken.Transfer(opts, serviceAddr, big.NewInt(100))
	if err != nil {
		t.Fatalf("trasfer tokens to service: expected no error, got %v", err)
	}
	backend.Commit()

	//deposit 50 token to the contract
	opts = bind.NewKeyedTransactor(serviceKey)
	_, err = swearGame.Deposit(opts, big.NewInt(50))
	if err != nil {
		t.Fatalf("depost tokens to contract: expected no error, got %v", err)
	}
	backend.Commit()
}
func openClaimForReflectorGame(t *testing.T, clientContent string, serviceContent string) {
	backend := newTestBackend()

	_, swearGame, sampleToken := deployTheGame(t, backend)

	deposit(t, backend, sampleToken, swearGame)

	opts := bind.NewKeyedTransactor(serviceKey)
	_, err := swearGame.Register(opts, clientAddr)
	if err != nil {
		t.Fatalf("Register: expected no error, got %v", err)
	}
	backend.Commit()

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
	amountStakedBefore, err := swearGame.AmountStaked(&bind.CallOpts{})
	if err != nil {
		t.Fatalf("AmountStaked : expected no error, got %v", err)
	}
	balanceOfClientBefore, err := sampleToken.BalanceOf(&bind.CallOpts{}, clientAddr)
	if err != nil {
		t.Fatalf("balanceOfClientBefore : expected no error, got %v", err)
	}
	opts = bind.NewKeyedTransactor(clientKey)
	_, err = swearGame.OpenNewClaim(opts, [32]byte{1})
	if err != nil {
		t.Fatalf("OpenNewClaim: expected no error, got %v", err)
	}
	backend.Commit()
	claimid, err := swearGame.ClientsClaimsIds(&bind.CallOpts{}, clientAddr, big.NewInt(0))
	t.Log("claim", claimid)

	amountStakedAfter, err := swearGame.AmountStaked(&bind.CallOpts{})
	if err != nil {
		t.Fatalf("AmountStaked : expected no error, got %v", err)
	}
	balanceOfClientAfter, err := sampleToken.BalanceOf(&bind.CallOpts{}, clientAddr)
	if err != nil {
		t.Fatalf("balanceOfClientBefore : expected no error, got %v", err)
	}
	if clientContent != serviceContent {
		if amountStakedBefore.Int64()-amountStakedAfter.Int64() != 5 {
			t.Fatalf("After a valid claim proccess : amountStaked at the contract should reduce by 5  ( got %d)", (amountStakedBefore.Int64() - amountStakedAfter.Int64()))
		}
		if balanceOfClientAfter.Int64()-balanceOfClientBefore.Int64() != 5 {
			t.Fatalf("After a valid claim proccess : the balance of client should increase by 5   ( got %d)", (balanceOfClientAfter.Int64() - balanceOfClientBefore.Int64()))
		}
	} else {
		if amountStakedBefore.Int64() != amountStakedAfter.Int64() {
			t.Fatalf("non valid claim proccess : amountStaked at the contract should reduce by 5  ( got %d)", (amountStakedBefore.Int64() - amountStakedAfter.Int64()))
		}
		if balanceOfClientAfter.Int64() != balanceOfClientBefore.Int64() {
			t.Fatalf("non valid  claim proccess should end by no change to the client balance   ( got %d)", (balanceOfClientAfter.Int64() - balanceOfClientBefore.Int64()))
		}
	}

}
func TestOpenValidClaim(t *testing.T) {
	//client and service content are diffrent
	openClaimForReflectorGame(t, "1234", "4567")
}

func TestOpenNoneValidClaim(t *testing.T) {
	//client and service content are the same
	openClaimForReflectorGame(t, "1234", "1234")
}
func TestRegisterAndOpenNewClaim(t *testing.T) {
	backend := newTestBackend()

	_, swearGame, sampleToken := deployTheGame(t, backend)

	deposit(t, backend, sampleToken, swearGame)

	opts := bind.NewKeyedTransactor(serviceKey)
	_, err := swearGame.Register(opts, clientAddr)
	if err != nil {
		t.Fatalf("Register: expected no error, got %v", err)
	}
	backend.Commit()

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

	_, err = swearGame.SetENSAddress(opts, ensAddress)
	if err != nil {
		t.Fatal("SetENSAddress fail")
	}

	opts = bind.NewKeyedTransactor(clientKey)
	_, err = swearGame.OpenNewClaim(opts, [32]byte{1})
	if err != nil {
		t.Fatalf("Register: expected no error, got %v", err)
	}
	backend.Commit()

	claimid, err := swearGame.ClientsClaimsIds(&bind.CallOpts{}, clientAddr, big.NewInt(0))
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

func TestOpenNewClaimNotRegister(t *testing.T) {
	backend := newTestBackend()

	_, swearGame, _ := deployTheGame(t, backend)

	_, _, _, _, err := deployENS(ethKey, big.NewInt(0), backend, ethAddr)

	if err != nil {
		t.Fatalf("deployENS: expected no error, got %v", err)
	}
	backend.Commit()

	opts := bind.NewKeyedTransactor(clientKey)
	_, err = swearGame.OpenNewClaim(opts, [32]byte{1})
	if err != nil {
		t.Fatalf("Register: expected no error, got %v", err)
	}
	backend.Commit()
	claimid, err := swearGame.ClientsClaimsIds(&bind.CallOpts{}, clientAddr, big.NewInt(0))
	for i := 0; i < 32; i++ {
		if claimid[i] != 0 {
			t.Fatalf("ClaimId: expected 0, got %v", claimid)
		}
	}
}
func TestDoubleRegistration(t *testing.T) {
	backend := newTestBackend()

	_, swearGame, sampleToken := deployTheGame(t, backend)

	deposit(t, backend, sampleToken, swearGame)

	opts := bind.NewKeyedTransactor(serviceKey)
	_, err := swearGame.Register(opts, clientAddr)
	if err != nil {
		t.Fatalf("Register: expected no error, got %v", err)
	}
	backend.Commit()
	opts = bind.NewKeyedTransactor(serviceKey)
	_, err = swearGame.Register(opts, clientAddr)
	if err != nil {
		t.Fatalf("Register: expected no error, got %v", err)
	}
	backend.Commit()
	registerPlayerCounter, err := swearGame.RegisteredPlayersCounter(&bind.CallOpts{})
	if err != nil {
		t.Fatalf("RegisteredPlayersCounter : expected no error, got %v", err)
	}
	t.Log("registerPlayerCounter", registerPlayerCounter)
	if registerPlayerCounter.Int64() == 2 {
		t.Fatalf("RegisteredPlayersCounter : should not be 2 because cannot allowed multiple registeration for the same client ")
	}

}

//client (not the owner of the contract) try to register
func TestRegisterFromClient(t *testing.T) {
	backend := newTestBackend()

	_, swearGame, sampleToken := deployTheGame(t, backend)

	deposit(t, backend, sampleToken, swearGame)

	opts := bind.NewKeyedTransactor(ethKey)
	_, err := swearGame.Register(opts, clientAddr)
	if err != nil {
		t.Fatalf("Register: expected no error, got %v", err)
	}
	backend.Commit()

	registerPlayerCounter, err := swearGame.RegisteredPlayersCounter(&bind.CallOpts{})
	if err != nil {
		t.Fatalf("RegisteredPlayersCounter : expected no error, got %v", err)
	}
	t.Log("registerPlayerCounter", registerPlayerCounter)
	if registerPlayerCounter.Int64() == 1 {
		t.Fatalf("RegisteredPlayersCounter : should not be 1 because client cannot register ")
	}
}
func TestRegisterAndLeave(t *testing.T) {
	backend := newTestBackend()
	_, swearGame, sampleToken := deployTheGame(t, backend)
	opts := bind.NewKeyedTransactor(serviceKey)
	_, err := swearGame.Register(opts, clientAddr)
	if err != nil {
		t.Fatalf("Register: expected no error, got %v", err)
	}
	backend.Commit()

	registerPlayerCounter, err := swearGame.RegisteredPlayersCounter(&bind.CallOpts{})
	if err != nil {
		t.Fatalf("RegisteredPlayersCounter : expected no error, got %v", err)
	}
	t.Log("registerPlayerCounter", registerPlayerCounter)
	if registerPlayerCounter.Int64() == 1 {
		t.Fatalf("RegisteredPlayersCounter : should not be 1 because there is no amount staked yet")
	}
	deposit(t, backend, sampleToken, swearGame)
	//try to register again
	opts = bind.NewKeyedTransactor(serviceKey)
	_, err = swearGame.Register(opts, clientAddr)
	if err != nil {
		t.Fatalf("Register: expected no error, got %v", err)
	}
	backend.Commit()

	registerPlayerCounter, err = swearGame.RegisteredPlayersCounter(&bind.CallOpts{})
	if err != nil {
		t.Fatalf("RegisteredPlayersCounter : expected no error, got %v", err)
	}
	t.Log("registerPlayerCounter", registerPlayerCounter)
	if registerPlayerCounter.Int64() != 1 {
		t.Fatalf("RegisteredPlayersCounter : should  be 1 but got ", registerPlayerCounter.Int64())
	}
	//try to register with the same address again
	_, err = swearGame.Register(opts, clientAddr)
	if err != nil {
		t.Fatalf("Register: expected no error, got %v", err)
	}
	backend.Commit()

	registerPlayerCounter, err = swearGame.RegisteredPlayersCounter(&bind.CallOpts{})
	if err != nil {
		t.Fatalf("RegisteredPlayersCounter : expected no error, got %v", err)
	}
	t.Log("registerPlayerCounter", registerPlayerCounter)
	if registerPlayerCounter.Int64() != 1 {
		t.Fatalf("RegisteredPlayersCounter : should  be 1 but got ", registerPlayerCounter.Int64())
	}

	_, err = swearGame.LeaveGame(opts, clientAddr)
	if err != nil {
		t.Fatalf("LeaveGame : expected no error, got %v", err)
	}
	backend.Commit()
	registerPlayerCounter, err = swearGame.RegisteredPlayersCounter(&bind.CallOpts{})
	if err != nil {
		t.Fatalf("RegisteredPlayersCounter : expected no error, got %v", err)
	}
	t.Log("registerPlayerCounter", registerPlayerCounter)
	if registerPlayerCounter.Int64() != 0 {
		t.Fatalf("RegisteredPlayersCounter : should  be 0 but got ", registerPlayerCounter.Int64())
	}
}

func TestDeposit(t *testing.T) {
	backend := newTestBackend()
	_, swearGame, sampleToken := deployTheGame(t, backend)
	//trasfer 100 tokens to the Service
	opts := bind.NewKeyedTransactor(ethKey)
	_, err := sampleToken.Transfer(opts, serviceAddr, big.NewInt(100))
	if err != nil {
		t.Fatalf("trasfer tokens to service: expected no error, got %v", err)
	}
	backend.Commit()
	opts = bind.NewKeyedTransactor(serviceKey)

	_, err = swearGame.Deposit(opts, big.NewInt(50))
	if err != nil {
		t.Fatalf("depost tokens to contract: expected no error, got %v", err)
	}
	backend.Commit()

	amountStaked, err := swearGame.AmountStaked(&bind.CallOpts{})
	if err != nil {
		t.Fatalf("AmountStaked : expected no error, got %v", err)
	}

	if amountStaked.Int64() != big.NewInt(50).Int64() {
		t.Fatalf("AmountStaked ", amountStaked.Int64(), "is not equal to the deposit amount", big.NewInt(50))
	}

	_, err = swearGame.Deposit(opts, big.NewInt(50))
	if err != nil {
		t.Fatalf("depost tokens to contract: expected no error, got %v", err)
	}
	backend.Commit()
	amountStaked, err = swearGame.AmountStaked(&bind.CallOpts{})
	if err != nil {
		t.Fatalf("AmountStaked : expected no error, got %v", err)
	}
	//test it accumulate the deposits
	if amountStaked.Int64() != big.NewInt(50).Int64()*2 {
		t.Fatalf("AmountStaked ", amountStaked.Int64(), "is not equal to the deposit amount", big.NewInt(50))
	}
	t.Log("amountstaked", amountStaked)

}

func deployTheGame(t *testing.T, backend *backends.SimulatedBackend) (swearGameContractAddress common.Address, swearGame *SwearGame, sampleToken *SampleToken) {

	sampleTokenAddress, sampleToken, err := deploySampleTokenContract(ethKey, big.NewInt(0), backend, big.NewInt(1000))
	if err != nil {
		t.Fatalf("deploy contract: expected no error, got %v", err)
	}

	caseContractAddress, err := deployCaseContract(serviceKey, big.NewInt(0), backend)
	if err != nil {
		t.Fatalf("deploy contract: expected no error, got %v", err)
	}

	swearGameContractAddress, swearGame, err = deploySwearGame(serviceKey, big.NewInt(0), backend, caseContractAddress, sampleTokenAddress, big.NewInt(5))
	if err != nil {
		t.Fatalf("deploy contract: expected no error, got %v", err)
	}
	t.Log("address", sampleTokenAddress, caseContractAddress)
	return swearGameContractAddress, swearGame, sampleToken
}

func registerENSRecord(t *testing.T, backend *backends.SimulatedBackend, name, contentHash string, ens *contract.ENS, resolver *contract.PublicResolver, resolverAddr common.Address) error {

	label := strings.TrimSuffix(name, ".game")
	opts := bind.NewKeyedTransactor(ethKey)
	_, err := ens.SetSubnodeOwner(opts, Namehash("game"), SHA3(label), ethAddr)
	if err != nil {
		return err
	}
	backend.Commit()
	_, err = ens.SetResolver(opts, Namehash(name), resolverAddr)
	if err != nil {
		return err
	}
	backend.Commit()
	_, err = resolver.SetContent(opts, Namehash(name), common.HexToHash(contentHash))
	if err != nil {
		return err
	}
	backend.Commit()

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

// Cheque represents a payment promise to a single beneficiary.
type Promise struct {
	Contract    common.Address // address of swearGame contract, needed to avoid cross-contract submission
	Beneficiary common.Address // address of Beneficiary (client)
	BlockNumber *big.Int       // Until which block number this promise is valid
	Sig         []byte         // signature Sign(Keccak256(contract, beneficiary, blocknumber), prvKey)
}

//Issue creates a promise signed by the serive's private key .
//this signed promise could later be submited by the client of the service as an evident that the service promise to serve it
func Issue(prvKey *ecdsa.PrivateKey, beneficiary common.Address, blockNumber *big.Int, contractAddress common.Address) (promise *Promise, err error) {

	sig, err := crypto.Sign(sigHash(beneficiary, contractAddress, blockNumber), prvKey)
	if err == nil {
		promise = &Promise{
			Contract:    contractAddress,
			Beneficiary: beneficiary,
			BlockNumber: blockNumber,
			Sig:         sig,
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
