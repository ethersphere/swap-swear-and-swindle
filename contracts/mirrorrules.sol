pragma solidity ^0.4.0;
import "./abstracts/trialrulesabstract.sol";

contract MirrorRules is TrialRulesAbstract {


uint8 constant MIRROR_CHALLENGE       = uint8(TrialRulesAbstract.Status.NOT_GUILTY) +1;
uint8 constant VALID_MIRROR_CHALLENGE = uint8(TrialRulesAbstract.Status.NOT_GUILTY) +2;

uint constant MIRROR_CHALLENGE_GRACE_PERIOD = 35;//Grace period set to 35 blocks to submit evident.
uint constant VALID_MIRROR_CHALLENGE_GRACE_PERIOD = 35;
uint constant REWARD = 5; //plaintiff reward for the case of a valid case

mapping(uint8 => address) public witnesses;
mapping(uint8 => mapping(uint8 => uint8)) public transitions;
mapping(uint8 => uint) public gracePeriods;
//map caseId to map status to time elapse
mapping(bytes32 => mapping(uint8 => uint)) gracePeriodStartTime;


  function MirrorRules(address paymentValidatorContract,address ENSMirrotValidatorContract){

    witnesses[MIRROR_CHALLENGE]       = paymentValidatorContract;
    witnesses[VALID_MIRROR_CHALLENGE] = ENSMirrotValidatorContract;

    transitions[uint8(WitnessAbstract.Status.VALID)][MIRROR_CHALLENGE]         = VALID_MIRROR_CHALLENGE;
    transitions[uint8(WitnessAbstract.Status.INVALID)][MIRROR_CHALLENGE]       = uint8(TrialRulesAbstract.Status.NOT_GUILTY);
    transitions[uint8(WitnessAbstract.Status.VALID)][VALID_MIRROR_CHALLENGE]   = uint8(TrialRulesAbstract.Status.GUILTY);
    transitions[uint8(WitnessAbstract.Status.INVALID)][VALID_MIRROR_CHALLENGE] = uint8(TrialRulesAbstract.Status.NOT_GUILTY);

    gracePeriods[MIRROR_CHALLENGE] = MIRROR_CHALLENGE_GRACE_PERIOD;
    gracePeriods[VALID_MIRROR_CHALLENGE] = VALID_MIRROR_CHALLENGE_GRACE_PERIOD;
  }
  /// @notice getStatus - get next trial status according to witness state and the current trial state
  ///
  /// @param witnessStatus witness status (VALID , INVALID,PENDING)
  /// @param trialStatus current trial status
  /// @return status - next trial status - can be also GUILTY or NOT GUILTY.
  function getStatus(uint8 witnessStatus,uint8 trialStatus) returns (uint8 status){
    return transitions[witnessStatus][trialStatus];
  }
  /// @notice getInitialStatus - get initial trial status
  ///
  /// @return status -
  function getInitialStatus() public returns (uint8 status){
    return MIRROR_CHALLENGE;
  }
  /// @notice getWitness - get witness according to the trial status
  ///
  /// @param trialStatus current trial status
  /// @return WitnessAbstract - return a witness contract instance
  function getWitness(uint8 trialStatus) returns (WitnessAbstract){

    return WitnessAbstract(witnesses[trialStatus]);
  }
  /// @notice startGracePeriod - start counting for a grace period for a certain case and status.
  ///
  /// @return bool - true if it actually start counting for the grace period
  ///                false -if the grace period already started
  function startGracePeriod(bytes32 caseId,uint8 status ) returns (bool){
    if (gracePeriodStartTime[caseId][status] == 0){
      gracePeriodStartTime[caseId][status] = block.number;
      return true;
    }
    return false;//already started.
  }
  /// @notice expired - check expiration for a certain case and trial status
  ///
  /// @return bool - true if expiered otherwise false
  function expired(bytes32 caseId,uint8 status ) returns (bool){
    if (gracePeriodStartTime[caseId][status]!=0){
       if ((block.number - gracePeriodStartTime[caseId][status])> gracePeriods[status]) return true;
    }
    return false;
  }

  /// @notice getReward - return the reward for a valid case
  ///
  /// @return reward - the reward for a valid case
  function getReward() constant returns (uint reward){
    return REWARD;
  }

}
