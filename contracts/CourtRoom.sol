pragma solidity ^0.4.0;

import "./owned.sol";
import "./sampletoken.sol";
import "./abstract/trialtransitionsabstract.sol";


contract CaseContract is Owned {

	struct Case {
		bytes32 id;
		address plaintiff;
    bytes32 serviceId;
		uint8 status;
    uint expiery;
    uint8 valid;
	}

	//id map to _Case
  mapping(bytes32 => Case)  OpenCases;

	function CaseContract() {
	}

	function newClaim(address _plaintiff,bytes32 _serviceId,uint8 status,uint expiery) returns (bytes32 id) {

	   Case memory _case;
		_case.id = sha3(_plaintiff,_serviceId, now);
     if (OpenCases[_case.id].valid != 0)return 0x0;
		_case.plaintiff = _plaintiff;
    _case.serviceId = _serviceId;
    _case.status = status;
		_case.valid = 1;
    _case.expiery = now + expiery;
		 OpenCases[_case.id] = _case;
		return _case.id;
	}

  function isValid(bytes32 id) constant returns (bool){
    return (OpenCases[id].valid != 0);
  }

  function getStatus(bytes32 id) constant returns (uint8 status) {
		return OpenCases[id].status;
	}

  function setStatus(bytes32 id,uint8 status) constant returns (bool) {
	 OpenCases[id].status = status;
   return true;
	}

	function resolveClaim(bytes32 _id) {

		if (OpenCases[_id].id == 0x0 || OpenCases[_id].status == 0) throw;
		OpenCases[_id].id = 0x0;
		OpenCases[_id].plaintiff = 0;
    OpenCases[_id].valid = 0;
		OpenCases[_id].status = uint8(TrialTransistionsAbstract.Status.UNCHALLENGED);
	}
  function getClaim(bytes32 _id) returns (address plaintiff,bytes32 serviceId,uint256 expiery) {
          return (OpenCases[_id].plaintiff,OpenCases[_id].serviceId,OpenCases[_id].expiery);
  }
}

contract SwearGame is Owned {

	uint256 public deposit;
	uint public reward;
	uint  public playerCount;
	mapping(address => bool) public players;
	mapping(address => bytes32[]) public ids;

	CaseContract public caseContract;
	SampleToken public token;
  TrialTransistionsAbstract public trialTransistions;

	function SwearGame(address _CaseContract, address _token, address _trialTransistions, uint _reward) {
		caseContract = CaseContract(_CaseContract);
		token        = SampleToken(_token);
    trialTransistions    = TrialTransistionsAbstract(_trialTransistions);
		reward = _reward;
		playerCount = 0;
	}


  /* The function without name is the default function that is called whenever anyone sends funds to a contract */
  function () payable {
        uint amount = msg.value;
    		require(token.transferFrom(owner, address(this), amount));
    		deposit += amount;
    		DepositStaked(amount, deposit);
    }

	function verdict(bytes32 _id,uint8 status,address plaintiff) private returns(bool) {

    if (status == uint8(TrialTransistionsAbstract.Status.NOT_GUILTY)){
       ClaimResolved(_id, plaintiff, 0,status);
       return false;
     }

	  bool caseCompensated = compensate(plaintiff);
		caseContract.resolveClaim(_id);
    leaveGame(plaintiff);
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

  function leaveGame(address _player) private {
    require(players[_player]);

		PlayerLeftGame(_player);

		players[_player] = false;
		playerCount--;
  }

  function getStatus(bytes32 id) public constant returns (uint8 status){
    return caseContract.getStatus(id);
  }

	function newCase(bytes32 serviceId) public returns (bool) {

		require(players[msg.sender]);

		bytes32 id  = caseContract.newClaim(msg.sender,serviceId,uint8(trialTransistions.getInitialStatus()),uint(trialTransistions.getTrialExpiry()));

    if (id == 0x0) return false;

		ids[msg.sender].push(id);

    trial(id);

		return true;
	}

  function resumeCase(bytes32 id) public returns (bool) {

		require(players[msg.sender]);

    require(caseContract.isValid(id));

    trial(id);

		return true;
	}

 function expired(uint256 expiry) private returns(bool){
   return (now  >  expiry);
  }

 function proceed() private returns (WitnessAbstract.Status){
   return WitnessAbstract.Status.PENDING;

 }

 function trial(bytes32 id) private{

  uint8 status =   caseContract.getStatus(id);


  var(plaintiff,serviceId,expiery) = caseContract.getClaim(id);

  if (status == uint8(TrialTransistionsAbstract.Status.UNCHALLENGED)) {
      return;
   }
   for (;status != uint8(TrialTransistionsAbstract.Status.UNCHALLENGED);) {

        if (expired(expiery)) {
             caseContract.setStatus(id, uint8(TrialTransistionsAbstract.Status.UNCHALLENGED));
             return;
        }

        WitnessAbstract witness = trialTransistions.getWitness(status);


        WitnessAbstract.Status outcome;
        if (witness == WitnessAbstract(0x0)) {
            outcome = proceed();
            return;
        } else {

           outcome = witness.testimonyFor(serviceId,plaintiff);

       }
       if (outcome == WitnessAbstract.Status.PENDING) {
           return;
      }
      status = trialTransistions.getStatus(uint8(outcome),status);
      caseContract.setStatus(id,status);
      if ((status == uint8(TrialTransistionsAbstract.Status.GUILTY))||
         (status == uint8(TrialTransistionsAbstract.Status.NOT_GUILTY))){
        verdict(id,status,plaintiff);
        status = uint8(TrialTransistionsAbstract.Status.UNCHALLENGED);
        caseContract.setStatus(id,status);
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
