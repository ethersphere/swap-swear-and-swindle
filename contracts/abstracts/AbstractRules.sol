pragma solidity ^0.4.0;
import "./AbstractWitness.sol";
import "./AbstractConstants.sol";

/// @title AbstractRules - Swindle Trial Interface
contract AbstractRules is AbstractConstants {

  /// @notice return next status based on current status and outcome
  function nextStatus(AbstractWitness.TestimonyStatus witnessStatus,uint8 trialStatus) public view returns (uint8 status);
  /// @notice return witness for a given status
  function getWitness(uint8 trialStatus) public view returns (address witness, uint expiry);
  /// @notice return initial status for a trial
  function getInitialStatus() public view returns (uint8 status);

  /// @notice get minimal deposit for service
  function getDeposit() public view returns (uint deposit);

  /// @notice get minimal epoch for service
  function getEpoch() public view returns (uint epoch);

}
