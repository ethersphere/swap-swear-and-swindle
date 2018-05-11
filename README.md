# sw3 contracts

Contracts for Swap, Swear and Swindle.

**Please note that all contracts within this repository are considered highly experimental, contain critical flaws (missing checks, badly chosen timeouts, etc.) and will cause loss of money if used in production. Also Swear / Swindle are pure experimentation at this time and will probably replaced completely.**

## Tests

To run the tests first install `truffle`, then run `npm install` to set everything up.
Then run the tests with `truffle test`.

## Overview

## Swap

`Swap` is a chequebook-style contract with support for promissory notes and hard deposits.

#### Cheques

The owner can issue cheques. Those are a signed piece of data containing the swap address, a beneficiary, a cumulative amount and a serial number. If the beneficiary wishes to cash the cheque, it first needs to be submitted to the Swap (using `submitCheque`). After a security delay the cheque can then be cashed using `cashCheque`. If there is not enough liquid balance in the contract part of the cheque might bounce (see `_payout`). The beneficiary can later try cashing again to get the remaining amount.

During the security delay a newer cheque with a higher serial number and the same beneficiary can be presented by the beneficiary to replace the old one. The owner can also present a cheque but it needs to be signed by the beneficiary (`submitChequeLower`, this is needed in case there was a decrease in the amount). Replacing the cheque resets the security delay.

#### Hard Deposit

The owner can lock a certain amount of the balance to a specific beneficiary (`increaseHardDeposit`) to give some solvency guarantees. Decreasing a hard deposit is a two step process: First it needs to be prepared (`prepareDecreaseHardDeposit`) to start a security delay, which the beneficiary can use to cash the outstanding cheques. Afterwards the deposit can be decreased (`decreaseHardDeposit`).
The balance not covered by hard deposits can be withdrawn by the owner at any time.

#### Promissory Notes

Notes are a more advanced form of cheques (but without the replacement mechanism).
Notes have the following fields:

| Field        | Description   |
| ------------ | ------------- |
| index        | used as a nonce, must be non-zero |
| amount       | amount to be paid (or 0 if blank cheque) |
| beneficiary  | beneficiary of the note (or 0 if bounty) |
| witness      | witness used as an escrow condition (or 0 , see `Swindle` below) |
| validFrom    | earliest date the note can be cashed (or 0) |
| validUntil   | latest date the note can be cashed (or 0) |
| remark       | 32-bytes of arbitrary data (e.g. service data in `Swear`) |

Cashing notes is also a 2-step process. First they need to be submitted (`submitNote`) which starts a security delay (needed for invoices). Afterwards in can be cashed out (`cashNote`) in arbitrary small steps. If the amount is 0, arbitrary amounts can be withdrawn.

The note needs to be valid both at submission and cashing. This means:
* `validFrom` < now < `validUntil`
* `witness` (if not 0) testimony returns `VALID`

If the `beneficiary` is 0, it is set to the submitter.
Part of the note might bounce, just as with cheques (`_payout`).

#### Invoices

Fulfilled notes can be subsumed in the cheque mechanism. Instead of cashing the note, the beneficiary sends an invoice (a signed note containing the noteId and the amount and serial of the last cheque). The owner is then supposed to sign a cheque with the amount `lastCheque.amount + note.amount` and serial `serial + 1`. If the beneficiary does not receive this cheque the note can still be submitted using `submitNote` (and not accept any further cheques until the situation is resolved). If the beneficiary does this even though the cheque was sent, the owner can cancel the payout within the security delay by presenting the correct cheque (`submitPaidInvoice`). If the recorded serial is still lower, the cheque will also count as submitted.

## Swindle

`Swindle` keeps tracks of ongoing trials, calls the witnesses when necessary and notifies `Swear` of the outcome. It never handles any funds.

#### Witness

A `witness` is any contract supporting the following function:

`function testimonyFor(address owner, address beneficiary, bytes32 noteId) public view returns (TestimonyStatus)` where `TestimonyStatus` can be
* `VALID`: the submitted evidence was found to be valid
* `INVALID`: the submitted evidence was found to be invalid
* `PENDING`: no evidence has been submitted so far

For `Swap` owner is the contract owner, in `Swindle` it would be the service provider (who is also the `Swap` owner if the deposit is handled offchain). What needs to checked can be determined from the `noteId` which could be arbitrary data in case of onchain deposits or the hash of a `Swap` note (information can then be encoded in the `remark` field, see the `HashWitness` example).

#### Trial Rules

Trials are implemented as a finite state machine with state identified as `uint8`s (with values smaller than 3 being reserved for `NOT_STARTED`, `GUILTY`, `NOT_GUILTY`). The initial state is determined by the `getInitialStatus` function from the `rules` contract. The next state is determined by the state and witness outcome alone using `nextStatus` (which should be `pure`). For every state the rules provide an associated `Witness` (`getWitness`) and a maximum wait time for the witness.
The rules contract itself should ideally be stateless.

The rules also contain `getEpoch` and `getDeposit` to determine the required deposit and service duration. Both of those are ignored in case of offchain deposits (here it is the plaintiffs responsibility to check the `Swap` note before accepting the service to be guaranteed).

#### Trial Process

Anyone can start a trial by providing a service provider, a plaintiff, a `rules` (also called trial) contract, a `noteId` (case data for the witnesses) and a `commitmentHash` identifying the commitment with the caller (which is treated as a `Swear` contract). The trial is then identified by a `caseId`.

The trial can be advanced one state at a time using the `continueTrial` function. `Swindle` calls the next witness with the `noteId` and moves to the next state according to the `rules`. If the testimony is still `PENDING` nothing happens unless the witness has timed out in which case it is interpreted as `INVALID`.

Once a terminal state (`GUILTY` or `NOT_GUILTY`) is reached, the trial can be ended using `endTrial`, which informs the `Swear` contract about the outcome. In case of a `GUILTY` verdict Swindle will also notify `Swear` about the necessary compensation (which at the moment is always the entire deposit to the `plaintiff` or the note `beneficiary` in case of a `Swap` note).

## Swear

`Swear` can handle on-chain deposits for and initiate `Swindle` trials. It can also process `Swap` notes to initiate a trial.

#### Onchain service provisions

In the onchain case, the service provider puts a deposit into the `Swear` contract using `addCommitment` specifying the `noteId` (can be arbitrary 32-byte data), the `timeout` and the `rules`. The deposit size and the timeout need to conform with the `rules` for this to be accepted. The deposit can be withdraw again (`withdraw`) after the timeout if there are no open cases (and no trial was lost).

Trials can be started using `startTrial` which records the ongoing trial ( thus preventing withdraws) and instructs `Swindle` to start the trial.

At the end of the trial `notifyTrialEnd` will be called from `Swindle` which decreases the amount of ongoing trials (potentially unlocking withdraws). `Swindle` may also call `compensate` in which case `Swear` will decrease the deposit and send it to the `plaintiff`.

#### Offchain service provisions

In the offchain case, the service provider signs a `Swap` note with the `client` being the `beneficiary`, the `amount` as the security deposit, `Swear` as the `witness` and the service timeout as `validUntil`. If there is no dispute over the service the note will eventually expire and the provider is no longer at risk of losing the deposit. No onchain activity is required in this case (except potentially making sure everything is covered by hard deposits).

If the `client` is not satisfied with the service a trial can be started using `startTrialFromNote`. In this case `Swear` expects the `remark` field to equal `keccak256(rules, payload)` where `rules` is the address of the trial rules and payload can be arbitrary 32 bytes. `Swear` then records the commitment (with a special `note` flag) and instructs `Swindle` to start the trial (which proceeds in the same way as in the onchain case).

If `Swindle` calls compensate no payout takes place, instead `Swear` records that it should return `VALID` if being called as a `Witness` by the `Swap` with the corresponding `noteId`.

Starting multiple trials for the same note is currently broken. As hard deposits can currently be decreased to 0 within 2 days all offchain service provisions have de-facto no solvency guarantees.

## Tests

There are two test trials in this repository: `OracleTrial` and `SimpleTrial` (with `HashWitness`).

#### OracleTrial

In the `OracleTrial` there a two `OracleWitness`es. An `OracleWitness` is a `Witness` meant for testing whose testimony for some `noteId` can be set in advance. For the `OracleTrial` to reach a `GUILTY` verdict both of the witnesses need to return a `VALID` testimony. In all other cases the trial ends with `NOT_GUILTY`.

#### SimpleTrial with HashWitness

A `SimpleTrial` only has a single stage. The `witness` is passed in as a constructor argument.

In the `storage test` the witness is a `HashWitness`. A `HashWitness` expects `noteId` to be in the format used by offchain deposits. The `HashWitness` returns a `VALID` testimony if `data` was provided to the contract (where `keccak256(data) == payload`) otherwise it returns `PENDING`.

It can be seen as very primitive form of chunk insurance (and is compatible with Swarm POC-2 chunks).

## Deviations from the (unreleased) sw3 paper

#### Swap

* hard deposits can be decreased without a payout to the owner. that way one can easily return the locked amount to the liquid balance.
* hard deposit decreases on payout. I assumes this was the intent anyway, but it's not explicitly stated in the paper.
* no support for soft deposits (not sure if there is anything even needed from the contract)
* index in promissory notes is ignored
* remark is not used in the invoice process

#### Swear

* can work on-chain
* can work with Swap (but is insecure due to low hard deposit timeout)
* many other deviations

#### Swindle

* anyone can start a case and is then treated as a swear contract
