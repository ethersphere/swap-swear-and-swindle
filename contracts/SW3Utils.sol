pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;
import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "openzeppelin-solidity/contracts/math/Math.sol";
import "openzeppelin-solidity/contracts/cryptography/ECDSA.sol";
import "./abstracts/AbstractWitness.sol";

/// @title Common functions, ideally Swap or Swear libraries would inherit this
contract SW3Utils {

  struct Note {    
    address swap;
    uint index; /* only used as a nonce for now, 0 is invalid */
    uint amount; /* amount of the note */
    address payable beneficiary; /* total amount paid out */
    address witness; /* witness used as escrow */
    uint validFrom; /* earliest timestamp for submission and payout */
    uint validUntil; /* latest timestamp for submission and payout */
    bytes32 remark ;/* arbitrary 32-bytes, can be used to encode information for Swear and witnesses */
    uint timeout;
  }
  
  /// @dev compute hash for an invoice
  function invoiceHash(bytes32 noteId, uint swapBalance, uint serial) public pure returns (bytes32) {
    return keccak256(abi.encodePacked(noteId, swapBalance, serial));
  }

  function encodeNote(Note memory note)
  public pure returns (bytes memory) {
    return abi.encode(note);
  }

  function noteHash(Note memory note)
  public pure returns (bytes32) {
    return keccak256(encodeNote(note));
  }

  function recover(bytes32 hash, bytes memory sig) internal pure returns (address) {
    return ECDSA.recover(ECDSA.toEthSignedMessageHash(hash), sig);
  }

}
