pragma solidity ^0.4.19;
import "zeppelin/math/SafeMath.sol";
import "zeppelin/math/Math.sol";
import "./abstracts/AbstractWitness.sol";

contract Swap {
  using SafeMath for uint;

  event Deposit(address depositor, uint amount);
  event ChequeCashed(address indexed beneficiary, uint indexed serial, uint amount);
  event ChequeSubmitted(address indexed beneficiary, uint indexed serial, uint amount);
  event ChequeBounced(address indexed beneficiary, uint indexed serial, uint paid, uint bounced);

  event HardDepositChanged(address indexed beneficiary, uint amount);
  event HardDepositDecreasePrepared(address indexed beneficiary, uint diff);

  uint constant timeout = 1 days;

  struct HardDeposit {
    uint amount;
    uint timeout;
    uint diff;
  }

  struct ChequeInfo {
    uint serial;
    uint amount;
    uint paidOut;
    uint timeout;
  }

  struct NoteInfo {
    uint index;
    uint amount;
    uint paidOut;
    uint timeout;
    address beneficiary;
    address witness;
    uint validFrom;
    uint validUntil;
    bytes32 remark;
  }

  mapping (bytes32 => NoteInfo) public notes;
  mapping (address => ChequeInfo) public infos;
  mapping (address => HardDeposit) public hardDeposits;
  uint public totalDeposit;

  address public owner;

  /* constructor, allows setting the owner (needed for "setup wallet as payment") */
  function Swap(address _owner) public {
    owner = _owner;
  }

  function liquidBalance() public view returns(uint) {
    return address(this).balance.sub(totalDeposit);
  }

  function liquidBalanceFor(address beneficiary) public view returns(uint) {
    return liquidBalance().add(hardDeposits[beneficiary].amount);
  }

  function chequeHash(address beneficiary, uint serial, uint amount) public view returns (bytes32) {
    return keccak256(address(this), serial, beneficiary, amount);
  }

  function recoverSignature(bytes32 hash, bytes32 r, bytes32 s, uint8 v) public pure returns (address) {
    return ecrecover(keccak256("\x19Ethereum Signed Message:\n32", hash), v, r, s);
  }

  function _submitChequeInternal(address beneficiary, uint serial, uint amount) internal {
    /* ensure serial is increasing */
    ChequeInfo storage info = infos[beneficiary];
    require(serial > info.serial);

    /* update the stored info */
    info.serial = serial;
    info.amount = amount;
    /* the grace period ends timeout seconds in the future */
    info.timeout = now + timeout;

    emit ChequeSubmitted(beneficiary, serial, amount);
  }

  function submitCheque(address beneficiary, uint serial, uint amount, bytes32 r, bytes32 s, uint8 v) public {
    require(msg.sender == beneficiary);
    /* verify signature */
    require(owner ==  recoverSignature(chequeHash(beneficiary, serial, amount), r, s, v));
    require(amount > infos[beneficiary].amount);
    _submitChequeInternal(beneficiary, serial, amount);
  }

  /* TODO: security implications of anyone being able to call this and the resulting timeout delay */
  function submitChequeLower(address beneficiary, uint serial, uint amount, bytes32 r, bytes32 s, uint8 v, bytes32 r2, bytes32 s2, uint8 v2) public {
    /* verify signature */
    require(owner == recoverSignature(chequeHash(beneficiary, serial, amount), r, s, v));
    require(beneficiary == recoverSignature(chequeHash(beneficiary, serial, amount), r2, s2, v2));

    _submitChequeInternal(beneficiary, serial, amount);
  }

  function _payout(address beneficiary, uint value) internal returns (uint payout, uint bounced) {
    /* part of hard deposit used */
    payout = Math.min256(value, hardDeposits[beneficiary].amount);
    if(payout != 0) {
      hardDeposits[beneficiary].amount -= payout;
      totalDeposit -= payout;
    }

    /* amount of the cash not backed by a hard deposit */
    uint rest = value - payout;
    uint liquid = liquidBalance();

    if(liquid >= rest) {
      /* swap channel is solvent */
      payout = value;
    } else {
      /* part of the cheque bounces */
      payout += liquid;
      bounced = rest - liquid;
    }

    beneficiary.transfer(payout);
  }

  function cashCheque(address beneficiary) public {
    ChequeInfo storage info = infos[beneficiary];

    /* grace period must have ended */
    require(now >= info.timeout);

    /* ensure there is actually ether to be paid out */
    uint value = info.amount.sub(info.paidOut); /* throws if paidOut > amount */
    require(value > 0);

    uint payout;
    uint bounced;

    (payout, bounced) = _payout(beneficiary, value);

    if(bounced != 0) emit ChequeBounced(beneficiary, info.serial, payout, bounced);
    else emit ChequeCashed(beneficiary, info.serial, payout);

    info.paidOut = info.paidOut.add(payout);
  }

  function prepareDecreaseHardDeposit(address beneficiary, uint diff) public {
    require(msg.sender == owner);
    HardDeposit storage deposit = hardDeposits[beneficiary];
    require(diff < deposit.amount);

    /* timeout is twice the normal timeout to ensure users can submit and cash in time */
    deposit.timeout = now + timeout * 2;
    deposit.diff = diff;
    emit HardDepositDecreasePrepared(beneficiary, diff);
  }

  function decreaseHardDeposit(address beneficiary) {
    HardDeposit storage deposit = hardDeposits[beneficiary];

    require(deposit.timeout != 0);
    require(now >= deposit.timeout);

    deposit.amount = deposit.amount.sub(deposit.diff);
    deposit.timeout = 0;

    totalDeposit = totalDeposit.sub(deposit.diff);

    emit HardDepositChanged(beneficiary, deposit.amount);
  }

  function increaseHardDeposit(address beneficiary, uint amount) public {
    require(msg.sender == owner);
    /* ensure hard deposits don't exceed the global balance */
    require(totalDeposit.add(amount) <= address(this).balance);

    HardDeposit storage deposit = hardDeposits[beneficiary];
    deposit.amount = deposit.amount.add(amount);
    totalDeposit = totalDeposit.add(amount);
    /* disable any pending decrease */
    deposit.timeout = 0;
    emit HardDepositChanged(beneficiary, deposit.amount);
  }

  function withdraw(uint amount) public {
    require(msg.sender == owner);
    require(amount <= liquidBalance());
    owner.transfer(amount);
  }

  function() payable public {
    emit Deposit(msg.sender, msg.value);
  }

  function submitNote(uint index, uint amount, address beneficiary, address witness, uint validFrom, uint validUntil, bytes32 remark, bytes32 r, bytes32 s, uint8 v) public {
    bytes32 noteId = keccak256(address(this), index, amount, beneficiary, witness, validFrom, validUntil, remark);

    require(owner == recoverSignature(noteId, r, s, v));
    require(notes[noteId].index == 0);

    notes[noteId] = NoteInfo({
      index: index,
      amount: amount,
      beneficiary: beneficiary,
      paidOut: 0,
      witness: witness,
      validFrom: validFrom,
      validUntil: validUntil,
      remark: remark,
      timeout: now + timeout
    });
  }

  function cashNote(bytes32 noteId, uint amount) public {
    NoteInfo storage note = notes[noteId];

    require(now >= note.timeout);
    if(note.validFrom != 0) require(now >= note.validFrom);
    if(note.validUntil != 0) require(now <= note.validUntil);

    address beneficiary = note.beneficiary;

    if(beneficiary != address(0x0)) {
      require(msg.sender == beneficiary);
    } else beneficiary = msg.sender;

    if(note.witness != address(0x0)) {
      /* TODO: re-entrance considerations, should be called STATIC? */
      require(AbstractWitness(note.witness).testimonyFor(owner, beneficiary, noteId) == AbstractWitness.TestimonyStatus.VALID);
    }

    if(note.amount != 0) {
      require(note.paidOut.add(amount) <= note.amount);
    }

    uint payout;
    uint bounced;

    (payout, bounced) = _payout(beneficiary, amount);

    note.paidOut += payout;
  }

  function submitPaidInvoice(bytes32 noteId, uint swapBalance, uint serial, bytes32 r, bytes32 s, uint8 v, uint amount, bytes32 r2, bytes32 s2, uint8 v2) public {
    require(msg.sender == owner);
    bytes32 invoiceId = keccak256(noteId, swapBalance, serial);

    NoteInfo storage note = notes[noteId];
    require(note.index != 0);
    require(note.timeout != 0 && note.timeout < now);

    uint cumulativeTotal = swapBalance.add(amount);

    /* TODO: this breaks with note.beneficiary = 0 */
    require(note.beneficiary == recoverSignature(invoiceId, r, s, v));

    require(owner == recoverSignature(chequeHash(note.beneficiary, serial + 1, cumulativeTotal), r2, s2, v2));

    /* TODO: this breaks with note.amount = 0 */
    require(note.amount == amount);
    note.paidOut = amount;

    if(serial + 1 > infos[note.beneficiary].serial)
      _submitChequeInternal(note.beneficiary, serial + 1, cumulativeTotal);
  }

}
