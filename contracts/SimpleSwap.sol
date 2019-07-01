pragma solidity ^0.5.10;
import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "openzeppelin-solidity/contracts/math/Math.sol";
import "openzeppelin-solidity/contracts/cryptography/ECDSA.sol";

/// @title Swap Channel Contract
contract SimpleSwap {
  using SafeMath for uint;

  event Deposit(address depositor, uint amount);
  event ChequeCashed(address indexed beneficiary, uint indexed serial, uint payout, uint requestPayout);
  event ChequeSubmitted(address indexed beneficiary, uint indexed serial, uint amount, uint timeout);
  event ChequeBounced(address indexed beneficiary, uint indexed serial, uint paid, uint bounced);
  event HardDepositChanged(address indexed beneficiary, uint amount);
  event HardDepositDecreasePrepared(address indexed beneficiary, uint diff);

  /* magic timeout used throughout the code, cause of many security issues */
  uint constant hardDepositTimeout = 1 days;

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
  address payable public owner;

  /// @notice constructor, allows setting the owner (needed for "setup wallet as payment")
  constructor(address payable _owner) public {
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
  /// @param timeout the check can be cashed timeout seconds in the future
  function _submitChequeInternal(address beneficiary, uint serial, uint amount, uint timeout) internal {
    ChequeInfo storage cheque = cheques[beneficiary];
    /* ensure serial is increasing */
    require(serial > cheque.serial, "SimpleSwap: invalid serial");
    /* update the stored info */
    cheque.serial = serial;
    cheque.amount = amount;
    cheque.timeout = now + timeout;
    /* the channel participants should watch to this event to find out if an older cheque is being submitted */
    emit ChequeSubmitted(beneficiary, serial, amount, timeout);
  }

  /// @notice submit a cheque by the owner
  /// @param beneficiary the beneficiary of the cheque
  /// @param serial the serial number of the cheque
  /// @param amount the (cumulative) amount of the cheque
  /// @param timeout the check can be cashed timeout seconds in the future
  /// @param beneficiarySig signature of the owner
  function submitChequeOwner(address beneficiary, uint serial, uint amount, uint timeout, bytes memory beneficiarySig) public {
    require(msg.sender == owner, "SimpleSwap: not owner");
    /* verify signature of the beneficiary */
    require(beneficiary == recover(chequeHash(address(this), beneficiary, serial, amount, timeout), beneficiarySig),
     "SimpleSwap: invalid beneficiarySig");
    /* update the cheque data */
    _submitChequeInternal(beneficiary, serial, amount, timeout);
  }

  /// @notice submit a cheque by the beneficiary
  /// @param serial the serial number of the cheque
  /// @param amount the (cumulative) amount of the cheque
  /// @param timeout the check can be cashed timeout seconds in the future
  /// @param ownerSig signature of the owner
  function submitChequeBeneficiary(uint serial, uint amount, uint timeout, bytes memory ownerSig) public {
    /* verify signature of the owner */
    //emit LogAddress(recover(chequeHash(address(this), msg.sender, serial, amount, timeout), ownerSig));
    require(owner == recover(chequeHash(address(this), msg.sender, serial, amount, timeout), ownerSig),
     "SimpleSwap: invalid ownerSig");
    /* update the cheque data */
    _submitChequeInternal(msg.sender, serial, amount, timeout);
  }

  /// @notice submit a cheque by any party
  /// @param beneficiary the beneficiary of the cheque
  /// @param serial the serial number of the cheque
  /// @param amount the (cumulative) amount of the cheque
  /// @param timeout the check can be cashed timeout seconds in the future
  /// @param ownerSig signature of the owner
  /// @param beneficarySig signature of the beneficiary
  function submitCheque(address beneficiary, uint serial, uint amount, uint timeout, bytes memory ownerSig, bytes memory beneficarySig) public {
    /* verify signature of the owner */
    require(owner == recover(chequeHash(address(this), beneficiary, serial, amount, timeout), ownerSig),
    "SimpleSwap: invalid ownerSig");
    /* verify signature of the beneficiary */
    require(beneficiary == recover(chequeHash(address(this), beneficiary, serial, amount, timeout), beneficarySig),
    "SimpleSwap: invalid beneficiarySig");
    /* update the cheque data */
    _submitChequeInternal(beneficiary, serial, amount, timeout);
  }

  /// @notice attempt to cash latest cheque
  /// @param beneficiary beneficiary for whose cheque should be paid out
  /// @param requestPayout amount requested to pay out
  function cashCheque(address payable beneficiary, uint requestPayout) public returns (uint) {
    ChequeInfo storage cheque = cheques[beneficiary];
    /* grace period must have ended */
    require(now >= cheque.timeout,  "SimpleSwap: cheque not yet timed out");
    /* ensure there is a balance to claim */
    require(requestPayout < cheque.amount.sub(cheque.paidOut), "SimpleSwap: requestPayout more than owed");
     /* calculates hard-deposit usage */
    uint hardDepositUsage = Math.min(requestPayout, hardDeposits[beneficiary].amount);
    /* calculates acutal payout */
    uint payout = Math.min(requestPayout, liquidBalance() + hardDepositUsage);
    /* will throw if chequebook is not solvent */
    require(payout != 0, "SimpleSwap: contract not solvent");
      /* if there some of the hard deposit is used update the structure */
    if(hardDepositUsage != 0) {
      hardDeposits[beneficiary].amount = hardDepositUsage.sub(hardDepositUsage);
      totalDeposit = totalDeposit.sub(hardDepositUsage);
    }
    /* increase the stored paidOut amount to avoid double payout */
    cheque.paidOut = cheque.paidOut.add(payout);
    /* do the actual payment */
    beneficiary.transfer(payout);
    emit ChequeCashed(beneficiary, cheque.serial, payout, requestPayout);
    return payout;
  }

  /// @notice prepare to decrease the hard deposit
  /// @param beneficiary beneficiary whose hard deposit should be decreased
  /// @param diff amount that the deposit is supposed to be decreased by
  function prepareDecreaseHardDeposit(address beneficiary, uint diff) public {
    require(msg.sender == owner, "SimpleSwap: not owner");
    HardDeposit storage deposit = hardDeposits[beneficiary];
    /* cannot decrease it by more than the deposit */
    require(diff < deposit.amount, "SimpleSwap: balance insufficient");

    /* timeout is twice the normal timeout to ensure users can submit and cash in time */
    deposit.timeout = now + hardDepositTimeout;
    deposit.diff = diff;
    emit HardDepositDecreasePrepared(beneficiary, diff);
  }

  /* TODO: necessary to make sure no funds can be permanently locked but this also breaks security of off-chain Swear, so this needs to change */
  /// @notice actually decrease the hard deposit
  /// @param beneficiary beneficiary whose hard deposit should be decreased
  function decreaseHardDeposit(address beneficiary) public {
    HardDeposit storage deposit = hardDeposits[beneficiary];

    /* check that there was a timeout and that it has passed */
    require(deposit.timeout != 0, "SimpleSwap: no timeout set");
    require(now >= deposit.timeout, "SimpleSwap: deposit not yet timed out");

    /* decrease the amount */
    /* this throws if diff > amount */
    deposit.amount = deposit.amount.sub(deposit.diff);
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
    require(msg.sender == owner, "SimpleSwap: not owner");
    /* ensure hard deposits don't exceed the global balance */
    require(totalDeposit.add(amount) <= address(this).balance, "SimpleSwap: hard deposit cannot be more than balance ");

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
    require(msg.sender == owner, "SimpleSwap: not owner");
    /* ensure we don't take anything from the hard deposit */
    require(amount <= liquidBalance(), "SimpleSwap: liquidBalance not sufficient");
    owner.transfer(amount);
  }

   /// @dev helper function to calculate payout value while respecting hard deposits
  /// @param beneficiary the address to send to
  /// @param value maximum amount to send
  /// @return payout amount that was actually paid out
  /// @return payout amount that bounced
  function _computePayout(address payable beneficiary, uint value) internal returns (uint payout, uint bounced) {
    /* part of hard deposit used */
    payout = Math.min(value, hardDeposits[beneficiary].amount);
    /* if there some of the hard deposit is used update the structure */
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
  }

  /// @notice deposit ether
  function() payable external {
    emit Deposit(msg.sender, msg.value);
  }

  function recover(bytes32 hash, bytes memory sig) internal pure returns (address) {
    return ECDSA.recover(ECDSA.toEthSignedMessageHash(hash), sig);
  }

  function chequeHash(address swap, address beneficiary, uint serial, uint amount, uint timeout)
  public pure returns (bytes32) {
    return keccak256(abi.encodePacked(swap, serial, beneficiary, amount, timeout));
  }
}
