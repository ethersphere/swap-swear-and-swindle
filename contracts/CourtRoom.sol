pragma solidity ^0.4.0;

import "./abstracts/CourtroomAbstract.sol";
import "./sampletoken.sol";
import "./abstracts/trialrulesabstract.sol";


contract SwearGame is SwearGameAbstract {

	uint256 public deposit;
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

	/// @notice SwearGame - Swear game constructor this function is called along with
	/// the contract deployment time.
  ///
  /// @param _token address of the token contract
	/// @param _trialRules - address of the trial specific rules contract
  /// @return WitnessAbstract - return a witness contract instance
  function SwearGame(address _token, address _trialRules) {
		token        = SampleToken(_token);
    trialRules    = TrialRulesAbstract(_trialRules);
		playerCount = 0;
	}
	/// @notice _newCase - open a new case and add it to OpenCases
  ///
  /// @param _plaintiff  - the plaintiff address for the case
	/// @param _serviceId - service id related to the case
  /// @return _status - the status of the case
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

  function resolveCase(bytes32 _id) private {

    if (OpenCases[_id].status == uint8(TrialRulesAbstract.Status.UNCHALLENGED)) throw;
    OpenCases[_id].plaintiff = 0;
    OpenCases[_id].valid = 0;
    OpenCases[_id].status = uint8(TrialRulesAbstract.Status.UNCHALLENGED);
  }
  function getCase(bytes32 _id) private returns (address plaintiff,bytes32 serviceId) {
          return (OpenCases[_id].plaintiff,OpenCases[_id].serviceId);
  }

	/// @notice () - a payable function which is used by the service as a deposit functionality
  ///
	///The function without name is the default function that is called whenever anyone sends funds to a contract
	/// It is used by the service for deposit
  function () payable {
        uint amount = msg.value;
    		require(token.transferFrom(owner, address(this), amount));
    		deposit += amount;
    		DepositStaked(amount, deposit);
    }

	function verdict(bytes32 _id,uint8 status,address plaintiff) private returns(bool) {

    if (status == uint8(TrialRulesAbstract.Status.NOT_GUILTY)){
       CaseResolved(_id, plaintiff, 0,status);
       return false;
     }
    uint reward = trialRules.getReward();
	  bool caseCompensated = compensate(plaintiff,reward);
		resolveCase(_id);
    _leaveGame(plaintiff);
		CaseResolved(_id, plaintiff, reward,status);
		return caseCompensated;

	}

	function compensate(address _beneficiary,uint reward) private returns(bool compensated) {


		compensated = token.transferFrom(address(this), _beneficiary, reward);

		require(compensated);

		deposit -=reward;

    Compensate(_beneficiary,reward);

		return compensated;

	}
	/// @notice register - register a player to the game
  ///
	/// The function will throw if the player is already register or there is not
	/// enough deposit in the contract to ensure the player could be compensated for the
	/// case of a valid case.
  /// @param _player  - the player address
  /// @return bool registered - true for success registration.
	function register(address _player) onlyOwner public returns (bool registered) {

		require(!players[_player]);

		uint reward = trialRules.getReward();

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
	/// @notice leaveGame - dismiss a player from the game (unregister)
  /// allow only plaintiff which do not have openCases on it name to leave game
  /// @param _player  - the player address
	function leaveGame(address _player) {

		for (uint256 i=0;i<ids[msg.sender].length;i++){ //allow only plaintiff which do not have openCases on it name to leave game
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
	/// @notice getStatus - return the trial status of a case
  ///
	/// @param id  - case id
  /// @return  status  - the status of a case
  function getStatus(bytes32 id) public constant returns (uint8 status){
    return OpenCases[id].status;
  }
	/// @notice newCase - open a new case for a service id
  ///
	/// the function require that the msg sender is already register to the game.
  /// @param serviceId  - service id
	/// @return bool - true for successful operation.
	function newCase(bytes32 serviceId) public returns (bool) {

		require(players[msg.sender]);

		bytes32 id  = _newCase(msg.sender,serviceId,uint8(trialRules.getInitialStatus()));

    if (id == 0x0) return false;

		ids[msg.sender].push(id);

		return true;
	}

	/// @notice trial - initiate or restart a trial proccess for a certain case
  ///
	/// the function requiere that the case is a valid one.
  /// @param id  - case id
	/// @return bool - true for successful operation.
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

  var(plaintiff,serviceId) = getCase(id);

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
	event NewCaseOpened(bytes32 id, address plaintiff);
	event NewEvidenceSubmitted(bytes32 id, address plaintiff);
	event CaseResolved(bytes32 id, address plaintiff, uint reward,uint8 status);
	event Payment(address from,address to ,uint256 value);
	event AdditionalDepositRequired(uint256 deposit);

}
