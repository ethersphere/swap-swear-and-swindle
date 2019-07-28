pragma solidity ^0.5.10;
import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "openzeppelin-solidity/contracts/math/Math.sol";
import "openzeppelin-solidity/contracts/cryptography/ECDSA.sol";

/// @title Swap Channel Contract
contract SimpleSwap {
  using SafeMath for uint;

  event Deposit(address depositor, uint amount);
  event ChequeCashed(
    address indexed beneficiaryPrincipal,
    address indexed beneficiaryAgent,
    address indexed callee,
    uint serial,
    uint totalPayout,
    uint requestPayout,
    uint calleePayout
  );
  event ChequeSubmitted(address indexed beneficiary, uint indexed serial, uint amount, uint cashTimeout);
  event ChequeBounced();
  event HardDepositAmountChanged(address indexed beneficiary, uint amount);
  event HardDepositDecreasePrepared(address indexed beneficiary, uint decreaseAmount);
  event HardDepositDecreaseTimeoutChanged(address indexed beneficiary, uint decreaseTimeout);

  uint public DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT;
  /* structure to keep track of the hard deposits (on-chain guarantee of solvency) per beneficiary*/
  struct HardDeposit {
    uint amount; /* hard deposit amount allocated */
    uint decreaseAmount; /* decreaseAmount substranced from amount when decrease is requested */
    uint decreaseTimeout; /* issuer has to wait decreaseTimeout seconds after decrease is requested to decrease hardDeposit */
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

  /* issuer of the contract, set at construction */
  address payable public issuer;

  /// @notice constructor, allows setting the issuer (needed for "setup wallet as payment")
  constructor(address payable _issuer, uint defaultHardDepositTimeoutDuration) public payable {
    // DEFAULT_HARDDEPOSIT_TIMOUTE_DURATION will be one day or a whatever non-zero argument given as an argument to the constructor
    DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT = defaultHardDepositTimeoutDuration;
    issuer = _issuer;
    if(msg.value > 0) {
      emit Deposit(msg.sender, msg.value);
    }
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

  /// @notice submit a cheque by the issuer
  /// @param beneficiary the beneficiary of the cheque
  /// @param serial the serial number of the cheque
  /// @param amount the (cumulative) amount of the cheque
  /// @param cashTimeout the check can be cashed cashTimeout seconds in the future
  /// @param beneficiarySig signature of the issuer
  function submitChequeIssuer(address beneficiary, uint serial, uint amount, uint cashTimeout, bytes memory beneficiarySig) public {
    require(msg.sender == issuer, "SimpleSwap: not issuer");
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
  /// @param issuerSig signature of the issuer
  function submitChequeBeneficiary(uint serial, uint amount, uint cashTimeout, bytes memory issuerSig) public {
    /* verify signature of the issuer */
    //emit LogAddress(recover(chequeHash(address(this), msg.sender, serial, amount, cashTimeout), issuerSig));
    require(issuer == recover(chequeHash(address(this), msg.sender, serial, amount, cashTimeout), issuerSig),
     "SimpleSwap: invalid issuerSig");
    /* update the cheque data */
    _submitChequeInternal(msg.sender, serial, amount, cashTimeout);
  }

  /// @notice submit a cheque by any party
  /// @param beneficiary the beneficiary of the cheque
  /// @param serial the serial number of the cheque
  /// @param amount the (cumulative) amount of the cheque
  /// @param cashTimeout the check can be cashed cashTimeout seconds in the future
  /// @param issuerSig signature of the issuer
  /// @param beneficarySig signature of the beneficiary
  function submitCheque(address beneficiary, uint serial, uint amount, uint cashTimeout, bytes memory issuerSig, bytes memory beneficarySig) public {
    /* verify signature of the issuer */
    require(issuer == recover(chequeHash(address(this), beneficiary, serial, amount, cashTimeout), issuerSig),
    "SimpleSwap: invalid issuerSig");
    /* verify signature of the beneficiary */
    require(beneficiary == recover(chequeHash(address(this), beneficiary, serial, amount, cashTimeout), beneficarySig),
    "SimpleSwap: invalid beneficiarySig");
    /* update the cheque data */
    _submitChequeInternal(beneficiary, serial, amount, cashTimeout);
  }


  function _cashChequeInternal(address beneficiaryPrincipal, address payable beneficiaryAgent, uint requestPayout, uint calleePayout) public {
     ChequeInfo storage cheque = cheques[beneficiaryPrincipal];
    /* grace period must have ended */
    require(now >= cheque.cashTimeout,  "SimpleSwap: cheque not yet timed out");
    require(requestPayout <= cheque.amount.sub(cheque.paidOut), "SimpleSwap: not enough balance owed");
    /* ensure there is a balance to claim */
     /* calculates hard-deposit usage */
    uint hardDepositUsage = Math.min(requestPayout, hardDeposits[beneficiaryPrincipal].amount);
    /* calculates acutal payout */
    uint totalPayout = Math.min(requestPayout, liquidBalance() + hardDepositUsage);
      /* if there some of the hard deposit is used update the structure */
    if(hardDepositUsage != 0) {
      hardDeposits[beneficiaryPrincipal].amount = hardDeposits[beneficiaryPrincipal].amount.sub(hardDepositUsage);
      totalHardDeposit = totalHardDeposit.sub(hardDepositUsage);
    }
    /* increase the stored paidOut amount to avoid double payout */
    cheque.paidOut = cheque.paidOut.add(totalPayout);
    /* do the actual payments */
    beneficiaryAgent.transfer(totalPayout.sub(calleePayout));
    emit ChequeCashed(beneficiaryPrincipal, beneficiaryAgent, msg.sender, cheque.serial, totalPayout, requestPayout, calleePayout);
    if(requestPayout != totalPayout) {
      emit ChequeBounced();
    }
  }

  function cashCheque(
    address beneficiaryPrincipal,
    address payable beneficiaryAgent,
    uint requestPayout,
    bytes memory beneficiarySig,
    uint256 expiry,
    uint256 calleePayout
  ) public {
    require(now <= expiry, "SimpleSwap: beneficiarySig expired");
    require(beneficiaryPrincipal == recover(cashOutHash(
      address(this),
      msg.sender,
      requestPayout,
      beneficiaryAgent,
      expiry,
      calleePayout
      ), beneficiarySig), "SimpleSwap: invalid beneficiarySig");
    _cashChequeInternal(beneficiaryPrincipal, beneficiaryAgent, requestPayout, calleePayout);
    msg.sender.transfer(calleePayout);
  }
  /// @notice attempt to cash latest chequebeneficiary
  /// @param beneficiaryAgent agent (of the beneficiary) who receives the payment (i.e. other chequebook contract or the beneficiary)
  /// @param requestPayout amount requested to pay out
  function cashChequeBeneficiary(address payable beneficiaryAgent, uint requestPayout) public {
    _cashChequeInternal(msg.sender, beneficiaryAgent, requestPayout, 0);
  }

  /// @notice prepare to decrease the hard deposit
  /// @param beneficiary beneficiary whose hard deposit should be decreased
  /// @param decreaseAmount amount that the deposit is supposed to be decreased by
  function prepareDecreaseHardDeposit(address beneficiary, uint decreaseAmount) public {
    require(msg.sender == issuer, "SimpleSwap: not issuer");
    HardDeposit storage hardDeposit = hardDeposits[beneficiary];
    /* cannot decrease it by more than the deposit */
    require(decreaseAmount <= hardDeposit.amount, "SimpleSwap: hard deposit not sufficient");
    // if hardDeposit.decreaseTimeout was never set, we DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT. Otherwise we use the one which was set.
    uint decreaseTimeout = hardDeposit.decreaseTimeout == 0 ? DEFAULT_HARDDEPOSIT_DECREASE_TIMEOUT : hardDeposit.decreaseTimeout;
    hardDeposit.canBeDecreasedAt = now + decreaseTimeout;
    hardDeposit.decreaseAmount = decreaseAmount;
    emit HardDepositDecreasePrepared(beneficiary, decreaseAmount);
  }

  /// @notice actually decrease the hard deposit
  /// @param beneficiary beneficiary whose hard deposit should be decreased
  function decreaseHardDeposit(address beneficiary) public {
    HardDeposit storage hardDeposit = hardDeposits[beneficiary];

    require(now >= hardDeposit.canBeDecreasedAt && hardDeposit.canBeDecreasedAt != 0, "SimpleSwap: deposit not yet timed out");

    /* decrease the amount */
    /* this throws if decreaseAmount > amount */
    hardDeposit.amount = hardDeposit.amount.sub(hardDeposit.decreaseAmount);
    /* reset the canBeDecreasedAt to avoid a double decrease */
    hardDeposit.canBeDecreasedAt = 0;
    /* keep totalDeposit in sync */
    totalHardDeposit = totalHardDeposit.sub(hardDeposit.decreaseAmount);

    emit HardDepositAmountChanged(beneficiary, hardDeposit.amount);
  }

  /// @notice increase the hard deposit
  /// @param beneficiary beneficiary whose hard deposit should be decreased
  /// @param amount the new hard deposit
  function increaseHardDeposit(address beneficiary, uint amount) public {
    require(msg.sender == issuer, "SimpleSwap: not issuer");
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
    bytes memory beneficiarySig
  ) public {
    require(msg.sender == issuer, "SimpleSwap: not issuer");
    require(beneficiary == recover(keccak256(abi.encode(address(this), beneficiary, decreaseTimeout)), beneficiarySig));
    hardDeposits[beneficiary].decreaseTimeout = decreaseTimeout;
    emit HardDepositDecreaseTimeoutChanged(beneficiary, hardDeposits[beneficiary].decreaseTimeout);
  }

  /// @notice withdraw ether
  /// @param amount amount to withdraw
  function withdraw(uint amount) public {
    /* only issuer can do this */
    require(msg.sender == issuer, "SimpleSwap: not issuer");
    /* ensure we don't take anything from the hard deposit */
    require(amount <= liquidBalance(), "SimpleSwap: liquidBalance not sufficient");
    issuer.transfer(amount);
  }

  /// @notice deposit ether
  function() payable external {
    if(msg.value > 0) {
      emit Deposit(msg.sender, msg.value);
    }
  }

  function recover(bytes32 hash, bytes memory sig) internal pure returns (address) {
    return ECDSA.recover(ECDSA.toEthSignedMessageHash(hash), sig);
  }

  function chequeHash(address swap, address beneficiary, uint serial, uint amount, uint cashTimeout)
  public pure returns (bytes32) {
    return keccak256(abi.encodePacked(swap, beneficiary, serial, amount, cashTimeout));
  }

  function cashOutHash(address swap, address sender, uint requestPayout, address beneficiaryAgent, uint expiry, uint calleePayout)
  public pure returns (bytes32) {
    return keccak256(abi.encodePacked(swap, sender, requestPayout, beneficiaryAgent, expiry, calleePayout));
  }
}


