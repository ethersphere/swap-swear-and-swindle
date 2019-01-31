pragma solidity ^0.5.0;
import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "openzeppelin-solidity/contracts/math/Math.sol";
import "openzeppelin-solidity/contracts/cryptography/ECDSA.sol";
import "./abstracts/AbstractWitness.sol";

/// @title Common functions, ideally Swap or Swear libraries would inherit this
contract SW3Utils {

  struct Note {
    bytes32 id;
    address swap;
    uint index; /* only used as a nonce for now, 0 is invalid */
    uint amount; /* amount of the note */
    address payable beneficiary; /* total amount paid out */
    address witness; /* witness used as escrow */
    uint validFrom; /* earliest timestamp for submission and payout */
    uint validUntil; /* latest timestamp for submission and payout */
    bytes32 remark ;/* arbitrary 32-bytes, can be used to encode information for Swear and witnesses */
  }

  /// @dev compute hash for a cheque
  function chequeHash(address swap, address beneficiary, uint serial, uint amount, uint timeout)
  public pure returns (bytes32) {
    return keccak256(abi.encodePacked(swap, serial, beneficiary, amount, timeout));
  }

  /// @dev compute hash for a note
  function noteHash(address swap, address beneficiary, uint index, uint amount, address witness, uint validFrom, uint validUntil, bytes32 remark)
  public pure returns (bytes32) {
    return keccak256(encodeNote(swap, beneficiary, index, amount, witness, validFrom, validUntil, remark));
  }

  /// @dev compute hash for an invoice
  function invoiceHash(bytes32 noteId, uint swapBalance, uint serial) public pure returns (bytes32) {
    return keccak256(abi.encodePacked(noteId, swapBalance, serial));
  }

  /// @dev encode a note to bytes, this form can be useful to avoid StackTooDeep issues
  function encodeNote(address swap, address beneficiary, uint index, uint amount, address witness, uint validFrom, uint validUntil, bytes32 remark)
  public pure returns (bytes memory) {
    return abi.encode(swap, index, beneficiary, amount, witness, validFrom, validUntil, remark);
  }

  function decodeNote(bytes memory note)
  internal pure returns (Note memory n) {
    (address swap, uint index, address payable beneficiary, uint amount, address witness, uint validFrom, uint validUntil, bytes32 remark)
      = abi.decode(note, (address, uint, address, uint, address, uint, uint, bytes32));

    return Note({
      id: keccak256(note),
      swap: swap,
      index: index,
      beneficiary: beneficiary,
      amount: amount,
      witness: witness,
      validFrom: validFrom,
      validUntil: validUntil,
      remark: remark
    });
  }

  function recover(bytes32 hash, bytes memory sig) internal pure returns (address) {
    return ECDSA.recover(ECDSA.toEthSignedMessageHash(hash), sig);
  }

}
