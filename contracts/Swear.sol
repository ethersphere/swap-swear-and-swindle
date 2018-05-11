pragma solidity ^0.4.19;
import "./abstracts/AbstractRules.sol";
import "./abstracts/AbstractWitness.sol";
import "./Swindle.sol";
import "./SW3Utils.sol";
import "zeppelin/math/SafeMath.sol";

contract Swear is SW3Utils, AbstractWitness {
  using SafeMath for uint;

  event CommitmentAdded(bytes32 commitmentHash, address indexed provider, address rules);
  event TrialStarted(bytes32 commitmentHash, bytes32 caseId, bytes32 noteId);

  Swindle public swindle;

  constructor(address _swindle) public {
    swindle = Swindle(_swindle);
  }

  struct Commitment {
    bool valid;
    address provider;

    uint deposit;
    uint timeout;
    AbstractRules rules;
    bytes32 noteId;

    bool note;

    uint cases;
  }

  mapping (bytes32 => Commitment) public commitments;


  function addCommitment(address rules, uint timeout, bytes32 noteId) public payable {
    require(msg.value >= AbstractRules(rules).getDeposit());
    require(timeout >= now + AbstractRules(rules).getEpoch());

    bytes32 commitmentHash = keccak256(msg.sender, rules, noteId);

    require(!commitments[commitmentHash].valid);

    commitments[commitmentHash] = Commitment({
      valid: true,
      provider: msg.sender,
      deposit: msg.value,
      rules: AbstractRules(rules),
      noteId: noteId,
      timeout: timeout,
      cases: 0,
      note: false
    });

    emit CommitmentAdded(commitmentHash, msg.sender, rules);
  }
  function compensate(bytes32 commitmentHash, address beneficiary, uint reward) public {
    require(msg.sender == address(swindle));
    Commitment storage commitment = commitments[commitmentHash];
    if(!commitment.note) {
      commitment.deposit = commitment.deposit.sub(reward);
      beneficiary.transfer(reward);
    } else {
      guiltyNotes[commitment.provider][commitment.noteId] = true;
    }
  }

  function withdraw(bytes32 commitmentHash) public {
    Commitment storage commitment = commitments[commitmentHash];

    require(msg.sender == commitment.provider);
    require(now > commitment.timeout + commitment.rules.getEpoch());
    require(commitment.cases == 0);

    msg.sender.transfer(commitment.deposit);

    commitment.valid = false;
  }

  function startTrial(bytes32 commitmentHash) public {
    Commitment storage commitment = commitments[commitmentHash];
    require(commitment.valid);

    address plaintiff = msg.sender;
    address provider = commitment.provider;

    commitment.cases++;

    bytes32 caseId = swindle.startTrial(provider, plaintiff, commitment.noteId, commitmentHash, commitment.rules);

    emit TrialStarted(commitmentHash, caseId, commitment.noteId);
  }

  function notifyTrialEnd(bytes32 commitmentHash) public {
    require(msg.sender == address(swindle));
    Commitment storage commitment = commitments[commitmentHash];

    commitment.cases--;
  }

  /* SWAP */
  mapping (address => mapping (bytes32 => bool)) public guiltyNotes;

  function noteGuilty(address provider, bytes32 noteId) public view returns(bool) {
    return guiltyNotes[provider][noteId];
  }

  function testimonyFor(address owner, address beneficiary, bytes32 noteId) public view returns (TestimonyStatus) {
    return guiltyNotes[owner][noteId] ? AbstractWitness.TestimonyStatus.VALID : AbstractWitness.TestimonyStatus.INVALID;
  }

  function startTrialFromNote(bytes note, address trial, bytes32 payload, bytes sig) public returns(address) {
    bytes32 noteId = keccak256(note);
    bytes32 commitmentHash = keccak256(provider, trial, noteId);

    address provider = recoverSignature(noteId, sig);

    address beneficiary;
    bytes32 remark;

    (,beneficiary,,,,,,remark) = decodeNote(note);

    require(keccak256(abi.encodePacked(trial, payload)) == remark);

    commitments[commitmentHash] = Commitment({
      valid: true,
      provider: provider,
      deposit: 0,
      rules: AbstractRules(trial),
      noteId: noteId,
      timeout: 0,
      cases: 1,
      note: true
    });

    bytes32 caseId = swindle.startTrial(provider, beneficiary, noteId, commitmentHash, AbstractRules(trial));
    emit TrialStarted(commitmentHash, caseId, noteId);
  }
}
