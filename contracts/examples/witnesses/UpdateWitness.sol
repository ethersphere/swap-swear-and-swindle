pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "../../abstracts/AbstractWitness.sol";
import "../../SW3Utils.sol";
import "../../Swap.sol";

contract UpdateWitness is AbstractWitness, SW3Utils {

  struct UpdateMessage {    
    address destination;
    uint timeFrom;
    uint timeUntil;
    bytes32 updateRoot;
  }

  function encodeUpdateMessage(UpdateMessage memory message) 
  public view returns (bytes memory) {
    return abi.encode(message);
  }

  function updateHash(UpdateMessage memory message)
  public view returns (bytes32) {
    return keccak256(abi.encode(message));
  }

  function testimonyFor(bytes memory specification, bytes memory data)
  public view returns (TestimonyStatus, bytes memory) {
    (
      address expectedSigner,
      address expectedDestination,
      uint expectedTime
    ) = abi.decode(specification, (address, address, uint));
    (
      bytes memory encodedUpdate,
      bytes memory sig      
    ) = abi.decode(data, (bytes, bytes));

    require(expectedSigner == recover(keccak256(encodedUpdate), sig), "wrong ack signer");

    UpdateMessage memory update = abi.decode(encodedUpdate, (UpdateMessage));

    require(update.destination == expectedDestination, "wrong destination");
    require(update.timeUntil >= expectedTime, "message after update");
    require(update.timeFrom <= expectedTime, "message before update");

    bytes[] memory kvs = new bytes[](1);
    kvs[0] = abi.encode(update.updateRoot);    
    return (TestimonyStatus.VALID, abi.encode(kvs));
  }
}
