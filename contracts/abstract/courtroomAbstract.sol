pragma solidity ^0.4.0;

import "./owned.sol";
import "./sampletoken.sol";

contract CaseContractAbstract is Owned {

	struct claim {
		bytes32 claimId;
		address plaintiff;
		bytes32[] evidence;
		uint status;
		bool valid;
	}

	//claimid map to claim
  mapping(bytes32 => claim) public OpenClaims;
	function CaseContract();
	function newClaim(address _plaintiff, bytes32 _evidence) returns (bytes32 claimId);
	function submitEvidence(bytes32 _claimId,bytes32 _evident) returns (uint status);
	function getStatus(bytes32 claimId) constant returns (uint status);
	function resolveClaim(bytes32 _claimId);
	function getClaim(bytes32 _claimId) returns (address plaintiff,bool valid);
  function setClaimValid(bytes32 _claimId);

}

contract CourtroomAbstract is Owned {

	uint256 public amountStaked;
	uint public rewardCompensation;
	uint  public registeredPlayersCounter;
	mapping(address => bool) public registeredPlayers;
	mapping(address => bytes32[]) public clientsClaimsIds;

	CaseContractAbstract public caseContract;
	SampleToken public sampleToken;
	bytes32 public claimId;


	function SwearGame(address _caseContract, address _sampleToken, uint _rewardCompensation);
	function deposit(uint256 _depositAmount) onlyOwner payable public returns(bool);
	function makeJudgement(bytes32 _claimId) private returns(bool);
	function compensate(address _claimant) private returns(bool compensated);
	function register(address _player) onlyOwner public returns (bool registered);
	function leaveGame(address _player) onlyOwner public;
	function openNewClaim(bytes32 _evidence) public returns (bool);
	function takeDecision() private returns(bool);

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
