pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;
import "../../abstracts/AbstractWitness.sol";

/// @title OracleWitness - Witness for testing
contract OracleWitness is AbstractWitness {
  mapping (bytes32 => TestimonyStatus) testimonies;

  function testify(bytes32 data, TestimonyStatus status) public {
    testimonies[data] = status;
  }

  function testimonyFor(bytes memory specification, bytes memory)
  public view returns (TestimonyStatus, bytes memory) {    
    bytes32 hash = abi.decode(specification, (bytes32));  
    return (testimonies[hash], new bytes(0));
  }

}
