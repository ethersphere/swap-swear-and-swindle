pragma solidity ^0.4.0;

import "./abstracts/witnessabstract.sol";
import "./abstracts/ensabstract.sol";

contract MirrorENS is WitnessAbstract{

	struct ensNameHashPair{
    bytes32 clientNameHash;
		bytes32 serviceNameHash;
  }
  //map caseId to map serviceId to ensNameHashPair
	mapping(bytes32=> mapping(bytes32=> ensNameHashPair))  ensNameHashPairs;

	ENSAbstract ens;

	address public ensAddress = 0x8163bc885c2b14478b75f178ca76f31581dc967f;


  function MirrorENS() {
	}
	/// @notice testimonyFor - request for testimony for a specific case ,service and client
	///
	/// @param caseId case id
	/// @param serviceId the service id which
	/// @param clientAddress client address
	/// @return Status { VALID,INVALID, PENDING}
  function testimonyFor(bytes32 caseId,bytes32 serviceId,address clientAddress) returns (WitnessAbstract.Status){

    if (compareNameHashPair(caseId,serviceId)) return WitnessAbstract.Status.VALID;
    return WitnessAbstract.Status.INVALID;

  }

	function submitNameHashes(bytes32 caseId,bytes32 serviceId, bytes32 clientNameHash , bytes32 serviceNameHash) returns (bool) {
		if (ensNameHashPairs[caseId][serviceId].clientNameHash != bytes32(0x0)) return false; //do not allow override submition
		ensNameHashPairs[caseId][serviceId] = ensNameHashPair(clientNameHash,serviceNameHash);
		return true;
  }
	/// @notice isEvidenceSubmitted - check if an evidence was submitted for a specific case ,service and client
  ///
  /// @param caseId case id
  /// @param serviceId the service id which
  /// @param clientAddress client address
  /// @return bool - true or false
  function isEvidenceSubmitted(bytes32 caseId, bytes32 serviceId,address clientAddress) returns (bool){
	  return (ensNameHashPairs[caseId][serviceId].clientNameHash != bytes32(0x0));
  }
	/// @notice ensResolve - resolve the ens
  ///
  /// @param node - the node to resolve
  /// @return bytes32 - the hash of the resoved ENS node
	function ensResolve(bytes32 node) private constant returns(bytes32) {
			address resolverAddress = ens.resolver(node);
			ResolverAbstract resolver = ResolverAbstract(resolverAddress);
			bytes32 content = resolver.content(node);
			return content;
	}
	/// @notice compareNameHashPair - compare the ensNameHashPair submitted for a certain case and service.
  ///
  /// @param caseId - case id
	//  @param serviceId - service id
  /// @return bytes32 - true (equal) otherwise false
	function compareNameHashPair(bytes32 caseId, bytes32 serviceId) private returns(bool){
		  //check if the two nodes resolved ENS are equal
			//for each specific game the the decision should be take diffrently
			ens = ENSAbstract(ensAddress);
			if (ensResolve(ensNameHashPairs[caseId][serviceId].clientNameHash)!= ensResolve(ensNameHashPairs[caseId][serviceId].serviceNameHash)){
				return true;
			}
			return false;
	}
  //Important!!! This function should is here just for testing pouposes and to enable setting that for other enss.
	//the ens address for this game ("mirror") should be hard coded in the contract
	function setENSAddress(address _ensAddr) returns(bool){
    ensAddress = _ensAddr;
		return true;
	}

}
