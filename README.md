# sw3 contracts

Contracts for Swap, Swear and Swindle.

**Please note that all contracts within this repository are considered highly experimental, contain critical flaws and will cause loss of money if used in production. Also Swear / Swindle are pure experimentation at this time and will probably replaced completely.**

## Tests

To run the tests first install `truffle`, then run `npm install` to set everything up.
Then run the tests with `truffle test`.

## Overview

## Swap

Swap is a chequebook-style contract with support for promissory notes and hard deposits.

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
* `witness` (if not 0) testimony returns VALID

If the `beneficiary` is 0, it is set to the submitter.
Part of the note might bounce, just as with cheques (`_payout`).

#### Invoices

Fulfilled notes can be subsumed in the cheque mechanism. Instead of cashing the note, the beneficiary sends an invoice (a signed note containing the noteId and the amount and serial of the last cheque). The owner is then supposed to sign a cheque with the amount `lastCheque.amount + note.amount` and serial `serial + 1`. If the beneficiary does not receive this cheque the note can still be submitted using `submitNote` (and not accept any further cheques until the situation is resolved). If the beneficiary does this even though the cheque was sent, the owner can cancel the payout within the security delay by presenting the correct cheque (`submitPaidInvoice`). If the recorded serial is still lower, the cheque will also count as submitted.

## Swear



## Swindle



## Tests



#### OracleTrial



#### SimpleTrial with HashWitness



## Deviations from the (unreleased) sw3 paper

#### Swap

* hard deposits can be decreased without a payout to the owner. that way one can easily return the locked amount to the liquid balance.
* hard deposit decreases on payout. I assumes this was the intent anyway, but it's not explicitly stated in the paper.
* no support for soft deposits (not sure if there is anything even needed from the contract)
* index in promissory notes is ignored
* remark is not used in the invoice process

#### Swear

* can work on-chain
* can work with Swap (but contains ugly hacks)
* many other deviations

#### Swindle

* anyone can start a case and is then treated as a swear contract
