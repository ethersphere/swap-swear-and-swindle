pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "../../abstracts/AbstractWitness.sol";
import "../../SW3Utils.sol";
import "../../Swap.sol";
import "openzeppelin-solidity/contracts/cryptography/MerkleProof.sol";

contract MerkleWitness is AbstractWitness, SW3Utils {  

  function testimonyFor(bytes memory specification, bytes memory data)
  public view returns (TestimonyStatus, bytes memory) {
    (bytes32 root, bytes32 leaf) = abi.decode(specification, (bytes32, bytes32));
    bytes32[] memory proof = abi.decode(data, (bytes32[]));

    require(MerkleProof.verify(proof, root, keccak256(abi.encode(leaf))), "invalid proof");

    return (TestimonyStatus.VALID, new bytes(0));
  }

  function encodeProof(bytes32[] memory proof)
  public view returns (bytes memory) {
    return abi.encode(proof);
  }
}
