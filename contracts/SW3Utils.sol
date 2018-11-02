pragma solidity ^0.4.23;
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
    address beneficiary; /* total amount paid out */
    address witness; /* witness used as escrow */
    uint validFrom; /* earliest timestamp for submission and payout */
    uint validUntil; /* latest timestamp for submission and payout */
    bytes32 remark ;/* arbitrary 32-bytes, can be used to encode information for Swear and witnesses */
  }

  /// @dev compute hash for a cheque
  function chequeHash(address swap, address beneficiary, uint serial, uint amount)
  public pure returns (bytes32) {
    return keccak256(abi.encodePacked(swap, serial, beneficiary, amount));
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
  public pure returns (bytes) {
    return abi.encodePacked(swap, index, beneficiary, amount, witness, validFrom, validUntil, remark);
  }

  function decodeNote(bytes note)
  internal pure returns (Note memory n) {
    address swap;
    uint index;
    address beneficiary;
    uint amount;
    address witness;
    uint validFrom;
    uint validUntil;
    bytes32 remark;

    uint divisor = 2**96;
    assembly {
      swap := div(mload(add(note, 32)), divisor)
      index := mload(add(note, 52))
      beneficiary := div(mload(add(note, 84)), divisor)
      amount := mload(add(note, 104))
      witness := div(mload(add(note, 136)), divisor)
      validFrom := mload(add(note, 156))
      validUntil := mload(add(note, 188))
      remark := mload(add(note, 220))
    }

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

  function recover(bytes32 hash, bytes sig) public pure returns (address) {
    return ECDSA.recover(ECDSA.toEthSignedMessageHash(hash), sig);
  }

}
