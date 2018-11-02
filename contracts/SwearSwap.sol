pragma solidity ^0.4.23;
import "./abstracts/AbstractRules.sol";
import "./abstracts/AbstractWitness.sol";
import "./abstracts/AbstractSwear.sol";
import "./Swindle.sol";
import "./SW3Utils.sol";
import "openzeppelin-solidity/contracts/math/SafeMath.sol";

/// @title Swear Contract
contract SwearSwap is SW3Utils, AbstractWitness, AbstractSwear {
  using SafeMath for uint;

  /* fired when a trial with swindle is started */
  event TrialStarted(bytes32 commitmentHash, bytes32 caseId, bytes32 noteId);

  Swindle public swindle;

  /// @notice constructor, allows setting the swindle
  constructor(address _swindle) public {
    swindle = Swindle(_swindle);
  }

  /* structure for a commitment */
  struct Commitment {
    bool valid; /* indicates wether this structure is valid */
    address provider; /* provider of the service */
    bytes32 noteId;
    uint cases; /* number of open cases */
  }

  /* associates commitmentHash with the commitment */
  mapping (bytes32 => Commitment) public commitments;

  /// @notice callback for swindle when compensation should take place
  /// @dev either reduces the deposit if onchain or mark note as valid if offchain
  /// @param commitmentHash commitment to compensate from
  function compensate(bytes32 commitmentHash, address, uint) public {
    require(msg.sender == address(swindle));
    Commitment storage commitment = commitments[commitmentHash];
    guiltyNotes[commitment.provider][commitment.noteId] = true;
  }

  /// @notice callback for swindle at the end of the trial
  /// @param commitmentHash commitment
  function notifyTrialEnd(bytes32 commitmentHash) public {
    require(msg.sender == address(swindle));
    Commitment storage commitment = commitments[commitmentHash];
    /* reduce the number of cases */
    commitment.cases--;
  }

  /* mapping from provider and noteId to wether the noteId has been marked as guilty */
  mapping (address => mapping (bytes32 => bool)) public guiltyNotes;

  /// @notice witness implementation of swear
  function testimonyFor(address owner, address , bytes32 noteId) public view returns (TestimonyStatus) {
    return guiltyNotes[owner][noteId] ? AbstractWitness.TestimonyStatus.VALID : AbstractWitness.TestimonyStatus.INVALID;
  }

  /// @notice start a trial from a SWAP note
  /// @param trial trial rules for the note (needs to match the remark)
  /// @param payload payload (needs to match the remark)
  /// @param sig signature of the note
  function startTrialFromNote(bytes encoded, address trial, bytes32 payload, bytes sig) public returns(address) {
    Note memory note = decodeNote(encoded);

    /* get the provider from the signature */
    address provider = recover(note.id, sig);
    bytes32 commitmentHash = keccak256(abi.encodePacked(provider, trial, note.id));

    /* ensure that trial and payload match the remark */
    require(keccak256(abi.encodePacked(trial, payload)) == note.remark);

    if(!commitments[commitmentHash].valid) {
      /* store the commitment */
      commitments[commitmentHash] = Commitment({
        valid: true,
        provider: provider,
        noteId: note.id,
        cases: 1 /* initialize with 1 open case */
      });
    } else {
      commitments[commitmentHash].cases++;
    }

    /* initiate the swindle trial, swindle will call back once its over */
    bytes32 caseId = swindle.startTrial(provider, note.beneficiary, payload, commitmentHash, AbstractRules(trial));
    emit TrialStarted(commitmentHash, caseId, note.id);
  }
}
