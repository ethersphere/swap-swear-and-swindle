pragma solidity ^0.4.0;
import "./AbstractWitness.sol";

contract AbstractRules {

  uint8 constant public TRIAL_STATUS_UNCHALLENGED = 0;
  uint8 constant public TRIAL_STATUS_GUILTY = 1;
  uint8 constant public TRIAL_STATUS_NOT_GUILTY = 2;

  function nextStatus(AbstractWitness.TestimonyStatus witnessStatus,uint8 trialStatus) public view returns (uint8 status);
  function getWitness(uint8 trialStatus) public view returns (address witness, uint expiry);
  function getInitialStatus() public view returns (uint8 status);

  /* get minimal deposit for service */
  function getDeposit() public view returns (uint deposit);

  /* get minimal epoch for service */
  function getEpoch() public view returns (uint epoch);

}
