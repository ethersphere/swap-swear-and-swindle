# swap-swear-and-swindle-suite

This repository includes a basic example of a courtroom contracts suite.

## Courtroom structure

The courtroom suite includes an abstract generic [SwearGame contract](#sweargame-contract) , a specific [trial rules contract](#trial-rules-contract) and a specific [witnesses](#witness-contract) contracts.

### SwearGame contract

This is the main general swear contract which conducts the trial.

It calls the specific [trial rules contract](#trial-rules-contract) to get the [witnesses](#witness-contract),get the trial statuses and transitions
and to check for expiry time for a specific case and trial status.

It calls the specific game witnesses contracts to testimony and validates for a submitted evidence.

The SwearGame contract is aligned with the ABIs defined at SwearGameAbstract (courtroomabstract.sol).

/// @notice () - a payable function which is used by the service as a deposit functionality

///The function without name is the default function that is called when anyone sends funds to a contract

function () payable;

/// @notice register - register a player to the game

/// The function will throw if the player is already registered or there is not

/// enough deposit in the contract to make sure the player could be compensated for the

/// case of a valid case.

/// @param _player  - the player address

/// @return bool registered - true for success registration.

function register(address _player) onlyOwner public returns (bool registered);

/// @notice leaveGame - dismiss a player from the game (unregister)

/// only allow plaintiff which does not have open cases submitted on its behalf to leave game

/// @param _player  - the player address

function leaveGame(address _player);

/// @notice getStatus - return the trial status of a case
///

/// @param id  - case id

/// @return  status  - the status of a case

function getStatus(bytes32 id) public constant returns (uint8 status);

/// @notice newCase - open a new case for a service id

/// the function require that the msg sender is already registered to the game.

/// @param serviceId  - the accuse service id  

/// @return bool - true for successful operation.

function newCase(bytes32 serviceId) public returns (bool);

/// @notice trial - start or restart a trial process for a certain case
///
/// the function require that the case is a valid one.

/// @param id  - case id

/// @return bool - true for successful operation.

function trial(bytes32 id) public returns (bool);

### Witness contract

 A smart contract which responsible to testimony according to a certain type of evidence which is submitted to it.

 Testimony means it is doing a validation for the evidence submitted to it.  

 The witness contract is specific to the game which is offered by the service,
 which means that each service will have its own witnesses contracts.

 The witness contract is accessed by the main SwearGame contract and is aligned with  
 the ABIs defined at witnessAbstract.sol contract:

 /// @notice testimonyFor - ask for testimony for a specific case, service, and client

 ///
 /// @param caseId case id

 /// @param serviceId the service id which the witness is requested to testimony for.

 /// @param clientAddress client address

 /// @return Status { VALID,INVALID, PENDING}

 function testimonyFor(bytes32 caseId,bytes32 serviceId,address clientAddress) returns (Status);

 /// @notice isEvidenceSubmitted - check if evidence was submitted for a specific case ,service and client

 ///

 /// @param caseId case id

 /// @param serviceId the service id which the evidence was submitted for

 /// @param clientAddress client address

 /// @return bool - true or false

 function isEvidenceSubmitted(bytes32 caseId, bytes32 serviceId,address clientAddress) returns (bool);

### Trial rules contract

A contract which defines a specific set of rules for a game:

- The trial statuses transitions map. A scheme which defines the transition from a certain status to the next one.
- Grace periods for each trial status.
- The reward which will be transferred to the plaintiff as a compensation for the case of a valid claim.

While the main SwearGame contract conducts a trial to resolve a specific case it iterates between different trial statuses.

This TrialRules contract is aligned with TrialRulesAbstract.sol  and should implement the following ABIs:

/// @notice getStatus - get next trial status according to witness state and the current trial state

/// @param witnessStatus witness status (VALID , INVALID,PENDING)

/// @param trialStatus current trial status

/// @return status - next trial status - can be also GUILTY or NOT GUILTY.

function getStatus(uint8 witnessStatus,uint8 trialStatus) returns (uint8 status);

/// @notice getWitness - get witness according to the trial status

/// @param trialStatus current trial status

/// @return WitnessAbstract - return a witness contract instance

function getWitness(uint8 trialStatus) returns (WitnessAbstract);

/// @notice getInitialStatus - get first trial status

/// @return status -

function getInitialStatus() public returns (uint8 status);

/// @notice expired - check the expiration period of a certain case and trial status.

/// @return bool - true if expired otherwise false.

function expired(bytes32 caseId,uint8 status) returns (bool);

/// @notice startGracePeriod - start counting for a grace period for a certain case and status.

/// @return bool - true if it actually start counting for the grace period

///                false if the grace period already started

function startGracePeriod(bytes32 caseId,uint8 status) returns (bool);

/// @notice getReward - return the reward for a valid case

///

/// @return reward - the reward for a valid case

function getReward()returns (uint reward);



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
