pragma solidity ^0.4.0;

import "./abstracts/witnessabstract.sol";
import "./abstracts/ensabstract.sol";

contract MirrorENS is WitnessAbstract{

  //This is the specific game logic "mirror"
	//specific game "mirror"
	ENSAbstract ens;
	//client.game namehash
	bytes32 constant clientENSNameHash = 0x94c4860d894e91f2df683b61455630d721209c6265d2e80c86a1f92cab14b370;
	//reflector.game namehash
	bytes32 constant reflectorENSNameHash = 0xacd7f5ed7d93b1526477b93e6c7def60c40420a868e7f694a7671413d89bb9a5;

	address public ensAddress = 0x8163bc885c2b14478b75f178ca76f31581dc967f;


  function MirrorENS() {
	}

  function testimonyFor(bytes32 serviceId,address clientAddress) returns (WitnessAbstract.Status){

    if (guilty()) return WitnessAbstract.Status.VALID;
    return WitnessAbstract.Status.INVALID;

  }

	function ensResolve(bytes32 node) private constant returns(bytes32) {
			address resolverAddress = ens.resolver(node);
			ResolverAbstract resolver = ResolverAbstract(resolverAddress);
			bytes32 content = resolver.content(node);
			return content;
	}

	function guilty() private returns(bool){
		  //check if the two nodes resolved ENS are equal
			//for each specific game the the decision should be take diffrently
			ens = ENSAbstract(ensAddress);
			if (ensResolve(clientENSNameHash)!= ensResolve(reflectorENSNameHash)){
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
