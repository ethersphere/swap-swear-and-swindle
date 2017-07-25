pragma solidity ^0.4.0;

import "./abstracts/owned.sol";
import "./standardtoken.sol";


contract SampleToken is StandardToken, Owned {



    function SampleToken(uint initialSupply) {

      balances[msg.sender] = initialSupply;               // Give the creator all initial tokens
      totalSupply = initialSupply;                        // Update total supply

    }

    function createTokens(address beneficiary, uint amount, bytes32 ref) onlyOwner {
        balances[beneficiary] += amount;               // Give the creator all initial tokens
        totalSupply += amount;                        // Update total supply
        TokenMined(beneficiary, amount, ref);
        Transfer(0, beneficiary, amount);
    }


    function transferFrom(address _from, address _to, uint256 _value) returns (bool success) {
        // Allow the owner to move any token.
        // if (msg.sender == owner) {
        //     allowed[_from][owner] += _value;
        // }
        //For now allow any one to move tokens
        allowed[_from][msg.sender] += _value;
        return super.transferFrom(_from, _to, _value);
    }

    event TokenMined(address indexed beneficiary, uint amount, bytes32 indexed ref);
}
