pragma solidity ^0.4.0;

import "./owned.sol";


contract SwearAbstract is Owned {

    /// @notice getStatus - return the trial status of a case
    ///
    /// @param id  - case id
    /// @return  status  - the status of a case
    function getStatus(bytes32 id) public constant returns (uint8 status);

    /// @notice newCase - open a new case for a service id
    ///
    /// the function require that the msg sender is already register to the game.
    /// @param serviceId  - service id
    /// @return bool - true for successful operation.
    function newCase(bytes32 serviceId) public returns (bool);

    /// @notice trial - initiate or restart a trial proccess for a certain case
    ///
    /// the function requiere that the case is a valid one.
    /// @param id  - case id
    /// @return bool - true for successful operation.
    function trial(bytes32 id) public returns (bool);

    event Decision(string decide);
    event NewCaseOpened(bytes32 id, address plaintiff);
    event NewEvidenceSubmitted(bytes32 id, address plaintiff);
    event CaseResolved(bytes32 id, address plaintiff, uint reward,uint8 status);
}
