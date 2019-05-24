pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;
import "./AbstractWitness.sol";
import "./AbstractConstants.sol";

/// @title AbstractRules - Swindle Trial Interface
contract AbstractTrialRules is AbstractConstants {
  
  /// @notice return initial status for a trial
  function getInitialStatus() public view returns (uint8 status);
  
  function initialize(bytes32 caseId, bytes memory payload) public;
  function setRoles(bytes32 caseId, AbstractWitness.TestimonyStatus witnessStatus, uint8 status, bytes memory roles) public;

  /// @notice return next status based on current status and outcome
  function nextStatus(AbstractWitness.TestimonyStatus witnessStatus,uint8 trialStatus) public view returns (uint8 status);  

  /// @notice return witness for a given status
  function getWitness(uint8 trialStatus) 
  public view returns (address witness, uint expiry);
  
  function getWitnessPayload(uint8 trialStatus, bytes32 caseId, address payable provider, address payable plaintiff) 
  public view returns (bytes memory specification);  
}
