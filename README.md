# swap-swear-and-swindle-suite

This repository includes a basic example of a courtroom contracts suite.

The courtroom concept is based on [swarm swap swear and swindle](http://swarm-gateways.net/bzz:/theswarm.eth/ethersphere/orange-papers/1/sw%5E3.pdf)

## Courtroom structure

The courtroom suite includes an abstract generic [SwearGame contract](#sweargame-contract) , a specific [trial rules contract](#trial-rules-contract) , a specific [witnesses](#witness-contract) contracts and a specific [registrar](#trial-registrar).

### SwearGame contract

This is the main general swear contract which conducts the trial.

It calls the specific [trial rules contract](#trial-rules-contract) to get the [witnesses](#witness-contract),get the trial statuses and transitions
and to check for expiry time for a specific case and trial status.

It calls the specific game witnesses contracts to testimony and validates for a submitted evidence.

The SwearGame contract is aligned with the ABIs defined at [SwearGameAbstract](contracts/abstracts/courtroomAbstract.sol).

### Witness contract

 A smart contract which responsible to testimony according to a certain type of evidence which is submitted to it.

 Testimony means it is doing a validation for the evidence submitted to it.  

 The witness contract is specific to the game which is offered by the service,
 which means that each service will have its own witnesses contracts.

 The witness contract is accessed by the main SwearGame contract and is aligned with  
 the ABIs defined at [witnessAbstract](contracts/abstracts/witnessAbstract.sol) contract.


### Trial rules contract

A contract which defines a specific set of rules for a game:

- The trial statuses transitions map. A scheme which defines the transition from a certain status to the next one.
- Grace periods for each trial status.
- The reward which will be transferred to the plaintiff as a compensation for the case of a valid claim.

While the main SwearGame contract conducts a trial to resolve a specific case it iterates between different trial statuses.

This TrialRules contract is aligned with [TrialRulesAbstract](contracts/abstracts/trialrulesabstract.sol) contract.

### Trial registrar

A contract which manages registrations and deposits for the game players.
A player who wishes to register and/or deposit to the swear game should do that via this contract.
The registrar contract is accessed by the main SwearGame contract and is aligned with  
the ABIs defined at [RegistrarAbstract](contracts/abstracts/registrarabstract.sol) contract.


## Mirror game
A very simple basic game where the service offers to mirror the client ENS("client.game") content on its own ENS("service.game").

### game flow
-service will deposit some funds in the swear contact to make sure the service in the case of a court case/litigation

-client receives a signed promise from the service to mirror its ENS during the next X blocks.

-client updates the client.game ENS.

If after X blocks the service ENS does not "mirror" its client ENS, the client can submit a new case for the swear contract.

As evidence,the client submits to court the signed promise it got from the service and the ENSNameHashes for both client and service.

If the case is valid, a refund + compensation will be sent to the client from the contract.

### witnesses
 - PromiseValidator - testify that a signed promise which was submitted to it is a valid one.
 - MirrorENS        - resolves and compares two ENS entries, to check for the presence of a hash.


### repository structure and files
 /contracts/
 -courtroom.sol - SwearGame contract is SwearGameAbstract

 -mirrorens.sol - MirrorENS contract is WitnessAbstract

 -mirrorrules.sol - MirrorRules contract is TrialRulesAbstract

 -promisevalidator.sol -PromiseValidator contract is WitnessAbstract

 -mirrorregistrar.sol - MirrorRegistrar is RegistrarAbstract

 -Owned.sol -  Owned contract

 -sampletoken.sol -  SampleToken is StandardToken

 -standardtoken.sol -StandardToken is Token (ERC20)

 -courtroom_test.go - go contract functionality tests

 /contracts/abstracts - include all abstract contracts solidity files used by this project contracts

 /contracts/mirrorgame - include the automatically generated contract go binding files

 /vendor     - vendor dir needed for the go tests.

## rebuild solidity files
 go generate
## test
go test ./contracts/ -v --run Test
