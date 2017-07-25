pragma solidity ^0.4.0;

contract WitnessAbstract {

  enum Status { VALID,INVALID, PENDING}
  /// @notice testimonyFor - request for testimony for a specific case ,service and client
  ///
  /// @param caseId case id
  /// @param serviceId the service id which
  /// @param clientAddress client address
  /// @return Status { VALID,INVALID, PENDING}
  function testimonyFor(bytes32 caseId,bytes32 serviceId,address clientAddress) returns (Status);
  /// @notice isEvidentSubmited - check if an evident was submited for a specific case ,service and client
  ///
  /// @param caseId case id
  /// @param serviceId the service id which
  /// @param clientAddress client address
  /// @return bool - true or false
  function isEvidentSubmited(bytes32 caseId, bytes32 serviceId,address clientAddress) returns (bool);
}
