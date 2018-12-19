pragma solidity ^0.4.23;
import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "./SW3Utils.sol";

/// @title Swap Channel Contract
contract SimpleSwap is SW3Utils {
  using SafeMath for uint;

  event Deposit(address depositor, uint amount);
  event ChequeCashed(address indexed beneficiary, uint indexed serial, uint amount);
  event ChequeSubmitted(address indexed beneficiary, uint indexed serial, uint amount);
  event ChequeBounced(address indexed beneficiary, uint indexed serial);

  event HardDepositChanged(address indexed beneficiary, uint amount);
  event HardDepositDecreasePrepared(address indexed beneficiary, uint diff);

  /* magic timeout used throughout the code, cause of many security issues */
  uint constant timeout = 1 days;

  /* structure to keep track of the hard deposit for one beneficiary */
  struct HardDeposit {
    uint amount; /* current hard deposit */
    uint timeout; /* timeout of prepared HardDepositDecrease or 0 */
    uint diff; /* amount that will be removed on decrease */
  }

  /* structure to keep track of the lastest cheque for one beneficiary */
  struct ChequeInfo {
    uint serial; /* serial of the last submitted cheque */
    uint amount; /* cumulative amount of the last submitted cheque */
    uint paidOut; /* total amount paid out */
    uint timeout; /* timeout after which payout can happen */
  }

  /* associates every beneficiary with their ChequeInfo */
  mapping (address => ChequeInfo) public cheques;
  /* associates every beneficiary with their HardDeposit */
  mapping (address => HardDeposit) public hardDeposits;
  /* sum of all hard deposits */
  uint public totalDeposit;

  /* owner of the contract, set at construction */
  address public owner;

  /// @notice constructor, allows setting the owner (needed for "setup wallet as payment")
  constructor(address _owner) public {
    owner = _owner;
  }

  /// @return the part of the balance that is not covered by hard deposits
  function liquidBalance() public view returns(uint) {
    return address(this).balance.sub(totalDeposit);
  }

  /// @return the part of the balance usable for a specific beneficiary
  function liquidBalanceFor(address beneficiary) public view returns(uint) {
    return liquidBalance().add(hardDeposits[beneficiary].amount);
  }

  /// @dev helper function to process cheque after signatures have been checked
  /// @param beneficiary the beneficiary of the cheque
  /// @param serial the serial number of the cheque
  /// @param amount the (cumulative) amount of the cheque
  function _submitChequeInternal(address beneficiary, uint serial, uint amount) internal {
    /* ensure serial is increasing */
    ChequeInfo storage info = cheques[beneficiary];
    require(serial > info.serial);

    /* update the stored info */
    info.serial = serial;
    info.amount = amount;
    /* the check can be cashed timeout seconds in the future */
    info.timeout = now + timeout;

    /* the channel participants should watch to this event to find out if an older cheque is being submitted */
    emit ChequeSubmitted(beneficiary, serial, amount);
  }

  /* TODO: security implications of anyone being able to call this and the resulting timeout delay */
  /// @notice submit a cheque even if its lower
  /// @param beneficiary the beneficiary of the cheque
  /// @param serial the serial number of the cheque
  /// @param amount the (cumulative) amount of the cheque
  /// @param ownerSig signature of the owner
  /// @param beneficarySig signature of the beneficiary
  function submitChequeLower(address beneficiary, uint serial, uint amount, bytes ownerSig, bytes beneficarySig) public {
    /* verify signature of the owner */
    require(owner == recoverSignature(chequeHash(address(this), beneficiary, serial, amount), ownerSig));
    /* verify signature of the beneficiary */
    require(beneficiary == recoverSignature(chequeHash(address(this), beneficiary, serial, amount), beneficarySig));
    /* update the cheque data */
    _submitChequeInternal(beneficiary, serial, amount);
  }

  /// @dev helper function to payout value while respecting hard deposits
  /// @param beneficiary the address to send to
  /// @param value maximum amount to send
  /// @return payout amount that was actually paid out
  /// @return payout amount that bounced
  function _payout(address beneficiary, uint amount) internal returns (bool) {
    /* SWAP contract should be allowed to process the payout */
    if(amount >= liquidBalance()) {
      return false; // returning false signals a bounced cheque. TODO: consider reverting
    }

    /* Update internal accounting */
    hardDeposits[beneficiary].amount = hardDeposits[beneficiary].amount.sub(amount);
    totalDeposit = totalDeposit.sub(amount);

    /* transfer the payout */
    beneficiary.transfer(amount);
    return true;
  }

  /// @notice attempt to cash latest cheque
  /// @param beneficiary beneficiary for whose cheque should be paid out
  function cashCheque(address beneficiary, uint256 amount) public {
    ChequeInfo storage info = cheques[beneficiary];

    /* grace period must have ended */
    require(now >= info.timeout);

    /* ensure there is actually ether to be paid out */
    require(info.amount.sub(info.paidOut) >= amount);
    info.paidOut = info.paidOut.add(amount);

    /* Attempt the actual payout and emit events based on the result */
    if(!_payout(beneficiary, amount)) {
        emit ChequeBounced(beneficiary, info.serial);
    } else {
        emit ChequeCashed(beneficiary, info.serial, amount);
    }
  }

  /// @notice prepare to decrease the hard deposit
  /// @param beneficiary beneficiary whose hard deposit should be decreased
  /// @param diff amount that the deposit is supposed to be decreased by
  function prepareDecreaseHardDeposit(address beneficiary, uint diff) public {
    require(msg.sender == owner);
    HardDeposit storage deposit = hardDeposits[beneficiary];
    /* cannot decrease it by more than the deposit */
    require(diff <= deposit.amount);

    /* timeout is twice the normal timeout to ensure users can submit and cash in time */
    deposit.timeout = now + timeout * 2;
    deposit.diff = diff;
    emit HardDepositDecreasePrepared(beneficiary, diff);
  }

  /* TODO: necessary to make sure no funds can be permanently locked but this also breaks security of off-chain Swear, so this needs to change */
  /// @notice actually decrease the hard deposit
  /// @param beneficiary beneficiary whose hard deposit should be decreased
  function decreaseHardDeposit(address beneficiary) public {
    HardDeposit storage deposit = hardDeposits[beneficiary];

    /* check that there was a timeout and that it has passed */
    require(deposit.timeout != 0);
    require(now >= deposit.timeout);

    /* decrease the amount */
    /* diff can never be more than amount (require statement in prepareDecreaseHardDeposit) */
    deposit.amount = deposit.amount - deposit.diff;
    /* reset the timeout to avoid a double decrease */
    deposit.timeout = 0;
    /* keep totalDeposit in sync */
    totalDeposit = totalDeposit.sub(deposit.diff);

    emit HardDepositChanged(beneficiary, deposit.amount);
  }

  /// @notice increase the hard deposit
  /// @param beneficiary beneficiary whose hard deposit should be decreased
  /// @param amount the new hard deposit
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

  /// @notice withdraw ether
  /// @param amount amount to withdraw
  function withdraw(uint amount) public {
    /* only owner can do this */
    require(msg.sender == owner);
    /* ensure we don't take anything from the hard deposit */
    require(amount <= liquidBalance());
    owner.transfer(amount);
  }

  /// @notice deposit ether
  function() payable public {
    emit Deposit(msg.sender, msg.value);
  }
}
