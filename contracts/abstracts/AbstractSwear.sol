pragma solidity ^0.5.0;

/// @title AbstractSwear - the sw3 swear interface
contract AbstractSwear {

  /// @notice callback for swindle when compensation should take place
  /// @param commitmentHash commitment to compensate from
  /// @param beneficiary beneficiary to compensate
  /// @param reward amount to be compensated
  function compensate(bytes32 commitmentHash, address payable beneficiary, uint reward) public;

  /// @notice callback for swindle at the end of the trial
  /// @param commitmentHash commitment
  function notifyTrialEnd(bytes32 commitmentHash) public;

}
