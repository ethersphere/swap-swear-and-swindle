pragma solidity ^0.4.23;
import "../abstracts/AbstractWitness.sol";

/* low-quality implementation of a witness for BMT, TO BE REPLACED */
/// @title ChunkWitness - Witness expects data for some Swarm POC3 hash
contract ChunkWitness is AbstractWitness {
  uint constant span = 2048;
  uint constant section = 64;

  event Testified(bytes32 noteId, bytes32 hash);

  mapping (bytes32 => TestimonyStatus) testimonies;

  function testimonyFor(address, address, bytes32 noteId)
  public view returns (TestimonyStatus) {
    return testimonies[noteId];
  }

  /* remark is expected to be keccak256(trial, swarmHash(data)) */
  function testify(bytes data) public {
    testimonies[bmt(data)] = TestimonyStatus.VALID;
  }

  function bmt(bytes d) public returns(bytes32) {
    return keccak256(abi.encodePacked(swap_uint64(uint64(d.length)), h(d, span)));
  }

  /* based on RefHasher from go-ethereum */
  function h(bytes d, uint s) private returns (bytes32) {
    uint length = d.length;

    bytes memory left = d;
    bytes memory right = new bytes(0);

    if(length > section) {
        for(; s >= length; s /= 2) {}
        left = bytes32ToBytes(h(slice(d, 0, s), s));
        right = slice(d, s, length - s);
        if((length - s) > section / 2) {
          right = bytes32ToBytes(h(right, s));
        }
    }
    return keccak256(abi.encodePacked(left, right));
  }

  function bytes32ToBytes(bytes32 data) internal pure returns (bytes) {
    bytes memory result = new bytes(32);
    for (uint i = 0; i < 32; i++) {
        result[i] = data[i];
    }
    return result;
  }

  function swap_uint64(uint64 val) private pure returns (uint64)
  {
    val = ((val << 8) & 0xFF00FF00FF00FF00 ) | ((val >> 8) & 0x00FF00FF00FF00FF );
    val = ((val << 16) & 0xFFFF0000FFFF0000 ) | ((val >> 16) & 0x0000FFFF0000FFFF );
    return (val << 32) | (val >> 32);
  }

  /* copied from BytesLibrary */
  function slice(bytes _bytes, uint _start, uint _length) internal  pure returns (bytes) {
    require(_bytes.length >= (_start + _length));

    bytes memory tempBytes;

    assembly {
        switch iszero(_length)
        case 0 {
            // Get a location of some free memory and store it in tempBytes as
            // Solidity does for memory variables.
            tempBytes := mload(0x40)

            // The first word of the slice result is potentially a partial
            // word read from the original array. To read it, we calculate
            // the length of that partial word and start copying that many
            // bytes into the array. The first word we copy will start with
            // data we don't care about, but the last `lengthmod` bytes will
            // land at the beginning of the contents of the new array. When
            // we're done copying, we overwrite the full first word with
            // the actual length of the slice.
            let lengthmod := and(_length, 31)

            // The multiplication in the next line is necessary
            // because when slicing multiples of 32 bytes (lengthmod == 0)
            // the following copy loop was copying the origin's length
            // and then ending prematurely not copying everything it should.
            let mc := add(add(tempBytes, lengthmod), mul(0x20, iszero(lengthmod)))
            let end := add(mc, _length)

            for {
                // The multiplication in the next line has the same exact purpose
                // as the one above.
                let cc := add(add(add(_bytes, lengthmod), mul(0x20, iszero(lengthmod))), _start)
            } lt(mc, end) {
                mc := add(mc, 0x20)
                cc := add(cc, 0x20)
            } {
                mstore(mc, mload(cc))
            }

            mstore(tempBytes, _length)

            //update free-memory pointer
            //allocating the array padded to 32 bytes like the compiler does now
            mstore(0x40, and(add(mc, 31), not(31)))
        }
        //if we want a zero-length slice let's just return a zero-length array
        default {
            tempBytes := mload(0x40)

            mstore(0x40, add(tempBytes, 0x20))
        }
    }

    return tempBytes;
  }
}
