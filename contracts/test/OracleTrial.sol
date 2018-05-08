pragma solidity ^0.4.19;
import "../abstracts/AbstractRules.sol";
import "./OracleWitness.sol";

contract OracleTrial is AbstractRules {
  uint8 constant TRIAL_STATUS_WITNESS_1 = 3;
  uint8 constant TRIAL_STATUS_WITNESS_2 = 4;

  address public witness1 = new OracleWitness();
  address public witness2 = new OracleWitness();

  function nextStatus(AbstractWitness.TestimonyStatus witnessStatus, uint8 trialStatus)
  public view returns (uint8 status) {
    require(witnessStatus == AbstractWitness.TestimonyStatus.VALID
      || witnessStatus == AbstractWitness.TestimonyStatus.INVALID);

    if(trialStatus == TRIAL_STATUS_WITNESS_1) {
      if(witnessStatus == AbstractWitness.TestimonyStatus.VALID) return TRIAL_STATUS_WITNESS_2;
      else return TRIAL_STATUS_NOT_GUILTY;
    } else if(trialStatus == TRIAL_STATUS_WITNESS_2){
      if(witnessStatus == AbstractWitness.TestimonyStatus.VALID) return TRIAL_STATUS_GUILTY;
      else return TRIAL_STATUS_NOT_GUILTY;
    } else revert();
  }

  function getWitness(uint8 trialStatus)
  public view returns (address, uint) {
    if(trialStatus == TRIAL_STATUS_WITNESS_1) return  (witness1, 2 days);
    if(trialStatus == TRIAL_STATUS_WITNESS_2) return  (witness2, 2 days);
    revert();
  }

  function getInitialStatus()
  public view returns (uint8 status) {
    return TRIAL_STATUS_WITNESS_1;
  }

  /* get minimal deposit for service */
  function getDeposit() public view returns (uint deposit) {
    return 100;
  }

  /* get minimal epoch for service */
  function getEpoch() public view returns (uint epoch) {
    return 30 days;
  }
}
