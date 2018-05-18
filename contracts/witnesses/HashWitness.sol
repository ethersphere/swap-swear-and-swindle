pragma solidity ^0.4.23;
import "../abstracts/AbstractWitness.sol";

/// @title HashWitness - Witness expects data for some keccak256 hash to be submitted
contract HashWitness is AbstractWitness {
  event Testified(bytes32 noteId, bytes32 hash);

  mapping (bytes32 => TestimonyStatus) testimonies;

  /* returns VALID if data was submitted, PENDING otherwise */
  function testimonyFor(address, address, bytes32 noteId)
  public view returns (TestimonyStatus) {
    return testimonies[noteId];
  }

  /* remark is expected to be keccak256(trial, keccak256(data)) */
  function testify(bytes data) public {
    testimonies[keccak256(data)] = TestimonyStatus.VALID;    
  }

}
