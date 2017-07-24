# swap-swear-and-swindle-suite

This repository include a basic abstract example of a courtroom contracts game.

## Courtroom structure
The courtroom suite includes :
 - an abstract generic contract "SwearGame" which in a sense conduct the trial: iterate over the specific game witnesses ,trial statuses and verdict accordingly .
 - a specific game rules contract which defined the trial rules such as transitions states,grace periods per state..  .(in our example it is "MirrorRules")
 - a specific game witnesses contracts which in a sense are expert witnesses which can take decisions based on the submitted evidences.   


## Mirror game
A very simple basic game where the service offer to mirror the client ENS("client.game") content on its own ENS("service.game").

### game flow
-service will deposit some fund in the swear contact to ensure the service in the case of a claim

-client get a signed promise from the service to mirror its ens during the next X blocks.

-client update the client.game ENS.

If after X blocks the service ENS does not "mirror" its client ENS ..the client can submit a new case for the swear contract.

As an evident the client submit to court the signed promise it got from the service and the ENSNameHashes for both client and service.

If the case is valid a refund + compensation will send to the client from the contract.

### witnesses
 - PromiseValidator - validate a signed promise for the case it is submitted and it is ask for testimony .
 - MirrorENS        - resolve and compare the 2 enss for the case it is submitted and it is ask for testimony .


### repository structure and files
 /contracts/
 -courtroom.sol - SwearGame contract is SwearGameAbstract

 -mirrorens.sol - MirrorENS contract is WitnessAbstract

 -mirrorrules.sol - MirrorRules contract is TrialRulesAbstract

 -promisevalidator.sol -PromiseValidator contract is WitnessAbstract

 -Owned.sol -  Owned contract

 -sampletoken.sol -  SampleToke is StandardToken

 -standardtoken.sol -StandardToken is Token (erc20)

 -courtroom_test.go - go contract functionality tests

 /contracts/abstracts - include all abstract contracts solidity files used by this project contracts

 /contracts/mirrorgame - include the automaticly generated contract go binding files
 
 /vendor     - vendor dir needed for the go tests.

## rebuild solidity files
 go generate
## test
go test ./contracts/ -v --run Test
