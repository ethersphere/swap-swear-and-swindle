pragma solidity ^0.4.23;

/// @title AbstractSwear - the sw3 swear interface
contract AbstractSwear {

  /// @notice callback for swindle when compensation should take place
  /// @param commitmentHash commitment to compensate from
  /// @param beneficiary beneficiary to compensate
  /// @param reward amount to be compensated
  function compensate(bytes32 commitmentHash, address beneficiary, uint reward) public;

  /// @notice callback for swindle at the end of the trial
  /// @param commitmentHash commitment
  function notifyTrialEnd(bytes32 commitmentHash) public;

}
