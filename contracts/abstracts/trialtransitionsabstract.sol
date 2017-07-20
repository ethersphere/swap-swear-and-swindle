pragma solidity ^0.4.0;
import "./witnessabstract.sol";

contract TrialTransistionsAbstract {


  enum Status {UNCHALLENGED,GUILTY,NOT_GUILTY}

  function getStatus(uint8 witnessState,uint8 trialStatus) returns (uint8 status);
  function getWitness(uint8 trialStatus) returns (WitnessAbstract);
  function getInitialStatus() public returns (uint8 status);
  function getTrialExpiry() returns (uint expiery);
}
