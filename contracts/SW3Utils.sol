pragma solidity ^0.4.19;
import "zeppelin/math/SafeMath.sol";
import "zeppelin/math/Math.sol";
import "./abstracts/AbstractWitness.sol";

/// @title Common functions, ideally Swap or Swear libraries would inherit this
contract SW3Utils {

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

  /// @dev decode a note from bytes
  function decodeNote(bytes note)
  public pure returns (address swap, address beneficiary, uint index, uint amount, address witness, uint validFrom, uint validUntil, bytes32 remark) {
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
