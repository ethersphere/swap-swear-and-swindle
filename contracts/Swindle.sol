pragma solidity ^0.5.0;
import "./abstracts/AbstractSwear.sol";
import "./abstracts/AbstractRules.sol";
import "./abstracts/AbstractWitness.sol";
import "./abstracts/AbstractConstants.sol";

/// @title Swindle contract
contract Swindle is AbstractConstants {

  /* structure to keep track of trial */
  struct Trial {
    AbstractSwear swear; /* address that initiated this trial, should implement the Swear interface */
    AbstractRules rules; /* rules contract */
    address payable plaintiff;
    address provider;
    bytes32 noteId;
    bytes32 commitmentHash; /* commitmentHash from Swear */
    uint lastAction; /* timestamp of the last status change */
    uint8 status; /* current status */
  }

  /* fired whenever the state of a case changes */
  event StateTransition(bytes32 indexed caseId, uint from, uint to);

  /* map from caseId to trial structure */
  mapping (bytes32 => Trial) trials;

  /// @dev start a trial, should be called from a Swear contract
  /// @param provider service provider
  /// @param plaintiff plaintiff, could be trial initiator or beneficiary in a Swap note
  /// @param noteId data for the witnesses
  /// @param commitmentHash hash to identify the commitment with Swear
  /// @param rules to use
  /// @return the caseId of the new trial
  function startTrial(address provider, address payable plaintiff, bytes32 noteId, bytes32 commitmentHash, AbstractRules rules) public returns (bytes32) {
    /* derive a caseId, WARNING: horribly broken and insecure */
    bytes32 caseId = keccak256(abi.encodePacked(provider, plaintiff, noteId));

    trials[caseId] = Trial({
      swear: AbstractSwear(msg.sender),
      rules: rules,
      plaintiff: plaintiff,
      provider: provider,
      noteId: noteId,
      commitmentHash: commitmentHash,
      status: rules.getInitialStatus() /* TODO: should be STATIC_CALL, Solidity 0.5 */,
      lastAction: now
    });

    return caseId;
  }

  /// @notice try to advance the trial by one step
  /// @param caseId case to proceed
  function continueTrial(bytes32 caseId) public {
    Trial storage trial = trials[caseId];
    /* if the trial is not going on or there is already a verdict abort */
    if(trial.status <= TRIAL_STATUS_NOT_GUILTY) return;

    /* outcome will be written to this variable */
    AbstractWitness.TestimonyStatus outcome;

    /* get the next step from the rules */
    (address witness, uint expiry) = trial.rules.getWitness(trial.status);

    if(now - trial.lastAction > expiry) {
      /* if too much time has passed assume the testimony to be INVALID */
      outcome = AbstractWitness.TestimonyStatus.INVALID;
    } else {
      /* TODO: STATIC_CALL, Solidity 0.5 */
      outcome = AbstractWitness(witness).testimonyFor(trial.provider, trial.plaintiff, trial.noteId);
    }

    /* if the outcode is still PENDING abort */
    if(outcome == AbstractWitness.TestimonyStatus.PENDING) return;

    /* status change so we need to update the lastAction timestamp */
    trial.lastAction = now;

    /* get the next status from the rules, different variable because we still need the old status in the next line */
    uint8 next = trial.rules.nextStatus(outcome, trial.status);

    emit StateTransition(caseId, trial.status, next);
    /* update the status */
    trial.status = next;
  }

  /// @notice end a trial (there needs to be a verdict already)
  /// @param caseId case to end
  function endTrial(bytes32 caseId) public {
    Trial storage trial = trials[caseId];

    if(trial.status == TRIAL_STATUS_NOT_GUILTY) {
      /* no special code for a not guilty verdict for now */
    } else if(trial.status == TRIAL_STATUS_GUILTY) {
      /* if GUILTY instruct Swear to compensate the plaintiff with the entire deposit */
      trial.swear.compensate(trial.commitmentHash, trial.plaintiff, trial.rules.getDeposit());
    } else revert(); /* revert if we are not at a verdict or the trial is invalid */
    /* invalidate the trial */
    trial.status = 0;
    /* notify Swear of trial end regardless of verdict */
    trial.swear.notifyTrialEnd(trial.commitmentHash);
  }

}
