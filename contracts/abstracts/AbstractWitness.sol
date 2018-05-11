pragma solidity ^0.4.0;

/// @title AbstractWitness - the sw3 witness interface
contract AbstractWitness {

  /* valid testimony values */
  enum TestimonyStatus { PENDING, VALID, INVALID}

  /// @notice get testimony
  /// @param owner Swap channel owner or service provider
  /// @param beneficiary beneficiary or plaintiff
  /// @param noteId arbitrary data or Swap note hash
  function testimonyFor(address owner, address beneficiary, bytes32 noteId) public view returns (TestimonyStatus);

}
