pragma solidity ^0.4.0;

import "./owned.sol";
import "./sampletoken.sol";
import "./abstracts/trialtransitionsabstract.sol";


contract SwearGameAbstract is Owned {
    uint256 public deposit;
	uint public reward;
	uint  public playerCount;
	mapping(address => bool) public players;
	mapping(address => bytes32[]) public ids;

	CaseContract public caseContract;
	SampleToken public token;
    TrialTransistionsAbstract public trialTransistions;

	function SwearGame(address _CaseContract, address _token, address _trialTransistions, uint _reward);
	function register(address _player) onlyOwner public returns (bool registered);
    function getStatus(bytes32 id) public constant returns (uint8);
	function newCase(bytes32 serviceId) public returns (bool);
    function trial(bytes32 id) public returns (bool);

	event Decision(string decide);
	event DepositStaked(uint depositAmount, uint deposit);
	event Compensate(address recipient, uint reward);
	event NewPlayer(address playerId);
	event PlayerLeftGame(address playerId);
	event NewClaimOpened(bytes32 id, address plaintiff);
	event NewEvidenceSubmitted(bytes32 id, address plaintiff);
	event ClaimResolved(bytes32 id, address plaintiff, uint reward,uint8 status);
	event Payment(address from,address to ,uint256 value);
	event AdditionalDepositRequired(uint256 deposit);

}
