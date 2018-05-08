pragma solidity ^0.4.19;
import "./Swear.sol";
import "./abstracts/AbstractRules.sol";
import "./abstracts/AbstractWitness.sol";
import "./abstracts/AbstractConstants.sol";

contract Swindle is AbstractConstants {

  struct Trial {
    Swear swear;
    AbstractRules rules;
    address plaintiff;
    address provider;
    bytes32 noteId;
    bytes32 commitmentHash;
    uint lastAction;
    uint8 status;
  }

  event StateTransition(bytes32 indexed caseId, uint from, uint to);


  mapping (bytes32 => Trial) trials;

  function startTrial(address provider, address plaintiff, bytes32 noteId, bytes32 commitmentHash, AbstractRules rules) public returns (bytes32) {
    bytes32 caseId = keccak256(provider, plaintiff, noteId);

    trials[caseId] = Trial({
      swear: Swear(msg.sender),
      rules: rules,
      plaintiff: plaintiff,
      provider: provider,
      noteId: noteId,
      commitmentHash: commitmentHash,
      status: rules.getInitialStatus() /* TODO: re-entrance */,
      lastAction: now
    });

    return caseId;
  }

  function continueTrial(bytes32 caseId) public {
    Trial storage trial = trials[caseId];
    if(trial.status <= TRIAL_STATUS_NOT_GUILTY) return;

    AbstractWitness.TestimonyStatus outcome;

    address witness;
    uint expiry;
    (witness, expiry) = trial.rules.getWitness(trial.status);

    if(now - trial.lastAction > expiry) {
      outcome = AbstractWitness.TestimonyStatus.INVALID;
    } else {
      /* TODO: re-entrance */
      outcome = AbstractWitness(witness).testimonyFor(trial.provider, trial.plaintiff, trial.noteId);
    }

    if(outcome == AbstractWitness.TestimonyStatus.PENDING) return;

    trial.lastAction = now;
    uint8 next = trial.rules.nextStatus(outcome, trial.status);

    StateTransition(caseId, trial.status, next);

    trial.status = next;
  }

  function endTrial(bytes32 caseId) public {
    Trial storage trial = trials[caseId];

    if(trial.status == TRIAL_STATUS_NOT_GUILTY) {

    } else if(trial.status == TRIAL_STATUS_GUILTY) {
      trial.swear.compensate(trial.commitmentHash, trial.plaintiff, trial.rules.getDeposit());
    } else revert();

    trial.swear.notifyTrialEnd(trial.commitmentHash);

    trial.status = 0;
  }

}
