pragma solidity ^0.4.0;
import "./abstracts/trialtransitionsabstract.sol";

contract MirrorTransistions is TrialTransistionsAbstract {


uint8 constant MIRROR_CHALLENGE       = uint8(TrialTransistionsAbstract.Status.NOT_GUILTY) +1;
uint8 constant VALID_MIRROR_CHALLENGE = uint8(TrialTransistionsAbstract.Status.NOT_GUILTY) +2;

uint constant MIRROR_CHALLENGE_GRACE_PERIOD = 35;//Grace period set to 35 blocks to submit evident.
uint constant VALID_MIRROR_CHALLENGE_GRACE_PERIOD = 35;

mapping(uint8 => address) public witnesses;
mapping(uint8 => mapping(uint8 => uint8)) public transitions;
mapping(uint8 => uint) public gracePeriods;
//map caseId to map status to time elapse
mapping(bytes32 => mapping(uint8 => uint)) gracePeriodStartTime;

  function MirrorTransistions(address paymentValidatorContract,address ENSMirrotValidatorContract){

    witnesses[MIRROR_CHALLENGE]       = paymentValidatorContract;
    witnesses[VALID_MIRROR_CHALLENGE] = ENSMirrotValidatorContract;

    transitions[uint8(WitnessAbstract.Status.VALID)][MIRROR_CHALLENGE]         = VALID_MIRROR_CHALLENGE;
    transitions[uint8(WitnessAbstract.Status.INVALID)][MIRROR_CHALLENGE]       = uint8(TrialTransistionsAbstract.Status.NOT_GUILTY);
    transitions[uint8(WitnessAbstract.Status.VALID)][VALID_MIRROR_CHALLENGE]   = uint8(TrialTransistionsAbstract.Status.GUILTY);
    transitions[uint8(WitnessAbstract.Status.INVALID)][VALID_MIRROR_CHALLENGE] = uint8(TrialTransistionsAbstract.Status.NOT_GUILTY);

    gracePeriods[MIRROR_CHALLENGE] = MIRROR_CHALLENGE_GRACE_PERIOD;
    gracePeriods[VALID_MIRROR_CHALLENGE] = VALID_MIRROR_CHALLENGE_GRACE_PERIOD;
  }
  function getStatus(uint8 witnessState,uint8 trialStatus) returns (uint8 status){
    return transitions[witnessState][trialStatus];
  }

  function getInitialStatus() public returns (uint8 status){
    return MIRROR_CHALLENGE;
  }

  function getWitness(uint8 trialStatus) returns (WitnessAbstract){

    return WitnessAbstract(witnesses[trialStatus]);
  }

  function startGracePeriod(bytes32 caseId,uint8 status ) returns (bool){
    if (gracePeriodStartTime[caseId][status] == 0){
      gracePeriodStartTime[caseId][status] = block.number;
      return true;
    }
    return false;//already started.
  }

  function expired(bytes32 caseId,uint8 status ) returns (bool){
    if (gracePeriodStartTime[caseId][status]!=0){
       if ((block.number - gracePeriodStartTime[caseId][status])> gracePeriods[status]) return true;
    }
    return false;
  }

}
