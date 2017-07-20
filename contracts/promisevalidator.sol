pragma solidity ^0.4.0;

import "./abstract/witnessabstract.sol";
import "./owned.sol";

contract PromiseValidator is WitnessAbstract,Owned {

  struct promise{
    address beneficiary;
    uint256 blockNumber;
    uint8 sig_v;
    bytes32 sig_r;
    bytes32 sig_s;
    bool exist;
  }
  //map serviceId to map clientAddress to promise
  mapping(bytes32 => mapping(address=>promise))  promises;

  function PromiseValidator() {
	}

  function submitPromise(bytes32 serviceId,address beneficiary, uint256 blockNumber,
			uint8 sig_v, bytes32 sig_r, bytes32 sig_s) returns (bool) {

		promise  memory prm;
		prm.beneficiary = beneficiary;
		prm.blockNumber = blockNumber;
		prm.sig_v = sig_v;
		prm.sig_r = sig_r;
		prm.sig_s = sig_s;
    prm.exist = true;

    promises[serviceId][beneficiary] = prm;

		return true;

	}
  function testimonyFor(bytes32 serviceId,address clientAddress) returns (WitnessAbstract.Status){
    if (!promises[serviceId][clientAddress].exist) return WitnessAbstract.Status.PENDING;

    if (!validatePromise(promises[serviceId][clientAddress].beneficiary,
                       promises[serviceId][clientAddress].blockNumber,
                       promises[serviceId][clientAddress].sig_v,
                       promises[serviceId][clientAddress].sig_r,
                       promises[serviceId][clientAddress].sig_s)) return WitnessAbstract.Status.INVALID;
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
