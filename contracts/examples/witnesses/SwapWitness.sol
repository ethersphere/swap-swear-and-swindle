pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "../../abstracts/AbstractWitness.sol";
import "../../SW3Utils.sol";
import "../../Swap.sol";

contract SwapWitness is AbstractWitness, SW3Utils {

  function testimonyFor(bytes memory specification, bytes memory data)
  public view returns (TestimonyStatus, bytes memory) { 
    (      
      Note memory template,
      address expectedOwner
    ) = abi.decode(specification, (Note, address));

    (
      bytes memory encodedNote,
      bytes memory sig
    ) = abi.decode(data, (bytes, bytes));

    Note memory note = abi.decode(encodedNote, (Note));    

    if(template.swap != (address(0))) require(note.swap == template.swap, "invalid swap");
    if(template.index != 0) require(note.index == template.index, "invalid index");
    if(template.amount != 0) require(note.amount == template.amount, "invalid amount");
    if(template.beneficiary != address(0)) require(note.beneficiary == template.beneficiary, "invalid beneficiary");
    if(template.witness != address(0)) require(note.witness == template.witness, "invalid witness");
    if(template.validFrom != 0) require(note.validFrom == template.validFrom, "invalid validFrom");
    if(template.validUntil != 0) require(note.validUntil == template.validUntil, "invalid validUntil");
    if(template.remark != 0) require(note.remark == template.remark, "invalid remark");
    if(template.timeout != 0) require(note.timeout == template.timeout, "invalid timeout");

    bytes32 noteHash = keccak256(encodedNote);

    address owner = recover(noteHash, sig);

    require(owner == Swap(uint160(note.swap)).owner());
    if(expectedOwner != address(0)) require(owner == expectedOwner, "invalid owner");

    bytes[] memory kvs = new bytes[](10);
    kvs[0] = abi.encode(note.swap);
    kvs[1] = abi.encode(note.index);
    kvs[2] = abi.encode(note.amount);
    kvs[3] = abi.encode(note.beneficiary);
    kvs[4] = abi.encode(note.witness);
    kvs[5] = abi.encode(note.validFrom);
    kvs[6] = abi.encode(note.validUntil);
    kvs[7] = abi.encode(note.remark);
    kvs[8] = abi.encode(note.timeout);
    kvs[9] = abi.encode(noteHash);
    
    return (TestimonyStatus.VALID, abi.encode(kvs));
  }
}
