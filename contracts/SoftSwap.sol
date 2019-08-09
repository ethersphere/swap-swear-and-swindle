pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;
import "./SimpleSwap.sol";

/// @title Swap Channel Contract with Soft Deposits
contract SoftSwap is SimpleSwap {

  constructor(address payable _owner) SimpleSwap(_owner, 2 days) public { }

  uint public softDeposit;
  
  /// @return the part of the balance that is not covered by hard deposits or the soft deposit
  function liquidBalance() public view returns(uint) {
    return address(this).balance.sub(totalHardDeposit).sub(softDeposit);
  }
}