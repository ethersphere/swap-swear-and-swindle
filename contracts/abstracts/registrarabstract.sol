pragma solidity ^0.4.0;

import "./trialrulesabstract.sol";
import "./token.sol";
import "./owned.sol";


contract RegistrarAbstract is Owned {


    /// @notice register - register a player to the game
    ///
    /// The function will throw if the player is already register or there is not
    /// enough deposit in the contract to ensure the player could be compensated for the
    /// case of a valid case.
    /// @param _player  - the player address
    /// @return bool registered - true for success registration.
    function register(address _player) onlyOwner public returns (bool);

    /// @notice deposit - deposit to the contract
    ///
    /// @param epochs  - deposit epochs
    /// @return bool registered - true for success deposit.
    function deposit(uint epochs) payable returns (bool);

    /// @notice collectDeposit - collect back the caller deposit.
    ///
    ///The function check that there is no open case for the specific caller.
    ///The function check that there is enough deposit left in thec contract for the case of a valid case compensation.
    /// @return bool  true for success otherwise false.
    function collectDeposit() external returns (bool);

    /// @notice isRegistered - Check if a player is registered to the game
    ///
    /// @param player  - player address
    /// @return bool  true for success otherwise false.
    function isRegistered(address player) returns (bool);

    /// @notice compensate - compensate the beneficiary with the reward amount
    ///
    /// @param _beneficiary  - beneficiary address
    /// @param reward        - reward amount
    /// The function will throw if it is not called by the swear contract.
    /// @return bool  true for success otherwise false.
    function compensate(address _beneficiary,uint reward) returns(bool compensated);

    /// @notice unRegister - un register a player
    ///
    /// The function will throw if it is not called by the swear contract.
    /// @param _player  - player address
    function unRegister(address _player);

    /// @notice setSwearContractAddress - set the swear contract address
    ///
    /// The function will enable to set the swear contract address only one time.otherwise
    /// it will throw.
    /// @param _swearAddress  - swear contract address
    /// @return bool  true for success otherwise false.
    function setSwearContractAddress(address _swearAddress) returns(bool);

    /// @notice incrementOpenCases - increment open cases number for a specific address
    ///
    /// The function will throw if it is not called by the swear contract.
    /// @param _address  - address to which to increment open cases
    function incrementOpenCases(address _address);

    /// @notice decrementOpenCases - decrement open cases number for a specific address
    ///
    /// The function will throw if it is not called by the swear contract.
    /// @param _address  - address to which to decrement open cases
    function decrementOpenCases(address _address);
}
