pragma solidity ^0.4.0;

contract AbstractWitness {

    enum TestimonyStatus { PENDING, VALID, INVALID}

    function testimonyFor(address owner, address beneficiary, bytes32 noteId) public view returns (TestimonyStatus);

}
