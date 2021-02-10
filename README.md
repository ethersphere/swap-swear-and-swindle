# sw3 contracts

Contracts for Swap, Swear and Swindle.

**Please note that all contracts within this repository are considered highly experimental and may cause loss of funds if used in production.**

The `master` branch only contains the `SimpleSwap` contract for now. Everything else (full `Swap`, `Swear` and `Swindle`) can be found in the `experimental` branch.

## Tests

This is a regular truffle project. You can either use `truffle test` (if installed) or `npm test` which will use a locally installed truffle.

```sh
npm install
npm test
```

To also generate coverage information use `npm run coverage` instead.

## Linting

This repo currently uses `solhint` as linter. It can be called through npm:
```sh
npm run solhint
```

## Go-bindings

To generate go bindings use
```sh
npm run abigen
```

This will generate the bindings in the `bindings/` directory. Suitable versions of `solc` and `abigen` have to be installed for this to work.
Alternatively this can also be done through docker:

```sh
docker build -t sw3 -f Dockerfile.abigen .
docker run -v $(pwd)/bindings:/sw3/bindings sw3 npm run abigen
```

In addition to the file from `abigen` this will also generate a go file that includes the runtime bytecode.

The go bindings are automatically generated on every release git tag and the pushed to the [ethersphere/go-sw3](https://github.com/ethersphere/go-sw3) repository.

## Test Deployments

Swap is deployed on the following testnets:

| Network | Factory | Token |
| ------- | -------- | ------- |
| Goerli | [0xf0277caffea72734853b834afc9892461ea18474](https://goerli.etherscan.io/address/0xf0277caffea72734853b834afc9892461ea18474) | [0x2aC3c1d3e24b45c6C310534Bc2Dd84B5ed576335](https://goerli.etherscan.io/address/0x2aC3c1d3e24b45c6C310534Bc2Dd84B5ed576335) |

## Overview

### SimpleSwap

`SimpleSwap` is a chequebook-style contract with support hard deposits.

#### Cheques

The issuer can issue cheques. Those are a signed piece of data containing the swap address, a beneficiary, and a cumulative amount. If the beneficiary wishes to cash the cheque it needs to be sent to the contract using the `cashChequeBeneficiary` function. Alternatively anyone else can also cash the cheque using `cashCheque` provided they have a signature from the beneficiary. In that case they might also get a portion of the payout as a fee. If there is not enough liquid balance in the contract part of the cheque might bounce (whether the full cheque was paid out should be verified through the `ChequeCashed` event). Later the beneficiary can try cashing again to get the remaining amount.

#### Hard Deposit

The issuer can lock a certain amount of the balance to a specific beneficiary (`increaseHardDeposit`) to give solvency guarantees up to the amount of the hard deposit. Decreasing a hard deposit is a two step process: First it needs to be prepared (`prepareDecreaseHardDeposit`) to start a security delay, which the beneficiary can use to cash the outstanding cheques. Afterwards the deposit can be decreased (`decreaseHardDeposit`).
The balance not covered by hard deposits can be withdrawn by the issuer at any time.

### Data Formats

#### Signatures

For signing purposes the chequebook uses EIP-712 Ethereum typed structured data hashing and signing. The `EIP712Domain` domain name is `Chequebook` in version `1.0`. The `chainId` field is used, the `verifyingContract` and `salt` fields are omitted.

#### Cheques

| Field            | Type         | Description   |
| ---------------- | ------------ | ------------- |
| swap             | address      | the swap contract this is for |
| beneficiary      | address      | beneficiary of the cheque |
| cumulativePayout | uint256      | cumulative amount |

#### CashOut

| Field            | Type         | Description   |
| ---------------- | ------------ | ------------- |
| swap             | address      | the swap contract this is for |
| sender           | address      | the address allowed to submit this cheque |
| requestPayout    | uint256      | the maximum amount that should be paid out |
| recipient        | address      | the target of the payment |
| calleePayout     | uint256      | amount of the payout that should go the caller |

#### CustomHardDepositDecreaseTimeout

| Field            | Type         | Description   |
| ---------------- | ------------ | ------------- |
| swap             | address      | the swap contract this is for |
| beneficiary      | address      | beneficiary of the hard deposit |
| decreaseTimeout  | uint256      | new timeout |

#### Factory

`SimpleSwap` has an accompanying factory contract `SimpleSwapFactory`. Deploying through the factory guarantees that the `SimpleSwap `contract was deployed with the correct init code.

## Swap

`Swap` is `SimpleSwap` extended with support for promissory notes. You can find it in the `experimental` branch.

## Swear and Swindle

`Swear` and `Swindle` are the contracts for the trial system of sw3. You can find them in the `experimental` branch.