pragma solidity ^0.4.23;
import "./abstracts/AbstractRules.sol";
import "./abstracts/AbstractWitness.sol";
import "./Swindle.sol";
import "./SW3Utils.sol";
import "openzeppelin-solidity/contracts/math/SafeMath.sol";

/// @title Swear Contract
contract Swear is SW3Utils, AbstractWitness {
  using SafeMath for uint;

  /* fired when an onchain commitment is added */
  event CommitmentAdded(bytes32 commitmentHash, address indexed provider, address rules);
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

    uint deposit; /* amount that was deposited into this contract or 0 for Swap based commitments */
    uint timeout; /* end of the service or 0 for Swap based commitments  */
    AbstractRules rules; /* rules of the game */
    bytes32 payload;
    bytes32 noteId; /* noteId to be passed to the witness, arbitrary data in onchain case otherwise hash of the Swap note */

    uint cases; /* number of open cases */
  }

  /* associates commitmentHash with the commitment */
  mapping (bytes32 => Commitment) public commitments;

  /// @notice add an onchain commitment (sent amount needs to be according to the rules)
  /// @param rules rules contract to use
  /// @param timeout end time for the service (needs to be according to the rules)
  /// @param payload metadata for the witnesses (arbitrary 32 bytes)
  function addCommitment(address rules, uint timeout, bytes32 payload) public payable {
    /* check enough ether were sent */
    require(msg.value >= AbstractRules(rules).getDeposit());
    /* check the timeout satisfies the rules */
    require(timeout >= now + AbstractRules(rules).getEpoch());

    /* compute the commitmentHash identifying this commitment */
    bytes32 commitmentHash = keccak256(msg.sender, rules, payload);

    /* make sure the same commitment has not happened before */
    /* TODO: commitmentHash should probably include more things */
    require(!commitments[commitmentHash].valid);

    /* store the commitment information */
    commitments[commitmentHash] = Commitment({
      valid: true,
      provider: msg.sender,
      deposit: msg.value,
      rules: AbstractRules(rules),
      noteId: 0,
      payload: payload,
      timeout: timeout,
      cases: 0 /* there are no open cases in the beginning */
    });

    emit CommitmentAdded(commitmentHash, msg.sender, rules);
  }

  /// @notice callback for swindle when compensation should take place
  /// @dev either reduces the deposit if onchain or mark note as valid if offchain
  /// @param commitmentHash commitment to compensate from
  /// @param beneficiary beneficiary to compensate
  /// @param reward amount to be compensated
  function compensate(bytes32 commitmentHash, address beneficiary, uint reward) public {
    require(msg.sender == address(swindle));
    Commitment storage commitment = commitments[commitmentHash];
    if(commitment.noteId == 0) {
      /* if this is an onchain commitment, reduce the deposit and transfer the compensation */
      commitment.deposit = commitment.deposit.sub(reward);
      beneficiary.transfer(reward);
    } else {
      /* if this is an offchain commitment mark the note as valid */
      guiltyNotes[commitment.provider][commitment.noteId] = true;
    }
  }

  /// @notice withdraw the deposit from a commitment
  /// @param commitmentHash commitment to withdraw
  function withdraw(bytes32 commitmentHash) public {
    Commitment storage commitment = commitments[commitmentHash];
    /* ensure commitment is (still) valid */
    require(commitment.valid);
    /* only the provider can do this */
    require(msg.sender == commitment.provider);
    /* make sure the service period is over */
    require(now > commitment.timeout + commitment.rules.getEpoch());
    /* make sure there are no open cases */
    require(commitment.cases == 0);

    /* send out commitment */
    msg.sender.transfer(commitment.deposit);

    /* mark commitment as invalid */
    commitment.valid = false;
  }

  /// @notice start trial for an onchain commitment
  /// @param commitmentHash commitment
  function startTrial(bytes32 commitmentHash) public {
    Commitment storage commitment = commitments[commitmentHash];
    /* ensure commitment is (still) valid */
    require(commitment.valid);

    /* plaintiff is the sender, WARNING: plaintiff gets the reward, there should probably be a beneficiary associated with the commitment */
    address plaintiff = msg.sender;
    address provider = commitment.provider;

    /* increase number of cases */
    commitment.cases++;

    /* initiate the swindle trial, swindle will call back once its over */
    bytes32 caseId = swindle.startTrial(provider, plaintiff, commitment.payload, commitmentHash, commitment.rules);

    emit TrialStarted(commitmentHash, caseId, commitment.payload);
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
  function startTrialFromNote(bytes note, address trial, bytes32 payload, bytes sig) public returns(address) {
    bytes32 noteId = keccak256(note);
    bytes32 commitmentHash = keccak256(provider, trial, noteId);

    /* get the provider from the signature */
    address provider = recoverSignature(noteId, sig);

    address beneficiary;
    bytes32 remark;

    /* get beneficiary and remark from the note */
    (,beneficiary,,,,,,remark) = decodeNote(note);

    /* ensure that trial and payload match the remark */
    require(keccak256(abi.encodePacked(trial, payload)) == remark);

    /* store the commitment */
    commitments[commitmentHash] = Commitment({
      valid: true,
      provider: provider,
      deposit: 0, /* handled by Swap */
      rules: AbstractRules(trial),
      noteId: noteId,
      payload: payload,
      timeout: 0, /* handled by Swap */
      cases: 1 /* initialize with 1 open case, WARNING: breaks when this is called multiple times */
    });

    /* initiate the swindle trial, swindle will call back once its over */
    bytes32 caseId = swindle.startTrial(provider, beneficiary, payload, commitmentHash, AbstractRules(trial));
    emit TrialStarted(commitmentHash, caseId, payload);
  }
}
