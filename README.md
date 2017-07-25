# swap-swear-and-swindle-suite

This repository include a basic abstract example of a courtroom contracts game.


## Courtroom structure
The courtroom suite includes :
 - an abstract generic contract "SwearGame" which in a sense conduct the trial: iterate over the specific game witnesses ,trial statuses and verdict accordingly .
 - a specific game rules contract which defined the trial rules such as transitions statuses,grace periods per state..  .(in our example it is "MirrorRules")
 - A specific game witnesses contracts which in a sense are "expert witnesses" which can validate submitted evidences.  

### SwearGame contract

This is the main general swear contract which conduct the trial.

It calls the specific TrialRules contract to get the witnesses ,get the trial statuses and transitions
and to check for expiry time for a specific case and trial status.

It calls the specific game witnesses contracts to testimony and validate for a submitted evidence.

The SwearGame contract is align with the ABIs defined at SwearGameAbstract (courtroomabstract.sol).

/// @notice () - a payable function which is used by the service as a deposit functionality

///The function without name is the default function that is called whenever anyone sends funds to a contract

function () payable;

/// @notice register - register a player to the game

/// The function will throw if the player is already register or there is not

/// enough deposit in the contract to ensure the player could be compensated for the

/// case of a valid case.

/// @param _player  - the player address

/// @return bool registered - true for success registration.

function register(address _player) onlyOwner public returns (bool registered);

/// @notice leaveGame - dismiss a player from the game (unregister)

/// allow only plaintiff which do not have openCases on it name to leave game

/// @param _player  - the player address

function leaveGame(address _player);

/// @notice getStatus - return the trial status of a case
///

/// @param id  - case id

/// @return  status  - the status of a case

function getStatus(bytes32 id) public constant returns (uint8 status);

/// @notice newCase - open a new case for a service id

/// the function require that the msg sender is already register to the game.

/// @param serviceId  - service id

/// @return bool - true for successful operation.

function newCase(bytes32 serviceId) public returns (bool);

/// @notice trial - initiate or restart a trial process for a certain case
///
/// the function require that the case is a valid one.

/// @param id  - case id

/// @return bool - true for successful operation.

function trial(bytes32 id) public returns (bool);

### Expert witness contract

 A smart contract which responsible to testimony for a pre defined
 evidence submitted to it.

 Testimony means it is doing a validation for the evidence submitted to it.  

 The witness contract is specific to the game which is offers by the service ,
 which means that each service will have its own witnesses contracts.

 The witness contract is access by the main SwearGame contract and is align with  
 the ABIs defined at witnessAbstract.sol contract:

 /// @notice testimonyFor - request for testimony for a specific case ,service and client

 ///
 /// @param caseId case id

 /// @param serviceId the service id which

 /// @param clientAddress client address

 /// @return Status { VALID,INVALID, PENDING}

 function testimonyFor(bytes32 caseId,bytes32 serviceId,address clientAddress) returns (Status);

 /// @notice isEvidenceSubmitted - check if an evidence was submitted for a specific case ,service and client

 ///

 /// @param caseId case id

 /// @param serviceId the service id which

 /// @param clientAddress client address

 /// @return bool - true or false

 function isEvidenceSubmitted(bytes32 caseId, bytes32 serviceId,address clientAddress) returns (bool);

### Trial statuses,transitions and grace periods

While the main SwearGame contract conduct a trial to resolve a specific case it iterate between different
trial statuses.

For each game there is a specific pre defined trial statuses and a specific trial transitions scheme which defined
the transition for a certain status to the next one.

The trial statuses and their transitions are defined per each game in a TrialRules contract.

The TrialRules contract also defined the grace period per each trial status.

This TrialRules contract is align with TrialRulesAbstract.sol  and should implement the following ABIs:

/// @notice getStatus - get next trial status according to witness state and the current trial state

/// @param witnessStatus witness status (VALID , INVALID,PENDING)

/// @param trialStatus current trial status

/// @return status - next trial status - can be also GUILTY or NOT GUILTY.

function getStatus(uint8 witnessStatus,uint8 trialStatus) returns (uint8 status);

/// @notice getWitness - get witness according to the trial status

/// @param trialStatus current trial status

/// @return WitnessAbstract - return a witness contract instance

function getWitness(uint8 trialStatus) returns (WitnessAbstract);

/// @notice getInitialStatus - get initial trial status

/// @return status -

function getInitialStatus() public returns (uint8 status);

/// @notice expired - check expiration for a certain case and trial status.

/// @return bool - true if expired otherwise false.

function expired(bytes32 caseId,uint8 status) returns (bool);

/// @notice startGracePeriod - start counting for a grace period for a certain case and status.

/// @return bool - true if it actually start counting for the grace period

///                false -if the grace period already started

function startGracePeriod(bytes32 caseId,uint8 status) returns (bool);

/// @notice getReward - return the reward for a valid case

///

/// @return reward - the reward for a valid case

function getReward()returns (uint reward);



## Mirror game
A very simple basic game where the service offers to mirror the client ENS("client.game") content on its own ENS("service.game").

### game flow
-service will deposit some funds in the swear contact to ensure the service in the case of a court case/litigation

-client receive a signed promise from the service to mirror its ens during the next X blocks.

-client update the client.game ENS.

If after X blocks the service ENS does not "mirror" its client ENS ..the client can submit a new case for the swear contract.

As an evidence the client submits to court the signed promise it got from the service and the ENSNameHashes for both client and service.

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

 -sampletoken.sol -  SampleToken is StandardToken

 -standardtoken.sol -StandardToken is Token (erc20)

 -courtroom_test.go - go contract functionality tests

 /contracts/abstracts - include all abstract contracts solidity files used by this project contracts

 /contracts/mirrorgame - include the automatically generated contract go binding files

 /vendor     - vendor dir needed for the go tests.

## rebuild solidity files
 go generate
## test
go test ./contracts/ -v --run Test
