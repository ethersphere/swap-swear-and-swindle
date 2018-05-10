pragma solidity ^0.4.19;
import "zeppelin/math/SafeMath.sol";
import "zeppelin/math/Math.sol";
import "./abstracts/AbstractWitness.sol";

contract SW3Utils {

  function chequeHash(address swap, address beneficiary, uint serial, uint amount) public view returns (bytes32) {
    return keccak256(abi.encodePacked(swap, serial, beneficiary, amount));
  }

  function noteHash(address swap, address beneficiary, uint index, uint amount, address witness, uint validFrom, uint validUntil, bytes32 remark) public view returns (bytes32) {
    return keccak256(encodeNote(swap, beneficiary, index, amount, witness, validFrom, validUntil, remark));
  }

  function invoiceHash(bytes32 noteId, uint swapBalance, uint serial) public pure returns (bytes32) {
    return keccak256(abi.encodePacked(noteId, swapBalance, serial));
  }

  function encodeNote(address swap, address beneficiary, uint index, uint amount, address witness, uint validFrom, uint validUntil, bytes32 remark) public view returns (bytes) {
    return abi.encodePacked(swap, index, beneficiary, amount, witness, validFrom, validUntil, remark);
  }

  function decodeNote(bytes note)
  public view returns (address swap, address beneficiary, uint index, uint amount, address witness, uint validFrom, uint validUntil, bytes32 remark) {
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

  function decodeSignature(bytes sig) internal pure returns (bytes32 r, bytes32 s, uint8 v) {
    assembly {
      r := mload(add(sig, 32))
      s := mload(add(sig, 64))
      v := and(mload(add(sig, 65)), 0xff)
    }

    v += 27; /* TODO: mainnet? */
  }

  function recoverSignature(bytes32 hash, bytes sig) internal pure returns (address) {
    var (r, s, v) = decodeSignature(sig);
    return ecrecover(keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", hash)), v, r, s);
  }

}
