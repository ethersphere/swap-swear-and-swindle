pragma solidity ^0.4.23;
import "../abstracts/AbstractWitness.sol";

/// @title OracleWitness - Witness for testing
contract OracleWitness is AbstractWitness {
  mapping (bytes32 => TestimonyStatus) testimonies;

  function testify(bytes32 noteId, TestimonyStatus status) public {
    testimonies[noteId] = status;
  }

  function testimonyFor(address, address, bytes32 noteId)
  public view returns (TestimonyStatus) {
    return testimonies[noteId];
  }

}
