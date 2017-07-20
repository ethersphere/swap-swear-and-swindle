pragma solidity ^0.4.0;
import "./abstracts/trialtransitionsabstract.sol";

contract MirrorTransistions is TrialTransistionsAbstract {


uint8 constant MIRROR_CHALLENGE       = uint8(TrialTransistionsAbstract.Status.NOT_GUILTY) +1;
uint8 constant VALID_MIRROR_CHALLENGE = uint8(TrialTransistionsAbstract.Status.NOT_GUILTY) +2;

uint constant GRACE_PERIOD = 50;//Trial can last 50 blocks from first case submition.

mapping(uint8 => address) public witnesses;
mapping(uint8 => mapping(uint8 => uint8)) public transitions;


  function MirrorTransistions(address paymentValidatorContract,address ENSMirrotValidatorContract){

    witnesses[MIRROR_CHALLENGE] = paymentValidatorContract;
    witnesses[VALID_MIRROR_CHALLENGE] = ENSMirrotValidatorContract;

    transitions[uint8(WitnessAbstract.Status.VALID)][MIRROR_CHALLENGE] = VALID_MIRROR_CHALLENGE;
    transitions[uint8(WitnessAbstract.Status.INVALID)][MIRROR_CHALLENGE] = uint8(TrialTransistionsAbstract.Status.NOT_GUILTY);
    transitions[uint8(WitnessAbstract.Status.VALID)][VALID_MIRROR_CHALLENGE] = uint8(TrialTransistionsAbstract.Status.GUILTY);
    transitions[uint8(WitnessAbstract.Status.INVALID)][VALID_MIRROR_CHALLENGE] = uint8(TrialTransistionsAbstract.Status.NOT_GUILTY);

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

  function getTrialExpiry() returns (uint expiery){
    return GRACE_PERIOD;
  }

}
