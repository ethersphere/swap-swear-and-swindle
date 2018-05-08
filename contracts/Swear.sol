pragma solidity ^0.4.19;
import "./abstracts/AbstractRules.sol";
import "./Swindle.sol";
import "zeppelin/math/SafeMath.sol";

contract Swear {
  using SafeMath for uint;

  event CommitmentAdded(bytes32 commitmentHash, address indexed provider, address rules);
  event TrialStarted(bytes32 commitmentHash, bytes32 caseId);

  Swindle public swindle;

  function Swear(address _swindle) public {
    swindle = Swindle(_swindle);
  }

  struct Commitment {
    bool valid;
    address provider;

    uint deposit;
    uint timeout;
    AbstractRules rules;
    bytes32 noteId;

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
      cases: 0
    });

    CommitmentAdded(commitmentHash, msg.sender, rules);
  }

  function compensate(bytes32 commitmentHash, address beneficiary, uint reward) public {
    require(msg.sender == address(swindle));
    Commitment storage commitment = commitments[commitmentHash];
    commitment.deposit = commitment.deposit.sub(reward);
    beneficiary.transfer(reward);

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

    TrialStarted(commitmentHash, caseId);
  }

  function notifyTrialEnd(bytes32 commitmentHash) public {
    require(msg.sender == address(swindle));
    Commitment storage commitment = commitments[commitmentHash];

    commitment.cases--;
  }
}
