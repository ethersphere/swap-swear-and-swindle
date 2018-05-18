pragma solidity ^0.4.23;
import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "openzeppelin-solidity/contracts/math/Math.sol";
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

  /// @dev decode a signature
  function decodeSignature(bytes sig) internal pure returns (bytes32 r, bytes32 s, uint8 v) {
    assembly {
      r := mload(add(sig, 32))
      s := mload(add(sig, 64))
      v := and(mload(add(sig, 65)), 0xff)
    }

    v += 27; /* TODO: ganache and real clients might not be compatible here */
  }

  /// @dev recover signature from a web3.eth.sign() message
  function recoverSignature(bytes32 hash, bytes sig) internal pure returns (address) {
    bytes32 r;
    bytes32 s;
    uint8 v;
    (r, s, v) = decodeSignature(sig);
    return ecrecover(keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", hash)), v, r, s);
  }

}
