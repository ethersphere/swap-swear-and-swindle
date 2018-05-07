pragma solidity ^0.4.19;
import "./abstracts/AbstractRules.sol";
import "./Swindle.sol";

contract Swear {

  Swindle public swindle;

  function Swear(address _swindle) {
    swindle = Swindle(_swindle);
  }

}
