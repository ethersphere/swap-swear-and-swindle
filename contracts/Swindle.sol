pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;
import "./abstracts/AbstractTrialRules.sol";
import "./abstracts/AbstractWitness.sol";
import "./abstracts/AbstractConstants.sol";

/// @title Swindle contract
contract Swindle is AbstractConstants {

  // structure to keep track of trial
  struct Trial {    
    AbstractTrialRules rules; // trial contract    
    address payable plaintiff;
    address payable provider;
    bytes32 inputHash;
    bytes32 trialDataHash;
    uint64 lastAction; // timestamp of the last status change    
    uint8 status; // current status
  }  

  event TrialStarted(bytes32 indexed caseId, address provider, address plaintiff, AbstractTrialRules rules, bytes trialData);
  // fired whenever the state of a case changes
  event StateTransition(bytes32 indexed caseId, uint from, uint to, bytes trialData);

  // map from caseId to trial structure
  mapping (bytes32 => Trial) trials;

  function getTrialInfo(bytes32 caseId)
  public view returns (uint8 status, address rules, bytes32 inputHash) {
    Trial storage trial = trials[caseId];
    return (trial.status, address(trial.rules), trial.inputHash);
  }

  /// @dev start a trial, should be called from a Swear contract
  /// @param provider service provider
  /// @param plaintiff plaintiff, could be trial initiator or beneficiary in a Swap note
  /// @param rules to use
  /// @return the caseId of the new trial
  function startTrial(address payable provider, address payable plaintiff, AbstractTrialRules rules, bytes memory input)
  public returns (bytes32) {
    // derive a caseId, WARNING: horribly broken and insecure
    bytes32 caseId = keccak256(abi.encodePacked(msg.sender, provider, plaintiff, rules));

    (uint8 status, bytes memory trialData) = rules.getInitialStatus(input);

    trials[caseId] = Trial({
      rules: rules,
      plaintiff: plaintiff,
      provider: provider,
      inputHash: keccak256(input),
      status: status,
      trialDataHash: keccak256(trialData),
      lastAction: uint64(now)
    });

    emit TrialStarted(caseId, provider, plaintiff, rules, trialData);    

    return caseId;
  }

  function continueTrial(bytes32 caseId, bytes memory trialData, bytes memory input) public {
    Trial storage trial = trials[caseId];
    // if the trial is not going on or there is already a verdict abort
    require(trial.status > TRIAL_STATUS_NOT_GUILTY);
    require(keccak256(trialData) == trial.trialDataHash);

    // outcome will be written to this variable
    AbstractWitness.TestimonyStatus outcome;

    // get the next step from the rules
    (address witness, uint expiry) = trial.rules.getWitness(trial.status);

    if(now - trial.lastAction > expiry) {      
      // if too much time has passed assume the testimony to be PENDING
      outcome = AbstractWitness.TestimonyStatus.PENDING;
    } else {
      bytes memory specification = trial.rules.getWitnessPayload(trial.status, trial.provider, trial.plaintiff, trialData);
      bytes memory kv;      
      
      (outcome, kv) = AbstractWitness(witness).testimonyFor(specification, input);      
      require(outcome != AbstractWitness.TestimonyStatus.PENDING, "still pending");      

      trialData = trial.rules.updateData(outcome, trial.status, trialData, kv);
      bytes32 trialDataHash = keccak256(trialData);
      if(trialDataHash != trial.trialDataHash) trial.trialDataHash = trialDataHash;
    }
      
    // get the next status from the rules, different variable because we still need the old status in the next line
    uint8 next = trial.rules.nextStatus(outcome, trial.status);

    emit StateTransition(caseId, trial.status, next, trialData);
    // status change so we need to update the lastAction timestamp
    trial.lastAction = uint64(now);
    // update the status
    trial.status = next;
  }
}
