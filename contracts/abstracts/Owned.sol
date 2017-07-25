pragma solidity ^0.4.0;

contract Owned {
    /// Allows only the owner to call a function
    modifier onlyOwner { require (msg.sender == owner); _; }

    address public owner;

    function Owned() { owner = msg.sender;}



    function changeOwner(address _newOwner) onlyOwner {
        owner = _newOwner;
    }
}
