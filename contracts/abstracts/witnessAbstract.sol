pragma solidity ^0.4.0;

contract WitnessAbstract {

  enum Status { VALID,INVALID, PENDING}
  function testimonyFor(bytes32 caseId,bytes32 serviceId,address clientAddress) returns (Status);

}
