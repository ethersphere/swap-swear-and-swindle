pragma solidity ^0.4.0;

import "./abstracts/witnessabstract.sol";
import "./abstracts/trialrulesabstract.sol";
import "./abstracts/owned.sol";

contract PromiseValidator is WitnessAbstract,Owned {

  struct promise{
    address beneficiary;
    uint256 blockNumber;
    uint8 sig_v;
    bytes32 sig_r;
    bytes32 sig_s;
    bool exist;
  }
  //map caseId to map serviceId to map clientAddress to promise
  mapping(bytes32 => mapping(bytes32 => mapping(address=>promise))) promises;
  //map caseid to pending time
  mapping(bytes32 => uint)  gracePeriods;

  function PromiseValidator() {
	}
  /// @notice submitPromise - submit a signed promise by client
  /// for this case it validate the service promise which is submitted by the client.
  ///
  /// @param caseId case id
  /// @param serviceId the service id
  /// @param beneficiary beneficiary address
  /// @param blockNumber the Promise is valid until this block number
  /// @param sig_v signature parameter v
  /// @param sig_r signature parameter r
  /// @param sig_s signature parameter s
  /// @return bool  - true if promise already submitted for this case,service and  beneficiary
  ///                 otherwise false
  function submitPromise(bytes32 caseId,bytes32 serviceId,address beneficiary, uint256 blockNumber,
			uint8 sig_v, bytes32 sig_r, bytes32 sig_s) returns (bool) {

    if (promises[caseId][serviceId][beneficiary].exist) return false;

    promises[caseId][serviceId][beneficiary] = promise(beneficiary,blockNumber,sig_v,sig_r,sig_s,true);

		return true;

	}
  /// @notice isEvidenceSubmited - check if an evidence was submited for a specific case ,service and client
  ///
  /// @param caseId case id
  /// @param serviceId the service id which
  /// @param clientAddress client address
  /// @return bool - true or false
  function isEvidentSubmited(bytes32 caseId, bytes32 serviceId,address clientAddress) returns (bool){
    return promises[caseId][serviceId][clientAddress].exist;
  }
  /// @notice testimonyFor - request for testimony for a specific case ,service and client
  /// for this case it validate the service promise which is submitted by the client.
  ///
  /// @param caseId case id
  /// @param serviceId the service id which
  /// @param clientAddress client address
  /// @return Status { VALID,INVALID, PENDING}
  function testimonyFor(bytes32 caseId, bytes32 serviceId,address clientAddress) returns (WitnessAbstract.Status){

    if (!validatePromise(promises[caseId][serviceId][clientAddress].beneficiary,
                       promises[caseId][serviceId][clientAddress].blockNumber,
                       promises[caseId][serviceId][clientAddress].sig_v,
                       promises[caseId][serviceId][clientAddress].sig_r,
                       promises[caseId][serviceId][clientAddress].sig_s)) return WitnessAbstract.Status.INVALID;
     return WitnessAbstract.Status.VALID;
  }
  /// @notice validatePromise
  ///
  /// @param beneficiary beneficiary address
  /// @param blockNumber the Promise is valid until this block number
  /// @param sig_v signature parameter v
  /// @param sig_r signature parameter r
  /// @param sig_s signature parameter s
  /// The digital signature is calculated on the concatenated triplet of contract address,beneficiary and blockNumber
  function validatePromise(address beneficiary, uint256 blockNumber,
      uint8 sig_v, bytes32 sig_r, bytes32 sig_s) private returns (bool){

      //check current block number is less than the Promise blocknumber
      if (block.number >= blockNumber ) return false;
      // Check the digital signature of the Promise.
      bytes32 hash = sha3(address(this), beneficiary, blockNumber);
      if(owner != ecrecover(hash, sig_v, sig_r, sig_s)) return false;
      return true;
  }

}
