pragma solidity ^0.4.0;

import "./owned.sol";

contract SwearGameAbstract is Owned {

  /// @notice () - open a new case and add it to OpenCases
  ///
	///The function without name is the default function that is called whenever anyone sends funds to a contract
	/// It is used by the service for deposit
  function () payable;
	/// @notice register - register a player to the game
  ///
	/// The function will throw if the player is already register or there is not
	/// enough deposit in the contract to ensure the player could be compensated for the
	/// case of a valid case.
  /// @param _player  - the player address
  /// @return bool registered - true for success registration.
	function register(address _player) onlyOwner public returns (bool registered);
	/// @notice leaveGame - dismiss a player from the game (unregister)
  /// allow only plaintiff which do not have openCases on it name to leave game
  /// @param _player  - the player address
	function leaveGame(address _player);
	/// @notice getStatus - return the trial status of a case
  ///
	/// @param id  - case id
  /// @return  status  - the status of a case
  function getStatus(bytes32 id) public constant returns (uint8 status);
	/// @notice newCase - open a new case for a service id
  ///
	/// the function require that the msg sender is already register to the game.
  /// @param serviceId  - service id
	/// @return bool - true for succesfull operation.
	function newCase(bytes32 serviceId) public returns (bool);

	/// @notice trial - initiate or restart a trial proccess for a certian case
  ///
	/// the function requiere that the case is a valid one.
  /// @param id  - case id
	/// @return bool - true for succesfull operation.
  function trial(bytes32 id) public returns (bool);

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
