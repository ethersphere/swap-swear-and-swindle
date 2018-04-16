pragma solidity ^0.4.19;
import "zeppelin/math/SafeMath.sol";
import "zeppelin/math/Math.sol";

contract Swap {
  using SafeMath for uint;

  event Deposit(address depositor, uint amount);
  event ChequeCashed(address indexed beneficiary, uint indexed serial, uint amount);
  event ChequeSubmitted(address indexed beneficiary, uint indexed serial, uint amount);
  event ChequeBounced(address indexed beneficiary, uint indexed serial, uint paid, uint bounced);

  event HardDepositChanged(address indexed beneficiary, uint amount);
  event HardDepositDecreasePrepared(address indexed beneficiary, uint amount);

  uint constant timeout = 1 days;

  struct HardDeposit {
    uint amount;
    uint timeout;
    uint next;
  }

  struct ChequeInfo {
    uint serial;
    uint amount;
    uint paidOut;
    uint timeout;
  }

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

  function chequeHash(address beneficiary, uint serial, uint amount) public pure returns (bytes32) {
    return keccak256(serial, beneficiary, amount);
  }

  function recoverSignature(bytes32 hash, bytes32 r, bytes32 s, uint8 v) public pure returns (address) {
    return ecrecover(keccak256("\x19Ethereum Signed Message:\n32", hash), v, r, s);
  }

  function submitCheque(address beneficiary, uint serial, uint amount, bytes32 r, bytes32 s, uint8 v) public {
    /* ensure serial is increasing */
    ChequeInfo storage info = infos[beneficiary];
    require(serial > info.serial);

    /* verify signature */
    address signer = recoverSignature(chequeHash(beneficiary, serial, amount), r, s, v);

    require(
      (signer == owner && amount > info.amount) || /* allow owner to increase the cheque value */
      (signer == beneficiary && amount < info.amount) /* allow beneficiary to decrease the cheque value */
    );

    /* update the stored info */
    info.serial = serial;
    info.amount = amount;
    /* the grace period ends timeout seconds in the future */
    info.timeout = now + timeout;

    emit ChequeSubmitted(beneficiary, serial, amount);
  }

  function cashCheque(address beneficiary) public {
    ChequeInfo storage info = infos[beneficiary];

    /* grace period must have ended */
    require(now >= info.timeout);

    /* ensure there is actually ether to be paid out */
    uint value = info.amount.sub(info.paidOut); /* throws if paidOut > amount */
    require(value > 0);

    /* part of hard deposit used */
    uint payout = Math.min256(value, hardDeposits[beneficiary].amount);
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
      emit ChequeCashed(beneficiary, info.serial, payout);
    } else {
      /* part of the cheque bounces */
      payout += liquid;
      emit ChequeBounced(beneficiary, info.serial, payout, rest - liquid);
    }

    info.paidOut = info.paidOut.add(payout);
    beneficiary.transfer(payout);
  }

  function prepareDecreaseHardDeposit(address beneficiary, uint amount) public {
    require(msg.sender == owner);
    HardDeposit storage deposit = hardDeposits[beneficiary];
    require(amount < deposit.amount);

    /* timeout is twice the normal timeout to ensure users can submit and cash in time */
    deposit.timeout = now + timeout * 2;
    deposit.next = amount;
    emit HardDepositDecreasePrepared(beneficiary, amount);
  }

  function decreaseHardDeposit(address beneficiary) {
    HardDeposit storage deposit = hardDeposits[beneficiary];

    require(deposit.timeout != 0);
    require(now >= deposit.timeout);

    uint diff = deposit.amount - deposit.next;

    deposit.amount = deposit.next;

    deposit.timeout = 0;

    totalDeposit = totalDeposit.sub(diff);

    emit HardDepositChanged(beneficiary, deposit.amount);
  }

  function increaseHardDeposit(address beneficiary, uint amount) public {
    require(msg.sender == owner);
    require(totalDeposit.add(amount) <= address(this).balance);
    hardDeposits[beneficiary].amount = hardDeposits[beneficiary].amount.add(amount);
    totalDeposit = totalDeposit.add(amount);
    emit HardDepositChanged(beneficiary, hardDeposits[beneficiary].amount);
  }

  function withdraw(uint amount) public {
    require(msg.sender == owner);
    require(amount <= liquidBalance());
    owner.transfer(amount);
  }

  function() payable public {
    emit Deposit(msg.sender, msg.value);
  }

}
