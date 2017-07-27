pragma solidity ^0.4.0;
import "./witnessabstract.sol";


contract TrialRulesAbstract {


    enum Status {UNCHALLENGED,GUILTY,NOT_GUILTY}
    /// @notice getStatus - get next trial status according to witness state and the current trial state
    ///
    /// @param witnessStatus witness status (VALID , INVALID,PENDING)
    /// @param trialStatus current trial status
    /// @return status - next trial status - can be also GUILTY or NOT GUILTY.
    function getStatus(uint8 witnessStatus,uint8 trialStatus) returns (uint8 status);

    /// @notice getWitness - get witness according to the trial status
    ///
    /// @param trialStatus current trial status
    /// @return WitnessAbstract - return a witness contract instance
    function getWitness(uint8 trialStatus) returns (WitnessAbstract);

    /// @notice getInitialStatus - get initial trial status
    ///
    /// @return status -
    function getInitialStatus() public returns (uint8 status);

    /// @notice expired - check expiration for a certain case and trial status
    ///
    /// @return bool - true if expired otherwise false
    function expired(bytes32 caseId,uint8 status) returns (bool);

    /// @notice startGracePeriod - start counting for a grace period for a certain case and status.
    ///
    /// @return bool - true if it actually start counting for the grace period
    ///                false -if the grace period already started
    function startGracePeriod(bytes32 caseId,uint8 status) returns (bool);

    /// @notice getReward - return the reward for a valid case
    ///
    /// @return reward - the reward for a valid case
    function getReward() constant returns (uint reward);

}
