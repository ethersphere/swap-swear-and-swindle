pragma solidity ^0.5.0;
import "../abstracts/AbstractRules.sol";
import "../abstracts/AbstractWitness.sol";

/// @title SimpleTrial - a simple test trial with a Witness as argument
contract SimpleTrial is AbstractRules {
  uint8 constant TRIAL_STATUS_WITNESS = 3;

  AbstractWitness public witness;

  constructor(AbstractWitness _witness) public {
    witness = _witness;
  }

  function nextStatus(AbstractWitness.TestimonyStatus witnessStatus, uint8 trialStatus)
  public view returns (uint8 status) {
    require(witnessStatus == AbstractWitness.TestimonyStatus.VALID
      || witnessStatus == AbstractWitness.TestimonyStatus.INVALID);

    if(trialStatus == TRIAL_STATUS_WITNESS) {
      if(witnessStatus == AbstractWitness.TestimonyStatus.VALID) return TRIAL_STATUS_NOT_GUILTY;
      else return TRIAL_STATUS_GUILTY;
    } else revert();
  }

  function getWitness(uint8 trialStatus)
  public view returns (address, uint) {
    if(trialStatus == TRIAL_STATUS_WITNESS) return  (address(witness), 2 days);
    revert();
  }

  function getInitialStatus()
  public view returns (uint8 status) {
    return TRIAL_STATUS_WITNESS;
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
