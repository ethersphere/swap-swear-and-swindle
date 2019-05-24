pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

/// @title AbstractWitness - the sw3 witness interface
contract AbstractWitness {
  /* valid testimony values */
  enum TestimonyStatus { PENDING, VALID, INVALID}

  /// @notice get testimony  
  /// @param payload arbitrary data
  /// @return status, encoded RoleAssignment[]
  function testimonyFor(bytes memory specification, bytes memory payload)
  public view returns (TestimonyStatus, bytes memory);

}
