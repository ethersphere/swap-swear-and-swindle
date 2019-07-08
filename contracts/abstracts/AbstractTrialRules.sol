pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;
import "./AbstractWitness.sol";
import "./AbstractConstants.sol";

/// @title AbstractRules - Swindle Trial Interface
contract AbstractTrialRules is AbstractConstants {
  
  /// @notice return initial status for a trial
  function getInitialStatus(bytes memory payload) public pure returns (uint8 status, bytes memory trialData);
  
  function updateData(AbstractWitness.TestimonyStatus witnessStatus, uint8 status, bytes memory trialData, bytes memory roles) 
  public pure returns (bytes memory);

  /// @notice return next status based on current status and outcome
  function nextStatus(AbstractWitness.TestimonyStatus witnessStatus,uint8 trialStatus) public pure returns (uint8 status);  

  /// @notice return witness for a given status
  function getWitness(uint8 trialStatus) 
  public view returns (address witness, uint expiry);
  
  function getWitnessPayload(uint8 trialStatus, address payable provider, address payable plaintiff, bytes memory trialData) 
  public view returns (bytes memory specification);  
}
