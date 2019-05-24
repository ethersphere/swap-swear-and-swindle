pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "../../abstracts/AbstractWitness.sol";
import "../../SW3Utils.sol";
import "../../Swap.sol";

contract AckWitness is AbstractWitness, SW3Utils {

  struct AckMessage {
    bytes32 dataHash;
    address destination;
    uint time;
  }

  function encodeAckMessage(AckMessage memory message) 
  public view returns (bytes memory) {
    return abi.encode(message);
  }

  function ackHash(AckMessage memory message)
  public view returns (bytes32) {
    return keccak256(abi.encode(message));
  }

  function testimonyFor(bytes memory specification, bytes memory data)
  public view returns (TestimonyStatus, bytes memory) {
    (
      address expectedSigner, 
      address expectedDestination
    ) = abi.decode(specification, (address, address));
    (
      bytes memory encodedAck,
      bytes memory sig      
    ) = abi.decode(data, (bytes, bytes));

    require(expectedSigner == recover(keccak256(encodedAck), sig), "wrong ack signer");

    AckMessage memory ack = abi.decode(encodedAck, (AckMessage));

    require(ack.destination == expectedDestination, "wrong destination");

    return (TestimonyStatus.VALID, abi.encode(ack.time, ack.dataHash));
  }
}
