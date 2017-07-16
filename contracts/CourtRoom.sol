pragma solidity ^0.4.0;

import "./owned.sol";
import "./sampletoken.sol";
import "./ensabstract.sol";


contract CaseContract is Owned {

  struct promise{
		address beneficiary;
		uint256 blockNumber;
		uint8 sig_v;
		bytes32 sig_r;
		bytes32 sig_s;
	}
	struct claim {
		bytes32 claimId;
		address plaintiff;
		bytes32[] evidence;
		promise   servicePromise;
		uint status;
		bool valid;
	}

	//claimid map to claim
  mapping(bytes32 => claim)  OpenClaims;


	function CaseContract() {
	}

	function newClaim(address _plaintiff, bytes32 _evidence) returns (bytes32 claimId) {

	  claim storage clm;
		clm.claimId = sha3(_plaintiff, _evidence, now);
		clm.plaintiff = _plaintiff;
		clm.evidence.push(_evidence);
		clm.status = 1;

		require(OpenClaims[clm.claimId].status == 0);

		OpenClaims[clm.claimId] = clm;

		return clm.claimId;
	}

	function submitEvidence(bytes32 _claimId,bytes32 _evident) returns (uint status) {

		// Maximum amount of evidence has been submitted
		require(OpenClaims[_claimId].evidence.length < 4);

		OpenClaims[_claimId].evidence.push(_evident);

		return OpenClaims[_claimId].status;

	}

	function submitPromise(bytes32 _claimId,address beneficiary, uint256 blockNumber,
			uint8 sig_v, bytes32 sig_r, bytes32 sig_s) returns (uint status) {

		promise  storage prm;
		prm.beneficiary = beneficiary;
		prm.blockNumber = blockNumber;
		prm.sig_v = sig_v;
		prm.sig_r = sig_r;
		prm.sig_s = sig_s;

		OpenClaims[_claimId].servicePromise = prm;

		return OpenClaims[_claimId].status;

	}


	function getStatus(bytes32 claimId) constant returns (uint status) {


		return OpenClaims[claimId].status;

	}

	function resolveClaim(bytes32 _claimId) {

		if (OpenClaims[_claimId].claimId == 0x0 || OpenClaims[_claimId].status == 0) throw;
		OpenClaims[_claimId].claimId = 0x0;
		OpenClaims[_claimId].plaintiff = 0;
		OpenClaims[_claimId].evidence.length = 0;
		OpenClaims[_claimId].status = 0;
		OpenClaims[_claimId].valid = false;

	}
	function getClaim(bytes32 _claimId) returns (address plaintiff,bool valid,address beneficiary, uint256 blockNumber,
			uint8 sig_v, bytes32 sig_r, bytes32 sig_s) {
        return (OpenClaims[_claimId].plaintiff,
                OpenClaims[_claimId].valid,
                OpenClaims[_claimId].servicePromise.beneficiary,
                OpenClaims[_claimId].servicePromise.blockNumber,
                OpenClaims[_claimId].servicePromise.sig_v,
                OpenClaims[_claimId].servicePromise.sig_r,
                OpenClaims[_claimId].servicePromise.sig_s
                );
    }

  function setClaimValid(bytes32 _claimId)  {
        OpenClaims[_claimId].valid = true;
  }

}

contract SwearGame is Owned {

	uint256 public amountStaked;
	uint public rewardCompensation;
	uint  public registeredPlayersCounter;
	mapping(address => bool) public registeredPlayers;
	mapping(address => bytes32[]) public clientsClaimsIds;

	CaseContract public caseContract;
	SampleToken public sampleToken;
	bytes32 public claimId;


	function SwearGame(address _caseContract, address _sampleToken, uint _rewardCompensation) {
		caseContract = CaseContract(_caseContract);
		sampleToken = SampleToken(_sampleToken);
		rewardCompensation = _rewardCompensation;
		registeredPlayersCounter = 0;

	}




	function deposit(uint256 _depositAmount) onlyOwner payable public returns(bool){

		require(sampleToken.balanceOf(owner) >= _depositAmount);

		require(sampleToken.transferFrom(owner, address(this), _depositAmount));

		amountStaked += _depositAmount;

		DepositStaked(_depositAmount, amountStaked);

		return true;

	}


	function makeJudgement(bytes32 _claimId) private returns(bool) {

	  bool claimantCompensated;
    bool decision = false;

    require(caseContract.getStatus(_claimId) != 0);
		// Somehow come to a resolution...
		var(plaintiff,valid,beneficiary,blockNumber,sig_v,sig_r,sig_s) = caseContract.getClaim(_claimId);

		// Case has already been compensated for
		require(!valid);

    decision = (validatePromise(beneficiary,blockNumber,sig_v,sig_r,sig_s) && takeDecision());

    caseContract.setClaimValid(_claimId);

    if (decision == true){
	    claimantCompensated = compensate(plaintiff);
			caseContract.resolveClaim(_claimId);
    }

		ClaimResolved(_claimId, plaintiff, rewardCompensation);
		return claimantCompensated;

	}

	function compensate(address _claimant) private returns(bool compensated) {

		compensated = sampleToken.transferFrom(address(this), _claimant, rewardCompensation);

		require(compensated);

		amountStaked -=rewardCompensation;

		return compensated;

	}


	function register(address _player) onlyOwner public returns (bool registered) {

		require(!registeredPlayers[_player]);

		if (registeredPlayersCounter == 0){
			require(amountStaked >= rewardCompensation);
		}else if ((amountStaked / registeredPlayersCounter) < rewardCompensation) {
			AdditionalDepositRequired(amountStaked);
			throw;
		}

		registeredPlayers[_player] = true;
		registeredPlayersCounter++;

		NewPlayer(_player);

		return true;

	}


	function leaveGame(address _player) onlyOwner public {

		// If the player is not registered to the game throw
		require(registeredPlayers[_player]);

		PlayerLeftGame(_player);

		registeredPlayers[_player] = false;
		registeredPlayersCounter--;

	}
	function openNewClaim(address beneficiary, uint256 blockNumber,
			uint8 sig_v, bytes32 sig_r, bytes32 sig_s,bytes32 _evidence) public returns (bool) {

		require(registeredPlayers[msg.sender]);

		claimId  = caseContract.newClaim(msg.sender,bytes32(_evidence));

		caseContract.submitPromise(claimId,beneficiary, blockNumber,sig_v, sig_r, sig_s);

		clientsClaimsIds[msg.sender].push(claimId);

		makeJudgement(claimId);

		return true;
	}

	/// @notice Cash cheque
	///
	/// @param beneficiary beneficiary address
	/// @param blockNumber the promise is valid until this block number
	/// @param sig_v signature parameter v
	/// @param sig_r signature parameter r
	/// @param sig_s signature parameter s
	/// The digital signature is calculated on the concatenated triplet of contract address,beneficiary and blockNumber
	function validatePromise(address beneficiary, uint256 blockNumber,
			uint8 sig_v, bytes32 sig_r, bytes32 sig_s) private returns (bool){

		  //check the exitance of the proof
      if (beneficiary == 0x0) return false;
      //check current block number is less than the promise blocknumber
      if (block.number >= blockNumber ) return false;
      // Check the digital signature of the promise.
			bytes32 hash = sha3(address(this), beneficiary, blockNumber);
			if(owner != ecrecover(hash, sig_v, sig_r, sig_s)) return false;
		  return true;
	}
	//This is the specific game logic "reflector"
	//specific game "reflector"
	ENSAbstract ens;
	//client.game namehash
	bytes32 constant clientENSNameHash = 0x94c4860d894e91f2df683b61455630d721209c6265d2e80c86a1f92cab14b370;
	//reflector.game namehash
	bytes32 constant reflectorENSNameHash = 0xacd7f5ed7d93b1526477b93e6c7def60c40420a868e7f694a7671413d89bb9a5;

	address public ensAddress = 0x8163bc885c2b14478b75f178ca76f31581dc967f;
	///////////////////////////////////////////////////////////////////////////////////////////////////////
	function ensResolve(bytes32 node) private constant returns(bytes32) {
			address resolverAddress = ens.resolver(node);
			ResolverAbstract resolver = ResolverAbstract(resolverAddress);
			bytes32 content = resolver.content(node);
			return content;
	}

	function takeDecision() private returns(bool){
		  //check if the two nodes resolved ENS are equal
			//for each specific game the the decision should be take diffrently
			ens = ENSAbstract(ensAddress);
			bytes32 contentHash1 = ensResolve(clientENSNameHash);
			bytes32 contentHash2 = ensResolve(reflectorENSNameHash);
			bool res = (contentHash1 != contentHash2);
			if (res){
				Decision("guilty");
			}
			else {
				Decision("not guilty");
			}
			return res;
	}
  //Important!!! This function should is here just for testing pouposes and to enable setting that for other enss.
	//the ens address for this game ("reflector") should be hard coded in the contract
	function setENSAddress(address _ensAddr) returns(bool){
    ensAddress = _ensAddr;
		return true;
	}


	event Decision(string decide);

	event DepositStaked(uint depositAmount, uint amountStaked);
	event Compensate(address recipient, uint rewardCompensation);
	event NewPlayer(address playerId);
	event PlayerLeftGame(address playerId);

	event NewClaimOpened(bytes32 caseId, address plaintiff);
	event NewEvidenceSubmitted(bytes32 claimId, address plaintiff);
	event ClaimResolved(bytes32 claimId, address plaintiff, uint rewardCompensation);
	event Payment(address from,address to ,uint256 value);
	event AdditionalDepositRequired(uint256 amountStaked);

}
