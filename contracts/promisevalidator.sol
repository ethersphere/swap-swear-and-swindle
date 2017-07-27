pragma solidity ^0.4.0;

import "./abstracts/witnessabstract.sol";
import "./abstracts/trialrulesabstract.sol";
import "./abstracts/owned.sol";


contract PromiseValidator is WitnessAbstract,Owned {

    struct Promise {
        address beneficiary;
        uint256 blockNumber;
        uint8 sigV;
        bytes32 sigR;
        bytes32 sigS;
        bool exist;
    }
    //map caseId to map serviceId to map clientAddress to Promise
    mapping(bytes32 => mapping(bytes32 => mapping(address=>Promise))) Promises;
    //map caseid to pending time
    mapping(bytes32 => uint)  gracePeriods;

    function PromiseValidator() {
     }

    /// @notice submitPromise - submit a signed Promise by client
    /// for this case it validate the service Promise which is submitted by the client.
    ///
    /// @param caseId case id
    /// @param serviceId the service id
    /// @param beneficiary beneficiary address
    /// @param blockNumber the Promise is valid until this block number
    /// @param sigV signature parameter v
    /// @param sigR signature parameter r
    /// @param sigS signature parameter s
    /// @return bool  - true if Promise already submitted for this case,service and  beneficiary
    ///                 otherwise false
    function submitPromise(
        bytes32 caseId,
        bytes32 serviceId,
        address beneficiary,
        uint256 blockNumber,
        uint8 sigV,
        bytes32 sigR,
        bytes32 sigS) returns (bool)
        {
        if (Promises[caseId][serviceId][beneficiary].exist)
        return false;
        Promises[caseId][serviceId][beneficiary] = Promise(
            beneficiary,
            blockNumber,
            sigV,
            sigR,
            sigS,
            true
        );
        return true;
    }

    /// @notice isEvidenceSubmitted - check if an evidence was submitted for a specific case ,service and client
    ///
    /// @param caseId case id
    /// @param serviceId the service id which
    /// @param clientAddress client address
    /// @return bool - true or false
    function isEvidenceSubmitted(bytes32 caseId, bytes32 serviceId,address clientAddress) returns (bool) {
        return Promises[caseId][serviceId][clientAddress].exist;
    }

    /// @notice testimonyFor - request for testimony for a specific case ,service and client
    /// for this case it validate the service Promise which is submitted by the client.
    ///
    /// @param caseId case id
    /// @param serviceId the service id which
    /// @param clientAddress client address
    /// @return Status { VALID,INVALID, PENDING}
    function testimonyFor(bytes32 caseId, bytes32 serviceId,address clientAddress) returns (WitnessAbstract.Status) {
        if (!validatePromise(
            Promises[caseId][serviceId][clientAddress].beneficiary,
            Promises[caseId][serviceId][clientAddress].blockNumber,
            Promises[caseId][serviceId][clientAddress].sigV,
            Promises[caseId][serviceId][clientAddress].sigR,
            Promises[caseId][serviceId][clientAddress].sigS))
                return WitnessAbstract.Status.INVALID;
        return WitnessAbstract.Status.VALID;
    }

    /// @notice validatePromise
    ///
    /// @param beneficiary beneficiary address
    /// @param blockNumber the Promise is valid until this block number
    /// @param sigV signature parameter v
    /// @param sigR signature parameter r
    /// @param sigS signature parameter s
    /// The digital signature is calculated on the concatenated triplet of contract address,beneficiary and blockNumber
    function validatePromise(
        address beneficiary,
        uint256 blockNumber,
        uint8 sigV,
        bytes32 sigR,
        bytes32 sigS) private returns (bool)
        {

        //check current block number is less than the Promise blocknumber
        if (block.number >= blockNumber )
        return false;
        // Check the digital signature of the Promise.
        bytes32 hash = sha3(address(this), beneficiary, blockNumber);
        if (owner != ecrecover(
            hash,
            sigV,
            sigR,
            sigS
            ))
            return false;
        return true;
    }
}
