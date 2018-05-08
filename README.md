# sw3 contracts

Contracts for Swap, Swear and Swindle.

## Tests

To run the tests first install `truffle`, then run `npm install` to set everything up.
Then run the tests with `truffle test`.

## Deviations from the sw3 paper

#### Swap

* hard deposits can be decreased without a payout to the owner. that way one can easily return the locked amount to the liquid balance.
* hard deposit decreases on payout. I assumes this was the intent anyway, but it's not explicitly stated in the paper.
* no support for soft deposits (not sure if there is anything even needed from the contract)
* index in promissory notes is ignored
* remark is not used in the invoice process

#### Swear

* does not use SWAP at the moment
* heavy on-chain footprint

#### Swindle

* anyone can start a case and is then treated as a swear contract
