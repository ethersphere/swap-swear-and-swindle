pragma solidity ^0.5.10;
import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "openzeppelin-solidity/contracts/math/Math.sol";
import "openzeppelin-solidity/contracts/cryptography/ECDSA.sol";

/// @title Swap Channel Contract
contract SimpleSwap {
  using SafeMath for uint;

  event Deposit(address depositor, uint amount);
  event ChequeCashed(address indexed beneficiary, uint indexed serial, uint payout, uint requestPayout);
  event ChequeSubmitted(address indexed beneficiary, uint indexed serial, uint amount, uint cashTimeout);
  event ChequeBounced();
  event HardDepositAmountChanged(address indexed beneficiary, uint amount);
  event HardDepositDecreasePrepared(address indexed beneficiary, uint decreaseAmount);
  event HardDepositDecreaseTimeoutChanged(address indexed beneficiary, uint decreaseTimeout);

  uint DEFAULT_HARDDEPPOSIT_DECREASE_TIMEOUT;
  /* structure to keep track of the hard deposits (on-chain guarantee of solvency) per beneficiary*/
  struct HardDeposit {
    uint amount; /* hard deposit amount allocated */
    uint decreaseAmount; /* decreaseAmount substranced from amount when decrease is requested */
    uint decreaseTimeout; /* owner has to wait decreaseTimeout seconds after decrease is requested to decrease hardDeposit */
    uint canBeDecreasedAt; /* point in time after which harddeposit can be decreased*/
  }
  /* structure to keep track of the latest cheque for one beneficiary */
  struct ChequeInfo {
    uint serial; /* serial of the last submitted cheque */
    uint amount; /* cumulative amount of the last submitted cheque */
    uint paidOut; /* total amount paid out */
    uint cashTimeout; /* timeout after which payout can happen */
  }
  /* associates every beneficiary with their ChequeInfo */
  mapping (address => ChequeInfo) public cheques;
  /* associates every beneficiary with their HardDeposit */
  mapping (address => HardDeposit) public hardDeposits;
  /* sum of all hard deposits */
  uint public totalHardDeposit;

  /* owner of the contract, set at construction */
  address payable public owner;

  /// @notice constructor, allows setting the owner (needed for "setup wallet as payment")
  constructor(address payable _owner, uint defaultHardDepositTimeoutDuration) public {
    // DEFAULT_HARDDEPOSIT_TIMOUTE_DURATION will be one day or a whatever non-zero argument given as an argument to the constructor
    DEFAULT_HARDDEPPOSIT_DECREASE_TIMEOUT = defaultHardDepositTimeoutDuration == 0 ? 1 days : defaultHardDepositTimeoutDuration;
    owner = _owner;
  }

  /// @return the part of the balance that is not covered by hard deposits
  function liquidBalance() public view returns(uint) {
    return address(this).balance.sub(totalHardDeposit);
  }

  /// @return the part of the balance usable for a specific beneficiary
  function liquidBalanceFor(address beneficiary) public view returns(uint) {
    return liquidBalance().add(hardDeposits[beneficiary].amount);
  }

  /// @dev helper function to process cheque after signatures have been checked
  /// @param beneficiary the beneficiary of the cheque
  /// @param serial the serial number of the cheque
  /// @param amount the (cumulative) amount of the cheque
  /// @param cashTimeout the check can be cashed cashTimeout seconds in the future
  function _submitChequeInternal(address beneficiary, uint serial, uint amount, uint cashTimeout) internal {
    ChequeInfo storage cheque = cheques[beneficiary];
    /* ensure serial is increasing */
    require(serial > cheque.serial, "SimpleSwap: invalid serial");
    /* update the stored info */
    cheque.serial = serial;
    cheque.amount = amount;
    cheque.cashTimeout = now + cashTimeout;
    /* the channel participants should watch to this event to find out if an older cheque is being submitted */
    emit ChequeSubmitted(beneficiary, serial, amount, cashTimeout);
  }

  /// @notice submit a cheque by the owner
  /// @param beneficiary the beneficiary of the cheque
  /// @param serial the serial number of the cheque
  /// @param amount the (cumulative) amount of the cheque
  /// @param cashTimeout the check can be cashed cashTimeout seconds in the future
  /// @param beneficiarySig signature of the owner
  function submitChequeOwner(address beneficiary, uint serial, uint amount, uint cashTimeout, bytes memory beneficiarySig) public {
    require(msg.sender == owner, "SimpleSwap: not owner");
    /* verify signature of the beneficiary */
    require(beneficiary == recover(chequeHash(address(this), beneficiary, serial, amount, cashTimeout), beneficiarySig),
     "SimpleSwap: invalid beneficiarySig");
    /* update the cheque data */
    _submitChequeInternal(beneficiary, serial, amount, cashTimeout);
  }

  /// @notice submit a cheque by the beneficiary
  /// @param serial the serial number of the cheque
  /// @param amount the (cumulative) amount of the cheque
  /// @param cashTimeout the check can be cashed cashTimeout seconds in the future
  /// @param ownerSig signature of the owner
  function submitChequeBeneficiary(uint serial, uint amount, uint cashTimeout, bytes memory ownerSig) public {
    /* verify signature of the owner */
    //emit LogAddress(recover(chequeHash(address(this), msg.sender, serial, amount, cashTimeout), ownerSig));
    require(owner == recover(chequeHash(address(this), msg.sender, serial, amount, cashTimeout), ownerSig),
     "SimpleSwap: invalid ownerSig");
    /* update the cheque data */
    _submitChequeInternal(msg.sender, serial, amount, cashTimeout);
  }

  /// @notice submit a cheque by any party
  /// @param beneficiary the beneficiary of the cheque
  /// @param serial the serial number of the cheque
  /// @param amount the (cumulative) amount of the cheque
  /// @param cashTimeout the check can be cashed cashTimeout seconds in the future
  /// @param ownerSig signature of the owner
  /// @param beneficarySig signature of the beneficiary
  function submitCheque(address beneficiary, uint serial, uint amount, uint cashTimeout, bytes memory ownerSig, bytes memory beneficarySig) public {
    /* verify signature of the owner */
    require(owner == recover(chequeHash(address(this), beneficiary, serial, amount, cashTimeout), ownerSig),
    "SimpleSwap: invalid ownerSig");
    /* verify signature of the beneficiary */
    require(beneficiary == recover(chequeHash(address(this), beneficiary, serial, amount, cashTimeout), beneficarySig),
    "SimpleSwap: invalid beneficiarySig");
    /* update the cheque data */
    _submitChequeInternal(beneficiary, serial, amount, cashTimeout);
  }

  /// @notice attempt to cash latest cheque
  /// @param beneficiary beneficiary for whose cheque should be paid out
  /// @param requestPayout amount requested to pay out
  function cashCheque(address payable beneficiary, uint requestPayout) public {
    ChequeInfo storage cheque = cheques[beneficiary];
    /* grace period must have ended */
    require(now >= cheque.cashTimeout,  "SimpleSwap: cheque not yet timed out");
    require(requestPayout <= cheque.amount.sub(cheque.paidOut), "SimpleSwap: not enough balance owed");
    /* ensure there is a balance to claim */
     /* calculates hard-deposit usage */
    uint hardDepositUsage = Math.min(requestPayout, hardDeposits[beneficiary].amount);
    /* calculates acutal payout */
    uint payout = Math.min(requestPayout, liquidBalance() + hardDepositUsage);
      /* if there some of the hard deposit is used update the structure */
    if(hardDepositUsage != 0) {
      hardDeposits[beneficiary].amount = hardDeposits[beneficiary].amount.sub(hardDepositUsage);
      totalHardDeposit = totalHardDeposit.sub(hardDepositUsage);
    }
    /* increase the stored paidOut amount to avoid double payout */
    cheque.paidOut = cheque.paidOut.add(payout);
    /* do the actual payment */
    beneficiary.transfer(payout);
    emit ChequeCashed(beneficiary, cheque.serial, payout, requestPayout);
    if(requestPayout != payout) {
      emit ChequeBounced();
    }
  }

  /// @notice prepare to decrease the hard deposit
  /// @param beneficiary beneficiary whose hard deposit should be decreased
  /// @param decreaseAmount amount that the deposit is supposed to be decreased by
  function prepareDecreaseHardDeposit(address beneficiary, uint decreaseAmount) public {
    require(msg.sender == owner, "SimpleSwap: not owner");
    HardDeposit storage hardDeposit = hardDeposits[beneficiary];
    /* cannot decrease it by more than the deposit */
    require(decreaseAmount <= hardDeposit.amount, "SimpleSwap: hard deposit not sufficient");
    // if hardDeposit.decreaseTimeout was never set, we DEFAULT_HARDDEPPOSIT_DECREASE_TIMEOUT. Otherwise we use the one which was set.
    uint decreaseTimeout = hardDeposit.decreaseTimeout == 0 ? DEFAULT_HARDDEPPOSIT_DECREASE_TIMEOUT : hardDeposit.decreaseTimeout;
    hardDeposit.decreaseTimeout = now + decreaseTimeout;
    hardDeposit.decreaseAmount = decreaseAmount;
    emit HardDepositDecreasePrepared(beneficiary, decreaseAmount);
  }

  /// @notice actually decrease the hard deposit
  /// @param beneficiary beneficiary whose hard deposit should be decreased
  function decreaseHardDeposit(address beneficiary) public {
    HardDeposit storage hardDeposit = hardDeposits[beneficiary];

    require(now >= hardDeposit.decreaseTimeout && hardDeposit.decreaseTimeout != 0, "SimpleSwap: deposit not yet timed out");

    /* decrease the amount */
    /* this throws if decreaseAmount > amount */
    hardDeposit.amount = hardDeposit.amount.sub(hardDeposit.decreaseAmount);
    /* reset the decreaseTimeout to avoid a double decrease */
    hardDeposit.decreaseTimeout = 0;
    /* keep totalDeposit in sync */
    totalHardDeposit = totalHardDeposit.sub(hardDeposit.decreaseAmount);

    emit HardDepositAmountChanged(beneficiary, hardDeposit.amount);
  }

  /// @notice increase the hard deposit
  /// @param beneficiary beneficiary whose hard deposit should be decreased
  /// @param amount the new hard deposit
  function increaseHardDeposit(address beneficiary, uint amount) public {
    require(msg.sender == owner, "SimpleSwap: not owner");
    /* ensure hard deposits don't exceed the global balance */
    require(totalHardDeposit.add(amount) <= address(this).balance, "SimpleSwap: hard deposit cannot be more than balance ");

    HardDeposit storage hardDeposit = hardDeposits[beneficiary];
    hardDeposit.amount = hardDeposit.amount.add(amount);
    // we don't explicitely set timeoutDuration, as zero means using the DEFAULT_HARDDEPOSIT_TIMEOUT_DURATION
    totalHardDeposit = totalHardDeposit.add(amount);
    /* disable any pending decrease */
    hardDeposit.decreaseTimeout = 0;
    emit HardDepositAmountChanged(beneficiary, hardDeposit.amount);
  }


  function setCustomHardDepositDecreaseTimeout(
    address beneficiary,
    uint decreaseTimeout,
    bytes memory ownerSig,
    bytes memory beneficiarySig
  ) public {
    require(owner == recover(keccak256(abi.encodePacked(address(this), beneficiary, decreaseTimeout)), ownerSig));
    require(beneficiary == recover(keccak256(abi.encodePacked(address(this), beneficiary, decreaseTimeout)), beneficiarySig));
    hardDeposits[beneficiary].decreaseTimeout = decreaseTimeout;
    emit HardDepositDecreaseTimeoutChanged(beneficiary, hardDeposits[beneficiary].decreaseTimeout);
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

  /// @notice deposit ether
  function() payable external {
    emit Deposit(msg.sender, msg.value);
  }

  function recover(bytes32 hash, bytes memory sig) internal pure returns (address) {
    return ECDSA.recover(ECDSA.toEthSignedMessageHash(hash), sig);
  }

  function chequeHash(address swap, address beneficiary, uint serial, uint amount, uint cashTimeout)
  public pure returns (bytes32) {
    return keccak256(abi.encodePacked(swap, serial, beneficiary, amount, cashTimeout));
  }
}
