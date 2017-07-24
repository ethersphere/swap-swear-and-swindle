pragma solidity ^0.4.0;

import "./owned.sol";
import "./sampletoken.sol";
import "./abstracts/trialrulesabstract.sol";

contract SwearGame is Owned {

	uint256 public deposit;
	uint public reward;
	uint  public playerCount;
	SampleToken public token;
  TrialRulesAbstract public trialRules;
  struct Case {
    address plaintiff;
    bytes32 serviceId;
    uint8 status;
    uint8 valid;
  }
  //id map to _Case
  mapping(bytes32 => Case)  OpenCases;
  mapping(address => bool) public players;
	mapping(address => bytes32[]) public ids;

  function SwearGame(address _token, address _trialRules, uint _reward) {
		token        = SampleToken(_token);
    trialRules    = TrialRulesAbstract(_trialRules);
		reward = _reward;
		playerCount = 0;
	}

  function _newCase(address _plaintiff,bytes32 _serviceId,uint8 _status) private returns (bytes32 id) {
     id = sha3(_plaintiff,_serviceId, now);
     if (OpenCases[id].valid != 0)return 0x0;
     OpenCases[id] = Case(_plaintiff,_serviceId,_status,1);
    return id;
  }

  function isValid(bytes32 id) private constant returns (bool){
    return (OpenCases[id].valid != 0);
  }

  function setStatus(bytes32 id,uint8 status) private constant  returns (bool) {
    if (msg.sender == owner) throw;
   OpenCases[id].status = status;
   return true;
  }

  function resolveClaim(bytes32 _id) private {

    if (OpenCases[_id].status == uint8(TrialRulesAbstract.Status.UNCHALLENGED)) throw;
    OpenCases[_id].plaintiff = 0;
    OpenCases[_id].valid = 0;
    OpenCases[_id].status = uint8(TrialRulesAbstract.Status.UNCHALLENGED);
  }
  function getClaim(bytes32 _id) private returns (address plaintiff,bytes32 serviceId) {
          return (OpenCases[_id].plaintiff,OpenCases[_id].serviceId);
  }

  /* The function without name is the default function that is called whenever anyone sends funds to a contract */
  function () payable {
        uint amount = msg.value;
    		require(token.transferFrom(owner, address(this), amount));
    		deposit += amount;
    		DepositStaked(amount, deposit);
    }

	function verdict(bytes32 _id,uint8 status,address plaintiff) private returns(bool) {

    if (status == uint8(TrialRulesAbstract.Status.NOT_GUILTY)){
       ClaimResolved(_id, plaintiff, 0,status);
       return false;
     }

	  bool caseCompensated = compensate(plaintiff);
		resolveClaim(_id);
    _leaveGame(plaintiff);
		ClaimResolved(_id, plaintiff, reward,status);
		return caseCompensated;

	}

	function compensate(address _beneficiary) private returns(bool compensated) {

		compensated = token.transferFrom(address(this), _beneficiary, reward);

		require(compensated);

		deposit -=reward;

    Compensate(_beneficiary,reward);

		return compensated;

	}

	function register(address _player) onlyOwner public returns (bool registered) {

		require(!players[_player]);

		if (playerCount == 0){
			require(deposit >= reward);
		}else if ((deposit / playerCount) < reward) {
			AdditionalDepositRequired(deposit);
			throw;
		}

		players[_player] = true;
		playerCount++;

		NewPlayer(_player);

		return true;

	}
	function leaveGame(address _player) {

		for (uint256 i=0;i<ids[msg.sender].length;i++){ //allow only plaintiff wihich do not have openCases on it name to leave game
			require(OpenCases[ids[msg.sender][i]].valid == 0);
		}
		return _leaveGame(msg.sender);

  }

  function _leaveGame(address _player) private {

    require(players[_player]);

		PlayerLeftGame(_player);

		players[_player] = false;
		playerCount--;

  }

  function getStatus(bytes32 id) public constant returns (uint8 status){
    return OpenCases[id].status;
  }

	function newCase(bytes32 serviceId) public returns (bool) {

		require(players[msg.sender]);

		bytes32 id  = _newCase(msg.sender,serviceId,uint8(trialRules.getInitialStatus()));

    if (id == 0x0) return false;

		ids[msg.sender].push(id);

		return true;
	}

  function trial(bytes32 id) public returns (bool){
    require(players[msg.sender]);

    require(isValid(id));

    _trial(id);

		return true;
  }

 function proceed() private returns (WitnessAbstract.Status){
   return WitnessAbstract.Status.PENDING;
 }

 function _trial(bytes32 id) private{

  uint8 status =  getStatus(id);

  var(plaintiff,serviceId) = getClaim(id);

  if (status == uint8(TrialRulesAbstract.Status.UNCHALLENGED)) {
      return;
   }
   for (;status != uint8(TrialRulesAbstract.Status.UNCHALLENGED);) {

        WitnessAbstract witness = trialRules.getWitness(status);

        WitnessAbstract.Status outcome;
        if (witness == WitnessAbstract(0x0)) {
            outcome = proceed();
            return;
        } else {
					 bool expired = trialRules.expired(id,status);
					 if (witness.isEvidentSubmited(id,serviceId,plaintiff) && !expired){
					   outcome = witness.testimonyFor(id,serviceId,plaintiff);
					 }else{
						  if(trialRules.startGracePeriod(id,status)||(!expired)){
								 outcome = WitnessAbstract.Status.PENDING;
							}else{
							outcome = WitnessAbstract.Status.INVALID;
						}
					 }
       }
       if (outcome == WitnessAbstract.Status.PENDING) {
           return;
       }
      status = trialRules.getStatus(uint8(outcome),status);
      setStatus(id,status);
      if ((status == uint8(TrialRulesAbstract.Status.GUILTY))||
         (status == uint8(TrialRulesAbstract.Status.NOT_GUILTY))){
        verdict(id,status,plaintiff);
        status = uint8(TrialRulesAbstract.Status.UNCHALLENGED);
        setStatus(id,status);
      }
  }
}

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
